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

// in case of panic
defer func() {
 if r := recover(); r != nil {
  fmt.Println("Panic: ",r)
 }
}()

// preparing the ground
fmt.Println("Enter an expression to evaluate, q to quit:")
var s string
var history []string
var historyPointer int
var isCommand bool
  
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

  } else if b[0] == 27 { // ESCAPE SECUENCE
    _, err = reader.ReadByte()
    if err != nil { fmt.Println("Error: ",err) }
    last, err := reader.ReadByte()
    if err != nil { fmt.Println("Error: ",err) }
    switch last {
     case 65:              // UP
      // history call
     for i:=0; i < len(s); i++ {
      fmt.Print("\b \b") // delete char
     }
     s = ""

     // Searching
     if len(history) > 0 {
      fmt.Print(history[historyPointer])
      s = history[historyPointer]
      if historyPointer == 0 {
       historyPointer = len(history) - 1
      } else {
       historyPointer--
      }
     } 
     case 66:              // DOWN
      // do nothing
     case 67:              // RIGHT
      // ?
     case 68:              // LEFT
      // ?
     default:
    }

  } else {
   s+=string(b)
   fmt.Print(string(b))
   if len(history) > 15 {
    history = history[:len(history)-1]
   }
  }
  
}

  if s == "quit" || s == "q" || s == "exit" { break }

  switch s {
   case "?":
    fmt.Println()
    for _, t := range utils.HelpMenu {
     fmt.Println(t)
    }
    isCommand=true

   case "clear_history":
    history=[]string{}
    historyPointer=0
    isCommand=true
    fmt.Println()

   case "show_history":
    fmt.Println()
    for i, t := range history {
     fmt.Printf("%d: %s\n",i+1,t)
    }
    isCommand=true

   default:
    isCommand=false
  }

  if !isCommand {
/*** Parsing the expression ***/
   fmt.Println()
   rpn,err := utils.InfixToRPN(s)
 
   if err != nil {
    fmt.Println("Error: ",err.Error())
   } else { 

  result := utils.EvaluateRPN(rpn)

    // printing the result
    fmt.Println(result) 
   }
  }
    // adding the correct expression to the history
    history = append(history, s)
    historyPointer = len(history) - 1
 } // principal loop
} // main


