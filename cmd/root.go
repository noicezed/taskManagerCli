/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task-manager",
	Short: "Add a new task to your task list",
	Long: `Task Manager CLI is a lightweight command-line tool to help you organize and track your tasks.
You can easily add, complete, list, and prioritize tasks using simple commands.

Features:
- Add tasks with a description, due date, and priority
- Mark tasks as completed or pending
- List tasks based on their status or priority
- Edit and delete tasks to keep your task list up-to-date

Example Commands:
- Add a new task: taskcli add --title "Complete report" --due "2024-09-10" --priority high
- List all tasks: taskcli list
- Mark a task complete: taskcli complete 1

Make task management simple and efficient with Task Manager CLI!`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}



func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.task-manager.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

