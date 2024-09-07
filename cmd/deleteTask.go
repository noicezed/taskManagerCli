/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteTaskCmd represents the deleteTask command
var deleteTaskCmd = &cobra.Command{
	Use:   "deleteTask",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteTask called")
		id, _ := cmd.Flags().GetInt("id")
		deleteTask(id)
	},
}

func init() {
	rootCmd.AddCommand(deleteTaskCmd)
	deleteTaskCmd.Flags().IntP("id", "i", 0, "Task ID")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deleteTask(id int){
	task, err := ReadTask()
	if err != nil {
        fmt.Println("Error reading task:", err)
        return
    }
	if id <= 0 {
        fmt.Println("Invalid task ID")
        return
    }
	found := false
	for i, t := range task {
		if t.ID == id {
            task = append(task[:i], task[i+1:]...)
            found = true
			break
        }
	}
	if!found {
        fmt.Println("Task not found")
        return
    }
	err = WriteTask(task)
	if err != nil {
        fmt.Println("Error writing task:", err)
        return
    }
	fmt.Println("Task with ID: %d deleted successfully", id)
}