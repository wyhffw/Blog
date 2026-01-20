# 个人博客系统

基于 Go 后端 + HTML 前端的个人博客系统，支持 CI/CD 自动部署。

## 功能特性

- ✅ Go 后端 RESTful API
- ✅ 简约现代的 HTML/CSS/JavaScript 前端
- ✅ 文章 CRUD 操作
- ✅ 用户认证系统（仅登录用户可写文章，访客只能查看）
- ✅ GitHub Actions CI/CD 自动部署
- ✅ 文件存储（JSON格式）

## 项目结构

```
.
├── cmd/
│   └── server/          # Go 服务器主程序
├── web/
│   └── dist/            # 前端静态文件
├── deploy/               # 部署配置文件
│   ├── wyhblog.service   # systemd 服务文件
│   ├── nginx-wyhblog.conf # Nginx 配置
│   └── deploy.sh         # 部署脚本
├── .github/
│   └── workflows/
│       └── deploy.yml    # CI/CD 配置
└── go.mod
```

## 本地开发

### 1. 安装依赖

```bash
go mod download
```

### 2. 运行服务器

```bash
export DATA_DIR=./data/posts
export PUBLIC_DIR=./web/dist
export ADMIN_USER=admin          # 设置你的管理员用户名
export ADMIN_PASS=your_password  # 设置你的管理员密码
go run ./cmd/server
```

访问 http://localhost:8080

**注意**：这是个人博客系统，只有一个管理员账号（你自己）。管理员账号通过环境变量 `ADMIN_USER` 和 `ADMIN_PASS` 配置，不需要注册功能。访客只能查看文章，只有管理员可以写文章。

## 部署到服务器

### 方式一：手动部署

1. 编译二进制：
```bash
go build -o bin/wyhblog ./cmd/server
```

2. 上传到服务器：
```bash
scp -r bin/wyhblog web/dist deploy/* user@your-server:/opt/wyhblog/
```

3. 在服务器上执行部署脚本：
```bash
ssh user@your-server
cd /opt/wyhblog
chmod +x deploy/deploy.sh
sudo ./deploy/deploy.sh
```

4. 配置 Nginx（Ubuntu/Debian）：
```bash
sudo cp /opt/wyhblog/deploy/nginx-wyhblog.conf /etc/nginx/sites-available/wyhblog
sudo ln -s /etc/nginx/sites-available/wyhblog /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

### 方式二：CI/CD 自动部署

1. 在 GitHub 仓库设置 Secrets：
   - `SERVER_HOST`: 服务器 IP 或域名
   - `SERVER_USER`: SSH 用户名
   - `SERVER_SSH_KEY`: SSH 私钥（完整内容）
   - `SERVER_PORT`: SSH 端口（可选，默认22）

2. 推送代码到 main 分支，GitHub Actions 会自动部署

## API 接口

### 公开接口（无需认证）
- `GET /api/health` - 健康检查
- `GET /api/posts` - 获取所有文章
- `GET /api/posts/{id}` - 获取单篇文章

### 需要认证的接口（需在 Header 中添加 `Authorization: Bearer <token>`）
- `POST /api/login` - 登录（返回 token）
- `POST /api/logout` - 登出
- `POST /api/posts` - 创建文章
- `PUT /api/posts/{id}` - 更新文章
- `DELETE /api/posts/{id}` - 删除文章

## 环境变量

- `ADDR`: 服务器监听地址（默认:8080）
- `DATA_DIR`: 文章存储目录（默认:./data/posts）
- `PUBLIC_DIR`: 前端静态文件目录（默认:./web/dist）
- `ADMIN_USER`: 管理员用户名（默认:admin）
- `ADMIN_PASS`: 管理员密码（默认:admin123，**生产环境必须修改！**）

## 许可证

MIT
