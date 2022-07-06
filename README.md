```
生成模型文件
cd core
goctl model pg datasource --url="user=postgres password=postgres host=127.0.0.1 port=5432 dbname=react_demo_backend sslmode=disable" --table="*"  --dir="./model"
go run user.go -f etc/config.yaml
```