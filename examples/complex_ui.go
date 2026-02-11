package main

import (
    "fmt"
    "strings"
    "github.com/ph4mished/color"
)

func main() {
    
    // Table with colored headers
    headerTemplate := color.Parse("[bold fg=cyan][0][reset]")
    rowTemplate := color.Parse("[0]  [fg=yellow][1][reset]  [fg=green][2][reset]")
    
    fmt.Println(headerTemplate.Apply(strings.Repeat("─", 40)))
    fmt.Println(headerTemplate.Apply("USER MANAGEMENT"))
    fmt.Println(headerTemplate.Apply(strings.Repeat("─", 40)))
    
    fmt.Println(rowTemplate.Apply("Alice", "admin", "active"))
    fmt.Println(rowTemplate.Apply("Bob", "user", "active"))
    fmt.Println(rowTemplate.Apply("Charlie", "guest", "inactive"))
    
    // Nested templates
    errorTemplate := color.Parse("[bold fg=red][0][reset]: [1]")
    suggestionTemplate := color.Parse("[fg=yellow]Suggestion: [0][reset]")
    
    errors := []struct{
        code string
        msg string
        suggestion string
    }{
        {"E001", "File not found", "Check the file path"},
        {"E002", "Permission denied", "Run with sudo or check permissions"},
        {"E003", "Out of memory", "Close other applications"},
    }
    
    for _, err := range errors {
        fmt.Println(errorTemplate.Apply(err.code, err.msg))
        fmt.Println("  " + suggestionTemplate.Apply(err.suggestion))
        fmt.Println()
    }
}
