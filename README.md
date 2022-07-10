```
cd core
goctl 生成 api
goctl  api go -api core.api -dir . -style gozero

生成模型文件
goctl model pg datasource --url="user=postgres password=postgres host=127.0.0.1 port=5432 dbname=react_demo_backend sslmode=disable" --table="*"  --dir="./model"
go run user.go -f etc/config.yaml
```