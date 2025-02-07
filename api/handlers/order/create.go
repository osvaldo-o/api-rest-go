package order

import (
	"flowers-mago/api/api/core"
	"flowers-mago/api/internal/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) Create(c *gin.Context) {

	var orderCreateParams core.OrderCreateParams

	if err := c.ShouldBindJSON(&orderCreateParams); err != nil {
		log.Println("error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"data":   "error",
		})
		return
	}

	order := &domain.Order{
		Name:          orderCreateParams.Name,
		Phone:         orderCreateParams.Phone,
		DeliveryDate:  orderCreateParams.DeliveryDate,
		DeliveryTime:  orderCreateParams.DeliveryTime,
		PlaceDelivery: orderCreateParams.PlaceDelivery,
		Description:   orderCreateParams.Description,
		Price:         orderCreateParams.Price,
		Comment:       orderCreateParams.Comment,
	}

	order, err := h.OrderService.Create(order)
	if err != nil {
		log.Println("error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"data":   "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   order,
	})

}
