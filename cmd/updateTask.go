/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// updateTaskCmd represents the updateTask command
var updateTaskCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateTask called")
		id := cmd.Flag("id").Value.String()
		field := cmd.Flag("field").Value.String()
		newValue := cmd.Flag("new").Value.String() 
		tasks, err := ReadTask()
		if err!= nil {
            fmt.Println("Error reading tasks:", err)
            return

        }
		intId, err := strconv.Atoi(id)
		if err!= nil {
            fmt.Printf("Invalid task ID: %v\n", err)
            return
        }
		updated := updateTask(tasks, intId, field, newValue)
		if updated {
			fmt.Println("Task updated successfully")
		} else {
			fmt.Println("Task not found")
		}
		// Save the updated tasks back to the JSON file
		err = saveTask(tasks)
		if err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
	},
}


func init() {
	rootCmd.AddCommand(updateTaskCmd)
	updateTaskCmd.PersistentFlags().String("id", "", "The name or description of the task")
	updateTaskCmd.PersistentFlags().String("field", "", "Set a due date for the task (format: YYYY-MM-DD)")
	updateTaskCmd.PersistentFlags().String("new", "", "Specify the task's priority(low, medium, high)")
}

func updateTask(tasks []Task, id int, field string, newValue string) bool {
	for i, task := range tasks {
        if task.ID == id {
            switch field {
            case "title":
                tasks[i].Title = newValue
            case "due":
                tasks[i].DueDate = newValue
            case "priority":
                tasks[i].Priority = newValue
			case "notes":
                tasks[i].Notes = newValue
			case "status":
				tasks[i].Status = newValue
            default:
                return false
            }
            return true
        }
    }
    return false
}

func saveTask(tasks []Task) error {
    return WriteTask(tasks)
}
