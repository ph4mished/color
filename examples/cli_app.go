package main
import(
    "fmt"
    "os"
    "github.com/ph4mished/color"
)
// Best practice for CLI applications
func main() {
    // Check for --no-color flag
    noColorFlag := false
    for _, arg := range os.Args {
        if arg == "--no-color" {
            noColorFlag = true
            break
        }
    }
    
    // Respect both flag and environment variable
    useColor := !noColorFlag && os.Getenv("NO_COLOR") == ""
    
    // Create toggle
    toggle := color.NewColorToggle(useColor)
    
    // All templates use this toggle
    templates := struct {
        Success color.CompiledTemplate
        Error   color.CompiledTemplate
        Header  color.CompiledTemplate
    }{
        Success: toggle.Parse("[fg=green]✓ [0][reset]"),
        Error:   toggle.Parse("[fg=red]✗ [0][reset]"),
        Header:  toggle.Parse("[bold][0][reset]"),
    }
    
    // Use templates - they'll respect the toggle
    fmt.Println(templates.Header.Apply("My Application"))
    fmt.Println(templates.Success.Apply("Started successfully"))
    
    // If --no-color was used or NO_COLOR is set,
    // outputs will be plain text without escape codes
}
