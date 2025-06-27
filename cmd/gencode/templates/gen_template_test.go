package templates

import (
	"ginp-api/cmd/gencode/desc"
	"ginp-api/internal/gen"
	"testing"
)

// 快速将现有的文件做成模板文件
func TestReplace(t *testing.T) {

	demoEntityLineName := "demo_table"
	demoEntityBigName := "DemoTable"
	demoPackageName := "demotable"

	demoApiBigName := "DemoAddApi"
	demoApiLineName := "demo_add_api"

	replaceData := map[string]string{
		demoEntityLineName: desc.ReplaceLineName,
		demoEntityBigName:  desc.ReplaceEntityName,
		demoPackageName:    desc.ReplacePackageName,
		demoApiBigName:     desc.ReplaceApiNameBig,
		demoApiLineName:    desc.ReplaceApiNameLine,
	}

	//1.开始生成：entity文件
	tPathEntity := desc.TemplatePathEntity()
	oPathEntity := desc.PathEntity(demoEntityLineName)
	gen.ReplaceAndWriteTemplate(oPathEntity, tPathEntity, replaceData)

	//2.开始生成：router文件
	tPathRouter := desc.TemplatePathRouter()
	oPathRouter := desc.PathRouter(demoEntityLineName)
	gen.ReplaceAndWriteTemplate(oPathRouter, tPathRouter, replaceData)

	//3.开始生成：controller文件
	tPathController := desc.TemplatePathController()
	oPathController := desc.PathController(demoEntityLineName)
	gen.ReplaceAndWriteTemplate(oPathController, tPathController, replaceData)

	//4.开始生成：service文件
	tPathService := desc.TemplatePathService()
	oPathService := desc.PathService(demoEntityLineName)
	gen.ReplaceAndWriteTemplate(oPathService, tPathService, replaceData)

	//5.开始生成：model文件
	tPathRepository := desc.TemplatePathModel()
	oPathRepository := desc.PathModel(demoEntityLineName)
	gen.ReplaceAndWriteTemplate(oPathRepository, tPathRepository, replaceData)

	//6.开始生成：api文件
	// tPathApi := desc.TemplatePathAddApi()
	// oPathApi := desc.PathAddApi(demoEntityLineName, demoApiLineName)
	// println(oPathApi + "\n" + tPathApi)
	// gen.ReplaceAndWriteTemplate(oPathApi, tPathApi, replaceData)
}
