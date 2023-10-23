## Оглавление

* [Подготовка](#подготовка)
* [Запуск](#запуск)
* [Инициализация БД](#инициализация-бд)
* [API](#api)


### Подготовка

Перед запуском [установите docker](https://docs.docker.com/engine/install/).

### Запуск

Для запуска используйте команду:
```
make up
```
Она поднимет контейнеры с приложением и необходимым для него экземпляром postgres.

### Инициализация БД

Для поднятия миграций, используйте команду:
```
make goose-up
```

### API

Чтобы открыть swagger, перейдите по ссылке:
http://localhost:8080/swagger/index.html