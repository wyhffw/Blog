#!/bin/bash
# 清理博客部署脚本

set -e

echo "=== 开始清理博客部署 ==="
echo ""

# 1. 停止并删除 systemd 服务
echo "1. 停止并删除 systemd 服务..."
if systemctl is-active --quiet wyhblog 2>/dev/null; then
    sudo systemctl stop wyhblog
    echo "   服务已停止"
fi

if systemctl is-enabled --quiet wyhblog 2>/dev/null; then
    sudo systemctl disable wyhblog
    echo "   服务已禁用"
fi

if [ -f /etc/systemd/system/wyhblog.service ]; then
    sudo rm /etc/systemd/system/wyhblog.service
    sudo systemctl daemon-reload
    echo "   服务文件已删除"
fi

# 2. 删除 Nginx 配置
echo ""
echo "2. 删除 Nginx 配置..."

# Ubuntu/Debian
if [ -f /etc/nginx/sites-enabled/wyhblog ]; then
    sudo rm /etc/nginx/sites-enabled/wyhblog
    echo "   已删除 /etc/nginx/sites-enabled/wyhblog"
fi

if [ -f /etc/nginx/sites-available/wyhblog ]; then
    sudo rm /etc/nginx/sites-available/wyhblog
    echo "   已删除 /etc/nginx/sites-available/wyhblog"
fi

# CentOS/RHEL
if [ -f /etc/nginx/conf.d/wyhblog.conf ]; then
    sudo rm /etc/nginx/conf.d/wyhblog.conf
    echo "   已删除 /etc/nginx/conf.d/wyhblog.conf"
fi

# 宝塔面板
if [ -f /www/server/panel/vhost/nginx/wyhblog.conf ]; then
    sudo rm /www/server/panel/vhost/nginx/wyhblog.conf
    echo "   已删除 /www/server/panel/vhost/nginx/wyhblog.conf"
fi

# 重载 Nginx
if command -v nginx &> /dev/null; then
    sudo nginx -t && sudo systemctl reload nginx 2>/dev/null || true
    echo "   Nginx 配置已重载"
fi

# 3. 删除应用文件
echo ""
echo "3. 删除应用文件..."
read -p "   是否删除应用文件 /opt/wyhblog? (y/N): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    if [ -d /opt/wyhblog ]; then
        sudo rm -rf /opt/wyhblog
        echo "   已删除 /opt/wyhblog"
    fi
fi

# 4. 删除数据文件
echo ""
echo "4. 删除数据文件..."
read -p "   是否删除数据文件 /var/lib/wyhblog? (y/N): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    if [ -d /var/lib/wyhblog ]; then
        sudo rm -rf /var/lib/wyhblog
        echo "   已删除 /var/lib/wyhblog"
    fi
fi

# 5. 删除环境变量文件
echo ""
echo "5. 删除环境变量文件..."
if [ -f /etc/wyhblog/wyhblog.env ]; then
    read -p "   是否删除环境变量文件 /etc/wyhblog/wyhblog.env? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        sudo rm -rf /etc/wyhblog
        echo "   已删除 /etc/wyhblog"
    fi
fi

echo ""
echo "=== 清理完成 ==="
echo ""
echo "注意："
echo "- 如果使用了 HTTPS，证书文件通常不会被删除"
echo "- 如果修改了防火墙规则，需要手动清理"
echo "- 检查是否有其他相关文件需要清理"
