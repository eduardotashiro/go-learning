package app

import "github.com/urfave/cli"

// gerar retorna app command line
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Net"
	return app
}
