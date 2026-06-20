
# Subscription Management API

API для управления подписками пользователей.

Запустить:

```bash
docker-compose up --build
```

Сервер: `http://localhost:8080`
Swagger: `http://localhost:8080/swagger/`

## API Endpoints
![Uploading image.png…]()

### POST /created-sub
Создать подписку

**Request body:**
```json
{
    "service_name": "Netflix",
    "price": 1000,
    "user_id": "d2f6cf95-ea12-4d74-b7fb-4215c94475a7",
    "start_date": "05-2026",
    "end_date": "06-2026"
}
```

**Response:**
- `201` — подписка успешно создана
- `400` — ошибка валидации данных
- `500` — внутренняя ошибка сервера


### GET /get-sub
Получить все подписки

**Response:**
```json
[
    {
        "id_sub": "1",
        "service_name": "Netflix",
        "price": 1000,
        "user_id": "d2f6cf95-ea12-4d74-b7fb-4215c94475a7",
        "start_date": "05-2026",
        "end_date": "06-2026"
    }
]
```
- `200` — список подписок
- `500` — ошибка получения данных


### GET /get-sub/{id}
Получить подписку по ID

**Response:**
- `200` — данные подписки
- `400` — неверный формат ID
- `404` — подписка не найдена


### GET /get-sub-sum
Получить сумму подписок за период + флитры
обязательно
 "end_date": "06-2026"
 "start_date": "05-2026"
 опционально
 "service_name": "test"
 "user_id": "d2f6cf95-ea12-4d74-b7fb-4215c94475a7" - только UUID пользователя

**Пример:**
```
/get-sub-sum?start_date=05-2026&end_date=08-2026
```

**Response:**
```json
{
    "total_sum": 3000,
    "currency": "RUB",
    "period": "05-2026 - 08-2026"
}
```
- `200` — сумма подписок
- `400` — отсутствуют обязательные параметры
- `500` — ошибка получения суммы


### PUT /update-sub/{id}
Полное обновление подписки

**Request body:** как в POST /created-sub

**Response:**
- `200` — подписка успешно обновлена
- `400` — ошибка валидации данных
- `404` — подписка не найдена
- `500` — ошибка обновления


### PATCH /update-sub/{id}
Частичное обновление подписки

**Request body (все поля опциональны):**
```json
{
    "price": 1400
}
```

**Response:**
- `200` — подписка обновлена
- `400` — ошибка валидации или все поля пустые
- `404` — подписка не найдена
- `500` — ошибка обновления


### DELETE /delete-sub/{id}
Удалить подписку

**Response:**
- `200` — подписка успешно удалена
- `400` — неверный формат ID
- `500` — ошибка удаления

## Структура проекта

```
src/
  main.go                 # точка входа
  config/database.go      # подключение к PostgreSQL
  routes/routes.go        # маршруты API
  subscription/
    controller/           # обработчики HTTP запросов
    dto/                  # модели данных
    service/              # бизнес-логика
  utils/log.go            # логгер ошибок
migrations/               # SQL миграции
docs/                     # Swagger документация
```

## Технологии

- Go 1.26+
- PostgreSQL 15
- Docker, Docker Compose
- Swagger (OpenAPI 2.0)

