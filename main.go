package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ls() {
	dirPath, _ := os.Getwd()

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	fmt.Printf("Contents of directory '%s':\n", dirPath)
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}

func writeAFile() {
	content := "Hello, Go!\nThis is a new line."
	err := os.WriteFile("output.txt", []byte(content), 0644) // 0644 are file permissions
	if err != nil {
		log.Fatal(err)
	}
	log.Println("String successfully written to output.txt")
}

func readTheFile() {
	// Read the file content into a byte slice
	contentBytes, err := os.ReadFile("output.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Convert the byte slice to a string
	contentString := string(contentBytes)

	fmt.Println("File content as string:")
	fmt.Println(contentString)
}

func insertStringToFile(filename string, offset int, content string) error {
	// Read the entire file content
	originalContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Calculate the new content size
	newContentSize := len(originalContent) + len(content)

	// Create a new byte slice to hold the modified content
	newContent := make([]byte, newContentSize)

	// Copy the part before the offset
	copy(newContent[:offset], originalContent[:offset])

	// Copy the string to be inserted
	copy(newContent[offset:offset+len(content)], []byte(content))

	// Copy the part after the offset
	copy(newContent[offset+len(content):], originalContent[offset:])

	// Write the new content back to the file, truncating the original
	err = ioutil.WriteFile(filename, newContent, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func main() {
	writeAFile()
	readTheFile()
	insertStringToFile("output.txt", 8, "->new<-")
	readTheFile()
}
