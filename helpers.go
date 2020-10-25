package helpers


// IsIn checks whether an element is present within a string. Change the types to slices as required.
// Returns a bool and the index of where the match occured (which can be useful).
func IsIn(b byte, s string) (bool, int) {
    for i := range s {
        if s[i] == b {
            return true, i
        }
    }
    return false, -1
}

