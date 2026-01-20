# 检查服务器部署状态

## 本地测试（开发环境）

### 1. 检查服务器是否运行

**Windows PowerShell:**
```powershell
# 检查端口 8080 是否被占用
netstat -ano | findstr :8080

# 或者检查 Go 进程
Get-Process | Where-Object {$_.ProcessName -like "*go*"}
```

**或者直接访问:**
```
http://localhost:8080
```

如果能看到博客页面，说明服务器正在运行。

### 2. 检查 API 是否正常

在浏览器访问或使用 curl:
```
http://localhost:8080/api/health
```

应该返回 JSON:
```json
{"ok":true,"time":"2024-01-19T..."}
```

---

## 生产环境（服务器）

### 1. 检查 systemd 服务状态

```bash
# 查看服务状态
sudo systemctl status wyhblog

# 查看服务日志
sudo journalctl -u wyhblog -n 50 --no-pager

# 查看实时日志
sudo journalctl -u wyhblog -f
```

**正常状态应该显示:**
```
Active: active (running)
```

### 2. 检查端口监听

```bash
# 检查 8080 端口是否在监听
sudo ss -lntp | grep :8080
# 或
sudo netstat -tlnp | grep :8080
```

应该看到类似:
```
127.0.0.1:8080  LISTEN  .../wyhblog
```

### 3. 测试 API 端点

```bash
# 在服务器上测试
curl http://127.0.0.1:8080/api/health

# 从外网测试（如果域名已配置）
curl http://your-domain.com/api/health
```

### 4. 检查 Nginx 配置

```bash
# 测试 Nginx 配置
sudo nginx -t

# 检查 Nginx 状态
sudo systemctl status nginx

# 查看 Nginx 错误日志
sudo tail -f /var/log/nginx/error.log
```

### 5. 检查文件是否存在

```bash
# 检查二进制文件
ls -l /opt/wyhblog/bin/wyhblog

# 检查前端文件
ls -l /opt/wyhblog/web/dist/index.html

# 检查数据目录
ls -l /var/lib/wyhblog
```

### 6. 检查进程

```bash
# 查看 wyhblog 进程
ps aux | grep wyhblog

# 查看进程详细信息
sudo systemctl status wyhblog --no-pager -l
```

---

## 常见问题排查

### 服务无法启动

```bash
# 查看详细错误日志
sudo journalctl -u wyhblog -n 100 --no-pager

# 检查配置文件
cat /etc/systemd/system/wyhblog.service

# 检查环境变量文件
cat /etc/wyhblog/wyhblog.env
```

### 端口被占用

```bash
# 查看占用 8080 端口的进程
sudo lsof -i :8080
# 或
sudo ss -lntp | grep :8080
```

### 权限问题

```bash
# 检查文件权限
ls -l /opt/wyhblog/bin/wyhblog
ls -ld /var/lib/wyhblog

# 修复权限（如果需要）
sudo chown -R www-data:www-data /var/lib/wyhblog
sudo chmod +x /opt/wyhblog/bin/wyhblog
```

---

## 快速检查脚本

创建一个检查脚本 `check-deploy.sh`:

```bash
#!/bin/bash
echo "=== 检查博客部署状态 ==="
echo ""

echo "1. 检查服务状态:"
sudo systemctl status wyhblog --no-pager | head -10
echo ""

echo "2. 检查端口监听:"
sudo ss -lntp | grep :8080 || echo "端口 8080 未监听"
echo ""

echo "3. 测试 API:"
curl -s http://127.0.0.1:8080/api/health || echo "API 无响应"
echo ""

echo "4. 检查文件:"
ls -l /opt/wyhblog/bin/wyhblog 2>/dev/null || echo "二进制文件不存在"
ls -l /opt/wyhblog/web/dist/index.html 2>/dev/null || echo "前端文件不存在"
echo ""

echo "5. 检查 Nginx:"
sudo nginx -t 2>&1 | tail -1
echo ""

echo "检查完成！"
```

运行:
```bash
chmod +x check-deploy.sh
./check-deploy.sh
```
