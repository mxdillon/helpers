package helpers

import (
    "testing"
)

// TestIsInString ensures that various permutations of inputs return the desired tuple.
func TestIsInString(t *testing.T) {
    type input struct {
        b []byte
        s string
    }
    type expected struct {
        b bool
        i int
    }

    tests := []struct {
        name string
        i    input
        e    expected
    }{
        {"Slice where element is not present", input{[]byte("c"), "hello max!"}, expected{false, -1}},
        {"Slice where element is present", input{[]byte("m"), "hello max!"}, expected{true, 6}},
        {"Slice where element is present multiple times", input{[]byte("l"), "hello lollipop lady!"}, expected{true, 2}},
        {"Slice with multiple characters", input{[]byte("ce"), "nice to meet you"}, expected{true, 2}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            actualBool, actualInt := IsInString(tt.i.b, tt.i.s)
            if actualBool != tt.e.b {
                t.Errorf("boolean incorrect, want %v, got %v", tt.e.b, actualBool)
            }
            if actualInt != tt.e.i {
                t.Errorf("index incorrect, want %v, got %v", tt.e.i, actualInt)
            }
        })
    }
}

// TestIsInInt ensures that various permutations of inputs return the desired tuple.
func TestIsInInt(t *testing.T) {
    type input struct {
        b int
        s []int
    }
    type expected struct {
        b bool
        i int
    }

    tests := []struct {
        name string
        i    input
        e    expected
    }{
        {"Slice where element is not present", input{34, []int{21, 12, 3, 2, 4, 1}}, expected{false, -1}},
        {"Slice where element is present", input{34, []int{21, 12, 34, 2, 4, 1}}, expected{true, 2}},
        {"Slice where element is present multiple times", input{34, []int{21, 12, 4, 34, 2, 4, 34, 1}}, expected{true, 3}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            actualBool, actualInt := IsInInt(tt.i.b, tt.i.s)
            if actualBool != tt.e.b {
                t.Errorf("boolean incorrect, want %v, got %v", tt.e.b, actualBool)
            }
            if actualInt != tt.e.i {
                t.Errorf("index incorrect, want %v, got %v", tt.e.i, actualInt)
            }
        })
    }
}
