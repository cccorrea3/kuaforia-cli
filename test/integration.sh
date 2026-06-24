#!/bin/sh
# Tests de integración — requieren servidor Kuaforia corriendo
set -e

SERVER=${KUAFORIA_SERVER:-http://localhost:8000}
API_KEY=${KUAFORIA_API_KEY:?KUAFORIA_API_KEY required}
TENANT=${KUAFORIA_TENANT:?KUAFORIA_TENANT required}
CLI=${CLI:-./bin/kuaforia}
PASS=0
FAIL=0

run() {
    desc="$1"
    shift
    if $CLI "$@" --server "$SERVER" --api-key "$API_KEY" --tenant "$TENANT" >/dev/null 2>&1; then
        echo "[PASS] $desc"
        PASS=$((PASS + 1))
    else
        echo "[FAIL] $desc"
        FAIL=$((FAIL + 1))
    fi
}

echo "=== Kuaforia Integration Tests ==="
echo "Server: $SERVER  Tenant: $TENANT"
echo ""

run "Health check" auth test
run "Create case" cases create "Test case integracion" --description "Creado por test de integracion"
run "List cases" cases list
run "Get case 1" cases get 1
run "Update case 1" cases update 1 --title "Test actualizado"
run "Delete case 1" cases delete 1 --force

echo ""
echo "Resultados: $PASS passed, $FAIL failed"
exit $FAIL
