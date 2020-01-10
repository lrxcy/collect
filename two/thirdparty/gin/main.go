package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Create ... 應該要替換func(c *gin.Context)為實際上希望Create的檔案
	router.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test_Create_with_POST_method",
		})
	})

	// Read ... 一般的GET請求會直接拿到對應的資料
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test_Read_with_GET_method",
		})
	})

	// Update ... 對應到curl裡面的PUT方法
	router.PUT("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test_Update_with_PUT_method",
		})
	})

	// Delete ... 對應到curl裡面的DELETE方法
	router.DELETE("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test_Delete_with_DELETE_method",
		})
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run() // listen and serve on 0.0.0.0:8080
}
