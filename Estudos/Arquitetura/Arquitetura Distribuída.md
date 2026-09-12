# Falasias da arquitetura distribuida

- circut breaks - arquitetura distribuida
- diminiua o stamp coupling em arquitetura distribuida utilizando endpoints da api restful privados, seletores de campo no contrato, graphql para desacolar os contratos, contratos baseados em valores com contrato baseados no consumo CDCs e endpoints de mensageria interna.
- Sobre a questão da segurança de rede em arquitetura distribuidas devemos assegurar cada endpoint.
	|Conceito|Pergunta que responde|
	|---|---|
	|**HTTPS/TLS**|A comunicação está protegida?|
	|**Authentication**|Quem é você?|
	|**JWT**|Como carregamos a identidade/claims?|
	|**Authorization**|Você pode fazer isso?|
	|**RBAC/Policies**|Qual regra determina a permissão?|
	|**Validation**|Os dados enviados são aceitáveis?|
	|**Rate Limiting**|Você está fazendo requisições demais?|
	|**API Gateway**|Onde podemos centralizar parte dessas proteções?|
	|**mTLS/service identity**|Esse serviço realmente é quem diz ser?|

- Splunk - log distribuído
- Transações ACID para assegurar que os dados sejam atualizados de um modo correto para garantir uma alta consistência e integridade deles.
- Consistência Eventual
- Sagas transacionais5

