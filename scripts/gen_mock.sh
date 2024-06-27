#!/bin/bash

if [ -z "$BASH" ]; then echo "Please run this script with bash"; exit 1; fi

SCRIPT_PATH=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
ROOT_PATH="$(cd $SCRIPT_PATH && cd ../internal && pwd)"
MOCKS_PATH="$ROOT_PATH/.gen/mock"

echo $ROOT_PATH
echo $MOCKS_PATH

mkdir -p "$MOCKS_PATH"

# mocks that interests us
MOCKS_TO_KEEP=(pkg domain)

echo "clearing existing mocks..."
find "$MOCKS_PATH" -mindepth 1 -maxdepth 1 -type d -exec rm -rv {} \;

echo "generating project mocks..."
mockery --all --keeptree --with-expecter --dir "$ROOT_PATH/domain" --output "$MOCKS_PATH"

#for mock in ${MOCKS_TO_KEEP[@]}; do
#  mv $(find "$MOCKS_PATH/internal" -maxdepth 2 -type d -name "$mock") "$MOCKS_PATH/"
#done
#rm -fr "$MOCKS_PATH/internal"

echo "done."