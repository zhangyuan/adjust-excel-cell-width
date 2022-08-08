package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"unicode/utf8"

	"github.com/xuri/excelize/v2"
)

func main() {
	if err := invoke(); err != nil {
		log.Fatalln(err)
	}
}

func invoke() error {
	if len(os.Args) != 2 {
		return errors.New("invalid argument")
	}

	filePath := os.Args[1]

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	for _, sheetName := range f.GetSheetList() {
		cols, err := f.GetCols(sheetName)
		if err != nil {
			return err
		}

		for idx, col := range cols {
			maxWidth := 0
			for _, rowCell := range col {
				cellWidth := utf8.RuneCountInString(rowCell) + 2
				if cellWidth > maxWidth {
					maxWidth = cellWidth
				}
			}

			name, err := excelize.ColumnNumberToName(idx + 1)
			if err != nil {
				return err
			}

			f.SetColWidth(sheetName, name, name, float64(maxWidth))
		}
	}

	f.Save()
	return nil
}
