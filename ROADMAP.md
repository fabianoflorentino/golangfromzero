# Golang Roadmap – Evolução do Projeto golangfromzero

Este roadmap guia a evolução técnica do projeto e o aprofundamento em Golang,
com foco em código idiomático, backend profissional e arquitetura sustentável.

---

## 📍 Fase 1 – Consolidar Fundamentos em Go (estado atual)

🎯 **Objetivo:** fortalecer o uso correto dos fundamentos já presentes no projeto.

### Estudar e revisar

- Structs e métodos
- Organização de pacotes
- Erros como valores (`errors.Is`, `%w`)
- Uso consistente de `context.Context`
- `defer` e ciclo de vida de recursos
- Logging com `slog`

### Ações no projeto

- Revisar propagação de `context`
- Padronizar tratamento de erros HTTP
- Remover código duplicado em handlers

📌 **Critério de conclusão**

- Código fácil de seguir sem comentários excessivos
- Fluxo de erro explícito e previsível

---

## 📍 Fase 2 – Go Idiomático e Design Simples

🎯 **Objetivo:** escrever código que “parece Go”, não Ruby ou Node.

### Estudar

- Interfaces pequenas
- Aceitar interfaces, retornar structs
- Composição > herança
- Evitar abstrações prematuras
- Nomeação idiomática de pacotes e funções

### Ações no projeto

- Avaliar:
  - quais interfaces realmente fazem sentido
  - quais podem ser removidas
- Reduzir dependência de pacotes genéricos (`helper`)
- Aproximar funções de onde são usadas

📌 **Critério de conclusão**

- Menos camadas
- Menos arquivos “utilitários”
- Código mais direto

---

## 📍 Fase 3 – Testes de Verdade em Go

🎯 **Objetivo:** aumentar confiança sem criar complexidade.

### Estudar

- `testing` package
- Table-driven tests
- Subtests
- Testes sem mocks pesados
- Fakes simples

### Ações no projeto

- Criar testes para:
  - handlers HTTP
  - regras de validação
- Testar erros explicitamente
- Evitar mocks de infra quando possível

📌 **Critério de conclusão**

- Testes legíveis
- Testes explicam o comportamento do sistema

---

## 📍 Fase 4 – Separação de Responsabilidades (Use Cases)

🎯 **Objetivo:** evitar lógica de negócio dentro de controllers.

### Estudar

- Application services / use cases
- Ports & adapters (hexagonal em Go)
- Dependências fluindo para dentro

### Ações no projeto

- Introduzir camada de `usecase`
- Controllers apenas:
  - recebem request
  - chamam use case
  - retornam response
- Repositórios usados via abstrações mínimas

📌 **Critério de conclusão**
- Controller não conhece regras de negócio
- Domínio não conhece HTTP nem banco

---

## 📍 Fase 5 – Backend Pronto para Produção

🎯 **Objetivo:** tornar o serviço resiliente e observável.

### Estudar

- Middlewares HTTP
- Graceful shutdown
- Healthchecks
- Configuração via env
- Timeouts e cancelamento

### Ações no projeto

- Garantir shutdown limpo do servidor
- Padronizar timeouts
- Melhorar respostas de erro
- Logs estruturados em pontos críticos

📌 **Critério de conclusão**

- Serviço pode rodar em produção sem surpresas

---

## 📍 Fase 6 – Concorrência e Performance (Avançado)

🎯 **Objetivo:** entender o comportamento do Go em runtime.

### Estudar

- Goroutines
- Channels e `select`
- Mutex vs channels
- Garbage Collector
- Escape analysis
- `pprof`
- Benchmarks

### Ações no projeto

- Introduzir concorrência apenas onde faz sentido
- Medir antes de otimizar
- Criar benchmarks simples

📌 **Critério de domínio**

- Decisões baseadas em métricas, não achismo

---

## 🧠 Princípios que guiam o projeto

- Simplicidade > abstração
- Clareza > flexibilidade
- Código explícito > código inteligente
- Resolver bem o problema atual

---

## ✅ Sinais de domínio em Go

- Você remove mais código do que adiciona
- Interfaces existem por necessidade real
- O código é fácil de explicar
- Testes e compilador fazem o trabalho pesado

---
