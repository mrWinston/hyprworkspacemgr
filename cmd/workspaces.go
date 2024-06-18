/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mrWinston/hyprworkspacemgr/pkg/hyprland"
	"github.com/mrWinston/hyprworkspacemgr/pkg/yuck"
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)

// workspacesCmd represents the workspaces command
var workspacesCmd = &cobra.Command{
	Use:   "workspaces",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
    hl := hyprland.NewClient() 

    out, err := yuck.WsExpression(hl)
    if err != nil {
      log.Errorf("Can't get workspaces: %v", err)
    }

    fmt.Println(out)
	},

}

func init() {
	ewwCmd.AddCommand(workspacesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workspacesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workspacesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
