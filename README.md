# Сервис Управления Товарами

## Обзор

Сервис Управления Товарами - это легкий сервис, разработанный для облегчения взаимодействия с товарами на складе. Он предоставляет простые конечные точки для получения информации о запасах и резервирования продукции.

## Установка

1. **Клонируйте репозиторий:**

   ```bash
   git clone git@github.com:mynamedust/simple-stock.git

2. **Перейдите в директорию проекта:**

   ```bash
   cd simple-stock

3. **Установите зависимости:**

   ```bash
   go mod download  
## Запуск

1. **Настройте доступ к базе данных в файле config.yml**

2. **Запустите приложение (будет доступно по адресу http://localhost:80):**
   ```bash
   go build -o ./stock ./cmd/ && ./stock

3. **Воспользуйтесь Postman или curl командами.**

## Конечные точки
1. GET /products/stock: Получение информации о доступных товарах на складе.
2. POST /products/reserve: Резервирование товаров с предоставлением необходимых деталей в теле запроса.
2. POST /products/release: Освобождение резерва товаров с предоставлением необходимых деталей в теле запроса.

## Примеры запросов
1. **Получение информации о товарах на складе:**
  ```bash
  curl -X GET http://localhost:80/products/stock \
  -H "Content-Type: application/vnd.api+json" \
  -d '{
    "data": {
      "type": "storehouse",
        "id": "2"
    }
  }'
```
  Ожидаемый ответ:
  ```bash
  {
    "data": {
      "type": "storehouse",
      "id": "2",
      "attributes": {
        "count": 615
      }
    }
  }
```
2. **Резервация товаров на складе:**
  ```bash
  curl -X POST http://localhost:80/products/reserve \
  -H "Content-Type: application/vnd.api+json" \
  -d '{
    "data": [
      {
        "type": "product",
        "id": "1",
        "attributes": {
          "storehouse_id": 1,
          "code": "ABC123"
        }
      },
      {
        "type": "product",
        "id": "2",
        "attributes": {
          "storehouse_id": 2,
          "code": "XYZ789"
        }
      }
    ]
  }'
```
  Ожидаемый ответ:
  ```bash
  {
    "errors": [
      {
        "id": "1",
        "title": "Product reservation failed",
        "detail": "Storehouses are not available",
        "status": "500"
      }
    ]
  }
  ```
3. **Освобождение резерва товаров:**
  ```bash
  curl -X POST http://localhost:80/products/reserve \
  -H "Content-Type: application/vnd.api+json" \
  -d '{
    "data": [
      {
        "type": "product",
        "id": "1",
        "attributes": {
          "storehouse_id": 1,
          "code": "P0012"
        }
      },
      {
        "type": "product",
        "id": "2",
        "attributes": {
          "storehouse_id": 1,
          "code": "P0005"
        }
      }
    ]
  }'
```
## Использование Makefile и Docker

Вы можете использовать `Makefile` для удобного управления различными задачами в вашем приложении. Вот несколько команд, которые могут вам пригодиться:
1. **Запус приложения в режиме отладки с использованием Docker Compose:**
```bash
  make up
```
2. **Остановка приложения:**
```bash
  make down
```
3. **Очистка ресурсов Docker:**
```bash
  make down
```
>Важно: Эта команда удаляет все контейнеры и тома, связанные с вашим приложением. Убедитесь, что вы действительно хотите выполнить эту команду, прежде чем запускать ее.

   
