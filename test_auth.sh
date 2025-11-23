#!/bin/bash

echo "=== Testing Authentication Flow ==="

# Test register
echo -e "\n1. Testing registration..."
REGISTER_RESPONSE=$(curl -s -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@test.com","password":"testpass"}')
echo "Register response: $REGISTER_RESPONSE"

# Test login
echo -e "\n2. Testing login..."
LOGIN_RESPONSE=$(curl -s -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"testpass"}')
echo "Login response: $LOGIN_RESPONSE"

# Extract token
TOKEN=$(echo $LOGIN_RESPONSE | grep -o '"token":"[^"]*' | cut -d'"' -f4)
USER_ID=$(echo $LOGIN_RESPONSE | grep -o '"user_id":[^,}]*' | cut -d':' -f2)

echo -e "\nExtracted token: $TOKEN"
echo "Extracted user_id: $USER_ID"

# Test protected endpoint with token
echo -e "\n3. Testing protected endpoint with token..."
STATS_RESPONSE=$(curl -s http://localhost:8080/api/user/$USER_ID/stats \
  -H "Authorization: $TOKEN")
echo "Stats response: $STATS_RESPONSE"

echo -e "\n=== Test Complete ==="
