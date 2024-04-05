package main

import (
	"fmt"
	"gin_mysql/enum"
	"gin_mysql/model"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CreateUserType struct {
	Id        int        `json:"id" gorm:"column:id"`
	Username  string     `json:"username" gorm:"column:username"`
	Password  string     `json:"password" gorm:"column:password"`
	Role      string     `json:"role" gorm:"default:USER;not null;column:role"`
	Status    string     `json:"status" gorm:"column:status;default:ACTIVE"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
}

type UpdateUserType struct {
	Username  *string    `json:"username,omitempty" gorm:"column:username"`
	Password  *string    `json:"password,omitempty" gorm:"column:password"`
	Role      *string    `json:"role,omitempty" gorm:"default:USER;not null;column:role"`
	Status    *string    `json:"status,omitempty" gorm:"column:status"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

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
		users.GET("/:id", GetUser(db))
		users.POST("", CreateUser(db))
		users.PUT("/:id", UpdateUser(db))
		users.DELETE("/:id", DeleteUser(db))
	}

	r.Run(":8080")
}

func GetUsers(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data []model.Users

		if err := db.Table(string(enum.USERS_TABLE)).Find(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

func GetUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.Users
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Table(string(enum.USERS_TABLE)).Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

func CreateUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data CreateUserType

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Table(string(enum.USERS_TABLE)).Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

func UpdateUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data UpdateUserType
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

		if err := db.Table(string(enum.USERS_TABLE)).Where("id = ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
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

		if err := db.Table(string(enum.USERS_TABLE)).Where("id = ?", id).Delete(nil).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}
