package models

import "aplicacao-web/database"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscarProdutos() []Produto {

	db := database.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CreateProdict(nome, descricao string, preco float64, quantidade int) {
	db := database.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}
func DeleteProduct(id string) {
	db := database.ConectaComBancoDeDados()

	delete, err := db.Prepare("delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}

func EditProduct(id string) Produto {
	db := database.ConectaComBancoDeDados()

	productDB, err := db.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	productUpdate := Produto{}

	for productDB.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = productDB.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		productUpdate.Id = id
		productUpdate.Nome = nome
		productUpdate.Descricao = descricao
		productUpdate.Preco = preco
		productUpdate.Quantidade = quantidade
	}

	defer db.Close()

	return productUpdate
}

func UpdateProduct(id int, nome, descricao string, preco float64, quantidade int) {
	db := database.ConectaComBancoDeDados()

	updateProduct, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
