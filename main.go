package main

import (
	"fmt"
	"color"
)

func main(){
	template := color.Parse("[bold fg=yellow]Hello[reset]")
	fmt.Println(template.Apply())
}
