#!/usr/bin/env bash
set -euox pipefail

# Get protoc executions
go get github.com/regen-network/cosmos-proto/protoc-gen-gocosmos #2>/dev/null

# Get cosmos sdk from github
go get github.com/cosmos/cosmos-sdk@v0.45.11 #2>/dev/null

echo "Linting gogo proto code"
cd proto
proto_dirs=$(find ./ -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq) # ./proto/canine_chain/jklmint
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep go_package $file &>/dev/null; then
      buf format --config buf.yaml $file -w
    fi
  done
done


