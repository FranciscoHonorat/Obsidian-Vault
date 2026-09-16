O escopo determina a visibilidade e a vida útil das variáveis e funções dentro de um programa. Em linguagens de programação, o escopo pode ser global ou local. Variáveis globais são acessíveis em qualquer parte do código, enquanto variáveis locais só podem ser acessadas dentro do bloco de código onde foram declaradas. Compreender o escopo é essencial para evitar conflitos de nomes e garantir que as variáveis sejam usadas corretamente dentro do programa.

- Escopo de Pacote (Global no arquivo): Declarar variáveis e funções fora de qualquer função ou bloco de código cria um escopo global dentro do arquivo, permitindo que essas variáveis e funções sejam acessadas em qualquer parte do arquivo. No entanto, é importante ter cuidado ao usar variáveis globais, pois elas podem levar a efeitos colaterais indesejados se não forem gerenciadas corretamente. Outro detalhe é que as variáveis globais em Go são visíveis apenas dentro do pacote em que foram declaradas, a menos que sejam exportadas (começando com letra maiúscula), permitindo que outros pacotes as acessem.

- Escopo de Função: Variáveis declaradas dentro de uma função têm escopo local, o que significa que só podem ser acessadas dentro dessa função. Isso ajuda a evitar conflitos de nomes e mantém o código mais organizado. Em Go, as variáveis locais são criadas quando a função é chamada e destruídas quando a função termina sua execução.

- Sombreamento (Shadowing): O sombreamento ocorre quando uma variável local tem o mesmo nome de uma variável global ou de outro escopo mais amplo. Nesse caso, a variável local "sombreia" a variável global, tornando-a inacessível dentro do escopo da função. Isso pode levar a confusão e erros se não for gerenciado corretamente, portanto, é importante escolher nomes de variáveis distintos para evitar sombreamento indesejado.


Como isso funciona em Go? Em Go, o escopo é determinado pelo local onde as variáveis e funções são declaradas. Variáveis globais são declaradas fora de qualquer função e são visíveis em todo o pacote, enquanto variáveis locais são declaradas dentro de funções e só podem ser acessadas dentro dessas funções. O sombreamento também é possível em Go, e os programadores devem estar atentos a isso para evitar confusões no código.

Isso se aplica a funções também. Funções declaradas dentro de um pacote têm escopo global dentro desse pacote, enquanto funções declaradas dentro de outras funções (funções aninhadas) têm escopo local e só podem ser chamadas dentro da função que as contém.

É interessante pensar nisso, pois se tentamos fazer um "package x_test" e chamamos uma função que está dentro do pacote "x", não conseguiremos acessar a função, pois ela não é exportada. Isso significa que ela não é visível fora do pacote "x". Para tornar uma função visível fora do pacote, devemos exportá-la, o que é feito iniciando o nome da função com uma letra maiúscula.


