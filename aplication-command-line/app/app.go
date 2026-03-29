package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// gerar retorna app-command-line pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Net"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços da net",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name: "servidores",
			Usage: "Busca nome do servidor na net",
			Flags: flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range servidores {
		fmt.Println(s.Host)
	}
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}


