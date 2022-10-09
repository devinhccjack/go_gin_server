package Models

import (
	"fmt"
	"go_server/Databases"
)

type Test struct {
    Id int 
    Testcol string `gorm:"column:testcol"`
}

// 设置Test的表名为`test`
// func (Test) TableName() string {
//     return "test"
// }

// 添加
func (this *Test) Insert() (id int, err error) {
    result := Mysql.DB.Create(&this)
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
    // result := Mysql.DB.Find(&users)
    // fmt.Println("result-----",users,)

    // var firstList = Test{}
    // firstList := new(Test)
    // first := Mysql.DB.First(&firstList)
    // var firstList = []Test{}
    first := Mysql.DB.Select("id,testcol").Find(&firstList)
    //查看返回条数
    fmt.Println("first.RowsAffected-----",first.RowsAffected)
    //判断返回值的错误
    fmt.Println("first.Error---",first.Error) 
    // fmt.Println("查询到的对象为", *firstList)
    //相对于上面的语句建议使用这个输出
	fmt.Println("查询到的对象为", firstList) 
    
    if first.Error != nil {
        err = first.Error
        return 
    }
    return 
}