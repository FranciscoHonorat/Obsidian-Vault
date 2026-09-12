```
func (o *Order) Place () {
	evt := event.NewOrderPlaced()
	o.event = append(o.evnt, evt)
}
```
Perguntas:

1 - O que acontece se eu chamar order.Place() logo depois de NewOrder(), sem nunca ter chamado AddItem() ? Existe alguma invariante impedindo isso?

eu acredito que isso não deveria ser possível, pois o Place deveria ser um evento interno de order e mostrar o resultado do evento gerado ao additem ser chamado, deveria ter alguma invariante que iria impedir isso, mas eu não coloquei. - Correto

2 - Depois que Place() roda, o que muda no estado da order? Ela ainda está OrderStatusPending? Isso faz sentido para você, semanticamente, uma order "colocada" continuar com status de pendente?

nada, pois o place() é apenas um evento, ela tem a função de registrar as coisas que vão sair do order... Errado

Um Domain Event não é uma noticia solta que sai da order sem afetar ela. Ele é o registro de uma fator que já aconteceu com o agregado. Se o evento se chama OrderPlaced - passando, fato consumado, então o agregado precisa ter mudado de estado para esse fato ser verdade. Caso contrário você teria duas fontes de verdade discordando entre si:
- O evento diz "Essa order foi colocada"
- O campo status diz "Essa order ainda está pendente"
Isso não é só uma inconsistência teórica. Pensa em consequências práticas: se o serviço de Inventory ou Notification reagir ao evento OrderPlaced (via kafka, mais para frente no seu roadmap), ele vai agir como se a order estivesse "colocada" - mas se alguém no seu próprio sistema perguntar order.Status() nesse meio tempo, a resposta vai ser Pending. O agregado está mentindo pra uma das duas partes.

A regra geral: o evento e o campo de estado são duas visões do mesmo fato - nunca podem divergir. Place() não é "só disparar um evento". Place() é uma transição de estado que, como consequência, também produz um evento registrado que a transição acontonteceu.

Novas perguntas:

Pergunta para pensar antes da próxima resposta: dado isso, o que Place() precisa fazer, em que orderm - mudar o status primeiro e depois criar um evento, ou o contrário? tem alguma razão de design para ordem importar aqui?

Quais duas condições, no mínimo, Place() precisa verificar antes de permitir a transição? Pensa em "tem conteúdo" e "está no estado certo pra ser colocada".

