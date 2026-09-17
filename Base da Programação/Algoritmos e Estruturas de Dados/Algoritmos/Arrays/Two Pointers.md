
eu tenho 2 arrays de inteiros e ordernado em ordem não descrescente, e dois inteiros de numeros inteiros, m e n que representam o numero de elementos em nums1 e nums2 respectivamente.

combine nums1 e nums2 em um unico array ordenado em ordem não decrescente

o array final ordenado não deve ser retornado pela função em vez disso, deve ser armazenado dentro do array "nums1". Para acomodar isso, nums1 possui tamanho m+n, onde os primeiros m elementos representam os elementos que devem ser combinados e os ultimos n elementos estão definidos como ) e devem ser ignorados. nums2 possui tamanho n.

Vamos fazer um passo a passo:

1 crie três ponteiros:
	- um no final dos elementos válidos de nums1 (posição m-1)
	- um no final de nums2 (posição n-1)
	- um no final do array nums1 total (posição m+n-1)
2 comparar de trás para frente:
	- compare o elemento em nums1[p1] com nums2[p2]
	- o maior vai para nums1[p] (final)
	- decremente o ponteiro correspondente
3 continue até processar todos os elementos:
	- se nums2 acaba primeiro, nums1 já está pronto
	- se nums1 acaba primeiro, copie os restantes de nums2

Por que funciona?
- nums1 já tem espaço reservado no final
- não há risco de sobrescrever dados importantes
- complexidade: O(m+n) tempo, O(1) espaço extra

```
func merge(nums1 []int, m int, nums2 []int, n int) {
    p1 := m - 1       // Último elemento válido em nums1
    p2 := n - 1       // Último elemento em nums2
    p := m + n - 1    // Última posição de nums1
    
    for p1 >= 0 && p2 >= 0 {
        if nums1[p1] > nums2[p2] {
            nums1[p] = nums1[p1]
            p1--
        } else {
            nums1[p] = nums2[p2]
            p2--
        }
        p--
    }
    
    // Se nums2 ainda tem elementos, copia-os
    // (nums1 já está no lugar certo)
    for p2 >= 0 {
        nums1[p] = nums2[p2]
        p2--
        p--
    }
}
```

