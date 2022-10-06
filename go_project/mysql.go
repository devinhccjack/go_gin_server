package main

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func newPool() *sql.DB {
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "root123456"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "testdb"
	dsn := cfg.FormatDSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

var pool = newPool()

// User 用户
type User struct {
	ID   int64
	Name string
}

// QueryUser 根据id查询用户
// 注意：如果结果集为空集，这个代码是有问题的
func QueryUser(id int64) (*User, error) {
	rows, err := pool.Query("select `id`, `name` from `users` where `id` = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // 注意这里，一定要关闭
	user := User{}
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		break
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		fmt.Printf("ddd0")
		QueryUser(1)
		
		c.JSON(200, gin.H{
			"message": "pong0",
		})
		// c.String(200, "Hello, Geektutu")
	})
	//设定请求url不存在的返回值

	router.Run() // listen and serve on 0.0.0.0:8080

}