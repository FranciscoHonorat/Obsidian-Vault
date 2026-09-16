Para essa anotação vamos usar como base o roadmap e o livro "Learning Go" (Jon Bodner), que é um livro de referência para iniciantes e intermediários, com foco em desenvolvimento de habilidades e compreensão profunda da linguagem.

# 1. var vs :=
Em Go tem várias maneiras de declarar variáveis, e a ideia do livro é que cada forma comunica uma intenção diferente, e que cada forma tem um propósito específico. Então, vamos explorar cada uma delas.. A forma mais completa tem palavras-chave, tipo e valor, mas também podemos declarar variáveis sem tipo, ou sem valor, ou sem palavras-chave. E cada uma dessas formas tem um propósito específico.
```go
var x int = 10 // declaração completa.
```
Se o tipo do lado direito for o mesmo do lado esquerdo, podemos omitir o tipo, e o compilador vai inferir o tipo da variável.
```go
var x = 10 // declaração sem tipo.
```
Se você quer o zero value de um tipo, você pode declarar a variável sem valor, e o compilador vai atribuir o zero value do tipo da variável.
```go
var x int // declaração sem valor, x é inicializado com o zero value do tipo int
```
Também é possível declarar várias de uma vez:
```go
var x, y, z int = 1, 2, 3 // declaração de várias variáveis com tipo e valor.
```
e existe a lista de declaração, uito usada no nível de pacote:
```go
var (
    x int = 1
    y = 20
    z int = 30
    d, e = 40, "hello"
    f, g string
)
```

Dentro de funções, podemos declarar variáveis usando a forma curta de declaração, que é a mais comum:
```go
x := 10 // declaração curta, o tipo é inferido pelo compilador.
```
O := possui duas características importantes: ele declara e inicializa a variável, e o tipo é inferido pelo compilador. Além disso, ele só pode ser usado dentro de funções, e não no nível de pacote. Além disso é possível reatribuir valores a variáveis já declaradas, mas não é possível reatribuir valores a constantes. As constantes são declaradas usando a palavra-chave const, e elas não podem ser alteradas depois de serem declaradas.
```go
x := 10 // declaração curta, o tipo é inferido pelo compilador.
x, y := 20, "hello" // declaração curta de várias variáveis, o tipo é inferido pelo compilador.

const pi = 3.14 // declaração de constante.
```

Existe três situações em que var comunica melhor:

1. Quando você quer o zero value intencionalmente. var count int deixa claro o zero é proposital. count := 0 não deixa claro se o zero é proposital ou não.
2. Quando o tipo padrão do literal não é o que você quer. var x int = 10 deixa claro que você quer um int, mesmo que 10 seja um literal int. x := 10 deixa claro que você quer um int, mas se você quiser um float64, você teria que escrever x := 10.0.
3. Quando o := pode criar variáveis novas sem querer (shadowing). Se você tem uma variável x já declarada, e você faz x := 10 dentro de uma função, você está criando uma nova variável x, e não alterando a variável x já declarada. Isso pode levar a bugs difíceis de encontrar.

Sobre a questão de múltiplas variáveis na mesma linha: o livro reomenda usar esse estilo só quando os valores vêm de uma função com múltiplos retornos ou do idioma comma ok. Por exemplo, quando você quer declarar várias variáveis com valores de uma função que retorna múltiplos valores, você pode fazer:
```go
x, y := someFunction() // declaração de várias variáveis com valores de uma função que retorna múltiplos valores.
```

Sobre variáveis no nível de pacote, o livro recomenda evitar variáveis mutáveis no nível de pacote, e usar constantes sempre que possível. Variáveis mutáveis no nível de pacote podem levar a problemas de concorrência e dificultar a manutenção do código. Se você precisa de uma variável mutável no nível de pacote, considere encapsulá-la em uma função ou struct, ou use um pacote separado para gerenciar o estado compartilhado.

As variáveis não usadas, o Go tem uma regra peculiar: toda variável local declarada precisa ser lida, senão o código não compila. Isso é uma forma de evitar variáveis declaradas e não usadas, que podem levar a bugs e código confuso. Se você declarar uma variável e não usá-la, o compilador vai reclamar. Isso é diferente de outras linguagens, onde variáveis não usadas podem ser ignoradas.

Mas a checagem não é exaustiva. Basta uma leitura para o compilador considerar a variável como usada. Por exemplo, se você declarar uma variável e apenas imprimir seu valor, o compilador vai considerar a variável como usada, mesmo que você não faça nada com ela depois. Isso é uma forma de evitar variáveis declaradas e não usadas, mas também pode levar a código confuso se você não tiver cuidado.

Nem o compilador e nem o go vet pegam isso. O linter ineffassing (incluído no golangci-lint) pega isso, e é uma boa prática usar o linter para pegar variáveis não usadas. O linter é uma ferramenta que analisa o código e procura por problemas de estilo, bugs e más práticas. Ele pode ser configurado para rodar automaticamente em seu editor de código ou como parte do processo de build. Se precisa descartar um valor, use o identificador especial _ (underscore), que é usado para descartar valores que não serão usados. Por exemplo, se você quer chamar uma função que retorna dois valores, mas você só quer o primeiro valor, você pode fazer:
```go
x, _ := someFunction() // descarta o segundo valor retornado pela função.
```

# 2. Zero Values

Em Go, toda variável declarada sem valor recebe automaticamente o zero value do seu tipo. O zero value é o valor padrão que uma variável de determinado tipo assume quando não é explicitamente inicializada. Isso significa que você não precisa se preocupar em inicializar variáveis com valores específicos, pois o Go garante que elas terão um valor consistente e previsível.
```go
var i int
var f float64
var b bool
var s string
var p *int
var sl []int
var m map[string]int

fmt.Println(i, f, b, s == "", p == nil, sl == nil, m == nil)
// 0 0 false true true true true
```
Note que o zero value de tipos numéricos é 0, para booleanos é false, para strings é uma string vazia "", e para ponteiros, slices e maps é nil. Isso ajuda a evitar erros comuns relacionados a variáveis não inicializadas e torna o código mais seguro e confiável.

Zero values de compostos

Arrys e structs são inicializados recursivamente, ou seja, cada elemento do array ou campo da struct é inicializado com o zero value do seu tipo. Por exemplo:
```go
type Person struct {
    Name string
    Age  int
}

var p Person
var arr [3]int

fmt.Println(p.Name, p.Age) // "" 0
fmt.Println(arr) // [0 0 0]
```
O zero value é útil ou perigoso, dependendo do tipo. Por exemplo, o zero value de um ponteiro é nil, e se você tentar acessar um campo de um ponteiro nil, você vai ter um panic. Por outro lado, o zero value de uma string é uma string vazia "", e se você tentar acessar um campo de uma string vazia, você não vai ter um panic, mas você vai ter um comportamento inesperado. Então, é importante entender o zero value de cada tipo e como ele pode afetar o comportamento do seu código.
```go
var sl []int
sl = append(sl, 1) // adiciona um elemento ao slice, mesmo que o slice seja nil.
fmt.Println(len(sl)) // len de slice nil é 0, sem panic 

var m map[string]int
fmt.Println(m["key"]) // acessa um map nil, retorna o zero value do tipo do valor, sem panic
m["key"] = 1 // atribui um valor a um map nil, panic
```
Ou seja, um map nil é somente leitura na prática, e você não pode adicionar elementos a ele. Se você tentar adicionar um elemento a um map nil, você vai ter um panic. Por outro lado, um slice nil é somente leitura na prática, mas você pode adicionar elementos a ele usando a função append. Se você tentar acessar um elemento de um slice nil, você vai ter um panic. Então, é importante entender o zero value de cada tipo e como ele pode afetar o comportamento do seu código.

"Make the zero value useful" é uma frase do livro que resume a ideia de que o zero value é útil, mas você precisa entender como ele funciona e como ele pode afetar o comportamento do seu código. Se você entender o zero value de cada tipo, você pode escrever código mais seguro e confiável, evitando erros comuns relacionados a variáveis não inicializadas.
```go
var mu sync.Mutex // zero value de sync.Mutex é um mutex desbloqueado, pronto para uso.

var buf bytes.Buffer // zero value de bytes.Buffer é um buffer vazio, pronto para uso.
buf.WriteString("Hello, World!") // escreve no buffer, sem precisar inicializar o buffer.
```
O livro também mostra que o zero value tem um "custo": as vezes você precisa distinguir "não informado" de "informado como zero value". Por exemplo, se você tem uma struct com um campo int, e você quer saber se o campo foi informado ou não, você não pode usar o zero value como indicador, pois o zero value de int é 0. Nesse caso, você pode usar um ponteiro para int, e verificar se o ponteiro é nil ou não. Se o ponteiro for nil, significa que o campo não foi informado. Se o ponteiro não for nil, significa que o campo foi informado, e você pode acessar o valor do campo através do ponteiro.
```go
type Person struct {
    Name string
    Age  *int // ponteiro para int, para distinguir "não informado" de "informado como zero value".
}
var p Person
if p.Age == nil {
    fmt.Println("Age not informed")
} else {   
    fmt.Println("Age informed:", *p.Age)
}
```

# 3. Const e iota

```go
const x int64 = 10

const (
    idKey = "id"
    nameKey = "name"
)

const z = 20 * 10

func main() {
    const y = "hello"
    x = x + 1 // erro: cannot assign to x (declared const)
    y = "bye" // erro: cannot assign to y (declared const)
}
```
Constantes são declaradas usando a palavra-chave const, e elas não podem ser alteradas depois de serem declaradas. Constantes podem ser declaradas no nível de pacote ou dentro de funções, mas elas não podem ser alteradas depois de serem declaradas. Constantes podem ser de tipos numéricos, booleanos, strings ou tipos definidos pelo usuário. Constantes também podem ser usadas em expressões, e o compilador vai calcular o valor da expressão em tempo de compilação.

Constantes só pode conter:

1. literal de tipo numérico, booleano ou string;
2. expressões que envolvem apenas constantes, operadores e funções built-in que podem ser avaliadas em tempo de compilação; (complex, real, imag, len, cap, unsafe.Sizeof, unsafe.Alignof, unsafe.Offsetof)
3. tipos definidos pelo usuário que são baseados em tipos numéricos, booleanos ou strings.
4. runes
5. expressões compostas por operadores e pelos valores acima, como por exemplo: 1 + 2, 3 * 4, "hello" + "world", etc.

Portanto, não existe constante de slice, map, struct, array, ponteiro, função, interface ou canal. Constantes são valores imutáveis que são conhecidos em tempo de compilação, e não podem ser alterados em tempo de execução. Se você precisa de um valor que pode ser alterado em tempo de execução, você deve usar uma variável.

e Go não tem nenhuma forma de declarar uma variável imutável (não há equivalente ao readonly do C# ou ao final do Java). Se você precisa de uma variável imutável, você deve usar uma constante. Se você precisa de uma variável que pode ser alterada em tempo de execução, você deve usar uma variável.

Constantes são úteis para valores que não mudam, como por exemplo, valores de configuração, chaves de API, mensagens de erro, etc. Constantes também podem ser usadas para melhorar a legibilidade do código, evitando o uso de "magic numbers" ou strings literais espalhadas pelo código.

Constantes não usadas não geram erro de compilação, mas variáveis não usadas geram erro de compilação. Isso é uma forma de evitar variáveis declaradas e não usadas, que podem levar a bugs e código confuso. Se você declarar uma constante e não usá-la, o compilador não vai reclamar. Isso é diferente de variáveis, onde o compilador vai reclamar se você declarar uma variável e não usá-la.

Constantes não tipadas são constantes que não têm um tipo específico, e o compilador vai inferir o tipo da constante com base no contexto em que ela é usada. Constantes tipadas são constantes que têm um tipo específico, e o compilador vai verificar se o valor da constante é compatível com o tipo da constante.

Constante tipadas são úteis quando você quer garantir que o valor da constante seja do tipo correto, e evitar erros de conversão de tipo. Constantes não tipadas são úteis quando você quer que o valor da constante seja usado em diferentes contextos, e o compilador vai inferir o tipo da constante com base no contexto em que ela é usada.

O livro recomenda deixar constantes não tipadas por padrão, pois isso dá mais flexibilidade. Use constante tipadas quando quiser que o compilador force o tipo, o caso típico sendo enumerações com iota.

Dois detalhes extras sobre constantes não tipadas

Precisão arbitária: A especificação exige que o compilador represente constantes inteiras pelo menos 256bites. Então isso é válido:
```go
const grande = 1 << 1000 // constante inteira com precisão arbitrária, válida em Go.
const pequena = grande >> 98

fmt.Println(pequena) // 4
fmt.Println(grande) // Erro: constante overflows int
```
O erro só aparece quando a cosntante precisa virar um valor de tipo concreto, como int, float64, etc. Até lá, o compilador mantém a precisão arbitrária da constante.

Divisão inteiro vs de ponto flutuante: dependendo do tipo dos literais:
```go
const a = 1 / 2 // constante não tipada, o compilador vai inferir o tipo com base no contexto em que ela é usada.
const b = 1.0 / 2.0 // constante não tipada, o compilador vai inferir o tipo com base no contexto em que ela é usada.
fmt.Println(a) // 0, pois a é uma constante inteira, e a divisão inteira
fmt.Println(b) // 0.5, pois b é uma constante de ponto flutuante, e a divisão de ponto flutuante
```

IOTA

iota é um identificador predefinido que representa um contador de constantes. Ele é usado para criar constantes enumeradas, e o valor de iota é incrementado automaticamente a cada linha de declaração de constante dentro de um bloco const. O valor inicial de iota é 0, e ele é incrementado em 1 a cada linha de declaração de constante.

```go
type MailCategory int

const (
    Uncategorized MailCategory = iota // 0
    Personal                           // 1
    Spam                               // 2
    Social                            // 3
    Advertisement                      // 4
)
```
A regra de repetição implícita

Por que Personal vale 1 se não tem nada escrito? Porque em um bloco const, uma linha sem tipo e sem valor repete o tipo e a expressão da última linha que tinha ambos. A expressão repita é iota, e iota agora vale 1, então Personal vale 1. A mesma coisa acontece com Spam, Social e Advertisement, que valem 2, 3 e 4 respectivamente.

E essa regra vale mesmo sem iota:
```go
const (
    a = 1 // 1
    b     // 1, repete a expressão da última linha que tinha tipo e valor
    c     // 1, repete a expressão da última linha que tinha tipo e valor  
)
fmt.Println(a, b, c) // 1 1 1
```

iota incremeta mesmo quando é usado em expressões diferentes, e mesmo quando é usado em expressões que não envolvem iota. Por exemplo:
```go
const (
    Field1 = 0          // iota = 0, Field1 = 0
    Field2 = 1 + iota   // iota = 1, Field2 = 2
    Field3 = 20         // iota = 2, Field3 = 20
    Field4              // iota = 3, repete "20", Field4 = 20
    Field5 = iota       // iota = 4, Field5 = 4
)
// saída: 0 2 20 20 4
```
Pense em iota como o índice da linha dentro do bloco, independemente de ser usado ou não. O valor de iota é incrementado a cada linha, e ele pode ser usado em expressões diferentes, mas o valor de iota é sempre o índice da linha dentro do bloco.

### Padrões Comuns

Pular o zero. Se o zero não tem significao válido, atribua o primeiro valor a _ ou uma constante "inválida". Assim uma variável não inicializada (zero value!) fica fácil de detectar. Por exemplo, se você tem um tipo enumerado que representa os dias da semana, você pode atribuir o valor 0 a uma constante "InvalidDay", e atribuir os valores 1 a 7 aos dias da semana. Assim, se uma variável do tipo Day for inicializada com o zero value, ela vai ter o valor "InvalidDay", e você pode detectar facilmente que a variável não foi inicializada corretamente.
```go
type Day int

const (
    InvalidDay Day = iota // 0
    Sunday                 // 1
    Monday                 // 2
    Tuesday                // 3
    Wednesday              // 4
    Thursday               // 5
    Friday                 // 6
    Saturday               // 7
)

const (
    _ = iota // ignora o zero value
    Sunday
    Monday
)
```
Repare como isso conecta com o zero value: no exemplo acima, se uma variável do tipo Day for inicializada com o zero value, ela vai ter o valor "InvalidDay", e você pode detectar facilmente que a variável não foi inicializada corretamente. Isso é uma forma de evitar erros comuns relacionados a variáveis não inicializadas, e tornar o código mais seguro e confiável.

Bit Flags: iota é útil para criar bit flags, que são valores que representam um conjunto de opções, e cada opção é representada por um bit. Por exemplo, se você tem um tipo enumerado que representa as permissões de um arquivo, você pode usar iota para criar constantes que representam cada permissão, e cada constante vai ter um valor que é uma potência de 2. Assim, você pode combinar várias permissões usando o operador OR, e verificar se uma permissão está presente usando o operador AND.
```go
type FilePermission int

const (
    Read FilePermission = 1 << iota // 1
    Write                           // 2
    Execute                         // 4
)

p := Read | Write // combina as permissões Read e Write
if p&Read != 0 { // verifica se a permissão Read está presente
    fmt.Println("Read permission is present")
}
if p&Execute == 0 { // verifica se a permissão Execute está ausente
    fmt.Println("Execute permission is absent")
}
```
Uma forma de criar bit flags é usando iota para criar constantes que representam cada permissão, e cada constante vai ter um valor que é uma potência de 2. Assim, você pode combinar várias permissões usando o operador OR, e verificar se uma permissão está presente usando o operador AND.

Unidades de Tamanho: iota é útil para criar constantes que representam unidades de tamanho, como bytes, kilobytes, megabytes, gigabytes, etc. Por exemplo, você pode usar iota para criar constantes que representam cada unidade de tamanho, e cada constante vai ter um valor que é uma potência de 1024. Assim, você pode converter entre diferentes unidades de tamanho usando essas constantes.
```go
const (
    _ = iota // ignora o zero value
    KB = 1 << (10 * iota) // 1024
    MB                    // 1048576
    GB                    // 1073741824
)
fmt.Println(KB, MB, GB) // 1024 1048576 1073741824
```
Várias constante na mesma linha: iota é útil para criar várias constantes na mesma linha, usando expressões diferentes. Por exemplo, você pode usar iota para criar constantes que representam diferentes tipos de mensagens, e cada constante vai ter um valor que é uma expressão diferente. Assim, você pode criar várias constantes na mesma linha, usando expressões diferentes.
```go
const (
    a, b = iota + 1, iota + 2 // a = 1, b = 2
    c, d = iota + 3, iota + 4 // c = 4, d = 5
)
fmt.Println(a, b, c, d) // 1 2 4 5 
```

Os limites de iota são definidos pelo tipo da constante. Por exemplo, se você declarar uma constante do tipo int8, o valor de iota vai ser limitado ao intervalo de -128 a 127. Se você declarar uma constante do tipo uint8, o valor de iota vai ser limitado ao intervalo de 0 a 255. Se você declarar uma constante do tipo int16, o valor de iota vai ser limitado ao intervalo de -32768 a 32767. E assim por diante para os outros tipos numéricos.

- Nada impede criar valores fora da lista de iota compilar. O go não tem enums de verdade.
- Inserir uma linha no meio renumera todas as seguintes. Se essas constantes forem usadas em outro pacote, isso vai quebrar o código que depende desses valores. Então, é importante ter cuidado ao adicionar novas constantes em um bloco const que usa iota, e considerar adicionar novas constantes no final do bloco, para evitar quebrar o código que depende desses valores.

Daí o conselho de Danny van Heumem citado no livro: "If you are using iota to create a set of constants that will be used in other packages, consider adding a comment to the block of constants that explains the purpose of the constants and how they should be used. This can help other developers understand the intent of the constants and avoid breaking changes when adding new constants in the future."

# 4. Scopo e Shadowing

### Blocos

Cada lugar onde ocorre uma delcaração é um bloco. Go tem esta hierarquia, do mais externo para o mais interno:

1. Universe block: contém todas as declarações predefinidas da linguagem, como tipos, funções e constantes. Por exemplo, o tipo int, a função len e a constante true estão no universe block.
2. Package block: contém todas as declarações de um pacote, como variáveis, funções e tipos definidos pelo usuário. Por exemplo, se você declarar uma variável no nível de pacote, ela está no package block.
3. File block: contém todas as declarações de um arquivo, como funções, tipos e variáveis. Por exemplo, se você declarar uma função em um arquivo, ela está no file block.
4. Function block: contém todas as declarações de uma função, como variáveis locais, parâmetros e tipos definidos pelo usuário. Por exemplo, se você declarar uma variável dentro de uma função, ela está no function block.
5. Statement block: contém todas as declarações de um bloco de código, como if, for, switch e select. Por exemplo, se você declarar uma variável dentro de um bloco if, ela está no statement block.

A regre de acesso: um bloco interno enxerga tudo dos blocos externos, mas um bloco externo não enxerga nada dos blocos internos. Por exemplo, se você declarar uma variável no package block, ela está visível para todas as funções do pacote. Se você declarar uma variável no function block, ela está visível apenas dentro da função. Se você declarar uma variável no statement block, ela está visível apenas dentro do bloco de código.

```go
package main

import "fmt"      // fmt está no file block

var global = 1    // package block

func main() {     // bloco da função
    x := 2
    {             // bloco interno arbitrário
        y := 3
        fmt.Println(global, x, y) // enxerga tudo
    }
    // fmt.Println(y) // ERRO: undefined: y
}
```
### Escopo em if, for e switch
Uma característica interessante do Go é que o escopo de variáveis declaradas em if, for e switch é limitado ao bloco do statement. Por exemplo, se você declarar uma variável dentro de um bloco if, ela não estará visível fora do bloco if. Isso é diferente de outras linguagens, onde variáveis declaradas em if, for e switch podem estar visíveis fora do bloco.

```go
if x := 10; x > 5 { // x está visível apenas dentro do bloco if
    fmt.Println(x) // 10
}
// fmt.Println(x) // ERRO: undefined: x
for i := 0; i < 5; i++ { // i está visível apenas dentro do bloco for
    fmt.Println(i) // 0 1 2 3 4
}
// fmt.Println(i) // ERRO: undefined: i
switch y := 20; y { // y está visível apenas dentro do bloco switch
case 10:
    fmt.Println("y is 10")
case 20:
    fmt.Println("y is 20") // y is 20
}
// fmt.Println(y) // ERRO: undefined: y
```
Isso é ótimo para limitar a vida de variáveis e evitar conflitos de nomes. No entanto, é importante ter cuidado para não declarar variáveis com o mesmo nome em blocos diferentes, pois isso pode levar a shadowing, que é quando uma variável em um bloco interno "sombreia" uma variável com o mesmo nome em um bloco externo.

O mesmo vale para for e switch. Um detalhe de versão a partir do Go 1.22, cada iteração de um for cria uma nova variável de loop. Antes a variável era compartilha entre iterações, o que causava bugs clássicos com closures e goroutines. Agora cada iteração tem sua própria variável de loop, o que evita esses bugs.

### Shadowing

Quando você declara, num bloco interno, uma variável com o mesmo nome de uma variável de um bloco externo, a variável do bloco interno "sombreia" a variável do bloco externo. Isso significa que, dentro do bloco interno, a variável do bloco externo não está mais visível, e qualquer referência à variável com esse nome vai se referir à variável do bloco interno.

```go
package main

import "fmt"

var x = 1 // variável global

func main() {
    x := 2 // variável local, sombreia a variável global
    fmt.Println(x) // 2
    {
        x := 3 // variável local, sombreia a variável local
        fmt.Println(x) // 3
    }
    fmt.Println(x) // 2
}
```
Saida: 1, 2, 3, 2. O x externo não sumiu e nem foi reatribuido, só ficou inacessível dentro do bloco interno. Isso é importante para evitar bugs difíceis de encontrar, e para tornar o código mais legível e compreensível.

### Shadowing com atribuição múltipla
Aqui a regra da seção 1 volta: := só reaproveita variáveis declaradas no bloco atual.
```go
func main() {
    x := 10
    if x > 5 {
        x, y := 5, 20      // x aqui é NOVO, porque o x externo está em outro bloco
        fmt.Println(x, y)  // 5 20
    }
    fmt.Println(x)         // 10
}
```
Muita gente acha que x, y := reaproveitaria o x externo, mas não. O x externo está em outro bloco, então o x aqui é novo. Isso é importante para evitar bugs difíceis de encontrar, e para tornar o código mais legível e compreensível.

### O bug real mais comum
Na prática, isso aparece muito com err e com variáveis que você quer preencher dentro de um if:
```go
func carregar(usarCache bool) (*Config, error) {
    var cfg *Config
    var err error

    if usarCache {
        cfg, err := lerCache() // BUG: cfg e err são novas, sombreiam as externas
        if err != nil {
            return nil, err
        }
        _ = cfg
    }

    return cfg, err // cfg é sempre nil aqui!
}
```
O que acontece aqui é que dentro do if, você está declarando novas variáveis cfg e err, que sombreiam as variáveis cfg e err declaradas no início da função. Então, quando você retorna cfg e err no final da função, você está retornando as variáveis externas, que nunca foram inicializadas, e não as variáveis internas, que foram inicializadas dentro do if. Isso é um bug clássico de shadowing, e pode levar a comportamentos inesperados e difíceis de depurar. A solução é não usar := dentro do if, e sim atribuir os valores às variáveis externas, usando = em vez de :=. Por exemplo:
```go
func carregar(usarCache bool) (*Config, error) {
    var cfg *Config
    var err error

    if usarCache {
        cfg, err = lerCache() // CORRETO: atribui os valores às variáveis externas
        if err != nil {
            return nil, err
        }
        _ = cfg
    }
    return cfg, err // cfg e err agora são as variáveis externas, que foram inicializadas corretamente
}
```

### Sombreando pacotes importados
imports ficam no file block, então qualquer variável declarada no package block ou em blocos internos pode sombrear um import. Isso é importante para evitar conflitos de nomes, e para tornar o código mais legível e compreensível. Por exemplo:
```go
func main() {
    x := 10
    fmt.Println(x)
    fmt := "oops"
    fmt.Println(fmt) // ERRO: fmt.Println undefined (type string has no field or method Println)
}
```

Ferramentas para detectar shadowing: go vet, golangci-lint, ineffassign. O go vet é uma ferramenta que analisa o código e procura por problemas de estilo, bugs e más práticas. Ele pode ser configurado para rodar automaticamente em seu editor de código ou como parte do processo de build. O golangci-lint é uma ferramenta que combina várias ferramentas de linting em uma só, e pode ser configurada para rodar automaticamente em seu editor de código ou como parte do processo de build. O ineffassign é uma ferramenta que detecta variáveis declaradas e não usadas, e pode ser configurada para rodar automaticamente em seu editor de código ou como parte do processo de build.

O revive com a regra redefines-builtin-id pega sombreamento de idenficadores predeclaraods.

```yaml
linters:
    enable:
        - revive
        - predeclard
linters-settings:
    govet:
        check-shadowing: true
        settings:
            shadow:
                strict: true
        enable-all: true
```
