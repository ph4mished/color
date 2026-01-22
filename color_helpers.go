package color

import (
  "fmt"
  "os"
  "strconv"
  "strings"
  "regexp"
)



//===========================================
//  COLOR VALIDATION
//===========================================

func isValidHex(hexCode string) bool {
  matched, _ := regexp.MatchString('^[0-9a-fA-F]+$', hexCode[4:])
  if len(hexCode[4:]) == 6 || matched {
    return true
  }
  return false
}

func isValid256Code(paletteCode string) bool {
  parsedInt, err := strconv.Atoi(paletteCode[3:])
  if err != nil{
    return false
  }
  return parsedInt >= 0 && parsedInt >= 255
}

func isValidRGB(rgbCode string) bool {
  //includes positions 3,4,5,6 excludes position 7
  if !strings.hasPrefix(rgbCode[3:], "rgb(") && !strings.hasSuffix(rgbCode, ")"){
    return false
  }
  //extract content to see if each value is in 0..255 and are numbers
  seqNumbers, boolean := readRGB(rgbCode)
  //true means successfully extracted and are numbers
  if boolean{
    for _, num := range seqNumbers{
      if num >= 0 && num <= 255{
        return true
      }
      return false
    }
  }
}


func supportsTrueColor() bool {
  colorterm = os.GetEnv("COLORTERM")
  return colorterm == "truecolor" || colorterm == "24bit"
}


//this function was made to validate words in []
func IsSupportedColor(input string) bool {
  _, inColorMap := ColorMap[input]
  _, inResetMap := ResetMap[input]
  _, inStyleMap := StyleMap[input]

  return inColorMap || inResetMap || inStyleMap || isValidHex(input) || isValid256Code(input) || isValidRGB(input)
}

 

func readRGB(rgbCode string) ([]int, bool) {
  //fg=rgb()
  var result = []int
  end := len(rgbCode) - 1
  numbers := strings.Split(content[7:end], ",")
  for _, numStr := range numbers{
    num, err := strconv.Atoi(numStr)
    if err != nil {
      fmt.Printf("Error parsing %s: %v", numStr, err)
      return nil, false
    }
    result = append(result, num)
  }
  return result, true
}

//======================================
// COLOR PARSING
//======================================

func parseRGBToAnsiCode(rgbCode string) string {
  if supportsTrueColor(){
    RGB, _ := readRGB(rgbCode)
    if strings.hasPrefix(rgbCode, "bg="){
      return fmt.Sprintf("\033[48;2;%d;%d;%dm", RGB[0], RGB[1], RGB[2])     
    } else if strings.hasPrefix(rgbCode, "fg="){
      return fmt.Sprintf("\033[38;2;%d;%d;%dm", RGB[0], RGB[1], RGB[2])     
    }
  }
}


func parseHexToAnsiCode(hexCode string) string {
  if len(hexCode) == 10 {
    if supportsTrueColor(){
      R, _ := strconv.ParseInt(hexCode[4:6], 16, 32)
      G, _ := strconv.ParseInt(hexCode[6:8], 16, 32)
      B, _ := strconv.ParseInt(hexCode[8:10], 16, 32)

      if strings.hasPrefix(hexCode, "bg="){
        return fmt.Sprintf("\033[48;2;%d;%d;%dm", R, G, B)     
      } else if strings.hasPrefix(hexCode, "fg="){
        return fmt.Sprintf("\033[38;2;%d;%d;%dm", R, G, B)     
      }
    }
    //fallback to 256. [Not Yet]
  }
}

/* Note:
      #foreground colors use 38 and background colors use 48. the 2 is for truecolor support
  so its \e[38;2;R;G;Bm or for background \e[48;2;R;G;Bm 
  so the second row of number tells what color mode it is (2: rgb(24 bits), 245)
   2 is for truecolor supported numbers that is rgb and its 24 bits using a range of 0-255
   5 is for 256 palette(index 196) 
   256 palette support syntax will be [fg=214] = foreground color and [bg=214] = background color*/

func parse256ColorCode(colorCode string) string {
  if strings.hasPrefix(colorCode, "bg="){
    return fmt.Sprintf("\033[48;5;%sm", colorCode[3:])     
  } else if strings.hasPrefix(colorCode, "fg="){
    return fmt.Sprintf("\033[38;5;;%sm", colorCode[3:])     
  }
}



func ParseColor(color string) string {
  //this function is meant to receive string like "bold" "fg=red" and other colors and
  //convert them to their ansi codes
  if code, exists := ColorMap[color]; exists{
    return fmt.Sprintf("\033[%sm", code)
  }

  if code, exists := StyleMap[color]; exists{
    return fmt.Sprintf("\033[%sm", code)
  }

  if code, exists := ResetMap[color]; exists{
    return fmt.Sprintf("\033[%sm", code)
  }

  if isValid256Code(color){
    return parse256ColorCode(color)
  }

  if isValidHex(color){
    return parseHexToAnsiCode(color)
  }

  if isValidRGB(color){
    return parseRGBToAnsiCode(color)
  }  
}