package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
	"github.com/unidoc/unipdf/v3/model"
)

func main() {
	loadEnv()
	readPdf()

}

func loadEnv() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load file .env:", err)
	}
}

func readPdf() {
	// Open PDF file.

	pdfFile, err := os.Open(os.Getenv("PDF_PATH"))
	if err != nil {
		fmt.Println("Failed to open file:", err)

	}

	// Create a new reader of page.
	reader, err := model.NewPdfReader(pdfFile)

	if err != nil {
		fmt.Println("Failed to create a pdf file:", err)

	}

	// Get number of pages
	numPages, err := reader.GetNumPages()
	if err != nil {
		fmt.Println("Failed to get number of pages:", err)

	}

	for i := 2; i <= numPages; i++ {
		// Get main Page
		page, err := reader.GetPage(i)
		if err != nil {
			fmt.Println("Failed to get main page", err)
			continue
		}

		//Get text from page.
		lines, err := page.GetContentStreams()
		if err != nil {
			fmt.Println("Failed to extract text from page:", err)
			continue
		}

		result := filter(lines)
		fmt.Println(result)

	}
}

func filter(lines []string) []string {
	regexData := regexp.MustCompile(`^\d{2}/\d{2}$`) // Format DD/MM
	arrayCount := 0

	catchedLines := []string{}
	for _, line := range lines {
		var buffer bytes.Buffer

		re := regexp.MustCompile(`\([^)]+\)`)
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			extrated := match[1 : len(match)-1] // Remove parentheses due there are parentheses when read a pdf file.
			fmt.Fprintln(&buffer, extrated)

		}
		lineWithFilter := buffer.String()
		lineIntoArray := strings.Split(lineWithFilter, "\n")

		for _, line := range lineIntoArray {
			array := lineIntoArray
			arrayCount++

			if haveCard(line) {
				fmt.Println(line)
			}

			if regexData.MatchString(line) {
				position := arrayCount

				store := array[position]
				value := array[position+1]
				fmt.Println(store)
				fmt.Println(value)
				fmt.Println(line)

			}
		}

	}

	return catchedLines

}

func haveCard(line string) bool {
	cards := []string{os.Getenv("FIRST_USER_CARD"), os.Getenv("FIRST_USER_VIRTUAL_CARD"),
		os.Getenv("SECOND_USER_CARD"), os.Getenv("SECOND_USER_VIRTUAL_CARD")}

	for _, card := range cards {
		if strings.Contains(line, card) {
			return true
		}
	}
	return false
}

//TODO: Find "Detalhamento da Fatura na leitura" and use this as a header.
//Fields
// Compra
// Data
// Descri��o
// Parcela
// R$
// US$
