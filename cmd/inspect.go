/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wizzardich/transcode-cli/v1/internal/pkg/ffmpeg"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:     "inspect PATH",
	Aliases: []string{"i"},
	Short:   "inspect file/directory contents",
	Long: `Show available file tracks and output concise important information about the file.
For example:

inspect file.mkv`,
	Run: func(cmd *cobra.Command, args []string) {
		target := &ffmpeg.Target{Path: args[0]}
		cmd.Println(target.Describe())
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inspectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inspectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
