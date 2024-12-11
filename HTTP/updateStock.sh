# ...existing code...

BASE_URL="http://localhost:60000"
NEW_STOCK=5
PRODUCT_ID=20

curl -X PUT "${BASE_URL}/api/products/${PRODUCT_ID}/stock" \
     -H "Content-Type: application/json" \
     -d "{\"new_stock\": ${NEW_STOCK}}" | jq

# ...existing code...
