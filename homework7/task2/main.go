package main

import (
	"fmt"
	"log"

	"go/ast"
	"go/parser"
	"go/token"
)

//Написать функцию, которая принимает на вход имя файла и название функции.
//Необходимо подсчитать в этой функции количество вызовов асинхронных функций.
//Результат работы должен возвращать количество вызовов int и ошибку error.
//Разрешается использовать только go/parser, go/ast и go/token.
func main() {
	count, err := CountAsyncFunctions("someGoFileWithCode.go", "TheFunctionWithSomeAsyncFunctionsInside")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Async functions found:", count)
}

func CountAsyncFunctions(fileName, funcName string) (count int, err error) {
	functionExists := false
	set := token.NewFileSet()
	astFile, err := parser.ParseFile(set, fileName, nil, 0)
	if err != nil {
		return
	}

	for _, d := range astFile.Decls {
		if fn, isFunc := d.(*ast.FuncDecl); isFunc {
			if fn.Name.String() != funcName {
				continue
			}
			functionExists = true
			for _, stmt := range fn.Body.List {
				switch stmt.(type) {
				case *ast.GoStmt:
					count++
				}
			}
			break
		}
	}
	if !functionExists {
		err = fmt.Errorf("function with name \"%s\" was not found", funcName)
	}
	return
}
