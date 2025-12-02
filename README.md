# coord-gin-rest

Минималистичный REST API сервер на Go с Gin, Elasticsearch (заглушка) и zerolog.

## Стек

- **Go 1.24**
- **Gin Web Framework** - легкий HTTP роутер
- **Zerolog** - структурированное логирование
- **Elasticsearch** (подготовлено, пока mock)
- **Docker + Docker Compose** - контейнеризация

## Структура проекта

```
coord-gin-rest/
├── cmd/
│   └── api/
│       └── main.go              # Точка входа
├── internal/
│   ├── config/              # Конфигурация
│   ├── logger/              # Логгер
│   ├── middleware/          # Middleware
│   ├── db/                  # БД (mock)
│   ├── model/               # Модели
│   ├── service/             # Сервисы
│   └── handler/             # HTTP хэндлеры
├── docker/
│   └── Dockerfile
├── docker-compose.yml
├── .env.example
└── go.mod
```

## Эндпоинты

| Метод | Путь | Описание |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/v1/api` | Данные (mock) |

## Быстрый старт

### 1. Клонирование

```bash
git clone https://github.com/UberionAI/coord-gin-rest.git
cd coord-gin-rest
```

### 2. Настройка .env

```bash
cp .env.example .env
```

### 3. Запуск

**Docker Compose:**
```bash
docker-compose up --build
```

**Локально:**
```bash
go mod tidy
go run cmd/api/main.go
```

## Тестирование

```bash
# Health check
curl http://localhost:8080/health

# Получение данных
curl http://localhost:8080/v1/api
```

## Логирование

Уровень `WARNING` по умолчанию. Пример:

```
2025/12/02 11:20:01 INF Starting coord-gin-rest service
2025/12/02 11:20:01 WRN Elasticsearch not configured, using mock data
2025/12/02 11:20:01 INF Server starting port=8080
```

## Graceful Shutdown

Поддерживается корректное завершение при SIGINT/SIGTERM (timeout 5s).

## Лицензия

MIT
