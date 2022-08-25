package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	Porta = 5000
	StringConexao = ""

)

//Iniciando as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro := strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 3000
		fmt.Sprintf("Porta modificada para %s", Porta)
	}
	fmt.Println("Conectando na porta:", Porta)

	StringConexao = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}
