// // main.go
// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/go-sql-driver/mysql"
// )

// var db *sql.DB

// func main() {
// 	router := gin.Default()
// 	router.LoadHTMLFiles("./login.html")   //解析网页
// 	router.LoadHTMLFiles("web/index.html") //解析网页

// 	// initDB()
// 	err := initDB()
// 	if err != nil {
// 		fmt.Printf("init db error:%v \n ", err)
// 	} else {
// 		fmt.Println("连接成功db")
// 	}
// 	defer db.Close()

// 	router.GET("/list", getList) //get方法
// 	v1_admin := router.Group("admin")
// 	v1_admin.GET("/list", LoginPage)
// 	v1_admin.POST("/login", LoginForm)
// 	v1_admin.GET("/morelist", queryMultiRowDemo)

// 	router.GET("/", func(c *gin.Context) {
// 		fmt.Printf("ddd0")
// 		c.JSON(200, gin.H{
// 			"message": "pong0",
// 		})
// 		// c.String(200, "Hello, Geektutu")
// 	})
// 	//设定请求url不存在的返回值
// 	router.NoRoute(NoResponse)
// 	defer db.Close()

// 	router.Run() // listen and serve on 0.0.0.0:8080

// }

// func mysqlConnect() {

// 	dsn := "root:root123456@tcp(127.0.0.1:3306)/testdb"

// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		panic(err)
// 		log.Fatal(err)
// 	}
// 	// defer db.Close()
// 	if err := db.Ping(); err != nil {
// 		fmt.Println("连接失败")
// 		panic(err)
// 	}
// 	fmt.Println("连接成功")
// 	//根据实际业务设置数值
// 	// db.SetConnMaxLifetime(time.Second*10) // 设置连接可以重复使用的最长时间
// 	// db.SetMaxOpenConns(500) // 最大连接数
// 	// db.SetMaxIdleConns(200) // 最大空闲连接数
// }

// type userInfo struct {
// 	id        int
// 	user_name string
// 	pass_word string
// 	phone     string
// 	remark    string
// }

// //查询单条数据示例
// func queryRowDemo() {
// 	fmt.Printf("scan failed ---")
// 	var name string
// 	err := db.QueryRow("select user_name from test_user where id = ?", 1).Scan(&name)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			// 没有查询到就结果，不认为是错误
// 		} else {
// 			log.Fatal(err)
// 		}
// 	}
// 	fmt.Println("name---", name)

// 	// sqlStr := "SELECT id,user_name,pass_word,phone,remark FROM test_user WHERE id=1"
// 	// var u userInfo
// 	// //非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
// 	// err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.user_name, &u.pass_word, &u.phone, &u.remark)
// 	// fmt.Printf("---", err)

// 	// if err != nil {
// 	// 	fmt.Printf("scan failed， err: %v\n", err)
// 	// 	return
// 	// }
// 	// fmt.Printf("id:%d,user_name:%s,pass_word:%s\n", u.id, u.user_name, u.pass_word)
// }

// //Client提交的数据
// type SqlUser struct {
// 	Name    string `json:"name"`
// 	Age     int    `json:"age"`
// 	Address string `json:"address"`
// }

// //应答体（响应client的请求）
// type SqlResponse struct {
// 	Code    int         `json:"code"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data"`
// }

// // 查询多条数据示例
// func queryMultiRowDemo(c *gin.Context) {
// 	sqlStr := "SELECT id,user_name,pass_word FROM test_user WHERE id>?"
// 	// var u userInfo
// 	var sqlResponse SqlResponse

// 	rows, err := db.Query(sqlStr, 0)
// 	if err != nil {
// 		fmt.Printf("query failed, err:%v\n", err)
// 		sqlResponse.Code = http.StatusBadRequest
// 		sqlResponse.Message = "查询错误"
// 		sqlResponse.Data = "error"
// 		c.JSON(http.StatusOK, sqlResponse)
// 		return
// 	}
// 	//非常重要，关闭rows释放持有的数据库链接
// 	defer rows.Close()

// 	//读出查询出的列字段名
// 	cols, _ := rows.Columns()
// 	//values是每个列的值，这里获取到byte里
// 	values := make([][]byte, len(cols))
// 	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
// 	scans := make([]interface{}, len(cols))
// 	//让每一行数据都填充到[][]byte里面,狸猫换太子
// 	for i := range values {
// 		scans[i] = &values[i]
// 	}

// 	results := make([]map[string]string, 0, 10)
// 	for rows.Next() {
// 		err := rows.Scan(scans...)
// 		if err != nil {
// 			return //nil, err
// 		}
// 		row := make(map[string]string, 10)
// 		for k, v := range values { //每行数据是放在values里面，现在把它挪到row里
// 			key := cols[k]
// 			row[key] = string(v)
// 		}
// 		results = append(results, row)
// 	}

// 	resUser := make([]userInfo, 0, 10)
// 	// var resUser []userInfo
// 	//初始化切片，用户存储查询结果
// 	// resUser := []*userInfo{}

// 	for rows.Next() {
// 		// var userTemp userInfo
// 		var u userInfo
// 		// 获取各列的值，放到对应的地址中
// 		err := rows.Scan(&u.id, &u.user_name, &u.pass_word)
// 		if err != nil {
// 			fmt.Printf("scan failed, err: %v\n", err)
// 			sqlResponse.Code = http.StatusBadRequest
// 			sqlResponse.Message = "查询错误"
// 			sqlResponse.Data = "error"
// 			c.JSON(http.StatusOK, sqlResponse)
// 			return
// 		}
// 		// userTemp.id = u.id
// 		// userTemp.user_name = u.user_name
// 		// userTemp.pass_word = u.pass_word
// 		u.phone = "phone"
// 		u.remark = "remark"

// 		resUser = append(resUser, u)

// 		fmt.Println("id:%d,user_name:%s,pass_word:%s\n", u.id, u.user_name, u.pass_word)
// 	}
// 	fmt.Println("resUser---", resUser,results)

// 	// return &u
// 	sqlResponse.Code = http.StatusOK
// 	sqlResponse.Message = "读取成功"
// 	sqlResponse.Data = results
// 	fmt.Println("sqlResponse.Data---", sqlResponse.Data)

// 	// c.JSON(http.StatusOK, gin.H{"sqlResponse":sqlResponse})
// 	c.JSON(http.StatusOK, sqlResponse)


// }

// // QueryUser1 单行查询
// func QueryUser1(id int) (*userInfo, error) {
// 	row := db.QueryRow("select `id`, `user_name`,`pass_word`,`phone`,`remark` from `test_user` where `id` = ?", id)
// 	user := userInfo{}
// 	// var uesr userInfo
// 	if err := row.Scan(&user.id, &user.user_name, &user.pass_word, &user.phone, &user.remark); err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil // 返回 (*User)(nil) 表示查询结果不错在
// 		}
// 		return nil, err
// 	}
// 	fmt.Println("sql----", &user)
// 	return &user, nil
// }

// // 请求
// // 获取多条数据
// func GetMoreList(c *gin.Context) {
// 	// 查询
// 	// queryMultiRowDemo()
// 	// userRes, errRes := queryMultiRowDemo()
// 	// fmt.Println("查询结果为：", userRes,userRes.user_name,userRes.pass_word, " error为：", errRes)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "pong1",
// 		"code":    0,
// 		"list":    "sql",
// 		// "name":    name,
// 		// "userObj":   userObj,
// 	})
// }

// func getList(c *gin.Context) {
// 	name := c.DefaultQuery("name", "Guest") //找不到name给它一个默认值Guest
// 	// queryRowDemo()
// 	id := c.DefaultQuery("id", "0")

// 	// defer QueryUser1(1)
// 	// x := 1
// 	x, err := strconv.Atoi(id)
// 	if err != nil {
// 		fmt.Printf("%v 转换失败！", id)
// 	} else {
// 		fmt.Printf("type:%T value:%#v\n", id, x)
// 	}

// 	// int, err := strconv.Atoi(id)
// 	fmt.Println("id----", x, id)

// 	// 查询
// 	userRes, errRes := QueryUser1(x)
// 	fmt.Println("查询结果为：", userRes, userRes.user_name, userRes.pass_word, " error为：", errRes)

// 	userObj := []string{userRes.user_name, userRes.pass_word, userRes.phone, userRes.remark}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "pong1",
// 		"code":    0,
// 		"list":    "sql",
// 		"name":    name,
// 		"userObj": userObj,
// 	})
// }

// func LoginPage(c *gin.Context) {
// 	c.HTML(http.StatusOK, "login.html", nil)

// 	// c.JSON(200, gin.H{
// 	// 	"message": "pong1",
// 	// 	"code":    0,
// 	// 	"list":    "admin",
// 	// })
// }

// //  登录处理
// func LoginForm(c *gin.Context) {
// 	username := c.PostForm("username")
// 	password := c.PostForm("password")
// 	fmt.Println("接收到参数username=" + username + ",password=" + password)

// 	sqlStr := "insert into test_user( user_name, pass_word,phone,remark) values (?,?,?,?)"
// 	var userform userInfo
// 	// userform.id = 2
// 	userform.user_name = username
// 	userform.pass_word = password
// 	userform.phone = "phone"
// 	userform.remark = "remark"

// 	ret, err := db.Exec(sqlStr, userform.user_name, userform.pass_word, userform.phone, userform.remark)

// 	if err != nil {

// 		fmt.Printf("insert failed, err:%v\n", err)

// 	} else {
// 		// c.JSON(http.StatusOK,gin.H{
// 		// 	"msg":"error",
// 		// })
// 		// c.PureJSON(http.StatusOK, gin.H{
// 		// 	"html": "<b>error!</b>",
// 		// })
// 	}
// 	fmt.Println(ret.LastInsertId()) //打印结果

// 	// formInfo := username + "-" + password
// 	// c.String(http.StatusOK,"data",formInfo)
// 	// c.JSON(http.StatusOK, gin.H{
// 	// 	"message": "success",
// 	// 	"code":    0,
// 	// 	"list":    "login inset",
// 	// 	"data":    formInfo,
// 	// })
// 	// c.PureJSON(200, gin.H{
// 	// 	"html": "<b>success!</b>",
// 	// })
// 	c.HTML(http.StatusOK, "web/index.html", gin.H{
// 		"title": "后台",
// 		"desc":  "web",
// 	})
// }

// func NoResponse(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "not url",
// 		"code":    0,
// 		"list":    "url 404",
// 	})
// }

// func initDB() (err error) {
// 	dsn := "root:root123456@tcp(127.0.0.1:3306)/testdb"
// 	// 这里db不要使用:= , 全局变量赋值 , 在main中使用
// 	// 这里不会校验账号密码是否正确
// 	db, err = sql.Open("mysql", dsn)
// 	if err != nil {
// 		return err
// 	}
// 	//尝试数据库连接
// 	err = db.Ping()
// 	if err != nil {
// 		return err
// 	}
// 	db.SetConnMaxLifetime(time.Second * 10) //连接存活最大时间
// 	db.SetMaxIdleConns(200)                 //最大空闲连接数
// 	db.SetMaxOpenConns(10)                  // 最大连接数
// 	return nil
// }

package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}