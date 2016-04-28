package JSLog

import (
    "log"
    "net/http"
    "encoding/json"
)

//Validator is a function that validates a passed token
var Validator TokenValidator

//Logger is a function that receives logs
var Logger LogReceiver

//Handle registeres the http log path
func Handle(path string){
    Logger = defaultLogger
    Validator = defaultValidator
    http.HandleFunc(path, logRequest)
}

func logRequest(w http.ResponseWriter, r *http.Request){
    if !Validator(r.Header.Get("Token")){
        http.Error(w, "Token was rejected", 403)
        return
    }
    
    decoder := json.NewDecoder(r.Body)
    item := new(LogItem)
    
    if err := decoder.Decode(item); err != nil{
        http.Error(w, "Request could not be decoded", 500)
        return
    }
    
    if err := Logger(item); err != nil{
        http.Error(w, "The request could not be logged", 500)
        return
    }
    
    w.Write([]byte("Request has been logged."))
}

//TokenValidator takes a passed token and validates it. If false is returned the log request is rejected. Default behaviour is to accept all requests.
type TokenValidator func(token string) bool

//LogReceiver takes a log item and processes it. Default behaviour is to print the log using the go "log" package
type LogReceiver func(item *LogItem) error

type LogItem struct{
    Level string
    Message string
    URL string
    Line int
}

func defaultLogger(item *LogItem) error{
    switch item.Level {
    case "Error":
        log.Fatalln(item)
        break
    case "Warning":
        log.Println(item)
        break
    default:
        log.Println(item)
    }
    return nil
}

func defaultValidator(token string) bool{
    return true
}