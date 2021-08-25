package main

import (
	"fmt"

	"github.com/xruterx/golang/lesson10/adress_book"
)

func main() {
	db, err := adress_book.ConnToDB("mongodb://127.0.0.1:27017")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = db.AddContact("Ann", "111-111-111", "first")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = db.AddContact("BOb", "111-111-111", "first")
	if err != nil {
		fmt.Println(err)
		return
	}
	con1, err := db.AddContact("George", "111-111-111", "second")
	if err != nil {
		fmt.Println(err)
		return
	}
	con2, err := db.AddContact("Ashley", "111-111-111", "third")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = db.AddContact("Martin", "111-111-111", "second")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.AssignToGroup(con1, "third")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.AssignToGroup(con2, "second")
	if err != nil {
		fmt.Println(err)
		return
	}
	conList, err := db.ListAllByGroup("first")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conList)
	conList, err = db.ListAllByGroup("second")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conList)
	conList, err = db.ListAllByGroup("third")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conList)
}
