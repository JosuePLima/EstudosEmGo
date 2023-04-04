/* Na linguagem GO, as funções são uma parte fundamental. Elas são usadas para agrupar um conjunto de instruões que podem ser reutilizadas em diferentes partes do programa. A sintaxe para definir uma função em GO é a seguinte: */

func nomeDaFuncao(argumentos) tipoDeRetorno {
	// corpo da função
	return valorDeRetorno
}

/* Aqui, "nomeDaFuncao" é o nome da função que você está definindo, "argumentos" são valores de entrada que a função espera, "tipoDeRetorno" é o tipo de dado que a função retorna e "valorDeRetorno" é o valor que a função retorna. Por exemplo, se você quiser definir uma função que recebe dois números inteiros e retorna sua soma, você pode fazer isso da seguinte forma: */

func soma(a int, b int) int {
	return a + b
}

/* Você também pode definir funções sem argumentos e sem valor de retorno: */

func imprimirMensagem() {
	fmt.Println("Olá, mundo!")
}

/* E para chamar uma função, você simplesmente escreve o nome dela seguido dos argumentos( se houver): */

soma := soma(2,3) // retorna 5
imprimirMensagem() // imprime "Olá, mundo!"

/* Além disso, em GO, é possível definir funções anônimas, que são funções que não tem um nome específico e são usadas principalmente como valores de variáveis ou como argumentos para outras funções. */

func main() {
		// define uma função anônima e a atribui a uma variável
	multiplicar := func(a, b int) int {
		return a* b
	}

	resultado := multiplicar(3,4)
	fmt.Println(resultado) //imprime 12

}