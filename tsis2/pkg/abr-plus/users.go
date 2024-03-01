package model

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
)

type User struct {
	UserID             int    `json:"userID"`
	UserEmail          string `json:"userEmail"`
	Username           string `json:"username"`
	Password           string `json:"-"`
	NumberOfPhoneUser  string `json:"numberOfPhoneUser"`
	ProfilePictureUser string `json:"profilePictureUser"`
	Role               string `json:"role"`
}

type UserModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

var users = []User{
	{
		UserID:             1,
		UserEmail:          "user1@example.com",
		Username:           "user1",
		Password:           "hashed_password1",
		NumberOfPhoneUser:  "123456789",
		ProfilePictureUser: "user1.jpg",
		Role:               "User",
	},
	{
		UserID:             2,
		UserEmail:          "user2@example.com",
		Username:           "user2",
		Password:           "hashed_password2",
		NumberOfPhoneUser:  "987654321",
		ProfilePictureUser: "user2.jpg",
		Role:               "Admin",
	},
}

func GetUsers() []User {
	return users
}

func GetUserByID(id string) (*User, error) {
	for _, u := range users {
		// Convert id to an integer
		userID, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}

		if u.UserID == userID {
			return &u, nil
		}
	}
	return nil, errors.New("User not found")
}
