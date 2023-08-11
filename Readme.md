Datachain Админ Панель
# ВХОДНЫЕ ДАННЫЕ
либо логин admin пароль admin, либо логин test пароль test

# Локальный запуск, для разработки
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

Запросы для фетча данных
- https://localhost:8000/about/get получить данные для секции about
- https://localhost:8000/stack/all получить изображения для секции stack
- https://localhost:8000/solution/all получить данные для секции solution
- https://localhost:8000/team/get получить каждого члена команды (имя, позиция, картинка) для секции team
- https://localhost:8000/footer/get получить данные для секции footer
