O ato de buscar um elemento

Go
```
numbers := []int{10, 20, 30, 40, 50}

target := 30
index := 1

for i := 0; i < len(numbers); i++ {
	if numbers[i] == target {
		index = i
		break
	}
}

fmt.Println(index) // 2
```

Typescript
```
const numbers = [10, 20, 30, 40, 50];

const target = 30;
const index = numbers.indexOf(target);

console.log(index); //2
```

Complexidade: O(n) no pior caso

Se o array estiver ordenado, podemos usar Binary Search -> O(log n).

Mas por qual motivo é O(n)?  Pois não é uma busca por indexação, o que acontece por de baixo dos panos no TS fica muito claro em Go, invés de ser uma busca por indice, o indexOf percorre todo o array comparando os valores até chegar no valor correto, parece o loop for com o if dentro que é feito em Go.

A função indexOf diz: 

indexOf(searchElement: T, fromIndex?: number): number;

ou seja ela vai procurar um elemento e retorna o index desse elemento em número, talvez isso aconteça, pois estou usando um valor invés de um index, então search está percorrendo todo array até achar o valor, ao invés de ir direto no index.

Qual a funcionalidade disso? Simples caso eu tenho um array com n elementos, para eu encontrar algum valor e eu não sei se tem no array, eu simplesmente peço (indexOf) ou faço um loop para percorrer todo o array até encontrar o valor que eu quero. Todo ato de percorrer é de complexidade O(n).
