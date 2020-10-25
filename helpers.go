package helpers


// IsInString checks whether a substring is present within a string. Substring can be >1 character.
// Returns a bool and the index of where the first match occurred (which can be useful).
func IsInString(b []byte, s string) (bool, int) {
    for i := range s {
        if s[i] == b[0] {
            return true, i
        }
    }
    return false, -1
}

