package Models

import (
	"fmt"
    "time"
	"go_server/Databases"
    // "gorm.io/gorm"

)

type Test struct {
    Id int `gorm:"column:id;primarykey;AUTO_INCREMENT" json:"id"`
    Testcol string `gorm:"column:testcol" json:"testcol"`
	Title       string `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Description string `gorm:"column:description;type:varchar(255);" json:"description"`
	Status      int `gorm:"column:status;type:tinyint;" json:"status"`
	CreatedAt *time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;" json:"deleted_at"`
}

// 设置Test的表名为`test`
// func (Test) TableName() string {
//     return "test"
// }

// 添加
func (this *Test) Insert() (id int, err error) {
    // 数据库的表与模型建立关系
    Mysql.GetDB().AutoMigrate(Test{})

    result := Mysql.GetDB().Create(&this)

    id = this.Id
    if result.Error != nil {
        err = result.Error
        return
    }
    return
}

// 查询
func (this *Test) GetTestDataModel() ( firstList []Test, err error) {
    // users := []Test{} //实例化对象    
     // 数据库的表与模型建立关系
      //获取DB
    db := Mysql.GetDB()

    err = db.AutoMigrate(Test{})
    if err != nil {
        fmt.Println(err)
    }
    first := db.Find(&firstList)
    // fmt.Println("result-----",users,)

    // var firstList = Test{}
    // firstList := new(Test)
    // first := Mysql.DB.First(&firstList)
    // var firstList = []Test{}
    // first := Mysql.DB.Select("id,testcol").Find(&firstList)
    //查看返回条数
    fmt.Println("first.RowsAffected-----",first.RowsAffected)
    //判断返回值的错误
    // fmt.Println("first.Error---",first.Error) 
    // fmt.Println("查询到的对象为", *firstList)
    //相对于上面的语句建议使用这个输出
	fmt.Println("查询到的对象为", firstList) 
    
    if first.Error != nil {
        err = first.Error
        return 
    }
    return 
}