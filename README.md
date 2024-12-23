# Go API for Current Toronto Time with MySQL Database Logging

This project is a simple Go API that provides the current time in
Toronto in JSON format. Additionally, it logs each request's timestamp
to a MySQL database.

## Features

-   **Current Time Endpoint**: Returns the current time in Toronto
    (America/Toronto timezone).

-   **Database Logging**: Each request logs the current time to a MySQL
    database.

-   **All Times Endpoint**: Retrieves all the logged times from the
    database.

## Requirements

Before you start, ensure you have the following installed:

1.  **Go** (version 1.23 or higher)

2.  **MySQL** (to create a database and table for logging times)

3.  **Docker** (optional for containerizing the app and database)

## Getting Started

### 1. Clone the Repository

First, clone the repository to your local machine:

git clone https://github.com/yourusername/toronto-time-api.git
cd toronto-time-api

### 2. Set Up MySQL Database

1.  Install MySQL if you don't have it installed.

2.  Log in to MySQL:
    ```
    mysql -u root -p
    ```

3.  Create a new database for this project:
    ```
    CREATE DATABASE time_api;
    USE time_api;
    ```

4.  Create the time_log table:
    ```
    CREATE TABLE time_log ( 
        id INT AUTO_INCREMENT PRIMARY KEY, 
        timestamp DATETIME NOT NULL 
        );
        111

### 3. Run the Go Application

1.  Make sure you have Go installed. If not, download and install Go
    from <https://golang.org/dl/>.

2.  In the project folder, run the following commands to build and run the app:
    ```
    go mod tidy
    go run main.go
    ```
    
    The server will start on port 8080.

### 4. Test the API

Once the server is running, you can test the endpoints:

-   Open a browser or use curl to send a GET request to:
    ```
    http://localhost:8080/current-time
    ```

    It will return the current time in Toronto in JSON format.

-   To see all the timestamps logged to the database, send a GET requestto:
    ```
    http://localhost:8080/all-times
    ```

    It will return a list of all the timestamps in JSON format.

## Dockerization (Optional)

You can run the Go application and MySQL in Docker containers for easy
deployment.

1.  Build and run the Docker container:

    ```
    docker build -t toronto-time-api .
    docker run -p 8080:8080 toronto-time-api
    ```

2.  If you'd like to run MySQL in a Docker container, you can use the following:
   
    ```
    docker run --name mysql-container -e
    MYSQL_ROOT_PASSWORD=root -e 
    MYSQL_DATABASE=time_api -p 3306:3306 -d mysql:8
    ```

    Adjust the dsn in your main.go if you are running MySQL inside a container.

## Error Handling

The application handles errors such as:

-   Database connection failures

-   Time zone loading issues

-   Query and data retrieval errors

These errors are logged in the console.

## Conclusion

This project demonstrates how to create a simple Go API that:

-   Retrieves the current time from a specific time zone.

-   Logs the request to a MySQL database.

-   Allows you to view all the logged times via a second API endpoint.

