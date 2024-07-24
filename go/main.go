package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

// User represents a user in the database
type User struct {
    EmpNo     int
    FirstName string
    LastName  string
}

var tmpl = template.Must(template.ParseFiles("index.html"))

func main() {
    dsn := "admin:admin@tcp(db:3306)/govsphp"
    maxAttempts := 60

    // Function to connect to the database with retry logic
    connectWithRetry := func(dsn string, maxAttempts int) (*sql.DB, error) {
        var db *sql.DB
        var err error
        for attempts := 1; attempts <= maxAttempts; attempts++ {
            db, err = sql.Open("mysql", dsn)
            if err != nil {
                log.Printf("Attempt %d: failed to open database: %v", attempts, err)
            } else {
                err = db.Ping()
                if err == nil {
                    log.Printf("Successfully connected to the database on attempt %d", attempts)
                    return db, nil
                }
                log.Printf("Attempt %d: failed to connect to database: %v", attempts, err)
            }
            time.Sleep(time.Duration(attempts) * time.Second)
        }
        return nil, err
    }

    // Connect to the database
    db, err := connectWithRetry(dsn, maxAttempts)
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }
    defer db.Close()

    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT emp_no, first_name, last_name FROM employees")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var users []User
        for rows.Next() {
            var user User
            if err := rows.Scan(&user.EmpNo, &user.FirstName, &user.LastName); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            users = append(users, user)
        }

        if err := tmpl.Execute(w, users); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    http.HandleFunc("/json_users", func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT emp_no, first_name, last_name FROM employees")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var users []User
        for rows.Next() {
            var user User
            if err := rows.Scan(&user.EmpNo, &user.FirstName, &user.LastName); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            users = append(users, user)
        }

        // Respond with JSON
        w.Header().Set("Content-Type", "application/json")
        jsonResponse, err := json.Marshal(users)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Write(jsonResponse)
    })

    // Start the server
    fmt.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
