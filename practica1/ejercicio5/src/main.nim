import std/strutils
import std/tables

type
  Stack*[T] = ref object
    data*: array[1024, T]
    top*: int

proc push*[T](stack: var Stack[T], newval: T): void =
  stack.top += 1
  stack.data[stack.top] = newval

proc pop*[T](stack: var Stack[T]): T =
  var res = stack.data[stack.top]
  stack.top -= 1
  res

proc peek*[T](stack: var Stack[T]): T =
  if stack.top == -1:
    raise newException(ValueError, "stack is empty")
  stack.data[stack.top]

proc peek*[T](stack: var Stack[T], offset: int): T =
  if stack.top - offset < 0:
    raise newException(StackOverflowError, "stack underflow")
  stack.data[stack.top-offset]

proc lessPrecedence(a, b: string): bool =
  var precedence = {
    "^": 1,
    "*": 2,
    "/": 2,
    "+": 3,
    "-": 3
  }.toTable

  precedence[a] > precedence[b]


proc isOperator(val: string): bool =
  val in @["^", "/", "+", "-", "*"]

# proc cmpPrecedence(a, b: string): int =

proc infixToPostfix(infix: string): string =

  var operators: Stack[string] = Stack[string]()
  operators.top = -1

  var output: Stack[string] = Stack[string]()
  output.top = -1

  var datum = split(infix)
  for idx, value in datum:
    if value == "(":
      discard
    elif isOperator(value):
      try:
        var top = peek[string](operators)
        if lessPrecedence(value, top):
          var le = operators.top
          for idx in countdown(le, 0):
            if lessPrecedence(value, peek[string](operators)):
              var temp = pop[string](operators)
              push[string](output, temp)
            else:
              break
        else:
          push[string](operators, value)
      except ValueError:
        push[string](operators, value)
    else:
      push[string](output, value)
  echo operators.data, output.data
  join(output.data, " ")

proc evaluatePostfix(postfix: string): int =
  echo "todo"

proc evaluateExpression(expression: string): int =
  var postFix = infixToPostfix(expression)
  echo postFix
  return evaluatePostfix(postFix)

var
  s: Stack[string] = Stack[string]()

when isMainModule:
  s.top = -1

  echo evaluateExpression("45 + 98 / 34 * 34 - 0")
