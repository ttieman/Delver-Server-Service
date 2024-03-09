package User_Model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	UserName    string    `json:"username"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	DOB         time.Time `json:dob`
	FirstName   string    `json:"first_name`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(userPassword, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	return err == nil
}
func NewUser(username, password, firstName, lastName, phoneNumber string, dob time.Time) (*User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:          uuid.New(),
		UserName:    username,
		Password:    hashedPassword,
		CreatedAt:   time.Now(),
		DOB:         dob,
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
	}, nil
}
