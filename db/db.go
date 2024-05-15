package db

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db         *sql.DB
	once       sync.Once
	dbInitFunc func() (*sql.DB, error) // inicialização da conexão
)

// inicializa o pool de conexões com o banco de dados
func initDB() (*sql.DB, error) {
	conexao := "user=postgres dbname=loja password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// retorna o pool de conexões e cria se necessário
func PoolDeConexoes() *sql.DB {
	once.Do(func() {
		// se a função de inicialização não for fornecida, usa a função padrão de inicialização
		if dbInitFunc == nil {
			dbInitFunc = initDB
		}

		// inicia o pool de conexões
		var err error
		db, err = dbInitFunc()
		if err != nil {
			log.Fatal(err)
		}
	})

	return db
}

func ConectaComBancoDeDados() *sql.DB {
	return PoolDeConexoes()
}
