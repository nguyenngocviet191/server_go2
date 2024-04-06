package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var count int = 0

// Model cho các mục dữ liệu
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items []Item

// Lấy tất cả các mục
func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

// Lấy một mục theo ID
func GetItem(c *gin.Context) {
	id := c.Param("id")
	for _, item := range items {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

// Tạo một mục mới
func CreateItem(c *gin.Context) {
	var item Item
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	items = append(items, item)
	c.JSON(http.StatusCreated, item)
}
func update_count() {
	for {
		count = count + 1
		time.Sleep(1 * time.Second)
	}
}
func display_count(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"count2": count})
}

//	func helloHandler(w http.ResponseWriter, r *http.Request) {
//	    fmt.Fprintf(w, "Hello\n")
//	}
//
// Hàm main
func main() {
	// Khởi tạo router
	router := gin.Default()

	// Khởi tạo dữ liệu
	items = append(items, Item{ID: "1", Name: "Item 1"})
	items = append(items, Item{ID: "2", Name: "Item 2"})

	// Định tuyến các endpoints
	router.GET("/count", display_count)
	router.GET("/items", GetItems)
	router.GET("/items/:id", GetItem)
	router.POST("/items", CreateItem)
	go update_count()
	// Khởi động server trên cổng 8000
	router.Run(":8000")
}
