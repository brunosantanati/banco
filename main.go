package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	aula1()
}

func aula1() {
	fmt.Println("######### Aula 1 #########")
	usarVariaveis()
	usarStruct()
}

func usarVariaveis() {
	//usando variáveis
	var titular string = "Bruno"
	var numeroAgencia int = 589
	var numeroConta int = 123456
	var saldo float64 = 125.5

	//printar variáveis simples
	fmt.Println(titular, numeroAgencia, numeroConta, saldo)
}

func usarStruct() {
	//cria uma instância da struct ContaCorrente com valores default
	fmt.Println(ContaCorrente{})

	//cria uma instância passando nomes dos campos
	contaDoGuilherme := ContaCorrente{titular: "Guilherme",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	//cria uma instância passando todos campos na ordem, então não precisa passar os nomes
	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)
}
