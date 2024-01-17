package database

import (
	"API/fundamentos/internal/entity"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TestCreateNewProduct é uma função de teste que verifica se a função Create do repositório de produto funciona corretamente.
func TestCreateNewProduct(t *testing.T) {
	// dsn é uma string que contém as informações de conexão com o banco de dados MySQL.
	dsn := "root:root@tcp(localhost:3305)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	// db é uma variável que armazena a instância do GORM, que é um ORM (Object-Relational Mapping) para Go.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// Se houver um erro ao abrir a conexão com o banco de dados, o teste é encerrado com uma mensagem de erro.
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// db.AutoMigrate é uma função do GORM que cria ou atualiza automaticamente as tabelas do banco de dados de acordo com as structs definidas no pacote entity.
	db.AutoMigrate(&entity.Product{})

	// product é uma variável que armazena uma instância da struct Product, que representa um produto no sistema.
	// A função NewProduct é uma função do pacote entity que cria um novo produto com um nome e um preço.
	product, err := entity.NewProduct("Leite de vaca", 10.00)

	// assert é uma biblioteca que fornece funções para verificar se as condições esperadas são atendidas nos testes.
	// assert.Nil verifica se o valor passado como argumento é nil, ou seja, se não há erro ao criar o produto.
	assert.Nil(t, err)

	// productDb é uma variável que armazena uma instância do repositório de produto, que é uma camada de abstração que permite interagir com o banco de dados usando o GORM.
	productDb := NewProduct(db)

	// A função Create é uma função do repositório de produto que insere um novo produto no banco de dados.
	err = productDb.Create(product)

	// assert.Nil verifica se o valor passado como argumento é nil, ou seja, se não há erro ao inserir o produto no banco de dados.
	assert.Nil(t, err)
	// assert.NoError é equivalente a assert.Nil, mas com uma mensagem de erro mais clara.
	assert.NoError(t, err)
	// assert.NotEmpty verifica se o valor passado como argumento não está vazio, ou seja, se o produto tem um ID gerado pelo banco de dados.
	assert.NotEmpty(t, product.ID)
}

// TestFindProductById é uma função de teste que verifica se a função FindById do repositório de produto funciona corretamente.
func TestFindProductById(t *testing.T) {
	// dsn é uma string que contém as informações de conexão com o banco de dados MySQL.
	dsn := "root:root@tcp(localhost:3305)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	// db é uma variável que armazena a instância do GORM, que é um ORM (Object-Relational Mapping) para Go.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// Se houver um erro ao abrir a conexão com o banco de dados, o teste é encerrado com uma mensagem de erro.
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// db.AutoMigrate é uma função do GORM que cria ou atualiza automaticamente as tabelas do banco de dados de acordo com as structs definidas no pacote entity.
	db.AutoMigrate(&entity.Product{})

	// product é uma variável que armazena uma instância da struct Product, que representa um produto no sistema.
	// A função NewProduct é uma função do pacote entity que cria um novo produto com um nome e um preço.
	product, err := entity.NewProduct("Leite de vaca", 10.00)

	// assert é uma biblioteca que fornece funções para verificar se as condições esperadas são atendidas nos testes.
	// assert.Nil verifica se o valor passado como argumento é nil, ou seja, se não há erro ao criar o produto.
	assert.Nil(t, err)

	// productDb é uma variável que armazena uma instância do repositório de produto, que é uma camada de abstração que permite interagir com o banco de dados usando o GORM.
	productDb := NewProduct(db)

	// A função Create é uma função do repositório de produto que insere um novo produto no banco de dados.
	err = productDb.Create(product)

	// assert.Nil verifica se o valor passado como argumento é nil, ou seja, se não há erro ao inserir o produto no banco de dados.
	assert.Nil(t, err)

	// productFound é uma variável que armazena uma instância da struct Product, que representa o produto encontrado no banco de dados.
	// A função FindById é uma função do repositório de produto que busca um produto no banco de dados pelo seu ID.
	productFound, err := productDb.FindById(product.ID.String())

	// assert.Nil verifica se o valor passado como argumento é nil, ou seja, se não há erro ao buscar o produto no banco de dados.
	assert.Nil(t, err)

	// assert.Equal verifica se os valores passados como argumentos são iguais, ou seja, se os atributos do produto criado e do produto encontrado são os mesmos.
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)

	// assert.NotEmpty verifica se o valor passado como argumento não está vazio, ou seja, se os atributos do produto encontrado não são nulos ou vazios.
	assert.NotEmpty(t, productFound.ID)
	assert.NotEmpty(t, productFound.Name)
	assert.NotEmpty(t, productFound.Price)
}

// func TestFindByEmail(t *testing.T) {
// 	dsn := "root:root@tcp(localhost:3305)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("Failed to connect to database: %v", err)
// 	}

// 	db.AutoMigrate(&entity.Product{})
// }

// TestFindAllProducts é uma função de teste que verifica se a função FindAll do repositório de produto funciona corretamente.
func TestFindAllProducts(t *testing.T) {
	// dsn é uma string que contém as informações de conexão com o banco de dados MySQL.
	dsn := "root:root@tcp(localhost:3305)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	// db é uma variável que armazena a instância do GORM, que é um ORM (Object-Relational Mapping) para Go.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// Se houver um erro ao abrir a conexão com o banco de dados, o teste é encerrado com uma mensagem de erro.
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// db.AutoMigrate é uma função do GORM que cria ou atualiza automaticamente as tabelas do banco de dados de acordo com as structs definidas no pacote entity.
	db.AutoMigrate(&entity.Product{})

	// Este loop for cria 23 produtos com nomes e preços aleatórios e os insere no banco de dados.
	for i := 1; i < 24; i++ {
		// product é uma variável que armazena uma instância da struct Product, que representa um produto no sistema.
		// A função NewProduct é uma função do pacote entity que cria um novo produto com um nome e um preço.
		// A função fmt.Sprintf formata uma string com um valor substituído, neste caso o valor de i.
		// A função rand.Float64 gera um número aleatório entre 0 e 1, que é multiplicado por 100 para obter um preço aleatório.
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		// assert é uma biblioteca que fornece funções para verificar se as condições esperadas são atendidas nos testes.
		// assert.NoError verifica se o valor passado como argumento é nil, ou seja, se não há erro ao criar o produto.
		assert.NoError(t, err)
		// db.Create é uma função do GORM que insere um registro no banco de dados.
		db.Create(product)
	}

	// productDB é uma variável que armazena uma instância do repositório de produto, que é uma camada de abstração que permite interagir com o banco de dados usando o GORM.
	productDB := NewProduct(db)
	// products é uma variável que armazena uma fatia de instâncias da struct Product, que representa todos os produtos encontrados no banco de dados.
	// A função FindAll é uma função do repositório de produto que busca todos os produtos no banco de dados e os ordena de acordo com um parâmetro.
	// O parâmetro "asc" indica que os produtos devem ser ordenados em ordem ascendente pelo nome.
	products, err := productDB.FindAll("asc")

	// assert.NoError verifica se o valor passado como argumento é nil, ou seja, se não há erro ao buscar os produtos no banco de dados.
	assert.NoError(t, err)
	// assert.Len verifica se o valor passado como primeiro argumento tem o comprimento esperado, que é passado como segundo argumento.
	// Neste caso, verifica se a fatia de produtos tem o mesmo comprimento que a fatia de produtos, o que deve ser verdadeiro.
	assert.Len(t, products, len(products))
	// assert.Equal verifica se os valores passados como argumentos são iguais, ou seja, se os atributos dos produtos são os esperados.
	// Neste caso, verifica se os nomes dos produtos na posição 0, 9 e 20 da fatia são "Product 1", "Product 10" e "Product 21", respectivamente.
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)
	assert.Equal(t, "Product 21", products[20].Name)
}

// TestUpdateProduct é uma função de teste que verifica se a função Update do repositório de produto funciona corretamente.
func TestUpdateProduct(t *testing.T) {
	// dsn é uma string que contém as informações de conexão com o banco de dados MySQL.
	dsn := "root:root@tcp(localhost:3305)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	// db é uma variável que armazena a instância do GORM, que é um ORM (Object-Relational Mapping) para Go.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// Se houver um erro ao abrir a conexão com o banco de dados, o teste é encerrado com uma mensagem de erro.
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// db.AutoMigrate é uma função do GORM que cria ou atualiza automaticamente as tabelas do banco de dados de acordo com as structs definidas no pacote entity.
	db.AutoMigrate(&entity.Product{})

	// product é uma variável que armazena uma instância da struct Product, que representa um produto no sistema.
	// A função NewProduct é uma função do pacote entity que cria um novo produto com um nome e um preço.
	product, err := entity.NewProduct("Leite de vaca x", 13.00)

	// assert é uma biblioteca que fornece funções para verificar se as condições esperadas são atendidas nos testes.
	// assert.Nil verifica se o valor passado como argumento é nil, ou seja, se não há erro ao criar o produto.
	assert.Nil(t, err)

	// db.Create é uma função do GORM que insere um registro no banco de dados.
	db.Create(product)

	// productDb é uma variável que armazena uma instância do repositório de produto, que é uma camada de abstração que permite interagir com o banco de dados usando o GORM.
	productDb := NewProduct(db)

	// Altera o nome do produto para "leite de vaca w".
	product.Name = "leite de vaca w"

	// A função Update é uma função do repositório de produto que atualiza um produto existente no banco de dados.
	err = productDb.Update(product)

	// assert.NoError verifica se o valor passado como argumento é nil, ou seja, se não há erro ao atualizar o produto no banco de dados.
	assert.NoError(t, err)

	// Busca o produto atualizado no banco de dados pelo seu ID.
	product, err = productDb.FindById(product.ID.String())

	// assert.Nil verifica se o valor passado como argumento é nil, ou seja, se não há erro ao buscar o produto no banco de dados.
	assert.Nil(t, err)

	// assert.NotEmpty verifica se o valor passado como argumento não está vazio, ou seja, se o produto tem um ID válido.
	assert.NotEmpty(t, product.ID)

	// assert.Equal verifica se os valores passados como argumentos são iguais, ou seja, se o nome do produto é o esperado após a atualização.
	assert.Equal(t, product.Name, "leite de vaca w")
}

func TestDeleteProdutc(t *testing.T) {
	dsn := "root:root@tcp(localhost:3305)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("porco 1", 12.0)

	assert.NoError(t, err)

	db.Create(product)

	ProductDb := NewProduct(db)

	err = ProductDb.Delete(product.ID.String())

	assert.NoError(t, err)

	product, err = ProductDb.FindById(product.ID.String())

	assert.Error(t, err)

	assert.Empty(t, product)
}
