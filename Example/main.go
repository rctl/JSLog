package main

import (
    "net/http"
    "github.com/rasmusj-se/JSLog"
)

func main() {
    JSLog.Handle("/log")
    JSLog.Logger = logger
    JSLog.Validator = authenticator
    http.ListenAndServe(":80", nil)
}

func logger(item *JSLog.LogItem) error{
    //Do something with item
    return nil
}

func authenticator(token string) bool{
    //Do something to validate the token if needed
    return true
}