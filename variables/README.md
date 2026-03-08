# Variáveis em Go
- Primeiramente, não podemos declarar uma var ou const e não inicializar ela, nem compila

## **Declaração explícita**

- sintaxe `var <nome-da-variável> string = <valor>`

## **Multiplas declarações**
```go
var(
    <nome-da-variavel> string = "Ken Thompson"
    <nome-da-variavel> string = "jhon"
)
fmt.println(<nome das variaveis separando-as por vírgulas>)
```
## **Constantes**

- Mesmos conceitos só muda a palavra reservada de `var` para `const`

- sintaxe `const <nome-da-variável> string = <valor>`

## **Multiplas declarações**
```go
const(
    <nome-da-variavel> string = "Ken Thompson"
    <nome-da-variavel> string = "jhon"
)
```

## **Inferencia**

- Só funcionda dentro de `func(){...}`

- sintaxe `<nome-da-variavel := "<valor"` ex p/ str


