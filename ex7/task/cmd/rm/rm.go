package rm

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.etcd.io/bbolt"
)

// RmCommand: removes a task from the task list
var RmCommand = &cobra.Command{
	Use:   "rm [task to remove]",
	Short: "Remove task from task list",
	Long:  "Remove task from task list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("do command ran!")
		fmt.Println(args[0])

		db, err := bbolt.Open("task.db", 0600, nil)
		if err != nil {
			return err
		}

		db.Update(func(tx *bbolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte("tasks"))
			if err != nil {
				fmt.Println("Error when retriving or creating tasks bucket")
				return err
			}
			if value := b.Get([]byte(args[0])); value != nil {
				err := b.Delete([]byte(args[0]))
				if err != nil {
					return err
				}
				fmt.Println("Task removed from task list")
			} else {
				fmt.Println("No task of that name exists.")
			}
			return err
		})
		return nil
	},
}
