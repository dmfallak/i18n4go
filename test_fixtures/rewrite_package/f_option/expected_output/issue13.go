package input_files

import (
	"fmt"
	"os"
	"path/filepath"
)

func Issue13() string {
	someString := T("hello")
	fmt.Println(someString, T("world"))
	fmt.Println(someString, T("hello"), T("world"))

	fmt.Println(someString, T("Hello world {{.Arg0}}", map[string]interface{}{"Arg0": someString}))

	return fmt.Sprint(10, T("world"))
}
