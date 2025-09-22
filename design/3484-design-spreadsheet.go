package design

import (
	"strings"
)

// https://leetcode.com/problems/design-spreadsheet/

type Spreadsheet struct {
	grid map[string]int
}

func NewSpreadsheet(rows int) Spreadsheet {
	_ = rows

	return Spreadsheet{grid: make(map[string]int)}
}

func (s *Spreadsheet) SetCell(cell string, value int) {
	s.grid[cell] = value
}

func (s *Spreadsheet) ResetCell(cell string) {
	s.SetCell(cell, 0)
}

func (s *Spreadsheet) GetValue(formula string) int {
	getValue := func(address string) int {
		if address[0] >= 'A' && address[0] <= 'Z' {
			return s.grid[address]
		}

		var value int
		for _, c := range address {
			value = value*10 + int(c-'0')
		}

		return value
	}
	pair := strings.Split(formula[1:], "+")
	x, y := getValue(pair[0]), getValue(pair[1])

	return x + y
}

// /**
//  * Your Spreadsheet object will be instantiated and called as such:
//  * obj := Constructor(rows);
//  * obj.SetCell(cell,value);
//  * obj.ResetCell(cell);
//  * param_3 := obj.GetValue(formula);
//  */
