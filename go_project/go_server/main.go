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