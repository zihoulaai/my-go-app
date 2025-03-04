package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// APIResponse 统一响应结构
type APIResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Book 数据库模型
type Book struct {
	gorm.Model
	Title  string `json:"title" gorm:"size:255;not null"`
	Author string `json:"author" gorm:"size:255;not null"`
	ISBN   string `json:"isbn" gorm:"size:13;uniqueIndex"`
}

var db *gorm.DB

func main() {
	// 初始化数据库
	initDB()

	router := gin.Default()

	// 全局中间件
	router.Use(corsMiddleware())
	router.Use(loggingMiddleware())

	// API 版本分组
	v1 := router.Group("/api/v1")
	{
		// 书籍路由组
		books := v1.Group("/books")
		{
			books.GET("", listBooks)
			books.POST("", createBook)
			books.GET("/:id", getBook)
			books.PUT("/:id", updateBook)
			books.DELETE("/:id", deleteBook)
		}
	}

	// 健康检查端点
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, APIResponse{
			Code:    http.StatusOK,
			Data:    gin.H{"status": "available"},
			Message: "Service is healthy",
		})
	})

	err := router.Run(":8888")
	if err != nil {
		return
	}
}

func loggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("Request received")
		c.Next()
	}
}

// 自定义中间件示例
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Next()
	}
}

// 数据库操作示例
func createBook(c *gin.Context) {
	var input struct {
		Title  string `json:"title" binding:"required"`
		Author string `json:"author" binding:"required"`
		ISBN   string `json:"isbn" binding:"required,len=13"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid input: " + err.Error(),
		})
		return
	}

	book := Book{Title: input.Title, Author: input.Author, ISBN: input.ISBN}
	result := db.Create(&book)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Database error: " + result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, APIResponse{
		Code:    http.StatusCreated,
		Data:    book,
		Message: "Book created successfully",
	})
}

func deleteBook(context *gin.Context) {
	// 从 URL 参数中获取 ID
	id := context.Param("id")

	// 从数据库中删除记录
	result := db.Delete(&Book{}, id)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Database error: " + result.Error.Error(),
		})
		return
	}

	// 返回成功响应
	context.JSON(http.StatusOK, APIResponse{
		Code:    http.StatusOK,
		Message: "Book deleted successfully",
	})
}

func updateBook(context *gin.Context) {
	// 从 URL 参数中获取 ID
	id := context.Param("id")

	// 从数据库中查找记录
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		context.JSON(http.StatusNotFound, APIResponse{
			Code:    http.StatusNotFound,
			Message: "Book not found",
		})
		return
	}

	// 从请求体中获取更新数据
	var input struct {
		Title  string `json:"title" binding:"required"`
		Author string `json:"author" binding:"required"`
		ISBN   string `json:"isbn" binding:"required,len=13"`
	}

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid input: " + err.Error(),
		})
		return
	}
}

func getBook(context *gin.Context) {
	// 从 URL 参数中获取 ID
	id := context.Param("id")

	// 从数据库中查找记录
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		context.JSON(http.StatusNotFound, APIResponse{
			Code:    http.StatusNotFound,
			Message: "Book not found",
		})
		return
	}

	// 返回查询结果
	context.JSON(http.StatusOK, APIResponse{
		Code: http.StatusOK,
		Data: book,
	})
}

func listBooks(context *gin.Context) {
	// 从数据库中查询所有记录
	var books []Book
	result := db.Find(&books)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Database error: " + result.Error.Error(),
		})
		return
	}

	// 返回查询结果
	context.JSON(http.StatusOK, APIResponse{
		Code: http.StatusOK,
		Data: books,
	})
}

// 初始化数据库连接
func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(&Book{}); err != nil {
		return
	}
}
