package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error
	dsn := "root:87102011@tcp(127.0.0.1:3306)/time_api"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v\n", err)
	}
	fmt.Println("Connected to the database!")
}

func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		log.Printf("Error loading timezone: %v\n", err)
		http.Error(w, "Error loading timezone", http.StatusInternalServerError)
		return
	}
	torontoTime := time.Now().In(loc)

	_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", torontoTime)
	if err != nil {
		log.Printf("Error logging time to database: %v\n", err)
		http.Error(w, "Error logging time to database", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"current_time": torontoTime.Format(time.RFC3339)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func allTimesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT timestamp FROM time_log")
	if err != nil {
		log.Printf("Error fetching data: %v\n", err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var times []string
	for rows.Next() {
		var timestamp time.Time
		if err := rows.Scan(&timestamp); err != nil {
			log.Printf("Error scanning data: %v\n", err)
			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			return
		}
		times = append(times, timestamp.Format(time.RFC3339))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(times)
}

func main() {
	initDB()
	defer db.Close()

	http.HandleFunc("/current-time", currentTimeHandler)
	http.HandleFunc("/all-times", allTimesHandler)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
