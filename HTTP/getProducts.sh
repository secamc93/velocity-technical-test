BASE_URL="http://localhost:60000"
# ...existing code...
curl -H "Accept: application/json" "${BASE_URL}/api/products" | jq
# ...existing code...
