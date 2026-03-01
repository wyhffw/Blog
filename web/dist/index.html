<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Blog · 极简笔记</title>
    
    <!-- 字体和图标 -->
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600&family=JetBrains+Mono:wght@400;500&display=swap" rel="stylesheet">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css">
    
    <!-- Markdown 渲染库 -->
    <script src="https://cdn.jsdelivr.net/npm/marked/marked.min.js"></script>
    <!-- 代码高亮 -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.8.0/styles/github.min.css">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.8.0/highlight.min.js"></script>
    
    <style>
        /* ===== 设计变量 ===== */
        :root {
            --bg: #fafafa;
            --surface: #ffffff;
            --text: #1a1a1a;
            --text-secondary: #4a4a4a;
            --text-tertiary: #8a8a8a;
            --border: #eaeaea;
            --accent: #000000;
            --accent-hover: #333333;
            --selection: #f0f0f0;
            --shadow: 0 10px 30px -15px rgba(0,0,0,0.1);
            --radius: 12px;
            --radius-sm: 6px;
            --font-sans: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
            --font-mono: 'JetBrains Mono', 'SF Mono', monospace;
        }
        
        /* 深色模式 */
        @media (prefers-color-scheme: dark) {
            :root {
                --bg: #0a0a0a;
                --surface: #141414;
                --text: #ededed;
                --text-secondary: #a0a0a0;
                --text-tertiary: #6a6a6a;
                --border: #2a2a2a;
                --accent: #ffffff;
                --accent-hover: #cccccc;
                --selection: #2a2a2a;
                --shadow: 0 10px 30px -15px rgba(0,0,0,0.3);
            }
        }
        
        /* ===== 基础重置 ===== */
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: var(--font-sans);
            background: var(--bg);
            color: var(--text);
            line-height: 1.7;
            font-size: 16px;
            -webkit-font-smoothing: antialiased;
            transition: background-color 0.3s, color 0.3s;
        }
        
        /* ===== 布局 ===== */
        .container {
            max-width: 720px;
            margin: 0 auto;
            padding: 0 24px;
        }
        
        /* ===== 导航栏 ===== */
        .site-header {
            padding: 32px 0 48px;
            border-bottom: 1px solid var(--border);
            margin-bottom: 48px;
        }
        
        .nav {
            display: flex;
            justify-content: space-between;
            align-items: center;
        }
        
        .logo {
            font-size: 20px;
            font-weight: 500;
            letter-spacing: -0.02em;
            text-decoration: none;
            color: var(--text);
        }
        
        .logo span {
            color: var(--text-tertiary);
            font-weight: 300;
        }
        
        .nav-links {
            display: flex;
            gap: 24px;
            align-items: center;
        }
        
        .nav-btn {
            background: none;
            border: none;
            color: var(--text-secondary);
            font-size: 14px;
            cursor: pointer;
            padding: 8px 12px;
            border-radius: var(--radius-sm);
            transition: all 0.2s;
            font-family: var(--font-sans);
        }
        
        .nav-btn:hover {
            color: var(--text);
            background: var(--border);
        }
        
        .nav-btn.primary {
            background: var(--accent);
            color: var(--bg);
        }
        
        .nav-btn.primary:hover {
            background: var(--accent-hover);
            opacity: 1;
        }
        
        /* ===== 文章列表 ===== */
        .posts-list {
            display: flex;
            flex-direction: column;
            gap: 64px;
        }
        
        .post-item {
            cursor: pointer;
            transition: all 0.2s;
            border-bottom: 1px solid var(--border);
            padding-bottom: 48px;
        }
        
        .post-item:last-child {
            border-bottom: none;
            padding-bottom: 0;
        }
        
        .post-meta {
            display: flex;
            align-items: center;
            gap: 16px;
            margin-bottom: 16px;
            font-size: 13px;
            color: var(--text-tertiary);
            text-transform: uppercase;
            letter-spacing: 0.5px;
        }
        
        .post-date {
            position: relative;
        }
        
        .post-category {
            background: var(--selection);
            padding: 4px 10px;
            border-radius: 20px;
            color: var(--text-secondary);
            font-size: 12px;
        }
        
        .post-title {
            font-size: 32px;
            font-weight: 500;
            line-height: 1.3;
            margin-bottom: 20px;
            letter-spacing: -0.02em;
            color: var(--text);
        }
        
        .post-excerpt {
            color: var(--text-secondary);
            font-size: 16px;
            line-height: 1.7;
            margin-bottom: 24px;
        }
        
        .post-footer {
            display: flex;
            align-items: center;
            gap: 24px;
        }
        
        .read-more {
            font-size: 14px;
            color: var(--text);
            text-decoration: none;
            display: inline-flex;
            align-items: center;
            gap: 6px;
            font-weight: 500;
        }
        
        .read-more i {
            font-size: 12px;
            transition: transform 0.2s;
        }
        
        .read-more:hover i {
            transform: translateX(4px);
        }
        
        /* ===== 文章详情页 ===== */
        .post-detail {
            max-width: 720px;
            margin: 0 auto;
        }
        
        .post-detail-header {
            margin-bottom: 48px;
            text-align: left;
        }
        
        .post-detail-title {
            font-size: 42px;
            font-weight: 500;
            line-height: 1.2;
            margin: 16px 0 24px;
            letter-spacing: -0.02em;
        }
        
        .post-detail-meta {
            display: flex;
            justify-content: space-between;
            align-items: center;
            color: var(--text-tertiary);
            font-size: 14px;
            padding: 24px 0;
            border-top: 1px solid var(--border);
            border-bottom: 1px solid var(--border);
        }
        
        .post-detail-content {
            font-size: 17px;
            line-height: 1.8;
            color: var(--text);
            margin-top: 48px;
        }
        
        /* Markdown 样式 */
        .post-detail-content h1 {
            font-size: 32px;
            margin: 40px 0 16px;
            font-weight: 500;
        }
        
        .post-detail-content h2 {
            font-size: 26px;
            margin: 36px 0 14px;
            font-weight: 500;
        }
        
        .post-detail-content h3 {
            font-size: 22px;
            margin: 28px 0 12px;
            font-weight: 500;
        }
        
        .post-detail-content p {
            margin-bottom: 24px;
        }
        
        .post-detail-content code {
            font-family: var(--font-mono);
            font-size: 14px;
            background: var(--selection);
            padding: 3px 6px;
            border-radius: 4px;
            border: 1px solid var(--border);
        }
        
        .post-detail-content pre {
            background: var(--surface);
            border: 1px solid var(--border);
            border-radius: var(--radius);
            padding: 20px;
            overflow-x: auto;
            margin: 32px 0;
        }
        
        .post-detail-content pre code {
            background: none;
            border: none;
            padding: 0;
            font-size: 14px;
            line-height: 1.6;
        }
        
        .post-detail-content blockquote {
            margin: 32px 0;
            padding: 20px 24px;
            border-left: 4px solid var(--accent);
            background: var(--selection);
            font-style: italic;
            border-radius: 0 var(--radius) var(--radius) 0;
        }
        
        .post-detail-content ul, 
        .post-detail-content ol {
            margin: 24px 0;
            padding-left: 24px;
        }
        
        .post-detail-content li {
            margin-bottom: 8px;
        }
        
        .post-detail-content img {
            max-width: 100%;
            border-radius: var(--radius);
            margin: 32px 0;
            border: 1px solid var(--border);
        }
        
        /* ===== 空状态 ===== */
        .empty-state {
            text-align: center;
            padding: 80px 0;
        }
        
        .empty-icon {
            font-size: 48px;
            color: var(--border);
            margin-bottom: 24px;
        }
        
        .empty-state h2 {
            font-size: 20px;
            font-weight: 400;
            color: var(--text-secondary);
            margin-bottom: 8px;
        }
        
        .empty-state p {
            color: var(--text-tertiary);
            font-size: 14px;
        }
        
        /* ===== 模态框 ===== */
        .modal {
            display: none;
            position: fixed;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
            background: rgba(0, 0, 0, 0.6);
            backdrop-filter: blur(8px);
            z-index: 1000;
            align-items: center;
            justify-content: center;
            padding: 20px;
        }
        
        .modal.active {
            display: flex;
        }
        
        .modal-content {
            background: var(--surface);
            border-radius: var(--radius);
            width: 100%;
            max-width: 600px;
            max-height: 90vh;
            overflow-y: auto;
            padding: 40px;
            box-shadow: var(--shadow);
            border: 1px solid var(--border);
            animation: modalSlide 0.3s ease;
        }
        
        @keyframes modalSlide {
            from {
                opacity: 0;
                transform: translateY(20px);
            }
            to {
                opacity: 1;
                transform: translateY(0);
            }
        }
        
        .modal-header {
            margin-bottom: 32px;
        }
        
        .modal-title {
            font-size: 24px;
            font-weight: 500;
            margin-bottom: 8px;
        }
        
        /* ===== 表单 ===== */
        .form-group {
            margin-bottom: 28px;
        }
        
        label {
            display: block;
            font-size: 13px;
            font-weight: 500;
            color: var(--text-secondary);
            margin-bottom: 10px;
            text-transform: uppercase;
            letter-spacing: 0.5px;
        }
        
        input, textarea {
            width: 100%;
            padding: 12px 0;
            border: none;
            border-bottom: 2px solid var(--border);
            font-size: 16px;
            font-family: var(--font-sans);
            color: var(--text);
            background: transparent;
            transition: border-color 0.2s;
        }
        
        input:focus, textarea:focus {
            outline: none;
            border-bottom-color: var(--accent);
        }
        
        textarea {
            min-height: 300px;
            resize: vertical;
            font-family: var(--font-mono);
            font-size: 14px;
            line-height: 1.7;
        }
        
        .form-actions {
            display: flex;
            gap: 12px;
            margin-top: 32px;
            padding-top: 24px;
            border-top: 1px solid var(--border);
        }
        
        /* ===== 按钮 ===== */
        .btn {
            padding: 12px 24px;
            border: 1px solid var(--border);
            background: var(--surface);
            color: var(--text);
            font-size: 14px;
            cursor: pointer;
            border-radius: var(--radius-sm);
            transition: all 0.2s;
            font-weight: 500;
            font-family: var(--font-sans);
        }
        
        .btn:hover {
            background: var(--border);
        }
        
        .btn-primary {
            background: var(--accent);
            color: var(--bg);
            border-color: var(--accent);
        }
        
        .btn-primary:hover {
            background: var(--accent-hover);
            border-color: var(--accent-hover);
        }
        
        .btn-danger {
            border-color: #ff4444;
            color: #ff4444;
        }
        
        .btn-danger:hover {
            background: #ff4444;
            color: white;
        }
        
        /* ===== 错误提示 ===== */
        .error {
            background: rgba(255, 68, 68, 0.1);
            color: #ff4444;
            padding: 12px 16px;
            border-radius: var(--radius-sm);
            margin-bottom: 20px;
            font-size: 14px;
            border: 1px solid rgba(255, 68, 68, 0.3);
        }
        
        /* ===== 加载状态 ===== */
        .loading {
            text-align: center;
            padding: 60px 0;
            color: var(--text-tertiary);
        }
        
        .loading::after {
            content: '';
            display: inline-block;
            width: 16px;
            height: 16px;
            border: 2px solid var(--border);
            border-top-color: var(--accent);
            border-radius: 50%;
            margin-left: 8px;
            animation: spin 0.8s linear infinite;
            vertical-align: middle;
        }
        
        @keyframes spin {
            to { transform: rotate(360deg); }
        }
        
        /* ===== 返回按钮 ===== */
        .back-btn {
            background: none;
            border: none;
            color: var(--text-tertiary);
            cursor: pointer;
            font-size: 14px;
            padding: 8px 0;
            margin-bottom: 24px;
            display: inline-flex;
            align-items: center;
            gap: 8px;
            transition: color 0.2s;
        }
        
        .back-btn:hover {
            color: var(--text);
        }
        
        /* ===== 响应式 ===== */
        @media (max-width: 768px) {
            .container {
                padding: 0 20px;
            }
            
            .post-title {
                font-size: 26px;
            }
            
            .post-detail-title {
                font-size: 32px;
            }
            
            .modal-content {
                padding: 24px;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <header class="site-header">
            <nav class="nav">
                <a href="/" class="logo" id="homeLink">
                    Blog <span>· 极简笔记</span>
                </a>
                <div class="nav-links">
                    <button class="nav-btn" id="loginBtn" onclick="showLoginModal()">
                        <i class="fa-regular fa-user" style="margin-right: 6px;"></i>登录
                    </button>
                    <button class="nav-btn" id="logoutBtn" onclick="logout()" style="display: none;">
                        <i class="fa-regular fa-sign-out" style="margin-right: 6px;"></i>登出
                    </button>
                    <button class="nav-btn primary" id="writeBtn" onclick="showCreateModal()" style="display: none;">
                        <i class="fa-regular fa-pen-to-square" style="margin-right: 6px;"></i>写文章
                    </button>
                </div>
            </nav>
        </header>

        <main id="app">
            <!-- 动态内容 -->
        </main>
    </div>

    <!-- 登录模态框 -->
    <div id="loginModal" class="modal">
        <div class="modal-content">
            <div class="modal-header">
                <h2 class="modal-title">管理员登录</h2>
                <p style="font-size: 14px; color: var(--text-tertiary); line-height: 1.6; margin-top: 8px;">
                    登录后可以发布、编辑和删除文章
                </p>
            </div>
            <div id="loginError" class="error" style="display: none;"></div>
            <form id="loginForm" onsubmit="login(event)">
                <div class="form-group">
                    <label>用户名</label>
                    <input type="text" id="loginUsername" required autocomplete="username" placeholder="输入用户名">
                </div>
                <div class="form-group">
                    <label>密码</label>
                    <input type="password" id="loginPassword" required autocomplete="current-password" placeholder="输入密码">
                </div>
                <div class="form-actions">
                    <button type="submit" class="btn btn-primary">
                        <i class="fa-regular fa-arrow-right-to-bracket" style="margin-right: 6px;"></i>登录
                    </button>
                    <button type="button" class="btn" onclick="closeLoginModal()">取消</button>
                </div>
            </form>
        </div>
    </div>

    <!-- 编辑文章模态框 -->
    <div id="editModal" class="modal">
        <div class="modal-content">
            <div class="modal-header">
                <h2 class="modal-title" id="modalTitle">写新文章</h2>
            </div>
            <form id="postForm" onsubmit="savePost(event)">
                <div class="form-group">
                    <label>标题</label>
                    <input type="text" id="postTitle" required placeholder="给你的文章起个标题">
                </div>
                <div class="form-group">
                    <label>内容 (支持 Markdown)</label>
                    <textarea id="postContent" required placeholder="用 Markdown 写作..."></textarea>
                </div>
                <div class="form-actions">
                    <button type="submit" class="btn btn-primary">
                        <i class="fa-regular fa-floppy-disk" style="margin-right: 6px;"></i>保存
                    </button>
                    <button type="button" class="btn" onclick="closeModal()">取消</button>
                    <button type="button" class="btn btn-danger" id="deleteBtn" onclick="deletePost()" style="display:none; margin-left: auto;">
                        <i class="fa-regular fa-trash-can"></i>
                    </button>
                </div>
            </form>
        </div>
    </div>

    <script>
        // 配置 marked 和 highlight.js
        marked.setOptions({
            highlight: function(code, lang) {
                if (lang && hljs.getLanguage(lang)) {
                    try {
                        return hljs.highlight(code, { language: lang }).value;
                    } catch (err) {}
                }
                return hljs.highlightAuto(code).value;
            },
            breaks: true,
            gfm: true
        });

        let currentPostId = null;
        let authToken = localStorage.getItem('authToken');
        let currentView = 'list'; // 'list' 或 'post'

        // 检查登录状态
        function checkAuth() {
            const loginBtn = document.getElementById('loginBtn');
            const logoutBtn = document.getElementById('logoutBtn');
            const writeBtn = document.getElementById('writeBtn');
            
            if (authToken) {
                loginBtn.style.display = 'none';
                logoutBtn.style.display = 'inline-block';
                writeBtn.style.display = 'inline-block';
            } else {
                loginBtn.style.display = 'inline-block';
                logoutBtn.style.display = 'none';
                writeBtn.style.display = 'none';
            }
        }

        // 登录
        async function login(event) {
            event.preventDefault();
            const username = document.getElementById('loginUsername').value;
            const password = document.getElementById('loginPassword').value;
            const errorDiv = document.getElementById('loginError');

            errorDiv.style.display = 'none';
            errorDiv.textContent = '';

            try {
                const res = await fetch('/api/login', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ username, password })
                });

                if (!res.ok) {
                    let errorMsg = '登录失败';
                    try {
                        const errorText = await res.text();
                        if (errorText) errorMsg = errorText;
                    } catch (e) {}
                    errorDiv.textContent = errorMsg;
                    errorDiv.style.display = 'block';
                    return;
                }

                const data = await res.json();
                if (!data.token) {
                    errorDiv.textContent = '登录失败: 服务器未返回token';
                    errorDiv.style.display = 'block';
                    return;
                }

                authToken = data.token;
                localStorage.setItem('authToken', authToken);
                closeLoginModal();
                checkAuth();
                
                if (currentView === 'list') {
                    loadPosts();
                }
            } catch (error) {
                errorDiv.textContent = '登录失败: ' + error.message;
                errorDiv.style.display = 'block';
            }
        }

        // 登出
        async function logout() {
            if (authToken) {
                try {
                    await fetch('/api/logout', {
                        method: 'POST',
                        headers: { 'Authorization': 'Bearer ' + authToken }
                    });
                } catch (e) {}
            }
            authToken = null;
            localStorage.removeItem('authToken');
            checkAuth();
            
            if (currentView === 'list') {
                loadPosts();
            } else {
                showHome();
            }
        }

        // 显示登录模态框
        function showLoginModal() {
            document.getElementById('loginModal').classList.add('active');
            document.getElementById('loginError').style.display = 'none';
            document.getElementById('loginForm').reset();
        }

        // 关闭登录模态框
        function closeLoginModal() {
            document.getElementById('loginModal').classList.remove('active');
        }

        // 显示写文章模态框
        function showCreateModal() {
            if (!authToken) {
                showLoginModal();
                return;
            }
            currentPostId = null;
            document.getElementById('modalTitle').textContent = '写新文章';
            document.getElementById('postTitle').value = '';
            document.getElementById('postContent').value = '';
            document.getElementById('deleteBtn').style.display = 'none';
            document.getElementById('editModal').classList.add('active');
        }

        // 加载文章列表
        async function loadPosts() {
            currentView = 'list';
            const app = document.getElementById('app');
            
            try {
                const res = await fetch('/api/posts');
                const posts = await res.json();
                
                if (posts.length === 0) {
                    app.innerHTML = `
                        <div class="empty-state">
                            <div class="empty-icon"><i class="fa-regular fa-pen-to-square"></i></div>
                            <h2>还没有文章</h2>
                            <p>${authToken ? '点击"写文章"开始你的第一篇博客' : '登录后可以发布文章'}</p>
                        </div>
                    `;
                    return;
                }

                app.innerHTML = `
                    <div class="posts-list">
                        ${posts.map(post => {
                            const date = new Date(post.created_at);
                            const dateStr = date.toLocaleDateString('zh-CN', { 
                                year: 'numeric', 
                                month: 'long', 
                                day: 'numeric' 
                            });
                            const excerpt = post.content.substring(0, 200).replace(/\n/g, ' ');
                            
                            return `
                                <article class="post-item" onclick="viewPost('${post.id}')">
                                    <div class="post-meta">
                                        <span class="post-date"><i class="fa-regular fa-calendar" style="margin-right: 6px;"></i>${dateStr}</span>
                                    </div>
                                    <h2 class="post-title">${escapeHtml(post.title)}</h2>
                                    <div class="post-excerpt">${escapeHtml(excerpt)}${post.content.length > 200 ? '...' : ''}</div>
                                    <div class="post-footer">
                                        <span class="read-more">
                                            阅读全文 <i class="fa-regular fa-arrow-right"></i>
                                        </span>
                                    </div>
                                </article>
                            `;
                        }).join('')}
                    </div>
                `;
            } catch (error) {
                app.innerHTML = `<div class="error">加载失败: ${error.message}</div>`;
            }
        }

        // 查看单篇文章
        async function viewPost(id) {
            currentView = 'post';
            const app = document.getElementById('app');
            
            try {
                const res = await fetch(`/api/posts/${id}`);
                const post = await res.json();
                
                const date = new Date(post.created_at);
                const dateStr = date.toLocaleDateString('zh-CN', { 
                    year: 'numeric', 
                    month: 'long', 
                    day: 'numeric',
                    hour: '2-digit',
                    minute: '2-digit'
                });

                const html = marked.parse(post.content);

                app.innerHTML = `
                    <div class="post-detail">
                        <button class="back-btn" onclick="loadPosts()">
                            <i class="fa-regular fa-arrow-left"></i> 返回列表
                        </button>
                        <article>
                            <header class="post-detail-header">
                                <h1 class="post-detail-title">${escapeHtml(post.title)}</h1>
                                <div class="post-detail-meta">
                                    <span><i class="fa-regular fa-calendar" style="margin-right: 6px;"></i>${dateStr}</span>
                                    ${authToken ? `
                                        <button class="nav-btn" onclick="editPost('${post.id}')">
                                            <i class="fa-regular fa-pen-to-square"></i> 编辑
                                        </button>
                                    ` : ''}
                                </div>
                            </header>
                            <div class="post-detail-content">
                                ${html}
                            </div>
                        </article>
                    </div>
                `;

                // 代码高亮
                document.querySelectorAll('pre code').forEach(block => {
                    hljs.highlightElement(block);
                });

            } catch (error) {
                app.innerHTML = `<div class="error">加载失败: ${error.message}</div>`;
            }
        }

        // 编辑文章
        function editPost(id) {
            if (!authToken) {
                showLoginModal();
                return;
            }
            viewPostForEdit(id);
        }

        // 加载文章到编辑框
        async function viewPostForEdit(id) {
            try {
                const res = await fetch(`/api/posts/${id}`);
                const post = await res.json();
                currentPostId = post.id;
                document.getElementById('modalTitle').textContent = '编辑文章';
                document.getElementById('postTitle').value = post.title;
                document.getElementById('postContent').value = post.content;
                document.getElementById('deleteBtn').style.display = 'inline-block';
                document.getElementById('editModal').classList.add('active');
            } catch (error) {
                alert('加载失败: ' + error.message);
            }
        }

        // 保存文章
        async function savePost(event) {
            event.preventDefault();
            if (!authToken) {
                showLoginModal();
                return;
            }

            const title = document.getElementById('postTitle').value;
            const content = document.getElementById('postContent').value;

            try {
                const url = currentPostId ? `/api/posts/${currentPostId}` : '/api/posts';
                const method = currentPostId ? 'PUT' : 'POST';
                
                const res = await fetch(url, {
                    method: method,
                    headers: { 
                        'Content-Type': 'application/json',
                        'Authorization': 'Bearer ' + authToken
                    },
                    body: JSON.stringify({ title, content })
                });

                if (!res.ok) {
                    if (res.status === 401) {
                        logout();
                        showLoginModal();
                        return;
                    }
                    throw new Error('保存失败');
                }
                
                closeModal();
                
                if (currentPostId) {
                    viewPost(currentPostId);
                } else {
                    loadPosts();
                }
            } catch (error) {
                alert('保存失败: ' + error.message);
            }
        }

        // 删除文章
        async function deletePost() {
            if (!authToken) return;
            if (!confirm('确定要删除这篇文章吗？')) return;
            
            try {
                const res = await fetch(`/api/posts/${currentPostId}`, { 
                    method: 'DELETE',
                    headers: { 'Authorization': 'Bearer ' + authToken }
                });
                if (!res.ok) {
                    if (res.status === 401) {
                        logout();
                        showLoginModal();
                        return;
                    }
                    throw new Error('删除失败');
                }
                closeModal();
                loadPosts();
            } catch (error) {
                alert('删除失败: ' + error.message);
            }
        }

        // 返回首页
        function showHome() {
            loadPosts();
        }

        // 关闭模态框
        function closeModal() {
            document.getElementById('editModal').classList.remove('active');
        }

        // HTML转义
        function escapeHtml(text) {
            const div = document.createElement('div');
            div.textContent = text;
            return div.innerHTML;
        }

        // 点击模态框外部关闭
        document.getElementById('editModal').addEventListener('click', function(e) {
            if (e.target === this) closeModal();
        });

        document.getElementById('loginModal').addEventListener('click', function(e) {
            if (e.target === this) closeLoginModal();
        });

        // 首页链接
        document.getElementById('homeLink').addEventListener('click', function(e) {
            e.preventDefault();
            loadPosts();
        });

        // 页面加载
        checkAuth();
        loadPosts();
    </script>
</body>
</html>
