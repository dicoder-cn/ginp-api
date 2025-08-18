package desc

import (
	"ginp-api/internal/gen"
	"ginp-api/pkg/filehelper"
	"strings"
)

func AddImportRouterPackage(lineName string, parentDir ...string) {
	allSmallName := gen.NameToAllSmall(lineName)
	// packgeName := "r" + allSmallName
	content, err := filehelper.ReadContent(PathRouterEntry())
	if err != nil {
		println(err.Error())
		return
	}

	// 构建导入路径
	importPath := ""
	if len(parentDir) > 0 && parentDir[0] != "" {
		// Remove 'c' prefix from parent directory to fix import path
		importPath = parentDir[0] + "/c" + allSmallName
	} else {
		importPath = "c" + allSmallName
	}

	if strings.Contains(content, importPath) {
		println("outers_entry.go 已经存在", importPath, "不再添加")
		return
	}

	placeHolder := PlaceholderRouterImport
	importStr := ""
	if len(parentDir) > 0 && parentDir[0] != "" {
		// Fix the import path by ensuring we don't have 'c' prefix in parent directory
		// This ensures imports like 'admin/ctestentity8' instead of 'cadmin/ctestentity8'
		importStr = RouterReplaceStr + parentDir[0] + "/c" + allSmallName + `"`
	} else {
		importStr = RouterReplaceStr + "c" + allSmallName + `"`
	}

	newContent := strings.Replace(content, placeHolder, importStr+"\n\t"+placeHolder, -1)
	err = filehelper.WriteContent(PathRouterEntry(), newContent)
	if err != nil {
		println("outers_entry 写入失败" + err.Error())
		return
	} else {
		println("outers_entry import写入成功")
	}
}
