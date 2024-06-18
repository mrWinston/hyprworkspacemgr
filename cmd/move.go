/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	client "github.com/labi-le/hyprland-ipc-client/v3"
	"github.com/mrWinston/hyprworkspacemgr/pkg/grid"
	"github.com/mrWinston/hyprworkspacemgr/pkg/hyprland"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var takeWindow bool

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:   "move [flags] <direction>",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: validateArgs,
	Run:  move,
}

func init() {
	rootCmd.AddCommand(moveCmd)
	moveCmd.Flags().BoolVarP(&takeWindow, "take", "t", false, "Take current window with you")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func validateArgs(cmd *cobra.Command, args []string) error {
	err := cobra.ExactArgs(1)(cmd, args)
	if err != nil {
		return err
	}
	// TODO: Make sure we pass in a valid direction
	return nil
}

func move(_ *cobra.Command, args []string) {
	direction := args[0]
	hl := hyprland.NewClient()

	monitors, err := hl.Monitors()
	if err != nil {
		fmt.Printf("Error listing monitors: %v", err)
		return
	}

	currentWorkspace, err := hl.ActiveWorkspace()
	if err != nil {
		fmt.Printf("Error getting workspace: %v", err)
		return
	}

	currentWindow, err := hl.ActiveWindow()
	if err != nil {
		fmt.Printf("Error getting currentWindow: %v", err)
		return
	}

	err = hyprland.SetAnimation(hl, direction)
	if err != nil {
		fmt.Printf("Error setting animation: %v", err)
		return
	}

	activeWorkspaces := lo.Map(monitors, func(item client.Monitor, _ int) int { return item.ActiveWorkspace.Id })

	nextID := getNextFreeWS(currentWorkspace.Id, direction, activeWorkspaces)

	takeWindowAddress := ""
	if takeWindow {
		takeWindowAddress = currentWindow.Address
	}

	err = hyprland.MoveToWorkspace(hl, nextID, takeWindowAddress)

	if err != nil {
		fmt.Printf("Error Moving: %v", err)
		return
	}
}

func getNextFreeWS(sourceID int, direction string, activeWS []int) int {
	nextID := sourceID
	for {
		fmt.Printf("nextId: %v\n", nextID)
		currentCoordinate := grid.IdxToCoord(nextID)
		fmt.Printf("currentCoordinate: %v\n", currentCoordinate)
		tmpID := grid.NextCoordInDirection(currentCoordinate, direction).ToIdx()
		fmt.Printf("tmpId: %v\n", tmpID)

		if tmpID == nextID {
			break
		}

		nextID = tmpID

		if !lo.Contains(activeWS, nextID) {
			break
		}
	}

  // Don't move to the workspace if it's already occupied. This is an egde case
  if lo.Contains(activeWS, nextID) {
    return sourceID
  }

	return nextID
}
