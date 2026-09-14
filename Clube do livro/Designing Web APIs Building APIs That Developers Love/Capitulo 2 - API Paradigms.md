Capitulo 2 - API Paradigms
Nessa introdução de capitulo, temos uma breve discussão como em muitos momentos acabamos por meio de diversos motivos empresarias por assim dizer construindo uma api sem ter todos os trade-offs e essa capitulo, vai ajudar a gente a pelo menos ter um norte para escolher o melhor paradigma o que vai facilitar futuras alterações no API.

request-response API:

é a mais comum para serviços web,  onde o client faz a requisição para algum endpoint e é respondido com dados em json ou xml pelo servidor. Existe alguns paradigmas comuns usado para expor esse tipo de serviço request-response API: Rest, RPC e GraphQL

(Dúvida, é possível usar Rest e GraphQL juntos ou Rest e RPC, acredito que sim)

Representational State Transfer (REST)

REST a escolha mais popular para a criação de API web, muito usada por diversas empresas. 

Em REST tudo é sobre resources. E o que seria? podemos dizer que é um entidade que pode ser identificada por nome, endereço ou qualquer coisa na web. REST APIs geralmente deixa exposto esse resources para serem requisitados e retornados pelos metodos HTTP (create, read, update e delete) conhecidos como CRUD

Que geralmente seguem uma regra básica:

Resources tem que ser parte da URLs tipo /users
Para da cada resoursces deve ter pelo menos duas implementações que são uma para coleção e uma para um elemento especifico (tipo getByID e getList)
Não usar verbos
Cada metodo HTTP possui uma função, diferentes requisições podem serem feitas na mesma URL com diferentes funcionalidades.

Create - Use POST for creating new resources
Read - Use Get para leitura, Get nunca vai mudar um estado.
Update - Se usa PUT para recolocar um novo recurso e patch para atualizar recursos já existentes
Delete - né

Cada código HTTP possui um sentido: 2xx indica sucesso, 3xx indica que o serviço mudou, 4xx problema do lado do cliente e 5xx problema indica do lado do servidor 

REST APIs pode retorna as requisições em JSON ou XML 

Fora as operações CRUD, o autor também cita que em alguns momentos temos que fazer algumas requisições foram do CRUD ou Non-CRUD e comumente nesses caso

Ele cita o que na API do Github, ele usam "archived" que é um paramentro de entrada para editar um repositorio representado pelo arquivo do respositorio ativo.

Uma ação como um subsource. Que a API do Github usa para trancar ou destrancar uma issue

Algumas operações como as de pesquisa são mais dificeis no paradigma do REST, tipicamente utilizamos um caso ativo de verbo na URL da API para achar arquivos no Github e que corresponde a uma query.

Remote Procedure Call (RPC)

difernete do REST que funciona por resources, RPC funciona por ações

RPC API geralmente se segue duas regras básicas:

-> Todos os endpoints tem que ter o nome daquela operação para ela ser executada
-> API vai chama o método HTTP mais adequado para executar aquele chamado

o estilo de trabalho do RPC é muito bom para API que vão expor as variáveis a ações que pode ser ter mais nuancias ou complicações do que simples CRUD. Por isso APIs no estilo RPC também acomodam modelos de recursos complexos ou ações que envolvem múltiplos tipos de recursos.

O que isso significa? 

Que dizer que esse paradigma é utilizado para fazer tarefas muito mais complexas do que um simples CRUD, como arquivar uma conversa ou abrir novas conversas no exemplo do Slack.

Bem, como dito pelo autor, RPC APIs não usa exclusivamente o HTTP, ele também pode usar outros protocolos de alta perfomance, como o Apache Thrift e gRPC. No caso do gRPC também pode retorna um JSON. Porém Thrift  e gRPC  tem suas requisições serializadas, possuem estruturada de dados e é mais claro em definir interfaces que permitem serialização. Eles também são ferramentas para edição de estruturas de dados.

Posso criar projetos com o Thrift, gRPC e o próprio HTTP

GraphQL  é um linguagem de query ou seja todas as suas requisições são get e post que vão retorna somente o solicitado pelo cliente e preencher funciona da mesma maneira. Achei interessante olhar o tanto de vantagens que ela possui sobre REST e RPC. Porém é legal também perceber o quanto é dificil otimizar o backend de sistemas dom GraphQL o que pode comprometer a perfomance, além disso, assim como cada aplicação é única, não devemos sair a torto e a direita implementando sem antes entender o domínio do negócio, é um paradigma dificil de trabalhar, pois é preciso tomar cuidado nas querys de requisições e como os dtos vão retorna elas. 

Mas por que é dificil de otimizar? 

Bem, isso se da ao usuarios externos, pois para os usuarios internos essa otimização pode ser testada e debugada, porém para os externos não conseguimos ter esse controle. Ou seja ao tercerizar o GraphQL deixamos de gerenciar multiplas requisições para gerenciar queries mais complexas, o que pode impactar a performance e a infrastructure de varias maneiras.

Event-Drive APIs

Ou seja, eles tinha um problema com a frequente mudança de dados e a resposta rápida a mudanças de estado. Então fazendo com que os desenvolvedores tivessem que consultar a API com muita frequencia.

No geral essa frequente chamada da APi raramente retornava um novo Data o que virou estudo de caso do Zapier que encontrou que apenas 1.5% de todas as chamadas retorna uma novo Data.

E para compartilhar isso existe três mecanismo comum: WebHooks, WebSockets e HTTP Streaming.

O que é WebHooks?

WebHooks é um mecanismo de receber atualizações em tempo real, sua URL geralmente aceita HTTP POST, a API fornece uma implementação  WebHooks para configurar  o envio um e receber um post toda vez que algum evento acontecer. A maior diferença de req-res apis é que WebHooks podemos receber atualizações em tempo real. Geralmente para a gente ficar recebendo atualizações tempos que fica trackeando o channel o tempo inteiro, porém com WebHooks podemos ser notificados a qualquer momento sobre o novo canal, sem precisar se manter trackeando um canal como acontece geralmente onde precisamos ficar o tempo inteiro fazendo requisição na API sobre o novo canal.

Quais a complexidades em termos um suporte WebHooks?

Temos algumas complexidades como falhas e diversas tentativas que funciona da seguinte maneira para assegurar que a entrega no WebHook foi bem sucessidade devemos criar um sistema que deve ficar tentando entregar o WebHook em caso de falhas. Outra complexidade é a segurança que ainda está evoluindo, em WebHooks a maior responsabilidade é garantir que o que foi entregue é legitimo, Firewalls geralmente em aplicativos que estão rodando é comum o firewall acessar a API para verificação, porém como estamos falando de recebimento em trafego, infelizmente não é possível, o que torna dificulta muito a vida de aplicações que usam WebHooks e frequentemente inviaveis, ruido é porque tipicamente o WebHooks representa apenas um evento e o problema começa quando acontece muitplos eventos em um curto periodo, pois o WebHook vai enviar tudo como um evento só o que pode gerar ruidos.

WebSockets 

é um protoco que utiliza dois caminhos de canais para se comunicar e se conectar com um único canal TCP. Esse é um dos melhores protocolos para web cliente and server e em alguns momentos para server-to-server também possui uma boa comunicação.

Como funciona?

Bem ele funciona permitindo comunicação simultanea entre server e cliente se comunicando um com o outro simultaneamente.  Eles tem um designed que funciona além da porta 80 ou 443 permitindo que também funcione bem quando o firewall bloquear outra porta

WebSockets é possui uma boa velocidade, transmissão de dados em tempo real e conexões duradoras. 

HTTP Streaming

o cliente envia a requesição e o server retorna um resposta HTTP de um comprimento finito. E agora isso é possível de se fazer pois o tamanho da resposta é indefinido. Com o HTTP Streaming, o server pode continuar a puxar novos dados com uma única e duradora conexão aberta pelo cliente. Essa transmissão de dados continua persistindo conectada pelo servidor até o cliente por duas opções: primeira o serve que vai definitir o Transfer-Encoding header em blocos. E isso indica para o cliente que aquelas dados devem está chegando em blocos definidos em uma nova linha delimitada em texto. A outra opção é ennviar os dados via server-set events (SSE). Essa opção é uma boa escolha para clientes que estão consumindo aqueles eventos no navegador por isso eles podem usar uma API de eventos padronizada.

Com isso fechamos o segundo capitulo pensando em tudo que temos que melhorar e o que podemos explorar!



