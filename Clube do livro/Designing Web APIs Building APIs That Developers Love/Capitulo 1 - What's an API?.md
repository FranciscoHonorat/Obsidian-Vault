API vai muito além do proprio nome dela, frequentemente usamos essas siglas sem entender o real poder dela e esse primeiro tópico retrata isso muito bem, pois em diversos momentos não estamos pensando no quão grandioso é uma API, ela é realmente um peça que pode processar seus dados bancarios em minutos, resolvendo questões de leis e burocracias, além de pode conectar o usuario com a nuvem e diversos outros serviços. 

Um trecho que me chamou bastante atenção: 

API é um componente chave para o sucesso e escalabilidade de companhias como Amazon, Stripe, Google e Facebook. Para essas companhias observar e criar plataformas de negocios é está sempre expandindo seus negocios, API é uma peça importante  dessa desafio.

Então, podemos entender que no tópico dois é discutido como as API são pensadas para resolver problemas que geralmente são obtidos a partir de informações obetidas por interação entre usuário e plataforma e as companhias não tem um time especifico para resolver tais problemas.

Outro trecho que chama atenção é:

API permite negócios desenvolver produtos exclusivos de maneira rápida.

API não são apenas sobre interação técnica, mas sobre transferência de responsabilidade ou seja uma boa API não resolve apenas o problema, mas também deve resolver com segurança embutida.

Pensar no designing da API e como ela vai ser usada, por quem vai usar, por qual motivo vai usada, é muito importante, pois, é uma forma de prevenir futuras manutenções, o motivo disso é que mudar de um designing de API para outro é algo muito dificil e desafiador. Por isso é importante pensar na especificidade da nossa API e validar ela antes de começa a sua implementação, pois o custo de mudar de um design para o outro é extremamente alto para a maioria dos desenvolvedores.

fala sobre como a web é muito poderosa para a questão de inovação de produtos e tecnologias e o resultado disso é a importancia das API para obtenção de resultados e dito isso temos diversas maneira de incorporar os mais diversos modelos de API em um produto, podendo ser fundamental para campanhas, ser um peça chave para solucionar um problema critico e também podendo dar suporte para outros produtos.

Uma API está alinhada com o core busniess?

Essa parte fala sobre como a API deve ser um suporte para a aplicação final e não um concorrente.

Essa parte fala sobre como em mutos momentos as API são criadas para serem usadas internamente e só depois de algum tempo passa para o externo e o autor cita também alguns motivos para isso sendo o primeiro o potêncial que ela vai ter para o desenvolvedores externos e fala sobre a questão das criações de ecosistemas de desenvolvimento, drives para novas demandas de um determinado produto, ou habilitar novas companhias a criar novos produtos que ela não quer construir sozinha.

Dois modelos:

interno -> externo - Desenvolvendo API para os próprios clientes, depois percebendo que integrando com ferramentas externas era critico para o sucesso.
Trade-off: APIs já testadas interamente, mas com atritos quando os usos internos e externos divergiram

externo -> interno - Postura oposta desde o inicio, construiram para desenvolvedores terceirizados acessarem os dados, criou um sistema de ferramentas ao redor da plataforma e o GaphQL veio depois, com foco em resolver problemas dos desenvolvedores externos (payloads enormes, campos desnecessários.).

Ou seja fala sobre como as vezes as empresas desenvolve API internas e logo em seguida vende essas API ou distribui essas API para outras empresas poderem usar.

Outro trecho também ele fala sobre a criação da Slack's API que foi cirada para uso interno e depois foi para uso externo.

é realmente interessante notar que o Github adotou GaphQL como resposta às necessidades externas. Permitindo que devs externos especificassem exatamente quais campos precisavam, em vez de receber tudo de uma só vez que era um problema de quem trabalhava com REST e dados.

Porém as desvantagens foi que o troubleshooting é mais complexo porque diferentes clientes podem usar diferentes access pattern. Com REST, é mais previsível (um endpoint = um padrão)

A API é o produto, porém cada empresa utiliza para um caso diferente, Stripe pagamentos, Twilio para comunicação SMS, ligação de voz e mensagem. Ou seja no caso para essas duas empresas a construção de uma API está totalmente alinhada aos um produto único para os clientes deles. Então podemos entender que API é um produto que tem que está alinhado desde o inicio com o negócio da empresa.  Tanto o gerenciamento quanto ao atendimento dos usuarios, um API é um produto modelado o mais organizacional possível.

Dito isso, podemos entender que a API é feita com bases nas necessidades dos usuários.

Ao perguntas para expertes da industria: como criar uma boa API? 

As respostas que vamos receber é o que essa API deve fazer?

E ele fala tipo que uma boa API deve ser usabilidade, escalabilidade, perfomance, documentação e recursos para o desenvolvimento são importante e a configurações para o sucesso. E também sobre como é impossivel implementar todos esses fatores, mas o mais importante é o usuário final.

E outro detalhe importante é bem testada ao longo do tempo e mudanças geralmente são dificeis e inevitaveis. APIs são plataformas flexiveis para se conectar a empresas e o ritmo de mudança é variável em contexto empresarias, as mudanças são lentas e pequenas  do que em empresas pequenas que ainda não encontraram seu lugar no mercado, porém starups apresentam API mais valiosas do que grandes empresas podem usar.

Isso que dizer que em empresas grandes, vamos construir uma API e vamos ficar só apenas cuidado de alguns aspectos tecnicos ou mudando uma coisa ali e aqui, então em empresas pequenas o tempo inteiro vamos ta construindo APIs novas e implementando coisas novas o tempo todo o que pode nos torna mais flexiveis e capazes de trabalhar com diferentes dominios de negócio.