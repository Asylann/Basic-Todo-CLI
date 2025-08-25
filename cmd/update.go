/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"tasks-cli/intern"

	"github.com/spf13/cobra"
)

var Id int

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")

		id := Id

		updTask, err := intern.UpdateTask(id, args[0])
		if err != nil {
			log.Println("Error during updating task")
			return
		}

		fmt.Printf("Task by id = %v was changed\n", id)
		fmt.Println(updTask)
	},
}

func init() {
	updateCmd.Flags().IntVarP(&Id, "id", "i", 0, "Task`s id that will be updated")

	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
