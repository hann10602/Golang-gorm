package main

import (
	"fmt"
	"gin_mysql/common"
	"gin_mysql/enum"
	"gin_mysql/modules/user/model"
	ginuser "gin_mysql/modules/user/transport/gin"
	"log"
	"net/http"
	"strconv"

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
		users.GET("", GetUsers(db))
		users.GET("/:id", ginuser.GetUser(db))
		users.POST("", ginuser.CreateUser(db))
		users.PUT("/:id", UpdateUser(db))
		users.DELETE("/:id", DeleteUser(db))
	}

	r.Run(":8080")
}

func GetUsers(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		var data []model.Users

		if err := db.Table(enum.USERS_TABLE).Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Table(enum.USERS_TABLE).Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}

func UpdateUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.UpdateUserDTO
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Table(enum.USERS_TABLE).Where("id = ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func DeleteUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		// if err := db.Table(enum.USERS_TABLE).Where("id = ?", id).Delete(nil).Error; err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"error": err.Error(),
		// 	})

		// 	return
		// }

		if err := db.Table(enum.USERS_TABLE).Where("id = ?", id).Updates(map[string]interface{}{
			"status": "DELETED",
		}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
