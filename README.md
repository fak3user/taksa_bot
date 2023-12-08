# бот такса

## develop

Стек среды разработки

```
go 1.21.3
entr
```

Для поднятия дев-окружения, находясь в папке с проектом, нужно запустить:

```
git ls-files -cdmo --exclude-standard | entr -cr go run main.go
```

Также нужно создать файл `.env` в корневой директории (можно использовать `.env.example` как пример)

## глоссарий

-   _чат_ - отдельный чат в телеграме, куда добавлен бот
-   _счёт_ - событие или череда событий, во время которого считается сумма потраченных средств состоит из _транзакций_
-   _участник_ - пользователь из _чата_, участвубщий в _счёте_
-   _транзакция_ - запись о конкретной трате c указанием суммы, названия, заплатившего и _участников_

## запросы и фичи

-   создание нового _счета_ внутри _чата_

    -   Endpoint: /api/chats/{chatId}/events
    -   Type: POST
    -   JSON Request Body:
        ```yaml
        { "title": "Trip to Rome", "users": ["@user1", "@user2"] }
        ```
    -   TypeScript object:
        ```typescript
        interface Event {
            id: string;
            title: string;
            chatId: string;
            users: User[];
            transactions: Transaction[];
            closed: boolean;
        }
        ```

-   состояние _счета_
    -   Endpoint: /api/events/{eventId}/summary
    -   Method: GET
    -   Typescript object: такой же как в создании нового _счета_
-   добавление _участника_ в _счет_

    -   Endpoint: /api/events/{eventId}/users
    -   Method: POST
    -   Request Body:
        ```yaml
        { "username": "string" }
        ```
    -   TypeScript object:
        ```typescript
        interface User {
            id: string;
            username: string;
        }
        ```

-   добавление _транзакций_ для каждого из участников

    -   Endpoint: Endpoint: /api/events/{eventId}/transactions
    -   Method: POST
    -   Request Body:
        ```yaml
        {
            "title": "string",
            "amount": "number",
            "splitType": "equal|exact|percentage",
            "splits": [
                {
                "userId": "string",
                "amount": "number", // Used if splitType is 'exact'
                "percentage": "number" // Used if splitType is 'percentage'
                }
            ]
        }
        ```
    -   TypeScript object:

        ```typescript
        interface Transaction {
            id: string;
            title: string;
            notes?: string;
            paidBy: string;
            amount: number;
            splitType: "equal" | "exact" | "percentage";
            splits: Split[];
        }

        interface Split {
            userId: string;
            amount?: number; // Optional because it's used only for 'exact' splitType
            percentage?: number; // Optional because it's used only for 'percentage' splitType
        }
        ```

-   редактирование _транзакции_
    -   Endpoint: /api/events/{eventId}/transactions/{transactionId}
    -   Method: PUT
    -   Request Body:
        ```yaml
        {
            "title": "string",
            "amount": "number",
            "splitType": "equal|exact|percentage",
            "splits": [
                {
                "userId": "string",
                "amount": "number", // Used if splitType is 'exact'
                "percentage": "number" // Used if splitType is 'percentage'
                }
            ]
        }
        ```
-   получение сумм _счета_
    -   Endpoint: /api/events/{eventId}/summary
    -   Method: GET
    -   TypeScript:
        ```typescript
        interface EventSummary {
            [userId: string]: {
                owes: {
                    // Amounts this user owes to others, could be negative
                    [otherUserId: string]: number;
                };
                // The total amount this user owes to others, negative
                // if user owns, positive if other users owns to this user
                totalOwed: number;
            };
        }
        ```
-   получение списка _транзакций_
    -   Endpoint: /api/events/{eventId}/transactions
    -   Method: GET
-   возможность указать, что
    TODO
-   закрытие счета
    -   Endpoint: /api/events/{eventId}/close
    -   Method: PATCH
    -   Request Body:
        ```yaml
        { "closed": "boolean" }
        ```
    -   TypeScript: тогглит в объекте аккаунта closed
-   нотифаи для каждого участника при добавлении _транзакции_
    TODO

## База данных

```json
{
  "definitions": {
    "User": {
      "type": "object",
      "properties": {
        "id": { "type": "string" },
        "username": { "type": "string" }
      },
      "required": ["id", "username"]
    },
    "Split": {
      "type": "object",
      "properties": {
        "userId": { "type": "string" },
        "amount": { "type": "number" },
        "percentage": { "type": "number" }
      },
      "required": ["userId"]
    },
    "Transaction": {
      "type": "object",
      "properties": {
        "id": { "type": "string" },
        "title": { "type": "string" },
        "amount": { "type": "number" },
        "paidBy": { "type": "string"},
        "splitType": { "type": "string", "enum": ["equal", "exact", "percentage"] },
        "splits": {
          "type": "array",
          "items": { "$ref": "#/definitions/Split" }
        }
      },
      "required": ["id", "title", "amount", "paidBy" "splitType", "splits"]
    },
    "Event": {
      "type": "object",
      "properties": {
        "id": { "type": "string" },
        "title": { "type": "string" },
        "chatId": { "type": "string" },
        "users": {
          "type": "array",
          "items": { "$ref": "#/definitions/User" }
        },
        "transactions": {
          "type": "array",
          "items": { "$ref": "#/definitions/Transaction" }
        },
        "closed": { "type": "boolean", "default": false }
      },
      "required": ["id", "title", "chatId", "users", "transactions"]
    }
  },
  "type": "object",
  "properties": {
    "events": {
      "type": "array",
      "items": { "$ref": "#/definitions/Event" }
    }
  },
  "required": ["events"]
}

```

## пайплайн

-   Создание нового счета:
    -   Пользователь создает новый счет в чате, указывая название счета.
    -   Бот создает новый объект счета с уникальным идентификатором, названием и пустым списком транзакций и участников.
-   Добавление участников:
    -   Участники в чате могут добавить себя в счет, чтобы участвовать в распределении расходов.
    -   Бот обновляет список участников в объекте счета.
-   Добавление транзакций:
    -   Участники могут добавлять транзакции, указывая название, сумму и участников, которые участвуют в этой транзакции.
    -   Бот создает новую транзакцию и добавляет ее в список транзакций для данного счета.
-   Редактирование транзакции:
    -   Участники могут редактировать уже добавленные транзакции, изменяя сумму или участников.
    -   Бот обновляет информацию о транзакции в объекте счета.
-   Личный расчет суммы:
    -   Участники могут запросить личный расчет, чтобы узнать, кто кому должен или кому нужно вернуть деньги.
    -   Бот анализирует все транзакции и вычисляет, сколько каждый должен или должны вернуть друг другу.
-   Закрытие счета:
    -   Когда все расходы учтены и участники выполнили расчеты, счет может быть закрыт.
    -   Бот помечает счет как закрытый и больше не позволяет добавлять новые транзакции.
-   Уведомления:
    -   Бот отправляет уведомления каждому участнику при добавлении новой транзакции, чтобы участники всегда были в курсе изменений в счете.
