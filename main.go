package main

import (
	"fmt"
	"gin_mysql/modules/user/model"
	ginuser "gin_mysql/modules/user/transport/gin"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(172.17.0.2:3306)/gin_mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(model.Users{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connected to database successfully...")

	r := gin.Default()

	users := r.Group("api/v1/users")
	{
		users.GET("", ginuser.GetUsers(db))
		users.GET("/:id", ginuser.GetUser(db))
		users.POST("", ginuser.CreateUser(db))
		users.PUT("/:id", ginuser.UpdateUser(db))
		users.DELETE("/:id", ginuser.DeleteUser(db))
	}

	r.Run(":8080")
}
