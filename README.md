# JSLog - A JavaScript logging framework

JavaScript Log catches errors from JavaScript and posts them to a backend server.
Server implementation in Go is available.

## Usage (Frontend)

Include the jslog

`<script src="js/log.js" type="text/javascript"></script>`

JSLog is automatically initialized to post logs to path /log and will listen for js errors that occurr.

To change to custom settings you can run:

`log.Configure("/backend/log/path", "MyUserToken");`

You can also manually log events by doing:

`log.Info("This is some info");`
`log.Warning("This is a warning");`
`log.Error("This is some error");`

## Usage (Backend)

With the JSLog there is a provided Go framework available to use.
The following is a simple example of how to use is:

```
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
```

## License

This project is licensed under MIT License. Use it freely.

## Contrib

If you wish to improve existing or implement new features just create a pull request.
It would be nice to evolve the project to add handlers for more languages such as Node, Python, Ruby etc 
