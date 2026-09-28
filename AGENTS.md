# AGENTS.md — plugin-authoring

Standalone plugin repo for the `charly box` authoring verbs
(`command:*:box`). The plugin is a Go module at `candy/plugin-authoring/`
(module path `github.com/opencharly/plugin-authoring/candy/plugin-authoring`);
the root `charly.yml` only declares `discover: candy` so the repo is a project
and its candy is scanned.

Canonical files:

- `candy/plugin-authoring/charly.yml` — the `plugin-authoring:` candy entity
  and the embedded `authoring-skill:` skill entity.
- `candy/plugin-authoring/` — the Go source: `plugin.go`, `authoring.go`,
  `authoring_edit.go`, `cmd/serve/main.go`.
- `.github/workflows/tag-on-merge.yml` — CalVer tag + `CHANGELOG/` on merge.
- `README.md` — user overview only; never agent guidance.

## Load these skills first (R0)

- `/charly-authoring:authoring` — the box-authoring verbs reference (projected
  from the embedded `authoring-skill:` entity). Load before changing a verb.
- `/charly-image:image` — box composition and `charly.yml` shapes the verbs
  mutate.
- `/charly-internals:plugin` — the plugin authoring reference: the `plugin:`
  block, the unified Provider model, nested command parents. Load before
  touching the provider.
- `/charly-internals:git-workflow` — before any git/PR action.

## Build / validate / test

- `go build ./...` in `candy/plugin-authoring/` — compile the plugin module.
- `go test ./...` in `candy/plugin-authoring/` — the plugin's Go tests.
- `charly box validate` at the repo root — the structural check (the candy +
  `plugin:` block).
- The merge gate is the **org-wide** `charly/pr-validator` (required check
  `validate / validate`, defined in `opencharly/.github`); this repo has **no**
  per-repo candy gate.
- The live R10 witness is `check-commands-local` in `opencharly/charly` (the
  full `charly box set/add-candy/rm-candy/write/cat/fetch/refresh` end-to-end).

## Modify this repo

- These are **nested command** providers (`command:<word>:box`); keep the parent
  identity on the wire when adding a verb.
- The `authoring-skill:` entity is the projected source for
  `/charly-authoring:authoring`; keep it in step with any verb change.

## Landing

- PR-only. Every change lands through a pull request; the org-required
  `charly/pr-validator` validates the diff and body and arms native auto-merge on
  PASS. Direct pushes to `main` are blocked.
- History lives in `CHANGELOG/` (written by `tag-on-merge` at merge time); the PR
  body IS the changelog.
- The authoritative rulebook is the umbrella `AGENTS.md` in
  `opencharly/opencharly` and `charly/AGENTS.md` in the charly repo. Do not
  restate its rules here.
