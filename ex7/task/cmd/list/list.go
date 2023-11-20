package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.etcd.io/bbolt"
)

// List command: Lists all remaining tasks
var ListCommand = &cobra.Command{
	Use:   "list",
	Short: "List all remaining tasks",
	Long:  "List all remaining tasks",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("List command ran!")

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

			cursor := b.Cursor()
			for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
				taskName := string(k)
				completed := string(v)

				if completed == "false" {
					fmt.Println("[ ] ", taskName)
				} else {
					fmt.Println("[x] ", taskName)
				}
			}

			return err
		})

		return nil
	},
}
