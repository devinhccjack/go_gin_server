package main

import (
    "go_server/Router"
    "go_server/Databases"
)

func main() {
    defer Mysql.DB.Close()
    Router.InitRouter()
}