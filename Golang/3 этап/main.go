package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"tracker/internal/handlers"
	"tracker/internal/repository"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "expenses.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	schema := `
	CREATE TABLE IF NOT EXISTS expenses (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	amount REAL NOT NULL,
	description TEXT NOT NULL,
	date DATE NOT NULL
	);`
	_, err = db.Exec(schema)
	if err != nil {
		panic(err)
	}

	repo := repository.NewExpenseRepository(db)
	h := handlers.NewHandler(repo)

	http.HandleFunc("/", h.ListExpenses)
	http.HandleFunc("/expenses", h.CreateExpense)
	http.HandleFunc("/edit", h.EditExpense)
	http.HandleFunc("/update", h.UpdateExpense)
	http.HandleFunc("/delete", h.DeleteExpense)
	http.HandleFunc("/about", h.AboutHandler)
	http.HandleFunc("/ping", h.PingHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
