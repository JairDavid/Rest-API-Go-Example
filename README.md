
# API Reference

A restful API example using Gorm (ORM) and postgreSQL. The purpose for which I made this API was just for practice
## General Table Views
![Logo](https://imgdb.net/storage/uploads/c63b9da5526062e5825e16a73e0019bbc07416eb64ddfa797317b89b259fadc9.png)

#### Category Enpoints

```bash
  GET /api/v1/category/
  GET /api/v1/category/:id
  POST /api/v1/category/
  PATCH /api/v1/items/:id
  DELETE /api/v1/items/:id
```
#### Customer Enpoints

```bash
  GET /api/v1/customer/
  GET /api/v1/customer/:id
  POST /api/v1/customer/
  PATCH /api/v1/customer/:id
  DELETE /api/v1/customer/:id
```
#### Product Enpoints

```bash
  GET /api/v1/product/
  GET /api/v1/product/:id
  POST /api/v1/product/
  PATCH /api/v1/product/:id
  DELETE /api/v1/product/:id
```

## Configurations

Create a .env file, inside export URL connection where "market" is the database name.
```bash
export DATABASE_URI_DEV="postgres://postgres:root@localhost:5432/market?sslmode=disable"
```
before run migrations, make you sure to create the database on your database engine.

![Logo](https://imgdb.net/storage/uploads/9ce2f9737d291a4aa4f46358fbda20573674c1e452bc5ca8c69043d90cf910bd.png)


#### Migrations
Inside server.go, uncomment the method "StartMigrations()"
```bash
//initial migrations
connection.StartMigrations()
```

Run it once
```bash
go run ./server.go
```
![Logo](https://imgdb.net/storage/uploads/69bb6b48d94a888868f9b10cca4e8793d1bb09d914495645efd783904ccfd5ab.png)

if you do the above settings correctly, everything should be fine, now if you look the database engine you'll see the migrated tables
(Comment again "StartMigrations()" method).

## Tools used
- [Gin Framework)](https://gin-gonic.com/)
- [Gorm (ORM)](https://gorm.io/index.html)
- [Godotenv](https://github.com/joho/godotenv)
- [Postgres driver](https://github.com/lib/pq)
