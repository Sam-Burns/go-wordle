package words

import (
	"fmt"
	"sort"
)

type FrequencyTable struct {
	letterTotals    map[rune]int
	cumulativeTotal int
}

type PrintableFrequencyTableRow struct {
	Character  rune
	usageRatio float64
}

func MakeFrequencyTableFromWordlist(wl Wordlist) *FrequencyTable {

	letterTotals := make(map[rune]int)

	cumulativeTotal := 0

	for _, word := range wl.Words {
		for _, character := range word.Characters {
			letterTotals[character] += 1
			cumulativeTotal += 1
		}
	}

	return &FrequencyTable{letterTotals, cumulativeTotal}
}

func (frequencyTable FrequencyTable) getFrequencyRatioForCharacter(character rune) float64 {
	return float64(frequencyTable.letterTotals[character]) / float64(frequencyTable.cumulativeTotal)
}

func (frequencyTable FrequencyTable) GetPrintableFrequencyTableRows() []PrintableFrequencyTableRow {
	unsortedRows := getUnsortedPrintableFrequencyTableRows(frequencyTable)
	sortedRows := sortRows(unsortedRows)
	return sortedRows
}

func getUnsortedPrintableFrequencyTableRows(frequencyTable FrequencyTable) map[rune]PrintableFrequencyTableRow {

	rows := make(map[rune]PrintableFrequencyTableRow)

	for character := 'A'; character <= 'Z'; character++ {
		usageRatio := frequencyTable.getFrequencyRatioForCharacter(character)
		rows[character] = PrintableFrequencyTableRow{character, usageRatio}
	}

	return rows
}

func (printableFrequencyTableRow PrintableFrequencyTableRow) GetPrintablePercentage() string {
	percentage := fmt.Sprintf("%.2f", printableFrequencyTableRow.usageRatio*100)

	if len(percentage) < 5 {
		percentage = " " + percentage
	}

	percentage += "%"

	return percentage
}

func sortRows(rows map[rune]PrintableFrequencyTableRow) []PrintableFrequencyTableRow {

	rowSlice := []PrintableFrequencyTableRow{}

	for _, row := range rows {
		rowSlice = append(rowSlice, row)
	}

	sort.Slice(rowSlice, func(i, j int) bool {
		return rowSlice[i].usageRatio > rowSlice[j].usageRatio
	})

	return rowSlice
}
