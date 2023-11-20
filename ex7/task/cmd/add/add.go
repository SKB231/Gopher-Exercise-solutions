package add

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.etcd.io/bbolt"
)

// Add command: adds a new task to our list
var AddTaskCommand = &cobra.Command{
	Use:   "add [taskname]",
	Short: "Adds a new task to our list",
	Long:  "Adds a new task to our list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("add command ran!")
		fmt.Println("adding task: ", args[0])
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

			err = b.Put([]byte(args[0]), []byte("false"))

			fmt.Println("Added task: ", args[0])
			return err
		})

		return nil
	},
}
