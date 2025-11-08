# Todo-app-plan

## Metadata

|Metadata |             |
|---      |---          |
| Author  | AnjeZenda   |
| Created | 07.11.2025  |
| Changed | 08.11.2025  |
| Tags    | arc, plan   |


## Backend

### First stage

Реализация простого сервиса с CRUD-операциями по упралению задачами:

- `POST /tasks` создание задач
- `GET /tasks` получение задач (при подгрузке информации)
- `UPDATE /tasks/{task_id}` обновление состояния задачи, либо изменение задачи в случае ошибок
- `DELETE /tasks/{task_id}` удаление задачи

Реализовать сервис статистики который будет отображать текущую статистику по задачам

- `POST /statistics`, с передачей фильтров в body

Реализовать сервис аунтификации и авторизации:

- `POST /register` - создание учетной записи
- `POST /login` - вход в аккаунт

---

В рамках текущей стадии предлагается придерживаться следующей иерархии

```
service
    cmd
        main.go
    internal
        adapter
        cases
        entities
        port
    pkg
        application
```


## Mobile x Frontend

## ML

## Infra

### First stage

Создание сущностей Task и User удовлетворяющих диаграммам ниже. В базы данных будет использоваться `postgres`

```mermaid
erDiagram
    Task }o--|| User: Operate
    User {
        uuid id PK
        string name 
        string email 
        string password
    }
    Task {
        uuid id PK
        uuid user_id FK 
        string name 
        timestamp created_at 
        bool is_finished 
        string description 
    }
```

Необходимо создать docker-compose, который поднимал бы:
- Client
- Database
- Server (в виде трех приложений: TaskManager, Statistics, Auth)
