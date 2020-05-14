// Parse this CSV file, putting two fields from the contents of the CSV file into a data structure.

// Parse an html template, then pass the data from step 1 into the CSV template; have the html template display the CSV data in a web page.
package main

// OpenClose struct contains the date and opening and closing values.
type OpenClose struct {
	Date  string
	Open  string
	Close string
}

type tickerTape []OpenClose
