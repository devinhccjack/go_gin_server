package Mysql

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "fmt"
    // "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

var DB *gorm.DB

func Init()  {
    var err error
    //使用dsn连接到数据库，grom自带的数据库池
	//账号:密码@连接方式(ip地址:端口号)/数据库？语言方式，时区（未设置时区的话采用8小时制度）
	dsn := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
    //使用mysq连接数据库，第二个参数可以增加更多的配置(可有可无)
	// conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) 
	// conn.AutoMigrate(&Student{}) //创建表?判断是否表结构存在
    //不设置，创建生成的表名都是默认是复数形式后面带s的，设置上这一句话，就不会默认带s了
    // db.SingularTable(true) 


    DB, err = gorm.Open("mysql", dsn)
    if err != nil {
        fmt.Printf("mysql connect error %v", err)
    }
    if DB.Error != nil {
        fmt.Printf("database error %v", DB.Error)
    }

    // dsn := "root:123456@tcp(127.0.0.1:3306)/gormDB?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.New(mysql.Config{
	// 	DSN: dsn,
	// }), &gorm.Config{
	// 	SkipDefaultTransaction: false, //跳过默认事务
	// 	NamingStrategy: schema.NamingStrategy{
	// 		SingularTable: false,  // 复数形式 User的表名应该是users
	// 		TablePrefix:   "gva_", //表名前缀 User的表名应该是t_users
	// 	},
	// 	DisableForeignKeyConstraintWhenMigrating: true, //设置成为逻辑外键(在物理数据库上没有外键，仅体现在代码上)
	// })
	
	// // 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	// sqlDB, err := db.DB()
	// // SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	// sqlDB.SetMaxIdleConns(10)
	// // SetMaxOpenConns 设置打开数据库连接的最大数量。
	// sqlDB.SetMaxOpenConns(100)
	// // SetConnMaxLifetime 设置了连接可复用的最大时间。
	// sqlDB.SetConnMaxLifetime(time.Hour)

}