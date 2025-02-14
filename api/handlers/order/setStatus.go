package order

import (
	"flowers-mago/api/api/core"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SetStatus(c *gin.Context) {

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

	var body *core.SetStatusBody
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"data":   "error",
		})
		return
	}

	error := h.OrderService.SetStatus(&body.Status, &id)
	if error != nil {
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
