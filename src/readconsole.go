package src

import (
	"bufio"
	"fmt"
	"os"
)

func ReadConsole() string {
	var text string
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter file location (ex: C:\\Users\\User\\Documents\\MT202.txt): ")
		scanner.Scan()

		text = scanner.Text()

		// text = `C:\Users\ASUS\Downloads\newgens\0_Assessment-Preparation\202MEP_52A_57A_58A.txt`

		valid, err := isValid(text)
		if err != nil {
			printError(err.Error())
			continue
		}
		if !valid {
			printError("Not valid path: " + text)
			continue
		}

		break
	}

	return text
}

func printError(err string) {
	fmt.Printf("Error: %s\n\n", err)
}
func isValid(fp string) (bool, error) {
	// Check if file already exists
	if _, err := os.Stat(fp); err != nil {
		return false, err
	}

	return true, nil
}
