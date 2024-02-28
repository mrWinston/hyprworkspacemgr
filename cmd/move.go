/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	client "github.com/labi-le/hyprland-ipc-client"
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
	Run: move,
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


func move(cmd *cobra.Command, args []string) {
  fmt.Println("move called")

  direction := args[0]
  fmt.Printf("direction: %v\n", direction)
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

  if direction == "left" || direction == "right" {
    err = hyprland.SetAnimation(hl, "slide")
  }else {
    err = hyprland.SetAnimation(hl, "slidevert")
  }
  if err != nil {
    fmt.Printf("Error setting animation: %v", err)
    return 
  }


  activeWorkspaces := lo.Map[client.Monitor, int](monitors, func(item client.Monitor, index int) int { return item.ActiveWorkspace.Id}) 

  // get next in dir, check if taken, else get next until it's the same  
  nextId := currentWorkspace.Id
  for {
    fmt.Printf("nextId: %v\n", nextId)
    currentCoordinate := grid.IdxToCoord(nextId)
    fmt.Printf("currentCoordinate: %v\n", currentCoordinate)
    tmpId := grid.NextCoordInDirection(currentCoordinate, direction).ToIdx()
    fmt.Printf("tmpId: %v\n", tmpId)

    if tmpId == nextId {
      return
    }

    nextId = tmpId

    if ! lo.Contains(activeWorkspaces, nextId) {
      break
    }
  }

  err = hyprland.MoveToWorkspace(hl, nextId)

  if err != nil {
    fmt.Printf("Error Moving: %v", err)
    return 
  }

  if ! takeWindow {
    return
  }
  a := client.Args{}
  a.Push("movetoworkspace")
  a.Push(fmt.Sprintf("%d,address:%s", nextId, currentWindow.Address))
  _, err = hl.Dispatch(a)
  if err != nil {
    fmt.Printf("Error Moving Window: %v", err)
    return 
  }
}
