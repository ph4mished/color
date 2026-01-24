package main

import (
    "fmt"
    "github.com/ph4mished/color"
)

func main() {
    // Parse and use color codes directly
    red := color.ParseColor("fg=red")
    bold := color.ParseColor("bold")
    reset := color.ParseColor("reset")
    
    fmt.Printf("%sThis is red and bold!%s\n", red + bold, reset)

    // Check if a color is supported
    if color.IsSupportedColor("fg=#FF0000") {
        fmt.Println("Hex colors are supported!")
    }
    
    // Or use the main functions
    color.Parse("[fg=blue]Hello in blue![reset]").Apply()
    color.Parse("[bg=yellow fg=black bold]Bold black text on yellow background.[reset]").Apply()


    //Or pre-parse the color template with placeholders for reuse. This is the heart of the library's performance.

    // Parse once
    template := color.Parse("[fg=red bold]Error: [0][reset]")

    // Reuse multiple times
    fmt.Println(template.Apply("File not found"))
    fmt.Println(template.Apply("Permission denied"))
    fmt.Println(template.Apply("Network timeout"))
}
