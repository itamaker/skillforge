# skillforge

`skillforge` scaffolds OpenClaw-ready skill folders from a compact JSON spec.

## Example

```bash
go run . init -spec examples/skill.json -out /tmp/research-skill
```

The command creates:

- `SKILL.md`
- `manifest.json`
- `bin/README.md`
- `examples/usage.md`

## Why it is useful

- Standardizes skill authoring for agent teams.
- Speeds up internal tooling rollout without building a full UI.
- Produces portable folders that can be moved into OpenClaw-style workspaces.

## Install

From source:

```bash
go install github.com/YOUR_GITHUB_USER/skillforge@latest
```

From Homebrew after you publish a tap formula:

```bash
brew tap itamaker/tap https://github.com/itamaker/homebrew-tap
brew install itamaker/tap/skillforge
```

## Repo-Ready Files

- `.github/workflows/ci.yml`
- `.github/workflows/release.yml`
- `.goreleaser.yaml`
- `PUBLISHING.md`
- `scripts/render-homebrew-formula.sh`

## Release

```bash
git tag v0.1.0
git push origin v0.1.0
```

The tagged release workflow publishes multi-platform binaries and `checksums.txt`, which you can feed into the Homebrew formula renderer.
The generated formula should be committed to `https://github.com/itamaker/homebrew-tap`.
