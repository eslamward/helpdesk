package store

import (
	"database/sql"
	"fmt"

	"github.com/eslamward/helpdesk/models"
)

type TicketStore interface {
	CreateTicket(ticket models.Ticket) (models.Ticket, error)
	GetAllTickets() ([]models.Ticket, error)
}

type TicketStorage struct {
	db *sql.DB
}

func NewTicketStore(db *sql.DB) *TicketStorage {
	return &TicketStorage{
		db: db,
	}
}

func (ts *TicketStorage) CreateTicket(ticket models.Ticket) (models.Ticket, error) {

	var id int
	statement := `INSERT INTO tickets(
		title,category,description,status,client,assign,created_at,updated_at
	)VALUES($1,$2,$3,$4,$5,$6,$7,$8)RETURNING id`

	err := ts.db.QueryRow(statement,
		ticket.Title,
		ticket.Category,
		ticket.Description,
		ticket.Status,
		ticket.Client,
		ticket.Assign,
		ticket.CreatedAt,
		ticket.UpdatedAt,
	).Scan(&id)

	if err != nil {
		return ticket, err
	}
	ticket.ID = id
	return ticket, nil
}

func (ts *TicketStorage) GetAllTickets() ([]models.Ticket, error) {
	var tickets []models.Ticket

	statement := "SELECT * FROM tickets"

	rows, err := ts.db.Query(statement)

	if err != nil {
		fmt.Println("Error Query")
		return nil, err

	}
	defer rows.Close()
	for rows.Next() {
		var ticket models.Ticket
		err := rows.Scan(
			&ticket.ID,
			&ticket.Title,
			&ticket.Category,
			&ticket.Description,
			&ticket.Status,
			&ticket.Client,
			&ticket.Assign,
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		tickets = append(tickets, ticket)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("error AfterScan")
		return nil, err
	}
	return tickets, nil
}
