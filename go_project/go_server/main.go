package main

import (
    "gin/Router"
    "gin/Databases"
)

func main() {
    defer Mysql.DB.Close()
    Router.InitRouter()
}