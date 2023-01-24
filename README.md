﻿# REST API реализация списка задач
 
 ## Запуск
 1) Создать Docker контейнер БД
 ```sh
 docker run --name=todo-db -e POSTGRES_PASSWORD='root' -p 5436:5432 -d --rm postgres
```
 2) Применить файл миграции
 ```sh
 migrate -path ./schema -database 'postgres://postgres:root@localhost:5436/postgres?sslmode=disable' up
 ```
 3) Скомпилировать и запустить программу
 ```sh
 go run cmd/main.go
 ```
 
 В случае создания контейнера с другим портом необходимо изменить порт в config/config.yml
 
 ## Roadmap
 - [ ] 1. Документация Swagger
 - [ ] 2. Приоритет у задач и подзадач
 - [ ] 3. Время выполнения
