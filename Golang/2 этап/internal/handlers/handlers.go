package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"tracker/internal/models"
)

var (
	expenses = []models.Expense{}
	nextID   = 1
	tmpl     = template.Must(template.ParseGlob("web/templates/*.html"))
)

func ListExpenses(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Expenses []models.Expense
	}{
		Expenses: expenses,
	}

	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, "Ошибка рендеринга: "+err.Error(), http.StatusInternalServerError)
	}
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	amountStr := r.FormValue("amount")
	description := r.FormValue("description")
	dateStr := r.FormValue("date")

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Неверный формат суммы", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Неверный формат даты", http.StatusBadRequest)
		return
	}

	expenses = append(expenses, models.Expense{
		ID:          nextID,
		Amount:      amount,
		Description: description,
		Date:        date,
	})
	nextID++

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Это учебный проект по веб-разработке на Go.")
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "pong")
}
