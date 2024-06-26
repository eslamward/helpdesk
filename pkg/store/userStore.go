package store

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/eslamward/helpdesk/models"
	"github.com/eslamward/helpdesk/pkg/utils"
)

type UserStore interface {
	RegisterUser(user models.User) (models.User, error)
	Login(user models.User) (models.User, error, bool)
	EmailAlreadyExists(string) (bool, error)
	GetUserByEmail(string) (models.User, error)
	ResetPassword(models.ResetPasswordObject) (bool, error)
}

type UserStorage struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (us *UserStorage) RegisterUser(user models.User) (models.User, error) {

	isUserExist, err := us.EmailAlreadyExists(user.Email)
	if err != nil {
		return user, err
	}
	if isUserExist {
		return user, errors.New("this email already exists")
	}

	var id int
	statement := `INSERT INTO users(
		email,
		password,
		type,
		created_at,
		updated_at)
		VALUES($1,$2,$3,$4,$5)RETURNING id`
	err = us.db.QueryRow(statement,
		user.Email,
		user.Password,
		user.Type,
		user.CreatedAt,
		user.UpdatedAt).Scan(&id)
	if err != nil {
		return user, err
	}
	user.ID = id
	return user, nil
}

func (us UserStorage) Login(user models.User) (models.User, error, bool) {

	emailExist, err := us.EmailAlreadyExists(user.Email)
	if err != nil {
		return user, err, false
	}
	if !emailExist {
		return user, errors.New("your aren't registered yet"), false
	}

	fetshedUser, err := us.GetUserByEmail(user.Email)
	if err != nil {
		return user, err, false

	}
	vaild := utils.ComparePassword(fetshedUser.Password, user.Password)

	return fetshedUser, nil, vaild

}

func (us UserStorage) ResetPassword(resetObj models.ResetPasswordObject) (bool, error) {

	user, err := us.GetUserByEmail(resetObj.Email)
	if err != nil {
		return false, err
	}
	oldPasswordMatch := utils.ComparePassword(user.Password, resetObj.OldPassword)
	if !oldPasswordMatch {
		return false, errors.New("the old password didn't match")
	}
	if resetObj.NewPassword != resetObj.ReNewPassword {
		return false, errors.New("the new password didn't match ")
	}

	password, err := utils.HashPassword(resetObj.NewPassword)
	if err != nil {
		return false, err
	}

	statement := `UPDATE users SET password = $1 , updated_at = $2 where id = $3`

	res, err := us.db.Exec(statement, password, time.Now(), user.ID)

	if err != nil {
		return false, err
	}
	rowAffected, err := res.RowsAffected()

	if err != nil {
		return false, err
	}

	if rowAffected == 0 {
		return false, errors.New("the password didn't change")
	}

	return true, nil
}

func (us *UserStorage) EmailAlreadyExists(email string) (bool, error) {

	statement := "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"

	var exist bool

	row := us.db.QueryRow(statement, email)

	err := row.Scan(&exist)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	fmt.Println(exist)
	return exist, nil

}
func (us *UserStorage) GetUserByEmail(email string) (models.User, error) {
	var fetshedUser models.User

	statement := `SELECT * FROM users WHERE email = $1`

	row := us.db.QueryRow(statement, email)

	err := row.Scan(
		&fetshedUser.ID,
		&fetshedUser.Email,
		&fetshedUser.Password,
		&fetshedUser.Type,
		&fetshedUser.CreatedAt,
		&fetshedUser.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return fetshedUser, errors.New("this email not registered")
		}
		return fetshedUser, err
	}
	return fetshedUser, nil
}
