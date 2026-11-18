#!/bin/bash
# generated with CMTools 0.0.1 7534e43

#
# Release script for fdx on GitHub using gh cli.
#
# Usage: release.bash [--dry-run]
#   --dry-run  do everything up to the prompt (notes, checksums), say what
#              would be uploaded and from which commit, and change nothing.
#
# The release is made from the commit HEAD (workspace DR-0012). The script
# commits and pushes nothing: it stops unless the tree is committed and HEAD is
# pushed (make check-committed), and creates the draft with --target HEAD.
#
DRY_RUN=0
for ARG in "$@"; do
    case "${ARG}" in
        -n|--dry-run) DRY_RUN=1 ;;
        -h|--help) sed -n '2,/^DRY_RUN/p' "$0" | grep '^#'; exit 0 ;;
        *) echo "usage: release.bash [--dry-run]" >&2; exit 2 ;;
    esac
done

REPO_ID="${PWD##*/}"
GROUP_ID="$(git config --get remote.origin.url | sed -E 's#.*github\.com[:/]([^/]+)/.*#\1#')"
REPO_URL="https://github.com/${GROUP_ID}/${REPO_ID}"
echo "REPO_URL -> ${REPO_URL}"

#
# Generate a new draft release using jq and gh
#
RELEASE_TAG="v$(jq -r .version codemeta.json)"
if ! printf '%s' "${RELEASE_TAG}" | grep -qE '^v[0-9a-zA-Z._-]+$'; then
    echo "error: version contains unexpected characters: ${RELEASE_TAG}"
    exit 1
fi
echo "tag: ${RELEASE_TAG}, notes:"
jq -r .releaseNotes codemeta.json >release_notes.tmp
cat release_notes.tmp

#
# Generate checksums for distribution zip files
#
CHECKSUM_FILE="${REPO_ID}-${RELEASE_TAG}-checksums.txt"
if command -v sha256sum >/dev/null 2>&1; then
    sha256sum dist/*.zip | sed 's|dist/||' > "dist/${CHECKSUM_FILE}"
else
    shasum -a 256 dist/*.zip | sed 's|dist/||' > "dist/${CHECKSUM_FILE}"
fi
echo "Checksums written to dist/${CHECKSUM_FILE}"

# The release is built from one known commit. Stop unless it is committed and
# pushed, and do not create a second draft for the tag.
if ! make -s check-committed; then
    rm -f release_notes.tmp
    exit 1
fi
RELEASE_COMMIT="$(git rev-parse HEAD)"
if gh release view "${RELEASE_TAG}" >/dev/null 2>&1; then
    echo "error: a release for ${RELEASE_TAG} already exists on GitHub; delete the draft first"
    rm -f release_notes.tmp
    exit 1
fi
if [ "${DRY_RUN}" = "1" ]; then
    echo "dry run: would create draft ${RELEASE_TAG} at commit ${RELEASE_COMMIT} and upload:"
    ls dist/*.zip "dist/${CHECKSUM_FILE}"
    rm -f release_notes.tmp
    exit 0
fi
# shellcheck disable=SC2162
read -r -p "Create the draft release on GitHub with gh? (y/N) " YES_NO
if [ "${YES_NO}" = "y" ]; then
	# Now generate a draft release
	echo "Pushing release ${RELEASE_TAG} to GitHub"
	gh release create "${RELEASE_TAG}" \
		--draft \
		--target "${RELEASE_COMMIT}" \
		-F release_notes.tmp \
		--generate-notes
	echo "Uploading distribution files and checksums"
	echo "Starting upload: dist/${CHECKSUM_FILE}
	"
	gh release upload "${RELEASE_TAG}" "dist/${CHECKSUM_FILE}"
	echo "  Completed upload: dist/${CHECKSUM_FILE}"
	for FILE in dist/*.zip; do
		echo "Starting upload: ${FILE}
	"
		gh release upload "${RELEASE_TAG}" "${FILE}"
		echo "  Completed upload: ${FILE}"
	done

	cat <<EOT

Now goto repo release and finalize draft.

	${REPO_URL}/releases

EOT
	rm release_notes.tmp

fi
