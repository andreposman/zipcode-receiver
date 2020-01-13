package data

import (
	"fmt"
	"os"
	"strings"
)

//SaveToFile receives the messages and saves to file
func SaveToFile(messages []string) {
	file, err := os.Create("messages.txt")
	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(strings.Join(messages, "\n"))
}
