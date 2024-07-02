myproject/
|-- main.go
|-- config/
|   |-- database.go
|-- handlers/
|   |-- employer_handler.go
|-- middlewares/
|   |-- logger.go
|   |-- request_id.go
|-- models/
|   |-- models.go
|-- repository/
|   |-- employer_repository.go
|   |-- employer_memory.go
|   |-- employer_postgres.go
|-- server/
|   |-- server.go


CREATE TABLE persons (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100)
);

CREATE TABLE employers (
    id SERIAL PRIMARY KEY,
    company VARCHAR(100),
    person_id INT REFERENCES persons(id)
);
