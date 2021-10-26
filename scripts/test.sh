#!/bin/bash

echo "Creating new account"

curl --location --request POST 'http://localhost:8080/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "testUser",
    "password": "12345"
}'
