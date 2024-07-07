package main 

import(
  "os"
  "bufio"
  "fmt"
  "utils"
  "golang.org/x/term"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
func main() {

fmt.Println("Enter an expression to evaluate, q to quit:")
var s string
  
 for {
  s=""
  fmt.Print("-> ")

/*** Reading the expression ***/ 
  oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
  if err != nil { panic(err) } 

for {
  b := make([]byte, 1)
  _, err = os.Stdin.Read(b)
  if err != nil {
   fmt.Println(err)
   return
  }


  if b[0] == 13 {  // ENTER
   term.Restore(int(os.Stdin.Fd()),oldState) // return back to normal
   break
  } else if b[0] == 8 { // ITS BACKSPACE, DELETE LAST CHAR
    if len(s) > 0 {     // Trying to avoid that the terminal gets broken
     s=s[:len(s)-1] 
     fmt.Print("\b \b") // delete char
    }
  } else {
   s+=string(b)
   fmt.Print(string(b))
  }
}

  if s == "quit" || s == "q" {
   break
  } 

/*** Parsing the expression ***/
//  fmt.Println()
//  fmt.Println(s)
  fmt.Println()
  rpn,err := utils.InfixToRPN(s)

  if err != nil {
   fmt.Println("Error: ",err.Error())
  } else { 
   result := utils.EvaluateRPN(rpn)
//   fmt.Println("Expresión:", s)
//   fmt.Println("RPN:", rpn)
   fmt.Println(result)
  }
 }
}


