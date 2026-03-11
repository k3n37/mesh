# distributed-systems-lab

Small, serious experiments for queues, retries, idempotency, and coordination in distributed systems.

## Purpose

Explore the operational behaviors that show up once systems cross process and network boundaries.

## Role in the ecosystem

- Applies theory from `foundations-lab`
- Supports architecture decisions in `master-platform`
- Complements `system-design-lab` with runnable code

## Status

Starter lab with a simple idempotent in-memory queue worker and roadmap for deeper experiments.

## Tech stack

- Go
- Markdown

## Structure

```text
distributed-systems-lab/
├── cmd/
│   └── idempotent-worker/
│       └── main.go
├── docs/
│   └── experiments.md
├── .editorconfig
├── .gitignore
├── README.md
├── ROADMAP.md
└── go.mod
```

## Getting started

```bash
go run ./cmd/idempotent-worker
```

## Related repositories

- `system-design-lab`
- `infrastructure-platform`
- `streaming-analytics-engine`

## Future direction

Add bounded retries, backoff, leader election notes, and consistency experiments without turning the repo into a framework.
