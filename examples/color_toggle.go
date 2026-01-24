package main

import (
    "fmt"
    "os"
    "github.com/ph4mished/color"
)

func main() {
    // Create color toggle - respects NO_COLOR env var and when output is redirected by default
    toggle := color.NewColorToggle()
    
    // Parse templates using the toggle
    successTemplate := toggle.Parse("[fg=green]✓ [0][reset]")
    errorTemplate := toggle.Parse("[fg=red]✗ [0][reset]")
    
    // These will only show colors if appropriate
    fmt.Println(successTemplate.Apply("Operation completed"))
    fmt.Println(errorTemplate.Apply("Operation failed"))
    
    // Manual control
    //forceColors := color.NewColorToggle(true)   // Always show colors
    //noColors := color.NewColorToggle(false)     // Never show colors
    
    // Use in CLI applications
    useColor := os.Getenv("NO_COLOR") == ""
    appToggle := color.NewColorToggle(useColor)
    
    helpTemplate := appToggle.Parse("[bold fg=cyan][0][reset] [fg=green][1][reset]")
    fmt.Println(helpTemplate.Apply("Usage:", "myapp [options]"))
}
