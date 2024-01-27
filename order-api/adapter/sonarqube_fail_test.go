package sonarqube_fail_test

import (
	"fmt"
	"os"
)

func writeToFile(content, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	_, err = file.WriteString(content)
	// Recurso não é fechado adequadamente em caso de erro
	return err
}

func main() {
	content := "Hello, SonarQube!"
	filePath := "example.txt"

	err := writeToFile(content, filePath)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
