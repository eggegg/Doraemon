package main

import (
        "github.com/eggegg/Doraemon/handlers"
        "github.com/labstack/echo"
)

func main() {
        // create a new echo instance
        e := echo.New()
        // Route / to handler function
        e.GET("/health-check", handlers.HealthCheck)
       
        // start the server, and log if it fails             
		e.Logger.Fatal(e.Start(":8080")) 
}