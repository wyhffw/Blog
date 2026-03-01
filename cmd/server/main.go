package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostStore struct {
	DataDir string
}

type AuthStore struct {
	tokens map[string]time.Time
	mu     sync.RWMutex
}

var authStore = &AuthStore{
	tokens: make(map[string]time.Time),
}

func NewPostStore(dataDir string) *PostStore {
	os.MkdirAll(dataDir, 0755)
	return &PostStore{DataDir: dataDir}
}

func (s *PostStore) GetPost(id string) (*Post, error) {
	filePath := filepath.Join(s.DataDir, id+".json")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var post Post
	if err := json.Unmarshal(data, &post); err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *PostStore) GetAllPosts() ([]Post, error) {
	files, err := os.ReadDir(s.DataDir)
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			id := strings.TrimSuffix(file.Name(), ".json")
			post, err := s.GetPost(id)
			if err != nil {
				continue
			}
			posts = append(posts, *post)
		}
	}
	// 按创建时间倒序
	for i, j := 0, len(posts)-1; i < j; i, j = i+1, j-1 {
		posts[i], posts[j] = posts[j], posts[i]
	}
	return posts, nil
}

func (s *PostStore) SavePost(post *Post) error {
	if post.ID == "" {
		post.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	}
	if post.Slug == "" {
		post.Slug = strings.ToLower(strings.ReplaceAll(post.Title, " ", "-"))
	}
	if post.CreatedAt.IsZero() {
		post.CreatedAt = time.Now()
	}
	post.UpdatedAt = time.Now()

	filePath := filepath.Join(s.DataDir, post.ID+".json")
	data, err := json.MarshalIndent(post, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func (s *PostStore) DeletePost(id string) error {
	filePath := filepath.Join(s.DataDir, id+".json")
	return os.Remove(filePath)
}

func generateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (a *AuthStore) CreateToken() (string, error) {
	token, err := generateToken()
	if err != nil {
		return "", err
	}
	a.mu.Lock()
	a.tokens[token] = time.Now().Add(24 * time.Hour)
	a.mu.Unlock()
	return token, nil
}

func (a *AuthStore) ValidateToken(token string) bool {
	a.mu.RLock()
	defer a.mu.RUnlock()
	expiry, ok := a.tokens[token]
	if !ok {
		return false
	}
	if time.Now().After(expiry) {
		delete(a.tokens, token)
		return false
	}
	return true
}

func (a *AuthStore) DeleteToken(token string) {
	a.mu.Lock()
	delete(a.tokens, token)
	a.mu.Unlock()
}

func requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			token = r.URL.Query().Get("token")
		}
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		if !authStore.ValidateToken(token) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func main() {
	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		dataDir = "./data/posts"
	}

	publicDir := os.Getenv("PUBLIC_DIR")
	if publicDir == "" {
		publicDir = "./web/dist"
	}

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}

	adminUser := os.Getenv("ADMIN_USER")
	if adminUser == "" {
		adminUser = "admin"
	}
	adminPass := os.Getenv("ADMIN_PASS")
	if adminPass == "" {
		adminPass = "admin123"
		log.Printf("WARNING: Using default admin password. Set ADMIN_PASS environment variable!")
	}

	store := NewPostStore(dataDir)
	r := mux.NewRouter()

	// API路由
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"ok":   true,
			"time": time.Now().UTC(),
		})
	}).Methods("GET")

	// 登录接口
	api.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var creds struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if creds.Username != adminUser || creds.Password != adminPass {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		token, err := authStore.CreateToken()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}).Methods("POST")

	// 登出接口
	api.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")
		authStore.DeleteToken(token)
		w.WriteHeader(http.StatusNoContent)
	}).Methods("POST")

	// 获取所有文章（公开）
	api.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		posts, err := store.GetAllPosts()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	}).Methods("GET")

	// 创建文章（需要认证）
	api.HandleFunc("/posts", requireAuth(func(w http.ResponseWriter, r *http.Request) {
		var post Post
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := store.SavePost(&post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	})).Methods("POST")

	// 获取单篇文章（公开）
	api.HandleFunc("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		post, err := store.GetPost(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	}).Methods("GET")

	// 更新文章（需要认证）
	api.HandleFunc("/posts/{id}", requireAuth(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var post Post
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		post.ID = vars["id"]
		if err := store.SavePost(&post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	})).Methods("PUT")

	// 删除文章（需要认证）
	api.HandleFunc("/posts/{id}", requireAuth(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if err := store.DeletePost(vars["id"]); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})).Methods("DELETE")

	// 静态文件服务
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(publicDir)))

	log.Printf("Server starting on %s", addr)
	log.Printf("Data directory: %s", dataDir)
	log.Printf("Public directory: %s", publicDir)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
