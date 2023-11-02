package main

import "gobank/config"

func main() {
	server := NewAPIServer(":3000")

	_, err := config.NewMySQLDB()

	if err != nil {
		return
	}

	server.Run()
}
