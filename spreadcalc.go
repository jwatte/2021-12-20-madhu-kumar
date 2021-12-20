package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type valueType int
const (
	Empty valueType = iota
	String
	Number
	Formula
)

type Cell struct {
	typ valueType
	value interface{}
}

type Row struct {
	numInput int
	cells [26]Cell
}

type Spreadsheet struct {
	rows []Row
}

func NewSpreadSheet(size int) *Spreadsheet {
	return &Spreadsheet{rows : make([]Row, 0, size)}
}

func (sheet *Spreadsheet) readInput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Split(text, ",")

		var row Row
		row.numInput = len(fields)
		
		for i, field := range fields {
			tmpf := strings.TrimSpace(field)
			if (len(tmpf) == 0) {
				row.cells[i] = Cell{typ: Empty, value: nil} 
			} else if (tmpf[0] == '\'') {
				row.cells[i] = Cell{typ: String, value: tmpf[1:]}
			} else if tmpf[0] == '=' {
				row.cells[i] = Cell{typ: Formula, value: tmpf[1:]}
			} else {
				f, _ := strconv.ParseFloat(tmpf, 64)
				row.cells[i] = Cell{typ: Number, value: f}
			}
		}

		sheet.rows = append(sheet.rows, row)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Reading input")
		os.Exit(1)
	}
}

func (sheet *Spreadsheet) writeOutput() {
	var b strings.Builder
	for _, row := range sheet.rows {
		for i := 0; i < row.numInput; i++ {
			cell := row.cells[i]
			switch cell.typ {
			case Empty:
				fmt.Fprintf(&b, "%s,", "")
			case String:
				fmt.Fprintf(&b, "%s,", "'" + cell.value.(string))
			case Number:
				fmt.Fprintf(&b, "%.02f,", cell.value.(float64))
			case Formula:
				fmt.Fprintf(&b, "not implemented", cell.value.(string))
			}
		}
		fmt.Println(strings.TrimSuffix(b.String(), ","))
		b.Reset()
	}
}


func (sheet *Spreadsheet) evaluateFormula(formula string) Cell {


	// Formula evaluation is not done.
	// Idea is to use a stack to evaluate formula

	return Cell{}
}

func (sheet *Spreadsheet) evaluateRow(r Row) Row {

	var newRow Row
	newRow.numInput = r.numInput
	
	for i := 0; i < r.numInput; i++ {
		cell := r.cells[i]
		switch cell.typ  {
		case Empty:
			newRow.cells[i] = Cell{typ : Empty, value: nil}
		case String:
			newRow.cells[i] = Cell{typ: String, value: cell.value.(string)}
		case Number:
			newRow.cells[i] = Cell{typ: Number, value:cell.value.(float64)}
		case Formula:
			newRow.cells[i] = sheet.evaluateFormula(cell.value.(string))
		}
	}

	return newRow
}

func (sheet *Spreadsheet) evaluate() *Spreadsheet {

	newSheet:= NewSpreadSheet(len(sheet.rows))
	for _, row := range sheet.rows {
		row = sheet.evaluateRow(row)
		newSheet.rows = append(newSheet.rows, row)
	}

	return newSheet
}


func main() {
	sheet := NewSpreadSheet(100)
	sheet.readInput(os.Stdin)

	// To be done. Evaluate formulas
	newSheet := sheet.evaluate()
	newSheet.writeOutput()
	// fmt.Fprintf(os.Stderr, "ERROR: spreadcalc go is not implemented\n")
	// os.Exit(1)
}
