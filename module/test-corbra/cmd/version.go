package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var name string

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "版本号",
	//Args: cobra.ExactArgs(1),	// 参数个数
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cobrayr command version 1.0.1", args)
		fmt.Println("cobrayr version name ", name)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().StringVarP(&name, "name", "n", "", "version name")
}