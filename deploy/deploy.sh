#!/bin/bash
set -e

echo "开始部署..."

# 确保目录存在
sudo mkdir -p /opt/wyhblog
sudo mkdir -p /var/lib/wyhblog
sudo mkdir -p /etc/wyhblog

# 设置权限
sudo chown -R www-data:www-data /var/lib/wyhblog || sudo chown -R nginx:nginx /var/lib/wyhblog || true

# 复制systemd服务文件
sudo cp /opt/wyhblog/deploy/wyhblog.service /etc/systemd/system/wyhblog.service

# 如果环境变量文件不存在，从示例复制
if [ ! -f /etc/wyhblog/wyhblog.env ]; then
    sudo cp /opt/wyhblog/deploy/wyhblog.env.example /etc/wyhblog/wyhblog.env
    echo "请编辑 /etc/wyhblog/wyhblog.env 设置管理员密码"
fi

# 重载systemd并启动服务
sudo systemctl daemon-reload
sudo systemctl enable wyhblog
sudo systemctl restart wyhblog

echo "部署完成！"
echo "检查服务状态: sudo systemctl status wyhblog"
