package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente ou valor do saque negativo"
	}
}

func main() {
	aula1()

	fmt.Println()
	aula2()
}

func aula1() {
	fmt.Println("######### Aula 1 #########")

	fmt.Println()
	usarVariaveis()

	fmt.Println()
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

func aula2() {
	fmt.Println("######### Aula 2 #########")

	fmt.Println()
	usarStructComPonteiro()

	fmt.Println()
	compararTipos()

	fmt.Println()
	sacar()
}

func usarStructComPonteiro() {
	//cria uma variável que é um ponteiro
	var contaDaCris *ContaCorrente
	//cria uma instância do struct ContaCorrente e retorna um ponteiro para ela
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	fmt.Println(contaDaCris)
	fmt.Println(&contaDaCris)
	fmt.Println(*contaDaCris)
}

func compararTipos() {
	//comparação 1
	contaDoGuilherme := ContaCorrente{titular: "Guilherme",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}
	contaDoGuilherme2 := ContaCorrente{titular: "Guilherme",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}
	fmt.Println(contaDoGuilherme == contaDoGuilherme2)

	//comparação 2
	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
	contaDaBruna2 := ContaCorrente{"Bruna", 222, 111222, 200}
	fmt.Println(contaDaBruna == contaDaBruna2)

	//comparação 3
	fmt.Println()
	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.titular = "Cris"
	contaDaCris2.saldo = 500

	fmt.Println(&contaDaCris, &contaDaCris2)
	fmt.Println(contaDaCris, contaDaCris2)
	fmt.Println(contaDaCris == contaDaCris2)
	fmt.Println(&contaDaCris == &contaDaCris2)
	fmt.Println(*contaDaCris == *contaDaCris2)
}

func sacar() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)

	fmt.Println(contaDaSilvia.Sacar(400))
	fmt.Println(contaDaSilvia.saldo)

}
