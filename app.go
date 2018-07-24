/**
 *  main file for start running rest api
 *
 *  created by @viyancs, 25/07/2018
 *  make sure always create simple code and clean code,
 *  always give mark comment for easy maintenancen
 */
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
