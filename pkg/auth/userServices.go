package auth

import (
	"net/http"
	"time"

	"github.com/eslamward/helpdesk/models"
	"github.com/eslamward/helpdesk/pkg/store"
	"github.com/eslamward/helpdesk/pkg/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserServices struct {
	store store.UserStore
}

func NewUserServices(store store.UserStore) *UserServices {
	return &UserServices{
		store: store,
	}
}

func (us UserServices) RegisterUser(context *gin.Context) {

	session := sessions.Default(context)

	var user models.User
	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = userValidation(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = password
	registeredUser, err := us.store.RegisterUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	registeredUser.Password = ""
	session.Set("userID", registeredUser.ID)
	session.Set("userEmail", registeredUser.Email)
	session.Save()
	context.JSON(http.StatusCreated, gin.H{"user": registeredUser})

}

func (us UserServices) Login(context *gin.Context) {
	var user models.User
	session := sessions.Default(context)

	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = userValidation(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginedUser, err, validLogin := us.store.Login(user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !validLogin {
		context.JSON(http.StatusBadRequest, gin.H{"error": "email or password are incorrect"})
		return
	}
	loginedUser.Password = ""
	session.Set("userID", loginedUser.ID)
	session.Set("userEmail", loginedUser.Email)
	session.Save()
	context.JSON(http.StatusOK, gin.H{"user": loginedUser})

}

func (us UserServices) Logout(context *gin.Context) {

	//Is User Logged IN

	session := sessions.Default(context)
	isLogedIn := us.isUserLogedIn(session)
	if !isLogedIn {
		context.JSON(http.StatusBadRequest, gin.H{"error": "you already logedout"})
		return
	}
	session.Clear()
	session.Save()
	context.JSON(http.StatusOK, gin.H{"logout": true})
}

func (us UserServices) ResetPassword(context *gin.Context) {

	var resetObj models.ResetPasswordObject

	err := context.ShouldBind(&resetObj)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isExist, err := us.store.EmailAlreadyExists(resetObj.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !isExist {
		context.JSON(http.StatusBadRequest, gin.H{"error": "this email is not exist"})
		return
	}
	changed, err := us.store.ResetPassword(resetObj)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !changed {
		context.JSON(http.StatusBadRequest, gin.H{"message": "password didn't changed"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "password changed successfully"})

}

func (us UserServices) isUserLogedIn(session sessions.Session) bool {

	userEmail := session.Get("userEmail")
	return userEmail != nil
}
