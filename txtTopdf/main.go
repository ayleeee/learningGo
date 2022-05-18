package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

func main() {

	var file string
	fmt.Print("업로드 파일 : ")
	fmt.Scanln(&file)

	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("%s file not found", file)
	}

	fileName := strings.TrimSuffix(file, path.Ext(file))

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 14)
	pdf.MultiCell(190, 5, string(content), "0", "0", false)

	_ = pdf.OutputFileAndClose(fileName + ".pdf")
	fmt.Println("PDF Created")

}
