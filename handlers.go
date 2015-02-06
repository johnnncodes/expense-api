package main

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

var utils *Utils = &Utils{}

type RequestHandler struct {
}

func (rh *RequestHandler) ExpenseList(w http.ResponseWriter, r *http.Request) {
    utils.JsonResponse(w, expenses)
}

func (rh *RequestHandler) ExpenseCreate(w http.ResponseWriter, r *http.Request) {
    expense := &Expense{}

    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&expense)

    if utils.ErrorCheck(w, err) {
        return
    }

    lastExpenseId = lastExpenseId + 1
    expenses = append(expenses, Expense{expense.Name, expense.Amount, lastExpenseId})

    utils.JsonResponse(w, expenses)
}

func (rh *RequestHandler) ExpenseDetail(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])

    if utils.ErrorCheck(w, err) {
        return
    }

    var requestedExpense Expense
    requestedExpenseNotFound := true

    for _, expense := range expenses {
        if expense.Id == id {
            requestedExpense = expense
            requestedExpenseNotFound = false
        }
    }

    if requestedExpenseNotFound {
        w.WriteHeader(http.StatusNotFound)
    } else {
        utils.JsonResponse(w, requestedExpense)
    }
}
