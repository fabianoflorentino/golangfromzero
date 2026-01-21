# ✅ MVP Production Checklist

> **Goal:** Put an API safely into production with the *minimum required* engineering.
>
> This checklist is intentionally **small, practical, and risk‑driven**.
> If all **MVP items** are done, the service can go to production.

---

## 🎯 Definition of “Production‑Ready (MVP)”

An API is **production‑ready** when:

* It **does not crash** under normal load
* It **shuts down safely** during deploys
* It **does not leak internal errors** to clients
* It is **observable enough to debug incidents**
* It can be **monitored by infrastructure (LB / K8s)**

Anything beyond this is **iteration**, not a blocker.

---

## 🔴 MVP — REQUIRED BEFORE PRODUCTION

### 1. Database Connection Pool

**Why:** Prevents DB exhaustion and improves performance.

**Must have:**

* Single pool created at startup
* Pool injected into handlers / repositories
* Reasonable limits

**Checklist:**

* [ ] Pool created once on startup
* [ ] Max connections configured
* [ ] Min idle connections configured
* [ ] Pool closed on shutdown

---

### 2. Graceful Shutdown

**Why:** Prevents broken requests and corrupted state during deploys.

**Must have:**

* SIGINT / SIGTERM handling
* HTTP server shutdown with timeout
* DB pool closed after requests finish

**Checklist:**

* [ ] OS signal handling
* [ ] `http.Server.Shutdown(ctx)`
* [ ] Configurable shutdown timeout (10–30s)

---

### 3. Structured Logging

**Why:** Without logs, incidents are guesswork.

**Must have:**

* Structured logs (JSON)
* Log level control via env
* Request context propagation

**Checklist:**

* [ ] Structured logger (`slog`, `zerolog`, etc.)
* [ ] Log request start + end
* [ ] Log errors with context (server‑side only)
* [ ] Log level configurable

---

### 4. Safe Error Handling

**Why:** Prevents security leaks and improves client experience.

**Rules:**

* Internal errors → logs
* Client errors → clean messages
* Never expose stack traces

**Checklist:**

* [ ] Typed / sentinel errors
* [ ] HTTP status mapped explicitly
* [ ] No `err.Error()` leaked for 500s
* [ ] No string comparison on errors

---

### 5. Health & Readiness Endpoints

**Why:** Required by load balancers and orchestrators.

**Endpoints:**

* `/health` → liveness
* `/ready` → readiness (DB reachable)

**Checklist:**

* [ ] `/health` returns 200 if process is alive
* [ ] `/ready` checks DB connectivity
* [ ] Fast response (< 100ms)

---

## 🟡 CONTEXTUAL — DEPENDS ON API TYPE

### 🔐 API Type Matrix

| Feature             | Internal API    | Public API   |
| ------------------- | --------------- | ------------ |
| Rate limiting       | Optional        | **Required** |
| Auth                | Network / token | **Required** |
| Metrics             | Recommended     | **Required** |
| CORS                | Optional        | **Required** |
| Detailed validation | Basic           | **Strict**   |

---

### 6. Rate Limiting

**Why:** Protects the service from abuse.

**Checklist:**

* [ ] Per‑IP or per‑token limits
* [ ] Configurable via env
* [ ] HTTP 429 on limit exceeded

> **Public API:** REQUIRED
> **Internal API:** Optional

---

### 7. Basic Metrics (Minimum Observability)

**Why:** You can’t fix what you can’t see.

**Minimum metrics:**

* Request count
* Request duration
* Error count

**Checklist:**

* [ ] `/metrics` endpoint (Prometheus)
* [ ] Request duration histogram
* [ ] Error counter

> **Public API:** REQUIRED
> **Internal API:** Strongly recommended

---

### 8. Input Validation

**Why:** Prevents bad data and unnecessary load.

**Checklist:**

* [ ] Validate query params
* [ ] Validate body fields
* [ ] Reject malformed UUIDs
* [ ] Return clear 400 errors

---

### 9. Centralized Configuration

**Why:** Enables safe deploys and reproducibility.

**Checklist:**

* [ ] All config via environment variables
* [ ] Startup validation
* [ ] `.env.example` documented

---

### 10. Security Middleware (Baseline)

**Checklist:**

* [ ] Panic recovery
* [ ] Request timeout
* [ ] Request size limit
* [ ] Basic security headers

> **Public API:** REQUIRED
> **Internal API:** Recommended

---

## 🟢 POST‑MVP — ITERATE AFTER GO‑LIVE

These improve quality but **must not block production**.

* [ ] Unit tests (>50%)
* [ ] Integration tests
* [ ] OpenAPI / Swagger docs
* [ ] Dockerfile (multi‑stage)
* [ ] Database migrations

---

## 🔵 ADVANCED — ONLY WHEN YOU FEEL REAL PAIN

Do **not** implement these preemptively.

* [ ] Circuit breaker
* [ ] Automatic retries
* [ ] Distributed tracing
* [ ] 90%+ test coverage
* [ ] Complex CI/CD pipelines

> Rule: **Add only when a real incident justifies it.**

---

## 🧠 SRE Rule of Thumb

> **Production readiness is about risk, not perfection.**

If you have:

* Pooling
* Shutdown
* Logs
* Health
* Error boundaries

➡️ You are **ready to ship**.

---
