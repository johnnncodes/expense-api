package main

import (
    "encoding/json"
    "net/http"
)

type Utils struct {
}

func (u *Utils) JsonResponse(w http.ResponseWriter, data interface{}) {
    payload, err := json.Marshal(data)
    if u.ErrorCheck(w, err) {
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Write(payload)
}

func (u *Utils) ErrorCheck(w http.ResponseWriter, err error) bool {
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return true
    }
    return false
}
