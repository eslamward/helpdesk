package store

type MainStore struct {
	TicketStore TicketStore
	UserStore   UserStore
}

func NewMainStore(ticketStore TicketStore, userStore UserStore) *MainStore {
	return &MainStore{
		TicketStore: ticketStore,
		UserStore:   userStore,
	}
}
