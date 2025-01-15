/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/kar1mov-u/to-do-CLI/cmd"
	"github.com/kar1mov-u/to-do-CLI/db"
)

func main() {

	db.InitDB()
	defer db.DB.Close()

	cmd.Execute()
}
