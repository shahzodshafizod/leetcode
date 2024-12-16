import unittest

# https://leetcode.com/problems/parsing-a-boolean-expression/

# python3 -m unittest stacks/1106-parsing-a-boolean-expression.py

class Solution(unittest.TestCase):
    # # Approach: Recursive
    # # Time: O(N)
    # # Space: O(N)
    # def parseBoolExpr(self, expression: str) -> bool:
    #     s = expression
    #     idx = 0
    #     def parse() -> bool:
    #         nonlocal idx
    #         c = s[idx]
    #         idx += 1
    #         match c:
    #             case 't': return True
    #             case 'f': return False
    #             case '!':
    #                 idx += 1 # skipping the openning parenthesis
    #                 res = not parse()
    #                 idx += 1 # skipping the closing parenthesis
    #                 return res
    #         children = []
    #         idx += 1 # skip '('
    #         while s[idx] != ')':
    #             if s[idx] == ',':
    #                 idx += 1
    #             else:
    #                 children.append(parse())
    #         idx += 1 # skip ')'
    #         if c == '&':
    #             return all(children)
    #         return any(children)
    #     return parse()

    # Approach: Iterative (Using Stack)
    # Time: O(N)
    # Space: O(N)
    def parseBoolExpr(self, expression: str) -> bool:
        stack = []
        for c in expression:
            if c == ')':
                value = ''
                has_true, has_false = False, False
                while stack[-1] != '(':
                    value = stack.pop()
                    if value == 't': has_true = True
                    else: has_false = True
                stack.pop() # remove '('
                oper = stack.pop()
                if any([
                    oper == '!' and value == 'f',
                    oper == '&' and not has_false,
                    oper == '|' and has_true,
                ]):
                    stack.append('t')
                else:
                    stack.append('f')
            elif c != ',':
                stack.append(c)
        return stack[-1] == 't'

    def test(self):
        for expression, expected in [
            ("&(|(f))", False),
            ("|(f,f,f,t)", True),
            ("!(&(f,t))", True),
            ("!(&(!(t),&(f),|(f)))", True),
            ("!(&(!(&(f)),&(t),|(f,f,t)))", False),
            ("|(&(t,f,f,t,t,t,t,f,f,t),t,f)", True),
            ("|(&(t,f,t,!(&(|(f,f,f,f)))),t)", True),
            ("!(&(&(!(&(f)),&(t),|(f,f,t)),&(t),&(t,t,f)))", True),
        ]:
            output = self.parseBoolExpr(expression)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
