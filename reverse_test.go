package go_stringutil

import "testing"

func TestReverse(t *testing.T) {
    cases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {"Hello, 世界", "界世 ,olleH"},
        {
            "that gum you like is coming back in style",
            "elyts ni kcab gnimoc si ekil uoy mug taht",
        },
        {"", ""},
    }
    for _, c := range cases {
        got := Reverse(c.in)
        if got != c.want {
            t.Error("Reverse(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}
