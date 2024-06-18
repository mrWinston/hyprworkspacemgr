package yuck

import (
	"fmt"
	"testing"
)

//

// (defwindow wsover
// 	:geometry (geometry
// 						:width "400px"
// 						:height "300px"
//             :anchor "center")
// 	:monitor 0
//   :exclusive false
//   :stacking "fg"
// (wsgrid))
//

func Test_yuck1(t *testing.T) {
  e1 := YuckExpression{
  	Action:   "defwindow",
  	Name:     "wsover",
  	Args:     []string{},
  	Opts:     map[string]any{
      "geometry": YuckExpression{
        Action: "geometry",
        Opts: map[string]any{
          "width": "400px",
          "height": "300px",
          "anchor": "center",
        },
      },
      "monitor": 0,
      "exclusive": false,
      "stacking": "fg",
    },
  	Children: []any{
      YuckExpression{
        Action: "wsgrid",
      },
    },
  }
  out:= e1.String()

  fmt.Println(out)
}
