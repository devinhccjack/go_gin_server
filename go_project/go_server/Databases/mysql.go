package Mysql

import (
    // "github.com/jinzhu/gorm"
    // _ "github.com/jinzhu/gorm/dialects/mysql"
    "fmt"
    "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var _db *gorm.DB

func Init()  {
    var err error
    //使用dsn连接到数据库，grom自带的数据库池
	//账号:密码@连接方式(ip地址:端口号)/数据库？语言方式，时区（未设置时区的话采用8小时制度）
	dsn := "root:root123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
    //使用mysq连接数据库，第二个参数可以增加更多的配置(可有可无)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) 
	// conn.AutoMigrate(&Student{}) //创建表?判断是否表结构存在
    //不设置，创建生成的表名都是默认是复数形式后面带s的，设置上这一句话，就不会默认带s了
    // db.SingularTable(true) 


    // DB, err = gorm.Open("mysql", dsn)
    if err != nil {
        fmt.Printf("mysql connect error %v", err)
    }
    // if _db.Error != nil {
    //     fmt.Printf("database error %v", _db.Error)
    // }

	sqlDB, err := _db.DB()
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
 
    //设置数据库连接池参数
    sqlDB.SetMaxOpenConns(10)   //设置数据库连接池最大连接数
    sqlDB.SetMaxIdleConns(2)   //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
    sqlDB.SetConnMaxIdleTime(time.Minute) // 最大空闲时间为1分钟的连接池


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

//获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools.getDB()获取db对象即可。
//不用担心协程并发使用同样的db对象会共用同一个连接，db对象在调用他的方法的时候会从数据库连接池中获取新的连接
func GetDB() *gorm.DB {
	return _db
}
