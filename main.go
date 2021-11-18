package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type Payload struct {
	Stuff Data
}
type Data struct {
	Fruit Fruits
	Veggies Vegetables
}
type City struct {
	Id         int
	Name       string
	Population int
}

type Cidade struct {
	codCidade int
	nome string
	codestado int
}

type Fruits map[string]int
type Vegetables map[string]int


func main(){

//	test()
	var db, _ = conectarBanco()

	query("select * from cidade", db)


	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request ) {
		response := []byte("")
		response, _ = getJsonResponse()

		_,_ = fmt.Fprint(w,  string(response) )
		return

	})
	_ = http.ListenAndServe(":3000", r)

}


func getJsonResponse()([]byte, error) {
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

func conectarBanco() (*sql.DB, error){

	var banco, err = sql.Open("mysql","u475983679_aula:Senha@01!@tcp(sql395.main-hosting.eu:3306)/u475983679_aula")
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
	} else {
		log.Printf("Banco de dados inicializado com sucesso")
	}
	return banco, err
}

func query( comando string, db *sql.DB) (result *sql.Rows,error error){

	var res, err = db.Query(comando)

	defer res.Close()

	if err != nil{
		log.Printf("Erro ao recuperar dados")
		log.Printf(err.Error())
	} else {
		log.Printf("Dados recuperado")
		for res.Next(){

			var cidade Cidade
			err := res.Scan(&cidade.codCidade, &cidade.nome, &cidade.codestado)
			if err != nil {
				log.Printf(err.Error())
			}
			fmt.Printf("%v\n", cidade)
		}
	}
	if res == nil {
		log.Printf("sem dados para processar")
	}

	return res, err

}

func test(){
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