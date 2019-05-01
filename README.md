
# gin-gorm-sample

## How To Run

```
docker-compose up
```

or

```
docker-compose run -v $PWD:/app -p 8080:8080 app bash
cd src/gin-sample
go run cmd/app/main.go
```


## Query

```
# INDEX
curl localhost:8080/products

# SHOW
curl localhost:8080/products/1

# CREATE
curl localhost:8080/products -X POST -H "Content-Type: application/json" -d '{"code":"EFEFE", "price": 4000}'

# UPDATE
curl localhost:8080/products/1 -X PUT -H "Content-Type: application/json" -d '{"code":"EFEFE", "price": 4000}'

# DELETE
curl localhost:8080/products/1 -X DELETE
```
