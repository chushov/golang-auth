![Go](https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-8.0+-4479A1?style=flat&logo=mysql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-20.10+-2496ED?style=flat&logo=docker&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?style=flat)
![GORM](https://img.shields.io/badge/GORM-ORM-00ADD8?style=flat)
![JWT](https://img.shields.io/badge/JWT-Auth-000000?style=flat&logo=json-web-tokens&logoColor=white)
![CI](https://img.shields.io/badge/CI-GitHub%20Actions-2088FF?style=flat&logo=github-actions&logoColor=white)
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

# RESTful API для JWT-авторизации на Go

## 📌 О проекте

Это **учебный частный проект**, созданный для изучения и демонстрации следующих технологий и подходов:

- 🔐 JWT-авторизация и аутентификация пользователей
- 🌐 Построение RESTful API на фреймворке Gin
- 💾 Работа с базой данных MySQL через ORM (GORM)
- 🔒 Хеширование паролей с использованием bcrypt
- 🐳 Контейнеризация приложения с Docker и docker-compose
- ✅ Написание unit, интеграционных и E2E-тестов
- 🔄 Настройка CI/CD pipeline с GitHub Actions

### ⚠️ ВАЖНОЕ ПРЕДУПРЕЖДЕНИЕ

**Данный проект предназначен ИСКЛЮЧИТЕЛЬНО для образовательных и учебных целей.**

- ✅ Можно использовать для обучения и изучения технологий
- ✅ Можно использовать в личных экспериментальных проектах
- ❌ **НЕ используйте этот код в production-окружении**
- ❌ Проект не проходил security-аудит и может содержать уязвимости
- ❌ Не предназначен для коммерческого или критически важного использования

Для production-решений рекомендуется использовать проверенные библиотеки и следовать best practices безопасности.

## 🛠 Технологии

- **Язык:** Go 1.18+
- **Фреймворк:** Gin (HTTP web framework)
- **ORM:** GORM
- **База данных:** MySQL 8.0+
- **Аутентификация:** JWT (JSON Web Tokens)
- **Хеширование:** bcrypt
- **Контейнеризация:** Docker, docker-compose
- **Тестирование:** Go testing, testify
- **CI/CD:** GitHub Actions

## 🚀 Быстрый старт с Docker

### Требования

- Docker 20.10+
- docker-compose 1.29+

### Шаги запуска

1. **Клонируйте репозиторий:**

```bash
git clone https://github.com/chushov/golang-auth.git
cd golang-auth
```

2. **Создайте файл `.env` на основе примера:**

```bash
cp .env.example .env
```

3. **Отредактируйте `.env` файл:**

Обязательно измените следующие переменные:

```env
DB_HOST=mysql
DB_PORT=3306
DB_USER=authuser
DB_PASSWORD=your_secure_password  # Замените на свой пароль
DB_NAME=authdb

JWT_SECRET=your_jwt_secret_key_here  # Замените на случайную строку

SERVER_PORT=8080
```

4. **Запустите приложение:**

```bash
docker-compose up -d
```

5. **Проверьте статус контейнеров:**

```bash
docker-compose ps
```

Приложение будет доступно по адресу: `http://localhost:8080`

### Управление контейнерами

**Остановка контейнеров:**
```bash
docker-compose down
```

**Остановка и удаление всех данных (включая volumes):**
```bash
docker-compose down -v
```

**Просмотр логов:**
```bash
docker-compose logs -f app
```

**Перезапуск контейнеров:**
```bash
docker-compose restart
```

## 💻 Локальная установка (без Docker)

### Требования

- Go 1.18 или выше
- MySQL 8.0 или выше

### Шаг 1: Установка и настройка MySQL

**Установка в Ubuntu:**

```bash
sudo apt update
sudo apt install mysql-server
```

**Безопасная настройка MySQL:**

```bash
sudo mysql_secure_installation
```

**Создание базы данных и пользователя:**

```bash
sudo mysql
```

```sql
CREATE DATABASE authdb;
CREATE USER 'authuser'@'localhost' IDENTIFIED BY 'your_secure_password';
GRANT ALL PRIVILEGES ON authdb.* TO 'authuser'@'localhost';
FLUSH PRIVILEGES;
EXIT;
```

### Шаг 2: Настройка проекта

1. **Клонируйте репозиторий:**

```bash
git clone https://github.com/chushov/golang-auth.git
cd golang-auth
```

2. **Создайте `.env` файл:**

```bash
cp .env.example .env
```

3. **Настройте переменные окружения в `.env`:**

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=authuser
DB_PASSWORD=your_secure_password
DB_NAME=authdb

JWT_SECRET=your_jwt_secret_key_here

SERVER_PORT=8080
```

### Шаг 3: Установка зависимостей

```bash
go mod download
```

### Шаг 4: Запуск приложения

```bash
go run .
```

Приложение запустится на порту `8080`.

## 🧪 Тестирование

### Запуск всех тестов

```bash
make test
```

или

```bash
go test ./...
```

### Запуск тестов с покрытием

```bash
go test -cover ./...
```

### Линтинг кода

```bash
make lint
```

### Структура тестов

Проект включает:

- ✅ **Unit-тесты** для моделей, auth и database слоёв
- ✅ **Интеграционные тесты** для REST API контроллеров
- ✅ **E2E-тесты** полного flow: регистрация → логин → получение токена → доступ к защищённым эндпоинтам
- ✅ **Edge case тесты** для проверки граничных случаев и ошибок

### CI/CD

Проект использует GitHub Actions для автоматического запуска:

- Всех тестов
- Линтинга кода
- Проверки сборки

Все проверки должны пройти успешно для принятия pull request.

## 📡 API Эндпоинты

### Регистрация пользователя

```http
POST /api/user/register
Content-Type: application/json

{
  "name": "John Doe",
  "username": "johndoe",
  "email": "john@example.com",
  "password": "securePassword123"
}
```

### Вход (получение JWT токена)

```http
POST /api/user/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "securePassword123"
}
```

**Ответ:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Получение данных пользователя (защищённый эндпоинт)

```http
GET /api/user/profile
Authorization: Bearer <your_jwt_token>
```

## 📂 Структура проекта

```
.
├── .github/          # GitHub Actions CI/CD конфигурация
├── auth/             # Логика JWT авторизации
├── controllers/      # HTTP обработчики (handlers)
├── database/         # Настройка подключения к БД
├── models/           # Модели данных (GORM)
├── rest/             # Примеры REST запросов
├── service/          # Middleware и сервисные функции
├── .env.example      # Пример переменных окружения
├── .gitignore
├── Dockerfile        # Multi-stage Docker сборка
├── docker-compose.yml # Orchestration для app + MySQL
├── e2e_auth_test.go  # E2E тесты
├── go.mod            # Go модули
├── go.sum
├── main.go           # Точка входа приложения
├── Makefile          # Команды для сборки и тестов
├── README.md
└── swagger.yml       # API документация
```

## 🧪 Тестирование с REST Client

В директории `/rest/` находятся примеры HTTP запросов для тестирования API.

**Рекомендуемые инструменты:**
- VSCode с расширением REST Client
- JetBrains IDEs (GoLand, IntelliJ) с встроенным HTTP Client
- Postman
- curl

### Пример использования

1. Откройте файл из `/rest/` в редакторе
2. Используйте плагин REST Client для отправки запросов
3. Скопируйте полученный JWT токен
4. Используйте токен в заголовке `Authorization: Bearer <token>` для защищённых эндпоинтов

## 📝 Переменные окружения

Все необходимые переменные окружения описаны в файле `.env.example`:

| Переменная | Описание | Пример значения |
|------------|----------|------------------|
| `DB_HOST` | Хост базы данных | `localhost` или `mysql` (для Docker) |
| `DB_PORT` | Порт базы данных | `3306` |
| `DB_USER` | Пользователь БД | `authuser` |
| `DB_PASSWORD` | Пароль пользователя БД | `your_secure_password` |
| `DB_NAME` | Имя базы данных | `authdb` |
| `JWT_SECRET` | Секретный ключ для JWT | случайная строка (минимум 32 символа) |
| `SERVER_PORT` | Порт приложения | `8080` |

## 🤝 Вклад в проект

Это учебный проект, но pull request'ы приветствуются для:

- Исправления ошибок
- Улучшения документации
- Добавления новых тестов
- Примеров использования

**Важно:** Все PR должны проходить CI проверки (тесты + линтинг).

## 📄 Лицензия

MIT License - см. файл [LICENSE](LICENSE) для деталей.

## 👤 Автор

**@chushov**

- GitHub: [@chushov](https://github.com/chushov)

---

⭐ Если этот проект помог вам в изучении Go и JWT-авторизации, поставьте звезду на GitHub!
