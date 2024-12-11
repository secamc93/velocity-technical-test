BASE_URL="http://localhost:60000"


curl -X POST ${BASE_URL}/orders \
-H "Content-Type: application/json" \
-d '{
	"id": "1"
	"customer_name": "John Doe",
	"total_amount": 100.50
}'
