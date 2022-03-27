# transaction-risk

This application receive a list of transactions and return a list of rated risks of these transactions


### To run the app

 - `Make run`


### To run the unit tests

 - `Make run_tests`

### To call the rate risk endpoint

**Request**
```
curl --location --request POST 'localhost:8090/rate-risks' \
--header 'Content-Type: application/json' \
--data-raw '{
"transactions": [
{"id": 1, "user_id": 1, "amount_us_cents": 200000, "card_id": 1},
{"id": 2, "user_id": 1, "amount_us_cents": 600000, "card_id": 1},
{"id": 3, "user_id": 1, "amount_us_cents": 1100000, "card_id": 1},
{"id": 4, "user_id": 2, "amount_us_cents": 100000, "card_id": 2},
{"id": 5, "user_id": 2, "amount_us_cents": 100000, "card_id": 3},
{"id": 6, "user_id": 2, "amount_us_cents": 100000, "card_id": 4}
]
}'
```

**Response**
```
{
    "risk_ratings": [
        "low",
        "medium",
        "high",
        "low",
        "medium",
        "high"
    ]
}
```