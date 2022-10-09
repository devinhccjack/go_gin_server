package UserModels

import (
    "go_server/Databases"
    "time"
)

// Model default model
type Model struct {
    ID        uint       `json:"id" gorm:"primary_key"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}
 
// User 用户表
type User struct {
    Model
    Name     string `json:"name"`
    Password string `json:"password"`
}

func (this *User) Insert() (id int, err error) {
    result := Mysql.DB.Create(&this)
    id = this.ID
    if result.Error != nil {
        err = result.Error
        return
    }
    return
}

