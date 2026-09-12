
![[Pasted image 20260811003628.png]]
# Análise Completa: LitInGame à Luz de "Fundamentals of Software Architecture"

---

## 1. Características Arquiteturais: O Que Você Declarou vs O Que Faltou

No livro, **características arquiteturais** são as "-ilities" que dirigem decisões: performance, confiabilidade, escalabilidade, segurança, custo, aprendibilidade, etc.

### O Que Você Declarou (Bom):

|Característica|Como Aparece no Doc|Rigor|
|---|---|---|
|**Performance**|"Latência p95 < 1.5s para ML"|✓ Métrica clara|
|**Confiabilidade**|"Falhas em ML não derrubam site"|✓ Padrão descrito|
|**Modularidade**|"3 Quanta desacoplados"|✓ Estrutura visível|

### O Que Você Omitiu (Crítico):

|Característica|Por Que Importa|Gap no Doc|
|---|---|---|
|**Escalabilidade**|Quantos alunos simultâneos a plataforma suporta? 25? 250?|Nenhuma menção|
|**Elasticidade**|A carga varia por aula (8h-11h). PostgreSQL escala horizontalmente?|Ignorada|
|**Custo**|FastAPI/PostgreSQL é barato. Mas ML em produção (latência sub-segundo)?|Invisível|
|**Aprendibilidade**|Quando você formar um time, quão rápido um novo dev entende a arquitetura?|Não declarada|
|**Segurança**|LGPD é compliance, não característica. Mas **segurança de dados de menores**?|Ausente|

**O Problema:** Características arquiteturais não declaradas criam ambiguidade. A banca vai perguntar:

- _"Qual é o SLA da plataforma?"_ (reliability vs custo)
- _"Vocês planejam escalar para outras escolas?"_ (elasticidade)
- _"Como vocês deployam updates sem derrubar a atividade?"_ (deployability)

---

## 2. Decision Records (ADRs): Decisões Não Justificadas

O livro é rígido aqui: **toda decisão arquitetural deve ter um ADR**. Um ADR documenta:

1. **Status** (Proposed / Accepted / Deprecated)
2. **Context** (Por que precisamos decidir?)
3. **Decision** (O que decidimos?)
4. **Consequences** (Quais são os trade-offs?)
5. **Alternatives** (O que mais consideramos?)

Seu documento **não tem nenhum ADR formal**. Pior: as decisões críticas nem aparecem como decisões.

### ADRs Que Você Precisa:

#### ADR-01: Arquitetura Monolítica Modular vs Microserviços

**Status:** Proposed (você precisa aceitar explicitamente)

**Context:**  
O LitInGame é um projeto acadêmico de TCC com 4 meses. Equipe: 1 pessoa. Escala esperada: 25-50 alunos piloto.

**Decision:**  
Arquitetura monolito modular com 3 Quanta (Student, ML Engine, Admin), não microserviços.

**Consequences:**

- ✓ Mais rápido de desenvolver (sem overhead de communication/versioning)
- ✓ Mais fácil de debugar (stack trace é linear)
- ✗ Escalabilidade horizontal limitada (a ML Engine escala separadamente via job queue?)
- ✗ Deploy acoplado (um bug no frontend druba todo o serviço, incluso ML)

**Alternatives Considered:**

1. **Microserviços completos** (Student API, ML Service, Admin API, Message Queue)
    - Alternativa: escalabilidade, independência, mais operacional overhead para TCC
2. **Serverless (AWS Lambda)**
    - Alternativa: pay-per-use, Cold starts (crítico para ML latency)
3. **Monolito único** (tudo em uma classe FastAPI)
    - Alternativa: simplicidade máxima, mas nenhuma separação

---

#### ADR-02: Classificação ML: TF-IDF + SVM vs Transformers

**Status:** Proposed (falta justificativa)

**Context:**  
Você precisa classificar respostas de alunos em 3 categorias (0, 1, 2). Restrições:

- Latência sub-segundo em hardware escolar (CPU, não GPU)
- Interpretabilidade para professor
- Corpus inicial pequeno (~200-300 exemplos)

**Decision:**  
TF-IDF + SVM. Não BERT/BERTimbau.

**Consequences:**

- ✓ Roda em CPU, latência previsível
- ✓ Features (TF-IDF terms) interpretáveis
- ✓ Rápido de treinar (< 1min)
- ✗ Acurácia inferior a BERT em corpus pequeno (esperado: 65-75% vs 78-85%)
- ✗ Sem transfer learning (BERT traz conhecimento pré-treinado)

**Alternatives:**

1. **BERTimbau (BERT português)**
    - Melhor acurácia, mas GPU obrigatória (não viável em escolas)
2. **Zero-shot com GPT-4 mini (API)**
    - Melhor acurácia, mas dependência de API externa (custo, latência, LGPD)
3. **Naive Bayes**
    - Mais simples que SVM, acurácia pior, sem probabilidade calibrada

**Recomendação:**  
Coloque isso no documento. Banca vai perguntar. Você precisa ter resposta técnica.

---

#### ADR-03: Banco de Dados: PostgreSQL Monolítico

**Status:** Proposed

**Context:**  
25-50 alunos × 10 respostas/aluno ≈ 250-500 registros de resposta em Fase 4. Crescimento pós-TCC? Desconhecido.

**Decision:**  
PostgreSQL single-node. Sem particionamento, sem replicação, sem read replicas.

**Consequences:**

- ✓ Operacionalmente simples
- ✓ ACID transactions nativas
- ✓ Índices de texto full-search (útil para auditoria de professor)
- ✗ Ponto único de falha (DB cai = tudo cai)
- ✗ Escalabilidade vertical limitada (não horizontal)
- ✗ Backup: precisa de estratégia

**Alternatives:**

1. **PostgreSQL + Replicação (standby)**
    - Mais resiliente, mas complexidade operacional
2. **MongoDB + sharding**
    - Mais escalável, mas menos confiável para dados críticos
3. **Serverless (DynamoDB, Cloud Firestore)**
    - Menos operação, mais custo (LGPD: dados em datacenter Brasil?)

---

## 3. Fitness Functions Inadequadas

Você tem 3 fitness functions. Mas pelo framework Richards & Ford, você precisa de **uma FF para cada característica arquitetural crítica**.

### O Que Você Tem:

|FF|Métrica|Status|
|---|---|---|
|FF1|Latência ML p95 < 1500ms|✓ OK, medível|
|FF2|Complexidade ciclomática V(G) ≤ 10|⚠ OK, mas limiar alto|
|FF3|Confiança < 70% → flag revisão|✓ OK, automático|

### O Que Falta:

|FF|Característica|Como Medir|
|---|---|---|
|**FF4**|**Segurança: LGPD**|Teste: deletar aluno → logs auditados? Sem orphans?|
|**FF5**|**Confiabilidade: Fallback**|Teste: mata PostgreSQL → site ainda carrega vídeos e legendas?|
|**FF6**|**Modularidade**|Teste: consegue trocar TF-IDF por BERT sem mudar API? (Plugin interface)|
|**FF7**|**Custo**|Monitorar: GB PostgreSQL, chamadas spaCy/mês (meta: < R$ 50/mês)|

**Como adicionar ao documento:**

````markdown
### 4. Fitness Functions Expandidas

#### FF4 — Compliance LGPD
Objetivo: Garantir que dados de menores podem ser deletados sem deixar orphans.

Teste (pytest):
```python
def test_aluno_deletado_sem_orphans():
    aluno_id = create_test_aluno("João", "12345678901")
    resposta_id = create_test_resposta(aluno_id, "meu texto")
    
    # Deleta aluno
    db.delete_aluno(aluno_id)
    
    # Verifica cascata
    assert db.count_respostas(aluno_id=aluno_id) == 0
    assert db.audit_log.filter(action='delete', aluno_id=aluno_id).count() >= 1
````

Frequência: A cada commit (integração contínua).

#### FF5 — Confiabilidade: Fallback

Objetivo: Se PostgreSQL cair, alunos ainda conseguem assistir vídeos e legendas.

Teste (integração):

```python
def test_ml_engine_fails_gracefully():
    # Mata ML engine
    kill_process("spacy_worker")
    
    # Tenta enviar resposta dissertativa
    response = client.post("/resposta", data={...})
    
    # Não falha 500, apenas retorna flag 'requer_revisao=true'
    assert response.status_code == 200
    assert response.json()["confianca_modelo"] is None
    assert response.json()["requer_revisao"] == True
```

Frequência: Semanal (staging).

#### FF6 — Modularidade: Plugin Interface para ML

Objetivo: Trocar TF-IDF por BERTimbau sem mudar API.

Design:

```python
# Interface abstrata
class ClassifierPlugin(ABC):
    @abstractmethod
    def predict(self, texto: str) -> Tuple[int, float]:
        """Retorna (classe: 0-2, confiança: 0-1)"""
        pass

# Implementações
class SKLearnClassifier(ClassifierPlugin):
    def predict(self, texto: str):
        # TF-IDF + SVM
        ...

class BERTClassifier(ClassifierPlugin):
    def predict(self, texto: str):
        # BERT português
        ...

# Uso
classifier = SKLearnClassifier()  # Troca aqui
predicts = classifier.predict("meu texto")
```

Teste:

```python
@pytest.mark.parametrize("classifier_cls", [SKLearnClassifier, BERTClassifier])
def test_classifier_interface(classifier_cls):
    clf = classifier_cls()
    classe, conf = clf.predict("Exemplo de resposta")
    assert isinstance(classe, int)
    assert 0 <= conf <= 1
```

Frequência: A cada commit.

```

---

## 4. Trade-offs Não Explicitados (Crítico)

No livro, **trade-offs são a essência da arquitetura**. Exemplos:

### Trade-off 1: Monolito vs Escalabilidade

| Escolha | Pro | Con |
|---|---|---|
| **Monolito modular** (você escolheu) | Fácil de desenvolver, um só deploy | Escala só verticalmente, falha acoplada |
| Microserviços | Escala independente, resiliência | Complexidade, latência entre serviços, LGPD (replicação) |

**Seu trade-off:** Priorizou velocidade de desenvolvimento (TCC de 4 meses) sobre escalabilidade (pós-TCC).

**Você precisa documentar:**
```

Trade-off: Simplicidade Arquitetural vs Elasticidade

Decisão: Monolito modular (Quanta).

Justificativa:

- Timeline: TCC em 16 semanas, 1 pessoa.
- Escala piloto: 25-50 alunos, 1-2 turmas simultâneas.
- Horizonte: Pós-TCC, se expandir, refatorar para job queue + worker pool.

Custo da Decisão:

- Não escala horizontalmente (mas não é necessário ainda).
- Deploy afeta todo sistema (mitigação: testes + CI/CD rigoroso).

```

### Trade-off 2: SVM vs Acurácia

| Escolha | Pro | Con |
|---|---|---|
| **SVM + TF-IDF** (você escolheu) | CPU-only, rápido | Acurácia ~70% (esperado) |
| BERT | Acurácia ~85% | GPU obrigatória |

**Seu trade-off:** Priorizou infraestrutura viável sobre acurácia máxima.

**Você precisa documentar:**
```

Trade-off: Acurácia vs Infraestrutura Disponível

Decisão: SVM + TF-IDF (latência/custo sobre acurácia).

Métricas esperadas:

- F1-score: 0.65-0.75 (meta: ≥ 0.70)
- Acurácia: 68-78%

Mitigação de baixa acurácia:

- Human-in-the-loop: < 70% confiança → professor revisa
- Iteração: Se acurácia muito baixa, considerar BERT em Fase 2 (pós-TCC)

````

---

## 5. O Que Adicionar Agora Mesmo

### Seção 3.5: Architecture Decision Records (ADRs)

```markdown
## 3.5 Architecture Decision Records (ADRs)

Este projeto segue a metodologia ADR conforme "Fundamentals of Software Architecture" (Richards & Ford).
Cada decisão arquitetural crítica é documentada com contexto, alternativas e trade-offs.

### ADR-01: Monolithic Modular vs Microservices
[Completo conforme acima]

### ADR-02: SVM + TF-IDF vs BERT Classification
[Completo conforme acima]

### ADR-03: PostgreSQL Single-Node Storage
[Completo conforme acima]

### ADR-04: Latency Threshold: < 1.5s vs < 500ms
**Status:** Accepted

**Context:**  
Alunos em sala de aula esperam resposta rápida do feedback da ML.

**Decision:**  
p95 latência < 1.5s (1500ms).

**Reasoning:**  
- < 500ms: Requer cache agressivo ou modelo no edge (complexo)
- 1.0s: Praticamente impossível com SVM (parsing + prediction)
- 1.5s: Tolerable (professor pede "espera um segundo"), atingível com SVM

**Consequences:**
- ✓ Atingível com CPU
- ✗ UX não é "instantânea" (perceived latency)

**Mitigation:**  
Loading spinner + progress bar durante inferência.
````

### Seção 3.6: Trade-off Analysis

```markdown
## 3.6 Architectural Trade-offs

Todo projeto envolve trade-offs. Este documento explicita os principais:

| Decisão | Priorizado | Sacrificado | Justificativa |
|---------|-----------|-----------|-----------|
| Monolito modular | Velocidade dev | Escalabilidade horizontal | TCC em 4 meses |
| SVM + TF-IDF | Simplicidade operacional | Acurácia máxima | Sem GPU em escolas |
| PostgreSQL single-node | Simplicidade | Disponibilidade HA | Piloto < 50 alunos |
| FastAPI | Produtividade Python | Comunidade vs Node.js | Stack Letras/Dev já existe |

**Transições Futuras (Pós-TCC):**
- Se escalar para 10+ escolas: Microserviços ML + job queue
- Se acurácia baixa: Migrar SVM → BERTimbau (com GPU)
- Se dados críticos: PostgreSQL + replicação standby
```

---

## 6. Checklist: Seu Documento vs Livro

Agora você tem um framework claro. Antes de apresentar, verifique:

- [ ] **Características Arquiteturais Declaradas:** Pelo menos 5 explícitas (performance, confiabilidade, escalabilidade, custo, aprendibilidade)
- [ ] **ADRs Formais:** Mínimo 4 ADRs com Status/Context/Decision/Consequences
- [ ] **Fitness Functions:** 1 FF por característica crítica (≥ 6 FFs)
- [ ] **Trade-offs:** Todas as decisões têm trade-off documentado
- [ ] **Alternativas:** Cada ADR menciona 2-3 alternativas rejeitadas e por quê
- [ ] **Arquitetura:** Diagramas mostram padrão (monolito, modular, quantum, etc)

---

## 7. Última Coisa: Sua Vantagem Letras

O livro de Richards & Ford é **normativo** (como deve ser). Mas você tem um ângulo que a maioria dos engenheiros não tem: **linguagem e comunicação**.

Seu documento pode ser o primeiro a:

1. **Usar ADRs com prosa clara** (não jargão de eng)
2. **Justificar trade-offs em termos pedagógicos** ("Priorizei monolito porque preciso iterar rápido com feedback do professor")
3. **Mapear características arquiteturais para objetivos educacionais** ("Confiabilidade não é SLA, é garantir que aluno não perde resposta que digitou")

Isso impressiona banca. Não é "apenas mais um TCC de tech". É **engenharia aplicada com rigor e reflexão**.

Faça isso. Seus pares engenheiros vão reconhecer o livro. Seus pares Letras vão reconhecer a clareza.