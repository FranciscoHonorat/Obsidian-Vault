https://go.dev/blog/pipelines

Concorrencia primitiva em Go é muito fácil criar construções para transmissão de dados por pipelines criando uma eficiencia no uso de I/O e multiplos CPUs.  E esse artigo vai apresentar exemplos e maneiras de suavizar os casos de operações que cai e introduzir tecnicas para tratamento de falhas com mais clareza.

O que é uma pipelin?

O artigo apresenta duas definições:

Formal - Apenas uma das muitas formas de concorrencia em programas.
Informal - uma pipeline é uma serie estagios conectados por canais, onde cada estagio é um grupo de goroutines correndo na mesma função. Em cada estagio, a goroutine recebe os valores por upstream via inbound channeles, executa alguma função com base nos dados e produz novos valores, envia agora os novos valores pelo downstream via outbound channels.

E cada estagio tem um determinado numero de inbound e outbound channels, exceto o primeiro ou ultimo estágio, pois eles podem ter apenas o inbound ou outbound respectivamente. Geralmente o primeiro estagio possui o nome de source ou producer e o ultimo estagio de sink ou consumer.

Squaring numbers é o primeiro exemplo de pipelines que é citado para explicar ideias e tecnicas.

```
go

func gen(num ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
```

O que a função gen ta fazendo? simples ela vai converter a lista de inteiros para channel e emitir o inteiro em uma lista. Essa função inicia com um goroutine que está enviando o inteiro para canal e fechando assim que os dados são enviados.

```
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
```

No segundo estagio a função sq vai receber o inteiros pelo channel e retorna ao channal emitindo cada inteiro de maneira quadrada. Após o inbound channel está fechado e esse estagio enviar todos o valores pelo downstream, ele é fechado pelo outbound channel.

```
func main() {
	c := gen(2, 3)
	out := sq(c)
	
	fmt.Println(<-out)
	fmt.Println(<-out)
}
```
A função main vai definir a pipeline e começa o estagio final e então os dados vão ser recebidos do segundo estagio e vão ser printados um por um na tela, assim que o canal é fechado.

```
func main() {
	for n := range sq(sq(gen(2, 3))) {
		fmt.Prinln(n
	}
}
```

Como o inbound e outbound possui os mesmo tipos a gente pode refatorar colocando tudo dentro de um range loop que vai percorrer diversos numeros.

Fan-out, Fan-in

