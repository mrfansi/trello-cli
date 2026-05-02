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

```bash
make install
# or
go install github.com/mrfansi/trello-cli/cmd/trello-cli@latest
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
make build       # compile binary
make test        # go test -race -cover ./...
make vet         # go vet ./...
make gen         # regenerate Trello client from openapi.json
make gen-cmds    # regenerate cobra commands from openapi.json
```

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
internal/trello/            # typed generated client
tools/dedup/                # client codegen post-processor
tools/cmdgen/               # cobra command generator
openapi.json                # Trello OpenAPI spec
```
