/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/itozll/iflag/examples/iflag-test/cmd/options"

	"github.com/spf13/cobra"
)

// fooCmd represents the foo command
var fooCmd = &cobra.Command{
	Use:   "foo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("foo called")

		fmt.Println("verbose =", options.Verbose)
		fmt.Println("verbose =", options.OVerbose.Value())

		fmt.Println("toggle =", options.OToggle.Value())
	},
}

func init() {
	rootCmd.AddCommand(fooCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fooCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fooCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
