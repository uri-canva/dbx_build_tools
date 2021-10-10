#!/usr/bin/env bash
set -eu

# Use a var to avoid a pipe - that would make number of processes nondeterministic
out=$(ps ux)
procs1=$(wc -l <<< out)
$RUNFILES/../dbx_build_tools/build_tools/svcctl/cmd/svcctl/svcctl_/svcctl stop-all
$RUNFILES/../dbx_build_tools/build_tools/svcctl/cmd/svcctl/svcctl_/svcctl start-all
out=$(ps ux)
procs2=$(wc -l <<< out)
if [ "$procs1" -lt "$procs2" ]; then
  echo "Restarting all services leaked some processes. Before: $procs1\nAfter:$procs2"
  exit -1
fi
