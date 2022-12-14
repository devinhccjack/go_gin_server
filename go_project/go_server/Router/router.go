package Router

import (
    "github.com/gin-gonic/gin"
    "go_server/Controllers"
    "go_server/Middlewares"
    "go_server/Sessions"
    "github.com/gin-contrib/sessions"
)

func InitRouter() {
    router := gin.Default()
    // 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
    router.Use(Middlewares.Cors())
    // 使用 session(cookie-based)
    router.Use(sessions.Sessions("myyyyysession", Sessions.Store))
    v1 := router.Group("v1")
    {
         v1.POST("/testinsert", Controllers.TestInsert)
         v1.GET("/testlist", Controllers.GetTestList)
    }
    
    router.Run(":8080")
}