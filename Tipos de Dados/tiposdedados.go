/* Em Go, os tipos de dados são divididos em duas categorias principais: tipos básicos e tipos compostos.

- Os tipos básicos incluem:

-- bool: que representa um valor booleano verdadeiro ou falso;
string: que representa uma sequência de caracteres;
int, int8, int16, int32, int64: que representam números inteiros com diferentes tamanhos;
uint, uint8 (ou byte), uint16, uint32, uint64, uintptr: que representam números inteiros sem sinal com diferentes tamanhos;
float32, float64: que representam números de ponto flutuante com diferentes precisões;
complex64, complex128: que representam números complexos com diferentes precisões.
Os tipos compostos incluem:

array: que é uma sequência fixa de elementos do mesmo tipo;
slice: que é uma sequência variável de elementos do mesmo tipo, cujo tamanho pode ser alterado;
map: que é uma coleção de pares chave-valor, em que as chaves são exclusivas e os valores podem ser de qualquer tipo;
struct: que é uma coleção de campos nomeados, em que cada campo pode ter um tipo diferente;
interface: que é um tipo genérico que pode representar qualquer tipo de valor;
function: que é um tipo que representa uma função que pode receber e retornar valores de tipos específicos.
Além desses tipos básicos e compostos, Go também permite a definição de tipos personalizados, usando a palavra-chave type. Por exemplo, podemos definir um novo tipo Nome para representar nomes de pessoas:
*/

type Nome string

// Agora, podemos usar "Nome" como um tipo em nosso programa, como em://

var nome Nome = "João"

/* Isso torna nosso código mais legível e expressivo, permitindo que usemos nomes significativos em vez de tipos genéricos. */

