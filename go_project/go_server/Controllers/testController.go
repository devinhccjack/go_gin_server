package Controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	// "strconv"
	"go_server/Services"
)

// 添加
func TestInsert(c *gin.Context) {
    var testService Services.Test
    title := c.PostForm("title")
    testcol := c.PostForm("testcol")
	fmt.Println("title=" + title + ",testcol=" + testcol)

    err := c.ShouldBindJSON(&testService)
    err1 := c.BindJSON(&testService)

    fmt.Println("err---" ,err,err1,testService)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
            "Code": http.StatusBadRequest,
            "Msg": err.Error(),
        })
        return
    }

    id, err := testService.Insert()
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "code": -1,
            "message": "Insert() error!",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "code": 1,
        "message": "success",
        "data": id,
    })

}

// 查询
func GetTestList(c *gin.Context) {
    var testService Services.Test
    ids := c.DefaultQuery("id","0")
    testcol := c.DefaultQuery("testcol","testcol")
	fmt.Println("id=" + ids + ",testcol=" + testcol)

    list, err := testService.GetTestData()
    fmt.Println("res---",list,err)

    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "code": -1,
            "message": "Insert() error!",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "code": 1,
        "message": "success",
        "data": list,
    })

}