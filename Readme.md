Datachain Админ Панель
# ВХОДНЫЕ ДАННЫЕ
1) логин admin пароль admin
2) логин test пароль test
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
# Запуск проекта на удаленном сервере для продакшена
1) Клонируем проект на удаленный сервер, если его еще нет
```
    git clone https://github.com/wibedev-team/datachain-back.git datachain-admin
    cd datachain-admin
```
2) Запускаем проект для продакшена
```
    docker compose up
```
3) Проверяем проект
   Переходим по ссылке https://188.225.44.3:8080/, чтобы удостоверится в работоспособности
   Если не работает, то по этой http://188.225.44.3:8080/
```
    https://188.225.44.3:8080
```

