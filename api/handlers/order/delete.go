package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Delete(c *gin.Context) {

	id := c.Param("id")

	err := h.OrderService.Delete(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la orden"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Orden eliminada correctamente"})
}
