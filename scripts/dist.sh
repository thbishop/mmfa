#!/bin/bash
set -e

# borrowed from: https://github.com/mitchellh/packer

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

PROJECT=ec2_metadata_dump

# Change into that dir because we expect that
cd $DIR

# Get the git commit
GIT_COMMIT=$(git rev-parse HEAD)
GIT_DIRTY=$(test -n "`git status --porcelain`" && echo "+CHANGES" || true)

# Determine the version that we're building based on the contents
# of packer/version.go.
VERSION=$(grep "const Version " version.go | sed -E 's/.*"(.+)"$/\1/')
VERSIONDIR="${VERSION}"
PREVERSION=$(grep "const VersionPrerelease " version.go | sed -E 's/.*"(.*)"$/\1/')
if [ ! -z $PREVERSION ]; then
    PREVERSION="${PREVERSION}.$(date -u +%s)"
    VERSIONDIR="${VERSIONDIR}-${PREVERSION}"
fi

echo "Version: ${VERSION} ${PREVERSION}"

# Determine the arch/os combos we're building for
XC_ARCH=${XC_ARCH:-"386 amd64"}
XC_OS=${XC_OS:-linux}

echo "Arch: ${XC_ARCH}"
echo "OS: ${XC_OS}"

# Make sure that if we're killed, we kill all our subprocseses
trap "kill 0" SIGINT SIGTERM EXIT

# This function builds whatever directory we're in...
rm -fr pkg
mkdir -p pkg
gox \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -ldflags "-X github.com/thbishop/$PROJECT/$PROJECT ${GIT_COMMIT}${GIT_DIRTY}" \
    -output "pkg/{{.OS}}_{{.Arch}}/$PROJECT-{{.Dir}}" \
    ./...

# Make sure it is renamed properly
for PLATFORM in $(find ./pkg -mindepth 1 -maxdepth 1 -type d); do
    set +e
    mv ${PLATFORM}/$PROJECT-$PROJECT ${PLATFORM}/$PROJECT 2>/dev/null
    mv ${PLATFORM}/$PROJECT-$PROJECT.exe ${PLATFORM}/$PROJECT.exe 2>/dev/null
    set -e
done

# Zip all the packages
mkdir -p ./pkg/dist
for PLATFORM in $(find ./pkg -mindepth 1 -maxdepth 1 -type d); do
    PLATFORM_NAME=$(basename ${PLATFORM})
    ARCHIVE_NAME="${VERSIONDIR}_${PLATFORM_NAME}"

    if [ $PLATFORM_NAME = "dist" ]; then
        continue
    fi

    pushd ${PLATFORM}
    zip ${DIR}/pkg/dist/${ARCHIVE_NAME}.zip ./*
    popd
done

# # Make the checksums
# pushd ./pkg/${VERSIONDIR}/dist
pushd ./pkg/dist
shasum -a256 * > ./${VERSIONDIR}_SHA256SUMS
popd

exit 0
