package fixture
import "errors"
type Change struct { Action string; Content string }
func Apply(original string, exists bool, prev *Change, next string) (Change, error) {
	if !exists {
		return Change{Action: "create", Content: next}, nil
	}
	action := "update"
	if prev != nil && prev.Action == "create" {
		action = "create"
	}
	return Change{Action: action, Content: next}, nil
}
func Identity(s string) string {return s}
