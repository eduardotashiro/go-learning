# Tipos de Dados

## Num inteiros

- int8, int16, int32, int64
> numeros inteiros com tamanhos especificos, no caso quantos bits ele usa

**vou usar apenas quando TENHO CERTEZA de quanto espaço vou precisar (micro otimização que fala)**

```go
var numero int16 = 100000000000000000000000 //overflow :$
```

int
> usa a arquitetura do meu PC como base, mais simples, go cuida sozinho do tamanho


## uint [Unsigned Integer]

melhora a capacidade para guardar numero, desde que seja **positivo**

```go
var numero uint = 111
fmt.Println(numero) //111
```

## rune

É um alias para `int32`
```go
var numero rune = 111
fmt.Println(numero) //111
```

## byte

é um alias para `uint8`

```go
var num4 byte = 666
```

## Error

```go
var erro1 error //<nil>
var erro1 error = errors.New("deu ruim !!!")
```

***OBS:*** SE NÃO INICIALIZA, GO COLOCA UM DEFAULT

- string: "" vazio
- int: 0
- float: 0.0
- bool: false 
error: nil



- Erro é um tipo ! isso me causou certa estranhesa, mas tudo bem

