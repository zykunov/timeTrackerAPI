# REST API сервиса учета рабочих часов
Для роутинга используется framework GIN, для хранения данных Postgrsql c GORM.

**доступ к документации swagger** - http://localhost:8080/docs/index.html

перед запуском:
>swag init -d .\cmd\,.\handlers\,.\helpers\,.\models\,.\routers\

### Структура проекта

**/cmd/main.go** - точка входа

**/handlers/handlers.go** - хэндлеры 

**/models** - модели и репозитории для GORM 

**/routers** - роутеры

**/storage** - прокладка для общения с БД


