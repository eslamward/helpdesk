package main

import (
	"fmt"
	"log"
	"os"

	apiserve "github.com/eslamward/helpdesk/pkg/apiServe"
	"github.com/eslamward/helpdesk/pkg/database"
	"github.com/eslamward/helpdesk/pkg/store"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatal("Please Provide username and password od database")
	}
	username := os.Args[1]
	password := os.Args[2]
	fmt.Print(username, password)
	dbConnection := database.NewDatabaseConnection(username, password)
	db, err := dbConnection.Init()

	if err != nil {
		log.Fatal(err.Error())
	}
	router := gin.Default()
	sessionStore := cookie.NewStore([]byte("secret"))

	router.Use(sessions.Sessions("mysession", sessionStore))

	userStore := store.NewUserStore(db)
	ticketStore := store.NewTicketStore(db)
	mainStore := store.NewMainStore(ticketStore, userStore)
	api := apiserve.NewAPIServices(":8080", mainStore)
	api.Serve(router)

}
