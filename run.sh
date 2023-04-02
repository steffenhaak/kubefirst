#!/bin/zsh
go run . civo create --alerts-email haak@aucentiq.com  \
  --github-org steffenhaak \
  --domain-name getreactive.de \
  --cluster-name gitops \
  --cloud-region FRA1