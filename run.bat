@echo off
echo Starting blog server...
echo.
echo Default login credentials:
echo   Username: admin
echo   Password: admin123
echo.
echo You can change these by setting ADMIN_USER and ADMIN_PASS environment variables
echo.
set DATA_DIR=./data/posts
set PUBLIC_DIR=./web/dist
set ADMIN_USER=admin
set ADMIN_PASS=admin123
go run ./cmd/server
