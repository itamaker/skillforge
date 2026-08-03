---
name: skillforge
description: Draft and scaffold OpenClaw-ready agent skill folders with the skillforge Go CLI. Use when you need to draft a compact skill spec (JSON) from a natural-language brief plus an optional tool catalog, scaffold a full skill folder (SKILL.md, manifest.json, bin/README.md, examples/usage.md) from a JSON skill spec, or validate whether a skill spec or tool catalog file has the required shape before it is used. Trigger for requests like "draft a skill from this brief", "turn this brief and tool catalog into a skill spec", "scaffold a skill from this spec", "generate a SKILL.md and manifest from skill.json", or "check whether this skill spec is valid".
---

# skillforge

`skillforge` is a Go CLI (binary name `skillforge`) that drafts and scaffolds OpenClaw-ready agent skill folders. It has two non-interactive subcommands, `draft` and `init`, plus a Bubble Tea TUI that launches when it is run with no arguments. Always use the non-interactive subcommands in an agent context; the TUI requires a live terminal and is not scriptable.

## Workflow

1. Build or locate the `skillforge` binary.
   - From this repo: `go build -o skillforge .` (or `make build`, or `go run . <subcommand> ...` without a separate build step).
   - If installed via Homebrew (`brew install itamaker/tap/skillforge`) or a GitHub release, the `skillforge` binary is already on `PATH`.

2. If you only have a natural-language brief (and optionally a tool catalog), draft a skill spec first:
   - `skillforge draft -brief <brief.md> [-catalog <tools.json>] [-out <spec.json>] [-name "<Explicit Name>"] [-max-tools <N>]`
   - `-brief` is required; everything else is optional.
   - Without `-catalog`, the drafted spec gets a single placeholder tool (`replace-me`) that must be replaced with a real tool before the spec is useful.
   - Without `-out`, the drafted spec JSON prints to stdout instead of being written to a file.
   - `-max-tools` caps how many tools are retrieved from the catalog by keyword overlap with the brief (default `3`).
   - The drafted spec infers `name`, `slug`, `description`, `audience`, `category`, `tags`, `triggers`, `constraints`, `checks`, `examples`, and `workflow` from the brief text — review and edit these before scaffolding, since they are heuristic guesses, not guaranteed-accurate metadata.

3. Scaffold a skill folder from a JSON skill spec (either drafted in step 2, or hand-written to match `examples/skill.json`):
   - `skillforge init -spec <spec.json> -out <output-dir> [-force]`
   - Both `-spec` and `-out` are required.
   - The spec must have a non-empty `name`, a non-empty `description`, and at least one entry in `tools` where each tool has a non-empty `name` and `command`. `init` validates this before writing anything and exits with a clear one-line error (exit code `1`) if validation fails — this is the mechanism for checking whether a skill spec or tool catalog file is well-formed.
   - On success it creates `<output-dir>/SKILL.md`, `<output-dir>/manifest.json`, `<output-dir>/bin/README.md`, and `<output-dir>/examples/usage.md`, and prints `generated skill scaffold in <output-dir>`.

4. Report the generated skill folder path back to the user, and point out any drafted fields (from step 2) that read as generic placeholders worth tightening by hand.

## Safety Rules

- `init` refuses to overwrite an existing `-out` directory unless `-force` is passed; with `-force` it deletes the existing directory contents first (`os.RemoveAll`). Only pass `-force` when the user clearly wants the existing output directory replaced.

## Prompt Patterns

- "Draft a skill spec from this brief and tool catalog."
- "Turn `brief.md` into a skillforge spec, then scaffold it."
- "Scaffold a skill folder from `skill.json` into `/tmp/my-skill`."
- "Regenerate the skill folder at `/tmp/my-skill`, overwriting what's there."
- "Draft a skill for retrieval evaluation using these three tools, keep only the top one."
- "Check whether this skill spec has all the required fields before I scaffold it."

## Resources

- `examples/brief.md`: sample natural-language brief for `draft -brief`.
- `examples/tools.json`: sample tool catalog (a plain array of tool objects) for `draft -catalog`.
- `examples/skill.json`: sample hand-written skill spec for `init -spec`.
- `references/REFERENCE.md`: full CLI flag reference, the `SkillSpec`/`ToolSpec`/`WorkflowStep` JSON schema, the two accepted tool-catalog shapes, and generated-file details.
