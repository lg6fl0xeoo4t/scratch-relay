package main

import (
    "fmt"
    "os"
    "strings"
)

func clean(parts []string) []string {
    out := make([]string, 0, len(parts))
    for _, part := range parts {
        p := strings.TrimSpace(part)
        if p != "" {
            out = append(out, p)
        }
    }
    return out
}

func main() {
    args := clean(os.Args[1:])
    if len(args) == 0 {
        fmt.Println("scratch-relay: pass one or more tokens")
        return
    }
    fmt.Println(strings.Join(args, " | "))
}
