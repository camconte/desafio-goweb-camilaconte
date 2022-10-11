package handler

import (
	"net/http"
	"desafio-goweb-camilaconte/internal/tickets"
	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	service tickets.Service
}

func NewTicketHandler(s tickets.Service) *TicketHandler {
	return &TicketHandler{
		service: s,
	}
}

func (h *TicketHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		tickets, err := h.service.GetAll(c)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return 
		}

		c.JSON(http.StatusOK, tickets)
	}
}

func (h *TicketHandler) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := h.service.GetTotalTickets(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (h *TicketHandler) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := h.service.AverageDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, gin.H{
			"average": avg,
		})
	}
}
