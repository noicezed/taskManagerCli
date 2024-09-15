/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)
var all bool // declare the flag

// listTaskCmd represents the listTask command
var listTaskCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("listTask called")
		if all{
			tasks := listTask()
            for _, task := range tasks {
                fmt.Printf("Task ID: %v, Title: %s, Priority: %s, Notes: %s, Due Dates: %s, Status: %s\n", task.ID, task.Title, task.Priority, task.Notes, task.DueDate, task.Status)
            }
			// fmt.Println("All tasks listed.", tasks)
		}
	},
}


func init() {
	rootCmd.AddCommand(listTaskCmd)
	// listTaskCmd.PersistentFlags().String("all","","List all tasks in the cluster and their associated tasks")
	listTaskCmd.Flags().BoolVarP(&all, "all", "a", true, "default: flase.")
}

func listTask()([]Task){

	tasks, err := ReadTask()
	if err != nil {
		return nil
	}
	return tasks
}