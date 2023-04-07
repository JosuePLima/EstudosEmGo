/* Uma struct é uma coleção de campos com tipos de dados diferentes, que são agrupados em uma única unidade. É semelhante a uma clase em outras linguagens orientadas a objetos. Aqui está um exemplo simples de uma definição de struct em GO: */

type Pessoa struct {
	nome string
	idade int
	altura float64
}

/* Neste exemplo, estamos definindo um struct chamado "Pessoa" com três campos: nome (string), idade(int) e altura (float64). Agora podemos criar uma variável do tipo Pessoa e inicializa-la com valores: */

var pessoa1 Pessoa
pessoa1.nome = "Maria"
pessoa1.idade = 25
pessoa1.altura= 1.70

/* Também é possível inicializar uma struct na declaração: */

pessoa2 := Pessoa{"João", 30, 1.80}

/* Podemos acessar os campos de uma struct usando a notação de ponto: */

fmt.Println(pessoa1.nome) // imprime "Maria"
fmt.Println(pessoa2.idade) // imprime 30

/* Além disso, podemos criar um ponteiro para uma struct usando o operador "&": */

pessoa3 := &Pessoa{"Lucas", 28, 1.75}
fmt.Println((*pessoa3).altura) // imprime 1.75

/* Por padrão, os campos de uma struct são exportado(visíveis fora do pacote) se a primeira letra do nome do campos for maiúscula. Caso contrario, eles são considerado não exportados. Em GO, a exportação de um campo é controlada pelo seu nome. Se o nome do campo começa com uma letra maiúscula, ele é considerado exportado e pode ser acessado por outros pacotes. Caso contrário, o campor é considerado não exportado e só pode ser acessado pelo próprio pacote. Por exemplo, suponha que temos a seguinte definição de struct em um pacote chamado "pessoa": */
package pessoa

type Pessoa struct{
	Nome string // exportado
	idade int // não exportado
}

/* Nesse caso, o campo "Nome" é exportado pq a primeira letra de seu nome é maiúscula.Isso significa que outros pacotes podem acessá-lo diretamente. Diferente do campo "idade", que só pode ser acessado pelo prórprio pacote "pessoa".
Agora, suponha que temos outro pacote chamado "main" que deseja criar uma instância da struct "Pessoa" e definir seus campos: */

package main

import "pessoa"

func main () {
	var p pessoa.Pessoa
	p.Nome = "Maria" //ok, campo exportado
	p.idade = 25 // erro, camponão exportado
}

/* Neste exemplo, podemos acessar o campo "Nome" diretamente porque ele é exportado. Por outro lado, o campo "idade" não pode ser acessado diretamente porque é não exportado e só pode ser acessado pelo próprio pacote "pessoa".

Essa convenção de nomenclatura é útil para controlar a visibilidade dos campos de uma struct em diferentes pacotes. Ela ajuda a evitar que os campos sejam alterados inadvertidamente por outros pacotes e ajuda a manter a integridade dos dados.*/