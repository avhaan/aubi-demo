# aubi-demo

A small Go service used as the **target repository** for
[AUBI](https://github.com/Neilyoo98/GDSC-Hackathon) — an AI system that routes
GitHub issues to the developer who owns the affected code, proposes a fix,
verifies it against this test suite, and opens a pull request after a human
approves.

This repo is a fixture. It exists to be broken and fixed.

## Layout

| Package | What it does |
|---|---|
| `auth/` | Short-lived access tokens with an in-memory cache |
| `billing/` | Line items → invoice totals, in integer cents |
| `api/` | Pagination helpers shared by the HTTP handlers |
| `store/` | A small bounded cache with eviction |

## Running the tests

```bash
go test ./...
```

`main` currently has **one failing test** in `auth/` — that is deliberate, and
it is the issue AUBI is pointed at. Everything else is green.

## About the bugs

The defects here were planted on purpose, and they are the kind that survive
code review rather than the kind that are obvious:

- the failing one is a check that was never written, not a typo
- the latent ones are not covered by the current tests, which is exactly why
  they would survive in a real codebase

**This matters for interpreting AUBI's results.** Because the suite passes with
the latent bugs present, `go test ./...` proves a proposed fix *compiles and
breaks nothing* — it does not prove the fix is correct. Only the `auth/` case
has a test that actually fails first and passes after, which is the one place
verification is a real correctness signal.

A stronger evaluation uses a real open-source repository with real bug-fix
commits reverted, where the upstream test suite does cover the defect and the
maintainers' own fix is the ground truth.

## License

MIT
