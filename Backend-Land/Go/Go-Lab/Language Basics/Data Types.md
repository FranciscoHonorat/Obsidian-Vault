# Language Basics: Data Types em Go

Os tipos que aparecem no seu roadmap são os **tipos predeclarados** (predeclared types), que o livro do Jon Bodner cobre no Capítulo 2. Antes de entrar em cada um, dois conceitos valem para todos.

**Zero value.** Toda variável declarada sem valor recebe automaticamente um valor padrão. Não existe "lixo de memória" como em C, e não existe `null` para tipos básicos.

```go
var i int      // 0
var f float64  // 0
var b bool     // false
var s string   // "" (string vazia)
var r rune     // 0
```

**Literais são "untyped".** Um literal como `10` ou `3.14` não tem tipo fixo até ser usado. Ele tem um *tipo padrão* (default type): `10` vira `int`, `3.14` vira `float64`, `'a'` vira `rune`, `"oi"` vira `string`. Isso vai ser importante na parte de conversão.

---

## 1. Numeric Types

### 1.1 Integers (Signed e Unsigned)

**Signed** (com sinal) aceita negativos. **Unsigned** (sem sinal) aceita só zero e positivos, mas alcança o dobro do valor máximo com o mesmo número de bits.

| Tipo | Tamanho | Faixa |
|---|---|---|
| `int8` | 8 bits | -128 a 127 |
| `int16` | 16 bits | -32.768 a 32.767 |
| `int32` | 32 bits | -2.147.483.648 a 2.147.483.647 |
| `int64` | 64 bits | cerca de ±9,22 × 10¹⁸ |
| `uint8` | 8 bits | 0 a 255 |
| `uint16` | 16 bits | 0 a 65.535 |
| `uint32` | 32 bits | 0 a 4.294.967.295 |
| `uint64` | 64 bits | 0 a cerca de 1,84 × 10¹⁹ |

**Os nomes especiais:**

- `byte` é um alias de `uint8`. São literalmente o mesmo tipo, mas em código idiomático você escreve `byte`.
- `int` tem 32 bits em CPUs de 32 bits e 64 bits na maioria das CPUs de 64 bits. Por não ter tamanho fixo, `int` e `int64` **são tipos diferentes** e não se misturam sem conversão, mesmo que na sua máquina ambos tenham 64 bits.
- `uint` segue a mesma regra do `int`, só que sem sinal.
- `rune` é alias de `int32` (veremos adiante).
- `uintptr` é usado com o pacote `unsafe`; você praticamente não vai tocar nele.

**Qual usar?** A regra do livro é simples: use o tamanho específico quando um formato binário ou protocolo de rede exigir; use generics quando escrever uma função que deve aceitar qualquer inteiro; **em todos os outros casos, use `int`**.

**Literais inteiros** aceitam bases diferentes e underscore para legibilidade:

```go
decimal := 1_000_000
binario := 0b1010      // 10
octal   := 0o17        // 15
hexa    := 0xFF        // 255
```

Evite o octal no formato antigo (`017`), porque confunde com decimal.

**Operadores:** `+ - * / %`, as formas compostas (`+=`, `*=` etc.), comparações (`== != < <= > >=`) e operadores de bits (`<< >> & | ^ &^`).

Dois detalhes importantes na divisão:

```go
fmt.Println(7 / 2)   // 3   (divisão inteira, descarta a parte fracionária)
fmt.Println(-7 / 2)  // -3  (trunca em direção ao zero, não arredonda para baixo)
fmt.Println(-7 % 2)  // -1  (o resto herda o sinal do dividendo)
```

Dividir um inteiro por zero em tempo de execução causa **panic**.

**Overflow:** com variáveis, o valor "dá a volta" silenciosamente. Com constantes, o compilador recusa.

```go
var b byte = 255
b++
fmt.Println(b) // 0

var x int8 = 200 // erro de compilação: constant 200 overflows int8
```

### 1.2 Floating Points

Go tem dois: `float32` (cerca de 6 a 7 dígitos de precisão) e `float64` (cerca de 15 a 16 dígitos). Ambos seguem o padrão **IEEE 754**, o mesmo do `double` e `float` do C#.

**Regra prática:** use `float64`. É o tipo padrão dos literais com ponto decimal e reduz problemas de precisão. Só troque para `float32` se um formato externo exigir ou se o profiler mostrar que memória é um gargalo real.

**Literais:**

```go
a := 3.14
b := 6.03e23      // notação científica
c := 0x12.34p5    // hexadecimal com expoente p (= 582.5)
```

**O ponto crítico: floats são aproximações.** Eles não conseguem representar exatamente a maioria dos valores decimais:

```go
x, y := 0.1, 0.2
fmt.Println(x+y == 0.3) // false
fmt.Println(x + y)      // 0.30000000000000004
```

Por isso:

1. **Nunca use float para dinheiro** ou qualquer valor que exija exatidão decimal. Use inteiros (centavos) ou uma biblioteca decimal (o livro mostra uma no capítulo sobre módulos de terceiros). Isso vale para créditos de carbono, preços, saldos.
2. **Não compare floats com `==`**. Use uma tolerância (epsilon):

```go
const epsilon = 1e-9
if math.Abs((x+y)-0.3) < epsilon {
    fmt.Println("iguais o suficiente")
}
```

**Comportamentos especiais:**

- `%` não funciona com floats (use `math.Mod`).
- Dividir um float diferente de zero por zero dá `+Inf` ou `-Inf`, não panic.
- Dividir zero por zero dá `NaN` (Not a Number), e `NaN` não é igual nem a si mesmo.

```go
var zero float64
fmt.Println(1 / zero)            // +Inf
fmt.Println(zero / zero)         // NaN
nan := zero / zero
fmt.Println(nan == nan)          // false
fmt.Println(math.IsNaN(nan))     // true
```

(Se escrever `1.0 / 0.0` direto como constante, o compilador acusa erro; por isso uso uma variável.)

Para geoprocessamento isso é bem relevante: coordenadas são `float64`, e comparar dois pontos com `==` é um bug clássico.

### 1.3 Complex Numbers

Go tem suporte nativo a números complexos, algo raro. São dois tipos:

- `complex64`: parte real e imaginária em `float32`
- `complex128`: parte real e imaginária em `float64` (o padrão)

```go
x := complex(2.5, 3.1)   // complex128
y := 4 + 2i              // literal imaginário: sufixo i

fmt.Println(x + y)       // (6.5+5.1i)
fmt.Println(real(x))     // 2.5
fmt.Println(imag(x))     // 3.1
fmt.Println(cmplx.Abs(x)) // módulo, do pacote math/cmplx
```

Esse literal com `i` é o "quinto tipo de literal" que o livro menciona. O zero value é `0+0i`. Valem as mesmas limitações de precisão dos floats.

Na prática: você quase nunca vai usar. Existem porque o Ken Thompson achou interessante incluir. Para computação numérica séria em Go existe o pacote Gonum, mas o próprio autor recomenda considerar outras linguagens primeiro.

---

## 2. Boolean

O tipo `bool` só tem dois valores: `true` e `false`. Zero value: `false`.

Operadores: `&&` (E), `||` (OU), `!` (NÃO). Os dois primeiros fazem **short-circuit**: se o lado esquerdo já decide o resultado, o direito nem é avaliado.

```go
if usuario != nil && usuario.Ativo { ... } // seguro: não acessa Ativo se usuario for nil
```

**O ponto que diferencia Go:** não existe "truthiness". Em JavaScript ou Python, `0`, `""` e `null` são tratados como falso. Em Go, **nenhum tipo pode ser convertido para `bool`**, nem implícita nem explicitamente.

```go
n := 5
if n { }          // erro de compilação
if bool(n) { }    // erro de compilação também
if n != 0 { }     // correto

s := ""
if s == "" { }    // forma idiomática de checar string vazia
```

Você sempre usa um operador de comparação para produzir um `bool`.

---

## 3. Runes

Uma `rune` representa **um code point Unicode**, ou seja, um "caractere" no sentido do Unicode. Ela é alias de `int32`: internamente é só um número.

```go
var letra rune = 'J'
fmt.Println(letra)          // 74 (o número do code point)
fmt.Printf("%c\n", letra)   // J
fmt.Printf("%U\n", letra)   // U+004A
```

**Literais de rune** usam aspas simples e aceitam escapes:

```go
a := 'a'
b := '\n'        // quebra de linha
c := '\u00e7'    // ç (Unicode 16 bits)
d := '\U0001F600' // emoji (Unicode 32 bits)
e := '\x41'      // A (hexadecimal, 8 bits)
```

**Idiomático:** mesmo sendo o mesmo tipo para o compilador, use `rune` quando a intenção for um caractere, e `int32` quando for um número. O tipo comunica a intenção.

```go
var inicial rune = 'F'   // bom
var inicial2 int32 = 'H' // compila, mas confunde quem lê
```

Comparando com C#: o `char` do C# é UTF-16 (16 bits), então emojis ocupam dois `char`. A `rune` do Go tem 32 bits e cabe qualquer code point sozinha.

---

## 4. Strings

Uma `string` em Go é uma **sequência imutável de bytes**, que por convenção contém texto em UTF-8. Zero value: `""`.

Operações básicas:

```go
s := "Olá" + ", " + "mundo"   // concatenação com +
fmt.Println(s == "Olá, mundo") // comparação com ==, !=, <, > (ordem lexicográfica por bytes)
```

**Imutabilidade:** você pode reatribuir a variável, mas não alterar o conteúdo.

```go
s := "gato"
s = "rato"   // ok, a variável aponta para outra string
s[0] = 'p'   // erro de compilação
```

**O detalhe mais importante: bytes versus caracteres.** Como caracteres fora do ASCII ocupam mais de um byte em UTF-8, `len` conta bytes, não letras. Isso afeta diretamente textos em português:

```go
palavra := "ação"
fmt.Println(len(palavra))                    // 6 (ç e ã ocupam 2 bytes cada)
fmt.Println(utf8.RuneCountInString(palavra)) // 4
fmt.Println(palavra[1])                      // 195 (um byte, não o "ç")
fmt.Println(len([]rune(palavra)))            // 4
```

Para percorrer caractere por caractere, use `for range`, que decodifica UTF-8 e entrega runes:

```go
for i, r := range "ação" {
    fmt.Printf("byte %d: %c\n", i, r)
}
// byte 0: a
// byte 1: ç
// byte 3: ã
// byte 5: o
```

Repare que o índice pula de 1 para 3: ele é a posição em bytes. O livro aprofunda isso no Capítulo 3.

Existem duas formas de escrever literais de string:

### 4.1 Interpreted String Literals

Usam **aspas duplas**. As sequências de escape são interpretadas:

```go
s := "Linha 1\nLinha 2\tcom tab\t\"aspas\" e barra \\ e \u00e7"
```

Escapes comuns: `\n` (nova linha), `\t` (tab), `\\` (barra invertida), `\"` (aspas), `\u` e `\U` (Unicode), `\x` (byte hexadecimal).

Restrições: não podem conter quebra de linha literal no código nem aspas duplas ou barra invertida sem escape.

### 4.2 Raw String Literals

Usam **crases** (backticks). Nada é interpretado: o que você digita é exatamente o que fica na string, incluindo quebras de linha.

```go
caminho := `C:\Users\francisco\docs`   // sem precisar escapar as barras
regex   := `^\d{5}-\d{3}$`             // CEP, sem dobrar as barras
query   := `
    SELECT id, nome
    FROM especies
    WHERE bioma = $1
`
json := `{"lat": -6.76, "lon": -38.23}`
```

A única coisa que não pode aparecer dentro é a própria crase.

**Quando usar cada uma:** interpretada para strings do dia a dia; raw para regex, SQL, JSON, templates, caminhos do Windows e qualquer texto com muitas barras ou aspas. Você também vê raw strings em struct tags (`json:"nome"`), que aparecem mais adiante no roadmap.

---

## 5. Type Conversion

Aqui Go é bem diferente do C#. No C#, `int` vira `double` automaticamente (conversão implícita). **Em Go não existe promoção automática entre variáveis de tipos diferentes.** Toda conversão é explícita, com a sintaxe `Tipo(valor)`.

```go
var x int = 10
var y float64 = 30.2

soma1 := float64(x) + y  // 40.2
soma2 := x + int(y)      // 40
soma3 := x + y           // erro: mismatched types int and float64
```

Vale até para inteiros de tamanhos diferentes:

```go
var a int = 10
var b byte = 100
fmt.Println(a + int(b))   // 110
fmt.Println(byte(a) + b)  // 110 (como byte)
```

A justificativa do Go: regras de conversão automática são complexas e geram bugs sutis. Sendo explícito, quem lê sabe exatamente o que acontece.

### Comportamentos a memorizar

**Float para inteiro trunca em direção ao zero** (não arredonda):

```go
fmt.Println(int(3.9))   // 3   (com variável; veja abaixo)
f := -3.9
fmt.Println(int(f))     // -3
fmt.Println(int(math.Round(f))) // -4, se quiser arredondar
```

Atenção: `int(3.9)` com a constante direto dá erro de compilação ("truncated"), porque o compilador percebe a perda. Com variável, ele trunca.

**Inteiro maior para menor corta os bits excedentes:**

```go
v := 300
fmt.Println(int8(v))  // 44 (300 - 256)
```

**Constantes e literais untyped são a exceção.** Como ainda não têm tipo fixo, podem ser usados diretamente onde couberem:

```go
const valor = 10
var i int = valor       // ok
var f float64 = valor   // ok, sem conversão

var g float64 = 5       // o literal 5 vira float64 sem problema
```

**Cuidado com inteiro para string.** `string(65)` não produz `"65"`, produz `"A"` (interpreta o número como code point). O `go vet` inclusive avisa sobre isso. Para converter números em texto e vice-versa, use o pacote `strconv`:

```go
s := strconv.Itoa(65)          // "65"
n, err := strconv.Atoi("123")  // 123, nil
if err != nil {
    // "abc" daria erro aqui, então sempre trate
}
f, err := strconv.ParseFloat("-6.76", 64)
```

**Conversões entre string, bytes e runes** são permitidas e comuns:

```go
s := "ação"
bs := []byte(s)   // [97 195 167 195 163 111]
rs := []rune(s)   // [97 231 227 111]
fmt.Println(string(bs), string(rs)) // ação ação
```

**E lembrando:** nada converte para `bool`. Use comparações.

---

## Para fixar

O livro propõe três exercícios no fim do Capítulo 2 que cobrem exatamente esse conteúdo. Recomendo tentar antes de olhar a solução:

1. Declare um inteiro `i` com valor 20, atribua a um float `f` e imprima os dois. (Dica: o que o compilador exige?)
2. Declare uma constante `value` que possa ser atribuída tanto a um inteiro quanto a um float, sem conversão.
3. Crie `b` do tipo `byte`, `smallI` do tipo `int32` e `bigI` do tipo `uint64`, atribua o valor máximo de cada um, some 1 e imprima. Tente prever o resultado antes de rodar.

Se quiser, na próxima etapa posso revisar suas soluções ou montar flashcards Anki com os pontos mais traiçoeiros desse tópico (divisão inteira, `len` em bytes, truncamento, `string(65)`).