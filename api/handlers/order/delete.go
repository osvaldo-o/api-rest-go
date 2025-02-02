package order

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Delete(c *gin.Context) {

	idParam := c.Param("id")
	if idParam == "" {
		log.Println("error: es necesario el ID")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"data":   "error",
		})
		return
	}

	err := h.OrderService.Delete(&idParam)
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
		"data":   "Pedido eliminada correctamente",
	})
}
