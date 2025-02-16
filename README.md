# 📝 Сервис обработки текста ⚙️

![Go Version](https://img.shields.io/badge/go-1.21%2B-blue)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![SQLite](https://img.shields.io/badge/SQLite-3-green.svg)

Современный микросервис для обработки текста с мониторингом состояния и постоянным хранилищем. Создан с ❤️ на Go!

```bash
curl -X POST -H "Content-Type: text/plain" --data "Ваш секретный текст... 🤫"
```

## 🌟 Возможности

✅ **Мгновенная проверка состояния сервера**  
`GET /health` → `{"status": "server is running"}`

✨ **Умная обработка текста**  
Сохранение любых текстовых данных с автоматическим логированием IP клиента.

🔒 **Хранилище на основе SQLite**  
Постоянное хранение данных с автоматическим управлением схемой.

📊 **Структурированное логирование**  
Красивые логи с использованием `slog` для наблюдаемости.

---

## 🚀 Начало работы

### 🔧 Необходимые зависимости

- **Go** 1.21+ ([скачать](https://golang.org/))
- **SQLite3** ([подробнее](https://www.sqlite.org/index.html))

### ⚡ Быстрая установка

```bash
git clone https://github.com/makhkets/interceptor
cd interceptor
go mod tidy
sqlite3 storage/database.db < schema.sql
```

### 🏃 Запуск сервера

```bash
go run main.go
```

Сервер запустится по адресу [http://localhost:80](http://localhost:80) 🎉

---

## 📚 Документация API

### 🩺 Проверка состояния

```http
GET /health
```

#### Ответ:

```json
{"status": "server is running"}
```

---

### 💌 Обработка текста

```http
POST /steal
Content-Type: text/plain

Ваш конфиденциальный текст...
```

#### Успешный ответ:

```
Data saved successfully
```

---

# 📂 Структура проекта

Проект построен на **Go** и включает в себя модули для обработки текста, логирования и работы с базой данных.

```bash
Project:
│   📜 go.mod              # Модульные зависимости Go
│   📜 go.sum              # Контрольная сумма зависимостей
│   🚀 main.go             # Главная точка входа
│   🛠️ makefile            # Скрипты сборки и команд
│   📖 README.md           # Документация проекта
│
├── 🏗️ build/               # Собранные бинарные файлы
│   │   ⚙️ interceptor.exe  # Исполняемый файл сервиса
│   │
│   └── 🗄️ storage/
│           🛢️ database.db  # Файл базы данных SQLite
│
├── 📦 gen/                 # Сгенерированный код OpenAPI
│   │   📜 api.yaml         # OpenAPI спецификация
│   │   🛠️ generate.go      # Скрипт генерации
│   │
│   └── 📂 interceptor/
│           🛠️ oas_cfg_gen.go
│           🛠️ oas_client_gen.go
│           🛠️ oas_handlers_gen.go
│           🛠️ oas_interfaces_gen.go
│           🛠️ oas_json_gen.go
│           🛠️ oas_labeler_gen.go
│           🛠️ oas_middleware_gen.go
│           🛠️ oas_operations_gen.go
│           🛠️ oas_request_decoders_gen.go
│           🛠️ oas_request_encoders_gen.go
│           🛠️ oas_response_decoders_gen.go
│           🛠️ oas_response_encoders_gen.go
│           🛠️ oas_router_gen.go
│           🛠️ oas_schemas_gen.go
│           🛠️ oas_server_gen.go
│           🛠️ oas_unimplemented_gen.go
│
├── 🔧 internal/            # Внутренние модули проекта
│   ├── 📂 domain/
│   │   └── 📝 model.go     # Определение моделей данных
│   │
│   ├── 📂 lib/             # Библиотеки
│   │   └── 📂 logger/      # Логирование
│   │       │   📜 logs.log      # Файл логов
│   │       │   🛠️ setup.go      # Настройка логов
│   │       │
│   │       ├── 📂 handlers/
│   │       │   ├── 🗑️ slogdiscard/
│   │       │   │   └── 🛠️ slogdiscard.go
│   │       │   │
│   │       │   └── 🎨 slogpretty/
│   │       │       └── 🛠️ slogpretty.go
│   │       │
│   │       └── 🛠️ sl/
│   │           └── 🛠️ sl.go
│   │
│   ├── 📂 migrations/      # Миграции базы данных
│   │   └── 📜 2_init.up.sql
│   │
│   ├── 📂 pkg/             # Пакеты вспомогательных функций
│   │   └── 📂 directories/
│   │       └── 🛠️ util.go
│   │
│   ├── 📂 server/          # Логика работы сервера
│   │   └── 🌍 server.go
│   │
│   └── 📂 storage/         # Работа с базой данных
│       │   🛠️ init.go
│       │
│       └── 📂 sqlite/
│           └── 🛢️ db.go
│
├── 🔄 migrator/            # Миграция базы данных
│   └── 🛠️ main.go
│
└── 🗄️ storage/              # Хранилище базы данных
    └── 🛢️ database.db
```

### 📌  эмодзи:
- 📂 — Папка
- 📜 — Файл конфигурации
- 🛠️ — Кодовая логика
- 🛢️ — База данных
- 🌍 — Сервер
- 🎨 — UI/красивый вывод
- 🗑️ — Игнорируемый/ненужный лог
- 🔄 — Миграция базы данных

---

## 🛠️ Примеры использования

### 📩 Обработка текста

```bash
curl -X POST http://localhost/steal \
  -H "Content-Type: text/plain" \
  -d "Мой секретный рецепт: 1 чашка магии ✨"
```

### ✅ Проверка состояния

```bash
curl http://localhost/health
# {"status":"server is running"}
```

---

## 🔍 Схема базы данных

```sql
CREATE TABLE IF NOT EXISTS logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    text TEXT NOT NULL,
    remote_address TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

---

## 💠 Пример логов

```log
2023/10/15 12:00:00 INFO server is running on :80
2023/10/15 12:00:05 INFO received text from 127.0.0.1 length=42
```
---

