package main

import (
	"github.com/fatih/color"

	"github.com/alfahadmunna/gopractice/auth"
)

func main() {
	auth.LoginWithCredentials("fahad", "password")

	session := auth.GetSession()
	// fmt.Println("session", session)
	color.Red(session)
}
