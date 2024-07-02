# payd-interview-test

Solution for backend developer position at payd 

``Prerequisite``
- Golang
- PostgreSQL 
- Postman - for testing endpoints
- Payd API credentials ie ``username`` ``password``

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
![Screenshot from 2024-07-02 12-47-24](https://github.com/alexymumo/payd-interview-test/assets/56880898/3ed26f31-c66b-4975-a699-acb204ff603b)


- Authenticate registered user

``
{
    "email":"alex@gmail.com",
    "password":"1234"
}
``

![Screenshot from 2024-07-02 12-48-56](https://github.com/alexymumo/payd-interview-test/assets/56880898/4320501a-a2d5-481a-bcf3-b596def7bfdd)


``Payment``

