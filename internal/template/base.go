package template

const NotEditMark = `
// Code generated by github.com/zhaozhihom/gen. DO NOT EDIT.
// Code generated by github.com/zhaozhihom/gen. DO NOT EDIT.
// Code generated by github.com/zhaozhihom/gen. DO NOT EDIT.
`

const Header = NotEditMark + `
package {{.Package}}

import(
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/clause"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"

	"gorm.io/plugin/dbresolver"

	{{if .StructPkgPath}}"{{.StructPkgPath}}"{{end}}
	{{range .ImportPkgPaths}}{{.}}` + "\n" + `{{end}}
)
`

const UnitTestHeader = NotEditMark + `
package {{.Package}}

import(
	"context"
	"fmt"
	"strconv"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	
	{{if .StructPkgPath}}"{{.StructPkgPath}}"{{end}}
	{{range .ImportPkgPaths}}{{.}}` + "\n" + `{{end}}
)

`
