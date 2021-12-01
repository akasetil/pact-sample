#!/bin/bash
set -o pipefail

# Run the FastAPI server, using the pact_provider.py as the app to be able to
# inject the provider_states endpoint
uvicorn tests.pact_provider:app & &>/dev/null
FASTAPI_PID=$!

# Make sure the FastAPI server is stopped when finished to avoid blocking the port
function teardown {
  echo "Tearing down FastAPI server ${FASTAPI_PID}"
  kill -9 $FASTAPI_PID
}
trap teardown EXIT

# Wait a little in case FastAPI isn't quite ready
sleep 1


echo "Validating provider locally"

pact-provider-verifier --provider-base-url=http://127.0.0.1:8000 \
  --provider-states-setup-url=http://127.0.0.1:8000/_pact/provider_states \
  ../pactfile/payment-product.json
