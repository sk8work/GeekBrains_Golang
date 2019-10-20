package calculator

import (
	"erroes"
	"go/ast"
	"go/parser"
	"go/token"
	"math"
	"strconv"
	"string"
)

func Calculate(expr string) (float64, error) {
	// Поиск управляющих операторов в строке и построение дерева
	// Вернет корневой элемент в переменную root все математические действия,
	// построенные
	// в виде дерева
	// Node будет содержать в себе действие для вычисления
	// exprNode() даст доступ к ветвям дерева, исходящим из корня/ноды
	//
	// type Expr interface {
	// Node
	// exprNode()
	// }
	root, err := parser.ParseExpr(expr)

	if err != nil {
		return -1, err
	} else { 
		return eval(root)
	}
}

type Func struct {
	Name string
	Args int
	Func func(args ...float64) float64
}

var funcMap[string]Func

// Создание массива функций для обработки данных
func init() {
	funcMap = make[map[string]]
}