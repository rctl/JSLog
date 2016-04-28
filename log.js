function Logging(server, token) {
    
    var log = this;
    
    if (server === ""){
        this.Server = "/log";
    }else{
        this.Server = server;
    }
        
    this.Token = token;
    
    this.Configure = function(server, token){
       log.Token = token;
       log.Server = server;
    };

    this.PostLog = function(level, message, url, line){
        $.post(log.Server, JSON.stringify({ Level: level, Message: message, Url: url, Line: line }), { Token: log.Token }).success(function (data) {
            console.log("Logging: A log item has been transmitted to a backend server.");
        }).fail(function (data) {
            console.log("Logging: The error could not be logged.");
        });
    };
    
    this.Info = function(message){
        log.PostLog("Info", message, window.location, 0);
    };
    
    this.Warning = function(message){
        log.PostLog("Warning", message, window.location, 0);
    }
    
    this.Error = function(message){
        log.PostLog("Error", message, window.location, 0);
    };
    
    window.onerror = function(message, url, lineNumber) {  
        this.PostLog("Error", message, url, lineNumber);
        return true;
    };
   
}

var log = Logging("/log", "");