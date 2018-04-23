package main

import (
    "net/http"
    "log"
    "./routes"
)

func main() {
    router := routes.NewRouter()
    if err := http.ListenAndServe(":8080", router); err != nil {
       log.Fatal(err)
    } else {
        log.Println("Started server in port 8080")
    }
}