package main

import (
	"fmt"
	"github.com/ph4mished/color"
	//"color"
)

func main(){
	template := color.Parse("[bold fg=yellow]Hello[reset]")
	fmt.Println(template.Apply())
}
