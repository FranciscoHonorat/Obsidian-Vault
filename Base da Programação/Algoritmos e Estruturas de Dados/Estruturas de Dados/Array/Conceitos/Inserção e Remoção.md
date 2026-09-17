Inserir significa adicionar um elemento novo. Mas a complexidade depende de onde você quer inserir.

Como funciona a inserção no final: se o array tiver um espaço vazio disponível, não vamos precisar mover os elementos anteriores O(1). Em um array dinâmico, como o ArrayList em Java, essa operação é geralmente O(1) amortizado, porque ocasionalmente o Array precisa crescer e copiar os elementos.

Agora quando queremos inserir um array no início para termos que preservar a ordem a gente precisa deslocar os elementos ou seja quanto maior o array, mais elementos vão ser deslocados, portanto inserir no início é O(n) e o mesmo funciona quando queremos inserir no meio de um array também vamos ter uma complexidade O(n).

Porém a remoção no final funciona no mesmo mecanismo de inserir no final, complexidade O(1), pois não vamos precisar deslocar, porém no início e no meio funciona do mesmo modo que para inserir tendo complexidade O(n).

Bem, mas por que isso acontece?

Então a ideia principal é: Array armazenam seus elementos de forma contígua, e quando queremos manter a ordem, inserir ou remover no início/meio exige deslocar elementos.

Inserção e Remoção em array são O(1) no final, quando não há necessidade de  redimensionamento. No início ou no meio, são O(n) porque os elementos precisam ser deslocados para manter a ordem.

#Conceitos