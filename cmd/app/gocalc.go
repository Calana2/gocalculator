package main 

import(
  "os"
  "bufio"
  "fmt"
  "utils"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

var history []string
const HISTORY_LEN int = 15

func main() {

 fmt.Println("Enter an expression to evaluate, q to quit:")
  
 for {
  fmt.Print("-> ")

  s,err := reader.ReadString('\n')

  if err != nil {
   fmt.Println("Error reading last line: ", err)
  }

  history = append(history,s)   

  if len(history) > HISTORY_LEN {
   history = history[1:]
  }

  if s == "quit\n" || s == "q\n" {
   break
  } 

  res, err := utils.Parse(s)

  if err != nil {
   fmt.Println("Error: ",err.Error())
  } else { fmt.Println(res) }
 }
}
