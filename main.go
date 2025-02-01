package main

import (
	"flowers-mago/api/api/handlers/order"
	"flowers-mago/api/api/middleware"
	"flowers-mago/api/internal/repository/mysql_db"
	orderMysql "flowers-mago/api/internal/repository/mysql_db/order"
	orderService "flowers-mago/api/internal/services/order"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

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

	router.Use(middleware.BasicAuthMiddleware)

	router.GET("/orders", orderHandler.GetAllOrders)
	router.GET("/order/:id", orderHandler.GetById)
	router.POST("/order", orderHandler.Create)
	router.PUT("/order/:id", orderHandler.UpdateOrder)
	router.DELETE("/order/:id", orderHandler.Delete)

	router.Run()

}
