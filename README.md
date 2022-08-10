
# JWT-авторизация на языке Go
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

Попытка написать авторизацию на языке Go, с сохранение шифрованных данных в базу, генерацией и вызовом JWT-токена.

Проверка работоспособности происходит через REST-запросы, которые расположены в папке /rest/.



## Особенности

- jwt
- gin
- gorm
- bcrypt
- Проверка через REST-запросы
- Работает на простой тривиальной MySQL / MongoDB


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
Если вы решите использовать плагин валидации пароля, скрипт предложит вам выбрать уровень валидации пароля. Самый высокий уровень, который можно установить, указав 2, требует, чтобы ваш пароль был длиной не менее восьми символов и содержал строчные, заглавные буквы, цифры и специальные символы.
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
Не забудьте изменить password​​​ на более надежный пароль и убедитесь, что эта команда заменит пароль root, заданный на шаге 2:
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

Скачиваем все зависимости проекта. Их не много: jwt, gorm, gin, bcrypt и часть пакетов из стандартной библиотеки Go.

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
## Автор

- [@chushov](https://www.github.com/chushov)


