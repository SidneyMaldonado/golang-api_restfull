package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

type Payload struct {
	Stuff Data
}
type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}
type City struct {
	Id         int
	Name       string
	Population int
}

type Cidade struct {
	Codcidade int
	Nome      string
	Codestado int
}

type Coluna struct {
	ordem int
	nome  string
	tipo  string
}
type Registro struct {
	dados []string
}

type Tabela struct {
	estrutura []Coluna
	registros []Registro
}

type Fruits map[string]int
type Vegetables map[string]int

func main2() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) { processGet(w, r) })
	log.Printf("Server started...")
	_ = http.ListenAndServe(":3000", r)

}

func processGet(w http.ResponseWriter, r *http.Request) {

	log.Printf("Iniciou get..." + r.URL.Path)
	var response = []byte("")
	var err error
	var routeparts = strings.Split(r.URL.Path, "/")
	log.Print("Length:", len(routeparts))
	table := routeparts[1]
	log.Printf("Table: " + table)

	if table == "favicon.ico" {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if len(routeparts) != 2 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		response, err = buscarTabela(table)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	_, _ = fmt.Fprint(w, string(response))
	return
}
func buscarTabela(table string) ([]byte, error) {

	var cmd = "select * from " + table
	log.Printf(cmd)
	var db, _ = conectarBanco()
	var data []byte
	var err error
	log.Printf("tabela:", table)
	data, err = query2(cmd, db)
	db.Close()
	return data, err
}
func getJsonResponse() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Oranges"] = 10

	vegetables := make(map[string]int)
	vegetables["Cars"] = 10
	vegetables["Beets"] = 0

	d := Data{fruits, vegetables}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}

func conectarBanco() (*sql.DB, error) {

	//	var banco, err = sql.Open("mysql","u475983679_aula:Senha@01!@tcp(sql395.main-hosting.eu:3306)/u475983679_aula")
	var banco, err = sql.Open("mysql", "root:@tcp(localhost:3306)/sistema")
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
	} else {
		log.Printf("Banco de dados inicializado com sucesso")
	}
	return banco, err
}

func query(comando string, db *sql.DB) (result []byte, error error) {

	var res, err = db.Query(comando)

	var b []byte
	defer res.Close()
	var cidades []Cidade

	if err != nil {
		log.Printf("Erro ao recuperar dados")
		log.Printf(err.Error())
	} else {
		log.Printf("Dados recuperado")

		var i = 0
		for res.Next() {

			var nova Cidade
			err := res.Scan(&nova.Codcidade, &nova.Nome, &nova.Codestado)
			cidades = append(cidades, nova)
			i++
			if err != nil {
				fmt.Printf("Erro %s", err)
			}
		}
	}

	b, err = json.Marshal(cidades)
	if err != nil {
		fmt.Printf("Erro %s", err)
	}
	//fmt.Printf( string(b))

	return b, err

}
func query2(comando string, db *sql.DB) (result []byte, error error) {

	var res, err = db.Query(comando)

	var b []byte
	defer res.Close()
	var tabela Tabela

	var colunas, _ = res.Columns()
	var tipos, _ = res.ColumnTypes()

	for i, s := range colunas {
		var col Coluna
		col.nome = s
		col.tipo = tipos[i].Name()
		tabela.estrutura = append(tabela.estrutura, col)
	}

	log.Print("Estrutura Length:", len(tabela.estrutura))
	if err != nil {
		log.Printf("Erro ao recuperar dados")
		log.Printf(err.Error())
	} else {
		log.Printf("Dados recuperado")

		for res.Next() {
			var novo Registro
			cols, _ := res.Columns()
			dados := make([]interface{}, len(cols))

			if err = res.Scan(dados...); err != nil {
				log.Print("erro scan:", err)
			}

			log.Print("novo dados:", dados)
			novo.dados = dados
			tabela.registros = append(tabela.registros, novo)
		}
	}

	log.Print("Conteudo:", tabela.registros[0].dados)

	log.Print("Registros Length:", len(tabela.registros))

	b, err = json.Marshal(tabela.registros)
	if err != nil {
		fmt.Printf("Erro %s", err)
	}
	//fmt.Printf( string(b))

	return b, err

}

func test() {
	db, err := sql.Open("mysql", "u475983679_aula:Senha@01!@tcp(sql395.main-hosting.eu:3306)/u475983679_aula")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Query("SELECT * FROM cidade")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {

		var city City
		err := res.Scan(&city.Id, &city.Name, &city.Population)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", city)
	}
}
