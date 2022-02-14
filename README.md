# GOLEARN-AS
Переводы:
* [English](./README_en.md)
* [Русский](./README.md)

## Обзор
Автоматизированная система центра повышения квалификации

## Зависимости 
### Для разработки
* git >= 2.2
* golang >= 1.16
* sqlite3 >= 3.37

Опционально:
* make >= 4.0
* python >= 3.6
* docker >= 20.10
* docker-compose >= 2.2
* postgres >= 12

### Для развёртывания
* docker >= 20.10
* docker-compose >= 2.2
* postgres >= 12

Опционально:
* dpkg >= 1.20
* tar >= 1.30

### Для запуска
* postgres >= 12

## Разработка
```bash
sqlite3 dev.db < config/migrate.sql
cp config/server.toml.bak config/server.toml
cp config/client.toml.bak config/client.toml
make run_server
```

## Теги. Версии. Разработка
v0.0.1
* Инициализация проекта в Git
* Инициализация Docker-compose
* Инициализация сервера
* Структура проекта
* Создание пользователя
* Авторизация через JWT 
* Получение, валидация токена
