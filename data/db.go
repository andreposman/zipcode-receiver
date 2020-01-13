package data

import (
	"fmt"
	"os"
)

//SaveToFile receives the messages and saves to file
func SaveToFile(messages []string) {
	file, err := os.Create("messages.txt")
	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(messages)
}
