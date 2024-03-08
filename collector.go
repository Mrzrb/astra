package astra

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type MethodInfo struct {
	MethodName     string
	FileName       string
	Line           int
	MethodReceiver string
	PkgName        string
}

func Collector(rootPath string) map[string]MethodInfo {
	directory := rootPath
	methodsMap := make(map[string]MethodInfo)

	err := parseDirectoryRecursively(directory, func(path string) {
		fileMethods := ParseAndStoreMethods(path)
		for k, v := range fileMethods {
			methodsMap[k] = v
		}
	})
	if err != nil {
		log.Println(err)
	}

	return methodsMap
}

func parseDirectoryRecursively(directory string, processFile func(string)) error {
	return filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			processFile(path)
		}
		return nil
	})
}

func ParseAndStoreMethods(filename string) map[string]MethodInfo {
	result := make(map[string]MethodInfo)

	// 解析源码
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	// 遍历AST
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		// case *ast.TypeSpec:
		// 	// 找到struct
		// 	if st, ok := x.Type.(*ast.StructType); ok {
		// 		// 遍历struct的方法
		// 		for _, f := range st.Fields.List {
		// 			if f.Names != nil {
		// 				for _, fieldName := range f.Names {
		// 					fmt.Println("Found struct field:", fieldName.Name)
		// 				}
		// 			}
		// 		}
		// 	}
		case *ast.FuncDecl:
			// 找到方法
			if x.Recv != nil {
				methodName := x.Name.Name
				pos := fset.Position(x.Pos())
				methodReceiver := ""
				if len(x.Recv.List) > 0 {
					switch nn := x.Recv.List[0].Type.(type) {
					case *ast.StarExpr:
						switch nnn := nn.X.(type) {
						case *ast.Ident:
							recvIdent := nnn
							methodReceiver = "*" + recvIdent.Name
						}
					case *ast.IndexExpr:
						recvIdent := nn.X.(*ast.Ident)
						methodReceiver = recvIdent.Name
					case *ast.Ident:
						recvIdent := nn
						methodReceiver = recvIdent.Name
					}
				}
				pkgName := node.Name.Name
				methodInfo := MethodInfo{
					MethodName:     methodName,
					FileName:       pos.Filename,
					Line:           pos.Line,
					MethodReceiver: methodReceiver,
					PkgName:        pkgName,
				}
				result[methodName] = methodInfo
			}
		}
		return true
	})

	return result
}

func ExtraMethodAndName(testString string) (string, string, error) {
	re := regexp.MustCompile(`\((.*?)\)\.(.*)-fm`)
	matches := re.FindStringSubmatch(testString)
	if len(matches) == 3 {
		receiver := matches[1]
		functionName := matches[2]
		return receiver, functionName, nil
	} else {
		return "", "", errors.New("fail")
	}
}
