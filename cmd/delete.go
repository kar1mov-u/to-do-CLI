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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  `Used to delete a task`,
	Run: func(cmd *cobra.Command, args []string) {
		flag, _ := cmd.Flags().GetBool("all")
		if flag {
			db.DB.Exec("DELETE FROM tasks")
			fmt.Println("All tasks have been deleted")
		} else {
			if len(args) != 1 {
				fmt.Println("Pleace only specify 1 ID")
				return
			}

			taskID, _ := strconv.Atoi(args[0])
			res, err := db.DB.Exec("DELETE FROM tasks WHERE tasks.id == ?", taskID)
			if err != nil {
				panic(err)
			}
			rowsAffec, err := res.RowsAffected()
			if err != nil {
				panic(err)
			}
			if rowsAffec == 0 {
				fmt.Printf("No task was found with ID %d. \n", taskID)
				return
			}
			fmt.Println("Taks successfuly deleted")

		}

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	deleteCmd.Flags().BoolP("all", "a", false, "delete all")
}
