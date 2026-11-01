package fixture
type Change struct { Action string; Content string }
func Apply(original string, exists bool, prev *Change, next string) (Change, error) {
	if !exists {
		return Change{Action: "create", Content: next}, nil
	}
	if prev != nil && prev.Action == "create" {
		return Change{Action: "create", Content: next}, nil
	}
	return Change{Action: "update", Content: next}, nil
}
func Identity(s string) string {return s}
