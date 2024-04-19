package common

import "os"

func WriteLogFile(result []byte, name string) {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	// Append text to the file without adding a new line
	if _, err := file.Write(result); err != nil {
		log.Fatalf("failed to append to file: %v", err)
	}
}
// func WriteLogFile() {
// 	file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatalf("failed to open file: %v", err)
// 	}
// 	defer file.Close()

// 	// Text to append
// 	text := "This is a new line."

// 	// Append text to the file without adding a new line
// 	if _, err := file.Write([]byte(text)); err != nil {
// 		log.Fatalf("failed to append to file: %v", err)
// 	}
// }
