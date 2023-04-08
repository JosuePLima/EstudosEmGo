/* Em Golang, a herança é implementada através da composição, em vez da herança clássica encontrada em outras linguagens de programação orientada a objetos. Isso significa que em vez de criar uma classe filha que herda propriedades e métodos da classe pai, você cria um tipo estruturado que incorpora outro tipo.

Para isso, é possível utilizar a declaração type com uma struct anônima (sem nome) para incorporar os campos da struct pai. Por exemplo: */

type Animal struct {
    name string
    age  int
}

type Dog struct {
    Animal  // Incorpora os campos da struct Animal
    breed string
}
/* Na declaração do tipo Dog, utilizamos o nome do tipo Animal em vez de declarar seus campos individualmente. Assim, a struct Dog possui os mesmos campos que a struct Animal (name e age), além do campo breed.

É possível acessar os campos da struct pai diretamente através da instância do tipo filho: */

dog := Dog{Animal{"Fido", 3}, "Golden Retriever"}
fmt.Println(dog.name)  // "Fido"
fmt.Println(dog.age)   // 3

/* Além da composição, em Golang também é possível utilizar interfaces para simular herança. Uma interface define um conjunto de métodos que uma estrutura pode implementar, permitindo que tipos diferentes compartilhem comportamentos comuns. Quando um tipo implementa todos os métodos de uma interface, ele pode ser considerado uma implementação dessa interface.

Por exemplo, podemos definir uma interface Animal que define um método MakeSound() e uma estrutura Dog que implementa essa interface: */

type Animal interface {
    MakeSound() string
}

type Dog struct {
    name string
}

func (d Dog) MakeSound() string {
    return "Woof!"
}

/* Assim, a struct Dog pode ser tratada como uma instância da interface Animal, e chamando o método MakeSound() na instância de Dog retornará "Woof!". Dessa forma, as interfaces permitem que os tipos implementem comportamentos em comum sem necessariamente compartilhar campos de dados. */
