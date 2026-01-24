package main

import (
    "fmt"
    "github.com/ph4mished/color"
)

func main() {
    // Simple template with one placeholder
    greeting := color.Parse("[fg=green]Hello, [0][reset]!")
    
    fmt.Println(greeting.Apply("Alice"))
    fmt.Println(greeting.Apply("Bob"))
    fmt.Println(greeting.Apply("World"))
    
    // Complex template with multiple placeholders
    logTemplate := color.Parse("[0] [fg=blue][1][reset]: [fg=yellow][2][reset]")
    
    // Different log levels
    fmt.Println(logTemplate.Apply("[INFO]", "main", "Application started"))
    fmt.Println(logTemplate.Apply("[WARN]", "auth", "Token expiring soon"))
    fmt.Println(logTemplate.Apply("[ERROR]", "db", "Connection failed"))
}
