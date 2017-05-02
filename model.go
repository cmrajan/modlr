package modlr

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ModelMap struct {
	Models []*Model
	DB     *sqlx.DB
}

type Model struct {
	Name   string `db:"name" json:"name"`
	Fields []*Field
}

type Field struct {
	Name     string `db:"field" json:"field"`
	DBColumn string `db:"colname" json:"colname"`
	DBType   string `db:"coltype" json:"coltype"`
}

func NewModelMap() (*ModelMap, error) {

	db, err := sqlx.Open("sqlite3", "modeldata.db")
	if err != nil {
		fmt.Printf("Error in creating DB Conn: %s\n", err)
		return nil, err
	}
	return &ModelMap{DB: db}, nil

}

func InitModelMap() *ModelMap {
	mmap, err := NewModelMap()
	if err != nil {
		fmt.Printf("Error initiating Modelmap%s", err)
	}

	var modelNames []string
	type modelData struct {
		Name     string `db:"name"`
		Field    string `db:"field"`
		DBColumn string `db:"colname"`
		DBType   string `db:"coltype"`
	}

	md := []modelData{}

	err = mmap.DB.Select(&modelNames, "select distinct name from model")
	if err != nil {
		fmt.Printf("Couldn't get rows :%s\n", err)
		return nil

	}

	for _, name := range modelNames {

		md = []modelData{}
		err = mmap.DB.Select(&md, "SELECT name,field,colname,coltype FROM model where name =$1", name)
		if err != nil {
			fmt.Printf("Couldn't get model data for name %s :%s\n", name, err)
			return nil
		}
		colNums := len(md)
		modelData := &Model{Name: name}
		fld := &Field{}
		modelData.Fields = make([]*Field, 0, colNums)
		for j := 0; j < colNums; j++ {
			fld = &Field{Name: md[j].Field, DBColumn: md[j].DBColumn, DBType: md[j].DBType}

			modelData.Fields = append(modelData.Fields, fld)

		}

		mmap.Models = append(mmap.Models, modelData)

	}

	res, err := json.MarshalIndent(mmap, " ", " ")
	if err != nil {
		fmt.Println("could't print json")
	}
	fmt.Println("json data", string(res))
	return mmap
}
