package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

/**
参考：https://www.jianshu.com/p/63dd2075eb22
command subcommand [para1 para2 ...] [--flag flagvalue | -shortflag flagvalue]
 */
var rootCmd = &cobra.Command{
	Use:                        "cobrayr",
	Aliases:                    []string{"goyr1", "goyr2"},
	SuggestFor:                 []string{"goyr1-suggest", "goyr2-suggest"},
	Short:                      "cobrayr first command",
	Long:                       "cobrayr first command detail description",
	Example:                    "cobrayr list",
	Deprecated:                 "v1.0 has upgrade",
	Annotations:                nil,
	Version:                    "v1.0",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("cobra root cmd execute error")
	}
}
