package color

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T){
	tempAnsi := Parse("[bold fg=blue]Hello[reset]")
	temp256 := Parse("[bold fg=115]Hello[reset]")
	tempHex := Parse("[bold fg=#AABBCC]Hello[reset]")
	tempRGB := Parse("[bold fg=rgb(15,102,224)]Hello [reset]")
	fmt.Println("STATIC TEMPLATES")
	fmt.Println("    FOR ANSI: ", tempAnsi.Apply())
	fmt.Println("    FOR 256: ", temp256.Apply())
	fmt.Println("    FOR HEX: ", tempHex.Apply())
	fmt.Println("    FOR RGB: ", tempRGB.Apply())
}

func TestWithToggle(t *testing.T){
	toggle := NewColorToggle(false)

	tempAnsi := toggle.Parse("[bold fg=blue]Hello[reset]")
	temp256 := toggle.Parse("[bold fg=115]Hello[reset]")
	tempHex := toggle.Parse("[bold fg=#AABBCC]Hello[reset]")
	tempRGB := toggle.Parse("[bold fg=rgb(15,102,224)]Hello [reset]")
	fmt.Println("\n\nTEMPLATES WITH TOGGLE(COLOR OFF)")
	fmt.Println("    FOR ANSI: ", tempAnsi.Apply())
	fmt.Println("    FOR 256: ", temp256.Apply())
	fmt.Println("    FOR HEX: ", tempHex.Apply())
	fmt.Println("    FOR RGB: ", tempRGB.Apply())	

}

func TestForInterpolation(t *testing.T){
	tempAnsi := Parse("[bold fg=blue]Hello [fg=yellow][0][reset]")
	temp256 := Parse("[bold fg=115]Hello [fg=13][0][reset]")
	tempHex := Parse("[bold fg=#AABBCC]Hello [fg=#AAFFCC][0][reset]")
	tempRGB := Parse("[bold fg=rgb(15,102,224)]Hello [fg=rgb(10,94,104)][0][reset]")
	fmt.Println("\n\nTEMPLATES WITH PLACEHOLDER")
	fmt.Println("    FOR ANSI: ", tempAnsi.Apply("World"))
	fmt.Println("    FOR 256: ", temp256.Apply("World"))
	fmt.Println("    FOR HEX: ", tempHex.Apply("World"))
	fmt.Println("    FOR RGB: ", tempRGB.Apply("World"))	

}

//test for internal (unexportable functions) will also be made
