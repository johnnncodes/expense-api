package main

import (
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
)

var expenses []Expense = make([]Expense, 0)
var lastExpenseId int = 0

func main() {
    router := mux.NewRouter()
    requestHandler := &RequestHandler{}

    router.HandleFunc("/api/expenses/", requestHandler.ExpenseList).Methods("GET")
    router.HandleFunc("/api/expenses/", requestHandler.ExpenseCreate).Methods("POST")
    router.HandleFunc("/api/expenses/{id}/", requestHandler.ExpenseDetail).Methods("GET")

    n := negroni.Classic()
    n.UseHandler(router)
    n.Run(":3000")
}
