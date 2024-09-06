/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// addTaskCmd represents the addTask command
var addTaskCmd = &cobra.Command{
	Use:   "addTask",
	Short: "A brief description of your command",
	Long: `The 'add' command allows you to create a new task and add it to your task list.
Each task can have a title, due date, and priority to help you organize and manage your work efficiently.

Flags:
- --title: (required) The name or description of the task.
- --due: (optional) Set a due date for the task (format: YYYY-MM-DD).
- --priority: (required) Specify the task's priority (low, medium, high).
- --notes: (required) Specify additional notes about the task.

Examples:
- Add a basic task: taskcli add --title "Buy groceries"
- Add a task with a due date: taskcli add --title "Finish report" --due "2024-09-10"
- Add a high-priority task: taskcli add --title "Call client" --priority high

Make task management simple and efficient with Task Manager CLI!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addTask called")
		title := cmd.Flag("title").Value.String()
		priority := cmd.Flag("priority").Value.String()
		dueDate := cmd.Flag("due").Value.String()
		notes := cmd.Flag("notes").Value.String()
		if title == "" && priority == "" && notes == "" {
            fmt.Println("Please provide all required flags: --title, --priority, and --notes")
            os.Exit(1)
        }
		if dueDate == ""{
			dueDate = time.Now().AddDate(0, 0, 7).Format("2006-01-02")
		}

		err := addTask(title, priority, dueDate, notes)

		if err!= nil {
            fmt.Printf("Error adding task: %v\n", err)
            os.Exit(1)
        }
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
	addTaskCmd.PersistentFlags().String("title", "", "The name or description of the task")
	addTaskCmd.PersistentFlags().String("due", "", "Set a due date for the task (format: YYYY-MM-DD)")
	addTaskCmd.PersistentFlags().String("priority", "", "Specify the task's priority(low, medium, high)")
	addTaskCmd.PersistentFlags().String("notes", "", "Specify additional notes about the task")


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Task struct {
	ID          int    		`json:"id"`
	Title 		string 		`json:"title:"`
	Status 		string 		`json:"status:"`
	Priority 	string 		`json:"priority:"`
	DueDate     string      `json:"due_date:"`
	CreateDate 	string		`json:"create_date:"`
	Notes 		string 		`json:"notes:"`
}

var taskFile = "task.json"

func ReadTask() ([]Task, error) {
	var tasks []Task

	// Check if there is any task.json exits
	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		return tasks, nil
	}

	//  Read data from task.json
	data, err := os.ReadFile(taskFile)
	if err!= nil {
        return nil, fmt.Errorf("failed to reading task.json: %v", err)
    }

	// Unmarshal JSON data to struct
	err = json.Unmarshal(data, &tasks)
	if err!= nil {
        return nil, fmt.Errorf("failed to read the task file - %v", err)
    }

	return tasks, nil
}


func WriteTask(tasks []Task) error {
	//  Marshall task into JSON format
	data, err := json.MarshalIndent(tasks, "", " ")
	if err!= nil {
        return fmt.Errorf("error marshalling tasks to JSON: %v", err)
    }

    // Write JSON data to task.json
	err = os.WriteFile(taskFile, data, 0644)
	if err!= nil {
        return fmt.Errorf("error writing task.json: %v", err)
    }

	return nil
}

func addTask(title, priority, dueDate, note string) (error){

	//  Read task file
	tasks, err := ReadTask()
	if err!= nil {
        return err
    }


	//  Auto Generate id
	id := len(tasks) + 1
	createDate := time.Now().Format("2006-01-02")


	newTask := Task{
		ID: id,
		Title: title,
        Status: "pending",
        Priority: priority,
        DueDate: dueDate,
        CreateDate: createDate,
		Notes: note,
	}

	tasks = append(tasks, newTask)

	//  Write task file
	err = WriteTask(tasks)
	if err!= nil {
        return err
    }

    fmt.Printf("Task added successfully: %s\n", newTask.Title)
	return nil
}
