package mathoperation

import (
	"backend/models/stack"
	"backend/models/stackfloat"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Fungsi untuk mengecek apakah token merupakan angka
func IsNumber(token string) bool {
	pattern := regexp.MustCompile(`^[0-9]+$`)
	return pattern.MatchString(token)
}

// Fungsi untuk mengecek apakah notasi infix sudah valid
func IsInfixValid(infix string) bool {
	// Memisahkan tiap token
	tokens := strings.Split(infix, " ")

	// Menyimpan berapa banyak kurung yang masih terbuka
	openBracket := 0
	prevToken := "s" // Menyimpan token sebelumnya

	// Lakukan perulangan sebanyak token
	for _, token := range tokens {
		if openBracket < 0 {
			return false
		}

		switch token {
		case "(":
			openBracket++
			if IsNumber(prevToken) || prevToken == ")" || prevToken == "" {
				return false
			}
		case ")":
			openBracket--
			if prevToken == "(" || prevToken == "+" || prevToken == "-" || prevToken == "*" || prevToken == "/" || prevToken == "^" || prevToken == "" {
				return false
			}
		case "+", "-", "*", "/", "^":
			if prevToken == "" || prevToken == "(" || prevToken == "+" || prevToken == "-" || prevToken == "*" || prevToken == "/" || prevToken == "^" {
				return false
			}
		default:
			if !IsNumber(token) || IsNumber(prevToken) || prevToken == ")" || prevToken == "" {
				return false
			}
		}
		prevToken = token
	}

	// Jika masih ada kurung yang terbuka, maka infix tidak valid

	return openBracket == 0
}

func Precedence(operator string) int {
	switch operator {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^":
		return 3
	}
	return 0
}

func InfixToPostfix(infix string) string {
	output := []string{}
	stack := stack.Stack{}

	tokens := strings.Split(infix, " ")
	for _, token := range tokens {
		switch token {
		case "(":
			stack.Push(token)
		case ")":
			for !stack.IsEmpty() && stack.Peek() != "(" {
				output = append(output, stack.Pop())
			}
			if !stack.IsEmpty() && stack.Peek() == "(" {
				stack.Pop()
			}
		case "+", "-", "*", "/", "^":
			for !stack.IsEmpty() && Precedence(stack.Peek()) >= Precedence(token) {
				output = append(output, stack.Pop())
			}
			stack.Push(token)
		default:
			output = append(output, token)
		}
	}

	for !stack.IsEmpty() {
		output = append(output, stack.Pop())
	}

	return strings.Join(output, " ")
}

// Fungsi menghitung operasi matematika yang diberikan dalam bentuk postfix.
// Fungsi ini mengembalikan hasil perhitungan dan error jika terjadi kesalahan.
func CalculatePostfix(postfix string) (float64, error) {
	stack := stackfloat.StackFloat{}

	tokens := strings.Split(postfix, " ")
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/", "^":
			if stack.IsEmpty() {
				return 0, nil
			}
			operand2 := stack.Pop()
			if stack.IsEmpty() {
				return 0, nil
			}
			operand1 := stack.Pop()
			switch token {
			case "+":
				stack.Push(operand1 + operand2)
			case "-":
				stack.Push(operand1 - operand2)
			case "*":
				stack.Push(operand1 * operand2)
			case "/":
				stack.Push(operand1 / operand2)
			case "^":
				stack.Push(math.Pow(operand1, operand2))
			}
		default:
			operand, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			stack.Push(operand)
		}
	}

	if stack.IsEmpty() {
		return 0, nil
	}

	return stack.Pop(), nil
}
