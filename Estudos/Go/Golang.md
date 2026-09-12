Primeiro teste - Explique sem olhar

Primeira tentativa

O que é Golang?

Golange é uma linguagem compilada inspirada em C, possui caracterisiticas que são voltadas a ser uma linguagem de uma syntax simples e fácil, rápida e que lida com alta concorrência e baixa latência 

Por que Golang existe?

Golange foi criada para ser uma linguagem que lida com alta latencia de requisições e alta concorrência simultaneas, por meio do seu tratamento de contexto, GC e GMP se tornou uma linguagem muito usada por fintechs.

Como Golang funciona internamente?

Internamente Golange funciona por meio um sistema de GMP que é goroutine, machine e processor e um GC para lidar com a fulga de dados que sai da stack para heap 

Quando eu usaria Golang?

Utiliza-se Golang quando queremos gera um binario leve e também quando temos que lidar com uma alta concorrencia 

Quando eu não usaria Golang?

Geralmente em sistemas de analise da dados e sistemas onde não estamos lidando com altas concorrencias

Quais são os trade-offs?

Se estou lidando com sistemas que tem bastante concorrencia e também vai ter bastante processos paralelos com alto risco de race, podemos usar Golang pois a linguagem tem uma alta capacidade para lidar com essas situações


Revisão do conteúdo:

Resultado foi 2,5 de 5

> **1. Qual é a diferença entre concorrência e paralelismo?**

> **2. O que acontece quando uma variável escapa da stack para o heap? Quem decide isso e quem posteriormente gerencia essa memória?**

> **3. Explique G, M e P e descreva o caminho de uma goroutine até chegar a uma thread do sistema operacional.**


O que é concorrência?

é um conceito da programação para se referenciar a capacidade de um sistema executar multiplas tarefas aparementemente ao mesmo tempo, são dividas em dois tipos: concorrência real (Paralelismo) que é multiplas tarefas executando literalmente ao mesmo tempo em multiplos processos/cores e a outra é a concorrência intercalada (time-slicing) que é um único processo alternando rápidamente entre as tarefas para parecer simultaneo, mas na verdade é sequencial

O que é paralelismo?

É possível ter concorrência em uma única CPU?

É possível ter paralelismo sem goroutines?

Onde Go entra nisso?

**Meta:** Explicar isso sem decorar definição.

Concorrência é uma propriedade do código e Paralelismo é um propriedade do programa em execução.

Uma maquina com um único núcleo vai executar as coisas de maneira intercalada ou sequencial ou sejas concorrencias podem ser executadas de maneira intercalada e uma maquina com multiplos núcleos podem executar as coisas de maneira paralela.


O que é um Goroutine?

Qual é a diferença entre Goroutine e uma thread?

Quem gerencia as goroutines?

Quando você escreve:
```
go processoPayment()
```

o que acontece internamente?

E se eu criar:
```
for i := 0; i < 1_000_000; i++ {
    go process(i)
}
```
isso significa que terei 1 milhão de threads executando simultaneamente?

-------
Goroutine é um unidade de execução concorrente gerenciada pelo runtime de Go, um dos pontos mais importantes é que uma goroutine não é uma thread do sistema operacional e quando esccrevemos:
```
go processPayment()
```
estamos dizendo ao runtime do Go: "execute processPayment()" como uma atividade concorrente." 
O runtime do Go então fica responsavel por agendar essa goroutine para execução. 

O runtime do Go multiplexa muitas goroutines sobre threads do sistema operacional. O scheduler utiliza o modelo GMP, no qual G representa a Goroutine, M uma thread do sistema operacional e P um contexto em execução. Os P possuem trabalho para executar e o scheduler utilizar o work stealing para redistribuir trabalho quando necessário. Portanto, milhares de goroutines não significa milhares de threads.

Em Go, o algoritmo trabalha com o continuantions, e não simplesmente com uma fila de goroutines FIFO.

Uma representação simplicada seria:

              Go Runtime
                  │
        ┌─────────┴─────────┐
        ↓                   ↓
       P1                   P2
        │                   │
    work deque          work deque
        │                   │
     G1 G2 G3            G4 G5 G6
Se P1 ficar sem trabalho, ela pode roubar trabalho de outro contexto

P1                         P2
│                          │
│ sem trabalho             │ G4 G5 G6
│                          │
└──────── steal ──────────→│
                           │
                         G5 G6


e scheduler de Go utilizar uma estrategia de work stealing para multiplexar goroutines sobre threads  do sistema operacional.

o modelo do runtime é:

G = Goroutines
M = I/Os threads
P = contexto/Processor

no runtime:

M -> P -> G

Ou seja, as threads (M) hospedam P, e os P fazem o scheduler das G.

Agora vamos construir um modelo mental

Imagine 4 CPUs, então simplicando seria

P1 -> CPU1
P2 -> CPU2
P3 -> CPU3
P4 -> CPU4

e temos 100 goroutines, então não precisamos ter 100 threads para lidar com elas, pois essas goroutines podem ser escalonadas pelos contextos disponíveis

Go Runtime

P1 → G1 G2 G3 ...
P2 → G4 G5 G6 ...
P3 → G7 G8 G9 ...
P4 → G10 G11 ...

Se um P ficar sem trabalho, entra o mecanismo de work stealing.

E ainda existe uma situação importante:

Imagine que 
P1 -> G1

e G1 bloqueia esperando uma operação I/O. O runtime pode dissociar o contexto do thread bloqueado e permitir que outra thread execute goroutines usando aquele contexto, mantendo os CPUs ocupados.


Tenho varias goroutines e um numero limitado de contextos de execução (P). O scheduler do runtime distribui essas goroutines para execução sobre threads (M), que utilizam as CPUs disponíveis.

G = goroutines que é uma unidade de execução concorrente gerenciada pelo runtime do Go.
P = Processor, um contexto de execução do runtime do Go
M = Machine, uma thread do sistema operacional utilizada pelo runtime.
CPUs = recursos lógicos disponíveis para processamento.

O runtime possui um modelo de escalonamento e distribuição de trabalho, incluindo o work stealing

O P representa a capacidade de execução disponibilizada pelo runtime para as goroutines. Ele limita quantas G podem executar código Go simultaneamente, evitando que a quantidade de goroutines seja confundidas com a quantidade de execução simultanea.

M representa uma thread do sistema operacional usado pelo runtime do Go para executar goroutines.

Uma maquina pode ter varias threads de execução do sistema operacional, e a quantidade de M não representa diretamente a quantidade de CPUs/cores.

O runtime utiliza M (threads do SO) para executar Go, enquanto P fornece o contexto necessário para que um M execute o código Go.

P determina quantos contextos podem executar código Go simultaneamente. M são threads do So que executam o trabalho.

Thread (linha de execução) é uma unidade de execução gerenciada pelo sistema operacional dentro de um processo.

O scheduler organiza quais goroutines serão executadas e quando elas terão oportunidade de execução.

Work stealing ajuda a reduzir a ociosidade dos P, permitindo que um P sem trabalho procure trabalho disponivel associado a outra P.

Quando uma G bloqueia uma operação que pode bloquear a execução, o runtime pode descoplar a M que ficou bloqueada do P, permitindo que P seja associado a outra M e continue executando outras G.

o runitme desacopla a M bloqueada do P, permitindo que outra M execute o trabalho usando aquele P.


A questão é: **como poderíamos representar o estado "variável não existe" usando as ferramentas do** `**testing.T**`******?** acredito que não exista nenhuma, no os possui uma variável chama FOO que tem essa responsabilidade



Por que criamos `lookupNonEmptyEnv()` durante a refatoração, mas não criamos uma interface `Environment`?

Pois, não queriamos criar abstração desnecessária, mas não sei dizer de maneira aprofunda o motivo

Qual foi a diferença entre a primeira alteração que fizemos (adicionar os testes de whitespace/missing) e a segunda alteração (criar `lookupNonEmptyEnv`)?

os códigos tinham uma duplicação sendo usada, ao criar essa variavel auxiliar colocamos uma abstração tornando as funções de comportamentos mais limpas e sendo responsáveis apenas por um determinado comportamento