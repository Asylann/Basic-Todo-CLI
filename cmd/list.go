/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"tasks-cli/intern"
	"tasks-cli/model"
	"time"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")

		var lt []model.Task
		if Filter == "todo" {
			lt = intern.GetListOfTasksTodo()
		} else if Filter == "in-progress" {
			lt = intern.GetListOfTasksInProgress()
		} else if Filter == "done" {
			lt = intern.GetListOfTasksDone()
		} else if Filter == "all" {
			lt = intern.GetListOfTasks()
		} else {
			log.Println("Wrong filter, message the system admin")
			return
		}

		for _, v := range lt {
			fmt.Printf("ID = %v || ", v.Id)
			fmt.Printf("Description : %v || ", v.Description)
			fmt.Printf("CreatedAt : %v || ", v.CreatedDate.Format(time.RFC850))
			fmt.Printf("UpdatedAt : %v || ", v.UpdatedDate.Format(time.RFC850))
			var status string = "todo"
			if v.IsProgress {
				status = "in-progress"
			} else if v.IsDone {
				status = "done"
			}
			fmt.Printf("Status : %v", status)
			fmt.Println()
		}

		log.Println("All Task were received !!!")
	},
}

var Filter string

func init() {
	listCmd.Flags().StringVarP(&Filter, "filter", "f", "all", "Filter for filtering the output")

	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
