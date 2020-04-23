package models

import (
	"fmt"
	"gindemo/DB"
	"log"
)

type User struct {
	Name  string `json:"name"`
	Pwd   string `json:"password"`
	Email string `json:"email"`
	Code  string `json:"code"`
}

func InserUserData(name, pwd, email, code string) bool {
	result, err := DB.Db.Exec("insert into in_user(name, password, email, code) values(?,?,?,?)", name, pwd, email, code)
	if err != nil {
		log.Fatalln("InsertUserData err", err)
		return false
	}
	fmt.Println(result.LastInsertId())
	return true
}
