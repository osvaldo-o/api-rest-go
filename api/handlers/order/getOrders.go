package order

import "github.com/gin-gonic/gin"

func (h Handler) GetAllOrders(c *gin.Context) {

	orders, err := h.OrderService.Get()
	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, orders)

}
