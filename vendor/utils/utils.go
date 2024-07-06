package utils

import (
	"errors"
	"strings"
	"unicode"
  //"fmt"
)

const LIMIT int = 31 // expression symbols limit


func Parse(s string) (interface{}, error) {

 s = strings.TrimSpace(s)
 s = strings.ReplaceAll(s," ","")

 if len(s) > LIMIT { return []string{}, errors.New("Too much symbols.") } 

  arr, l := symbols(s) 

  switch l {
   case -1 :
   return arr, errors.New("Expression not allowed.")
  }
 return calc(arr), nil 
}


// Input: A expression
// Output: An array of strings, containing numbers and operators
func symbols(s string) (arr []string, l int) {
 for i:=0; i<len(s); i++ {
  if unicode.IsDigit(rune(s[i])) {     // is a number
   number := string(s[i]) 
    for ; i+1 < len(s) && unicode.IsDigit(rune(s[i+1])) ; {
      i++ 
      number+=string(s[i]) 
    }
    arr = append(arr,number)
   if number[0] == '0' {
    arr[len(arr)-1] = "0"
   }

  } else if inOps(rune(s[i])) {        // is an operator
    arr = append(arr,string(s[i]))

  } else {                          
    return arr,-1                     // -1 , error for not allowed expressions
  }
 }

 l = len(arr)
 return 
}


// Input: A character
// Output: TRUE if the character is an operator, otherwise FALSE
func inOps(i rune) bool {
operators := []rune{'+','-','*','/','%','(',')'}
 for _, c := range operators {
  if c == i {
   return true
  }
 }
 return false
}

func calc(arr []string) float64 {
 return 3.1415
}


