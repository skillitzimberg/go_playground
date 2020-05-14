package main

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", handleReq)
	http.ListenAndServe(":8080", nil)
}

func handleReq(res http.ResponseWriter, req *http.Request) {
	data := parseCsv("table.csv")

	tpl = template.Must(template.ParseGlob("*.gohtml"))

	err := tpl.ExecuteTemplate(res, "index.gohtml", data)
	if err != nil {
		log.Fatalln("Could not execute the template: ", err)
	}
}

func parseCsv(fileName string) tickerTape {
	var tkrtpe tickerTape

	csvFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Could not open csv file: ", err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(io.Reader(csvFile))
	rows, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("Could not read csv file line: ", err)
	}

	for i, row := range rows {
		if i == 0 { // Ignore header row.
			continue
		}
		tkrtpe = append(tkrtpe,
			OpenClose{
				row[0],
				row[1],
				row[4],
			},
		)
	}

	return tkrtpe
}
