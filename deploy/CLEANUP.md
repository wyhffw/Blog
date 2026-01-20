# 删除部署的网站

## 快速清理（推荐）

在服务器上运行清理脚本：

```bash
cd /opt/wyhblog  # 或你的项目目录
chmod +x deploy/cleanup.sh
sudo ./deploy/cleanup.sh
```

脚本会交互式地询问你是否删除各个部分。

---

## 手动清理步骤

### 1. 停止并删除 systemd 服务

```bash
# 停止服务
sudo systemctl stop wyhblog

# 禁用服务
sudo systemctl disable wyhblog

# 删除服务文件
sudo rm /etc/systemd/system/wyhblog.service

# 重载 systemd
sudo systemctl daemon-reload
sudo systemctl reset-failed
```

### 2. 删除 Nginx 配置

**Ubuntu/Debian:**
```bash
sudo rm /etc/nginx/sites-enabled/wyhblog
sudo rm /etc/nginx/sites-available/wyhblog
sudo nginx -t
sudo systemctl reload nginx
```

**CentOS/RHEL:**
```bash
sudo rm /etc/nginx/conf.d/wyhblog.conf
sudo nginx -t
sudo systemctl reload nginx
```

**宝塔面板:**
```bash
sudo rm /www/server/panel/vhost/nginx/wyhblog.conf
sudo nginx -s reload
```

### 3. 删除应用文件

```bash
# 删除应用目录（包含二进制、前端文件等）
sudo rm -rf /opt/wyhblog
```

### 4. 删除数据文件（可选）

**警告：这会删除所有文章数据！**

```bash
# 删除文章数据
sudo rm -rf /var/lib/wyhblog
```

### 5. 删除环境变量文件

```bash
sudo rm -rf /etc/wyhblog
```

### 6. 清理防火墙规则（如果需要）

**firewalld (CentOS/RHEL):**
```bash
sudo firewall-cmd --permanent --remove-service=http
sudo firewall-cmd --permanent --remove-service=https
sudo firewall-cmd --reload
```

**ufw (Ubuntu/Debian):**
```bash
sudo ufw delete allow 80/tcp
sudo ufw delete allow 443/tcp
```

### 7. 删除 SSL 证书（如果使用了 Let's Encrypt）

```bash
# 删除证书（certbot）
sudo certbot delete --cert-name wyhblog.xyz

# 或手动删除
sudo rm -rf /etc/letsencrypt/live/wyhblog.xyz
sudo rm -rf /etc/letsencrypt/renewal/wyhblog.xyz.conf
```

---

## 验证清理

检查是否还有残留：

```bash
# 检查服务
systemctl status wyhblog  # 应该显示 "not found"

# 检查进程
ps aux | grep wyhblog  # 应该没有输出

# 检查端口
sudo ss -lntp | grep :8080  # 应该没有输出

# 检查文件
ls -la /opt/wyhblog  # 应该显示 "No such file or directory"
ls -la /var/lib/wyhblog  # 如果删除了数据，应该显示 "No such file or directory"
```

---

## 完整清理命令（一键执行）

**警告：这会删除所有内容，包括数据！**

```bash
# 停止并删除服务
sudo systemctl stop wyhblog 2>/dev/null || true
sudo systemctl disable wyhblog 2>/dev/null || true
sudo rm -f /etc/systemd/system/wyhblog.service
sudo systemctl daemon-reload

# 删除 Nginx 配置
sudo rm -f /etc/nginx/sites-enabled/wyhblog
sudo rm -f /etc/nginx/sites-available/wyhblog
sudo rm -f /etc/nginx/conf.d/wyhblog.conf
sudo rm -f /www/server/panel/vhost/nginx/wyhblog.conf 2>/dev/null || true
sudo nginx -t && sudo systemctl reload nginx 2>/dev/null || sudo nginx -s reload 2>/dev/null || true

# 删除文件
sudo rm -rf /opt/wyhblog
sudo rm -rf /var/lib/wyhblog
sudo rm -rf /etc/wyhblog

echo "清理完成！"
```

---

## 注意事项

1. **备份数据**：删除前请确保已备份重要数据
   ```bash
   # 备份文章数据
   sudo tar -czf wyhblog-backup-$(date +%Y%m%d).tar.gz /var/lib/wyhblog
   ```

2. **保留数据**：如果只想停止服务但保留数据，只执行步骤 1 和 2

3. **域名配置**：删除后记得在 DNS 提供商处删除或修改域名解析

4. **证书清理**：如果使用了 HTTPS，证书文件通常需要单独清理
