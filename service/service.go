package service

import (
	"fmt"
	"log"
	"os"
	"simple-user-inventory/db"
	"simple-user-inventory/server"
	"simple-user-inventory/server/session"
	"simple-user-inventory/server/utils"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

func runInternal() (
	string,
	string,
	string,
	db.Orm,
	sessions.Store,
) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	orm, err := db.NewOrm()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Do you want record admin?")
	var aCreate string
	fmt.Scan(&aCreate)
	fmt.Print("\r\033[1A")
	if aCreate == "yes" || aCreate == "y" {
		fmt.Println("Admin Name?")
		var aName string
		fmt.Scan(&aName)
		fmt.Print("\r\033[1A")

		fmt.Println("                                                                ")
		fmt.Println("Admin Email?")
		var aEmail string
		fmt.Scan(&aEmail)
		fmt.Print("\r\033[1A")

		fmt.Println("                                                                ")
		fmt.Println("Admin Password?")
		var aPassword string
		fmt.Scan(&aPassword)
		fmt.Print("\r\033[7A")
		fmt.Println("                                                                ")
		fmt.Println("                                                                ")
		fmt.Println("                                                                ")
		fmt.Println("                                                                ")
		fmt.Println("                                                                ")
		fmt.Println("                                                                ")
		fmt.Println("                                                                ")

		err = orm.User().CreateAdmin(aName, aEmail, aPassword)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Print("\r\033[1A")
		fmt.Println("                                                                ")
		fmt.Println("                                                                ")
	}

	if utils.IsDev() {
		log.Println("starting service as development mode")
	}

	name := os.Getenv("SERVICE_NAME")
	if len(name) == 0 {
		log.Fatal("env param SERVICE_NAME is empty")
	}
	ver := os.Getenv("VERSION")
	if len(ver) == 0 {
		log.Fatal("env param VERSION is empty")
	}
	at := os.Getenv("SERVER_LISTEN_AT")
	if len(at) == 0 {
		log.Fatal("env param SERVER_LISTEN_AT is empty")
	}
	secret := os.Getenv("SESSION_SECRET")
	if len(secret) == 0 {
		log.Fatal("env param SESSION_SECRET is empty")
	}

	store := session.NewSessionStroe(secret)

	return name, ver, at, orm, store
}

func Run() {
	server.Run(runInternal())
}
