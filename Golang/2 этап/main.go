package main

import (
	"fmt"
	"net/http"

	"tracker/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.ListExpenses)
	http.HandleFunc("/expenses", handlers.CreateExpense)
	http.HandleFunc("/about", handlers.AboutHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера", err)
	}
}
