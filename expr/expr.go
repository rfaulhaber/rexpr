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

func NewNode() *Node {
	return &Node{}
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

	operandStack := &stack{}

	var nodeStack []*Node

	node := &Node{}

	for _, token := range tokens {
		if op, isOp := OperatorFromString(token); isOp {
			if node.IsLeaf() {
				rightStr := operandStack.Pop()
				rightVal, err := strconv.ParseFloat(rightStr, 2)

				if err != nil {
					return nil, fmt.Errorf("invalid operand: %s", rightStr)
				}

				leftStr := operandStack.Pop()
				leftVal, err := strconv.ParseFloat(leftStr, 2)

				if err != nil {
					return nil, fmt.Errorf("invalid operand: %s", leftStr)
				}

				newNode := &Node{
					left: &Node{ value: leftVal },
					right: &Node{ value: rightVal },
					operator: op,
				}

				node.right = newNode
			} else if node.left == nil {
				node.operator = op

				leftStr := operandStack.Pop()
				leftVal, err := strconv.ParseFloat(leftStr, 2)

				if err != nil {
					return nil, fmt.Errorf("invalid operand: %s", leftStr)
				}

				node.left = &Node{
					value: leftVal,
				}

				node = &Node{
					left: node,
				}
			} else {
				fmt.Println("trying to resolve issue", operandStack.Size())
				node = &Node{
					right: node,
					operator: op,
				}

				leftStr := operandStack.Pop()
				leftVal, err := strconv.ParseFloat(leftStr, 2)

				if err != nil {
					return nil, fmt.Errorf("invalid operand: %s", leftStr)
				}

				node.left = &Node{
					value: leftVal,
				}
			}
		} else if _, parseErr := strconv.ParseFloat(token, 2); parseErr == nil {
			operandStack.Push(token)
		} else {
			return nil, fmt.Errorf("token neither valid operator nor operand: %s", token)
		}
	}

	if node.right == nil {
		node = node.left
	} else if node.left == nil {
		node = node.right
	}

	return node, nil
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

func (n *Node) reparent(right *Node, operator Operator) *Node {
	return &Node{
		left:     n,
		right:    right,
		operator: operator,
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
