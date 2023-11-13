# Сервис Резервирования

## Обзор

Сервис резервирования - это легкий сервис, разработанный для облегчения взаимодействия с товарами на складе. Он предоставляет простые конечные точки для получения информации о запасах и резервирования продукции.

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
2. POST /products/release: Освобождение зарезервированных товаров с предоставлением необходимых деталей в теле запроса.

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
  make clean
```
>Важно: Эта команда удаляет все контейнеры и тома, связанные с вашим приложением. Убедитесь, что вы действительно хотите выполнить эту команду, прежде чем запускать ее.

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
    "data": [
        {
            "type": "product_dto",
            "id": "2",
            "attributes": {
                "quantity": 200
            }
        },
        {
            "type": "product_dto",
            "id": "2",
            "attributes": {
                "quantity": 75
            }
        },
        {
            "type": "product_dto",
            "id": "2",
            "attributes": {
                "quantity": 180
            }
        },
        {
            "type": "product_dto",
            "id": "2",
            "attributes": {
                "quantity": 90
            }
        },
        {
            "type": "product_dto",
            "id": "2",
            "attributes": {
                "quantity": 70
            }
        }
    ]
}
```
2. **Резервация товаров на складе:**
  ```bash
  curl -X POST http://localhost:80/products/reserve \
  -H "Content-Type: application/vnd.api+json" \
  -d '{
  "data": {
    "id": "1",
    "type": "storehouse_dto",
    "relationships": {
      "products": {
        "data": [
          {
            "type": "product_dto",
            "id": "1",
            "attributes": {
              "quantity": 1
            }
          },
          {
            "type": "product_dto",
            "id": "2",
            "attributes": {
              "quantity": 15
            }
          }
        ]
      }
    }
  }
}'
```
3. **Освобождение зарезервированных товаров:**
  ```bash
  curl -X POST http://localhost:80/products/release \
  -H "Content-Type: application/vnd.api+json" \
  -d '{
  "data": {
    "id": "1",
    "type": "storehouse_dto",
    "relationships": {
      "products": {
        "data": [
          {
            "type": "product_dto",
            "id": "1",
            "attributes": {
              "quantity": 1
            }
          }
        ]
      }
    }
  }
}'
```
