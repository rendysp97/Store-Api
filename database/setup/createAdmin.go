package setup

import (
	"log"
	"os"
	database "store-api/database/connection"
	"store-api/model"
	"store-api/repository"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func CreatedAdmin() {

	adminPhone := os.Getenv("ADMIN_PHONE")
	adminPass := os.Getenv("ADMIN_PASS")
	adminEmail := os.Getenv("ADMIN_EMAIL")

	user, err := repository.GetUserByPhoneAndEmail(database.Db, adminPhone, adminEmail)

	if err == nil && user != nil {
		log.Println("Admin Already Define")
		return
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(adminPass), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Fail Pass Hash", err)
		return
	}

	admin := model.User{
		Nama:       "Admin",
		Notelp:     adminPhone,
		Email:      adminEmail,
		Kata_sandi: string(hashed),

		Is_admin: true,
	}

	if err := repository.CreateUser(database.Db, &admin); err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			log.Println("Admin Already Define")

			return
		}

		log.Println("Fail Created Admin")
		return

	}

	log.Println("Succeed Created Admin")

}
