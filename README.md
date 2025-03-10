# Go Database Access Example

This project demonstrates how to access a MySQL database using Go. It follows the official Go tutorial for database access, which can be found [here](https://go.dev/doc/tutorial/database-access).

## Project Structure

- `main/main.go`: The main Go file that contains the code to connect to the database and perform CRUD operations.
- `create-tables.sql`: SQL script to create the `album` table and insert some initial data.

## Prerequisites

- Go 1.16 or later
- MySQL database

## Setup

1. **Clone the repository:**
   ```sh
   git clone <repository-url>
   cd data-acess
   ```

2. **Set up the MySQL database:**
   - Create a MySQL database named `recordings`.
   - Run the `create-tables.sql` script to create the `album` table and insert initial data.
     ```sh
     mysql -u <username> -p recordings < create-tables.sql
     ```

3. **Set environment variables:**
   ```sh
   export DBUSER=<your-database-username>
   export DBPASS=<your-database-password>
   ```

## Running the Application

Navigate to the `main` directory and run the Go application:
```sh
cd main
go run main.go
```

## Functionality

The application performs the following operations:

1. **Connect to the database:** Establishes a connection to the MySQL database using the provided credentials.
2. **Retrieve albums by artist:** Fetches albums by a specific artist.
3. **Retrieve album by ID:** Fetches an album by its ID.
4. **Add a new album:** Inserts a new album into the database.
5. **Remove an album by ID:** Deletes an album from the database by its ID.