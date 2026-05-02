# trello-cli

Go CLI over the Trello REST API. Resource commands are auto-generated
from `openapi.json`, giving 100% endpoint coverage out of the box.
Designed as the transport layer for AI agents (e.g. OpenClaw) to talk
to Trello, plus a daily-driver CLI for humans.

- **Human guide:** [`docs/USAGE.md`](docs/USAGE.md) — install, auth,
  workflows, troubleshooting.
- **Full command catalog:** [`docs/COMMANDS.md`](docs/COMMANDS.md) —
  every operation, every flag (auto-generated from the spec).
- **Agent skill (OpenClaw / AgentSkills):**
  [`skills/trello-cli/SKILL.md`](skills/trello-cli/SKILL.md).

## Install

Pick whichever fits your environment.

### Homebrew (macOS / Linux)

```bash
brew install mrfansi/tap/trello-cli
```

### `go install`

Requires Go ≥ the version pinned in [`go.mod`](go.mod).

```bash
go install github.com/mrfansi/trello-cli/cmd/trello-cli@latest
```

The binary lands in `$GOBIN` (or `$GOPATH/bin`). Add that to `PATH` if
it isn't already.

### Pre-built binaries

Download the archive for your platform from the
[GitHub Releases page](https://github.com/mrfansi/trello-cli/releases)
and drop the `trello-cli` binary on your `PATH`. Releases include
checksums.

### From source

```bash
git clone https://github.com/mrfansi/trello-cli.git
cd trello-cli
make            # builds ./bin/trello-cli
make install    # or: install into $GOBIN
```

`make` is the one-liner. It embeds `git describe`-derived version
metadata via `-ldflags`. Run `make help` for every available target.

### Verify

```bash
trello-cli --version
trello-cli me        # exercises auth (see below)
```

## Auth

Get an API key + token at <https://trello.com/app-key>.

```bash
export TRELLO_API_KEY=...
export TRELLO_TOKEN=...
```

Or write `~/.trello-cli/config.yaml`:

```yaml
api_key: ...
token: ...
```

## Usage

Command shape: `trello-cli <resource> <operation> [args] [flags]`.
Operation names mirror the OpenAPI `operationId` (kebab-case
`<method>-<path>`).

```bash
trello-cli me                                          # auth check (alias)
trello-cli boards get-boards-id <board-id>             # GET /boards/{id}
trello-cli boards get-boards-id-labels <board-id>      # GET /boards/{id}/labels
trello-cli boards post-boards --name "New" --idOrganization <org-id>
trello-cli cards post-cards --idList <list-id> --name "Task"
trello-cli cards put-cards-id <card-id> --data '{"name":"Renamed"}'
trello-cli cards delete-cards-id <card-id>
trello-cli lists get-lists-id-cards <list-id>
trello-cli labels get-labels-id <label-id>
trello-cli members get-members-id me
trello-cli search get-search --query "term" --modelTypes cards
```

Top-level groups: `actions`, `applications`, `batch`, `boards`,
`cards`, `checklists`, `customFields`, `emoji`, `enterprises`,
`labels`, `lists`, `members`, `notifications`, `organizations`,
`plugins`, `search`, `tokens`, `webhooks`.

`trello-cli <group> --help` lists every operation in that group.
Every operation `--help` lists path args, query flags, and (for
mutating endpoints) the `--data` body flag.

### Raw passthrough

For ad-hoc requests or quick experimentation:

```bash
trello-cli raw GET /members/me
trello-cli raw GET /boards/{id}/labels --path id=abc --query limit=10
trello-cli raw POST /cards --query idList=xyz --query name="New"
trello-cli raw PUT /cards/{id} --path id=abc --data @body.json
trello-cli raw DELETE /cards/{id} --path id=abc
```

Flags: `--path key=value`, `--query key=value`, `--header key=value`,
`--data <json|@file>`. Auth is auto-injected. Output is the raw
response body; non-2xx prints status to stderr and exits 1.

### Output

All resource and `raw` commands emit raw JSON. Pipe through `jq` for
filtering.

## Development

```bash
make             # default goal: builds ./bin/trello-cli with embedded version
make help        # list all targets
make test        # go test -race -cover ./...
make ci          # vet + test (used in CI)
make fmt         # gofmt -s -w .
make lint        # golangci-lint (auto-installs if needed via `make tools`)
make gen         # regenerate Trello client from openapi.json
make gen-cmds    # regenerate cobra commands + docs/COMMANDS.md from openapi.json
make snapshot    # local goreleaser snapshot (requires goreleaser)
make clean       # remove ./bin and ./dist
```

### Releasing

Tag a version on `main`:

```bash
scripts/release.sh v0.1.0
```

The script enforces a clean tree, the `main` branch, and an
unused tag, then pushes the tag. The
[release workflow](.github/workflows/release.yml) runs `goreleaser`
which builds linux/darwin/windows × amd64/arm64 archives, checksums,
and a GitHub Release.

## Testing

```bash
go test -race -cover ./internal/... ./tools/...
```

Logic-bearing packages (`config`, `output`, `cmdutil`, `commands`,
`tools/dedup`) have unit coverage. Generated command files are
mechanical wrappers around `auto.execRaw`; coverage of `execRaw`
(via `commands.rawCmd`) exercises the same code path.

## Regeneration pipelines

Two independent generators run from `openapi.json`:

1. **Typed Go client** (`make gen`) — runs `oapi-codegen` then
   `tools/dedup` to remove duplicate type aliases and replace
   anonymous-union path params with plain `string`.
2. **Cobra commands** (`make gen-cmds`) — runs `tools/cmdgen` which
   walks every `paths.<path>.<method>` and emits one cobra subcommand
   per operation into `internal/commands/auto/`.

The typed client is currently retained for callers that want
strongly-typed access; the CLI itself dispatches every request via
the raw HTTP path so coverage stays in lockstep with the spec.

## Layout

```
cmd/trello-cli/             # main entrypoint
internal/commands/          # cobra root + raw + me alias
internal/commands/auto/     # generated resource subcommand groups
internal/client/            # auth-injecting HTTP factory
internal/config/            # viper-backed credential loader
internal/cmdutil/           # shared command helpers (context, decode)
internal/output/            # JSON / table renderer (used by `me`)
internal/version/           # build-time version metadata (-ldflags)
internal/trello/            # typed generated client
tools/dedup/                # client codegen post-processor
tools/cmdgen/               # cobra command + docs generator
docs/USAGE.md               # human usage guide
docs/COMMANDS.md            # full command reference (auto-generated)
skills/trello-cli/          # OpenClaw / AgentSkills SKILL.md
scripts/release.sh          # release tag helper
.goreleaser.yaml            # multi-platform release config
.github/workflows/          # CI + release pipelines
openapi.json                # Trello OpenAPI spec
```
