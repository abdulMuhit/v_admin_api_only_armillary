package model

import (
	"fmt"
	"errors"
	"github.com/bandros/framework"
)

type Users struct {
	Id 			int
	Nama 		string
	Email  		string
	Password 	string
	Status string
}

func Login(email string, password string) (map[string]interface{}, error) {

	db := framework.Database{}
	defer db.Close()
	db.Select("email, password, status, parent, role, nama, id")
	db.From("pengguna")
	db.Where("email", email)
	db.Where("status", 1)
	user, err := db.Row()
	fmt.Println("user login ", user, " err ", err)
 /* 	if err != nil {
		return nil, err
	}  */

	if len(user) == 0 {
		//return nil, err
		return nil, errors.New("Invalid password/user")
	}

	if framework.ValidPassword(password, user["password"].(string)) {
		return user, nil
	} else {
		return nil, errors.New("Invalid password/user")
	}

	
}
