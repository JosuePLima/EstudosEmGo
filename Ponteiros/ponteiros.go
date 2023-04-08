/* Em Go, os ponteiros são usados para referenciar a memória de uma variável em vez de seu valor. Um ponteiro é um valor que contém o endereço de memória de uma variável. Quando você tem um ponteiro para uma variável, você pode acessar ou modificar o valor da variável referenciada diretamente, sem precisar copiar o valor para outra variável.

Em Go, um ponteiro é denotado pelo operador "&" (ampersand) antes de uma variável e pelo operador "*" (asterisco) antes de um ponteiro. Para declarar um ponteiro, usamos a seguinte sintaxe:*/

var p *int

/* Nesse caso, declaramos um ponteiro para um inteiro. Podemos inicializar esse ponteiro para apontar para uma variável específica usando o operador "&", como mostrado abaixo:*/

var i int = 42
p = &i

/* Nesse caso, inicializamos o ponteiro "p" com o endereço de memória da variável "i". Para acessar o valor da variável referenciada pelo ponteiro, usamos o operador "*" antes do ponteiro, como mostrado abaixo: */

fmt.Println(*p) // Output: 42

/* Nesse caso, usamos o operador "*" para desreferenciar o ponteiro "p" e acessar o valor da variável "i". Também podemos modificar o valor da variável referenciada pelo ponteiro, atribuindo um novo valor ao endereço de memória referenciado pelo ponteiro, como mostrado abaixo: */

*p = 21
fmt.Println(i) // Output: 21

/* Nesse caso, usamos o operador "*" para desreferenciar o ponteiro "p" e atribuir um novo valor ao endereço de memória referenciado pelo ponteiro, que modifica o valor da variável "i". */



