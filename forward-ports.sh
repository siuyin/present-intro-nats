#!/bin/sh

cmd="kubectl port-forward deploy/nats 4222"
$cmd 2>&1 >/dev/null &

echo "ps aux|grep 'kubectl port-forward' to find process"
