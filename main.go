package main

import (
	"fmt"

	"github.com/brunosantanati/banco/contas"
)

func main() {
	aula1()

	fmt.Println()
	aula2()

	fmt.Println()
	aula3()
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
	fmt.Println(contas.ContaCorrente{})

	//cria uma instância passando nomes dos campos
	contaDoGuilherme := contas.ContaCorrente{Titular: "Guilherme",
		NumeroAgencia: 589, NumeroConta: 123456, Saldo: 125.5}

	//cria uma instância passando todos campos na ordem, então não precisa passar os nomes
	contaDaBruna := contas.ContaCorrente{"Bruna", 222, 111222, 200}

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
	var contaDaCris *contas.ContaCorrente
	//cria uma instância do struct ContaCorrente e retorna um ponteiro para ela
	contaDaCris = new(contas.ContaCorrente)
	contaDaCris.Titular = "Cris"
	contaDaCris.Saldo = 500

	fmt.Println(contaDaCris)
	fmt.Println(&contaDaCris)
	fmt.Println(*contaDaCris)
}

func compararTipos() {
	//comparação 1
	contaDoGuilherme := contas.ContaCorrente{Titular: "Guilherme",
		NumeroAgencia: 589, NumeroConta: 123456, Saldo: 125.5}
	contaDoGuilherme2 := contas.ContaCorrente{Titular: "Guilherme",
		NumeroAgencia: 589, NumeroConta: 123456, Saldo: 125.5}
	fmt.Println(contaDoGuilherme == contaDoGuilherme2)

	//comparação 2
	contaDaBruna := contas.ContaCorrente{"Bruna", 222, 111222, 200}
	contaDaBruna2 := contas.ContaCorrente{"Bruna", 222, 111222, 200}
	fmt.Println(contaDaBruna == contaDaBruna2)

	//comparação 3
	fmt.Println()
	var contaDaCris *contas.ContaCorrente
	contaDaCris = new(contas.ContaCorrente)
	contaDaCris.Titular = "Cris"
	contaDaCris.Saldo = 500

	var contaDaCris2 *contas.ContaCorrente
	contaDaCris2 = new(contas.ContaCorrente)
	contaDaCris2.Titular = "Cris"
	contaDaCris2.Saldo = 500

	fmt.Println(&contaDaCris, &contaDaCris2)
	fmt.Println(contaDaCris, contaDaCris2)
	fmt.Println(contaDaCris == contaDaCris2)
	fmt.Println(&contaDaCris == &contaDaCris2)
	fmt.Println(*contaDaCris == *contaDaCris2)
}

func sacar() {
	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.Titular = "Silvia"
	contaDaSilvia.Saldo = 500

	fmt.Println(contaDaSilvia.Saldo)

	fmt.Println(contaDaSilvia.Sacar(400))
	fmt.Println(contaDaSilvia.Saldo)
}

func aula3() {
	fmt.Println("######### Aula 3 #########")

	fmt.Println()
	multiplosRetornos()

	fmt.Println()
	transferirEntreContas()
}

func multiplosRetornos() {
	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.Titular = "Silvia"
	contaDaSilvia.Saldo = 500

	fmt.Println(contaDaSilvia.Saldo)
	status, valor := contaDaSilvia.Depositar(2000)
	fmt.Println(status, valor)

	fmt.Println(contaDaSilvia.Depositar(2000))
}

func transferirEntreContas() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDoGustavo.Tranferir(50, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}
