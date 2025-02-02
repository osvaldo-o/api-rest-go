package order

import (
	"flowers-mago/api/internal/domain"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateOrder(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println("error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"data":   "error",
		})
		return
	}

	var order domain.Order
	order.ID = id
	if err := c.ShouldBindJSON(&order); err != nil {
		log.Println("error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"data":   "error",
		})
		return
	}

	if err := h.OrderService.Update(&order); err != nil {
		log.Println("error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"data":   "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   "Pedido actualizado correctamente",
	})

}
