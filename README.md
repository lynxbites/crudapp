# crudapp
Репозиторий тестового задания для SkillsRock.

Для контейнеризации сервера используется Docker.
Для миграций PostgreSQL используется [migrate](https://github.com/golang-migrate/migrate).

Документация API сгенерирована с помощью [swag](https://github.com/swaggo/swag).
Swagger можно найти по пути http://localhost:8000/swagger.

Спасибо что уделили внимание! c:

## Usage
Указать строку подключения к PostgreSQL в файле .env (Пример - .env.example):

    CONNSTR=postgres://username:password@localhost:5432/database

Мигрировать базу данных:

    migrate -database "postgres://username:password@localhost:5432/database?sslmode=disable" -path internal/db/migrations up
    
Собрать и запустить Docker контейнер:

    make 

