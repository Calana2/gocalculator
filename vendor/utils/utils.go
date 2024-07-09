package utils

import (
	"errors"
	"strconv"
	"unicode"
   "strings"
   "fmt"
)


var HelpMenu = [...]string{
 "gocalc beta by Calana2.",
 "",
 "Available commands:",
 "? - Help",
 "clear_history - Delete the history",
 "show_history  - Show the history",
}


var Banner string = `
..####....####....####....####...##.......####..
.##......##..##..##..##..##..##..##......##..##.
.##.###..##..##..##......######..##......##.....
.##..##..##..##..##..##..##..##..##......##..##.
..####....####....####...##..##..######...####..
................................................`




/*** Functions to decompose the expression ***/

func getSymbols(s string) (arr []string, l int) {
// Input: A expression
// Output: An array of strings, containing numbers and operators

 s = strings.TrimSpace(s)
 s = strings.ReplaceAll(s," ","")

 for i:=0; i<len(s); i++ {
  char := rune(s[i])

  if unicode.IsDigit(char) { // is a number
    number := string(s[i])  
    for ; i+1 < len(s) && ( unicode.IsDigit(rune(s[i+1])) ||
                            string(s[i+1]) == "." ) ; {
      i++ 
      number+=string(s[i])  // take all the chars of the number
    }
    arr = append(arr,number) 
    if number[0] == '0' && len(number) > 1 && number[1] != '.' {   // fixing a 0xxxx like number
     arr[len(arr)-1] = "0"
    } 
  }  else if isAritmeticOperator(string(s[i])) {                                                    // is an operator
    arr = append(arr,string(s[i]))
  } else {                          
    return arr,-1                                                             // -1 , error for not allowed expressions
  }
 }

 l = len(arr)
 return arr,l
}

func isAritmeticOperator(token string) bool {
// Input: A string with LEN=1
// Output; TRUE if is an operator, otherwise FALSE
 return token == "+" || token == "-" || token == "*" || token == "/" || token == "%" || token == "(" || token == ")"
}




/*** Functions to calculate the expression ***/



func InfixToRPN(infix string) ([]string, error) {

 var rpn []string
 stack := []string{}
 symbols,l := getSymbols(infix)

 if l <= 0 {
  return []string{},errors.New("Not allowed expression")
 }

 // when the first number is negative an overflow may be caused
 if symbols[0] == "-" {
  num := symbols[1]
  symbols = symbols[1:len(symbols)]
  symbols[0] = "-" + num
 }

 for _, token := range symbols {
  switch token {
  case "+", "-", "*", "/", "%":
   for len(stack) > 0 && precedence(stack[len(stack)-1]) >= precedence(token) {
    rpn = append(rpn, stack[len(stack)-1])
    stack = stack[:len(stack)-1]
   }
   stack = append(stack, token)
  case "(":
   stack = append(stack, token)
  case ")":
   for len(stack) > 0 && stack[len(stack)-1] != "(" {
    rpn = append(rpn, stack[len(stack)-1])
    stack = stack[:len(stack)-1]
   }
   if len(stack) > 0 && stack[len(stack)-1] == "(" {
    stack = stack[:len(stack)-1]
   }
  default:
   rpn = append(rpn, token)
  }
 }
 for len(stack) > 0 {
  rpn = append(rpn, stack[len(stack)-1])
  stack = stack[:len(stack)-1]
 }
 return rpn,nil
}

func EvaluateRPN(rpn []string) float64 {
/* Iterate over the elements in reverse polish notation
   IF it is a number it stores it in the stack, 
   ELSE IF it is an operator, it takes the nth and the n-1st value out of the stack, 
   therefore performs n OPERATOR n-1 and stores it in the stack
   FINALLY it RETURNS the final value
*/
var result float64
defer func() {
 if r := recover(); r != nil {
  fmt.Println("Error: ",r)               
  result = 0
 } 
}()

 stack := []float64{}
 for _, token := range rpn {
  switch token {
  case "+", "-", "*", "/", "%":
   if len(stack) < 2 {
    panic("Bad expression")
    return 0
   }
   operand2 := stack[len(stack)-1]
   stack = stack[:len(stack)-1]
   operand1 := stack[len(stack)-1]
   stack = stack[:len(stack)-1]
   switch token {
   case "+":
    stack = append(stack, operand1+operand2)
   case "-":
    stack = append(stack, operand1-operand2)
   case "*":
    stack = append(stack, operand1*operand2)
   case "/":
    if operand2 == 0 {
     panic("Divition by zero")
    }
    stack = append(stack, operand1/operand2)
   case "%":
  //  stack = append(stack, operand1%operand2)
   }
  default:
   num, _ := strconv.ParseFloat(token,64)
   stack = append(stack,num)
  }
 }
 result = stack[0]
 return result
}

func precedence(operator string) int {
// Input: An operator
// Output: The level of precedence (0 is top)
 switch operator {
 case "%":
  return 1
 case "+", "-":
  return 2
 case "*", "/":
  return 3
 }
 return 0   // for parenthesis 
}

