package routing

import (
	"github.com/eslamward/helpdesk/pkg/middelware"
	"github.com/eslamward/helpdesk/pkg/tickets"
	"github.com/gin-gonic/gin"
)

func RegisteTicketRouting(ticketServices *tickets.TicketServices, router *gin.Engine) {

	authorized := router.Group("/ticket")
	authorized.Use(middelware.AuthMiddelware())

	authorized.POST("/create", ticketServices.CreateTicket)
	authorized.GET("/alltickets", ticketServices.GetAllTickets)
}
