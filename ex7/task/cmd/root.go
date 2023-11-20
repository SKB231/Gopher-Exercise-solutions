package cmd

import (
	"fmt"
	"os"

	"github.com/SKB231/gopherex/ex7/task/cmd/add"
	"github.com/SKB231/gopherex/ex7/task/cmd/do"
	"github.com/SKB231/gopherex/ex7/task/cmd/list"
	"github.com/SKB231/gopherex/ex7/task/cmd/rm"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(add.AddTaskCommand)
	rootCmd.AddCommand(list.ListCommand)
	rootCmd.AddCommand(do.DoCommand)
	rootCmd.AddCommand(rm.RmCommand)
}

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a reliable cli task manager for your daily needs!",
	Long:  "Task is a reliable cli task manager for your daily needs!",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Task called!")
		fmt.Println("Args: ")
		fmt.Println(args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
