package yuck

import (
	"testing"

	"github.com/mrWinston/hyprworkspacemgr/pkg/hyprland"
)

func Test_WsExpression(t *testing.T) {
  hl := hyprland.NewClient() 

  out, err := WsExpression(hl)

  if err != nil {
    t.Errorf("WsExpression shouldn't error: %v", err)
    return
  }
  t.Log(out)
}
