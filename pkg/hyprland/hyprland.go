package hyprland

import (
	"fmt"
	"os"

	"github.com/labi-le/hyprland-ipc-client/v3"
)

// NewClient creates a new hyprland ipc client
func NewClient() client.IPC {
  return client.MustClient(os.Getenv("HYPRLAND_INSTANCE_SIGNATURE"))
}

// MoveToWorkspace moves the the currently focued window to the workspace with the given id
func MoveToWorkspace(hl client.IPC, workspaceID int, takeWindowAddress string) error {
	args := client.NewByteQueue()
	args.Add([]byte("focusworkspaceoncurrentmonitor"))
	args.Add([]byte(fmt.Sprintf("%d", workspaceID)))

	_, err := hl.Dispatch(args)
	if err != nil {
		return err
	}
	if takeWindowAddress == "" {
		return nil
	}

	a := client.NewByteQueue()
	a.Add([]byte("movetoworkspace"))
	a.Add([]byte(fmt.Sprintf("%d,address:%s", workspaceID, takeWindowAddress)))
	_, err = hl.Dispatch(a)

	return err
}

// SetAnimation sets the workspace switching animation to the given string
func SetAnimation(hl client.IPC, direction string) error {
	animation := "slidevert"
	if direction == "left" || direction == "right" {
		animation = "slide"
	}
	args := client.NewByteQueue()
	args.Add([]byte("animation"))
	args.Add([]byte(fmt.Sprintf("workspaces,1,3,default,%s", animation)))
	return hl.Keyword(args)
}
