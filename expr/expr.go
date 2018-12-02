package expr


import (
	"fmt"
	"github.com/pkg/errors"
	"math"
	"strconv"
	"strings"
)

type Node struct {
	left     *Node
	right    *Node
	value    float64
	operator Operator
}

func NewNode(value float64) *Node {
	return &Node{
		value: value,
	}
}

func (n *Node) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *Node) Evaluate() (float64, error) {
	if n.IsLeaf() {
		return n.value, nil
	} else {
		left, err := n.left.Evaluate()

		if err != nil {
			return 0.0, errors.Wrap(err, "left evaluation failed")
		}

		right, err := n.right.Evaluate()

		if err != nil {
			return 0.0, errors.Wrap(err, "right evaluation failed")
		}

		switch n.operator {
		case Add:
			return left + right, nil
		case Sub:
			return left - right, nil
		case Mult:
			return left * right, nil
		case Div:
			return left / right, nil
		case IDiv:
			return math.Floor(left / right), nil
		case Pow:
			return math.Pow(left, right), nil
		case Fact:
			return float64(factorial(int64(left))), nil
		default:
			return 0.0, errors.New("invalid operator found!")
		}
	}
}

func ParseString(s string) (*Node, error) {
	tokens := strings.Split(s, " ")

	nodes := nodeStack{}

	for _, token := range tokens {
		if op, isOp := OperatorFromString(token); isOp { // if operator
			right := nodes.Pop()
			left := nodes.Pop()

			nodes.Push(&Node{
				left: left,
				right: right,
				operator: op,
			})
		} else if number, isNumber := tokenIsNumber(token); isNumber { // if number
			nodes.Push(NewNode(number))
		} else { // we don't know what it is!
			return nil, fmt.Errorf("token neither valid operator nor operand: %s", token)
		}
	}

	return nodes.Pop(), nil
}

func (n *Node) String() string {
	if n.IsLeaf() {
		if n.value == float64(int64(n.value)) {
			return fmt.Sprintf("%d", int64(n.value))
		} else {
			return fmt.Sprintf("%f", n.value)
		}
	} else {
		return fmt.Sprintf("(%s %s %s)", n.left.String(), n.operator.String(), n.right.String())
	}
}

func factorial(i int64) int64 {
	if i <= 2 {
		return i
	} else {
		return i * factorial(i-1)
	}
}

type Operator int

const (
	Add = iota
	Sub
	Mult
	Div
	IDiv
	Fact
	Pow
)

var operators = [...]string{
	"+",
	"-",
	"*",
	"/",
	"//",
	"!",
	"^",
}

var strToOperator = map[string]Operator{
	"+":  Add,
	"-":  Sub,
	"*":  Mult,
	"/":  Div,
	"//": IDiv,
	"!":  Fact,
	"^":  Pow,
}

func (o Operator) String() string {
	return operators[o]
}

func OperatorFromString(s string) (Operator, bool) {
	op, ok := strToOperator[s]
	return op, ok
}

func tokenIsNumber(token string) (float64, bool) {
	float, parseErr := strconv.ParseFloat(token, 2)

	return float, parseErr == nil
}
