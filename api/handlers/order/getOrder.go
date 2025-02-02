package order

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetById(c *gin.Context) {

	idParam := c.Param("id")
	if idParam == "" {
		log.Println("error: es necesario el ID")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"data":   nil,
		})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println("error: id invalido")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"data":   nil,
		})
		return
	}

	order, err := h.OrderService.GetById(&id)
	if err != nil {
		log.Println("error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   order,
	})

}
