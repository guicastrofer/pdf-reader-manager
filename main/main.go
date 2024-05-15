package main

import (
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

func readPdf() {
	// Open PDF file.

	pdfFile, err := os.Open(os.Getenv("PDF_PATH"))
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)

	}

	// Create a new reader of page.
	reader, err := model.NewPdfReader(pdfFile)

	if err != nil {
		fmt.Println("Erro ao criar leitor de PDF:", err)

	}

	// Get number of pages
	numPages, err := reader.GetNumPages()
	if err != nil {
		fmt.Println("Erro ao obter número de páginas:", err)

	}

	// Read each page

	for i := 2; i <= numPages; i++ {
		// Get main Page
		page, err := reader.GetPage(i)
		if err != nil {
			fmt.Println("Erro ao obter página:", err)
			continue
		}

		//Get text from page.
		text, err := page.GetContentStreams()
		if err != nil {
			fmt.Println("Erro ao extrair texto da página:", err)
			continue
		}

		// Show text
		fmt.Println("Página", i, ":")

		filterValue(strings.Join(text, ", "))

	}

}

func filterValue(result string) {
	re := regexp.MustCompile(`\([^)]+\)`)
	matches := re.FindAllString(result, -1)

	// Show extracted infos
	for _, match := range matches {
		fmt.Println(match[1 : len(match)-1]) // Remove parentheses
	}
}

func loadEnv() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env:", err)
	}
}

//TODO: Find "Detalhamento da Fatura na leitura" and use this as a header.
//Fields
// Compra
// Data
// Descri��o
// Parcela
// R$
// US$
