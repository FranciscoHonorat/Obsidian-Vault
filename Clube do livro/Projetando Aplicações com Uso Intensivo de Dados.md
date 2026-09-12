# Capitulo 1 - Trade-offs na arquitetura de sistemas de dados
O primeiro tópico ele explora uma geral sobre o livro, falando de maneira bem geral sobre como é armazenado os dados atualmente,  as dificuldades trazidas pelo aumento de dados e também as soluções. Ele cita que uma aplicação que faz uso intensivo de dados quando o gerenciamento de dados é um dos principais desafios em seu desenvolvimento. Enquanto em sistemas com uso intensivo de computação o desafio está em parelelizar um cálculo muito grande, em aplicações intensivas em dados, normalmente nos preocupamos mais com aspectos como armazenar e processar grandes volumes de dados, gerenciar alterações nos dados, garantir consistência diante de falas e concorrência e asegurar que serviços mantenhma alta disponibilidade.

Ele fala que geralmente essas aplicações são construidas com componentes básicos padronizados que oferecem funcionalidades comumente necessárias:

Amazenar dados para que possam ser recuperados posteriormente
Lembrar o resultado de uma operação custosa para acelerar as leituras
Permitir que os usuários pesquisem dados por palavra chave ou filtrem de diversas maneiras
Tratar eventos e mudanças nos dados assim que ocorrerem
Processar periodicamente grandes volumes de daods acumulados

Ele fala como ao longo do livro vamos aprender a fazer perguntas para avaliar e comparar sistemas de dados, identificando a abordagem que melhor atende às necessidade de uma aplicação específica.

O segundo tópico ele trás uma comparação entre sistemas operacionais e sistemas analíticos e fala dos responsaveis por cuidar de cada um desses sistemas

Sistemas operacionais consistem em serviços de backend e na infraestrutura de dados, nos qauis os dados são gerados, por exemplo, ao atender usuários externos. OLTP
Sistemas analíticos atendem as necessidades de analistas de negócio e de cientistas de dados. Eles contêm apenas uma cópia read-only dos dados provenientes dos sistemas operacionais e são otmizados para os tpos de processamente necessários às análises. OLAP

E ambos os sistemas são importantes para o ciclo de vida dos dados.

Antigamente o processamento de dados de negócios, uma gravação no banco de dados normalmente correspondia a uma transação comercial em andamento: realizar uma venda, fazer um pedido a um fornecedor, pagar o salário de um funcionário etc. A médida que os bancos de dados cresceram para áreas sem movimentos financeiros o nome transação permaneceu e esses mesmo bancos passaram a serem usados para finitas coisas, o padrão básico de acesso permaneceu semelhante ao do processamento de transações de neǵocios. Um sistema operacional normalemnte recupera um pequeno conjunto de registros por meio de uma chave (chamada point query). Os registros podem ser apagados, atualizados ou inseridos com base nas entradas do usuário. Como essas aplicações são interativas, esse padroão de acessou a passou a ser conhecido como processamento de transações online (Online Transaction Processing, ou OLTP).

Porém os bancos de dados continuaram a serem utilizados cada vez mais para análise de dados que apresneta padrões de acesso bastante diferente em comparação com OLTP.  Normalmente, uma consulta analítica percorre um volume enorme de registros e calcula estíticas agreda em vez de retorna registros individuais ao usuário. Um analista de negócios de uma rede de supermercados pode querer responder a consultas analíticas como estas:

QUal foi a receita total de cada uam das nossas lojas em janeiro?
Quantas bananas a mais do que o normal vendemos na nossa última promoção?
Qual marca de alimento para bebês é mais frequentemente comparada junto com a marca X de fraldas?

Os relatórios resultantes desse tpo de consulta são importantes para o BI, ajudnaod a gestão a decidir quais ações tomar. Para diferenciar esse padrão de uso de banco de dados do processamento de transações, passou a ser chamado de processamento analítico online (ONline Analytical Processing, ou OLAP)

Sistemas OLTP executam principalmente conjuntos fixo de consultas incorporada ao código da aplicação, euqnato consultas personalizadas pontuais são utilizadas apenas ocasionalmente para manutenção ou solução de problemas. Sobre banco de dados analiticos geralmente oferecem ao seus usuários a liberdade de escrever consultas SQL arbitrárias manualmente ou de gerá-las automaticmaente por meio de ferramente de sualização de daods ou de um dashboard, como Tableu, Looker ou MIcrosft Power Bi.

Quando um srema é projeto apra carga de trabalho. Sistemas projetados para esse tipo de uso, conhecidos como análise de produto ou análise em tempo geral, incluem Pinot, Druid e ClickHouse. Esses sisteams ingerem dados em tempo real e são otimizados para respsppostas de consultas com baixa latência. Por outro lado sistemas OLAP tradicionais normalmente ingerem dados em lote e são otimizados para prcoessamento de consultas com alta taxa de transferência.