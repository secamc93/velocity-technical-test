BASE_URL="http://localhost:60000"
ORDER_ID=1

curl -H "Accept: application/json" "${BASE_URL}/api/orders/${ORDER_ID}" | jq
