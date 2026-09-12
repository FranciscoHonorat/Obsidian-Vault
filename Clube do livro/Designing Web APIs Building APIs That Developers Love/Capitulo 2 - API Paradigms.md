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
