package yuck

import (
	"fmt"

	client "github.com/labi-le/hyprland-ipc-client/v3"
	"github.com/mrWinston/hyprworkspacemgr/pkg/grid"
)

var classToGlyph map[string]string = map[string]string{
	"kitty":                  "",
	"google-chrome":          "󰊯",
	"Spotify":                "󰓇",
	"Slack":                  "󰒱",
	"com.github.flxzt.rnote": "󱦹",
	"vlc":                    "󰕼",
}

const defaultAppGlyph = "󰣆"

func GetGlyphForAppClass(appClass string) string {
	out, ok := classToGlyph[appClass]
	if !ok {
		return defaultAppGlyph
	}

	return out
}

type workspaceReference struct {
	id     int
	active bool
  glyphs string
}

// WsOverviewWidget contains the state of the workspace widget
type WsOverviewWidget struct {
	workspaces []*workspaceReference
}

// WsExpression returns the yuck expression to render the workspace overview widget
func WsExpression(hl client.IPC) (string, error) {

  workspaces := map[int]*workspaceReference{}
	allSpaces, err := hl.Workspaces()
	if err != nil {
		return "", err
	}

	for _, val := range allSpaces {
    workspaces[val.Id] = &workspaceReference{
    	id:     val.Id,
    	active: false,
    	glyphs: "",
    }
	}

	clients, err := hl.Clients()
	if err != nil {
		return "", err
	}

	for _, client := range clients {
    ws, ok := workspaces[client.Workspace.Id]
    if !ok {
      return "", fmt.Errorf("Client on unknown workspace: %d - %s", client.Workspace.Id, client.Title)
    }
    ws.glyphs += GetGlyphForAppClass(client.Class)
	}

	monitors, err := hl.Monitors()
	if err != nil {
		return "", err
	}

	for _, val := range monitors {
    ws, ok := workspaces[val.ActiveWorkspace.Id]
    if ! ok  {
      return "", fmt.Errorf("monitor with unknown workspace: %d - %s", val.ActiveWorkspace.Id, val.Name)
    }
    ws.active = true
	}

  rootEx := YuckExpression{
  	Action:   "box",
  	Opts:     map[string]any{
      "vexpand": true,
      "hexpand": true,
      "class": "workspace",
      "orientation": "v",
      "space-evenly": true,
    },
  	Children: []any{},
  }
  
  for i := 0; i < 3; i++ {
    row := YuckExpression{
    	Action:   "box",
    	Opts:     map[string]any{
        "class": "workspace_row",
        "orientation": "h",
        "space-evenly": true,
      },
    	Children: []any{},
    }
    for j := 0; j < 3; j++ {

      coord := grid.Coordinate([2]int{i, j})
      wsID := coord.ToIdx()
      ws, ok := workspaces[wsID]
      if ! ok {
        ws = &workspaceReference{id: wsID}
      }
      activeClass := "wsinactive"
      if ws.active {
        activeClass = "wsactive"
      }
      
      wsbut := YuckExpression{
      	Action:   "button",
      	Opts:     map[string]any{
          "class": fmt.Sprintf("wsbut %s", activeClass),
        },
      	Children: []any{ws.glyphs},
      }
      row.Children = append(row.Children, wsbut)
    }
    rootEx.Children = append(rootEx.Children, row)
  }

	return rootEx.String(), nil
}
