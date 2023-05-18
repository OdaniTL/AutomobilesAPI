
# Golang RestAPI  no framework

Creation of an automotive api rest without framework.

## Tech Stack

**Server:** Golang

**DB:** PostgresQL

## Features

- Light/dark mode toggle
- Live previews
- Fullscreen mode
- Cross platform

## Run Locally

Clone the project

```bash
  git clone https://link-to-project
```

Go to the project directory

```bash
  cd AutomobilesAPI
```

Start the server

```bash
  go run mai.go
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`P_SQL_HOST=yourHost`

`P_SQL_PORT=yourport`

`P_SQL_USER=yourUser`

`P_SQL_PASSWORD=yourPassword`

`P_SQL_DBNAME=dbName`

## API Reference

### Autos EndPoints

```http
  GET /api/autos/getAll
```

```http
  GET /api/autos/findById/:id
```

```http
  GET /api/autos/findByName/:name
```

```http
  POST  /api/autos/create
```

```http
  PUT  /api/autos/update
```

```http
  DELETE  /api/autos/delete
```

#### Model  Response for Get Methods

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `AutoId` | `int` | an integer representing the unique identifier of an automobile  |
| `Brand` | `string` | The brand or manufacturer of the automobile |
| `Model` | `string` | The model of the car. For example, "Civic" or "Accord" for a Honda car. |
| `Year` | `int` | The year in which the automobile was manufactured. |
| `BodyType` | `string` | refers to the type of body or design of the car, such as sedan, SUV, coupe, etc. |
| `EngineType` | `string` | The type of engine used in the automobile, such as diesel, petrol, electric, hybrid, etc. |
| `Displacement` | `float32` | The volume of the engine's cylinders, usually measured in liters or cubic centimeters (cc). |
| `State` | `bool` |Price is a property of the Autos struct which represents the price of the automobile. |
