package main

type Expense struct {
    Name   string `json:"name"`
    Amount string `json:"amount"`
    Id     int    `json:"id"`
}
