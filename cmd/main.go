package main

import (
	"encoding/json"
	"fmt"
	"github.com/vikkkichu/study_go/internal/model"
	"log"
	"os"
)

func main() {
	secondUser()
}

func firstUser() {
	fileContent, err := os.ReadFile("testdata/user.json")
	if err != nil {
		panic(err)
	}

	user := &model.User{}
	err = json.Unmarshal(fileContent, user)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}

func secondUser() {
	fileContent, err := os.ReadFile("testdata/users.json")
	if err != nil {
		panic(err)
	}
	users := []*model.User{}
	err = json.Unmarshal(fileContent, &users)
	if err != nil {
		panic(err)
	}
	users = append(users, &model.User{Name: "Akile", Password: "30"})
	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}
