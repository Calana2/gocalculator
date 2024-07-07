package utils

import (
	"errors"
	"strconv"
	"unicode"
  "strings"
  "fmt"
)

// Input: A expression
// Output: An array of strings, containing numbers and operators
func getSymbols(s string) (arr []string, l int) {

 s = strings.TrimSpace(s)
 s = strings.ReplaceAll(s," ","")

 for i:=0; i<len(s); i++ {
  char := rune(s[i])

  if unicode.IsDigit(char) { // is a number
    number := string(s[i])  
    for ; i+1 < len(s) && unicode.IsDigit(rune(s[i+1])) ; {
      i++ 
      number+=string(s[i])  // take all the chars of the number
    }
    arr = append(arr,number) 
    if number[0] == '0' {   // fixing a 0xxxx like number
     arr[len(arr)-1] = "0"
    } 
  }  else if isAritmeticOperator(string(s[i])) {                                                    // is an operator
    arr = append(arr,string(s[i]))
  } else {                          
    fmt.Println(string(s[i]))
    return arr,-1                                                             // -1 , error for not allowed expressions
  }
 }

 l = len(arr)
 return arr,l
}

// Input: A string with LEN=1
// Output; TRUE if is an operator, otherwise FALSE
func isAritmeticOperator(token string) bool {
 return token == "+" || token == "-" || token == "*" || token == "/" || token == "%"
}

func InfixToRPN(infix string) ([]string, error) {

 var rpn []string
 stack := []string{}
 symbols,l := getSymbols(infix)

 if l <= 0 {
  return []string{},errors.New("Not allowed expression")
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

func EvaluateRPN(rpn []string) int {
 stack := []int{}
 for _, token := range rpn {
  switch token {
  case "+", "-", "*", "/", "%":
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
    stack = append(stack, operand1/operand2)
   case "%":
    stack = append(stack, operand1%operand2)
   }
  default:
   num, _ := strconv.Atoi(token)
   stack = append(stack, num)
  }
 }
 return stack[0]
}


func precedence(operator string) int {
 switch operator {
 case "+", "-":
  return 1
 case "*", "/", "%":
  return 2
 }
 return 0   // for parenthesis
}

