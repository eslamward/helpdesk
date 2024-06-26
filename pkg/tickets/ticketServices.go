package tickets

import (
	"net/http"
	"time"

	"github.com/eslamward/helpdesk/models"
	"github.com/eslamward/helpdesk/pkg/store"
	"github.com/gin-gonic/gin"
)

type TicketServices struct {
	ticketStore store.TicketStore
}

func NewTicketServices(ticketStore store.TicketStore) *TicketServices {
	return &TicketServices{
		ticketStore: ticketStore,
	}
}

func (ts TicketServices) CreateTicket(context *gin.Context) {

	var ticket models.Ticket
	err := context.ShouldBind(&ticket)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO getSignnedUser function ticket.id = returned id

	ticket.Client = 1
	ticket.CreatedAt = time.Now()
	ticket.UpdatedAt = time.Now()
	ticket.Status = "Opened"

	inertedTicket, err := ts.ticketStore.CreateTicket(ticket)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"ticket": inertedTicket})
}

func (ts TicketServices) GetAllTickets(context *gin.Context) {

	tickets, err := ts.ticketStore.GetAllTickets()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"tickets": tickets})

}
