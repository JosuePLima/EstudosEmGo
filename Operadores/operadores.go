/* Em Go, os operadores são símbolos que permitem realizar operações em valores ou variáveis. Eles podem ser categorizados em diferentes tipos, dependendo da sua função.

Aqui estão alguns dos operadores mais comuns em Go:

Operadores Aritméticos:
São usados para realizar operações matemáticas em valores numéricos. Eles incluem:
+ (adição)
- (subtração)
* (multiplicação)
/ (divisão)
% (módulo)
Operadores de Comparação:
São usados para comparar valores e produzir um valor booleano (verdadeiro ou falso). Eles incluem:
== (igual a)
!= (diferente de)
< (menor que)
> (maior que)
<= (menor ou igual a)
>= (maior ou igual a)
Operadores Lógicos:
São usados para combinar valores booleanos e produzir um valor booleano resultante. Eles incluem:
&& (E lógico)
|| (OU lógico)
! (NÃO lógico)
Operadores de Atribuição:
São usados para atribuir valores a variáveis. Eles incluem:
= (atribuição simples)
+= (adicionar e atribuir)
-= (subtrair e atribuir)
*= (multiplicar e atribuir)
/= (dividir e atribuir)
%= (módulo e atribuir)
Além desses operadores, Go também possui operadores bit a bit, que são usados para manipular valores em nível de bit, operadores de incremento e decremento, que são usados para incrementar ou decrementar valores, e operadores de ponteiro, que são usados para manipular ponteiros de memória.

É importante notar que a precedência dos operadores em Go segue as regras matemáticas padrão. Se houver dúvidas sobre a ordem em que as operações são realizadas, é possível usar parênteses para controlar a ordem de execução.


Os operadores relacionais em Go são usados para comparar valores e produzir um resultado booleano (verdadeiro ou falso). Aqui estão os operadores relacionais em Go:

== - igual a: retorna verdadeiro se os valores comparados forem iguais.
!= - diferente de: retorna verdadeiro se os valores comparados forem diferentes.
< - menor que: retorna verdadeiro se o valor da esquerda for menor que o valor da direita.
> - maior que: retorna verdadeiro se o valor da esquerda for maior que o valor da direita.
<= - menor ou igual a: retorna verdadeiro se o valor da esquerda for menor ou igual ao valor da direita.
>= - maior ou igual a: retorna verdadeiro se o valor da esquerda for maior ou igual ao valor da direita.
Os operadores relacionais são frequentemente usados em estruturas de controle de fluxo, como if/else e loops, para tomar decisões com base em comparações de valores. Por exemplo:
*/

a := 10
b := 5

if a > b {
	fmt.Println("a é maior que b")
} else if a < b {
	fmt.Println("b é maior que a")
} else {
	fmt.Println("a é igual a b")
}

/* Neste exemplo, o operador relacional > é usado para comparar o valor de a com o valor de b e decidir qual mensagem deve ser impressa. */





