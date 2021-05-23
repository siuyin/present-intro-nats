# Intro to NATS

## Presentation
present-intro-nats.slide

## Kubernetes manifests
ns.yaml:
- ns siuyin

pv.yaml:
- storage class nats-storage
- pv nats1 nats2 nats2
- pvc nats1 nats2 nats2

bb.yaml:
- deploy bb # busybox to test pv usage

nats.yaml:
- deploy nats # with jetstream cluster
