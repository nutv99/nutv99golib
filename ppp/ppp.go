package ppp

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book999 struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Book2 struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	gorm.Model
}


type Book111 struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func GetSchoolAddress() string {
    return "กรุงเทพ";
}

func Aaaa() string {
	message := "Hello, "
	return message
}

func listBook2sHandler(c *gin.Context) {
	 
        message := "Hello Book2 "
	c.JSON(http.StatusOK, &message)
}
