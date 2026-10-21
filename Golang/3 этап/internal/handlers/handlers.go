package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"tracker/internal/models"
	"tracker/internal/repository"
)

type Handler struct {
	repo *repository.ExpenseRepository
	tmpl *template.Template
}

func NewHandler(repo *repository.ExpenseRepository) *Handler {
	return &Handler{
		repo: repo,
		tmpl: template.Must(template.ParseGlob("web/templates/*.html")),
	}
}

func (h *Handler) ListExpenses(w http.ResponseWriter, r *http.Request) {
	expenses, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, "Ошибка БД: "+err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		Expenses []models.Expense
	}{
		Expenses: expenses,
	}
	h.tmpl.ExecuteTemplate(w, "layout", data)
}

func (h *Handler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)
	description := r.FormValue("description")
	date, _ := time.Parse("2006-01-02", r.FormValue("date"))

	newExpense := &models.Expense{
		Amount:      amount,
		Description: description,
		Date:        date,
	}

	if err := h.repo.Create(newExpense); err != nil {
		http.Error(w, "Ошибка сохранения", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) EditExpense(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	expense, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, "Трата не найдена", http.StatusNotFound)
		return
	}
	h.tmpl.ExecuteTemplate(w, "edit", expense)
}

func (h *Handler) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	id, _ := strconv.Atoi(r.FormValue("id"))
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)
	description := r.FormValue("description")
	date, _ := time.Parse("2006-01-02", r.FormValue("date"))

	updateExpense := &models.Expense{
		ID:          id,
		Amount:      amount,
		Description: description,
		Date:        date,
	}

	if err := h.repo.Update(updateExpense); err != nil {
		http.Error(w, "Ошибка обновления", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, "Ошибка удаления", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Это учебный проект по веб-разработке на Go"))
}

func (h *Handler) PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("pong"))
}
