package order

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetAllOrders(c *gin.Context) {

	orders, err := h.OrderService.Get()
	if err != nil {
		log.Println("error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   orders,
	})

}
