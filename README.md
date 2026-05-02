# trello-cli

Curated Go CLI over the Trello REST API.

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

```bash
trello-cli me                                  # show authenticated user
trello-cli board ls                            # list your boards
trello-cli board get <board-id>
trello-cli board create "My Board"

trello-cli list ls --board <board-id>
trello-cli list create "Backlog" --board <board-id>
trello-cli list archive <list-id>

trello-cli card ls --list <list-id>
trello-cli card ls --board <board-id>
trello-cli card create "Task" --list <list-id> --desc "details"
trello-cli card update <card-id> --name "New name" --closed
trello-cli card rm <card-id>

trello-cli checklist ls --card <card-id>
trello-cli checklist create "Steps" --card <card-id>

trello-cli member get <username>
```

Add `--json` to any command for raw JSON output.

## Development

```bash
make build      # compile binary
make test       # go test -race -cover ./...
make vet        # go vet ./...
make gen        # regenerate Trello client from openapi.json
```

## Testing

```bash
go test -race -cover ./internal/... ./tools/...
```

Logic-bearing packages (`config`, `output`, `cmdutil`, `tools/dedup`)
have unit coverage. Cobra command wiring and the auth-injecting HTTP
factory are exercised via the live smoke flow (`trello-cli me`,
`board ls`, etc.) once credentials are configured.

## Regenerate client

```bash
make gen
```

The generator runs `oapi-codegen` then a small post-processor (`tools/dedup`)
that:

- removes duplicate type declarations from clashing operationIds in the spec
- replaces anonymous-union path params (`struct { union json.RawMessage }`)
  with plain `string` so endpoints like `GetMembersId` are callable.

Without that pass, the generated code does not compile and the affected
endpoints serialize empty path segments.

## Layout

```
cmd/trello-cli/      # main entrypoint
internal/commands/   # cobra subcommands (board, list, card, checklist, member, me)
internal/client/     # auth-injecting wrapper around generated client
internal/config/     # viper-backed config (env + ~/.trello-cli/config.yaml)
internal/cmdutil/    # shared command helpers (context, decode)
internal/output/     # table + json renderer
internal/trello/     # generated client (do not hand-edit)
tools/dedup/         # codegen post-processor
openapi.json         # Trello OpenAPI spec
```
