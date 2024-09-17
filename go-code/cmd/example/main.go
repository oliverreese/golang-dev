package main

import (
	"fmt"
	"os"

	"github.com/oliverreese/golang-dev/pkg/database"
    "github.com/oliverreese/golang-dev/pkg/repository"
	"github.com/oliverreese/golang-dev/pkg/service"
)

func main() {
    db, err := database.NewDatabase()
    if(err != nil) {
        fmt.Println(err)
    }
    defer db.DB.Close()

    messageService := &service.MessageService{}
    fmt.Println(messageService.GetMessage())

    UserRepository := repository.NewUserRepository(&db)
    UserService := service.NewUserService(UserRepository)

    user, err := UserService.GetOneUserById(1)
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Printf("%+v\n",user)
}

func init() {
    os.Setenv("TZ", "Europe/Berlin")
    fmt.Println("Start the machine ...")
}