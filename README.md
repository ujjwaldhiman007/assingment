# Go Fiber Car Management API

This is a RESTful API built with Go Fiber for managing car records. It provides basic CRUD operations for car entities using MongoDB as the database.

## API Endpoints

The API provides the following endpoints:

- `GET /api/cars`: Retrieve all cars
- `POST /api/cars`: Create a new car
- `PATCH /api/cars/:id`: Update an existing car
- `DELETE /api/cars/:id`: Delete a car

## Setup and Running

1. Ensure you have Go installed on your system.
2. Clone this repository.
3. Install dependencies:
   ```
   go mod tidy
   ```
4. Set up your MongoDB connection string in an `.env` file:
   ```
   MONGODB_URI=your_mongodb_connection_string
   DB_NAME=your_database_name
   ```
5. Run the application:
   ```
   go run main.go
   ```

## API Usage

### Get All Cars
```
GET /api/cars
```

### Create a Car
```
POST /api/cars
Content-Type: application/json

{
  "name": "Car Name"
}
```

### Update a Car
```
PATCH /api/cars/:id
Content-Type: application/json

{
  "name": "Updated Car Name"
}
```

### Delete a Car
```
DELETE /api/cars/:id
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the [MIT License](LICENSE).