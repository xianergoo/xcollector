package api_service

import (
	"log"

	"github.com/gin-gonic/gin"
)

func StartService() error {
	router := gin.Default()

	// 添加处理 OPTIONS 请求的中间件
	router.Use(corsMiddleware())

	// 定义路由
	router.GET("/devices", getDevices)
	router.POST("/devices", createDevice)
	router.DELETE("/devices/:id", deleteDevice)

	// 启动服务
	if err := router.Run(":8080"); err != nil {
		return err
	}

	log.Println("Service started successfully.")
	return nil
}

type Device struct {
	ID        int    `json:"id"`
	IPAddress string `json:"ip_address"`
	GroupNo   int    `json:"group"`
}

var devices = []Device{
	{ID: 1, IPAddress: "192.168.4.223", GroupNo: 1},
	{ID: 2, IPAddress: "192.168.4.179", GroupNo: 8},
}

// getDevices 获取设备列表
func getDevices(c *gin.Context) {
	// 实现获取设备列表的逻辑
	log.Print(devices)
	c.JSON(200, gin.H{"data": devices})
}

// createDevice 创建设备
func createDevice(c *gin.Context) {
	// 实现创建设备的逻辑

	c.JSON(201, gin.H{"message": "Device created"})
}

// deleteDevice 删除设备
func deleteDevice(c *gin.Context) {
	// 实现删除设备的逻辑
	c.JSON(200, gin.H{"message": "Device deleted"})
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 允许任何来源
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
