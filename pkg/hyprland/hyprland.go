package hyprland

import (
	"fmt"
	"os"

	"github.com/labi-le/hyprland-ipc-client"
)

func NewClient() client.IPC {
  return client.NewClient(os.Getenv("HYPRLAND_INSTANCE_SIGNATURE"))
}

func MoveToWorkspace(hl client.IPC, workspaceId int) error {
  args := client.Args{} 
  args.Push("focusworkspaceoncurrentmonitor")
  args.Push(fmt.Sprintf("%d", workspaceId))
  
  _, err := hl.Dispatch(args)
  return err
}

func SetAnimation(hl client.IPC, animation string) error {
	args := client.Args{}
	args.Push("animation")
  args.Push(fmt.Sprintf("workspaces,1,3,default,%s", animation))
  return hl.Keyword(args)
}
