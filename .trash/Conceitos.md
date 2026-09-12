Um array (ou Vetor) é uma estrutura de dados, ou seja um coleção ordenada, linear e homogênea de elmentos, o que significa que todos os itens armazenados pertencem ao mesmo tipo de dado.

Uma das característica físicas mais marcantes de um vetor é que seus elementos são armazenados em posíções sequenciais e contíguas de memória (lado a lado na memória RAM)

Cada posição é identificada por um indice numérico que começa por 0 (zero), ou seja o primeiro elemento começa no índice 0.

Como os dados estão em sequência contígua, o sistema computacional não precisa percorrer por elemento para encontrar uma posição, ele calcula o endereço exato de memória instantaneamento através de uma conta arimética simples a partir do endereço base.

Desempenho e complexidade Assintótica (Big - O)

A eficiência de um vetor varia drasticamente dependendo da operação realizada.

Acesso ou leitura direta O(1) (Tempo Constante): Como é possível pular diretamente para qualquer índice usando cálculo aritmético, a leitura de qualquer posição ocorre de forma instantânea.

Busca por Valor - O(n) (Tempo Linear): Se o vetor não esitver odernado e você precisar encontrar um valor específico, será necessário examinar cada elemento sequencialmente a partir do início até encontr-alo (pesquissa linear). Se o vetor estiver ordenado, pode-se utilizar a busca binária para reduzir o tempo para O(log n).

Inserção - O(n) (Tempo Linear) Inserir um item no início ou no meio do vetor exige mover fisicamente todos os elementos posteriores uma posição para frente para abrir espaço.

Remoção - O(n) (Tempo Linear): Deletar um elemento exige deslocar todos os itens subsequentes para trás para preencher a lacuna dexiada.

Os vetores se destacam em computadores modernos devido à arquitetura do processador. Quando a CPU lê uma posição de memória contigua do vetor, ela carrega automaticamente os blocos vizinhos para a memória cache. Isso reduz drasticamente os caches misses e faz com que algoritmos com o Quick Sort tenham um desempenho prático excelente quando executados sobre vetores. A sua simplicidade de acesso permitem manipular e iterar sobre conjuntos de dados de tamanho conhecido de forma rápida e intuitiva.

O tamanho fixo (alocação estática) faz com que em implementações estáticas, o tamanho do vetor deve ser definido no momento do desenvolvimento ou compilação. Se o número de dados exceder essa capacidade pré-alocada, o sistema precisará reservar um novo bloco de memória maior e copiar todos os dados antigos. Os espaços ociosos, se for reservado um espaço maior do que realmente é utilizado, posições de memória ficarão reservadas e sem uso. Alto cuso de reorganização o que faz operações frequentes de inserção e remoção tornam o uso de vetores custoso computacionalmente devido à necessidade constante de deslocar elementos.

