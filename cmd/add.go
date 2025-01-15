/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adding a new task",
	Long:  `Used to add new task`,
	Run: func(cmd *cobra.Command, args []string) {
		task, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Enter a new task").Show()

		db, err := sql.Open("sqlite3", "./tasks.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		// dropQuery := `DROP TABLE tasks;`
		// if _, err := db.Exec(dropQuery); err != nil {
		// 	panic(err)
		// }
		createTableQuery := `
			CREATE TABLE IF NOT EXISTS tasks(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			completed BOOLEAN,
			completed_at DATETIME
			);`

		if _, err := db.Exec(createTableQuery); err != nil {
			panic(err)
		}

		insertAuery := `INSERT INTO tasks (title,completed,completed_at) VALUES (?,?,CURRENT_TIMESTAMP);`
		if _, err := db.Exec(insertAuery, task, false); err != nil {
			panic(err)
		}
		fmt.Println("Task added successfully")

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
