package modlr

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

func DbInit() error {

	os.Remove("modeldata.db")
	db, err := sqlx.Open("sqlite3", "modeldata.db")
	if err != nil {
		fmt.Printf("Error in creating DB: %s\n", err)

		return err
	}
	err = db.Ping()

	//var m model.Model
	//var models []model.Model

	schema := `CREATE TABLE if not exists model (
    name varchar,
    field varchar,
	colname varchar,
    coltype varchar);
	`
	_, err = db.Exec(schema)

	schema = `CREATE TABLE if not exists data (
    model_name varchar,
    str1 varchar,
	str2 varchar,
    str3 varchar,
    str4 varchar,
	str5 varchar 
	);
	`

	_, err = db.Exec(schema)

	if err != nil {
		fmt.Printf("Couldn't run schema:%s\n", err)
		return err
	}

	_, err = db.Exec("insert into model values ('employee', 'empcode','str1','varchar')")
	_, err = db.Exec("insert into model values ('employee', 'empname','str2','varchar')")
	_, err = db.Exec("insert into model values ('address', 'address','str1','varchar')")
	_, err = db.Exec("insert into model values ('address', 'city','str2','varchar')")
	_, err = db.Exec("insert into model values ('address', 'state','str5','varchar')")

	_, err = db.Exec("insert into data values ('employee', '101','raj','rajan',null,null)")
	_, err = db.Exec("insert into data values ('employee', '102','ravi',null,null,null)")

	if err != nil {
		fmt.Printf("Couldn't run schema:%s\n", err)
		return err
	}
	return nil
}
