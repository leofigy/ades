go mod tidy

## Regular usage with local db (sqlite)
```
./ades
```
## Using mariadb/sqlite, tell the user/pass and the dbname
```
./ades -db "ades:pass$@tcp(127.0.0.1:3306)/ades?charset=utf8mb4&parseTime=True&loc=Local"
```