Esse foi um projeto super simples para se aprofundar de maneira prática tudo que foi estudado sobre Array e seus algoritmos relacionados.

Então vai funcionar da seguinte maneira:

1 - go run main.go:
	Ao roda esse comando, a API vai chamar minha função LoadCities que vai ler o arquivo cities.csv e estruturar os dados usando como base os dados de city que se localiza no dominio. Depois disso é selecionado um índice para iniciar e entramos em um laço for que vai percorrer cities por meio do metodo range do Go, dentro do for, vamos chamar a função distance que é um Haversine que vai calcular a distância entre dois pontos na terra. No final vai ser retornado no terminal, id, nome da cidade e a distância entre a cidade referenciada por ID antes do loop.