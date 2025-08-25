/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"tasks-cli/intern"

	"github.com/spf13/cobra"
)

// markInProgressCmd represents the markInProgress command
var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("markInProgress called")

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Println("Error during converting to int")
			return
		}

		err = intern.MarkInProgress(id)
		if err != nil {
			log.Println("Error during Marking")
			return
		}

		log.Printf("Task by id = %v was marked as InProgress", id)
	},
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markInProgressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markInProgressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
