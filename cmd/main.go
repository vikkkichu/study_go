package main

import (
	"encoding/json"
	"fmt"
	"github.com/vikkkichu/study_go/internal/model"
	"os"
)

func main() {

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
