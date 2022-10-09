package Services

import (
	"fmt"
	"go_server/Models/testModel"
)

type Test struct {
    Id int `json:"id"` 
    Testcol string `json:"testcol"`
}

// 添加
func (this *Test) Insert() (id int, err error) {
    var testModel Models.Test
    testModel.Id = this.Id
    testModel.Testcol = this.Testcol
    id, err = testModel.Insert()
    return
}

// 查询
func (this *Test) GetTestData() ( list []Models.Test, err error) {
    var testModel Models.Test
    // var list []Models.Test
    // list := make([]Models.Test, 0, 10)
    list, err = testModel.GetTestDataModel()
    fmt.Println("list-service---",list)
    return 
}

