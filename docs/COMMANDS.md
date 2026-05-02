# trello-cli command reference

Auto-generated from `openapi.json`. Do not edit by hand — re-run `make gen-cmds` to refresh.

**Coverage**: 18 resource groups, 255 operations.

## Resource groups

| Group | Operations |
|-------|-----------:|
| [`actions`](#actions) | 16 |
| [`applications`](#applications) | 1 |
| [`batch`](#batch) | 1 |
| [`boards`](#boards) | 36 |
| [`cards`](#cards) | 42 |
| [`checklists`](#checklists) | 12 |
| [`customFields`](#customfields) | 8 |
| [`emoji`](#emoji) | 1 |
| [`enterprises`](#enterprises) | 21 |
| [`labels`](#labels) | 5 |
| [`lists`](#lists) | 11 |
| [`members`](#members) | 44 |
| [`notifications`](#notifications) | 11 |
| [`organizations`](#organizations) | 26 |
| [`plugins`](#plugins) | 5 |
| [`search`](#search) | 2 |
| [`tokens`](#tokens) | 8 |
| [`webhooks`](#webhooks) | 5 |

Plus two handcrafted commands:

- `me` — alias for `members get-members-id me`.
- `raw <METHOD> <PATH>` — passthrough to any endpoint.

## actions

16 operations.

### `actions delete-actions-id`

`DELETE /actions/{id}`

Delete an Action

```bash
trello-cli actions delete-actions-id <id>
```

Path arguments:

- `<id>` — The ID of the Action

### `actions delete-actions-idaction-reactions-id`

`DELETE /actions/{idAction}/reactions/{id}`

Delete Action's Reaction

```bash
trello-cli actions delete-actions-idaction-reactions-id <idAction> <id>
```

Path arguments:

- `<idAction>` — The ID of the Action
- `<id>` — The ID of the reaction

### `actions get-actions-id`

`GET /actions/{id}`

Get an Action

```bash
trello-cli actions get-actions-id <id>
```

Path arguments:

- `<id>` — The ID of the Action

Query flags:

- `--display` — (no description)
- `--entities` — (no description)
- `--fields` — `all` or a comma-separated list of action [fields](/cloud/trello/guides/rest-api/object-definitions/#action-object)
- `--member` — (no description)
- `--member_fields` — `all` or a comma-separated list of member [fields](/cloud/trello/guides/rest-api/object-definitions/)
- `--memberCreator` — Whether to include the member object for the creator of the action
- `--memberCreator_fields` — `all` or a comma-separated list of member [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `actions get-actions-id-board`

`GET /actions/{id}/board`

Get the Board for an Action

```bash
trello-cli actions get-actions-id-board <id>
```

Path arguments:

- `<id>` — The ID of the action

Query flags:

- `--fields` — `all` or a comma-separated list of board fields

### `actions get-actions-id-card`

`GET /actions/{id}/card`

Get the Card for an Action

```bash
trello-cli actions get-actions-id-card <id>
```

Path arguments:

- `<id>` — The ID of the action

Query flags:

- `--fields` — `all` or a comma-separated list of card fields

### `actions get-actions-id-field`

`GET /actions/{id}/{field}`

Get a specific field on an Action

```bash
trello-cli actions get-actions-id-field <id> <field>
```

Path arguments:

- `<id>` — The ID of the Action
- `<field>` — An action field

### `actions get-actions-id-list`

`GET /actions/{id}/list`

Get the List for an Action

```bash
trello-cli actions get-actions-id-list <id>
```

Path arguments:

- `<id>` — The ID of the action

Query flags:

- `--fields` — `all` or a comma-separated list of list fields

### `actions get-actions-id-member`

`GET /actions/{id}/member`

Get the Member of an Action

```bash
trello-cli actions get-actions-id-member <id>
```

Path arguments:

- `<id>` — The ID of the Action

Query flags:

- `--fields` — `all` or a comma-separated list of member fields

### `actions get-actions-id-membercreator`

`GET /actions/{id}/memberCreator`

Get the Member Creator of an Action

```bash
trello-cli actions get-actions-id-membercreator <id>
```

Path arguments:

- `<id>` — The ID of the Action

Query flags:

- `--fields` — `all` or a comma-separated list of member fields

### `actions get-actions-id-organization`

`GET /actions/{id}/organization`

Get the Organization of an Action

```bash
trello-cli actions get-actions-id-organization <id>
```

Path arguments:

- `<id>` — The ID of the action

Query flags:

- `--fields` — `all` or a comma-separated list of organization fields

### `actions get-actions-idaction-reactions`

`GET /actions/{idAction}/reactions`

Get Action's Reactions

```bash
trello-cli actions get-actions-idaction-reactions <idAction>
```

Path arguments:

- `<idAction>` — The ID of the action

Query flags:

- `--member` — Whether to load the member as a nested resource. See [Members Nested Resource](/cloud/trello/guides/rest-api/nested-resources/#members-nested-resource)
- `--emoji` — Whether to load the emoji as a nested resource.

### `actions get-actions-idaction-reactions-id`

`GET /actions/{idAction}/reactions/{id}`

Get Action's Reaction

```bash
trello-cli actions get-actions-idaction-reactions-id <idAction> <id>
```

Path arguments:

- `<idAction>` — The ID of the Action
- `<id>` — The ID of the reaction

Query flags:

- `--member` — Whether to load the member as a nested resource. See [Members Nested Resource](/cloud/trello/guides/rest-api/nested-resources/#members-nested-resource)
- `--emoji` — Whether to load the emoji as a nested resource.

### `actions get-actions-idaction-reactionsummary`

`GET /actions/{idAction}/reactionsSummary`

List Action's summary of Reactions

```bash
trello-cli actions get-actions-idaction-reactionsummary <idAction>
```

Path arguments:

- `<idAction>` — The ID of the action

### `actions post-actions-idaction-reactions`

`POST /actions/{idAction}/reactions`

Create Reaction for Action

```bash
trello-cli actions post-actions-idaction-reactions <idAction>
```

Path arguments:

- `<idAction>` — The ID of the action

Body: `--data <json|@file>` (optional JSON request body).

### `actions put-actions-id`

`PUT /actions/{id}`

Update an Action

```bash
trello-cli actions put-actions-id <id>
```

Path arguments:

- `<id>` — The ID of the Action

Query flags:

- `--text` — The new text for the comment

### `actions put-actions-id-text`

`PUT /actions/{id}/text`

Update a Comment Action

```bash
trello-cli actions put-actions-id-text <id>
```

Path arguments:

- `<id>` — The ID of the action to update

Query flags:

- `--value` — The new text for the comment

## applications

1 operations.

### `applications applications-key-compliance`

`GET /applications/{key}/compliance`

Get Application's compliance data

```bash
trello-cli applications applications-key-compliance <key>
```

Path arguments:

- `<key>` — (no description)

## batch

1 operations.

### `batch get-batch`

`GET /batch`

Batch Requests

```bash
trello-cli batch get-batch
```

Query flags:

- `--urls` — A list of API routes. Maximum of 10 routes allowed. The routes should begin with a forward slash and should not include the API version number - e.g. "urls=/members/trello,/cards/[cardId]"

## boards

36 operations.

### `boards boards-id-checklists`

`GET /boards/{id}/checklists`

Get Checklists on a Board

```bash
trello-cli boards boards-id-checklists <id>
```

Path arguments:

- `<id>` — The ID of the board

### `boards boardsidmembersidmember`

`DELETE /boards/{id}/members/{idMember}`

Remove Member from Board

```bash
trello-cli boards boardsidmembersidmember <id> <idMember>
```

Path arguments:

- `<id>` — The id of the board to update
- `<idMember>` — The id of the member to add to the board.

### `boards delete-boards-id`

`DELETE /boards/{id}`

Delete a Board

```bash
trello-cli boards delete-boards-id <id>
```

Path arguments:

- `<id>` — The id of the board to delete

### `boards delete-boards-id-boardplugins`

`DELETE /boards/{id}/boardPlugins/{idPlugin}`

Disable a Power-Up on a Board

```bash
trello-cli boards delete-boards-id-boardplugins <id> <idPlugin>
```

Path arguments:

- `<id>` — The ID of the board
- `<idPlugin>` — The ID of the Power-Up to disable

### `boards get-board-id-plugins`

`GET /boards/{id}/plugins`

Get Power-Ups on a Board

```bash
trello-cli boards get-board-id-plugins <id>
```

Path arguments:

- `<id>` — The ID of the board

Query flags:

- `--filter` — One of: `enabled` or `available`

### `boards get-boards-id`

`GET /boards/{id}`

Get a Board

```bash
trello-cli boards get-boards-id <id>
```

Path arguments:

- `<id>` — (no description)

Query flags:

- `--actions` — This is a nested resource. Read more about actions as nested resources [here](/cloud/trello/guides/rest-api/nested-resources/).
- `--boardStars` — Valid values are one of: `mine` or `none`.
- `--cards` — This is a nested resource. Read more about cards as nested resources [here](/cloud/trello/guides/rest-api/nested-resources/).
- `--card_pluginData` — Use with the `cards` param to include card pluginData with the response
- `--checklists` — This is a nested resource. Read more about checklists as nested resources [here](/cloud/trello/guides/rest-api/nested-resources/).
- `--customFields` — This is a nested resource. Read more about custom fields as nested resources [here](#custom-fields-nested-resource).
- `--fields` — The fields of the board to be included in the response. Valid values: all or a comma-separated list of: closed, dateLastActivity, dateLastView, desc, descData, idMemberCreator, idOrganization, invi...
- `--labels` — This is a nested resource. Read more about labels as nested resources [here](/cloud/trello/guides/rest-api/nested-resources/).
- `--lists` — This is a nested resource. Read more about lists as nested resources [here](/cloud/trello/guides/rest-api/nested-resources/).
- `--members` — This is a nested resource. Read more about members as nested resources [here](/cloud/trello/guides/rest-api/nested-resources/).
- `--memberships` — This is a nested resource. Read more about memberships as nested resources [here](/cloud/trello/guides/rest-api/nested-resources/).
- `--pluginData` — Determines whether the pluginData for this board should be returned. Valid values: true or false.
- `--organization` — This is a nested resource. Read more about organizations as nested resources [here](/cloud/trello/guides/rest-api/nested-resources/).
- `--organization_pluginData` — Use with the `organization` param to include organization pluginData with the response
- `--myPrefs` — (no description)
- `--tags` — Also known as collections, tags, refer to the collection(s) that a Board belongs to.

### `boards get-boards-id-actions`

`GET /boards/{boardId}/actions`

Get Actions of a Board

```bash
trello-cli boards get-boards-id-actions <boardId>
```

Path arguments:

- `<boardId>` — (no description)

Query flags:

- `--fields` — The fields to be returned for the Actions. [See Action fields here](/cloud/trello/guides/rest-api/object-definitions/#action-object).
- `--filter` — A comma-separated list of [action types](/cloud/trello/guides/rest-api/action-types/).
- `--format` — The format of the returned Actions. Either list or count.
- `--idModels` — A comma-separated list of idModels. Only actions related to these models will be returned.
- `--limit` — The limit of the number of responses, between 0 and 1000.
- `--member` — Whether to return the member object for each action.
- `--member_fields` — The fields of the [member](/cloud/trello/guides/rest-api/object-definitions/#member-object) to return.
- `--memberCreator` — Whether to return the memberCreator object for each action.
- `--memberCreator_fields` — The fields of the [member](/cloud/trello/guides/rest-api/object-definitions/#member-object) creator to return
- `--page` — The page of results for actions.
- `--reactions` — Whether to show reactions on comments or not.
- `--before` — A date string in the form of YYYY-MM-DDThh:mm:ssZ or a mongo object ID. Only objects created before this date will be returned.
- `--since` — A date string in the form of YYYY-MM-DDThh:mm:ssZ or a mongo object ID. Only objects created since this date will be returned.

### `boards get-boards-id-boardplugins`

`GET /boards/{id}/boardPlugins`

Get Enabled Power-Ups on Board

```bash
trello-cli boards get-boards-id-boardplugins <id>
```

Path arguments:

- `<id>` — The ID of the Board

### `boards get-boards-id-boardstars`

`GET /boards/{boardId}/boardStars`

Get boardStars on a Board

```bash
trello-cli boards get-boards-id-boardstars <boardId>
```

Path arguments:

- `<boardId>` — (no description)

Query flags:

- `--filter` — Valid values: mine, none

### `boards get-boards-id-cards`

`GET /boards/{id}/cards`

Get Cards on a Board

```bash
trello-cli boards get-boards-id-cards <id>
```

Path arguments:

- `<id>` — (no description)

### `boards get-boards-id-cards-filter`

`GET /boards/{id}/cards/{filter}`

Get filtered Cards on a Board

```bash
trello-cli boards get-boards-id-cards-filter <id> <filter>
```

Path arguments:

- `<id>` — ID of the Board
- `<filter>` — One of: `all`, `closed`, `complete`, `incomplete`, `none`, `open`, `visible`

### `boards get-boards-id-customfields`

`GET /boards/{id}/customFields`

Get Custom Fields for Board

```bash
trello-cli boards get-boards-id-customfields <id>
```

Path arguments:

- `<id>` — The ID of the board

### `boards get-boards-id-field`

`GET /boards/{id}/{field}`

Get a field on a Board

```bash
trello-cli boards get-boards-id-field <id> <field>
```

Path arguments:

- `<id>` — The ID of the board.
- `<field>` — The field you'd like to receive. Valid values: closed, dateLastActivity, dateLastView, desc, descData, idMemberCreator, idOrganization, invitations, invited, labelNames, memberships, name, pinned, ...

### `boards get-boards-id-labels`

`GET /boards/{id}/labels`

Get Labels on a Board

```bash
trello-cli boards get-boards-id-labels <id>
```

Path arguments:

- `<id>` — The ID of the Board.

Query flags:

- `--fields` — The fields to be returned for the Labels.
- `--limit` — The number of Labels to be returned.

### `boards get-boards-id-lists`

`GET /boards/{id}/lists`

Get Lists on a Board

```bash
trello-cli boards get-boards-id-lists <id>
```

Path arguments:

- `<id>` — The ID of the board

Query flags:

- `--cards` — Filter to apply to Cards.
- `--card_fields` — `all` or a comma-separated list of card [fields](/cloud/trello/guides/rest-api/object-definitions/#card-object)
- `--filter` — Filter to apply to Lists
- `--fields` — `all` or a comma-separated list of list [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `boards get-boards-id-lists-filter`

`GET /boards/{id}/lists/{filter}`

Get filtered Lists on a Board

```bash
trello-cli boards get-boards-id-lists-filter <id> <filter>
```

Path arguments:

- `<id>` — The ID of the board
- `<filter>` — One of `all`, `closed`, `none`, `open`

### `boards get-boards-id-members`

`GET /boards/{id}/members`

Get the Members of a Board

```bash
trello-cli boards get-boards-id-members <id>
```

Path arguments:

- `<id>` — The ID of the board

### `boards get-boards-id-memberships`

`GET /boards/{id}/memberships`

Get Memberships of a Board

```bash
trello-cli boards get-boards-id-memberships <id>
```

Path arguments:

- `<id>` — The ID of the board

Query flags:

- `--filter` — One of `admins`, `all`, `none`, `normal`
- `--activity` — Works for premium organizations only.
- `--orgMemberType` — Shows the type of member to the org the user is. For instance, an org admin will have a `orgMemberType` of `admin`.
- `--member` — Determines whether to include a [nested member object](/cloud/trello/guides/rest-api/nested-resources/).
- `--member_fields` — Fields to show if `member=true`. Valid values: [nested member resource fields](/cloud/trello/guides/rest-api/nested-resources/).

### `boards post-boards`

`POST /boards/`

Create a Board

```bash
trello-cli boards post-boards
```

Query flags:

- `--name` — The new name for the board. 1 to 16384 characters long.
- `--defaultLabels` — Determines whether to use the default set of labels.
- `--defaultLists` — Determines whether to add the default set of lists to a board (To Do, Doing, Done). It is ignored if `idBoardSource` is provided.
- `--desc` — A new description for the board, 0 to 16384 characters long
- `--idOrganization` — The id or name of the Workspace the board should belong to.
- `--idBoardSource` — The id of a board to copy into the new board.
- `--keepFromSource` — To keep cards from the original board pass in the value `cards`
- `--powerUps` — The Power-Ups that should be enabled on the new board. One of: `all`, `calendar`, `cardAging`, `recap`, `voting`.
- `--prefs_permissionLevel` — The permissions level of the board. One of: `org`, `private`, `public`.
- `--prefs_voting` — Who can vote on this board. One of `disabled`, `members`, `observers`, `org`, `public`.
- `--prefs_comments` — Who can comment on cards on this board. One of: `disabled`, `members`, `observers`, `org`, `public`.
- `--prefs_invitations` — Determines what types of members can invite users to join. One of: `admins`, `members`.
- `--prefs_selfJoin` — Determines whether users can join the boards themselves or whether they have to be invited.
- `--prefs_cardCovers` — Determines whether card covers are enabled.
- `--prefs_background` — The id of a custom background or one of: `blue`, `orange`, `green`, `red`, `purple`, `pink`, `lime`, `sky`, `grey`.
- `--prefs_cardAging` — Determines the type of card aging that should take place on the board if card aging is enabled. One of: `pirate`, `regular`.

### `boards post-boards-id-boardplugins`

`POST /boards/{id}/boardPlugins`

Enable a Power-Up on a Board

```bash
trello-cli boards post-boards-id-boardplugins <id>
```

Path arguments:

- `<id>` — The ID of the Board

Query flags:

- `--idPlugin` — The ID of the Power-Up to enable

### `boards post-boards-id-calendarkey-generate`

`POST /boards/{id}/calendarKey/generate`

Create a calendarKey for a Board

```bash
trello-cli boards post-boards-id-calendarkey-generate <id>
```

Path arguments:

- `<id>` — The id of the board to update

### `boards post-boards-id-emailkey-generate`

`POST /boards/{id}/emailKey/generate`

Create a emailKey for a Board

```bash
trello-cli boards post-boards-id-emailkey-generate <id>
```

Path arguments:

- `<id>` — The id of the board to update

### `boards post-boards-id-idtags`

`POST /boards/{id}/idTags`

Create a Tag for a Board

```bash
trello-cli boards post-boards-id-idtags <id>
```

Path arguments:

- `<id>` — The id of the board to update

Query flags:

- `--value` — The id of a tag from the organization to which this board belongs.

### `boards post-boards-id-labels`

`POST /boards/{id}/labels`

Create a Label on a Board

```bash
trello-cli boards post-boards-id-labels <id>
```

Path arguments:

- `<id>` — The id of the board to update

Query flags:

- `--name` — The name of the label to be created. 1 to 16384 characters long.
- `--color` — Sets the color of the new label. Valid values are a label color or `null`.

### `boards post-boards-id-lists`

`POST /boards/{id}/lists`

Create a List on a Board

```bash
trello-cli boards post-boards-id-lists <id>
```

Path arguments:

- `<id>` — The ID of the board

Query flags:

- `--name` — The name of the list to be created. 1 to 16384 characters long.
- `--pos` — Determines the position of the list. Valid values: `top`, `bottom`, or a positive number.

### `boards post-boards-id-markedasviewed`

`POST /boards/{id}/markedAsViewed`

Mark Board as viewed

```bash
trello-cli boards post-boards-id-markedasviewed <id>
```

Path arguments:

- `<id>` — The id of the board to update

### `boards put-boards-id`

`PUT /boards/{id}`

Update a Board

```bash
trello-cli boards put-boards-id <id>
```

Path arguments:

- `<id>` — (no description)

Query flags:

- `--name` — The new name for the board. 1 to 16384 characters long.
- `--desc` — A new description for the board, 0 to 16384 characters long
- `--closed` — Whether the board is closed
- `--subscribed` — Whether the acting user is subscribed to the board
- `--idOrganization` — The id of the Workspace the board should be moved to
- `--prefs/permissionLevel` — One of: org, private, public
- `--prefs/selfJoin` — Whether Workspace members can join the board themselves
- `--prefs/cardCovers` — Whether card covers should be displayed on this board
- `--prefs/hideVotes` — Determines whether the Voting Power-Up should hide who voted on cards or not.
- `--prefs/invitations` — Who can invite people to this board. One of: admins, members
- `--prefs/voting` — Who can vote on this board. One of disabled, members, observers, org, public
- `--prefs/comments` — Who can comment on cards on this board. One of: disabled, members, observers, org, public
- `--prefs/background` — The id of a custom background or one of: blue, orange, green, red, purple, pink, lime, sky, grey
- `--prefs/cardAging` — One of: pirate, regular
- `--prefs/calendarFeedEnabled` — Determines whether the calendar feed is enabled or not.

### `boards put-boards-id-members`

`PUT /boards/{id}/members`

Invite Member to Board via email

```bash
trello-cli boards put-boards-id-members <id>
```

Path arguments:

- `<id>` — The ID of the board

Query flags:

- `--email` — The email address of a user to add as a member of the board.
- `--type` — Valid values: admin, normal, observer. Determines what type of member the user being added should be of the board.

Body: `--data <json|@file>` (optional JSON request body).

### `boards put-boards-id-members-idmember`

`PUT /boards/{id}/members/{idMember}`

Add a Member to a Board

```bash
trello-cli boards put-boards-id-members-idmember <id> <idMember>
```

Path arguments:

- `<id>` — The id of the board to update
- `<idMember>` — The id of the member to add to the board.

Query flags:

- `--type` — One of: admin, normal, observer. Determines the type of member this user will be on the board.
- `--allowBillableGuest` — Optional param that allows organization admins to add multi-board guests onto a board.

### `boards put-boards-id-memberships-idmembership`

`PUT /boards/{id}/memberships/{idMembership}`

Update Membership of Member on a Board

```bash
trello-cli boards put-boards-id-memberships-idmembership <id> <idMembership>
```

Path arguments:

- `<id>` — The id of the board to update
- `<idMembership>` — The id of a membership that should be added to this board.

Query flags:

- `--type` — One of: admin, normal, observer. Determines the type of member that this membership will be to this board.
- `--member_fields` — Valid values: all, avatarHash, bio, bioData, confirmed, fullName, idPremOrgsAdmin, initials, memberType, products, status, url, username

### `boards put-boards-id-myprefs-emailposition`

`PUT /boards/{id}/myPrefs/emailPosition`

Update emailPosition Pref on a Board

```bash
trello-cli boards put-boards-id-myprefs-emailposition <id>
```

Path arguments:

- `<id>` — The id of the board to update

Query flags:

- `--value` — Valid values: bottom, top. Determines the position of the email address.

### `boards put-boards-id-myprefs-idemaillist`

`PUT /boards/{id}/myPrefs/idEmailList`

Update idEmailList Pref on a Board

```bash
trello-cli boards put-boards-id-myprefs-idemaillist <id>
```

Path arguments:

- `<id>` — The id of the board to update

Query flags:

- `--value` — The id of an email list.

### `boards put-boards-id-myprefs-showsidebar`

`PUT /boards/{id}/myPrefs/showSidebar`

Update showSidebar Pref on a Board

```bash
trello-cli boards put-boards-id-myprefs-showsidebar <id>
```

Path arguments:

- `<id>` — The id of the board to update

Query flags:

- `--value` — Determines whether to show the side bar.

### `boards put-boards-id-myprefs-showsidebaractivity`

`PUT /boards/{id}/myPrefs/showSidebarActivity`

Update showSidebarActivity Pref on a Board

```bash
trello-cli boards put-boards-id-myprefs-showsidebaractivity <id>
```

Path arguments:

- `<id>` — The id of the board to update

Query flags:

- `--value` — Determines whether to show sidebar activity.

### `boards put-boards-id-myprefs-showsidebarboardactions`

`PUT /boards/{id}/myPrefs/showSidebarBoardActions`

Update showSidebarBoardActions Pref on a Board

```bash
trello-cli boards put-boards-id-myprefs-showsidebarboardactions <id>
```

Path arguments:

- `<id>` — The id of the board to update

Query flags:

- `--value` — Determines whether to show the sidebar board actions.

### `boards put-boards-id-myprefs-showsidebarmembers`

`PUT /boards/{id}/myPrefs/showSidebarMembers`

Update showSidebarMembers Pref on a Board

```bash
trello-cli boards put-boards-id-myprefs-showsidebarmembers <id>
```

Path arguments:

- `<id>` — The id of the board to update

Query flags:

- `--value` — Determines whether to show members of the board in the sidebar.

## cards

42 operations.

### `cards cardsidmembersvoted-1`

`POST /cards/{id}/membersVoted`

Add Member vote to Card

```bash
trello-cli cards cardsidmembersvoted-1 <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--value` — The ID of the member to vote 'yes' on the card

### `cards delete-cards-id`

`DELETE /cards/{id}`

Delete a Card

```bash
trello-cli cards delete-cards-id <id>
```

Path arguments:

- `<id>` — The ID of the Card

### `cards delete-cards-id-actions-id-comments`

`DELETE /cards/{id}/actions/{idAction}/comments`

Delete a comment on a Card

```bash
trello-cli cards delete-cards-id-actions-id-comments <id> <idAction>
```

Path arguments:

- `<id>` — The ID of the Card
- `<idAction>` — The ID of the comment action to update

### `cards delete-cards-id-checkitem-idcheckitem`

`DELETE /cards/{id}/checkItem/{idCheckItem}`

Delete checkItem on a Card

```bash
trello-cli cards delete-cards-id-checkitem-idcheckitem <id> <idCheckItem>
```

Path arguments:

- `<id>` — The ID of the Card
- `<idCheckItem>` — The ID of the checkitem

### `cards delete-cards-id-checklists-idchecklist`

`DELETE /cards/{id}/checklists/{idChecklist}`

Delete a Checklist on a Card

```bash
trello-cli cards delete-cards-id-checklists-idchecklist <id> <idChecklist>
```

Path arguments:

- `<id>` — The ID of the Card
- `<idChecklist>` — The ID of the checklist to delete

### `cards delete-cards-id-idlabels-idlabel`

`DELETE /cards/{id}/idLabels/{idLabel}`

Remove a Label from a Card

```bash
trello-cli cards delete-cards-id-idlabels-idlabel <id> <idLabel>
```

Path arguments:

- `<id>` — The ID of the Card
- `<idLabel>` — The ID of the label to remove

### `cards delete-cards-id-membersvoted-idmember`

`DELETE /cards/{id}/membersVoted/{idMember}`

Remove a Member's Vote on a Card

```bash
trello-cli cards delete-cards-id-membersvoted-idmember <id> <idMember>
```

Path arguments:

- `<id>` — The ID of the Card
- `<idMember>` — The ID of the member whose vote to remove

### `cards delete-cards-id-stickers-idsticker`

`DELETE /cards/{id}/stickers/{idSticker}`

Delete a Sticker on a Card

```bash
trello-cli cards delete-cards-id-stickers-idsticker <id> <idSticker>
```

Path arguments:

- `<id>` — The ID of the Card
- `<idSticker>` — The ID of the sticker

### `cards delete-id-idmembers-idmember`

`DELETE /cards/{id}/idMembers/{idMember}`

Remove a Member from a Card

```bash
trello-cli cards delete-id-idmembers-idmember <id> <idMember>
```

Path arguments:

- `<id>` — The ID of the Card
- `<idMember>` — The ID of the member to remove from the card

### `cards deleted-cards-id-attachments-idattachment`

`DELETE /cards/{id}/attachments/{idAttachment}`

Delete an Attachment on a Card

```bash
trello-cli cards deleted-cards-id-attachments-idattachment <id> <idAttachment>
```

Path arguments:

- `<id>` — The ID of the Card
- `<idAttachment>` — The ID of the attachment to delete

### `cards get-cards-id`

`GET /cards/{id}`

Get a Card

```bash
trello-cli cards get-cards-id <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--fields` — `all` or a comma-separated list of [fields](/cloud/trello/guides/rest-api/object-definitions/). **Defaults**: `badges, checkItemStates, closed, dateLastActivity, desc, descData, due, start, idBoard...
- `--actions` — See the [Actions Nested Resource](/cloud/trello/guides/rest-api/nested-resources/#actions-nested-resource)
- `--attachments` — `true`, `false`, or `cover`
- `--attachment_fields` — `all` or a comma-separated list of attachment [fields](/cloud/trello/guides/rest-api/object-definitions/)
- `--members` — Whether to return member objects for members on the card
- `--member_fields` — `all` or a comma-separated list of member [fields](/cloud/trello/guides/rest-api/object-definitions/). **Defaults**: `avatarHash, fullName, initials, username`
- `--membersVoted` — Whether to return member objects for members who voted on the card
- `--memberVoted_fields` — `all` or a comma-separated list of member [fields](/cloud/trello/guides/rest-api/object-definitions/). **Defaults**: `avatarHash, fullName, initials, username`
- `--checkItemStates` — (no description)
- `--checklists` — Whether to return the checklists on the card. `all` or `none`
- `--checklist_fields` — `all` or a comma-separated list of `idBoard,idCard,name,pos`
- `--board` — Whether to return the board object the card is on
- `--board_fields` — `all` or a comma-separated list of board [fields](/cloud/trello/guides/rest-api/object-definitions/#board-object). **Defaults**: `name, desc, descData, closed, idOrganization, pinned, url, prefs`
- `--list` — See the [Lists Nested Resource](/cloud/trello/guides/rest-api/nested-resources/)
- `--pluginData` — Whether to include pluginData on the card with the response
- `--stickers` — Whether to include sticker models with the response
- `--sticker_fields` — `all` or a comma-separated list of sticker [fields](/cloud/trello/guides/rest-api/object-definitions/)
- `--customFieldItems` — Whether to include the customFieldItems

### `cards get-cards-id-actions`

`GET /cards/{id}/actions`

Get Actions on a Card

```bash
trello-cli cards get-cards-id-actions <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--filter` — A comma-separated list of [action types](https://developer.atlassian.com/cloud/trello/guides/rest-api/action-types/).
- `--page` — The page of results for actions. Each page of results has 50 actions.

### `cards get-cards-id-attachments`

`GET /cards/{id}/attachments`

Get Attachments on a Card

```bash
trello-cli cards get-cards-id-attachments <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--fields` — `all` or a comma-separated list of attachment [fields](/cloud/trello/guides/rest-api/object-definitions/)
- `--filter` — Use `cover` to restrict to just the cover attachment

### `cards get-cards-id-attachments-idattachment`

`GET /cards/{id}/attachments/{idAttachment}`

Get an Attachment on a Card

```bash
trello-cli cards get-cards-id-attachments-idattachment <id> <idAttachment>
```

Path arguments:

- `<id>` — The ID of the Card
- `<idAttachment>` — The ID of the Attachment

Query flags:

- `--fields` — The Attachment fields to be included in the response.

### `cards get-cards-id-board`

`GET /cards/{id}/board`

Get the Board the Card is on

```bash
trello-cli cards get-cards-id-board <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--fields` — `all` or a comma-separated list of board [fields](/cloud/trello/guides/rest-api/object-definitions/#board-object)

### `cards get-cards-id-checkitem-idcheckitem`

`GET /cards/{id}/checkItem/{idCheckItem}`

Get checkItem on a Card

```bash
trello-cli cards get-cards-id-checkitem-idcheckitem <id> <idCheckItem>
```

Path arguments:

- `<id>` — The ID of the Card
- `<idCheckItem>` — The ID of the checkitem

Query flags:

- `--fields` — `all` or a comma-separated list of `name,nameData,pos,state,type,due,dueReminder,idMember`

### `cards get-cards-id-checkitemstates`

`GET /cards/{id}/checkItemStates`

Get checkItems on a Card

```bash
trello-cli cards get-cards-id-checkitemstates <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--fields` — `all` or a comma-separated list of: `idCheckItem`, `state`

### `cards get-cards-id-checklists`

`GET /cards/{id}/checklists`

Get Checklists on a Card

```bash
trello-cli cards get-cards-id-checklists <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--checkItems` — `all` or `none`
- `--checkItem_fields` — `all` or a comma-separated list of: `name,nameData,pos,state,type,due,dueReminder,idMember`
- `--filter` — `all` or `none`
- `--fields` — `all` or a comma-separated list of: `idBoard,idCard,name,pos`

### `cards get-cards-id-customfielditems`

`GET /cards/{id}/customFieldItems`

Get Custom Field Items for a Card

```bash
trello-cli cards get-cards-id-customfielditems <id>
```

Path arguments:

- `<id>` — The ID of the Card

### `cards get-cards-id-field`

`GET /cards/{id}/{field}`

Get a field on a Card

```bash
trello-cli cards get-cards-id-field <id> <field>
```

Path arguments:

- `<id>` — The ID of the Card
- `<field>` — The desired field.

### `cards get-cards-id-list`

`GET /cards/{id}/list`

Get the List of a Card

```bash
trello-cli cards get-cards-id-list <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--fields` — `all` or a comma-separated list of list [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `cards get-cards-id-members`

`GET /cards/{id}/members`

Get the Members of a Card

```bash
trello-cli cards get-cards-id-members <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--fields` — `all` or a comma-separated list of member [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `cards get-cards-id-membersvoted`

`GET /cards/{id}/membersVoted`

Get Members who have voted on a Card

```bash
trello-cli cards get-cards-id-membersvoted <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--fields` — `all` or a comma-separated list of member [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `cards get-cards-id-plugindata`

`GET /cards/{id}/pluginData`

Get pluginData on a Card

```bash
trello-cli cards get-cards-id-plugindata <id>
```

Path arguments:

- `<id>` — The ID of the Card

### `cards get-cards-id-stickers`

`GET /cards/{id}/stickers`

Get Stickers on a Card

```bash
trello-cli cards get-cards-id-stickers <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--fields` — `all` or a comma-separated list of sticker [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `cards get-cards-id-stickers-idsticker`

`GET /cards/{id}/stickers/{idSticker}`

Get a Sticker on a Card

```bash
trello-cli cards get-cards-id-stickers-idsticker <id> <idSticker>
```

Path arguments:

- `<id>` — The ID of the Card
- `<idSticker>` — The ID of the sticker

Query flags:

- `--fields` — `all` or a comma-separated list of sticker [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `cards post-cards`

`POST /cards`

Create a new Card

```bash
trello-cli cards post-cards
```

Query flags:

- `--name` — The name for the card
- `--desc` — The description for the card
- `--pos` — The position of the new card. `top`, `bottom`, or a positive float
- `--due` — A due date for the card
- `--start` — The start date of a card, or `null`
- `--dueComplete` — Whether the status of the card is complete
- `--idList` — The ID of the list the card should be created in
- `--idMembers` — Comma-separated list of member IDs to add to the card
- `--idLabels` — Comma-separated list of label IDs to add to the card
- `--urlSource` — A URL starting with `http://` or `https://`. The URL will be attached to the card upon creation.
- `--fileSource` — (no description)
- `--mimeType` — The mimeType of the attachment. Max length 256
- `--idCardSource` — The ID of a card to copy into the new card
- `--keepFromSource` — If using `idCardSource` you can specify which properties to copy over. `all` or comma-separated list of: `attachments,checklists,customFields,comments,due,start,labels,members,start,stickers`
- `--address` — For use with/by the Map View
- `--locationName` — For use with/by the Map View
- `--coordinates` — For use with/by the Map View. Should take the form latitude,longitude
- `--cardRole` — For displaying cards in different ways based on the card name. Board cards must have a name that is a link to a Trello board. Mirror cards must have a name that is a link to a Trello card.

### `cards post-cards-id-actions-comments`

`POST /cards/{id}/actions/comments`

Add a new comment to a Card

```bash
trello-cli cards post-cards-id-actions-comments <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--text` — The comment

### `cards post-cards-id-attachments`

`POST /cards/{id}/attachments`

Create Attachment On Card

```bash
trello-cli cards post-cards-id-attachments <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--name` — The name of the attachment. Max length 256.
- `--file` — The file to attach, as multipart/form-data
- `--mimeType` — The mimeType of the attachment. Max length 256
- `--url` — A URL to attach. Must start with `http://` or `https://`
- `--setCover` — Determines whether to use the new attachment as a cover for the Card.

### `cards post-cards-id-checklists`

`POST /cards/{id}/checklists`

Create Checklist on a Card

```bash
trello-cli cards post-cards-id-checklists <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--name` — The name of the checklist
- `--idChecklistSource` — The ID of a source checklist to copy into the new one
- `--pos` — The position of the checklist on the card. One of: `top`, `bottom`, or a positive number.

### `cards post-cards-id-idlabels`

`POST /cards/{id}/idLabels`

Add a Label to a Card

```bash
trello-cli cards post-cards-id-idlabels <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--value` — The ID of the label to add

### `cards post-cards-id-idmembers`

`POST /cards/{id}/idMembers`

Add a Member to a Card

```bash
trello-cli cards post-cards-id-idmembers <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--value` — The ID of the Member to add to the card

### `cards post-cards-id-labels`

`POST /cards/{id}/labels`

Create a new Label on a Card

```bash
trello-cli cards post-cards-id-labels <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--color` — A valid label color or `null`. See [labels](/cloud/trello/guides/rest-api/object-definitions/)
- `--name` — A name for the label

### `cards post-cards-id-markassociatednotificationsread`

`POST /cards/{id}/markAssociatedNotificationsRead`

Mark a Card's Notifications as read

```bash
trello-cli cards post-cards-id-markassociatednotificationsread <id>
```

Path arguments:

- `<id>` — The ID of the Card

### `cards post-cards-id-stickers`

`POST /cards/{id}/stickers`

Add a Sticker to a Card

```bash
trello-cli cards post-cards-id-stickers <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--image` — For custom stickers, the id of the sticker. For default stickers, the string identifier (like 'taco-cool', see below)
- `--top` — The top position of the sticker, from -60 to 100
- `--left` — The left position of the sticker, from -60 to 100
- `--zIndex` — The z-index of the sticker
- `--rotate` — The rotation of the sticker

### `cards put-cards-id`

`PUT /cards/{id}`

Update a Card

```bash
trello-cli cards put-cards-id <id>
```

Path arguments:

- `<id>` — The ID of the Card

Query flags:

- `--name` — The new name for the card
- `--desc` — The new description for the card
- `--closed` — Whether the card should be archived (closed: true)
- `--idMembers` — Comma-separated list of member IDs
- `--idAttachmentCover` — The ID of the image attachment the card should use as its cover, or null for none
- `--idList` — The ID of the list the card should be in
- `--idLabels` — Comma-separated list of label IDs
- `--idBoard` — The ID of the board the card should be on
- `--pos` — The position of the card in its list. `top`, `bottom`, or a positive float
- `--due` — When the card is due, or `null`
- `--start` — The start date of a card, or `null`
- `--dueComplete` — Whether the status of the card is complete
- `--subscribed` — Whether the member is should be subscribed to the card
- `--address` — For use with/by the Map View
- `--locationName` — For use with/by the Map View
- `--coordinates` — For use with/by the Map View. Should be latitude,longitude
- `--cover` — Updates the card's cover | Option | Values | About | |--------|--------|-------| | color | `pink`, `yellow`, `lime`, `blue`, `black`, `orange`, `red`, `purple`, `sky`, `green` | Makes the cover a s...

### `cards put-cards-id-actions-idaction-comments`

`PUT /cards/{id}/actions/{idAction}/comments`

Update Comment Action on a Card

```bash
trello-cli cards put-cards-id-actions-idaction-comments <id> <idAction>
```

Path arguments:

- `<id>` — The ID of the Card
- `<idAction>` — The ID of the comment action to update

Query flags:

- `--text` — The new text for the comment

### `cards put-cards-id-checkitem-idcheckitem`

`PUT /cards/{id}/checkItem/{idCheckItem}`

Update a checkItem on a Card

```bash
trello-cli cards put-cards-id-checkitem-idcheckitem <id> <idCheckItem>
```

Path arguments:

- `<id>` — The ID of the Card
- `<idCheckItem>` — The ID of the checkitem

Query flags:

- `--name` — The new name for the checklist item
- `--state` — One of: `complete`, `incomplete`
- `--idChecklist` — The ID of the checklist this item is in
- `--pos` — `top`, `bottom`, or a positive float
- `--due` — A due date for the checkitem
- `--dueReminder` — A dueReminder for the due date on the checkitem
- `--idMember` — The ID of the member to remove from the card

### `cards put-cards-id-stickers-idsticker`

`PUT /cards/{id}/stickers/{idSticker}`

Update a Sticker on a Card

```bash
trello-cli cards put-cards-id-stickers-idsticker <id> <idSticker>
```

Path arguments:

- `<id>` — The ID of the Card
- `<idSticker>` — The ID of the sticker

Query flags:

- `--top` — The top position of the sticker, from -60 to 100
- `--left` — The left position of the sticker, from -60 to 100
- `--zIndex` — The z-index of the sticker
- `--rotate` — The rotation of the sticker

### `cards put-cards-idcard-checklist-idchecklist-checkitem-idcheckitem`

`PUT /cards/{idCard}/checklist/{idChecklist}/checkItem/{idCheckItem}`

Update Checkitem on Checklist on Card

```bash
trello-cli cards put-cards-idcard-checklist-idchecklist-checkitem-idcheckitem <idCard> <idCheckItem> <idChecklist>
```

Path arguments:

- `<idCard>` — The ID of the Card
- `<idCheckItem>` — The ID of the checklist item to update
- `<idChecklist>` — The ID of the item to update.

Query flags:

- `--pos` — `top`, `bottom`, or a positive float

### `cards put-cards-idcard-customfield-idcustomfield-item`

`PUT /cards/{idCard}/customField/{idCustomField}/item`

Update Custom Field item on Card

```bash
trello-cli cards put-cards-idcard-customfield-idcustomfield-item <idCard> <idCustomField>
```

Path arguments:

- `<idCard>` — ID of the card that the Custom Field value should be set/updated for
- `<idCustomField>` — ID of the Custom Field on the card.

Body: `--data <json|@file>` (optional JSON request body).

### `cards put-cards-idcard-customfields`

`PUT /cards/{idCard}/customFields`

Update Multiple Custom Field items on Card

```bash
trello-cli cards put-cards-idcard-customfields <idCard>
```

Path arguments:

- `<idCard>` — (no description)

Body: `--data <json|@file>` (optional JSON request body).

## checklists

12 operations.

### `checklists delete-checklists-id`

`DELETE /checklists/{id}`

Delete a Checklist

```bash
trello-cli checklists delete-checklists-id <id>
```

Path arguments:

- `<id>` — ID of a checklist.

### `checklists delete-checklists-id-checkitems-idcheckitem`

`DELETE /checklists/{id}/checkItems/{idCheckItem}`

Delete Checkitem from Checklist

```bash
trello-cli checklists delete-checklists-id-checkitems-idcheckitem <id> <idCheckItem>
```

Path arguments:

- `<id>` — ID of a checklist.
- `<idCheckItem>` — ID of the check item to retrieve.

### `checklists get-checklists-id`

`GET /checklists/{id}`

Get a Checklist

```bash
trello-cli checklists get-checklists-id <id>
```

Path arguments:

- `<id>` — ID of a checklist.

Query flags:

- `--cards` — Valid values: `all`, `closed`, `none`, `open`, `visible`. Cards is a nested resource. The additional query params available are documented at [Cards Nested Resource](/cloud/trello/guides/rest-api/n...
- `--checkItems` — The check items on the list to return. One of: `all`, `none`.
- `--checkItem_fields` — The fields on the checkItem to return if checkItems are being returned. `all` or a comma-separated list of: `name`, `nameData`, `pos`, `state`, `type`, `due`, `dueReminder`, `idMember`
- `--fields` — `all` or a comma-separated list of checklist [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `checklists get-checklists-id-board`

`GET /checklists/{id}/board`

Get the Board the Checklist is on

```bash
trello-cli checklists get-checklists-id-board <id>
```

Path arguments:

- `<id>` — ID of a checklist.

Query flags:

- `--fields` — `all` or a comma-separated list of board [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `checklists get-checklists-id-cards`

`GET /checklists/{id}/cards`

Get the Card a Checklist is on

```bash
trello-cli checklists get-checklists-id-cards <id>
```

Path arguments:

- `<id>` — ID of a checklist.

### `checklists get-checklists-id-checkitems`

`GET /checklists/{id}/checkItems`

Get Checkitems on a Checklist

```bash
trello-cli checklists get-checklists-id-checkitems <id>
```

Path arguments:

- `<id>` — ID of a checklist.

Query flags:

- `--filter` — One of: `all`, `none`.
- `--fields` — One of: `all`, `name`, `nameData`, `pos`, `state`,`type`, `due`, `dueReminder`, `idMember`.

### `checklists get-checklists-id-checkitems-idcheckitem`

`GET /checklists/{id}/checkItems/{idCheckItem}`

Get a Checkitem on a Checklist

```bash
trello-cli checklists get-checklists-id-checkitems-idcheckitem <id> <idCheckItem>
```

Path arguments:

- `<id>` — ID of a checklist.
- `<idCheckItem>` — ID of the check item to retrieve.

Query flags:

- `--fields` — One of: `all`, `name`, `nameData`, `pos`, `state`, `type`, `due`, `dueReminder`, `idMember`,.

### `checklists get-checklists-id-field`

`GET /checklists/{id}/{field}`

Get field on a Checklist

```bash
trello-cli checklists get-checklists-id-field <id> <field>
```

Path arguments:

- `<id>` — ID of a checklist.
- `<field>` — Field to update.

### `checklists post-checklists`

`POST /checklists`

Create a Checklist

```bash
trello-cli checklists post-checklists
```

Query flags:

- `--idCard` — The ID of the Card that the checklist should be added to.
- `--name` — The name of the checklist. Should be a string of length 1 to 16384.
- `--pos` — The position of the checklist on the card. One of: `top`, `bottom`, or a positive number.
- `--idChecklistSource` — The ID of a checklist to copy into the new checklist.

### `checklists post-checklists-id-checkitems`

`POST /checklists/{id}/checkItems`

Create Checkitem on Checklist

```bash
trello-cli checklists post-checklists-id-checkitems <id>
```

Path arguments:

- `<id>` — ID of a checklist.

Query flags:

- `--name` — The name of the new check item on the checklist. Should be a string of length 1 to 16384.
- `--pos` — The position of the check item in the checklist. One of: `top`, `bottom`, or a positive number.
- `--checked` — Determines whether the check item is already checked when created.
- `--due` — A due date for the checkitem
- `--dueReminder` — A dueReminder for the due date on the checkitem
- `--idMember` — An ID of a member resource.

### `checklists put-checklists-id-field`

`PUT /checklists/{id}/{field}`

Update field on a Checklist

```bash
trello-cli checklists put-checklists-id-field <id> <field>
```

Path arguments:

- `<id>` — ID of a checklist.
- `<field>` — Field to update.

Query flags:

- `--value` — The value to change the checklist name to. Should be a string of length 1 to 16384.

### `checklists put-checlists-id`

`PUT /checklists/{id}`

Update a Checklist

```bash
trello-cli checklists put-checlists-id <id>
```

Path arguments:

- `<id>` — ID of a checklist.

Query flags:

- `--name` — Name of the new checklist being created. Should be length of 1 to 16384.
- `--pos` — Determines the position of the checklist on the card. One of: `top`, `bottom`, or a positive number.

## customFields

8 operations.

### `customFields delete-customfields-id`

`DELETE /customFields/{id}`

Delete a Custom Field definition

```bash
trello-cli customFields delete-customfields-id <id>
```

Path arguments:

- `<id>` — ID of the Custom Field.

### `customFields delete-customfields-options-idcustomfieldoption`

`DELETE /customFields/{id}/options/{idCustomFieldOption}`

Delete Option of Custom Field dropdown

```bash
trello-cli customFields delete-customfields-options-idcustomfieldoption <id> <idCustomFieldOption>
```

Path arguments:

- `<id>` — ID of the customfielditem.
- `<idCustomFieldOption>` — ID of the customfieldoption to retrieve.

### `customFields get-customfields-id`

`GET /customFields/{id}`

Get a Custom Field

```bash
trello-cli customFields get-customfields-id <id>
```

Path arguments:

- `<id>` — ID of the Custom Field.

### `customFields get-customfields-id-options`

`POST /customFields/{id}/options`

Add Option to Custom Field dropdown

```bash
trello-cli customFields get-customfields-id-options <id>
```

Path arguments:

- `<id>` — ID of the customfield.

### `customFields get-customfields-options-idcustomfieldoption`

`GET /customFields/{id}/options/{idCustomFieldOption}`

Get Option of Custom Field dropdown

```bash
trello-cli customFields get-customfields-options-idcustomfieldoption <id> <idCustomFieldOption>
```

Path arguments:

- `<id>` — ID of the customfielditem.
- `<idCustomFieldOption>` — ID of the customfieldoption to retrieve.

### `customFields post-customfields`

`POST /customFields`

Create a new Custom Field on a Board

```bash
trello-cli customFields post-customfields
```

Body: `--data <json|@file>` (optional JSON request body).

### `customFields post-customfields-id-options`

`GET /customFields/{id}/options`

Get Options of Custom Field drop down

```bash
trello-cli customFields post-customfields-id-options <id>
```

Path arguments:

- `<id>` — ID of the customfield.

### `customFields put-customfields-id`

`PUT /customFields/{id}`

Update a Custom Field definition

```bash
trello-cli customFields put-customfields-id <id>
```

Path arguments:

- `<id>` — ID of the Custom Field.

Body: `--data <json|@file>` (optional JSON request body).

## emoji

1 operations.

### `emoji emoji`

`GET /emoji`

List available Emoji

```bash
trello-cli emoji emoji
```

Query flags:

- `--locale` — The locale to return emoji descriptions and names in. Defaults to the logged in member's locale.
- `--spritesheets` — `true` to return spritesheet URLs in the response

## enterprises

21 operations.

### `enterprises delete-enterprises-id-organizations-idorg`

`DELETE /enterprises/{id}/organizations/{idOrg}`

Delete an Organization from an Enterprise.

```bash
trello-cli enterprises delete-enterprises-id-organizations-idorg <id> <idOrg>
```

Path arguments:

- `<id>` — ID of the enterprise to retrieve.
- `<idOrg>` — ID of the organization to be removed from the enterprise.

### `enterprises enterprises-id-members-idmember-deactivated`

`PUT /enterprises/{id}/members/{idMember}/deactivated`

Deactivate a Member of an Enterprise.

```bash
trello-cli enterprises enterprises-id-members-idmember-deactivated <id> <idMember>
```

Path arguments:

- `<id>` — ID of the enterprise to retrieve.
- `<idMember>` — ID of the Member to deactive.

Query flags:

- `--value` — Determines whether the user is deactivated or not.
- `--fields` — A comma separated list of any valid values that the [nested member field resource]() accepts.
- `--organization_fields` — Any valid value that the [nested organization resource](/cloud/trello/guides/rest-api/nested-resources/) accepts.
- `--board_fields` — Any valid value that the [nested board resource](/cloud/trello/guides/rest-api/nested-resources/) accepts.

### `enterprises enterprises-id-organizations-idmember`

`DELETE /enterprises/{id}/admins/{idMember}`

Remove a Member as admin from Enterprise.

```bash
trello-cli enterprises enterprises-id-organizations-idmember <id> <idMember>
```

Path arguments:

- `<id>` — ID of the Enterprise to retrieve.
- `<idMember>` — ID of the member to be removed as an admin from enterprise.

### `enterprises get-enterprises-id`

`GET /enterprises/{id}`

Get an Enterprise

```bash
trello-cli enterprises get-enterprises-id <id>
```

Path arguments:

- `<id>` — ID of the enterprise to retrieve.

Query flags:

- `--fields` — Comma-separated list of: `id`, `name`, `displayName`, `prefs`, `ssoActivationFailed`, `idAdmins`, `idMembers` (Note that the members array returned will be paginated if `members` is 'normal' or 'ad...
- `--members` — One of: `none`, `normal`, `admins`, `owners`, `all`
- `--member_fields` — One of: `avatarHash`, `fullName`, `initials`, `username`
- `--member_filter` — Pass a SCIM-style query to filter members. This takes precedence over the all/normal/admins value of members. If any of the member_* args are set, the member array will be paginated.
- `--member_sort` — This parameter expects a SCIM-style sorting value prefixed by a `-` to sort descending. If no `-` is prefixed, it will be sorted ascending. Note that the members array returned will be paginated if...
- `--member_sortBy` — Deprecated: Please use member_sort. This parameter expects a SCIM-style sorting value. Note that the members array returned will be paginated if `members` is `normal` or `admins`. Pagination can be...
- `--member_sortOrder` — Deprecated: Please use member_sort. One of: `ascending`, `descending`, `asc`, `desc`
- `--member_startIndex` — Any integer between 0 and 100.
- `--member_count` — 0 to 100
- `--organizations` — One of: `none`, `members`, `public`, `all`
- `--organization_fields` — Any valid value that the [nested organization field resource]() accepts.
- `--organization_paid_accounts` — Whether or not to include paid account information in the returned workspace objects
- `--organization_memberships` — Comma-seperated list of: `me`, `normal`, `admin`, `active`, `deactivated`

### `enterprises get-enterprises-id-admins`

`GET /enterprises/{id}/admins`

Get Enterprise admin Members

```bash
trello-cli enterprises get-enterprises-id-admins <id>
```

Path arguments:

- `<id>` — ID of the enterprise to retrieve.

Query flags:

- `--fields` — Any valid value that the [nested member field resource]() accepts.

### `enterprises get-enterprises-id-auditlog`

`GET /enterprises/{id}/auditlog`

Get auditlog data for an Enterprise

```bash
trello-cli enterprises get-enterprises-id-auditlog <id>
```

Path arguments:

- `<id>` — ID of the enterprise to retrieve.

### `enterprises get-enterprises-id-claimableorganizations`

`GET /enterprises/{id}/claimableOrganizations`

Get ClaimableOrganizations of an Enterprise

```bash
trello-cli enterprises get-enterprises-id-claimableorganizations <id>
```

Path arguments:

- `<id>` — ID of the enterprise to retrieve

Query flags:

- `--limit` — Limits the number of workspaces to be sorted
- `--cursor` — Specifies the sort order to return matching documents
- `--name` — Name of the enterprise to retrieve workspaces for
- `--activeSince` — Date in YYYY-MM-DD format indicating the date to search up to for activeness of workspace
- `--inactiveSince` — Date in YYYY-MM-DD format indicating the date to search up to for inactiveness of workspace

### `enterprises get-enterprises-id-members`

`GET /enterprises/{id}/members`

Get Members of Enterprise

```bash
trello-cli enterprises get-enterprises-id-members <id>
```

Path arguments:

- `<id>` — ID of the Enterprise to retrieve.

Query flags:

- `--fields` — A comma-seperated list of valid [member fields](/cloud/trello/guides/rest-api/object-definitions/#member-object).
- `--filter` — Pass a SCIM-style query to filter members. This takes precedence over the all/normal/admins value of members. If any of the below member_* args are set, the member array will be paginated.
- `--sort` — This parameter expects a SCIM-style sorting value prefixed by a `-` to sort descending. If no `-` is prefixed, it will be sorted ascending. Note that the members array returned will be paginated if...
- `--sortBy` — Deprecated: Please use `sort` instead. This parameter expects a SCIM-style sorting value. Note that the members array returned will be paginated if `members` is 'normal' or 'admins'. Pagination can...
- `--sortOrder` — Deprecated: Please use `sort` instead. One of: `ascending`, `descending`, `asc`, `desc`.
- `--startIndex` — Any integer between 0 and 9999.
- `--count` — SCIM-style filter.
- `--organization_fields` — Any valid value that the [nested organization field resource](/cloud/trello/guides/rest-api/nested-resources/) accepts.
- `--board_fields` — Any valid value that the [nested board resource](/cloud/trello/guides/rest-api/nested-resources/) accepts.

### `enterprises get-enterprises-id-members-idmember`

`GET /enterprises/{id}/members/{idMember}`

Get a Member of Enterprise

```bash
trello-cli enterprises get-enterprises-id-members-idmember <id> <idMember>
```

Path arguments:

- `<id>` — ID of the enterprise to retrieve.
- `<idMember>` — An ID of a member resource.

Query flags:

- `--fields` — A comma separated list of any valid values that the [nested member field resource]() accepts.
- `--organization_fields` — Any valid value that the [nested organization field resource](/cloud/trello/guides/rest-api/nested-resources/) accepts.
- `--board_fields` — Any valid value that the [nested board resource](/cloud/trello/guides/rest-api/nested-resources/) accepts.

### `enterprises get-enterprises-id-organizations`

`GET /enterprises/{id}/organizations`

Get Organizations of an Enterprise

```bash
trello-cli enterprises get-enterprises-id-organizations <id>
```

Path arguments:

- `<id>` — ID of the Enterprise to retrieve.

Query flags:

- `--fields` — comma-separated list of organization [fields](/cloud/trello/guides/rest-api/object-definitions/)
- `--filter` — (no description)
- `--startIndex` — Any integer greater than and equal to 1.
- `--count` — Any integer between 0 and 100.

### `enterprises get-enterprises-id-organizations-bulk-idorganizations`

`GET /enterprises/{id}/organizations/bulk/{idOrganizations}`

Bulk accept a set of organizations to an Enterprise.

```bash
trello-cli enterprises get-enterprises-id-organizations-bulk-idorganizations <id> <idOrganizations>
```

Path arguments:

- `<id>` — ID of the enterprise to retrieve.
- `<idOrganizations>` — An array of IDs of the organizations to be removed from the enterprise.

### `enterprises get-enterprises-id-pendingorganizations`

`GET /enterprises/{id}/pendingOrganizations`

Get PendingOrganizations of an Enterprise

```bash
trello-cli enterprises get-enterprises-id-pendingorganizations <id>
```

Path arguments:

- `<id>` — ID of the enterprise to retrieve

Query flags:

- `--activeSince` — Date in YYYY-MM-DD format indicating the date to search up to for activeness of workspace
- `--inactiveSince` — Date in YYYY-MM-DD format indicating the date to search up to for inactiveness of workspace

### `enterprises get-enterprises-id-signupurl`

`GET /enterprises/{id}/signupUrl`

Get signupUrl for Enterprise

```bash
trello-cli enterprises get-enterprises-id-signupurl <id>
```

Path arguments:

- `<id>` — ID of the enterprise to retrieve.

Query flags:

- `--authenticate` — (no description)
- `--confirmationAccepted` — (no description)
- `--returnUrl` — Any valid URL.
- `--tosAccepted` — Designates whether the user has seen/consented to the Trello ToS prior to being redirected to the enterprise signup page/their IdP.

### `enterprises get-enterprises-id-transferrable-bulk-idorganizations`

`GET /enterprises/{id}/transferrable/bulk/{idOrganizations}`

Get a bulk list of organizations that can be transferred to an enterprise.

```bash
trello-cli enterprises get-enterprises-id-transferrable-bulk-idorganizations <id> <idOrganizations>
```

Path arguments:

- `<id>` — ID of the Enterprise to retrieve.
- `<idOrganizations>` — An array of IDs of an Organization resource.

### `enterprises get-enterprises-id-transferrable-organization-idorganization`

`GET /enterprises/{id}/transferrable/organization/{idOrganization}`

Get whether an organization can be transferred to an enterprise.

```bash
trello-cli enterprises get-enterprises-id-transferrable-organization-idorganization <id> <idOrganization>
```

Path arguments:

- `<id>` — ID of the Enterprise to retrieve.
- `<idOrganization>` — An ID of an Organization resource.

### `enterprises get-users-id`

`GET /enterprises/{id}/members/query`

Get Users of an Enterprise

```bash
trello-cli enterprises get-users-id <id>
```

Path arguments:

- `<id>` — ID of the enterprise to retrieve.

Query flags:

- `--licensed` — When true, returns members who possess a license for the corresponding Trello Enterprise; when false, returns members who do not. If unspecified, both licensed and unlicensed members will be returned.
- `--deactivated` — When true, returns members who have been deactivated for the corresponding Trello Enterprise; when false, returns members who have not. If unspecified, both active and deactivated members will be r...
- `--collaborator` — When true, returns members who are guests on one or more boards in the corresponding Trello Enterprise (but do not possess a license); when false, returns members who are not. If unspecified, both ...
- `--managed` — When true, returns members who are managed by the corresponding Trello Enterprise; when false, returns members who are not. If unspecified, both managed and unmanaged members will be returned.
- `--admin` — When true, returns members who are administrators of the corresponding Trello Enterprise; when false, returns members who are not. If unspecified, both admin and non-admin members will be returned.
- `--activeSince` — Returns only Trello users active since this date (inclusive).
- `--inactiveSince` — Returns only Trello users active since this date (inclusive).
- `--search` — Returns members with email address or full name that start with the search value.
- `--cursor` — Cursor to return next set of results, use cursor returned in the response to query the next batch.

### `enterprises post-enterprises-id-tokens`

`POST /enterprises/{id}/tokens`

Create an auth Token for an Enterprise.

```bash
trello-cli enterprises post-enterprises-id-tokens <id>
```

Path arguments:

- `<id>` — ID of the enterprise to retrieve.

Query flags:

- `--expiration` — One of: `1hour`, `1day`, `30days`, `never`

### `enterprises put-enterprises-id-admins-idmember`

`PUT /enterprises/{id}/admins/{idMember}`

Update Member to be admin of Enterprise

```bash
trello-cli enterprises put-enterprises-id-admins-idmember <id> <idMember>
```

Path arguments:

- `<id>` — ID of the enterprise to retrieve.
- `<idMember>` — ID of member to be made an admin of enterprise.

### `enterprises put-enterprises-id-enterprisejoinrequest-bulk`

`PUT /enterprises/${id}/enterpriseJoinRequest/bulk`

Decline enterpriseJoinRequests from one organization or a bulk list of organizations.

```bash
trello-cli enterprises put-enterprises-id-enterprisejoinrequest-bulk <id>
```

Path arguments:

- `<id>` — ID of the Enterprise to retrieve.

Query flags:

- `--idOrganizations` — An array of IDs of an Organization resource.

### `enterprises put-enterprises-id-members-idmember-licensed`

`PUT /enterprises/{id}/members/{idMember}/licensed`

Update a Member's licensed status

```bash
trello-cli enterprises put-enterprises-id-members-idmember-licensed <id> <idMember>
```

Path arguments:

- `<id>` — ID of the Enterprise.
- `<idMember>` — The ID of the Member

Query flags:

- `--value` — Boolean value to determine whether the user should be given an Enterprise license (true) or not (false).

### `enterprises put-enterprises-id-organizations`

`PUT /enterprises/{id}/organizations`

Transfer an Organization to an Enterprise.

```bash
trello-cli enterprises put-enterprises-id-organizations <id>
```

Path arguments:

- `<id>` — ID of the Enterprise to retrieve.

Query flags:

- `--idOrganization` — ID of Organization to be transferred to Enterprise.

## labels

5 operations.

### `labels delete-labels-id`

`DELETE /labels/{id}`

Delete a Label

```bash
trello-cli labels delete-labels-id <id>
```

Path arguments:

- `<id>` — The ID of the Label

### `labels get-labels-id`

`GET /labels/{id}`

Get a Label

```bash
trello-cli labels get-labels-id <id>
```

Path arguments:

- `<id>` — The ID of the Label

Query flags:

- `--fields` — all or a comma-separated list of [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `labels post-labels`

`POST /labels`

Create a Label

```bash
trello-cli labels post-labels
```

Query flags:

- `--name` — Name for the label
- `--color` — The color for the label.
- `--idBoard` — The ID of the Board to create the Label on.

### `labels put-labels-id`

`PUT /labels/{id}`

Update a Label

```bash
trello-cli labels put-labels-id <id>
```

Path arguments:

- `<id>` — The ID of the Label

Query flags:

- `--name` — The new name for the label
- `--color` — The new color for the label. See: [fields](/cloud/trello/guides/rest-api/object-definitions/) for color options

### `labels put-labels-id-field`

`PUT /labels/{id}/{field}`

Update a field on a label

```bash
trello-cli labels put-labels-id-field <id> <field>
```

Path arguments:

- `<id>` — The id of the label
- `<field>` — The field on the Label to update.

Query flags:

- `--value` — The new value for the field.

## lists

11 operations.

### `lists get-lists-id`

`GET /lists/{id}`

Get a List

```bash
trello-cli lists get-lists-id <id>
```

Path arguments:

- `<id>` — The ID of the list

Query flags:

- `--fields` — `all` or a comma separated list of List field names.

### `lists get-lists-id-actions`

`GET /lists/{id}/actions`

Get Actions for a List

```bash
trello-cli lists get-lists-id-actions <id>
```

Path arguments:

- `<id>` — The ID of the list

Query flags:

- `--filter` — A comma-separated list of [action types](https://developer.atlassian.com/cloud/trello/guides/rest-api/action-types/).

### `lists get-lists-id-board`

`GET /lists/{id}/board`

Get the Board a List is on

```bash
trello-cli lists get-lists-id-board <id>
```

Path arguments:

- `<id>` — The ID of the list

Query flags:

- `--fields` — `all` or a comma-separated list of board [fields](/cloud/trello/guides/rest-api/object-definitions/#board-object)

### `lists get-lists-id-cards`

`GET /lists/{id}/cards`

Get Cards in a List

```bash
trello-cli lists get-lists-id-cards <id>
```

Path arguments:

- `<id>` — The ID of the list

### `lists post-lists`

`POST /lists`

Create a new List

```bash
trello-cli lists post-lists
```

Query flags:

- `--name` — Name for the list
- `--idBoard` — The long ID of the board the list should be created on
- `--idListSource` — ID of the List to copy into the new List
- `--pos` — Position of the list. `top`, `bottom`, or a positive floating point number

### `lists post-lists-id-archiveallcards`

`POST /lists/{id}/archiveAllCards`

Archive all Cards in List

```bash
trello-cli lists post-lists-id-archiveallcards <id>
```

Path arguments:

- `<id>` — The ID of the list

### `lists post-lists-id-moveallcards`

`POST /lists/{id}/moveAllCards`

Move all Cards in List

```bash
trello-cli lists post-lists-id-moveallcards <id>
```

Path arguments:

- `<id>` — The ID of the list

Query flags:

- `--idBoard` — The ID of the board the cards should be moved to
- `--idList` — The ID of the list that the cards should be moved to

### `lists put-id-idboard`

`PUT /lists/{id}/idBoard`

Move List to Board

```bash
trello-cli lists put-id-idboard <id>
```

Path arguments:

- `<id>` — The ID of the list

Query flags:

- `--value` — The ID of the board to move the list to

### `lists put-lists-id`

`PUT /lists/{id}`

Update a List

```bash
trello-cli lists put-lists-id <id>
```

Path arguments:

- `<id>` — The ID of the list

Query flags:

- `--name` — New name for the list
- `--closed` — Whether the list should be closed (archived)
- `--idBoard` — ID of a board the list should be moved to
- `--pos` — New position for the list: `top`, `bottom`, or a positive floating point number
- `--subscribed` — Whether the active member is subscribed to this list

### `lists put-lists-id-closed`

`PUT /lists/{id}/closed`

Archive or unarchive a list

```bash
trello-cli lists put-lists-id-closed <id>
```

Path arguments:

- `<id>` — The ID of the list

Query flags:

- `--value` — Set to true to close (archive) the list

### `lists put-lists-id-field`

`PUT /lists/{id}/{field}`

Update a field on a List

```bash
trello-cli lists put-lists-id-field <id> <field>
```

Path arguments:

- `<id>` — The ID of the list
- `<field>` — The field on the List to be updated

Query flags:

- `--value` — The new value for the field

## members

44 operations.

### `members delete-members-id-boardbackgrounds-idbackground`

`DELETE /members/{id}/boardBackgrounds/{idBackground}`

Delete a Member's custom Board background

```bash
trello-cli members delete-members-id-boardbackgrounds-idbackground <id> <idBackground>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<idBackground>` — The ID of the board background

### `members delete-members-id-boardstars-idstar`

`DELETE /members/{id}/boardStars/{idStar}`

Delete Star for Board

```bash
trello-cli members delete-members-id-boardstars-idstar <id> <idStar>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<idStar>` — The ID of the board star

### `members delete-members-id-customboardbackgrounds-idbackground`

`DELETE /members/{id}/customBoardBackgrounds/{idBackground}`

Delete custom Board Background of Member

```bash
trello-cli members delete-members-id-customboardbackgrounds-idbackground <id> <idBackground>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<idBackground>` — The ID of the custom background

### `members delete-members-id-customstickers-idsticker`

`DELETE /members/{id}/customStickers/{idSticker}`

Delete a Member's custom Sticker

```bash
trello-cli members delete-members-id-customstickers-idsticker <id> <idSticker>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<idSticker>` — The ID of the uploaded sticker

### `members delete-members-id-savedsearches-idsearch`

`DELETE /members/{id}/savedSearches/{idSearch}`

Delete a saved search

```bash
trello-cli members delete-members-id-savedsearches-idsearch <id> <idSearch>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<idSearch>` — The ID of the saved search to delete

### `members get-members-id-actions`

`GET /members/{id}/actions`

Get a Member's Actions

```bash
trello-cli members get-members-id-actions <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--filter` — A comma-separated list of [action types](https://developer.atlassian.com/cloud/trello/guides/rest-api/action-types/).

### `members get-members-id-boardbackgrounds`

`GET /members/{id}/boardBackgrounds`

Get Member's custom Board backgrounds

```bash
trello-cli members get-members-id-boardbackgrounds <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--filter` — One of: `all`, `custom`, `default`, `none`, `premium`

### `members get-members-id-boardbackgrounds-idbackground`

`GET /members/{id}/boardBackgrounds/{idBackground}`

Get a boardBackground of a Member

```bash
trello-cli members get-members-id-boardbackgrounds-idbackground <id> <idBackground>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<idBackground>` — The ID of the board background

Query flags:

- `--fields` — `all` or a comma-separated list of: `brightness`, `fullSizeUrl`, `scaled`, `tile`

### `members get-members-id-boards`

`GET /members/{id}/boards`

Get Boards that Member belongs to

```bash
trello-cli members get-members-id-boards <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--filter` — `all` or a comma-separated list of: `closed`, `members`, `open`, `organization`, `public`, `starred`
- `--fields` — `all` or a comma-separated list of board [fields](/cloud/trello/guides/rest-api/object-definitions/)
- `--lists` — Which lists to include with the boards. One of: `all`, `closed`, `none`, `open`
- `--organization` — Whether to include the Organization object with the Boards
- `--organization_fields` — `all` or a comma-separated list of organization [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `members get-members-id-boardsinvited`

`GET /members/{id}/boardsInvited`

Get Boards the Member has been invited to

```bash
trello-cli members get-members-id-boardsinvited <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--fields` — `all` or a comma-separated list of board [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `members get-members-id-boardstars`

`GET /members/{id}/boardStars`

Get a Member's boardStars

```bash
trello-cli members get-members-id-boardstars <id>
```

Path arguments:

- `<id>` — The ID or username of the member

### `members get-members-id-boardstars-idstar`

`GET /members/{id}/boardStars/{idStar}`

Get a boardStar of Member

```bash
trello-cli members get-members-id-boardstars-idstar <id> <idStar>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<idStar>` — The ID of the board star

### `members get-members-id-cards`

`GET /members/{id}/cards`

Get Cards the Member is on

```bash
trello-cli members get-members-id-cards <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--filter` — One of: `all`, `closed`, `complete`, `incomplete`, `none`, `open`, `visible`

### `members get-members-id-customboardbackgrounds`

`GET /members/{id}/customBoardBackgrounds`

Get a Member's custom Board Backgrounds

```bash
trello-cli members get-members-id-customboardbackgrounds <id>
```

Path arguments:

- `<id>` — The ID or username of the member

### `members get-members-id-customboardbackgrounds-idbackground`

`GET /members/{id}/customBoardBackgrounds/{idBackground}`

Get custom Board Background of Member

```bash
trello-cli members get-members-id-customboardbackgrounds-idbackground <id> <idBackground>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<idBackground>` — The ID of the custom background

### `members get-members-id-customemoji`

`GET /members/{id}/customEmoji`

Get a Member's customEmojis

```bash
trello-cli members get-members-id-customemoji <id>
```

Path arguments:

- `<id>` — The ID or username of the member

### `members get-members-id-customstickers`

`GET /members/{id}/customStickers`

Get Member's custom Stickers

```bash
trello-cli members get-members-id-customstickers <id>
```

Path arguments:

- `<id>` — The ID or username of the member

### `members get-members-id-customstickers-idsticker`

`GET /members/{id}/customStickers/{idSticker}`

Get a Member's custom Sticker

```bash
trello-cli members get-members-id-customstickers-idsticker <id> <idSticker>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<idSticker>` — The ID of the uploaded sticker

Query flags:

- `--fields` — `all` or a comma-separated list of `scaled`, `url`

### `members get-members-id-field`

`GET /members/{id}/{field}`

Get a field on a Member

```bash
trello-cli members get-members-id-field <id> <field>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<field>` — One of the member [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `members get-members-id-notificationchannelsettings`

`GET /members/{id}/notificationsChannelSettings`

Get a Member's notification channel settings

```bash
trello-cli members get-members-id-notificationchannelsettings <id>
```

Path arguments:

- `<id>` — The ID or username of the member

### `members get-members-id-notificationchannelsettings-channel`

`GET /members/{id}/notificationsChannelSettings/{channel}`

Get blocked notification keys of Member on this channel

```bash
trello-cli members get-members-id-notificationchannelsettings-channel <id> <channel>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<channel>` — Channel to block notifications on

### `members get-members-id-notifications`

`GET /members/{id}/notifications`

Get Member's Notifications

```bash
trello-cli members get-members-id-notifications <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--entities` — (no description)
- `--display` — (no description)
- `--filter` — (no description)
- `--read_filter` — One of: `all`, `read`, `unread`
- `--fields` — `all` or a comma-separated list of notification [fields](/cloud/trello/guides/rest-api/object-definitions/)
- `--limit` — Max 1000
- `--page` — Max 100
- `--before` — A notification ID
- `--since` — A notification ID
- `--memberCreator` — (no description)
- `--memberCreator_fields` — `all` or a comma-separated list of member [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `members get-members-id-organizations`

`GET /members/{id}/organizations`

Get Member's Organizations

```bash
trello-cli members get-members-id-organizations <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--filter` — One of: `all`, `members`, `none`, `public` (Note: `members` filters to only private Workspaces)
- `--fields` — `all` or a comma-separated list of organization [fields](/cloud/trello/guides/rest-api/object-definitions/)
- `--paid_account` — Whether or not to include paid account information in the returned workspace object

### `members get-members-id-organizationsinvited`

`GET /members/{id}/organizationsInvited`

Get Organizations a Member has been invited to

```bash
trello-cli members get-members-id-organizationsinvited <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--fields` — `all` or a comma-separated list of organization [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `members get-members-id-savedsearches`

`GET /members/{id}/savedSearches`

Get Member's saved searched

```bash
trello-cli members get-members-id-savedsearches <id>
```

Path arguments:

- `<id>` — The ID or username of the member

### `members get-members-id-savedsearches-idsearch`

`GET /members/{id}/savedSearches/{idSearch}`

Get a saved search

```bash
trello-cli members get-members-id-savedsearches-idsearch <id> <idSearch>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<idSearch>` — The ID of the saved search to delete

### `members get-members-id-tokens`

`GET /members/{id}/tokens`

Get Member's Tokens

```bash
trello-cli members get-members-id-tokens <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--webhooks` — Whether to include webhooks

### `members get-members=id`

`GET /members/{id}`

Get a Member

```bash
trello-cli members get-members=id <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--actions` — See the [Actions Nested Resource](/cloud/trello/guides/rest-api/nested-resources/#actions-nested-resource)
- `--boards` — See the [Boards Nested Resource](/cloud/trello/guides/rest-api/nested-resources/#boards-nested-resource)
- `--boardBackgrounds` — One of: `all`, `custom`, `default`, `none`, `premium`
- `--boardsInvited` — `all` or a comma-separated list of: closed, members, open, organization, pinned, public, starred, unpinned
- `--boardsInvited_fields` — `all` or a comma-separated list of board [fields](/cloud/trello/guides/rest-api/object-definitions/)
- `--boardStars` — Whether to return the boardStars or not
- `--cards` — See the [Cards Nested Resource](/cloud/trello/guides/rest-api/nested-resources/#cards-nested-resource) for additional options
- `--customBoardBackgrounds` — `all` or `none`
- `--customEmoji` — `all` or `none`
- `--customStickers` — `all` or `none`
- `--fields` — `all` or a comma-separated list of member [fields](/cloud/trello/guides/rest-api/object-definitions/)
- `--notifications` — See the [Notifications Nested Resource](/cloud/trello/guides/rest-api/nested-resources/#notifications-nested-resource)
- `--organizations` — One of: `all`, `members`, `none`, `public`
- `--organization_fields` — `all` or a comma-separated list of organization [fields](/cloud/trello/guides/rest-api/object-definitions/)
- `--organization_paid_account` — Whether or not to include paid account information in the returned workspace object
- `--organizationsInvited` — One of: `all`, `members`, `none`, `public`
- `--organizationsInvited_fields` — `all` or a comma-separated list of organization [fields](/cloud/trello/guides/rest-api/object-definitions/)
- `--paid_account` — Whether or not to include paid account information in the returned member object
- `--savedSearches` — (no description)
- `--tokens` — `all` or `none`

### `members membersidavatar`

`POST /members/{id}/avatar`

Create Avatar for Member

```bash
trello-cli members membersidavatar <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--file` — (no description)

### `members membersidcustomboardbackgrounds-1`

`POST /members/{id}/customBoardBackgrounds`

Create a new custom Board Background

```bash
trello-cli members membersidcustomboardbackgrounds-1 <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--file` — (no description)

### `members membersidcustomemojiidemoji`

`GET /members/{id}/customEmoji/{idEmoji}`

Get a Member's custom Emoji

```bash
trello-cli members membersidcustomemojiidemoji <id> <idEmoji>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<idEmoji>` — The ID of the custom emoji

Query flags:

- `--fields` — `all` or a comma-separated list of `name`, `url`

### `members post-members-id-boardbackgrounds-1`

`POST /members/{id}/boardBackgrounds`

Upload new boardBackground for Member

```bash
trello-cli members post-members-id-boardbackgrounds-1 <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--file` — (no description)

### `members post-members-id-boardstars`

`POST /members/{id}/boardStars`

Create Star for Board

```bash
trello-cli members post-members-id-boardstars <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--idBoard` — The ID of the board to star
- `--pos` — The position of the newly starred board. `top`, `bottom`, or a positive float.

### `members post-members-id-customemoji`

`POST /members/{id}/customEmoji`

Create custom Emoji for Member

```bash
trello-cli members post-members-id-customemoji <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--file` — (no description)
- `--name` — Name for the emoji. 2 - 64 characters

### `members post-members-id-customstickers`

`POST /members/{id}/customStickers`

Create custom Sticker for Member

```bash
trello-cli members post-members-id-customstickers <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--file` — (no description)

### `members post-members-id-onetimemessagesdismissed`

`POST /members/{id}/oneTimeMessagesDismissed`

Dismiss a message for Member

```bash
trello-cli members post-members-id-onetimemessagesdismissed <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--value` — The message to dismiss

### `members post-members-id-savedsearches`

`POST /members/{id}/savedSearches`

Create saved Search for Member

```bash
trello-cli members post-members-id-savedsearches <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--name` — The name for the saved search
- `--query` — The search query
- `--pos` — The position of the saved search. `top`, `bottom`, or a positive float.

### `members put-members-id`

`PUT /members/{id}`

Update a Member

```bash
trello-cli members put-members-id <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Query flags:

- `--fullName` — New name for the member. Cannot begin or end with a space.
- `--initials` — New initials for the member. 1-4 characters long.
- `--username` — New username for the member. At least 3 characters long, only lowercase letters, underscores, and numbers. Must be unique.
- `--bio` — (no description)
- `--avatarSource` — One of: `gravatar`, `none`, `upload`
- `--prefs/colorBlind` — (no description)
- `--prefs/locale` — (no description)
- `--prefs/minutesBetweenSummaries` — `-1` for disabled, `1`, or `60`

### `members put-members-id-boardbackgrounds-idbackground`

`PUT /members/{id}/boardBackgrounds/{idBackground}`

Update a Member's custom Board background

```bash
trello-cli members put-members-id-boardbackgrounds-idbackground <id> <idBackground>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<idBackground>` — The ID of the board background

Query flags:

- `--brightness` — One of: `dark`, `light`, `unknown`
- `--tile` — Whether the background should be tiled

### `members put-members-id-boardstars-idstar`

`PUT /members/{id}/boardStars/{idStar}`

Update the position of a boardStar of Member

```bash
trello-cli members put-members-id-boardstars-idstar <id> <idStar>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<idStar>` — The ID of the board star

Query flags:

- `--pos` — New position for the starred board. `top`, `bottom`, or a positive float.

### `members put-members-id-customboardbackgrounds-idbackground`

`PUT /members/{id}/customBoardBackgrounds/{idBackground}`

Update custom Board Background of Member

```bash
trello-cli members put-members-id-customboardbackgrounds-idbackground <id> <idBackground>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<idBackground>` — The ID of the custom background

Query flags:

- `--brightness` — One of: `dark`, `light`, `unknown`
- `--tile` — Whether to tile the background

### `members put-members-id-notificationchannelsettings`

`PUT /members/{id}/notificationsChannelSettings`

Update blocked notification keys of Member on a channel

```bash
trello-cli members put-members-id-notificationchannelsettings <id>
```

Path arguments:

- `<id>` — The ID or username of the member

Body: `--data <json|@file>` (optional JSON request body).

### `members put-members-id-notificationchannelsettings-channel`

`PUT /members/{id}/notificationsChannelSettings/{channel}`

Update blocked notification keys of Member on a channel

```bash
trello-cli members put-members-id-notificationchannelsettings-channel <id> <channel>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<channel>` — Channel to block notifications on

Body: `--data <json|@file>` (optional JSON request body).

### `members put-members-id-savedsearches-idsearch`

`PUT /members/{id}/savedSearches/{idSearch}`

Update a saved search

```bash
trello-cli members put-members-id-savedsearches-idsearch <id> <idSearch>
```

Path arguments:

- `<id>` — The ID or username of the member
- `<idSearch>` — The ID of the saved search to delete

Query flags:

- `--name` — The new name for the saved search
- `--query` — The new search query
- `--pos` — New position for saves search. `top`, `bottom`, or a positive float.

## notifications

11 operations.

### `notifications get-notifications-id`

`GET /notifications/{id}`

Get a Notification

```bash
trello-cli notifications get-notifications-id <id>
```

Path arguments:

- `<id>` — The ID of the notification

Query flags:

- `--board` — Whether to include the board object
- `--board_fields` — `all` or a comma-separated list of board [fields](/cloud/trello/guides/rest-api/object-definitions/)
- `--card` — Whether to include the card object
- `--card_fields` — `all` or a comma-separated list of card [fields](/cloud/trello/guides/rest-api/object-definitions/)
- `--display` — Whether to include the display object with the results
- `--entities` — Whether to include the entities object with the results
- `--fields` — `all` or a comma-separated list of notification [fields](/cloud/trello/guides/rest-api/object-definitions/)
- `--list` — Whether to include the list object
- `--member` — Whether to include the member object
- `--member_fields` — `all` or a comma-separated list of member [fields](/cloud/trello/guides/rest-api/object-definitions/)
- `--memberCreator` — Whether to include the member object of the creator
- `--memberCreator_fields` — `all` or a comma-separated list of member [fields](/cloud/trello/guides/rest-api/object-definitions/)
- `--organization` — Whether to include the organization object
- `--organization_fields` — `all` or a comma-separated list of organization [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `notifications get-notifications-id-board`

`GET /notifications/{id}/board`

Get the Board a Notification is on

```bash
trello-cli notifications get-notifications-id-board <id>
```

Path arguments:

- `<id>` — The ID of the notification

Query flags:

- `--fields` — `all` or a comma-separated list of board[fields](/cloud/trello/guides/rest-api/object-definitions/)

### `notifications get-notifications-id-card`

`GET /notifications/{id}/card`

Get the Card a Notification is on

```bash
trello-cli notifications get-notifications-id-card <id>
```

Path arguments:

- `<id>` — The ID of the notification

Query flags:

- `--fields` — `all` or a comma-separated list of card [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `notifications get-notifications-id-field`

`GET /notifications/{id}/{field}`

Get a field of a Notification

```bash
trello-cli notifications get-notifications-id-field <id> <field>
```

Path arguments:

- `<id>` — The ID of the notification
- `<field>` — A notification [field](/cloud/trello/guides/rest-api/object-definitions/)

### `notifications get-notifications-id-list`

`GET /notifications/{id}/list`

Get the List a Notification is on

```bash
trello-cli notifications get-notifications-id-list <id>
```

Path arguments:

- `<id>` — The ID of the notification

Query flags:

- `--fields` — `all` or a comma-separated list of list [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `notifications get-notifications-id-membercreator`

`GET /notifications/{id}/memberCreator`

Get the Member who created the Notification

```bash
trello-cli notifications get-notifications-id-membercreator <id>
```

Path arguments:

- `<id>` — The ID of the notification

Query flags:

- `--fields` — `all` or a comma-separated list of member [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `notifications get-notifications-id-organization`

`GET /notifications/{id}/organization`

Get a Notification's associated Organization

```bash
trello-cli notifications get-notifications-id-organization <id>
```

Path arguments:

- `<id>` — The ID of the notification

Query flags:

- `--fields` — `all` or a comma-separated list of organization [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `notifications notificationsidmember`

`GET /notifications/{id}/member`

Get the Member a Notification is about (not the creator)

```bash
trello-cli notifications notificationsidmember <id>
```

Path arguments:

- `<id>` — The ID of the notification

Query flags:

- `--fields` — `all` or a comma-separated list of member [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `notifications post-notifications-all-read`

`POST /notifications/all/read`

Mark all Notifications as read

```bash
trello-cli notifications post-notifications-all-read
```

Query flags:

- `--read` — Boolean to specify whether to mark as read or unread (defaults to `true`, marking as read)
- `--ids` — A comma-seperated list of IDs. Allows specifying an array of notification IDs to change the read state for. This will become useful as we add grouping of notifications to the UI, with a single butt...

### `notifications put-notifications-id`

`PUT /notifications/{id}`

Update a Notification's read status

```bash
trello-cli notifications put-notifications-id <id>
```

Path arguments:

- `<id>` — The ID of the notification

Query flags:

- `--unread` — Whether the notification should be marked as read or not

### `notifications put-notifications-id-unread`

`PUT /notifications/{id}/unread`

Update Notification's read status

```bash
trello-cli notifications put-notifications-id-unread <id>
```

Path arguments:

- `<id>` — The ID of the notification

Query flags:

- `--value` — (no description)

## organizations

26 operations.

### `organizations delete-organizations-id`

`DELETE /organizations/{id}`

Delete an Organization

```bash
trello-cli organizations delete-organizations-id <id>
```

Path arguments:

- `<id>` — The ID or name of the Organization

### `organizations delete-organizations-id-logo`

`DELETE /organizations/{id}/logo`

Delete Logo for Organization

```bash
trello-cli organizations delete-organizations-id-logo <id>
```

Path arguments:

- `<id>` — The ID or name of the organization

### `organizations delete-organizations-id-members`

`DELETE /organizations/{id}/members/{idMember}`

Remove a Member from an Organization

```bash
trello-cli organizations delete-organizations-id-members <id> <idMember>
```

Path arguments:

- `<id>` — The ID or name of the organization
- `<idMember>` — The ID of the Member to remove from the Workspace

### `organizations delete-organizations-id-prefs-associateddomain`

`DELETE /organizations/{id}/prefs/associatedDomain`

Remove the associated Google Apps domain from a Workspace

```bash
trello-cli organizations delete-organizations-id-prefs-associateddomain <id>
```

Path arguments:

- `<id>` — The ID or name of the organization

### `organizations delete-organizations-id-prefs-orginviterestrict`

`DELETE /organizations/{id}/prefs/orgInviteRestrict`

Delete the email domain restriction on who can be invited to the Workspace

```bash
trello-cli organizations delete-organizations-id-prefs-orginviterestrict <id>
```

Path arguments:

- `<id>` — The ID or name of the organization

### `organizations delete-organizations-id-tags-idtag`

`DELETE /organizations/{id}/tags/{idTag}`

Delete an Organization's Tag

```bash
trello-cli organizations delete-organizations-id-tags-idtag <id> <idTag>
```

Path arguments:

- `<id>` — The ID or name of the organization
- `<idTag>` — The ID of the tag to delete

### `organizations get-organizations-id`

`GET /organizations/{id}`

Get an Organization

```bash
trello-cli organizations get-organizations-id <id>
```

Path arguments:

- `<id>` — The ID or name of the Organization

### `organizations get-organizations-id-actions`

`GET /organizations/{id}/actions`

Get Actions for Organization

```bash
trello-cli organizations get-organizations-id-actions <id>
```

Path arguments:

- `<id>` — The ID or name of the organization

### `organizations get-organizations-id-boards`

`GET /organizations/{id}/boards`

Get Boards in an Organization

```bash
trello-cli organizations get-organizations-id-boards <id>
```

Path arguments:

- `<id>` — The ID or name of the organization

Query flags:

- `--filter` — `all` or a comma-separated list of: `open`, `closed`, `members`, `organization`, `public`
- `--fields` — `all` or a comma-separated list of board [fields](/cloud/trello/guides/rest-api/object-definitions/)

### `organizations get-organizations-id-exports`

`GET /organizations/{id}/exports`

Retrieve Organization's Exports

```bash
trello-cli organizations get-organizations-id-exports <id>
```

Path arguments:

- `<id>` — The ID or name of the Workspace

### `organizations get-organizations-id-field`

`GET /organizations/{id}/{field}`

Get field on Organization

```bash
trello-cli organizations get-organizations-id-field <id> <field>
```

Path arguments:

- `<id>` — The ID or name of the organization
- `<field>` — An organization [field](/cloud/trello/guides/rest-api/object-definitions/)

### `organizations get-organizations-id-members`

`GET /organizations/{id}/members`

Get the Members of an Organization

```bash
trello-cli organizations get-organizations-id-members <id>
```

Path arguments:

- `<id>` — The ID or name of the Organization

### `organizations get-organizations-id-memberships`

`GET /organizations/{id}/memberships`

Get Memberships of an Organization

```bash
trello-cli organizations get-organizations-id-memberships <id>
```

Path arguments:

- `<id>` — The ID or name of the organization

Query flags:

- `--filter` — `all` or a comma-separated list of: `active`, `admin`, `deactivated`, `me`, `normal`
- `--member` — Whether to include the Member objects with the Memberships

### `organizations get-organizations-id-memberships-idmembership`

`GET /organizations/{id}/memberships/{idMembership}`

Get a Membership of an Organization

```bash
trello-cli organizations get-organizations-id-memberships-idmembership <id> <idMembership>
```

Path arguments:

- `<id>` — The ID or name of the organization
- `<idMembership>` — The ID of the membership to load

Query flags:

- `--member` — Whether to include the Member object in the response

### `organizations get-organizations-id-newbillableguests-idboard`

`GET /organizations/{id}/newBillableGuests/{idBoard}`

Get Organizations new billable guests

```bash
trello-cli organizations get-organizations-id-newbillableguests-idboard <id> <idBoard>
```

Path arguments:

- `<id>` — The ID or name of the organization
- `<idBoard>` — The ID of the board to check for new billable guests.

### `organizations get-organizations-id-plugindata`

`GET /organizations/{id}/pluginData`

Get the pluginData Scoped to Organization

```bash
trello-cli organizations get-organizations-id-plugindata <id>
```

Path arguments:

- `<id>` — The ID or name of the organization

### `organizations get-organizations-id-tags`

`GET /organizations/{id}/tags`

Get Tags of an Organization

```bash
trello-cli organizations get-organizations-id-tags <id>
```

Path arguments:

- `<id>` — The ID or name of the Organization

### `organizations organizations-id-members-idmember-all`

`DELETE /organizations/{id}/members/{idMember}/all`

Remove a Member from an Organization and all Organization Boards

```bash
trello-cli organizations organizations-id-members-idmember-all <id> <idMember>
```

Path arguments:

- `<id>` — The ID or name of the organization
- `<idMember>` — The ID of the member to remove from the Workspace

### `organizations post-organizations`

`POST /organizations`

Create a new Organization

```bash
trello-cli organizations post-organizations
```

Query flags:

- `--displayName` — The name to display for the Organization
- `--desc` — The description for the organizations
- `--name` — A string with a length of at least 3. Only lowercase letters, underscores, and numbers are allowed. If the name contains invalid characters, they will be removed. If the name conflicts with an exis...
- `--website` — A URL starting with `http://` or `https://`

### `organizations post-organizations-id-exports`

`POST /organizations/{id}/exports`

Create Export for Organizations

```bash
trello-cli organizations post-organizations-id-exports <id>
```

Path arguments:

- `<id>` — The ID or name of the Workspace

Query flags:

- `--attachments` — Whether the CSV should include attachments or not.

### `organizations post-organizations-id-logo`

`POST /organizations/{id}/logo`

Update logo for an Organization

```bash
trello-cli organizations post-organizations-id-logo <id>
```

Path arguments:

- `<id>` — The ID or name of the Workspace

Query flags:

- `--file` — Image file for the logo

### `organizations post-organizations-id-tags`

`POST /organizations/{id}/tags`

Create a Tag in Organization

```bash
trello-cli organizations post-organizations-id-tags <id>
```

Path arguments:

- `<id>` — The ID or name of the Organization

### `organizations put-organizations-id`

`PUT /organizations/{id}`

Update an Organization

```bash
trello-cli organizations put-organizations-id <id>
```

Path arguments:

- `<id>` — The ID or name of the Organization

Query flags:

- `--name` — A new name for the organization. At least 3 lowercase letters, underscores, and numbers. Must be unique
- `--displayName` — A new displayName for the organization. Must be at least 1 character long and not begin or end with a space.
- `--desc` — A new description for the organization
- `--website` — A URL starting with `http://`, `https://`, or `null`
- `--prefs/associatedDomain` — The Google Apps domain to link this org to.
- `--prefs/externalMembersDisabled` — Whether non-workspace members can be added to boards inside the Workspace
- `--prefs/googleAppsVersion` — `1` or `2`
- `--prefs/boardVisibilityRestrict/org` — Who on the Workspace can make Workspace visible boards. One of `admin`, `none`, `org`
- `--prefs/boardVisibilityRestrict/private` — Who can make private boards. One of: `admin`, `none`, `org`
- `--prefs/boardVisibilityRestrict/public` — Who on the Workspace can make public boards. One of: `admin`, `none`, `org`
- `--prefs/orgInviteRestrict` — An email address with optional wildcard characters. (E.g. `subdomain.*.trello.com`)
- `--prefs/permissionLevel` — Whether the Workspace page is publicly visible. One of: `private`, `public`

### `organizations put-organizations-id-members`

`PUT /organizations/{id}/members`

Update an Organization's Members

```bash
trello-cli organizations put-organizations-id-members <id>
```

Path arguments:

- `<id>` — The ID or name of the organization

Query flags:

- `--email` — An email address
- `--fullName` — Name for the member, at least 1 character not beginning or ending with a space
- `--type` — One of: `admin`, `normal`

### `organizations put-organizations-id-members-idmember`

`PUT /organizations/{id}/members/{idMember}`

Update a Member of an Organization

```bash
trello-cli organizations put-organizations-id-members-idmember <id> <idMember>
```

Path arguments:

- `<id>` — The ID or name of the organization
- `<idMember>` — The ID or username of the member to update

Query flags:

- `--type` — One of: `admin`, `normal`

### `organizations put-organizations-id-members-idmember-deactivated`

`PUT /organizations/{id}/members/{idMember}/deactivated`

Deactivate or reactivate a member of an Organization

```bash
trello-cli organizations put-organizations-id-members-idmember-deactivated <id> <idMember>
```

Path arguments:

- `<id>` — The ID or name of the organization
- `<idMember>` — The ID or username of the member to update

Query flags:

- `--value` — (no description)

## plugins

5 operations.

### `plugins get-plugins-id`

`GET /plugins/{id}/`

Get a Plugin

```bash
trello-cli plugins get-plugins-id <id>
```

Path arguments:

- `<id>` — The ID or name of the organization

### `plugins get-plugins-id-compliance-memberprivacy`

`GET /plugins/{id}/compliance/memberPrivacy`

Get Plugin's Member privacy compliance

```bash
trello-cli plugins get-plugins-id-compliance-memberprivacy <id>
```

Path arguments:

- `<id>` — The ID of the Power-Up

### `plugins post-plugins-idplugin-listing`

`POST /plugins/{idPlugin}/listing`

Create a Listing for Plugin

```bash
trello-cli plugins post-plugins-idplugin-listing <idPlugin>
```

Path arguments:

- `<idPlugin>` — The ID of the Power-Up for which you are creating a new listing.

Body: `--data <json|@file>` (optional JSON request body).

### `plugins put-plugins-id`

`PUT /plugins/{id}/`

Update a Plugin

```bash
trello-cli plugins put-plugins-id <id>
```

Path arguments:

- `<id>` — The ID or name of the organization

### `plugins put-plugins-idplugin-listings-idlisting`

`PUT /plugins/{idPlugin}/listings/{idListing}`

Updating Plugin's Listing

```bash
trello-cli plugins put-plugins-idplugin-listings-idlisting <idPlugin> <idListing>
```

Path arguments:

- `<idPlugin>` — The ID of the Power-Up whose listing is being updated.
- `<idListing>` — The ID of the existing listing for the Power-Up that is being updated.

Body: `--data <json|@file>` (optional JSON request body).

## search

2 operations.

### `search get-search`

`GET /search`

Search Trello

```bash
trello-cli search get-search
```

Query flags:

- `--query` — The search query with a length of 1 to 16384 characters
- `--idBoards` — `mine` or a comma-separated list of Board IDs
- `--idOrganizations` — A comma-separated list of Organization IDs
- `--idCards` — A comma-separated list of Card IDs
- `--modelTypes` — What type or types of Trello objects you want to search. all or a comma-separated list of: `actions`, `boards`, `cards`, `members`, `organizations`
- `--board_fields` — all or a comma-separated list of: `closed`, `dateLastActivity`, `dateLastView`, `desc`, `descData`, `idOrganization`, `invitations`, `invited`, `labelNames`, `memberships`, `name`, `pinned`, `power...
- `--boards_limit` — The maximum number of boards returned. Maximum: 1000
- `--board_organization` — Whether to include the parent organization with board results
- `--card_fields` — all or a comma-separated list of: `badges`, `checkItemStates`, `closed`, `dateLastActivity`, `desc`, `descData`, `due`, `idAttachmentCover`, `idBoard`, `idChecklists`, `idLabels`, `idList`, `idMemb...
- `--cards_limit` — The maximum number of cards to return. Maximum: 1000
- `--cards_page` — The page of results for cards. Maximum: 100
- `--card_board` — Whether to include the parent board with card results
- `--card_list` — Whether to include the parent list with card results
- `--card_members` — Whether to include member objects with card results
- `--card_stickers` — Whether to include sticker objects with card results
- `--card_attachments` — Whether to include attachment objects with card results. A boolean value (true or false) or cover for only card cover attachments.
- `--organization_fields` — all or a comma-separated list of billableMemberCount, desc, descData, displayName, idBoards, invitations, invited, logoHash, memberships, name, powerUps, prefs, premiumFeatures, products, url, website
- `--organizations_limit` — The maximum number of Workspaces to return. Maximum 1000
- `--member_fields` — all or a comma-separated list of: avatarHash, bio, bioData, confirmed, fullName, idPremOrgsAdmin, initials, memberType, products, status, url, username
- `--members_limit` — The maximum number of members to return. Maximum 1000
- `--partial` — By default, Trello searches for each word in your query against exactly matching words within Member content. Specifying partial to be true means that we will look for content that starts with any ...

### `search get-search-members`

`GET /search/members/`

Search for Members

```bash
trello-cli search get-search-members
```

Query flags:

- `--query` — Search query 1 to 16384 characters long
- `--limit` — The maximum number of results to return. Maximum of 20.
- `--idBoard` — (no description)
- `--idOrganization` — (no description)
- `--onlyOrgMembers` — (no description)

## tokens

8 operations.

### `tokens delete-token`

`DELETE /tokens/{token}/`

Delete a Token

```bash
trello-cli tokens delete-token <token>
```

Path arguments:

- `<token>` — (no description)

### `tokens delete-tokens-token-webhooks-idwebhook`

`DELETE /tokens/{token}/webhooks/{idWebhook}`

Delete a Webhook created by Token

```bash
trello-cli tokens delete-tokens-token-webhooks-idwebhook <token> <idWebhook>
```

Path arguments:

- `<token>` — (no description)
- `<idWebhook>` — ID of the [Webhooks](ref:webhooks) to retrieve.

### `tokens get-tokens-token`

`GET /tokens/{token}`

Get a Token

```bash
trello-cli tokens get-tokens-token <token>
```

Path arguments:

- `<token>` — (no description)

Query flags:

- `--fields` — `all` or a comma-separated list of `dateCreated`, `dateExpires`, `idMember`, `identifier`, `permissions`
- `--webhooks` — Determines whether to include webhooks.

### `tokens get-tokens-token-member`

`GET /tokens/{token}/member`

Get Token's Member

```bash
trello-cli tokens get-tokens-token-member <token>
```

Path arguments:

- `<token>` — (no description)

Query flags:

- `--fields` — `all` or a comma-separated list of valid fields for [Member Object](/cloud/trello/guides/rest-api/object-definitions/).

### `tokens get-tokens-token-webhooks`

`GET /tokens/{token}/webhooks`

Get Webhooks for Token

```bash
trello-cli tokens get-tokens-token-webhooks <token>
```

Path arguments:

- `<token>` — (no description)

### `tokens get-tokens-token-webhooks-idwebhook`

`GET /tokens/{token}/webhooks/{idWebhook}`

Get a Webhook belonging to a Token

```bash
trello-cli tokens get-tokens-token-webhooks-idwebhook <token> <idWebhook>
```

Path arguments:

- `<token>` — (no description)
- `<idWebhook>` — ID of the [Webhooks](ref:webhooks) to retrieve.

### `tokens post-tokens-token-webhooks`

`POST /tokens/{token}/webhooks`

Create Webhooks for Token

```bash
trello-cli tokens post-tokens-token-webhooks <token>
```

Path arguments:

- `<token>` — (no description)

Query flags:

- `--description` — A description to be displayed when retrieving information about the webhook.
- `--callbackURL` — The URL that the webhook should POST information to.
- `--idModel` — ID of the object to create a webhook on.

### `tokens tokenstokenwebhooks-1`

`PUT /tokens/{token}/webhooks/{idWebhook}`

Update a Webhook created by Token

```bash
trello-cli tokens tokenstokenwebhooks-1 <token> <idWebhook>
```

Path arguments:

- `<token>` — (no description)
- `<idWebhook>` — ID of the [Webhooks](ref:webhooks) to retrieve.

Query flags:

- `--description` — A description to be displayed when retrieving information about the webhook.
- `--callbackURL` — The URL that the webhook should `POST` information to.
- `--idModel` — ID of the object that the webhook is on.

## webhooks

5 operations.

### `webhooks delete-webhooks-id`

`DELETE /webhooks/{id}`

Delete a Webhook

```bash
trello-cli webhooks delete-webhooks-id <id>
```

Path arguments:

- `<id>` — ID of the webhook to retrieve.

### `webhooks get-webhooks-id`

`GET /webhooks/{id}`

Get a Webhook

```bash
trello-cli webhooks get-webhooks-id <id>
```

Path arguments:

- `<id>` — ID of the webhook to retrieve.

### `webhooks post-webhooks`

`POST /webhooks/`

Create a Webhook

```bash
trello-cli webhooks post-webhooks
```

Query flags:

- `--description` — A string with a length from `0` to `16384`.
- `--callbackURL` — A valid URL that is reachable with a `HEAD` and `POST` request.
- `--idModel` — ID of the model to be monitored
- `--active` — Determines whether the webhook is active and sending `POST` requests.

### `webhooks put-webhooks-id`

`PUT /webhooks/{id}`

Update a Webhook

```bash
trello-cli webhooks put-webhooks-id <id>
```

Path arguments:

- `<id>` — ID of the webhook to retrieve.

Query flags:

- `--description` — A string with a length from `0` to `16384`.
- `--callbackURL` — A valid URL that is reachable with a `HEAD` and `POST` request.
- `--idModel` — ID of the model to be monitored
- `--active` — Determines whether the webhook is active and sending `POST` requests.

### `webhooks webhooksidfield`

`GET /webhooks/{id}/{field}`

Get a field on a Webhook

```bash
trello-cli webhooks webhooksidfield <id> <field>
```

Path arguments:

- `<id>` — ID of the webhook.
- `<field>` — Field to retrieve. One of: `active`, `callbackURL`, `description`, `idModel`

