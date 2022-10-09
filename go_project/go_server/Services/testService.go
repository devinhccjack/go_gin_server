package Services

import (
	"fmt"
    "time"
	"go_server/Models/testModel"
)

type Test struct {
    Id int `json:"id"` 
    Testcol string `json:"testcol"`
    Title       string `json:"title"`
	Description string `json:"description"`
	Status      int `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// 添加
func (this *Test) Insert() (id int, err error) {
    var testModel Models.Test
    // testModel.Id = this.Id
    testModel.Testcol = this.Testcol
    testModel.Title = this.Title
    testModel.Description = this.Description
    testModel.Status = this.Status

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

