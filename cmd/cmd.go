package cmd

import api_service "xCollector/internal/pkg/apiserivce"

// "github.com/yourusername/xCollector/internal/pkg/database"

// // InitDatabase 初始化数据库
// func InitDatabase() error {
// 	// 连接到数据库
// 	db, err := database.Connect()
// 	if err != nil {
// 		return err
// 	}

// 	// 执行数据库初始化操作
// 	err = db.Init()
// 	if err != nil {
// 		return err
// 	}

// 	log.Println("Database initialized successfully.")
// 	return nil
// }

func Run() {
	api_service.StartService()
}
