package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kar1mov-u/to-do-CLI/db"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list tasks",
	Long:  `Used to list tasks from the DB, can be used with different flags to different options`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		query := "SELECT id,title, completed, completed_at FROM tasks"
		if !all {
			query = "SELECT id,title, completed, completed_at FROM tasks WHERE tasks.completed = False"

		}
		rows, err := db.DB.Query(query)
		if err != nil {
			fmt.Printf("Failed to fetch tasks :%v \n", err)
			return
		}

		defer rows.Close()
		data := [][]string{{"ID", "Title", "Completed", "Completed At"}}
		// fmt.Println("Tasks: ")
		for rows.Next() {
			var id int
			var title string
			var complete bool
			var completed_at *time.Time
			err := rows.Scan(&id, &title, &complete, &completed_at)
			if err != nil {
				fmt.Printf("Error on scanning row %v\n", err)
			}

			status := "Incomplete"
			if complete {
				status = "Completed"
			}
			idStr := strconv.Itoa(id)
			var completedAt_Str string
			if complete {
				completedAt_Str = completed_at.Format("2006-01-02 15:04:05")
				// fmt.Printf("%d. %s [%s] %v \n", id, title, status, completed_at)
			} else {
				completedAt_Str = "N/A"
				// fmt.Printf("%d. %s [%s]\n", id, title, status)
			}
			data = append(data, []string{idStr, title, status, completedAt_Str})
		}
		pterm.DefaultTable.WithHasHeader(true).WithData(data).Render()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listCmd.Flags().BoolP("all", "a", false, "Get all tasks")
}
