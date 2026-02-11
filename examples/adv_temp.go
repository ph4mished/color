package main

import (
    "fmt"
    "time"
    "github.com/ph4mished/color"
)

func main() {
    // Status indicator with conditional colors
    statusTemplate := color.Parse("[0] [1][reset]")
    
    items := []struct{
        name string
        status string
    }{
        {"Database", "Online"},
        {"API Server", "Offline"},
        {"Cache", "Degraded"},
    }
    
    for _, item := range items {
        var statusColor string
        switch item.status {
        case "Online":
            statusColor = "[fg=green bold]"
        case "Offline":
            statusColor = "[fg=red bold]"
        default:
            statusColor = "[fg=yellow]"
        }
        
        statusColored := color.Parse(statusColor + item.status).Apply()
        fmt.Println(statusTemplate.Apply(item.name + ":", statusColored))
    }
    
    // Progress bar template
    progressTemplate := color.Parse("[fg=cyan][0][reset]/[fg=cyan][1][reset] [fg=green][2][reset]%")
    
    total := 100
    for i := 0; i <= total; i += 10 {
        percent := i * 100 / total
        fmt.Printf("\r%s", progressTemplate.Apply(i, total, percent))
        time.Sleep(100 * time.Millisecond)
    }
    fmt.Println()
}
