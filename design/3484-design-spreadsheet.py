from typing import List, Any, Optional, Dict
import unittest

# https://leetcode.com/problems/design-spreadsheet/

# python3 -m unittest design/3484-design-spreadsheet.py


class Spreadsheet:

    def __init__(self, rows: int):
        _ = rows
        self.grid: Dict[str, int] = {}

    def setCell(self, cell: str, value: int) -> None:
        self.grid[cell] = value

    def resetCell(self, cell: str) -> None:
        self.setCell(cell, 0)

    def getValue(self, formula: str) -> int:
        pair = formula[1:].split('+')
        x = self.grid.get(pair[0], 0) if pair[0][0].isalpha() else int(pair[0])
        y = self.grid.get(pair[1], 0) if pair[1][0].isalpha() else int(pair[1])
        return x + y


# Your Spreadsheet object will be instantiated and called as such:
# obj = Spreadsheet(rows)
# obj.setCell(cell,value)
# obj.resetCell(cell)
# param_3 = obj.getValue(formula)


class Solution(unittest.TestCase):
    def test(self):
        test_cases: List[tuple[List[str], List[Any], List[Any]]] = [
            (
                ["Spreadsheet", "getValue", "setCell", "getValue", "setCell", "getValue", "resetCell", "getValue"],
                [[3], ["=5+7"], ["A1", 10], ["=A1+6"], ["B2", 15], ["=A1+B2"], ["A1"], ["=A1+B2"]],
                [None, 12, None, 16, None, 25, None, 15],
            ),
        ]
        for commands, values, expected in test_cases:
            sheet: Optional[Spreadsheet] = None
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "Spreadsheet":
                        sheet = Spreadsheet(values[idx][0])
                    case "setCell":
                        if sheet is not None:
                            sheet.setCell(values[idx][0], values[idx][1])
                    case "resetCell":
                        if sheet is not None:
                            sheet.resetCell(values[idx][0])
                    case "getValue":
                        if sheet is not None:
                            output = sheet.getValue(values[idx][0])
                    case _:
                        pass
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")
