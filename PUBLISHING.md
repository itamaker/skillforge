# Publishing

## Repository Setup

1. Create a GitHub repository named `skillforge`.
2. Copy this directory so it becomes the repository root.
3. Push `main`.

## CI

The repository includes `.github/workflows/ci.yml` to run:

- `go test ./...`
- `go build ./...`

## Release

Tagging a semantic version triggers `.github/workflows/release.yml`.

```bash
git tag v0.1.0
git push origin v0.1.0
```

That workflow publishes release archives and `checksums.txt` through GoReleaser.

## Homebrew Tap

After the first release:

```bash
./scripts/render-homebrew-formula.sh --owner itamaker --version v0.1.0 > /path/to/homebrew-tap/Formula/skillforge.rb
```

Commit the rendered file to `https://github.com/itamaker/homebrew-tap` as `Formula/skillforge.rb`, then users can install with:

```bash
brew tap itamaker/tap https://github.com/itamaker/homebrew-tap
brew install itamaker/tap/skillforge
```
