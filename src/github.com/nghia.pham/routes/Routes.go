package routes

import (
    "net/http"
    "github.com/nghia.pham/handler"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
    Route{
        "Index",
        "GET",
        "/",
        handler.Index,
    },
    Route{
        "TodoIndex",
        "GET",
        "/todos",
        handler.TodoIndex,
    },
    Route{
        "TodoShow",
        "GET",
        "/todos/{todoId}",
        handler.TodoShow,
    },
}
