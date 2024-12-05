package calc

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToFloat64(str string) float64 {
	degree := float64(1)
	var res float64 = 0
	var invers bool = false
	for i := len(str); i > 0; i-- {
		if str[i-1] == '-' {
			invers = true
		} else {
			res += float64(9-int('9'-str[i-1])) * degree
			degree *= 10
		}
	}
	if invers {
		res = 0 - res
	}
	return res
}

func IsSign(value rune) bool {
	return value == '+' || value == '-' || value == '*' || value == '/'
}

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")

	if len(expression) < 3 {
		return 0, fmt.Errorf("expression is too short")
	}

	if IsSign(rune(expression[0])) || IsSign(rune(expression[len(expression)-1])) {
		return 0, fmt.Errorf("expression cannot start or end with an operator")
	}

	for {
		openIdx := strings.LastIndex(expression, "(")
		if openIdx == -1 {
			break
		}

		closeIdx := strings.Index(expression[openIdx:], ")")
		if closeIdx == -1 {
			return 0, fmt.Errorf("missing closing parenthesis")
		}

		subExpression := expression[openIdx+1 : openIdx+closeIdx]
		result, err := Calc(subExpression)
		if err != nil {
			return 0, err
		}

		expression = expression[:openIdx] + strconv.FormatFloat(result, 'f', 0, 64) + expression[openIdx+closeIdx+1:]
	}

	var res float64
	var b string
	var currentOp rune
	var resFlag bool

	for _, value := range expression + "s" {
		switch {
		case value == ' ':
			continue
		case value >= '0' && value <= '9':
			b += string(value)
		case IsSign(value) || value == 's':
			if resFlag {
				switch currentOp {
				case '+':
					res += StringToFloat64(b)
				case '-':
					res -= StringToFloat64(b)
				case '*':
					res *= StringToFloat64(b)
				case '/':
					if b == "0" {
						return 0, fmt.Errorf("division by zero")
					}
					res /= StringToFloat64(b)
				default:
					return 0, fmt.Errorf("unknown operator: %c", currentOp)
				}
			} else {
				resFlag = true
				res = StringToFloat64(b)
			}
			b = ""
			currentOp = value
		default:
			return 0, fmt.Errorf("invalid character in expression: %c", value)
		}
	}
	return res, nil
}
