package apiserve

import (
	"github.com/eslamward/helpdesk/pkg/auth"
	"github.com/eslamward/helpdesk/pkg/routing"
	"github.com/eslamward/helpdesk/pkg/store"
	"github.com/eslamward/helpdesk/pkg/tickets"
	"github.com/gin-gonic/gin"
)

type APIservices struct {
	Address string
	Store   *store.MainStore
}

func NewAPIServices(address string, store *store.MainStore) *APIservices {
	return &APIservices{
		Address: address,
		Store:   store,
	}
}

func (as APIservices) Serve(router *gin.Engine) error {

	ticketServices := tickets.NewTicketServices(as.Store.TicketStore)
	userServices := auth.NewUserServices(as.Store.UserStore)

	routing.RegisteTicketRouting(ticketServices, router)
	routing.RegisterUserRouting(userServices, router)

	return router.Run(as.Address)
}
