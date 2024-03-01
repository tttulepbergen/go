package model

import (
	"context"
	"database/sql"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Admin represents the structure of an admin entity.
type Admin struct {
	AdminID             int    `json:"adminID"`
	AdminEmail          string `json:"adminEmail"`
	AdminName           string `json:"adminName"`
	Password            string `json:"-"`
	NumberOfPhoneAdmin  string `json:"numberOfPhoneAdmin"`
	ProfilePictureAdmin string `json:"profilePictureAdmin"`
	Role                string `json:"role"`
}

// AdminModel represents methods to interact with the "admins" table in the database.
type AdminModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

// Insert inserts a new admin into the database.
func (a AdminModel) Insert(admin *Admin) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO admins (admin_email, adminame, password, number_of_phone_admin, profile_picture_admin, role) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING admin_id
	`
	args := []interface{}{
		admin.AdminEmail,
		admin.AdminName,
		hashedPassword,
		admin.NumberOfPhoneAdmin,
		admin.ProfilePictureAdmin,
		admin.Role,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = a.DB.QueryRowContext(ctx, query, args...).Scan(&admin.AdminID)
	return err
}

// Get retrieves an admin from the database based on the provided admin ID.
func (a AdminModel) Get(adminID int) (*Admin, error) {
	query := `
		SELECT admin_id, admin_email, adminame, number_of_phone_admin, profile_picture_admin, role
		FROM admins
		WHERE admin_id = $1
	`
	var admin Admin
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := a.DB.QueryRowContext(ctx, query, adminID)
	err := row.Scan(
		&admin.AdminID,
		&admin.AdminEmail,
		&admin.AdminName,
		&admin.NumberOfPhoneAdmin,
		&admin.ProfilePictureAdmin,
		&admin.Role,
	)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

// Similar methods for Update and Delete can be added based on your requirements.
