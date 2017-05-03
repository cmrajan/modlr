package modlr

import (
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestModelCreate(t *testing.T) {

	_ = DbInit()

	dbmap := initDbMap()

	_ = dbmap.AddModel("employee")
	_ = dbmap.AddModel("address")

	fmt.Printf("Created model %#v\n", dbmap.models)

}
