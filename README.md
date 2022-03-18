# GOLEARN-AS
Переводы:
* [English](./README_en.md)
* [Русский](./README.md)

## Обзор
Разработка автоматизированной системы в защищенном исполнении центра повышения квалификации научно-производственного предприятия

## Зависимости 
### Для разработки
* git >= 2.2
* golang >= 1.15
* python >= 3.6
* sqlite3 >= 3.37
* PyQt >= 5.15

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
make pre_run
make run_server
```
Go to http://localhost:5000/ping

### Клиент
[Примеры](./examples/README.md)
## Теги. Версии. Разработка
v0.1.8
* Добавлены функции безопасности в GUI
* Профиль вынесен в отдельный виджет
* Разработана встроенная система аудита и логирования
* В качестве данных для доступа на клиенте хранится временный токен JWT
* Добавлены ссылки на исходный код, поддержку по АС и документацию
* Исправлены ошибки (сервер)
* Произведено тестирование клиента

v0.1.5
* Разработан [GUI](./docs/GUI_README.md) на языке Python3/qt5
* GUI: авторизация, регистрация
* Исправлены ошибки

v0.1.4
* Разработана автоматическая сборка .deb пакета в Jenkins
* Проработаны некоторые связанности
* Добавлены функции
* Управление группами
* Управление заявлениями

v0.1.0
* Golang версия снижена до 1.15 для сборки на OS Debian
* Проработаны модели сущностей БД 
* Проработны связанности сущностей БД (TODO)
* Улучшен функционал клиента

v0.0.3
* Инициализация моделей сущностей
* Анализ связанностей
* Улучшен функционал клиента
* Драйвер управления БД переведен на ORM
* Разделены конфигурационные файлы клиента и сервера

v0.0.1
* Инициализация проекта в Git
* Инициализация Docker-compose
* Инициализация сервера
* Структура проекта
* Создание пользователя
* Авторизация через JWT 
* Получение, валидация токена