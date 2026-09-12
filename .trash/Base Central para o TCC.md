Esta revisão traz o **salto de maturidade final** para o seu trabalho. Integrar a literatura de **Richards & Ford ("Fundamentals of Software Architecture")** refuta qualquer argumento da banca de que o projeto é apenas uma "aplicação web comum". Ele passa a ser um estudo de **Arquitetura Evolutiva com Rigor de Engenharia de Software**.

  

Abaixo está o **Documento Consolidado e Expandido** pronto para inclusão no seu TCC/Projeto, incorporando as _-ilities_ omitidas, os **ADRs**, as **Fitness Functions expandidas** e a **Análise Formal de Trade-offs**.

  

# Documentação Arquitetural Refinada: LitInGame

_Conforme as diretrizes de "Fundamentals of Software Architecture" (Richards & Ford)_

  

## 1. Características Arquiteturais (_Architectural Characteristics / -ilities_)

As características arquiteturais definem as métricas de qualidade não-funcionais que governam a estrutura e as decisões do sistema.

  

|**Característica**|**Onde Aparece / Métrica**|**Justificativa & Mitigação de Trade-off**|
|---|---|---|
|**Performance**|Latência de Inferência $p95 < 1.5\text{s}$|Necessária para não quebrar o fluxo pedagógico em sala de aula. Mitigada no Frontend por meio de _loading state_ / progresso percebido.|
|**Confiabilidade**|Fallback gracioso da ML Engine ($0\%$ de _downtime_ do player de vídeo)|Se a ML falhar ou o banco oscilar, a atividade e a exibição do vídeo continuam funcionais; a avaliação é marcada para revisão posterior.|
|**Modularidade**|3 Quanta Arquiteturais isolados|Garante que o domínio do Estudante, da Engine de ML e do Painel Docente evoluam e sejam testados de forma independente.|
|**Escalabilidade & Elasticidade**|Piloto estático: 25 a 50 usuários simultâneos por turma|**Escala Delimitada:** Não há necessidade de elasticidade em tempo real na Fase 1. A carga ocorre em picos previsíveis (horários de aula).|
|**Custo-Efetividade**|Teto operacional: $< \text{R}\$ 50,00/\text{mês}$|Otimizado para infraestrutura acadêmica de baixo custo (CPU-only, sem dependência de GPUs caras ou APIs de LLM por token).|
|**Segurança & Compliance (LGPD)**|Anonimização de dados de menores e deleção em cascata|Garantia de conformidade legal para dados de estudantes da Educação Básica, com trilha de auditoria e ausência de dados órfãos.|
|**Aprendibilidade / Manutenibilidade**|_Onboarding_ de novos devs $< 2\text{ horas}$|Uso de código idiomático em Python (FastAPI + Pydantic) com documentação OpenAPI auto-gerada e padrões arquiteturais explícitos.|

## 2. Decision Records (ADRs Formais)

### **ADR-01: Escolha do Padrão Arquitetural — Monolito Modular vs. Microserviços**

- **Status:** Accepted
    
      
    
- **Context:** O _LitInGame_ é um TCC desenvolvido no prazo de 16 semanas por uma pessoa. O escopo do piloto prevê 25 a 50 alunos por aplicação.
    
      
    
- **Decision:** Adotar uma arquitetura de **Monolito Modular delimitada em 3 Quanta** (Student, ML Engine, Admin), utilizando rotas e módulos isolados no FastAPI, em vez de uma arquitetura de microserviços distribuída.
    
      
    
- **Consequences:**
    
      
    - $\checkmark$ Desenvolvimento e depuração acelerados (stack trace linear, sem overhead de comunicação via rede/mensageria).
        
          
        
    - $\checkmark$ Deploy unificado e de baixo custo.
        
          
        
    - $\times$ Escalabilidade horizontal limitada da Engine de ML no mesmo nó (aceitável para a escala do piloto).
        
          
        
    - $\times$ Risco de falha acoplada se não houver isolamento de exceções nas rotas.
        
          
        
- **Alternatives Considered:**
    
      
    - _Microserviços Completos:_ Rejeitado pelo alto overhead de infraestrutura, conteinerização múltipla e orquestração incompatíveis com a janela de 16 semanas.
        
          
        
    - _Serverless (AWS Lambda):_ Rejeitado devido ao risco de _cold start_ em bibliotecas de NLP (`spaCy`), prejudicando o requisito de latência.
        
          
        

### **ADR-02: Algoritmo de Classificação da ML — TF-IDF + SVM vs. Transformers (BERT/BERTimbau)**

- **Status:** Accepted
    
      
    
- **Context:** O sistema precisa avaliar a presença de marcadores discursivos nas respostas dos alunos com o corpus inicial pequeno (~200 a 300 amostras) e rodar em hardware escolar modesto (apenas CPU).
    
      
    
- **Decision:** Utilizar **Vetorização TF-IDF combinada com Support Vector Machines (SVM)** via `scikit-learn` e `spaCy` para o idioma português.
    
      
    
- **Consequences:**
    
      
    - $\checkmark$ Inferência ultrarrápida em CPU (latência na casa dos milissegundos).
        
          
        
    - $\checkmark$ Modelo leve ($< 50\text{ MB}$), de fácil deploy e baixo uso de memória RAM.
        
          
        
    - $\checkmark$ Interpretabilidade direta dos termos com maior peso para o professor.
        
          
        
    - $\times$ Acurácia esperada inferior ($68\% \text{ a } 78\%$) em comparação ao fine-tuning de modelos LLM/BERT.
        
          
        
- **Alternatives Considered:**
    
      
    - _BERTimbau (BERT Português):_ Rejeitado por exigir GPU para manter a latência sub-segundo em inferência e por requerer um volume de dados muito maior de treinamento.
        
          
        
    - _API Externa de LLM (GPT-4 mini / Claude):_ Rejeitada pelo custo recorrente por token, dependência de conexão de internet de alta velocidade nas escolas e potenciais atritos com a LGPD ao enviar dados de menores para terceiros.
        
          
        

### **ADR-03: Persistência de Dados — PostgreSQL Single-Node**

- **Status:** Accepted
    
      
    
- **Context:** O volume de dados estimado para a validação é de 250 a 500 registros de respostas. A prioridade é a integridade relacional entre Alunos, Turmas, Respostas e Logins.
    
      
    
- **Decision:** Adotar o **PostgreSQL** em nó único sem replicação ou sharding inicial.
    
      
    
- **Consequences:**
    
      
    - $\checkmark$ Suporte nativo a transações ACID e consultas relacionais complexas para o Dashboard do Professor.
        
          
        
    - $\checkmark$ Suporte a busca em texto (_Full-Text Search_) para auditoria rápida.
        
          
        
    - $\times$ Ponto único de falha (_Single Point of Failure_), aceitável para o ambiente acadêmico/piloto.
        
          
        
- **Alternatives Considered:**
    
      
    - _NoSQL (MongoDB):_ Rejeitado pela ausência de necessidade de esquemas flexíveis para os dados estruturados de alunos/respostas e perda de integridade relacional nativa.
        
          
        

### **ADR-04: Limiar de Latência e UX de Inferência — $p95 < 1.5\text{s}$**

- **Status:** Accepted
    
      
    
- **Context:** Alunos em sala de aula necessitam de um retorno rápido do sistema para manter o engajamento sem que a interface pareça travada.
    
      
    
- **Decision:** Fixar a meta de latência do percentil 95 ($p95$) em no máximo $1.5\text{ segundo}$ ($1500\text{ms}$) e implementar um _feedback visual_ (spinner/progress bar) no Frontend.
    
      
    
- **Consequences:**
    
      
    - $\checkmark$ Perfeita viabilidade de processamento síncrono no servidor via CPU.
        
          
        
    - $\times$ A experiência não é "instantânea" ($< 200\text{ms}$), exigindo gerenciamento de estado na interface do usuário.
        
          
        

## 3. Matriz e Análise de Trade-offs

Em conformidade com Richards & Ford, **toda decisão arquitetural é uma escolha de trade-offs**. A tabela abaixo explicita os sacrifícios conscientes feitos no projeto:

  

|**Decisão de Projeto**|**O Que Foi Priorizado**|**O Que Foi Sacrificado**|**Justificativa Contextual (TCC / Sala de Aula)**|
|---|---|---|---|
|**Monolito Modular**|Velocidade de desenvolvimento e simplicidade de deploy|Escalabilidade horizontal ilimitada|Restrição de tempo (16 semanas) e equipe (1 desenvolvedor). O impacto do sacrifício é nulo para o tamanho do piloto.|
|**TF-IDF + SVM**|Viabilidade de infraestrutura (CPU) e baixa latência|Acurácia máxima de NLP|Escolas públicas/privadas possuem infraestrutura limitada. É melhor ter um modelo leve rodando localmente com _Human-in-the-Loop_ do que dependência de GPU.|
|**PostgreSQL Single-Node**|Simplicidade operacional e consistência de dados|Alta Disponibilidade (HA) com Replicação|O custo e a complexidade de manter réplicas de banco de dados não se justificam para um volume de 50 alunos piloto.|
|**FastAPI (Python)**|Produtividade e ecossistema nativo de Ciência de Dados|Concorrência massiva I/O (comparado a Node.js/Go)|Integração perfeita com as bibliotecas de ML (`scikit-learn`, `spaCy`) no mesmo runtime da API.|

## 4. Fitness Functions Expandidas (_Suíte de Testes da Arquitetura_)

A suíte de testes de arquitetura automatizada garante que as características e limites do sistema sejam mantidos continuamente no CI/CD via `pytest`:

  

Python

```
# tests/test_architecture_fitness.py
import pytest
import time
from radon.complexity import cc_visit
from abc import ABC, abstractmethod

# --- FF1: Complexidade Ciclomática ≤ 10 ---
def test_ff1_cyclomatic_complexity():
    """Garante manutenibilidade limitando a complexidade das funções."""
    with open("app/ml/engine.py", "r") as f:
        blocks = cc_visit(f.read())
    for block in blocks:
        assert block.complexity <= 10, f"FF1 Violada! Função {block.name} com V(G)={block.complexity}"

# --- FF2: Latência p95 da ML < 1500ms ---
def test_ff2_ml_latency(ml_engine, dataset_amostra):
    """Garante que a latência de inferência em CPU permaneça dentro do SLA."""
    tempos = []
    for amostra in dataset_amostra:
        t0 = time.time()
        _ = ml_engine.predict(amostra.texto)
        tempos.append(time.time() - t0)
    tempos.sort()
    p95 = tempos[int(len(tempos) * 0.95)]
    assert p95 < 1.5, f"FF2 Violada! Latência p95 foi de {p95:.3f}s"

# --- FF3: Auditabilidade de Baixa Confiança (< 70%) ---
def test_ff3_low_confidence_fallback(client):
    """Respostas com confiança < 70% DEVEM ser marcadas para revisão docente."""
    response = client.post("/api/resposta", json={"texto": "resposta ambígua e muito curta"})
    data = response.json()
    if data["confianca_modelo"] < 0.70:
        assert data["requer_revisao_docente"] is True

# --- FF4: Compliance LGPD - Cascade Delete de Dados de Menores ---
def test_ff4_lgpd_cascade_delete(db_session, test_aluno):
    """Garante que a exclusão do aluno remove todos os seus dados sem deixar registros órfãos."""
    aluno_id = test_aluno.id
    db_session.delete(test_aluno)
    db_session.commit()
    
    respostas_orfas = db_session.query(Resposta).filter_by(aluno_id=aluno_id).count()
    assert respostas_orfas == 0, "FF4 Violada! Respostas órfãs encontradas no banco de dados."

# --- FF5: Confiabilidade - Degradação Graciosa (Fallback da ML) ---
def test_ff5_ml_graceful_degradation(client, monkeypatch):
    """Se a engine de ML falhar, a API deve responder 200 OK e marcar para revisão manual."""
    def mock_predict_failure(*args, **kwargs):
        raise RuntimeError("Engine de ML Indisponível")
    
    monkeypatch.setattr("app.ml.engine.MLEngine.predict", mock_predict_failure)
    
    response = client.post("/api/resposta", json={"texto": "Análise sobre o discurso do personagem"})
    assert response.status_code == 200
    assert response.json()["requer_revisao_docente"] is True
    assert response.json()["status"] == "salvo_sem_avaliacao_automatica"

# --- FF6: Interface de Plugin para o Classificador de ML ---
class ClassifierPlugin(ABC):
    @abstractmethod
    def predict(self, texto: str) -> tuple[int, float]:
        pass

def test_ff6_classifier_plugin_interface(classifier_instance):
    """Garante que qualquer futuro modelo (ex: BERT) respeite o contrato da interface."""
    assert isinstance(classifier_instance, ClassifierPlugin)
    classe, confianca = classifier_instance.predict("Texto de teste")
    assert isinstance(classe, int) and classe in [0, 1, 2]
    assert 0.0 <= confianca <= 1.0
```

## 5. Estrutura do Documento para Apresentação ao Orientador

Ao compilar este material para o texto final do TCC ou para enviar ao orientador, organizing a estrutura da seguinte forma trará grande clareza ao texto:

  

1. **Introdução:** Apresentação do problema pedagógico (letramento crítico) e justificativa da solução (plataforma web com suporte de ML).
    
      
    
2. **Fundamentação Teórica:**
    
      
    - Análise do Discurso e Letramento Crítico na Educação Básica.
        
          
        
    - Conceito de _Transfer of Learning_ (Near vs. Far Transfer).
        
          
        
    - **Arquitetura de Software Evolutiva** (Fundamentada em Richards & Ford).
        
          
        
3. **Engenharia e Arquitetura de Software:**
    
      
    - Detalhamento dos **3 Quanta Arquiteturais**.
        
          
        
    - Apresentação das **Características Arquiteturais** e **Matriz de Trade-offs**.
        
          
        
    - Registros de Decisão Arquitetural (**ADRs 01 a 04**).
        
          
        
    - Implementação das **Fitness Functions** via automação de testes.
        
          
        
4. **Metodologia de Aplicação e Validação:**
    
      
    - Estrutura das Atividades em 3 Etapas (Jogo $\rightarrow$ Desacoplamento $\rightarrow$ Vida Real).
        
          
        
    - Coleta de Métricas do Modelo e Desempenho dos Alunos.
        
          
        
5. **Considerações Finais e Trabalhos Futuros.**
    
      
    

Com esta estrutura, o seu projeto une **teoria linguística**, **prática pedagógica** e **Engenharia de Software de alto nível**, estando devidamente blindado contra questionamentos da banca examinadora.