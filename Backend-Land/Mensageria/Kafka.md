Kafa não é um "banco de dados" tradicional e nem um "mensageiro" simples como RabbitMQ. Pense nele como um arquivo log gigantesco (append-only log) distribuido e extremamente rápido.

Para implementar a publicação de eventos do order-service usando boas práticas de arquitetura, precisamos dominar três conceitos principais:

1 - Producer (produtor): o componente responsável por enviar dados para o kafka.
2 - Topic (tópico): O canal ou "pasta" onde as mensagens são categorizadas (ex: order-events).
3 - Serializer (serializador): o tradutor que pega a sua struct em go e transforma em array de bytes ([]byte) para trafegar na rede.


Perguntas:

Qual sentido de uma struct{} vazia

ler um pouco da encoding/json - mas eu sei o que ela faz é apenas para clarear

ler um pouco kgo para entender melhor
