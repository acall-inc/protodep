package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var depCmd = &cobra.Command{
	Use:   "dep",
	Short: "get proto dependencies",
	RunE: func(cmd *cobra.Command, args []string) error {

		isUpdate, err := cmd.Flags().GetBool("update")
		if err != nil {
			return err
		}
		fmt.Println("dep!!")
		fmt.Printf("isUpdate=%v\n", isUpdate)

		p, _ := os.Getwd()
		fmt.Println(p)
		return nil
	},
}

func initDepCmd() {
	depCmd.PersistentFlags().BoolP("update", "u", false, "update locked file and vendors")
}
