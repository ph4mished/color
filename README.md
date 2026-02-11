
# Color

A comprehensive Go library for adding colors, styles, and formatting to terminal output with support for multiple color formats and truecolor detection. [This is a Go port of the Spectra color library](https://github.com/ph4mished/spectra)


# Installation

```bash
go get github.com/ph4mished/color
```

# Features

- Multiple Color Systems: Named colors, hex codes, RGB, 256-color palette
- TrueColor Detection: Automatic detection of terminal truecolor support
- Terminal Safe: Graceful fallbacks when color not supported(no-color fallback)
- Simple API: Easy-to-use functions for text styling and coloring.
- Comprehensive Styles: Bold, italic, underline, blink, reverse, hidden, strike-through
- Granular Resets: Individual and full reset codes for precise control
- No Escape: Texts in [] that aren't colors/styles are left as it is.

# Core Concepts

## Template System

The library follows a template-first approach: parse color templates once with or without placeholders([0], [1], etc), then reuse them with different data to replace placeholders.
**Placeholders are like slots**

## Color Toggling

Respects the NO_COLOR environment variable and detects when output is redirected. It can be manually controlled to suit user preference.

---

# Quick Start

```go
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
```

# Complete Usage Examples

## Basic Template with Placeholders

```go
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
```

## Basic Text Coloring
```go
package main

import "github.com/ph4mished/color"

func main(){
  // Simple colored text
  color.Parse("[fg=green]Success message![reset]").Apply()
  color.Parse("[fg=red bold]Error: Something went wrong![reset]").Apply()
  color.Parse("[fg=cyan italic]Info message[reset]").Apply()

  // Background colors
  color.Parse("[bg=blue fg=white]White text on blue background[reset]").Apply()
  color.Parse("[bg=lightgreen fg=black]Black text on light green[reset]").Apply()
}
```

## Advanced Color Formats

```go
package main

import "github.com/ph4mished/color"

func main(){
// Hex colors (requires truecolor support)
color.Parse("[fg=#FF5733]Orange hex color[reset]").Apply()
color.Parse("[bg=#3498db]Blue background[reset]").Apply()

// RGB colors
color.Parse("[fg=rgb(255,105,180)]Hot pink text[reset]").Apply()
color.Parse("[bg=rgb(50,205,50)]Lime green background[reset]").Apply()

// 256-color palette
color.Parse("[fg=214]Orange from 256-color palette[reset]").Apply()
color.Parse("[bg=196]Red background from palette[reset]").Apply()
}
```

## Text Styles

```go
package main

import "github.com/ph4mished/color"

func main(){
    // Combine styles
    color.Parse("[bold underline=single] Bold and underlined[reset]").Apply()
    color.Parse("[italic dim]", "Dim italic text. [italic=reset dim=reset][strike]Strikethrough  text only[reset]").Apply()
    color.Parse("[blink=slow hidden]Slow blinking hidden text[reset]").Apply()

    // Reset specific attributes
    color.Parse("[bold fg=blue]Blue bold text. [bold=reset]No longer bold, but still blue. [fg=reset]No color, but other styles remain[reset]").Apply()
}
```

## Color Toggling

```go
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
    forceColors := color.NewColorToggle(true)   // Always show colors
    noColors := color.NewColorToggle(false)     // Never show colors
    
    // Use in CLI applications
    useColor := os.Getenv("NO_COLOR") == ""
    appToggle := color.NewColorToggle(useColor)
    
    helpTemplate := appToggle.Parse("[bold fg=cyan][0][reset] [fg=green][1][reset]")
    fmt.Println(helpTemplate.Apply("Usage:", "myapp [options]"))
}
```

## Advanced Template Examples

```go
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
```

## Building Complex UIs

```go
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
```

## Project Structure Example

```go
// file: styles/styles.go - Define your color scheme
package styles

import "github.com/ph4mished/color"

var Toggle = color.NewColorToggle()

var Templates = struct {
    Success  color.CompiledTemplate
    Error    color.CompiledTemplate
    Warning  color.CompiledTemplate
    Info     color.CompiledTemplate
    Header   color.CompiledTemplate
    Flag     color.CompiledTemplate
}{
    Success:  Toggle.Parse("[fg=green bold]✓ [0][reset]"),
    Error:    Toggle.Parse("[fg=red bold]✗ [0][reset]"),
    Warning:  Toggle.Parse("[fg=yellow bold]⚠ [0][reset]"),
    Info:     Toggle.Parse("[fg=blue][0][reset]"),
    Header:   Toggle.Parse("[bold fg=cyan][0][reset]"),
    Flag:     Toggle.Parse("[fg=yellow][0][reset], [fg=yellow][1][reset]: [2]"),
}
```



```go
package cmd

//file: cmd/help.go - Use the color templates
import (
    "fmt"
    "yourproject/styles"
)

func ShowHelp() {
    fmt.Println(styles.Templates.Header.Apply("MyApp Help"))
    fmt.Println()
    
    fmt.Println(styles.Templates.Header.Apply("Usage:"))
    fmt.Println("  myapp [command] [options]")
    fmt.Println()
    
    fmt.Println(styles.Templates.Header.Apply("Commands:"))
    fmt.Println(styles.Templates.Flag.Apply("start", "", "Start the application"))
    fmt.Println(styles.Templates.Flag.Apply("stop", "", "Stop the application"))
    fmt.Println(styles.Templates.Flag.Apply("status", "", "Check application status"))
    fmt.Println()
    
    fmt.Println(styles.Templates.Header.Apply("Options:"))
    fmt.Println(styles.Templates.Flag.Apply("-h", "--help", "Show this help"))
    fmt.Println(styles.Templates.Flag.Apply("-v", "--version", "Show version"))
    fmt.Println(styles.Templates.Flag.Apply("-d", "--debug", "Enable debug mode"))
}
```


```go
package main

//file: main.go - Main application
import (
    "fmt"
    "os"
    "yourproject/styles"
    "yourproject/cmd"
)

func main() {
    if len(os.Args) > 1 && os.Args[1] == "--help" {
        cmd.ShowHelp()
        return
    }
    
    // Use color templates throughout
    fmt.Println(styles.Templates.Success.Apply("Application started"))
    
    // Process...
    
    fmt.Println(styles.Templates.Info.Apply("Processing completed"))
    fmt.Println(styles.Templates.Success.Apply("All tasks finished"))
}
```


## CLI Applications

```go
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
```

## Error Handling in Templates

```go
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
```

## Pattern to avoid
```go
// Good pattern
var appTemplates struct {
    Success color.Template
    Error   color.Template
}

func init() {
    toggle := color.NewColorToggle()
    appTemplates.Success = toggle.Parse("[fg=green] [0][reset]")
    appTemplates.Error = toggle.Parse("[fg=red] [0][reset]")
}

// Bad pattern (parsing in hot loop)
func processItems(items []string) {
    for _, item := range items {
        // DON'T DO THIS - parses every iteration!
        tmpl := color.Parse("[fg=blue]" + item + "[reset]")
        fmt.Println(tmpl.Apply())
    }
}
```

## Performance Comparison

```go
package main

import (
    "fmt"
    "time"
    "github.com/ph4mished/color"
)

func main() {
    const iterations = 1000000
    
    // Method 1: Parse once, apply many
    template := color.Parse("[bold fg=red][0][reset] [fg=green][1][reset]")
    
    start := time.Now()
    for i := 0; i < iterations; i++ {
        template.Apply(fmt.Sprintf("Item%d", i), fmt.Sprintf("Value%d", i))
    }
    fmt.Printf("Template reuse: %v\n", time.Since(start))
    
    // Method 2: Parse every time
    start = time.Now()
    for i := 0; i < iterations; i++ {
        color.Parse(fmt.Sprintf("[bold fg=red]Item%d[reset] [fg=green]Value%d[reset]", i, i)).Apply()
    }
    fmt.Printf("Parse every time: %v\n", time.Since(start))
    
    // Method 3: Manual concatenation
    start = time.Now()
    for i := 0; i < iterations; i++ {
        _ = color.ParseColor("fg=red bold") + fmt.Sprintf("Item%d", i) + 
            color.ParseColor("reset") + " " + 
            color.ParseColor("fg=green") + fmt.Sprintf("Value%d", i) + 
            color.ParseColor("reset")
    }
    fmt.Printf("Manual concatenation: %v\n", time.Since(start))
}
```

## Performance Comparison Result

```bash
Template reuse: 2.540160479s
Parse every time: 24.301094388s
Manual concatenation: 2.651576062s
```


# Spectra Syntax Reference

## Basic Colors
**Foreground Colors**
| Command | Effect |
|---------|--------|
| `fg=black` | Black text |
| `fg=red` | Red text |
| `fg=green` | Green text |
| `fg=yellow` | Yellow text |
| `fg=blue` | Blue text |
| `fg=magenta` | Magenta text |
| `fg=cyan` | Cyan text |
| `fg=white` | White text |
| `fg=darkgray` | Dark gray text |
| `fg=lightred` | Light red text |
| `fg=lightgreen` | Light green text |
| `fg=lightyellow` | Light yellow text |
| `fg=lightblue` | Light blue text |
| `fg=lightmagenta` | Light magenta text |
| `fg=lightcyan` | Light cyan text |
| `fg=lightwhite` | Light white text |


**Background Colors**
| Command | Effect |
|---------|--------|
| `bg=black` | Black background |
| `bg=red` | Red background |
| `bg=green` | Green background |
| `bg=yellow` | Yellow background |
| `bg=blue` | Blue background |
| `bg=magenta` | Magenta background |
| `bg=cyan` | Cyan background |
| `bg=white` | White background |
| `bg=darkgray` | Dark gray background |
| `bg=lightred` | Light red background |
| `bg=lightgreen` | Light green background |
| `bg=lightyellow` | Light yellow background |
| `bg=lightblue` | Light blue background |
| `bg=lightmagenta` | Light magenta background |
| `bg=lightcyan` | Light cyan background |
| `bg=lightwhite` | Light white background |


## Text Styles
| Command | Effect |
|---------|--------|
| `bold` | Bold/bright text |
| `dim` | Dim/faint text |
| `italic` | Italic text |
| `underline=single` | Single underlined text |
| `underline=double` | Double underlined text |
| `blink=slow` | Slow blinking text |
| `blink=fast` | Fast blinking text |
| `reverse` | Reverse video (swap foreground and background colors) |
| `hidden` | Hidden text |
| `strike` | Strikethrough text |

## Reset Commands
| Command | Effect |
|---------|--------|
| `reset` | Reset all colors and styles |
| `fg=reset` | Reset foreground color only |
| `bg=reset` | Reset background color only |
| `bold=reset` | Reset bold style only |
| `dim=reset` | Reset dim style only |
| `italic=reset` | Reset italic style only |
| `underline=reset` | Reset underline style only |
| `blink=reset` | Reset blink style only |
| `blinkfast=reset` | Reset fast blink style only |
| `reverse=reset` | Reset reverse style only |
| `hidden=reset` | Reset hidden style only |
| `strike=reset` | Reset strikethrough style only |


## Advanced Features
| Command | Effect |
|---------|--------|
| `fg=#RRGGBB` | Hex color for foreground |
| `bg=#RRGGBB` | Hex color for background |
| `fg=rgb(RR,GG,BB)` |RGB color for foreground |
| `bg=rgb(RR,GG,BB)` | RGB color for background |
| `fg=NNN` | 256-color palette (0-255) for foreground |
| `bg=NNN` | 256-color palette (0-255) for background |



# Tips and Best Practices

1. Parse Once: Always parse templates at initialization, not in loops
2. Use Toggles: Respect user preferences with color toggling
3. Template Reuse: Create templates for consistent styling
4. Placeholder Limits: The current implementation supports [0] through [999]
5. Testing: Test both color and no-color outputs


# Limitations

1. Terminal Dependency: Colors only work in terminals that support ANSI escape codes(Unix/Linux platforms)
2. TrueColor Requirement: Hex and RGB colors require terminal with truecolor support
3. Style Support: Some styles (blink, double underline) may not work in all terminals
4. Color Detection: Fallback from truecolor to 256-color not yet implemented
5. Windows: May require additional setup on Windows terminals

# Platform Support

- Linux/macOS terminals (full support)
- Windows Terminal/WSL (good support probably)
- Legacy Windows CMD (not supported)
- iTerm2, GNOME Terminal, Kitty (not tested yet)

# Contributing

We welcome contributions! Here's how you can help:

1. Report Bugs: Open an issue with reproduction steps
2. Suggest Features: Share your ideas for improvements
3. Submit PRs:
   - Fork the repository
   - Create a feature branch
   - Add tests for your changes
   - Ensure code follows Go conventions
   - Submit a pull request


# Development Setup

```bash
# Clone the repository
git clone https://github.com/ph4mished/color.git
cd color

# Run tests
go test 

```

# Areas Needing Improvement

1. Better Windows compatibility
2. 256-color fallback for truecolor
3. Performance optimization


# License

MIT License - see LICENSE file for details.

# Acknowledgments

- ANSI escape code specifications
- The Go community for testing and feedback
- All contributors who have helped improve this library

---

**Note**: Always test color output in different terminals to ensure compatibility with your users' environments. Consider providing a --no-color flag in your applications for users who prefer plain text.


