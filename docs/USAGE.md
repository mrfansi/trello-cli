# trello-cli usage guide

A Go CLI over the Trello REST API. Resource commands are auto-generated
from `openapi.json` (255 operations across 18 groups), giving 100%
endpoint coverage. Designed as a transport layer for AI agents (e.g.,
OpenClaw) to interact with Trello, plus a daily-driver CLI for humans.

## Install

```bash
make install
# or
go install github.com/mrfansi/trello-cli/cmd/trello-cli@latest
```

The binary is named `trello-cli` and lands on `$GOBIN` (or
`$GOPATH/bin`). Confirm it is on your `PATH`:

```bash
trello-cli --help
```

## Authenticate

Trello requires an **API key** + **token** pair. Both go on the query
string of every request — `trello-cli` injects them automatically.

1. Visit <https://trello.com/app-key> and copy the **Key**.
2. On the same page, click "Token" to authorize the app and copy the
   generated token.

Choose one of:

### Option A — environment variables (transient)

```bash
export TRELLO_API_KEY=...
export TRELLO_TOKEN=...
```

### Option B — config file (persistent)

```bash
mkdir -p ~/.trello-cli
cat > ~/.trello-cli/config.yaml <<'EOF'
api_key: ...
token: ...
EOF
chmod 600 ~/.trello-cli/config.yaml
```

The config file path is fixed at `~/.trello-cli/config.yaml`. Env vars
take precedence when both are set.

### Verify auth

```bash
trello-cli me
```

Returns the authenticated user's profile as JSON. Non-2xx responses go
to stderr and exit 1.

## Command shape

```text
trello-cli <group> <operation> [path-args...] [flags]
trello-cli raw <METHOD> <PATH> [flags]
trello-cli me
```

- **`<group>`** — resource family from the OpenAPI path's first
  segment (e.g., `boards`, `cards`, `members`). 18 total.
- **`<operation>`** — kebab-case `<method>-<path>` mirroring the
  OpenAPI `operationId` (e.g., `get-boards-id`, `post-cards`).
- **Path args** — positional, in the order they appear in the path
  template (e.g., `<id>` for `/boards/{id}`).
- **Query flags** — one `--<name>` per query parameter declared in
  the spec (e.g., `--fields`, `--filter`).
- **Body** — for endpoints with a request body, pass `--data <json>`
  or `--data @file.json`.

## All output is JSON

Every resource command, plus `raw`, prints the raw response body to
stdout. Non-2xx responses go to stderr (exit 1). Pipe through `jq` for
filtering:

```bash
trello-cli boards get-boards-id <id> --fields name,url | jq .name
trello-cli members get-members-id me | jq -r '.idBoards[]'
```

## Discovery

Each cobra layer exposes its own help:

```bash
trello-cli --help                              # 18 groups + me + raw
trello-cli boards --help                       # all board operations
trello-cli boards get-boards-id --help         # path args, query flags, body
```

For a flat reference of every operation, see
[`docs/COMMANDS.md`](./COMMANDS.md). It is auto-generated from the spec
and refreshes via `make gen-cmds`.

## Common workflows

Concrete recipes for typical Trello tasks. Replace `<board-id>` /
`<list-id>` / `<card-id>` with real IDs.

### Identify the authenticated user

```bash
trello-cli me | jq '{id, username, fullName}'
```

### List boards

```bash
trello-cli members get-members-id me \
  | jq -r '.idBoards[]' \
  | while read id; do
      trello-cli boards get-boards-id "$id" --fields name,url \
        | jq -r '[.id, .name, .url] | @tsv'
    done
```

Or in one call (most boards are returned in the member object):

```bash
trello-cli members get-members-id-boards me --fields name,url
```

### Get a specific board

```bash
trello-cli boards get-boards-id <board-id> --fields name,url,desc
```

### List a board's lists

```bash
trello-cli boards get-boards-id-lists <board-id> --fields name,pos,closed
```

### List cards in a list

```bash
trello-cli lists get-lists-id-cards <list-id> --fields name,desc,due,idMembers
```

### Create a card

```bash
trello-cli cards post-cards \
  --idList <list-id> \
  --name "New task" \
  --desc "details here"
```

For complex bodies, use `--data`:

```bash
trello-cli cards post-cards --data '{"idList":"<list-id>","name":"Task","desc":"..."}'
```

### Update / move / archive a card

```bash
# Rename
trello-cli cards put-cards-id <card-id> --data '{"name":"Renamed"}'

# Move to another list
trello-cli cards put-cards-id <card-id> --data '{"idList":"<other-list-id>"}'

# Archive
trello-cli cards put-cards-id <card-id> --data '{"closed":true}'
```

### Delete a card

```bash
trello-cli cards delete-cards-id <card-id>
```

### Add a checklist to a card

```bash
trello-cli checklists post-checklists \
  --data '{"idCard":"<card-id>","name":"Steps"}'
```

### Add labels to a board

```bash
trello-cli labels post-labels \
  --data '{"idBoard":"<board-id>","name":"Bug","color":"red"}'
```

### Search across the workspace

```bash
trello-cli search get-search --query "term" --modelTypes cards,boards
```

### Webhooks

```bash
# Create a webhook subscribing to a model (board / card / list / etc.)
trello-cli webhooks post-webhooks \
  --data '{"idModel":"<model-id>","callbackURL":"https://example.com/hook"}'

# List webhooks for a token
trello-cli tokens get-tokens-token-webhooks <your-token>
```

## Raw passthrough

For ad-hoc requests, prototyping, or endpoints with awkward
codegen names:

```bash
trello-cli raw GET /members/me
trello-cli raw GET /boards/{id}/labels --path id=abc --query limit=10
trello-cli raw POST /cards --query idList=xyz --query name="Task"
trello-cli raw PUT /cards/{id} --path id=abc --data @body.json
trello-cli raw DELETE /cards/{id} --path id=abc
```

Flags:

- `--path key=value` (repeatable) — substitutes `{key}` in the path
- `--query key=value` (repeatable) — appends to the query string
- `--header key=value` (repeatable) — extra request headers
- `--data <json|@file>` — JSON body literal or `@` + file path
- `-X, --method <METHOD>` — alternative to the positional method arg

Auth is auto-injected on the query string. Output is the raw response
body.

## Error handling

| Symptom | Cause | Fix |
|---------|-------|-----|
| `missing credentials: set TRELLO_API_KEY and TRELLO_TOKEN env vars or write ~/.trello-cli/config.yaml` | Neither env vars nor config file are set | Configure auth (above) |
| `trello api 401 Unauthorized: invalid token` | Token revoked, expired, or wrong key/token pair | Regenerate at <https://trello.com/app-key> |
| `trello api 404 Not Found` | Resource ID does not exist or you lack access | Verify the ID and your membership |
| `trello api 400 Bad Request: invalid id` | Path arg is not a valid Trello ID format | Pass a 24-character hex ID or `me` |
| `path placeholder {x} not found` | `--path x=...` for a key not present in the path template | Drop the flag or use the right key |
| `non-2xx response` (exit 1) | Any 3xx/4xx/5xx | Stderr already prints status + body |
| Hang for ~30s then fail | Network or timeout | Each request has a 30 s deadline; check connectivity |

## Regeneration

Two independent generators run from `openapi.json`:

```bash
make gen        # typed Go client (oapi-codegen + tools/dedup)
make gen-cmds   # cobra commands (tools/cmdgen) + docs/COMMANDS.md
```

Run both after editing the spec. Patches to the spec live in
`openapi.json`; document why each patch exists in the commit message.

## Design notes

- **Auth on the query string** — Trello uses `?key=...&token=...`,
  not `Authorization` headers. The CLI injects via a request editor on
  every outbound HTTP request.
- **One JSON in, one JSON out** — every command emits the raw response
  body; combine with `jq` for filtering. No bespoke output formatting
  per resource.
- **Codegen is mechanical** — `tools/cmdgen` walks the spec and emits
  one cobra subcommand per `paths.<path>.<method>`. Subcommand names
  mirror `operationId`. No hand-written ergonomics layer; predictable
  rather than pretty.
- **Typed client retained** — `internal/trello/client.gen.go` is still
  generated for callers that want strongly-typed access from Go code.
  The CLI itself dispatches via the raw HTTP path so coverage stays in
  lockstep with the spec.

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
tools/cmdgen/               # cobra command + docs generator
docs/USAGE.md               # this file
docs/COMMANDS.md            # full command reference (auto-generated)
skills/trello-cli/          # OpenClaw / AgentSkills SKILL.md
openapi.json                # Trello OpenAPI spec
```
