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
	result := readPdf()
	filterValue(result)
}

func readPdf() string {
	// Abre o arquivo PDF

	pdfFile, err := os.Open(os.Getenv("PDF_PATH"))
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return ""
	}

	// Cria um novo leitor de PDF
	reader, err := model.NewPdfReader(pdfFile)

	if err != nil {
		fmt.Println("Erro ao criar leitor de PDF:", err)
		return ""
	}

	// Obtém o número de páginas
	numPages, err := reader.GetNumPages()
	if err != nil {
		fmt.Println("Erro ao obter número de páginas:", err)
		return ""
	}

	// Lê cada página
	var alltext string
	for i := 2; i <= numPages; i++ {
		// Obtém a página atual
		page, err := reader.GetPage(i)
		if err != nil {
			fmt.Println("Erro ao obter página:", err)
			continue
		}

		// Extrai o texto da página
		text, err := page.GetContentStreams()
		if err != nil {
			fmt.Println("Erro ao extrair texto da página:", err)
			continue
		}

		// Exibe o texto da página
		fmt.Println("Página", i, ":")
		fmt.Println(text)

		alltext = strings.Join(text, ", ")

	}
	return alltext
}

func filterValue(result string) {
	re := regexp.MustCompile(`\([^)]+\)`)
	matches := re.FindAllString(result, -1)

	// Exibe itens extraídos
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
