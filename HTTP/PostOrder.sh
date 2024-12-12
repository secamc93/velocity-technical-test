BASE_URL="http://localhost:60000"
IDEMPOTENCY_KEY="123e4567-e89b-12d3-a456-426614174000"

curl -X POST ${BASE_URL}/api/orders \
-H "Content-Type: application/json" \
-H "Idempotency-Key: ${IDEMPOTENCY_KEY}" \
-d '{
  "customer_name": "John Doe",
  "total_amount": 1549.92,
  "items": [
    {
      "product_id": 1,
      "quantity": 2,
      "subtotal": 999.98
    },
    {
      "product_id": 2,
      "quantity": 1,
      "subtotal": 129.99
    },
    {
      "product_id": 3,
      "quantity": 3,
      "subtotal": 239.97
    },
    {
      "product_id": 4,
      "quantity": 5,
      "subtotal": 174.95
    },
    {
      "product_id": 5,
      "quantity": 4,
      "subtotal": 199.96
    }
  ]
}' | jq
