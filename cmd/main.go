package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/vikkkichu/study_go/internal/model"
	"github.com/vikkkichu/study_go/internal/task10"
)

func main() {
	task10.String()
}

// var (
// 	rubles          float64 = 1500
// 	rublesToDollars float64 = 92.5
// 	dollars         float64 = rubles / rublesToDollars
// )
// _, _ = fmt.Println(fmt.Sprintf(" %0.2f RUB = %0.2f USD", rubles, dollars))
// var (
//	name        string  = "Вика"
//	age         int     = 26
//	temperature float64 = 25.5
// )
//	// var (
//	name        string  = "Вика"
//	age         int     = 26
//	temperature float64 = 25.5
//)

//fmt.Println(fmt.Sprintf("Привет! Меня зовут %v. Мне %v лет. Сегодня на улице %v градусов.", name, age, temperature))

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
	fileContent, err := os.ReadFile("../testdata/users.json")
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
