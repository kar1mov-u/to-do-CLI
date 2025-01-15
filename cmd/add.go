/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/kar1mov-u/to-do-CLI/db"

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

		insertAuery := `INSERT INTO tasks (title,completed,completed_at) VALUES (?,?,CURRENT_TIMESTAMP);`
		if _, err := db.DB.Exec(insertAuery, task, false); err != nil {
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
