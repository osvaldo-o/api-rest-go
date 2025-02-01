package order

import (
	"flowers-mago/api/api/core"
	"flowers-mago/api/internal/domain"

	"github.com/gin-gonic/gin"
)

func (h Handler) Create(c *gin.Context) {

	var orderCreateParams core.OrderCreateParams

	if err := c.ShouldBindJSON(&orderCreateParams); err != nil {
		c.JSON(400, err)
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

	err := h.OrderService.Create(order)
	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, "Success")

}
