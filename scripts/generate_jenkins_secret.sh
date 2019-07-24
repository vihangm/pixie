#!/bin/bash -e

# This script can be used to generate secret tokens used by Jenkins.
# Any cluster accessed by Jenkins needs credentials to be registered
# in the UI under: https://jenkins.pixielabs.ai/credentials.

namespace=pl
account_name=jenkins-robot

kubectl -n ${namespace} create serviceaccount ${account_name}

# The next line gives `${account_name}` administator permissions for this namespace.
kubectl -n ${namespace} create rolebinding "${account_name}-binding" \
        --clusterrole=cluster-admin --serviceaccount="${namespace}:${account_name}"

# Get the name of the token that was automatically generated for the ServiceAccount `${account_name}`.
token_name=$(kubectl -n ${namespace} get serviceaccount "${account_name}" \
                     -o go-template --template='{{range .secrets}}{{.name}}{{"\n"}}{{end}}')

# Retrieve the token and decode it using base64.
kubectl -n ${namespace} get secrets "${token_name}" -o go-template \
        --template '{{index .data "token"}}' | base64 -d
