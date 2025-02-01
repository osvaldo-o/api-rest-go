package main

import (
	"flowers-mago/api/api/handlers/order"
	"flowers-mago/api/internal/repository/mysql_db"
	orderMysql "flowers-mago/api/internal/repository/mysql_db/order"
	orderService "flowers-mago/api/internal/services/order"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	db, err := mysql_db.ConnectClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	orderRepo := &orderMysql.Repository{
		DB: db,
	}

	orderSrv := &orderService.Service{
		Repo: orderRepo,
	}

	orderHandler := &order.Handler{
		OrderService: orderSrv,
	}

	envError := godotenv.Load()
	if envError != nil {
		fmt.Println("Error cargando el archivo .env")
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/orders", orderHandler.GetAllOrders)
	router.GET("/order/:id", orderHandler.GetById)
	router.POST("/order", orderHandler.Create)
	router.PUT("/order/:id", orderHandler.UpdateOrder)
	router.DELETE("/order/:id", orderHandler.Delete)

	router.Run()

}
