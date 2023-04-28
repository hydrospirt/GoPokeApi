# Проект GoPokeApi
## Описание
Простой API Pokemon на языке Go
### Технологии:
- Go 1.20.3
- Gin Web Framework
- Viper v2
- GORM - The fantastic ORM library for Golang
### Как запустить проект
Клонировать репозиторий и перейти в него в командной строке:

```
git@github.com:hydrospirt/GoPokeApi.git
```
Установить модули:
```
go get github.com/spf13/viper
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/postgres
```
Прописать настройки для PostgreSQL:
GoPokeApi/pkg/common/envs/
```
DB_URL=postgres://DB_USER:DB_PASSWORD@DB_HOST:DB_PORT/DB_NAME
```
Запустить проект:
```
go run main.go
```
### Примеры запросов:
Добавление по id

Вывод всех добавленых Pokemon
