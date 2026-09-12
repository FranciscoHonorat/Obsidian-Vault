back/
front/
infra/
QA/
docs/             
devops/            
data/              
design/           
scripts/           
security/         

Na verdade, to aqui pensando

eu tenho diversos backend feitos, obvio que falta refatorar muita coisa e melhorar bastante coisa também, mas na verdade o que eu preciso é seguir meu roteiro de testes, pois, meus projetos estão feitos, porém eles não foram testados ao limite e eu não consigo tirar deles as métricas Z Y X e para isso eu fiz um roteiro simples para consegui e tenho que além de implementar tudo isso também fazer a parte do roteiro.

# Roteiro de Testes para Projetos Pessoais

## Francisco — Portfolio Verification Workflow

---

## FASE 0: PRÉ-REQUISITOS (antes de qualquer teste)

### ✅ Checklist Pré-Teste

- [ ] Projeto está deployado em URL pública (Vercel, Render, Railway, etc)
- [ ] README.md está atualizado com instruções
- [ ] `.env` configurado corretamente no deploy
- [ ] Banco de dados (TimescaleDB, Redis) está rodando e acessível
- [ ] Aplicação inicializa sem erro: `npm start` ou equivalente
- [ ] Você tem acesso a logs da aplicação (stdout/stderr)

**Se não tem deploy live:** PAUSAR testes até ter. Deploy + live link é pré-requisito, não bonus.

---

## FASE 1: TESTES DE CARGA — Throughput & Latência

### Objetivo

Medir quantos eventos/requisições seu sistema processa por segundo e com qual latência.

### Passo 1.1: Instalar ferramentas (local)

```bash
# macOS
brew install apache2-utils vegeta

# Ubuntu/Debian
sudo apt-get install apache2-utils
# Para vegeta
go install github.com/tsenart/vegeta@latest
# Ou via npm
npm install -g vegeta
```

**Alternativa simples (sem instalar):** Use `curl` + bash loop (abaixo).

---

### Passo 1.2: Teste com Apache Bench (Simples)

**O que mede:** Quantas requisições por segundo, latência média/max

#### Cenário A: Endpoint simples (ex: GET /health ou /api/data)

```bash
# Teste: 1000 requisições, 10 paralelas
ab -n 1000 -c 10 https://seu-projeto.com/api/health

# Teste intenso: 5000 requisições, 50 paralelas
ab -n 5000 -c 50 https://seu-projeto.com/api/data
```

**Output esperado:**

```
Requests per second:    1234.56 [#/sec] (mean)
Time per request:       8.109 [ms] (mean)
Time per request:       0.811 [ms] (mean, across all concurrent requests)
Failed requests:        0
```

**Capture isso em um arquivo:**

```bash
ab -n 5000 -c 50 https://seu-projeto.com/api/data > resultado-load-test.txt 2>&1
```

---

### Passo 1.3: Teste com Vegeta (Mais realista)

**O que mede:** Distribuição de latência, percentis (p50, p95, p99)

#### Teste básico

```bash
# Criar arquivo com requisições
cat > requisicoes.txt <<EOF
GET https://seu-projeto.com/api/events
GET https://seu-projeto.com/api/logs
GET https://seu-projeto.com/api/metrics
EOF

# Rodar vegeta: 100 requisições por segundo por 30 segundos
cat requisicoes.txt | vegeta attack -duration=30s -rate=100 -timeout=10s | vegeta report

# Saída mais detalhada (com percentis)
cat requisicoes.txt | vegeta attack -duration=30s -rate=100 | vegeta report -type=text
```

**Output esperado:**

```
Requests      [total, rate, throughput]    3000, 100.00, 99.50
Duration      [total, attack, wait]        30.002s, 30.001s, 456.123ms
Latencies     [min, mean, 50, 95, 99, max] 45.2ms, 156.4ms, 142ms, 285ms, 450ms, 2.1s
Bytes In      [total, mean]                150000, 50.00
Bytes Out     [total, mean]                0, 0.00
Success       [ratio]                      99.80%
Status Codes  [code:count]                 200:2994, 500:6
```

---

### Passo 1.4: Teste customizado (curl + bash)

**Se não quer instalar nada:**

```bash
#!/bin/bash
# test-api.sh

TARGET_URL="https://seu-projeto.com/api/events"
REQUESTS=100
CONCURRENT=10

echo "=== TESTE DE CARGA ==="
echo "URL: $TARGET_URL"
echo "Requisições: $REQUESTS, Concurrent: $CONCURRENT"
echo ""

# Rodar benchmark simples
time for i in {1..100}; do
  curl -s "$TARGET_URL" > /dev/null &
done
wait

echo "Teste concluído. Verifique os logs da aplicação."
```

```bash
chmod +x test-api.sh
./test-api.sh
```

---

## FASE 2: TESTES DE VOLUME DE DADOS

### Objetivo

Se seu projeto processa dados (geolocação, eventos, observabilidade), medir quantos consegue processar.

### Passo 2.1: Teste de ingestão (Redis Streams ou queue)

**Exemplo para Hermes (se processa eventos):**

```javascript
// test-ingest.js
const redis = require('redis');
const client = redis.createClient({
  host: process.env.REDIS_HOST,
  port: process.env.REDIS_PORT
});

async function testIngest() {
  const streamKey = 'hermes:events';
  const testSize = 10000; // 10k eventos
  
  console.time('Ingest 10k events');
  
  for (let i = 0; i < testSize; i++) {
    await client.xAdd(streamKey, '*', {
      timestamp: new Date().toISOString(),
      level: 'info',
      message: `Test event ${i}`,
      service: 'test-service'
    });
  }
  
  console.timeEnd('Ingest 10k events');
  
  // Contar eventos
  const count = await client.xLen(streamKey);
  console.log(`Total events in stream: ${count}`);
  
  client.quit();
}

testIngest().catch(console.error);
```

**Rodar:**

```bash
node test-ingest.js
```

**Output esperado:**

```
Ingest 10k events: 2345.67ms
Total events in stream: 10000
```

**Métrica extraída:** "10,000 eventos ingeridos em 2.3 segundos = ~4,300 eventos/segundo"

---

### Passo 2.2: Teste de processamento (query/agregação)

**Exemplo para TimescaleDB (se usa):**

```javascript
// test-query.js
const { Pool } = require('pg');

const pool = new Pool({
  connectionString: process.env.DATABASE_URL
});

async function testQuery() {
  console.time('Query 1M rows aggregation');
  
  const result = await pool.query(`
    SELECT 
      DATE_TRUNC('minute', timestamp) as minute,
      COUNT(*) as count,
      AVG(duration_ms) as avg_duration
    FROM events
    WHERE timestamp > NOW() - INTERVAL '24 hours'
    GROUP BY minute
    ORDER BY minute DESC
  `);
  
  console.timeEnd('Query 1M rows aggregation');
  console.log(`Rows returned: ${result.rows.length}`);
  
  pool.end();
}

testQuery().catch(console.error);
```

**Métrica extraída:** "Agregação de 1M de registros em Xms"

---

## FASE 3: TESTES DE MEMÓRIA & CPU

### Passo 3.1: Profile de memória (Node.js)

```bash
# Rodar aplicação com profiling
node --max-old-space-size=4096 seu-app.js &
APP_PID=$!

# Deixar rodando por 2 minutos
sleep 120

# Capturar estado de memória
ps aux | grep seu-app

# Parar
kill $APP_PID
```

**Output esperado:**

```
usuario  12345  0.5  2.3  1234567 456789  ...
           ↑    ↑    ↑     ↑       ↑
         CPU%  MEM%  VSZ   RSS (memória real em KB)
```

**Interpretação:**

- RSS = ~450MB = "Sistema usa ~450MB de memória em repouso"

### Passo 3.2: Teste de memória sob carga

```bash
# Terminal 1: rodar app
npm start

# Terminal 2: monitorar
watch -n 1 'ps aux | grep seu-app'

# Terminal 3: gerar carga (vegeta de novo)
cat requisicoes.txt | vegeta attack -duration=60s -rate=500 | vegeta report
```

**Capture:**

- Memória inicial: X MB
- Memória sob carga (500 req/s): Y MB
- Memória após carga (repouso): Z MB

**Métrica:** "Sistema estabiliza em ~600MB com 500 req/s, sem memory leaks"

---

## FASE 4: TESTE DE ENDPOINTS ESPECÍFICOS

### Para Hermes (exemplo)

```bash
# 1. Ingestão de eventos
ab -n 1000 -p event-payload.json \
  -T "application/json" \
  https://seu-hermes.com/api/events/ingest

# 2. Query de dashboard
ab -n 500 \
  "https://seu-hermes.com/api/metrics?service=app&range=1h"

# 3. Search de logs
ab -n 300 \
  "https://seu-hermes.com/api/search?q=error&limit=100"
```

**Salve resultados:**

```bash
ab -n 1000 -p event.json -T "application/json" https://seu-hermes.com/api/events > test-ingest-results.txt
```

---

## FASE 5: COMPILAR RESULTADOS

### Template de extração

```markdown
## Performance Metrics (Verificado via Load Test)

### Throughput
- **Ingestão de eventos:** X eventos/segundo
- **API de queries:** Y requisições/segundo  
- **Dashboard:** Z requisições/segundo

### Latência
- **P50:** A ms
- **P95:** B ms
- **P99:** C ms
- **Max:** D ms

### Confiabilidade
- **Success rate:** 99.X%
- **Failed requests:** 0
- **Error rate:** <0.5%

### Recursos
- **Memory (repouso):** X MB
- **Memory (500 req/s):** Y MB
- **CPU:** Z%

### Teste de Volume
- **10k eventos ingeridos:** X segundos
- **1M registros processados:** Y ms

### Capturado em
- Data: 2026-01-15
- Ferramenta: Apache Bench + Vegeta
- Ambiente: Production (Render/Vercel deploy)
- Comando: `ab -n 5000 -c 50 https://...`
```

---

## FASE 6: DOCUMENTAÇÃO NO README

### Formato para colocar no GitHub

````markdown
# Hermes — Observable Event Stream

## Performance

Tested via Apache Bench and Vegeta load testing (Jan 2026).

| Métrica | Valor | Teste |
|---------|-------|-------|
| Throughput | 5,200 req/s | `ab -n 5000 -c 50` |
| P99 Latency | 89 ms | Vegeta 100 req/s/30s |
| Memory (idle) | 340 MB | Node.js process |
| Memory (500 req/s) | 420 MB | Sustained 1 min |
| Event ingest | 8,500 events/s | Redis Streams |
| Query latency (p99) | 145 ms | TimescaleDB |

### Como reproduzir testes

```bash
# Setup
npm install
npm start

# Terminal 2: Load test
ab -n 5000 -c 50 http://localhost:3000/api/health
````

Ver test-results.txt para output completo.

```

---

## FASE 7: ROTEIRO COMPLETO (Dia a Dia)

### Dia 1: Preparar ambiente
- [ ] Deploy no Render/Vercel
- [ ] Testar acesso via URL pública
- [ ] Verificar logs funcionando
- [ ] Instalar Apache Bench + Vegeta localmente

### Dia 2: Teste básico
- [ ] Rodar 1 teste com Apache Bench (1000 requisições)
- [ ] Capturar output em arquivo
- [ ] Rodar Vegeta por 30s (100 req/s)
- [ ] Salvar resultados

### Dia 3: Teste de volume
- [ ] Teste de ingestão de dados (10k registros)
- [ ] Teste de query/agregação
- [ ] Teste de memória

### Dia 4: Compilação
- [ ] Organizar resultados em tabela
- [ ] Redigir parágrafo para README
- [ ] Adicionar comandos reproduzíveis

### Dia 5: Validação
- [ ] Revisar cada número (é honesto?)
- [ ] Re-rodar 1 teste para confirmar
- [ ] Atualizar CV/Portfolio com métricas

---

## REGRA DE OURO

**✅ Coloque no README:**
```

Latency: 89ms (p99) — Apache Bench, 5000 req, 50 concurrent

```

**❌ NÃO coloque:**
```

Pode processar milhões de eventos (sem testar) Usado em produção por X empresas (mentira)

```

**Se o teste não rodou, o número não entra.**

---

## Dúvidas Comuns

**P: E se meu projeto for lento?**
A: Melhor ter 500 req/s verificado do que 10k fabricado. Recruiter vê honestidade.

**P: Preciso de X requisições/segundo para parecer bom?**
A: Não. A métrica prova que você sabe como medir performance. Qualidade > quantidade.

**P: E se o projeto cair durante o teste?**
A: Ótimo feedback! Agora você sabe que precisa de scaling/cache. Isso também é aprendizado válido.

---

## Próximos Passos para Francisco

1. Qual projeto tem deploy pronto? (Hermes ou Sísifo)
2. Qual é a URL?
3. Qual endpoint você quer testar primeiro?

Vamos começar com UM teste simples e depois escalonamos.
```