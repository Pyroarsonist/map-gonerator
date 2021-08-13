package main

import (
	"fmt"
	"log"
	"os"
)

func printToFile(str string, fileName string) {
	f, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(str)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Written to", fileName)
	fmt.Println()
}

func printToConsole(str string) {
	fmt.Println()
	fmt.Println(str)
	fmt.Println()
}
