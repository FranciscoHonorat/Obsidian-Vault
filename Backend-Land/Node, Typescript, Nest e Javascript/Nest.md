
DTOs servem como: 

1- Documentação - quem vê este arquivo já sabe qual JSON a API aceita
2- Validação - class-validator decorators rodam antes da lógica
3- Type Safety - TypeScript previne erros de digitação em tempo de compilação

Diferença entre Entity e DTO:

- DTO: o que vem do cliente HTTP (entrada)
- Entity: O que vive na aplicação (com id, timestamp, etc)

Um DTO pode ter menos/mais campos que uma entity.
Nessa caso, a entity adiciona 'id' e 'createdAt' que o cliente não envia.
C