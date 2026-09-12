Passei as últimas semanas aprofundando arquitetura de microsserviços em Go  e documentando cada decisão, erro e correção pelo caminho.

O re-api-books começou como uma API de filmes simples, mas virou um projeto de estudo sério sobre o que realmente separa "funciona no meu Docker Compose" de "está pronto pra produção":

🔹 Arquitetura assíncrona de verdade  api-gateway (Gin) publica criações numa fila RabbitMQ e devolve 202 na hora; movies-service (gRPC + MongoDB) consome, processa e grava o status do job. Cliente consulta o resultado depois.

🔹 Teste de carga com números reais, não estimados,  rodei vegeta contra os endpoints e encontrei um índice ausente no MongoDB que colapsava a listagem de 1,45ms pra 1,22s de latência média sob 80 req/s. Corrigir o índice trouxe a mesma taxa de volta pra 1,45ms sem tocar em hardware.

🔹 Tracing distribuído com Jaeger para cada request gera uma trace OpenTelemetry que atravessa gateway → serviço → MongoDB, incluindo o caminho assíncrono via RabbitMQ (propaguei o contexto de trace nos próprios headers AMQP).

🔹 Resiliência validada contra falha real com circuit breaker + retry no gateway. Testei derrubando os containers de verdade, não com mocks, e isso revelou dois bugs reais: uma chamada gRPC sem timeout que travava por 20s, e um publisher que nunca reco

O que mais gostei de fazer diferente: toda decisão não-óbvia, o que foi te por quê e está documentada em ADRs no próprio repo, não só no código. Acheique isso mudou completamente como eu penso sobre esse tipo de trabalho.

Projeto de estudo, 100% open source, zero custo de infra (roda local com Docker/kind/LocalStack):
🔗 github.com/FranciscoHonorat/re-api-books

#golang #microservices #softwareengineering #grpc #observability
