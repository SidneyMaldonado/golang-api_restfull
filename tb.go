package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	test4()
}
func test1() {

	var db, err = sql.Open("mysql", "u475983679_aula:Senha@01!@tcp(sql395.main-hosting.eu:3306)/u475983679_aula")
	if err != nil {
		log.Println(err.Error())
	}
	var rows, err2 = db.Query("select * from cidade")
	if err2 != nil {
		log.Println(err.Error())
	}

	types, _ := rows.ColumnTypes()
	for rows.Next() {
		row := make([]interface{}, len(types))
		for i := range types {
			row[i] = new(interface{})
		}
		rows.Scan(row...)

		for i, _ := range row {
			log.Println("Dado:", row[i])
		}

	}

}

func test2() {
	var db, err = sql.Open("mysql", "u475983679_aula:Senha@01!@tcp(sql395.main-hosting.eu:3306)/u475983679_aula")
	if err != nil {
		log.Println(err.Error())
	}
	var rows, err2 = db.Query("select * from cidade")
	if err2 != nil {
		log.Println(err.Error())
	}

	cols, _ := rows.Columns()
	vals := make([]interface{}, len(cols))
	result := make(map[string]interface{}, len(cols))

	for i, key := range cols {
		switch key {
		case "id", "status":
			vals[i] = new(int)
		default:
			vals[i] = new(string)
		}

		result[key] = vals[i]
	}

	b, _ := json.Marshal(result)
	fmt.Println(string(b))

}

func test3() {
	var db, err = sql.Open("mysql", "u475983679_aula:Senha@01!@tcp(sql395.main-hosting.eu:3306)/u475983679_aula")
	if err != nil {
		log.Println(err.Error())
	}
	var rows, err2 = db.Query("select * from cidade")
	if err2 != nil {
		log.Println(err.Error())
	}

	cols, _ := rows.Columns()

	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			log.Println(err.Error())
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = fmt.Sprintf("%v", *val)
		}
		fmt.Print(m)
	}
}
func format(s string, v interface{}) string {
	t, b := new(template.Template), new(strings.Builder)
	template.Must(t.Parse(s)).Execute(b, v)
	fmt.Println(b.String())
	return b.String()
}

type Grupo struct {
	linhas []Algo
}

type Algo struct {
	dados []string
}

func test4() {
	var db, err = sql.Open("mysql", "u475983679_aula:Senha@01!@tcp(sql395.main-hosting.eu:3306)/u475983679_aula")
	if err != nil {
		log.Println(err.Error())
	}
	var rows, err2 = db.Query("select * from cidade")
	if err2 != nil {
		log.Println(err.Error())
	}
	cols, _ := rows.Columns()

	data := make(map[string]string)

	var grupo Grupo

	for rows.Next() {

		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		rows.Scan(columnPointers...)

		for i, colName := range cols {
			data[colName] = columns[i]
		}

		var lista Algo
		for _, v := range data {
			//	fmt.Printf("key[%s] value[%s]\n", k, v)
			lista.dados = append(lista.dados, v)
		}

		grupo.linhas = append(grupo.linhas, lista)

	}

	fmt.Println(grupo.linhas)
	for i, a := range grupo.linhas {
		fmt.Printf("registro: [%d] => [%s] [%s] [%s]  \n", i, a.dados[0], a.dados[2], a.dados[1])
	}
}
