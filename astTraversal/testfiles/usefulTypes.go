package testfiles

import (
	"fmt"
	"strings"

	"github.com/Mrzrb/astra/astTraversal/testfiles/otherpkg1"
)

// MyStruct is a struct.
type MyStruct struct {
	// Name is a string
	Name string

	PP *MyStruct
}

// MyInt is an int.
type MyInt int

// SayHello is a method on MyStruct.
func (m *MyStruct) SayHello() {
	fmt.Println("Hello from", strings.Join([]string{"MyStruct", m.Name}, " "))
}

// ExternalPackage is a method on MyStruct that uses a type from another package.
func (m *MyStruct) ExternalPackage() {
	_ = otherpkg1.Foo{}
}
