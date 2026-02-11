package main

import (
    "fmt"
    "github.com/ph4mished/color"
)

func main() {
    // Template for showing validation errors
    validationTemplate := color.Parse("[fg=red]• [0]: [1][reset]")
    
    errors := map[string]string{
        "username": "Must be at least 3 characters",
        "email":    "Invalid email format",
        "password": "Must contain uppercase and numbers",
    }
    
    fmt.Println(color.Parse("[bold fg=yellow]Validation Errors:[reset]").Apply())
    for field, message := range errors {
        fmt.Println(validationTemplate.Apply(field, message))
    }
    
    // Template with conditional formatting
    scoreTemplate := color.Parse("[0]: [1]")
    
    scores := []struct{
        name string
        score int
    }{
        {"Alice", 95},
        {"Bob", 75},
        {"Charlie", 45},
        {"Diana", 60},
    }
    
    for _, s := range scores {
        var scoreColor string
        switch {
        case s.score >= 90:
            scoreColor = "[fg=green bold]"
        case s.score >= 70:
            scoreColor = "[fg=yellow]"
        default:
            scoreColor = "[fg=red]"
        }
        
        coloredScore := color.Parse(scoreColor + fmt.Sprint(s.score)+ "[reset]").Apply()
        fmt.Println(scoreTemplate.Apply(s.name, coloredScore))
    }
}
