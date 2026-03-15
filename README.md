# skillforge

[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)

`skillforge` is a Go CLI that scaffolds OpenClaw-ready skill directories from a compact JSON spec.

It helps agent teams standardize skill packaging without building a separate UI or internal generator service.

![skillforge social preview](docs/images/social-preview.png)

## Support

[![Buy Me A Coffee](https://img.shields.io/badge/Buy%20Me%20A%20Coffee-FFDD00?style=for-the-badge&logo=buy-me-a-coffee&logoColor=black)](https://buymeacoffee.com/amaker)

## Quickstart

### Install

```bash
brew install itamaker/tap/skillforge
```

<details>
<summary>You can also download binaries from <a href="https://github.com/itamaker/skillforge/releases">GitHub Releases</a>.</summary>

Current release archives:

- macOS (Apple Silicon/arm64): `skillforge_0.1.1_darwin_arm64.tar.gz`
- macOS (Intel/x86_64): `skillforge_0.1.1_darwin_amd64.tar.gz`
- Linux (arm64): `skillforge_0.1.1_linux_arm64.tar.gz`
- Linux (x86_64): `skillforge_0.1.1_linux_amd64.tar.gz`

Each archive contains a single executable: `skillforge`.

</details>

### First Run

Run:

```bash
skillforge
```

This launches the interactive Bubble Tea terminal UI.

You can still use the direct command form:

```bash
skillforge init -spec examples/skill.json -out /tmp/research-skill
```

Or draft a richer spec from a short brief plus a tool catalog:

```bash
skillforge draft -brief examples/brief.md -catalog examples/tools.json -out /tmp/research-skill.json
```

The generated directory includes:

- `SKILL.md`
- `manifest.json`
- `bin/README.md`
- `examples/usage.md`

## Requirements

- Go `1.22+`

## Run

```bash
go run .
```

Or run the non-interactive command directly:

```bash
go run . init -spec examples/skill.json -out /tmp/research-skill
```

Use `-force` if you want to overwrite an existing output directory.

Generate a spec first:

```bash
go run . draft -brief examples/brief.md -catalog examples/tools.json
```

## Build From Source

```bash
make build
```

```bash
go build -o dist/skillforge .
```

## What It Does

1. Loads a compact JSON skill spec.
2. Validates required fields such as `name`, `description`, and `tools`.
3. Can draft a richer spec from a natural-language brief by retrieving relevant tools from a local catalog.
4. Generates a portable skill folder with workflow, constraints, checks, docs, and manifest files.
5. Produces output that can be moved into OpenClaw-style agent workspaces.

## Notes

- Use `examples/skill.json` as a starting point for new skill definitions.
- Maintainer release steps live in `PUBLISHING.md`.

## Contributors ✨

| [![Zhaoyang Jia][avatar-zhaoyang]][author-zhaoyang] |
| --- |
| [Zhaoyang Jia][author-zhaoyang] |



[author-zhaoyang]: https://github.com/itamaker
[avatar-zhaoyang]: https://images.weserv.nl/?url=https://github.com/itamaker.png&h=120&w=120&fit=cover&mask=circle&maxage=7d

## License

[MIT](LICENSE)
