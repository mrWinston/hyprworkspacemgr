package yuck

import (
	"fmt"
	"strconv"
	"strings"
)

type YuckExpression struct {
	Action   string
	Name     string
	Args     []string
	Opts     map[string]any
	Children []any
}

func (ye YuckExpression) String() string {
	sb := strings.Builder{}
	sb.WriteString("(")

	sb.WriteString(ye.Action)
	sb.WriteString(" ")

	if ye.Name != "" {
		sb.WriteString(ye.Name)
		sb.WriteString(" ")
	}

	if len(ye.Args) > 0 {
		sb.WriteString("[")
		sb.WriteString(strings.Join(ye.Args, " "))
		sb.WriteString("]")
		sb.WriteString(" ")
	}

	for key, val := range ye.Opts {
		 if okVal, ok := val.(fmt.Stringer); ok {
		 	sb.WriteString(fmt.Sprintf(":%s %s", key, okVal.String()))
		 } else {
      strVal := fmt.Sprint(val)
		 	sb.WriteString(fmt.Sprintf(":%s %s", key, strconv.Quote(strVal)))
		 }
			// sb.WriteString(fmt.Sprintf(":%s %s", key, fmt.Sprint(val)))
		sb.WriteString(" ")
	}

	for _, val := range ye.Children {
		if okVal, ok := val.(fmt.Stringer); ok {
		  sb.WriteString(okVal.String())
		} else {
		  sb.WriteString(fmt.Sprintf("\"%s\"", fmt.Sprint(val)))
		}
		sb.WriteString(" ")
	}
	sb.WriteString(")")

	return sb.String()
}
