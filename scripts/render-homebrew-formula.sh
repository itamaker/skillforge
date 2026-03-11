#!/usr/bin/env bash

set -euo pipefail

PROJECT="skillforge"
CLASS_NAME="Skillforge"
DESCRIPTION="Scaffold OpenClaw-ready skills from a compact JSON spec"
CHECKSUMS_FILE="dist/checksums.txt"
OWNER=""
REPO="${PROJECT}"
VERSION=""

usage() {
  echo "Render a formula for https://github.com/itamaker/homebrew-tap" >&2
  echo "Usage: $0 --owner <project-repo-owner> --version <v0.1.0> [--repo <project-repo-name>] [--checksums <path>]" >&2
}

while [[ $# -gt 0 ]]; do
  case "$1" in
    --owner)
      OWNER="$2"
      shift 2
      ;;
    --repo)
      REPO="$2"
      shift 2
      ;;
    --version)
      VERSION="${2#v}"
      shift 2
      ;;
    --checksums)
      CHECKSUMS_FILE="$2"
      shift 2
      ;;
    *)
      usage
      exit 1
      ;;
  esac
done

if [[ -z "${OWNER}" || -z "${VERSION}" ]]; then
  usage
  exit 1
fi

if [[ ! -f "${CHECKSUMS_FILE}" ]]; then
  echo "checksums file not found: ${CHECKSUMS_FILE}" >&2
  exit 1
fi

checksum_for() {
  local os="$1"
  local arch="$2"
  local extension="tar.gz"
  local artifact="${PROJECT}_${VERSION}_${os}_${arch}.${extension}"

  awk -v artifact="${artifact}" '$2 == artifact { print $1 }' "${CHECKSUMS_FILE}"
}

darwin_arm64="$(checksum_for darwin arm64)"
darwin_amd64="$(checksum_for darwin amd64)"
linux_arm64="$(checksum_for linux arm64)"
linux_amd64="$(checksum_for linux amd64)"

for checksum in "${darwin_arm64}" "${darwin_amd64}" "${linux_arm64}" "${linux_amd64}"; do
  if [[ -z "${checksum}" ]]; then
    echo "missing required checksum in ${CHECKSUMS_FILE}" >&2
    exit 1
  fi
done

cat <<EOF
class ${CLASS_NAME} < Formula
  desc "${DESCRIPTION}"
  homepage "https://github.com/${OWNER}/${REPO}"
  version "${VERSION}"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/${OWNER}/${REPO}/releases/download/v${VERSION}/${PROJECT}_${VERSION}_darwin_arm64.tar.gz"
      sha256 "${darwin_arm64}"
    else
      url "https://github.com/${OWNER}/${REPO}/releases/download/v${VERSION}/${PROJECT}_${VERSION}_darwin_amd64.tar.gz"
      sha256 "${darwin_amd64}"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/${OWNER}/${REPO}/releases/download/v${VERSION}/${PROJECT}_${VERSION}_linux_arm64.tar.gz"
      sha256 "${linux_arm64}"
    else
      url "https://github.com/${OWNER}/${REPO}/releases/download/v${VERSION}/${PROJECT}_${VERSION}_linux_amd64.tar.gz"
      sha256 "${linux_amd64}"
    end
  end

  def install
    bin.install "${PROJECT}"
  end

  test do
    system "#{bin}/${PROJECT}", "--help"
  end
end
EOF
