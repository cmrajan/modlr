package modlr

import (
	"fmt"
	"testing"
)

func TestModelCreate(t *testing.T) {

	DbInit()
	dbmap := initDbMap()

	m := dbmap.AddModel("employee")

	fmt.Printf("Created model %#v\n", m)

}
