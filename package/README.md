# Package - Mód e pacotes em Go

Organizando código em **pacotes**, usando **módulos**, além de **pacotes externos**, que por sinal achei incrivel

## Conceitos

- As funções vão ser publicas e privadas, dependendo se o nome delas começam com maiusculas ou não, começou com maiúscula (PÚBLICA), começou com minúscula(PRIVADA), o macete é jogar uma privada dentro de uma func publica e exportar a publica que funciona igual

### Criando um módulo

```
go mod init <nome do modulo>
```
### Compilando

```bash
go build
```
Isso compila todo o código e cria um executável

-- *deve ser por isso que go e docker casam demais* --

## `helpers.go`

```go
package utils 

import "fmt"

// Escrever registra uma mensagem na tela
func Escrever() { // Começa com MAIÚSCULA para ser PÚBLICA
	fmt.Println("escrevendo do pacote utils")
	escrever2()
}
```

- **package utils** define que este arquivo pertence ao pacote `utils`
===============
- **`func Escrever(){...}`** Função publica (começa com MAIÚSCULA), pode ser importada para outros pacotes
===============
- **`escrever2()`** chama uma função privada do mesmo pacote (UTILS)
===============

## `main go`

```go
package main

import (
    "modulo/utils"
    "fmt"
    "github.com/badoux/checkmail" // CMD: go get gitgub.com/badoux/checkmail
)

func main(){
    fmt.Println("escrevendo do arquivo main")
}
ultil.Escrever()

//usando pacotes externos
erro:= checkmail.ValidateFormat("bugzilla@gmail.com")
fmt.Println(erro)
```
- **package main**: define o pacote principal(ponto de entrada)

- **`import(...)`**: Importa três pacotes:
- `"modulo/utils"`: pacote local
- `"fmt"`: pacote padrão, esta em tudo..
- `"github.com/badoux/checkmail"`: pacote externo do github

- `"utils.Escrever()"`: :fire: a função pública do pacote utils
- **`checkmail.ValidateFormat`**: valida um email, deve ter um regex por de baixo dos panos

## Instalando um pacote externo:

```bash
go get github.com/badoux/checkmail
```
isso baixa o pacote e add no go.mod direto

---

### `go.mod`

```go.mod
module modulo

go 1.25.0 //versão do go

require github.com/badoux/checkmail v1.2.4 // indirect
```

- **`module modulo`**: define o nome do módulo(deve ser o mesmo valor usado no import)
> esta estranho[module modulo] pq dei um go mod init modulo

### `go.sum`

```go.sum
github.com/badoux/checkmail v1.2.4 h1:4zMjdYDjE2Q7xF06VNfyN8P9JGU7epLjNb+Yu5OThVI=
github.com/badoux/checkmail v1.2.4/go.mod h1:XroCOBU5zzZJcLvgwU15I+2xXyCdTWXyR9MGfRhBYy0=
```

- **arquivo de verificação**:hashes criptográficos (SHA-XXX) de cada dependência
- **NÃO MEXER**
- **segurança**: garante que ngm modificou os pacotes baixados(espero estudar mais afundo isso)


### IMPORTANTE

- go mod init <nome> : cria novo modulo
- go get <pacote> : baixa uma dependência
- go build : compila o projeto em um binário (só por jogar no docker)
- go run main.go : executa sem compilar

> basicamente go run .. para desenvolvimento
> basicamente go build ... para prod