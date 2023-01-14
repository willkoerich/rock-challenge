# Desafio Go


Desafio da Stone para equipe de desenvolvimento implementado em Go e utilizando a base de dados Postgres.

# API

Rota *POST /login*

```
curl --location --request POST 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "cpf": "{cpf}",
    "secret": "{plain_text_secret}"
}'
```

Rota *GET /accounts*

```
curl --location --request GET 'http://localhost:8080/accounts'
```

Rota *POST /accounts*
```
curl --location --request POST 'http://localhost:8080/accounts' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "{name}",
    "cpf": "{cpf}",
    "password": "{plain_text_password}",
    "balance": {balance}
}'
```

Rota *GET /accounts/{account_id}/balance*
```
curl --location --request GET 'http://localhost:8080/accounts/{account_id}/balance' 
```

Rota *GET /transfers*

```
curl --location --request GET 'http://localhost:8080/transfers'
```

Rota *POST /transfers*

```
curl --location --request POST 'http://localhost:8080/transfers' \
--header 'Authorization: Bearer {access_token} \
--header 'Content-Type: application/json' \
--data-raw '{
    "account_destination_id": {account_destination_id},
    "amount": {amount_to_transfer}
}'
```


# Como executar

Para executar a aplicação, o projeto consta com um arquivo do docker-compose que quando executado, executa o banco de dados postgres e o client pgAdmin.
A execução da aplicação foi mantida em separado, se fazendo a necessidade de execução por linha de comando ou IDE de escolha.

```bash
> docker-compose up
 ```