package helpers

import(
    "testing"
)

// testIsIn
func testIsIn(t *testing.T) {
	type input struct {
		b byte
        s string
	}
	type expected struct {
		b bool
        i int
	}
	tests := []struct {
		name      string
		i      input
		e      expected
	}{
		{"Slice where element is not present", input{byte("c"), "hello max!"}, expected{false, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualBool, actualInt := IsIn(tt.i.b, tt.i.s)
            if actualBool != tt.e.b{
				t.Errorf("Boolean incorrect, want %v, got %v", tt.e.b, actualBool)
			}
		})
	}
}

