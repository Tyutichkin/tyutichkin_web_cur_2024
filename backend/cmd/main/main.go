package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"main/internal/middleware"
	"main/internal/repository/sqlite"
	"main/internal/service"
	"time"
)

var (
	host = ":8083"
	DB   *sql.DB
)

func main() {
	err := initDatabase()
	defer DB.Close()
	if err != nil {
		panic(fmt.Sprintf("initDatabase err: %w", err))
	}
	repo := sqlite.NewRepository(DB)
	svc := service.NewService(repo)

	router, err := initApi(svc)
	if err != nil {
		panic(fmt.Sprintf("initApi error: %w", err))
	}

	err = router.Run(host)
	if err != nil {
		panic(fmt.Sprintf("GIN router run err: %w", err))
	}
}

// API V1
func initApi(svc *service.Service) (router *gin.Engine, err error) {
	router = gin.Default()
	config := cors.Config{
		AllowAllOrigins:  true, // Allow all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))

	v1 := router.Group("/api/v1")
	{
		v1.POST("user/login", svc.Login)

		v1.Use(middleware.AuthMiddleware())

		v1.GET("good/all", svc.GetAllGoods)
		v1.GET("stock/all", svc.GetAllStocks)
		v1.GET("user/all", svc.GetAllUsers)

		v1.PUT("good", svc.EditGood)
		v1.PUT("good_stock", svc.EditGoodStock)
		v1.PUT("user", svc.EditUser)
		v1.PUT("stock", svc.EditStock)

		v1.DELETE("stock/:id", svc.DeleteStock)
		v1.DELETE("user/:id", svc.DeleteUser)
		v1.DELETE("good/:id", svc.DeleteGood)

		v1.POST("good", svc.AddGood)
		v1.POST("good_stock", svc.AddGoodStock)
		v1.POST("user", svc.AddUser)
		v1.POST("stock", svc.AddStock)

		v1.POST("search/good", svc.SearchGoods)
		v1.POST("search/user", svc.SearchUsers)
		v1.POST("search/stock", svc.SearchStocks)

		v1.POST("good/upload", svc.UploadGood)
		v1.GET("good/download/:id", svc.DownloadGood)
	}

	return router, nil
}

func initDatabase() error {
	db, err := connectDatabase()
	if err != nil {
		return fmt.Errorf("connectDatabase error: %w", err)
	}
	DB = db
	return nil
}

func connectDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../db.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
