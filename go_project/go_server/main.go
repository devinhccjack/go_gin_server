package main

import (
    "go_server/Router"
    "go_server/Databases"
)

func main() {
    Mysql.Init()
    defer Mysql.DB.Close()
    Router.InitRouter()
}

// test build docker image
// package main

// import "github.com/gin-gonic/gin"

// func main() {
//     r := gin.Default()
//     r.GET("/ping", func(c *gin.Context) {
//         c.JSON(200, gin.H{
//             "message": "pong",
//         })
//     })

//     r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
// }