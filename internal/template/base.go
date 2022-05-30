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
	"database/sql"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/clause"

	"github.com/zhaozhihom/gen"
	"github.com/zhaozhihom/gen/field"
	"github.com/zhaozhihom/gen/helper"

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