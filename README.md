# go_gin_api_demo

### 服務配置
1. Golang Ver : 1.19.1
2. Postgres Ver : 14.1
3. Redis Ver : 6.2

### 使用docker建置環境
```
docker-compose up -d --build
```
進入容器後 執行 go run main.go 可起啟動專案

### 執行migrate
相關文檔: https://github.com/golang-migrate/migrate

```
//生成migrate檔案
migrate create -ext sql -dir migrations/ -format unix create_{name}_table
```

```
//執行migrate
migrate -database postgres://{user}:{password}@{host}:{port}/{dbname}?sslmode=disable -path migrations/ up
```

### swgger api文件
相關文檔: https://github.com/swaggo/swag/blob/master/README_zh-CN.md

```
// 文件產出
swag i -g router/route.go -dir app/front_service,app/base --instanceName front_service -o=docs/front_service
```

可於以下路徑看到swagger文件 \
{localhost}/api/swagger/front/index.html