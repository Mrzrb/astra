package astTraversal

import (
	"go/ast"
	"go/doc"
)

type DeclarationTraverser struct {
	Traverser *BaseTraverser
	Decl      ast.Node
	File      *FileNode
	VarName   string
}

func (t *BaseTraverser) Declaration(node ast.Node, varName string) (*DeclarationTraverser, error) {
	return &DeclarationTraverser{
		Traverser: t,
		Decl:      node,
		File:      t.ActiveFile(),
		VarName:   varName, // The name of the variable on the LHS of the arrangement
	}, nil
}

func (d *DeclarationTraverser) Value() (ast.Node, error) {
	switch n := d.Decl.(type) {
	case *ast.ValueSpec:
		var index int
		for i, name := range n.Names {
			if d.VarName == d.Traverser.ExtractVarName(name).Type {
				index = i
				break
			}
		}

		return n.Values[index], nil
	case *ast.AssignStmt:
		var index int
		for i, name := range n.Lhs {
			if d.VarName == d.Traverser.ExtractVarName(name).Type {
				index = i
				break
			}
		}

		return n.Rhs[index], nil
	default:
		return nil, ErrInvalidNodeType
	}
}

func (d *DeclarationTraverser) GoDoc() (*doc.Value, error) {
	pkgDoc, err := d.Traverser.Packages.GoDoc(d.File.Package)
	if err != nil {
		return nil, err
	}

	switch n := d.Decl.(type) {
	case *ast.ValueSpec:
		var index int
		for i, name := range n.Names {
			if d.VarName == d.Traverser.ExtractVarName(name).Type {
				index = i
			}
		}

		for _, constDoc := range pkgDoc.Consts {
			for _, name := range constDoc.Names {
				if name == n.Names[index].Name {
					return constDoc, nil
				}
			}
		}

		for _, varDoc := range pkgDoc.Vars {
			for _, name := range varDoc.Names {
				if name == n.Names[index].Name {
					return varDoc, nil
				}
			}
		}
	case *ast.AssignStmt:
		// GoDoc for assignment statements is not supported
		return nil, nil
	}

	return nil, nil
}
