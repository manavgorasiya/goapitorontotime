Go API for Current Toronto Time with MySQL Database Logging
===========================================================

This project is a simple Go API that provides the current time in Toronto in JSON format. Additionally, it logs each request's timestamp to a MySQL database.

Features
--------

*   **Current Time Endpoint**: Returns the current time in Toronto (America/Toronto timezone).
    
*   **Database Logging**: Each request logs the current time to a MySQL database.
    
*   **All Times Endpoint**: Retrieves all the logged times from the database.
    

Requirements
------------

Before you start, ensure you have the following installed:

1.  **Go** (version 1.23 or higher)
    
2.  **MySQL** (to create a database and table for logging times)
    
3.  **Docker** (optional for containerizing the app and database)
    

Getting Started
---------------

### 1\. Clone the Repository

First, clone the repository to your local machine:

Plain textANTLR4BashCC#CSSCoffeeScriptCMakeDartDjangoDockerEJSErlangGitGoGraphQLGroovyHTMLJavaJavaScriptJSONJSXKotlinLaTeXLessLuaMakefileMarkdownMATLABMarkupObjective-CPerlPHPPowerShell.propertiesProtocol BuffersPythonRRubySass (Sass)Sass (Scss)SchemeSQLShellSwiftSVGTSXTypeScriptWebAssemblyYAMLXML`   bashCopy codegit clone https://github.com/yourusername/toronto-time-api.git  cd toronto-time-api   `

### 2\. Set Up MySQL Database

1.  Install MySQL if you don’t have it installed.
    
2.  bashCopy codemysql -u root -p
    
3.  sqlCopy codeCREATE DATABASE time\_api;USE time\_api;
    
4.  sqlCopy codeCREATE TABLE time\_log ( id INT AUTO\_INCREMENT PRIMARY KEY, timestamp DATETIME NOT NULL);
    

### 3\. Run the Go Application

1.  Make sure you have Go installed. If not, download and install Go from [https://golang.org/dl/](https://golang.org/dl/).
    
2.  bashCopy codego mod tidygo run main.goThe server will start on port 8080.
    

### 4\. Test the API

Once the server is running, you can test the endpoints:

*   Open a browser or use curl to send a GET request to:bashCopy codehttp://localhost:8080/current-timeIt will return the current time in Toronto in JSON format.
    
*   To see all the timestamps logged to the database, send a GET request to:bashCopy codehttp://localhost:8080/all-timesIt will return a list of all the timestamps in JSON format.
    

Dockerization (Optional)
------------------------

You can run the Go application and MySQL in Docker containers for easy deployment.

1.  bashCopy codedocker build -t toronto-time-api .docker run -p 8080:8080 toronto-time-api
    
2.  bashCopy codedocker run --name mysql-container -e MYSQL\_ROOT\_PASSWORD=root -e MYSQL\_DATABASE=time\_api -p 3306:3306 -d mysql:8Adjust the dsn in your main.go if you are running MySQL inside a container.
    

Error Handling
--------------

The application handles errors such as:

*   Database connection failures
    
*   Time zone loading issues
    
*   Query and data retrieval errors
    

These errors are logged in the console.

Conclusion
----------

This project demonstrates how to create a simple Go API that:

*   Retrieves the current time from a specific time zone.
    
*   Logs the request to a MySQL database.
    
*   Allows you to view all the logged times via a second API endpoint.#   g o a p i t o r o n t o t i m e 
 
 #   g o a p i t o r o n t o t i m e 
 
 #   g o a p i t o r o n t o t i m e  
 