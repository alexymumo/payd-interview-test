# payd-interview-test

Solution for backend developer intern position payd 

``Prerequisite``
- Golang
- PostgreSQL 
- Postman - for testing endpoints
- Payd API credentials 

``Technologies used``
- Golang
- Gin - HTTP web framework
- Gorm - ORM for Golang
- PostgreSQL
- JWT - JSON WEB TOKENS
- RabbitMQ

``Installation``

- Clone the repository 

``git clone https://github.com/alexymumo/payd-interview-test.git``

- Create a `.env` file in the root directory and add the following environment variables

```

- DB_HOST=
- DB_USER=
- DB_PASSWORD=
- DB_NAME=
- DB_PORT=
- JWT_SECRET=
- AMQP_URL=

```

- Running Locally using  ``go run main.go``

- Use postman to test the api endpoints


``Authentication``

- Register new user

Payload

``
{
    "email":"alex@gmail.com",
    "name":"Alex Test",
    "password":"1234"
}
``


- Authenticate registered user

``
{
    "email":"alex@gmail.com",
    "password":"1234"
}
``


``Payment``

