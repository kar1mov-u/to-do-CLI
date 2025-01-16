/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/kar1mov-u/to-do-CLI/db"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete <task_id>",
	Short: "used to mark a task as completed",
	Long:  `Use to complete a task`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Please only specify one task")
			return
		}
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task id ,please provide valid integer")
		}

		result, err := db.DB.Exec(`UPDATE tasks SET completed = TRUE, completed_at = CURRENT_TIMESTAMP WHERE id=?`, taskID)
		if err != nil {
			panic(err)
		}

		rowsAffec, err := result.RowsAffected()
		if err != nil {
			panic(err)
		}
		if rowsAffec == 0 {
			fmt.Printf("No task was found with ID %d. \n", taskID)
			return
		}
		fmt.Println("Succesfully completed task")

	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
