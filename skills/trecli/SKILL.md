---
name: trecli
description: Read and write Trello (boards, lists, cards, members, checklists, labels, webhooks, search, ...) via the trecli binary. 100% Trello REST coverage. Use whenever the user mentions Trello.
homepage: https://github.com/mrfansi/trecli
metadata: {"openclaw":{"emoji":"📋","requires":{"bins":["trecli"],"env":["TRELLO_API_KEY","TRELLO_TOKEN"]},"primaryEnv":"TRELLO_TOKEN","install":[{"id":"go","kind":"go","module":"github.com/mrfansi/trecli/cmd/trecli","bins":["trecli"],"label":"Install trecli (go install)"}]}}
---

# trecli

Use `trecli` to read or write data in Trello. Every Trello REST
endpoint (255 operations) is reachable as a CLI command. Output is
always JSON on stdout; non-2xx responses go to stderr with exit 1.

## When to use this skill

Trigger this skill whenever the user asks you to:

- list / read / inspect a Trello board, list, card, member, checklist,
  label, organization, webhook, action, notification, or token
- create / update / move / archive / delete any of the above
- search Trello content
- subscribe / unsubscribe webhooks against Trello models
- look up identifiers (board ID, list ID, card ID, member ID)

Do **not** invent endpoints from memory. If the user describes a task
that does not match a known recipe below, run `trecli <group>
--help` to discover the operation, or fall back to `trecli raw`.

## Setup contract

This skill assumes:

- Binary `trecli` exists on `PATH` (`requires.bins` enforced).
- `TRELLO_API_KEY` and `TRELLO_TOKEN` are present in the environment
  (`requires.env` enforced). They are injected by OpenClaw at run time
  via `skills.entries.trecli.env` or via `apiKey` SecretRef.

If both checks fail, do **not** call the CLI; ask the user to install
the binary or configure credentials. Suggested fixes:

```bash
go install github.com/mrfansi/trecli/cmd/trecli@latest
export TRELLO_API_KEY=...   # https://trello.com/app-key
export TRELLO_TOKEN=...     # click "Token" on that page
```

Verify auth in one round trip:

```bash
trecli me
```

## Command shape

```text
trecli <group> <operation> [path-args...] [flags]
trecli raw <METHOD> <PATH> [flags]
trecli me                              # alias for `members get-members-id me`
```

- `<group>` — one of: `actions`, `applications`, `batch`, `boards`,
  `cards`, `checklists`, `customFields`, `emoji`, `enterprises`,
  `labels`, `lists`, `members`, `notifications`, `organizations`,
  `plugins`, `search`, `tokens`, `webhooks`.
- `<operation>` — kebab-case `<method>-<path>`, derived from the
  OpenAPI `operationId`. Examples: `get-boards-id`, `post-cards`,
  `put-cards-id`, `delete-cards-id`.
- Path args are positional, in spec order.
- Query parameters become `--<name>` flags.
- Endpoints with a request body accept `--data <json|@file>`.

The full command catalog (255 ops, every flag) lives at
`{baseDir}/../../docs/COMMANDS.md` in this repo, or
`docs/COMMANDS.md` from the project root. Read it when you need the
exact flag for a less common endpoint.

## Output handling

Every successful call prints raw JSON to stdout. Pipe through `jq`
for filtering. Examples:

```bash
trecli boards get-boards-id <board-id> --fields name,url | jq .name
trecli members get-members-id me | jq -r '.idBoards[]'
```

On non-2xx, stderr contains `trello api <status>: <body>` and the
process exits 1. Handle the error before continuing.

## Curated recipes

Replace bracketed placeholders with real IDs.

### 1. Identify the user

```bash
trecli me
```

Returns the authenticated member object. Use to surface username,
email, board membership.

### 2. List the user's boards

```bash
trecli members get-members-id-boards me --fields name,url,closed
```

Returns an array. Filter open boards: `| jq '[.[] | select(.closed | not)]'`.

### 3. Read a board (with lists and cards in one call)

```bash
trecli boards get-boards-id <board-id> \
  --fields name,url \
  --lists open \
  --cards open
```

The `--lists` and `--cards` query params eagerly include the nested
collections in the same response.

### 4. List the lists on a board

```bash
trecli boards get-boards-id-lists <board-id> --fields name,pos,closed
```

### 5. List the cards on a list

```bash
trecli lists get-lists-id-cards <list-id> --fields name,desc,due,idMembers,labels
```

### 6. Find a list by name on a board

```bash
trecli boards get-boards-id-lists <board-id> --fields name \
  | jq -r --arg n "<list-name>" '.[] | select(.name == $n) | .id'
```

### 7. Create a card

Simple:

```bash
trecli cards post-cards \
  --idList <list-id> \
  --name "Task title" \
  --desc "Task body"
```

Complex body (attachments, members, position, due):

```bash
trecli cards post-cards --data '{
  "idList": "<list-id>",
  "name": "Task",
  "desc": "...",
  "pos": "top",
  "due": "2026-12-31T23:59:00.000Z",
  "idMembers": ["<member-id>"]
}'
```

### 8. Update a card

Use a JSON body so multiple fields update in one call:

```bash
trecli cards put-cards-id <card-id> --data '{
  "name": "Renamed",
  "desc": "Updated description",
  "idList": "<another-list-id>"
}'
```

Field reference: `name`, `desc`, `closed`, `idMembers` (array),
`idAttachmentCover`, `idList`, `idLabels` (array), `pos`, `due`,
`dueComplete`, `subscribed`.

### 9. Move or archive a card

```bash
# Move
trecli cards put-cards-id <card-id> --data '{"idList":"<list-id>"}'

# Archive
trecli cards put-cards-id <card-id> --data '{"closed":true}'

# Unarchive
trecli cards put-cards-id <card-id> --data '{"closed":false}'
```

### 10. Delete a card

```bash
trecli cards delete-cards-id <card-id>
```

Destructive. Confirm with the user before calling.

### 11. Add a checklist with items

```bash
# Create the checklist on a card
CL_ID=$(trecli checklists post-checklists \
  --data '{"idCard":"<card-id>","name":"Steps"}' | jq -r .id)

# Add items
trecli checklists post-checklists-id-checkitems "$CL_ID" \
  --data '{"name":"First step","pos":"bottom"}'
```

### 12. Manage labels

```bash
# Create
trecli labels post-labels --data '{"idBoard":"<board-id>","name":"Bug","color":"red"}'

# Add label to card
trecli cards post-cards-id-idlabels <card-id> --value <label-id>

# List labels on a board
trecli boards get-boards-id-labels <board-id> --fields name,color
```

### 13. Search across Trello

```bash
trecli search get-search --query "<term>" --modelTypes cards,boards
```

`--modelTypes` accepts a comma-separated list of `actions`, `boards`,
`cards`, `members`, `organizations`. Returns grouped hits.

### 14. Webhooks

```bash
# Create a webhook subscribed to a model
trecli webhooks post-webhooks --data '{
  "idModel": "<board|card|list-id>",
  "callbackURL": "https://example.com/hook",
  "description": "My webhook"
}'

# List webhooks tied to your token
trecli tokens get-tokens-token-webhooks "$TRELLO_TOKEN"

# Delete a webhook
trecli webhooks delete-webhooks-id <webhook-id>
```

### 15. Raw passthrough (escape hatch)

When a recipe above does not match, drop to `raw`:

```bash
trecli raw GET /boards/{id}/memberships --path id=<board-id>
trecli raw POST /cards/{id}/actions/comments \
  --path id=<card-id> \
  --query text="Comment body"
trecli raw PUT /cards/{id} --path id=<card-id> --data @update.json
```

Flags: `--path key=value`, `--query key=value`, `--header key=value`,
`--data <json|@file>`. Equivalent to the curl command, with auth
auto-injected.

## Decision rules for agents

1. Always prefer a curated recipe over `raw` when one fits.
2. If the user names a board / list / card by **title**, look up the
   ID first (recipes 2, 4, 6) before mutating.
3. **Never** delete or archive without an explicit user confirmation.
4. If a request fails with 401, do not retry; surface the auth error
   so the user can rotate the token.
5. If a request fails with 404, double-check the ID — Trello IDs are
   24-character hex; the literal string `me` is a valid alias on
   `/members/{id}`.
6. JSON output is the source of truth — never invent fields you did
   not see.
7. Time fields are ISO-8601 with milliseconds and `Z` suffix
   (`2026-12-31T23:59:00.000Z`).

## Common errors

| Stderr | Meaning | Action |
|--------|---------|--------|
| `missing credentials` | env vars / config file missing | Ask user to set them; do not call again |
| `trello api 401 Unauthorized: invalid token` | Token revoked / wrong | Surface to user; do not retry |
| `trello api 400 Bad Request: invalid id` | Path arg is not a Trello ID | Look up the correct ID |
| `trello api 404 Not Found` | Resource missing or no access | Verify ID + membership |
| `path placeholder {x} not found` | `--path x=...` flag for an unknown placeholder | Inspect the path template; drop or rename the flag |
| `non-2xx response` (exit 1) | Any 3xx-5xx | Read stderr line for status + body |

## Reference

- Full command catalog: `{baseDir}/../../docs/COMMANDS.md` (or
  `docs/COMMANDS.md` from the repo root).
- Human usage guide: `{baseDir}/../../docs/USAGE.md`.
- Source: <https://github.com/mrfansi/trecli>.
- Trello REST docs: <https://developer.atlassian.com/cloud/trello/rest/>.
