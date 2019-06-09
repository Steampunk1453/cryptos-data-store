# cryptos-data-store

BackEnd with Go, JWT, PostgreSQL and Heroku

# URL

https://cryptos-data-store.herokuapp.com

# Endpoints

https://cryptos-data-store.herokuapp.com/api/user/new {POST}

https://cryptos-data-store.herokuapp.com/api/user/login {POST}

https://cryptos-data-store.herokuapp.com/api/crypto/new {POST}

https://cryptos-data-store.herokuapp.com/api/crypto/{id} {GET}

https://cryptos-data-store.herokuapp.com/api/crypto/user/{id} {GET}

https://cryptos-data-store.herokuapp.com/api/crypto {GET}

https://cryptos-data-store.herokuapp.com/api/crypto {PUT}

https://cryptos-data-store.herokuapp.com/api/crypto/{id} {DELETE}

# Local

Create .env and add:

db_name = ******

db_pass = ******

db_user = ******

db_type = ******

db_host = localhost

db_port = 5432

token_password = ******

PORT = 8000

Run

`go run main.go`