// Package stringutil contains utility functions for working with strings.
package stringutil

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
    r := []rune(s)
    // from the start and end of the string, increment until we reach the middle
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        // swap elements each time we increment
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
