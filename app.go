package main

import (
    "log"
    "github.com/go-gem/gem"
)

func index(ctx *gem.Context) {
    ctx.HTML(200, "hello world")
}

func main() {
    // Create server.
    srv := gem.New(":8282")

    // Create router.
    router := gem.NewRouter()
    // Register handler
    router.GET("/", index)

    // Start server.
    log.Println(srv.ListenAndServe(router.Handler()))
}
