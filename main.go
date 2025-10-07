package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
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

func insertStringsToFile(filename string, insertions map[int]string) error {
	// Read the entire file content
	originalContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Create a new byte slice to hold the modified content
	newContent := make([]byte, len(originalContent))

	// Copy original content to new content
	copy(newContent, originalContent)

	// Keep track of total length of inserted strings
	totalInsertedLength := 0

	// Sort the insertion map by keys (offsets) to ensure proper order
	offsets := make([]int, 0, len(insertions))
	for offset := range insertions {
		offsets = append(offsets, offset)
	}
	sort.Ints(offsets)

	for _, offset := range offsets {
		content := insertions[offset]
		insertionOffset := offset + totalInsertedLength

		// Create a new byte slice to hold modified content after each insertion
		newContentWithInsertions := make([]byte, len(newContent)+len(content))
		copy(newContentWithInsertions[:insertionOffset], newContent[:insertionOffset])
		copy(newContentWithInsertions[insertionOffset:], []byte(content))
		copy(newContentWithInsertions[insertionOffset+len(content):], newContent[insertionOffset:])

		// Update newContent to the newly constructed content
		newContent = newContentWithInsertions
		totalInsertedLength += len(content)
	}

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

	insertions := map[int]string{
		8:  "->new<-",
		20: "->another<-",
	}

	err := insertStringsToFile("output.txt", insertions)
	if err != nil {
		log.Fatalf("Error inserting strings: %v", err)
	}

	readTheFile()
}
