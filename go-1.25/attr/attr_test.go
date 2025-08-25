package attr

import "testing"

// BEGIN ATTR OMIT
func TestAttributes(t *testing.T) {
	t.Attr("subsystem", "foo")
	t.Attr("important", "maybe")
}

// END ATTR OMIT
