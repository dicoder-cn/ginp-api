package main

import (
	"ginp-api/cmd/gencode/desc"
	"ginp-api/internal/gen"
)

func main() {
	desc.GetPwd()
	result := gen.Input(desc.InputGenType, nil)

	switch result {
	case "1":
		desc.GenEntity()
	case "2":
		desc.GenFields()
	case "3":
		desc.GenAddApi()
	default:
		println("输入错误")
	}

}
