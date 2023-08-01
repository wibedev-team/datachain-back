Datachain Админ Панель
# ВХОДНЫЕ ДАННЫЕ
либо логин admin пароль admin, либо логин test пароль test
# Локальный запуск
Запуск проекта
```
    docker compose -f local.docker-compose.yaml up
```
Переходим по ссылке https://localhost:8080/, чтобы удостоверится в работоспособности
```
    https://localhost:8080/
```
Чтобы остановить проект, в той же директории
```
    docker compose -f local.docker-compose.yaml down
```