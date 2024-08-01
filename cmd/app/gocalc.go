package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	// extra
	"golang.org/x/term"
	"utils"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
const HISTORY_LEN = 20

func main() {

	// in case of panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic: ", r)
		}
	}()

	// preparing the ground
	fmt.Println(utils.Banner)
	fmt.Println("Enter an expression to evaluate, ? to help, q to quit:")
	var s string
	var history []string
	var historyPointer int
	var isCommand bool
	var file *os.File

	// looking for a history file
	var historyFile = ".gocalc_history"
	_, err := os.Stat(historyFile)
	if os.IsNotExist(err) {
		os.Create(".gocalc_history")
	} else if err != nil {
		fmt.Println(err)
		return
	}

	file, err = os.Open(historyFile)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			history = append(history, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
		historyPointer = len(history) - 1
	}

	for {
		s = ""
		fmt.Print("-> ")

		/*** Reading the expression ***/
		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}

		for {
			b := make([]byte, 1)
			_, err = os.Stdin.Read(b)
			if err != nil {
				fmt.Println(err)
				return
			}

			if b[0] == 13 { // ENTER
				term.Restore(int(os.Stdin.Fd()), oldState) // return back to normal
				break

			} else if b[0] == 8 || b[0] == 127 { // ITS BACKSPACE, DELETE LAST CHAR
				if len(s) > 0 { // Trying to avoid that the terminal gets broken
					s = s[:len(s)-1]
					fmt.Print("\b \b") // delete char
				}

			} else if b[0] == 27 { // ESCAPE SECUENCE
				_, err = reader.ReadByte()
				if err != nil {
					fmt.Println("Error: ", err)
				}
				last, err := reader.ReadByte()
				if err != nil {
					fmt.Println("Error: ", err)
				}
				switch last {
				case 65: // UP
					// history call
					for i := 0; i < len(s); i++ {
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
				case 66: // DOWN
					// do nothing
				case 67: // RIGHT
					// ?
				case 68: // LEFT
					// ?
				default:
				}

			} else {
				s += string(b)
				fmt.Print(string(b))
				if len(history) >= HISTORY_LEN {
          history = history[1:]
		      historyPointer = len(history) - 1
				}
			}

		}

		if s == "quit" || s == "q" || s == "exit" {
			fmt.Println()
			break
		}

		switch s {
		case "?":
			fmt.Println()
			for _, t := range utils.HelpMenu {
				fmt.Println(t)
			}
			isCommand = true

		case "clear_history":
			history = []string{}
			historyPointer = 0
			isCommand = true
			fmt.Println()

		case "show_history":
			fmt.Println()
			for i, t := range history {
				fmt.Printf("%d: %s\n", i+1, t)
			}
			isCommand = true

		default:
			isCommand = false
		}

		if !isCommand {
			/*** Parsing the expression ***/
			fmt.Println()
			rpn, err := utils.InfixToRPN(s)

			if err != nil {
				fmt.Println("Error: ", err.Error())
			} else {

				result := utils.EvaluateRPN(rpn)
				regex := regexp.MustCompile(`\.\d{10,}`)
				match := regex.FindString(fmt.Sprintf("%v", result))

				// printing the result
				if match != "" {
					fmt.Printf("= %.10f\n", result)
				} else {
					fmt.Printf("= %v\n", result)
				}
			}
		}
		// adding the correct expression to the history
		history = append(history, s)
		historyPointer = len(history) - 1
    fmt.Println(history)
	} // principal loop

	// Saving the history
  os.Remove(historyFile)
  os.Create(historyFile)
  file, err  = os.OpenFile(historyFile,os.O_WRONLY|os.O_TRUNC|os.O_CREATE,0644)

	for index, line := range history {
		_, err = file.WriteString(line+"\n")
		if err != nil {
			fmt.Printf("Error saving line %d\n", index+1)
			fmt.Println(err)
      }
    }
} // main
