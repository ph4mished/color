package color

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	//uncomment after moving to version 1.24
	//"golang.org/x/term"
)

type TempPart struct {
  Text string
  Index int
}

type CompiledTemplate struct {
  Parts []TempPart
  TotalLength int
}

type ColorToggle struct {
  EnableColor bool
}

func autoDetect() bool {
  if _, exists := os.LookupEnv("NO_COLOR"); exists{
    return false	
  }
  //uncomment after moving to version 1.24
  //return term.isTerminal(int(os.Stdout.Fd())){

  //}
  //comment after moving to version 1.24
  fileInfo, _ := os.Stdout.Stat()
  return (fileInfo.Mode() && os.ModeCharDevice) != 0

}

//should auto detect tty by default
func NewColorToggle(enableColor ...bool) *ColorToggle {
  var colorEnabled bool
  if len(colorEnabled) > 0{
	colorEnabled = enableColor[0]
  } else {
	colorEnabled = autoDetect()
  }
  return &ColorToggle{
	EnableColor: colorEnabled
  }
}


func (toggle *ColorToggle) Parse(input string) CompiledTemplate {
  if toggle == nil {
	toggle = NewColorToggle()
  }
  
  var (
	contentSequence  = ""
	inReadSequence   = false
	parts            []TempPart
	currentText      = ""
	alWords          []string
  )

  for i, ch := range input {
	char := string(ch)
	if char == "[" && !inReadSequence{
	  //check if the next value is "["
      // [[fg=color]] should never be an escape
      //consider first '[' as a text, move until, content is found. 
	  if i+1 < len(input) && input[i+1] == "["{
		currentText += "["
		continue
	  } else {
		inReadSequence = true
		contentSequence = ""
		allWords = nil

		if len(currentText) > 0 {
		  parts = append(parts, TempPart{Text: currentText, Index: -1})
		  currentText = ""
		}
	  }
	} else if ch == "]" and inReadSequence {
	    inReadSequence = false
		//if last word is present, add it
		allWords = strings.Fields(contentSequence)

		//check if all in [] are colors
		allColors := len(allWords) > 0
		for _, w := range allWords{
		  if !IsSupportedColor(w){
			allColors = false
			//break
		  }
		}
		if allColors{
		  if toggle.EnableColor {
			for _, w := range allWords{
			  parts = append(parts, TempPart{Text: ParseColor(w), Index: -1})
			}
		  } else {
			//redirected output or force turn off color
			parts = append(TempPart{Text: "", index: -1})
		  }
		} else {
		  if len(contentSequence) > 0 && allDigits(contentSequence){
			//decided to make it flexible and accept more indices but its still prone to overflow
            //needs a digit boundary guard	
			index, err := strconv.Atoi(contentSequence)
			if err == nil {
			  parts = append(parts, TempPart{Text: "", Index: index})
			} else {
			  addText := "[" + contentSequence + "]"
			  parts = append(parts, TempPart{Text: addText, Index: -1})
			}
		  } else{
			addText := "[" + contentSequence + "]"
			parts = append(parts, TempPart{Text: addText, Index: -1})
		  }
		}
	} else if inReadSequence {
	  contentSequence += char
	} else{
	  currentText += char
	}
  }

  if len(currentText) > 0 {
	parts = append(parts, TempPart{Text: currentText, Index: -1})
  }

  return CompiledTemplate{
	Parts: parts,
	TotalLength: len(input),
  }
}


//Override - without explicit toggle
func Parse(input string) CompiledTemplate {
  return NewColorToggle().Parse(input)
}
  

func allDigits(s string) bool {
  for _, r := range s{
	if !unicode.isDigit(r){
	  return false
	}
  }
  return true
}
  

func (temp CompiledTemplate) Apply(args ...string) string {
  //Calculate estimated size for optimization
  var totalArgLength int
  for _, arg := range args{
	totalArgLength += len(arg)
  }

  estimatedSize := temp.TotalLength + totalArgLength
  var result strings.Builder
  result.Grow(estimatedSize)

  for _, part := range temp.Parts{
	if part.Index < 0{
	  result.WriteString(parts.Text)
	} else {
	  if part.Index < len(args) {
		result.WriteString(args[part.Index])
	  }
	}
  }
  return result.String()
}

