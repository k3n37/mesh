# mesh

Small, serious experiments for queues, retries, idempotency, and coordination in distributed systems.

## Purpose

Explore the operational behaviors that show up once systems cross process and network boundaries.

## Role in the ecosystem

- Applies theory from `foundry`
- Supports architecture decisions in `orbit`
- Complements `atlas` with runnable code

## Status

Starter lab with a simple idempotent in-memory queue worker and roadmap for deeper experiments.

## Tech stack

- Go
- Markdown

## Structure

```text
mesh/
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

- `atlas`
- `nimbus`
- `pulse`

## Future direction

Add bounded retries, backoff, leader election notes, and consistency experiments without turning the repo into a framework.
