# 部署指南

## 前置要求

- Ubuntu/Debian 或 CentOS 服务器
- Go 1.22+ 
- Nginx
- 域名已解析到服务器 IP

## 快速部署

### 1. 服务器准备

```bash
# 安装 Go（如果还没有）
cd /tmp
curl -fsSL -o go1.22.10.linux-amd64.tar.gz https://go.dev/dl/go1.22.10.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.22.10.linux-amd64.tar.gz
echo 'export PATH=/usr/local/go/bin:$PATH' | sudo tee /etc/profile.d/go.sh >/dev/null
source /etc/profile.d/go.sh
go version

# 安装 Nginx（Ubuntu/Debian）
sudo apt update && sudo apt install -y nginx

# 或 CentOS
sudo yum install -y epel-release && sudo yum install -y nginx
```

### 2. 上传代码到服务器

```bash
# 方式1: 使用 git
git clone <your-repo-url> /opt/wyhblog

# 方式2: 使用 scp
scp -r bin/wyhblog web/dist deploy/* user@server:/opt/wyhblog/
```

### 3. 编译（如果还没编译）

```bash
cd /opt/wyhblog
go build -o bin/wyhblog ./cmd/server
chmod +x bin/wyhblog
```

### 4. 配置 systemd 服务

```bash
sudo mkdir -p /var/lib/wyhblog /etc/wyhblog
sudo cp deploy/wyhblog.service /etc/systemd/system/
sudo cp deploy/wyhblog.env.example /etc/wyhblog/wyhblog.env

# 编辑管理员账号密码（必须！）
sudo nano /etc/wyhblog/wyhblog.env
```

**重要**：这是个人博客系统，只有一个管理员账号。在 `wyhblog.env` 文件中设置：
- `ADMIN_USER`: 你的管理员用户名（默认: admin）
- `ADMIN_PASS`: 你的管理员密码（**必须修改默认值！**）

这个账号用于登录后写文章。访客只能查看文章，无法注册新账号。

# 启动服务
sudo systemctl daemon-reload
sudo systemctl enable --now wyhblog
sudo systemctl status wyhblog
```

### 5. 配置 Nginx

#### Ubuntu/Debian（标准安装）

```bash
sudo cp deploy/nginx-wyhblog.conf /etc/nginx/sites-available/wyhblog
sudo ln -s /etc/nginx/sites-available/wyhblog /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

#### CentOS 或宝塔面板

```bash
# 宝塔面板路径
sudo cp deploy/nginx-wyhblog.conf /www/server/panel/vhost/nginx/wyhblog.conf

# 或 CentOS 标准路径
sudo cp deploy/nginx-wyhblog.conf /etc/nginx/conf.d/wyhblog.conf

sudo nginx -t
sudo systemctl reload nginx
```

**重要**: 修改配置文件中的 `server_name` 为你的域名！

### 6. 配置 HTTPS（推荐）

使用 certbot（Let's Encrypt）：

```bash
sudo apt install -y certbot python3-certbot-nginx  # Ubuntu
# 或
sudo yum install -y certbot python3-certbot-nginx  # CentOS

sudo certbot --nginx -d wyhblog.xyz -d www.wyhblog.xyz
```

### 7. 防火墙配置

```bash
# Ubuntu/Debian (ufw)
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp

# CentOS (firewalld)
sudo firewall-cmd --permanent --add-service=http
sudo firewall-cmd --permanent --add-service=https
sudo firewall-cmd --reload
```

## CI/CD 配置

### GitHub Actions 设置

1. 在 GitHub 仓库 Settings > Secrets and variables > Actions 中添加：

   - `SERVER_HOST`: 服务器 IP 或域名
   - `SERVER_USER`: SSH 用户名（如 root）
   - `SERVER_SSH_KEY`: SSH 私钥完整内容
   - `SERVER_PORT`: SSH 端口（可选，默认22）

2. 生成 SSH 密钥对（如果还没有）：

```bash
ssh-keygen -t rsa -b 4096 -C "github-actions"
# 将公钥添加到服务器 ~/.ssh/authorized_keys
cat ~/.ssh/id_rsa.pub | ssh user@server "mkdir -p ~/.ssh && cat >> ~/.ssh/authorized_keys"
```

3. 推送代码到 main 分支，GitHub Actions 会自动部署

## 验证部署

```bash
# 检查服务状态
sudo systemctl status wyhblog

# 检查 API
curl http://127.0.0.1:8080/api/health
curl http://your-domain.com/api/health

# 查看日志
sudo journalctl -u wyhblog -f
```

## 故障排查

### 服务无法启动

```bash
sudo journalctl -u wyhblog -n 50 --no-pager
```

### Nginx 404

- 检查配置文件是否正确加载：`nginx -T | grep wyhblog`
- 检查 server_name 是否匹配域名
- 检查防火墙是否开放 80/443

### 权限问题

```bash
sudo chown -R www-data:www-data /var/lib/wyhblog
sudo chown -R www-data:www-data /opt/wyhblog
```
