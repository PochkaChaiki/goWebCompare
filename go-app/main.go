package main

import (
	"encoding/json"
	"gowebcompare/todo"
	"log"
	"net/http"
	"strconv"
)

var TodoList *todo.TodoList

func main() {
	TodoList = todo.New()

	mux := http.NewServeMux()

	// mux.Handle("/api/")
	mux.HandleFunc("GET /api/get_list", getList)
	mux.HandleFunc("GET /api/get_todo/{id}", getTodo)
	mux.HandleFunc("POST /api/create_todo", createTodo)
	mux.HandleFunc("PUT /api/update_todo", updateTodo)
	mux.HandleFunc("DELETE /api/delete_todo/{id}", deleteTodo)

	s := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: RequestLoggerMiddleware(mux),
	}

	log.Printf("Server is running on %s\n", s.Addr)
	log.Fatal(s.ListenAndServe())
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func getList(w http.ResponseWriter, r *http.Request) {
	list := TodoList.GetList()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := TodoList.GetTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo todo.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := TodoList.CreateTodo(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	var todo todo.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := TodoList.UpdateTodo(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := TodoList.DeleteTodo(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
