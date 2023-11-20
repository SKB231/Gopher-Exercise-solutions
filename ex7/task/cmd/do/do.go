package do

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.etcd.io/bbolt"
)

// Do command: marks a task as complete
var DoCommand = &cobra.Command{
	Use:   "do [existingTaskName]",
	Short: "Mark task as complete",
	Long:  "Mark task as complete",
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
				if string(value) == "true" {
					fmt.Println("Task already marked complete.")
				} else {
					b.Put([]byte(args[0]), []byte("true"))
					fmt.Println(args[0], " Now marked as complete!")
				}
			} else {
				fmt.Println("No task of that name exists.")
			}
			return err
		})
		return nil
	},
}
