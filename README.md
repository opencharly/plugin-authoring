# plugin-authoring

Programmatic box-manifest authoring for OpenCharly — the `charly box` authoring
verbs, nested under the core `box` command group.

`charly box set/add-candy/rm-candy/write/cat` mutate a project's `charly.yml`
directly via the SDK kit (the comment-preserving `yaml.Node` utilities), with
zero core reentry. `charly box fetch/refresh` pre-prime / force-re-clone the
remote-repo cache by re-running the hidden core `__box-fetch` / `__box-refresh`
reentry over the generic `HostBuild("cli")` reverse channel. The plugin is
compiled into charly (`command:*:box`) and dispatches in-process; placement is
invisible.

## What it provides

| Capability | Surface |
|---|---|
| `command:set:box` | `charly box set <path> <value>` — set a dot-path value (parsed as YAML) |
| `command:add-candy:box` | `charly box add-candy <box> <candy>` |
| `command:rm-candy:box` | `charly box rm-candy <box> <candy>` |
| `command:write:box` | `charly box write <path> [--content X \| --from-stdin]` |
| `command:cat:box` | `charly box cat <path>` |
| `command:fetch:box` | `charly box fetch` — fetch remote `@github` refs |
| `command:refresh:box` | `charly box refresh` — refresh the project's resolved state |

`set`/`add-candy`/`rm-candy`/`write`/`cat` run entirely on `sdk/kit` + stdlib;
`fetch` and `refresh` reach the host-coupled repo resolver.

## How to use it

The commands are compiled in — no candy composition is needed:

```bash
charly box set box.my-box.description "My box"
charly box add-candy my-box layer-ripgrep
charly box write box/my-box/extra.yml --content $'foo: bar\n'
charly box cat box/my-box/charly.yml
charly box fetch
```

## Layout

- `candy/plugin-authoring/` — the plugin module: `plugin.go`, `authoring.go`,
  `authoring_edit.go`, `cmd/serve/main.go`.
- `candy/plugin-authoring/charly.yml` — the `plugin-authoring:` candy entity
  and the embedded `authoring-skill:` skill entity.
- `charly.yml` — the root project manifest (`discover: candy`).
- `.github/workflows/tag-on-merge.yml` — CalVer tag + `CHANGELOG/` on merge.

## Related

- Owning skill: `/charly-authoring:authoring` (projected from the embedded
  `authoring-skill:` entity).
- `/charly-image:image` — box composition and `charly.yml` shapes.
- `/charly-internals:plugin` — the plugin/provider model.
- [`opencharly/charly`](https://github.com/opencharly/charly) — the charly CLI.
