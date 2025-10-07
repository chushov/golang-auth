# JWT-авторизация на языке Go
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

Пилотный частный проект по сервису авторизации на языке Go, с сохранение шифрованных данных в базу, генерацией и вызовом JWT-токена.

Проверить сервис можно через REST-запросы, которые расположены в папке /rest/.

## Особенности

- jwt
- gin
- gorm
- bcrypt
- mysql, mongodb
- Docker (multi-stage сборка)

## Запуск с Docker

Проект поддерживает запуск в Docker-контейнерах с использованием docker-compose для простого развертывания приложения вместе с MySQL.

### Подготовка

1. Убедитесь, что у вас установлены Docker и docker-compose
2. Создайте файл `.env` на основе `.env.example`:

```bash
cp .env.example .env
```

3. Отредактируйте `.env` файл, установив необходимые переменные окружения

### Запуск

Для запуска всех сервисов (приложение и база данных) выполните:

```bash
docker-compose up -d
```

Приложение будет доступно на порту `8080`, MySQL на порту `3306`.

### Управление контейнерами

Остановка контейнеров:

```bash
docker-compose down
```

Остановка и удаление всех данных (включая volumes):

```bash
docker-compose down -v
```

**Важно:** Для production-окружения обязательно измените пароли и секретные ключи в `.env` файле!

## Установка и запуск

### Шаг 1 — Установка MySQL

Устанавливаем в Ubuntu

```bash
   sudo apt update
   sudo apt install mysql-server
```
Настройка MySQL
```bash
   sudo mysql_secure_installation
```
При использовании плагина валидации пароля скрипт предложит выбрать степень валидации. Самый высокий уровень который можно установить (2), требует, чтобы ваш пароль был не менее восьми символов и содержал строчные буквы, заглавные буквы, цифры и специальные символы.
```bash
    Output
    Securing the MySQL server deployment.

    Connecting to MySQL using a blank password.

    VALIDATE PASSWORD COMPONENT can be used to test passwords
    and improve security. It checks the strength of password
    and allows the users to set only those passwords which are
    secure enough. Would you like to setup VALIDATE PASSWORD component?

    Press y|Y for Yes, any other key for No: Y

    There are three levels of password validation policy:

    LOW    Length >= 8
    MEDIUM Length >= 8, numeric, mixed case, and special characters
    STRONG Length >= 8, numeric, mixed case, special characters and dictionary                  file

    Please enter 0 = LOW, 1 = MEDIUM and 2 = STRONG:
    2
```
Настройка аутентификации и прав пользователя
```bash
   sudo mysql
```
```bash
   SELECT user,authentication_string,plugin,host FROM mysql.user;
```
Не забыть изменить password​​​ на более надежный. Убедиться, что команда заменит пароль root, заданный на шаге 2:
```bash
   ALTER USER 'root'@'localhost' IDENTIFIED WITH caching_sha2_password BY 'password';
```
Примечание. Предыдущее выражение ALTER USER устанавливает аутентификацию root user MySQL с помощью плагина caching_sha2_password​​. Согласно официальной документации MySQL, caching_sha2_password​​​ считается предпочтительным плагином аутентификации MySQL, так как он обеспечивает более защищенное шифрование пароля, чем более старая, но все еще широко используемая версия mysql_native_password.

Однако многие приложения PHP, например phpMyAdmin, работают ненадежно с caching_sha2_password. Если вы планируете использовать эту базу данных с приложением PHP, возможно, вам потребуется установить аутентификацию root с помощью mysql_native_password​​:

```bash
   ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';
```
Затем выполните команду FLUSH PRIVILEGES, которая просит сервер перезагрузить предоставленные таблицы и ввести в действие изменения:
```bash
   FLUSH PRIVILEGES;
```
```bash
   mysql> exit
```
```bash
   sudo mysql
```

### Шаг 2 — База данных

Запускаем службу MySQL или запускаем в Докере. Внутри файла main.go прописываем коннект до базы: порт, пользователь в базе данных и его пароль. Рекомендую создавать отдельного пользователя и никогда ничего не делать от root-пользователя.

### Шаг 3 — Зависимости и запуск приложения

Скачиваем все зависимости проекта. 
Их не много: jwt, gorm, gin, bcrypt и часть пакетов из стандартной библиотеки Go.

```bash
   go get .
```
```bash
   go run .
```

### Шаг 4 — Тестирование REST-запросами

С целью тестирования написанного было принято решение вынести REST-запросы на регистрацию, получение данных и получение токена в отдельные файлы и положить в папку прямо в проекте.

Проще всего тестировать через VSCode \ Goland, с плагином REST Client, который позволяет отправлять запросы прямо из интерфейса и получать ответ в соседнем окне.

Файлы запросов находятся в директории /rest/, а эндпоинты запросов сгруппированы в файле main.go, через возможности фреймворка GIN, который умеет по признаку, в нашем случае /api/, группировать.

На каждый запрос запрос придет ответ и обработка ошибки, в случае ошибки. С помощью GIN терминал выведет каждый обработанный запрос и подсветит его.

Посмотреть на факт создания учетных записей и переданные зашифрованные пароли можно прямо в базе данных. Удобнее всего смотреть через возможности Goland, подключившись к локальной или удаленной базе напрямую из интерфейса, указав порт и созданного для авторизации юзера и имя таблицы. Если имя таблицы не нравится или имя требуется зашифровать, не забудьте поменять название в коннекторе к базе данных в файле main.go.

### Тестирование и CI

Для запуска тестов:

```bash
make test
```
или
```bash
go test ./ ...
```

Для проверки качества кода и статического анализа:

```bash
make lint
```

В проекте реализованы:
- Unit-тесты моделей, auth и базы (папки `models/`, `auth/`, `database/`),
- Интеграционные тесты REST API (controllers/),
- E2E-тест флоу регистрации – получение токена – защищённый эндпоинт,
- Автоматический запуск тестов и линтинга через GitHub Actions на каждый push/pull request в master.

Любой pull request считается принятым, только если все проверки CI проходят успешно.

#### Edge Case тесты

В каждый слой добавьте юнит- или интеграционный тест, проверяющий:

- Некорректный email (не-unique, невалидный формат).
- Пустой токен авторизации, невалидная схема заголовка.
- Ошибки миграции БД (например, неверная модель).

Пример: `TestRegisterUser_InvalidEmail`

```go
func TestRegisterUser_InvalidEmailFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestDB()
	router := gin.Default()
	router.POST("/api/user/register", RegisterUser)

	user := models.User{
		Name:     "Name",
		Username: "User",
		Email:    "wrongformat", // невалидный адрес
		Password: "password",
	}
	payload, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.NotEqual(t, http.StatusCreated, rec.Code)
}
```

## Автор

[@chushov](https://github.com/chushov)
