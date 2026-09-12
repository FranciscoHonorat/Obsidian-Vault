# Capitulo 1:



# Capitulo 10: Estilo de Arquitetura em Camadas
Perguntas de auto Avaliação:
1 - Qual é a diferença entre camada aberta e fechada?
2 - Descreva o conceito das camadas de isolamento e quais os benefício
3 - O que é o padrão sinkhole da arquitetura?
4 - Quais são as principais características da arquitetura que orientariam o uso de uma arquitetura em camadas?
5 - Por que a testabilidade não é bem suportada no estilo da arquitetura em camadas?
6 - Por que a agilidade não é bem suportada no estilo da arquitetura em camadas?

Anotações:

O autor explica os motivos pelo qual a arquitetura de camadas é comum, pois quase sempre que começamos a codificar quase sempre estamos utilizando uma arquitetura de camadas para tal. Ele também fala como ela é uma arquitetura padrão para diversas aplicações, basicamente por sua simplicidade, familiaridade e baixo custo. Também é citado a lei de Conway que determina que as organizações que projetam sistemas estão limitadas a produizr designs que são cópias das estruturas de comunicação dessas organizações.

Mas o que isso quer dizer? 

e ele complementa falando sobre como a maioria das organizações é dividida em  desenvolvedores UI, backend, regras e especialistas em DB. E essas camadas organizacionais se encaixam bem nos níveis de uma arquitetura em camadas tradicional. 

E esse estilo de arquitetura em camadas se enquadra bem em vários antipadrões da arquitetura, inclusive os da arquitetura por implicação e da arquitetura acidental.

Topologia

O auto nessa sessão afima que os componentes no estilo de arquitetura em camadas são organizadas em camdas horizontais lógicas ou seja cada camada possui sua função específica dentro da aplicação.  Embora não possua um limite de quantas camadas uma aplicação possa ter, ela segue um base consistente de quatro camadas bem definidas que são: apresentação, comercial, persistência e banco de dados. Ele também fala como em alguns casos algumas camadas podem virar uma só e cita o exemplo da camada comercial e persistência podem virar uma só chamada de camada de negócio e em particular quanto a lógica da persistência (como SQL ou HSQL) está incorporada nos componenetes da camada de negócio. Sendo assim aplicações menores podem ter apenas 3 camadas e aplicações maiores e  mais complexas possuindo uma quantidade maior de camadas.

Logo em seguida ele apresenta diversas variantes da topologia da perspectiva da camada física (implementação). sendo a primeira dela uma combinação das camadas de apresentação, comercial e persistência em uma única camada e banco de dados em uma camada separada geralmente externa. A segunda variação é uma camada de apresentação com sua própria unidade de implementação e uma segunda camada com comercial e persistência juntas e novamente o banco de dados separado com um banco externo. A terceira variação combina todas as camadas padrão em uma única implementação, essa variação pode ter um banco de dados interno ou na memória. Muitos produtos on-premises são criados e entregues nessa terceira variação.

Cada camada dentor do estilo da arquitetura de camadas possui uma função  e responsabilidade, por exemplo, a camada de apresentação seria responsável por lidar com a interface do usuário e com a lógica de comunicação do navegador. já a camda de negócio seria a responsável por executar as regras de comerciais específicas associadas à requisição. Cada acamada da arquitetura forma uma abstração em torno do trabalho que precisa ser feito para atender certa requisição comercial.

O conceito de seperação das preocupações no estilo de arquitetura em camadas facilita a criação de funções eficientes e de modelos de reponsabilidade dentro da arquitetura. O que faz cada camada cuidar únicamente da lógica que pertence a sua camada. Por exemplo, os componentes na camada de apresentação só lidam com a lógica de apresentação. O que permite aos desenvolvedores utilizem sua expertise técnica em particular para focar os aspectos técnicos do domínio. Contudo o trade-off desse benefício é uma falta de agilidade geral (a capacidade de responder rápido à mudanças).

A arquitetura em camadas é particionada tecnicamente (em oposição a uma arquitetura particionada por dominio). Ou seja os grupos de componentes invez de serem agrupados por dominio, são agrupados por sua função tecnica na arquitetura. Como resultado disso o dominio fica distribuido em todas as camadas da arquitetura. Por exemplo, o domínio "cliente" é contido nas camadas de apresentação, comercial e das regras, dos serviços e do banco de dados, dificultando fazer alterações nesse domínio, de modo que a abordagem de design orientada a domínios não funciona bem com o estilo de arquitetura em camadas.

As arquiteturas em camadas são dividas em dois tipos: fechada e aberta

Uma camada fechada significa que, conforme uma requisição desce de camda em cada, ela não pode pular nenhuma, mas pode passar pela camada imediatamente abaixo dela para chegar na próxima. Exemplo, em uma arquitetura fechada uma requisição que tem origem na camada de apresentação deve passar primeiro pela camada de negócio e depois a camada de persistência e finalmente o banco de dados. Ai fica a pergunta não seria mais fácil ir direto para a última camada sem passar nas outras? então para isso acontecer seria necessário a camada de negócio e persistência serem abertas, perimitindo que as solicitações evitassem as outras camadas.

Então qual é melhor o autor pergunta e ele responde apresentando um conceito-chave conhecido como camadas de isolamento. 

O conceito de camadas de isolamento significa que as alterações feitas em uma camada da arquitetura normalmente não impactam nem afetam os componentes das outras camadas, fazendo com que os contratos entre essas camadas continuem inalterados. Uma camada independente das outras, sendo assim pouco ou nenhum conhecimento dos trabalhos internos das outras camadas na arquitetura. Contudo, para dar suporte a cadama de isolamento, as envolvias no fluxo maior da requisição devem ser necessariamente estar fechadas. Se uma camada tivesse acesso direto a outra camadas, as alterações feitas em uma afetaria todas as outras o que tornaria uma aplicação muito acoplada com interdependências das camadas entre os componentes. O que torna esse tipo de arquitetura muito frágil, dificil e caro de alterar.

Esse conceito também permite que qualquer camada seja substituida sem impactar a outra camada (obivamente se os contratos estiverem bem definidos e o uso do padrão de delegação comercial) Por exemplo, é possivel trocar uma camada de apresentação antiga em JSF por uma em react sem impactar nenhuma outra camada na aplicação.

Aqui fala justamente de um dos motivos dessa arquitetura ter problemas que é a falta de administração e controle caso uma camada utilize componentes de outras camadas. e para resolver esse problema geralmente se adiciona um camada com elementos compartilhados que é uma camada aberta.

A utilização desse conceito de camadas abertas e fechadas ajuda a definir a relação entre as camadas da arquitetura e os fluxos da requisição e para isso é necessário comunicação para conseguir orientar os desenvolvedores as restrinções de cada camada dentro da arquitetura. Geralmente a falha nessa comunicação ou documentação correta sobre quais camadas estão abertas ou fechadas e o motivo que costuma causar uma arquitetura muito acoplada  e frágil que são dificeis de testar, manter e implementar. 

Bem, geralmente essa arquitetura é usada para gerar projetos baratos e rápidos, enquanto escolhemos a arquitetura mais adequeada para determinos dominios. Por isso devemos manter a reutilização no mínimo e manter hierarquias de objetos bem rasas para manter um bom nível de modularidade.

Fique de olhos atento ao antipadrão sinkhole da arquitetura. Esse antipadrão ocorre quando as solicitações passam de camda em camada como um processamento de passagem simples sem nnehum lógica de negócio realizada dentro de cada uma.

Toda arquitetura em camadas terá, pelo menos, alguns cenários que caem no antipadrão sinkhole da arquitetura. O segredo para determinar se esse anti padrão está em ação é analisar a porcentagem de solicitações que ficam nessa categoria. A regra 80-20 costuma ser uma boa prática a seguir. Por exemplo, é aceitável se apenas 20% das solicitações são sinkholes. Contudo, se 80% delas forem sinkhole, é um bom indicador de que a arquitetura em camadas não é o estilo correto apra o domínio do problema. Outra abordagem para resolver o antipadrão sinkhole é torna todas as camadas na arquitetura abertas, percebendo, claro que o trade-off é uma maior dificuldade em gerenciar a alteração na arquiteutra.

Então, por que usar esse estilo de arquitetura? O auto explica que o motivo é bem simples, caso eu tenha um site ou uma aplicação pequenas é uma boa escolha. Ou caso eu precise iniciar logo e tenho pouco tempo e orçamento para tal é uma boa escolha. Porém conforme as necessidades da aplicação for crescendo caracteristicas como manutenção, agilidade, testabilidade e implementabilidade são afetas negativamente.

Os pontos mais fortes dessa arquitetura é seu custo geral e simplicidade. Sendo monolíticas por natureza não têm as complexidades associadas aos estilos da arquitetura distribuída e são simples e fáceis de entender, e têm custo relativamente baixo para criar e manter. Porém nem tudo é flores e essas classificações começam a diminuir rápidamente conforme as arquiteturas em camadas monolíticas ficam maiores, e por consequência, mais complexas.

As taxas de implementabilidade e testabilidade são muito baixas. As taxas de implementação são baixas devido à formalidade da implementação, ao alto risco e a falta de implementações frequentes.. Uma simples alteração em 3 linhas quebra a aplicação inteira tendo que mudar quase todas as camadas. Começa com 3 linhas e termina com dezenas de linhas alteradas.
A baixa classificação em testabilidade também reflete esse cenário, com essa simples alteração de três linhas, a maioria dos desenvolvedores não passará horas executando o conjunto inteior de teste de regressão (memso que tal coisa existisse), em particular junto a dezenas de outras alterações sendo feitas na aplicação monolítica ao mesmo tempo. A testabilidade nesse modelo tem apenas duas estrelas e não uma, pois possui uma capacidade de simular ou esboçar componentes, o que facilita o esforço do teste geral.

Confiabilidade geral tem uma média nesse estilo de arquitetura devido a falta de tráfego da rede, largura de band e latência encontrados nas arquiteturas mais distribuídas. 

Elasticiade e escalabilidade têm classificação muito baixa para a arquitetura em camadas, basicamente devido as implementação monoliticas e à falta de modularidade da arquiteutra. Embora seja possível criar certas funções em escala monolítca mais do que outras, em geral esse esforço requeer técnicas de design muito complexas, como multithreading, mensageria interna e outras práticas paralelas de processamento, técnicas para as quais essa arquitetura não é muito adequada.

O desempenho é sempr euma das características interessante para classificar a arquitetura em camadas, A classificação de duas estrelas é porque o estilo de arquitetura simplesmente não serve para os sistemas de alto desempenho devido a falta de processamento paralelo, as camadas fechadas e ao antipadrão da arquitetura sinkhole. Como a escalabilidade, podemos lidar com o desempenho por meio do uso de ccache, do multithreading e outras, mas não é uma característica natural desse estilo de arquitetura.

As arquiteturas em camadas não tem suporte para a tolerância a falhas devido as implementações monoliticas e à falta de modularidade arquitetural. Se uma pequena parte de uma arquitetura em camadas tem uma condição de falha de memória, a unidade inteira da aplicação é impactada e falha. E mais, a disponibilidade geral é impactada devido ao alto MTTR (tempo medio para reparo) normalmente sentido pela maiorida das apalicações monolíticas, com os tempos de incialização variando de 2 minutos para aplicações menores até 15 minutos ou mais para aplicações maiores.