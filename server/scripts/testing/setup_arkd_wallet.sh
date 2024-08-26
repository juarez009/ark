#!/bin/bash

# Check if password is provided as an argument
if [ -z "$1" ]; then
  echo "Usage: $0 <password>"
  exit 1
fi

PASSWORD="$1"

# Authorization header
AUTH_HEADER="Authorization: Basic YWRtaW46YWRtaW4="

# Function to make a GET request
make_get_request() {
    local url="$1"
    curl -s -X GET "$url" -H "$AUTH_HEADER"
}

# Function to make a POST request
make_post_request() {
    local url="$1"
    local data="$2"
    curl -s -X POST "$url" -H "$AUTH_HEADER" -H "Content-Type: application/json" -d "$data"
}

# Step 1: Generate Seed
seed_response=$(make_get_request "http://localhost:6060/v1/admin/wallet/seed")

# Extract seed value using jq
seed=$(echo "$seed_response" | jq -r '.seed')

if [ -z "$seed" ]; then
    echo "Failed to generate seed"
    exit 1
fi

# Print the generated seed
echo "Generated wallet seed: $seed"

# Step 2: Create Wallet
create_wallet_response=$(make_post_request "http://localhost:6060/v1/admin/wallet/create" \
  "{\"seed\": \"$seed\", \"password\": \"$PASSWORD\"}")

# Step 3: Unlock Wallet
unlock_wallet_response=$(make_post_request "http://localhost:6060/v1/admin/wallet/unlock" \
  "{\"password\": \"$PASSWORD\"}")

# Step 4: Wait for a second
sleep 1

# Step 5: Get New Address
address_response=$(make_get_request "http://localhost:6060/v1/admin/wallet/address")

# Extract address value using jq
address=$(echo "$address_response" | jq -r '.address')

if [ -z "$address" ]; then
    echo "Failed to get new address"
    exit 1
fi

# Print the generated address
echo "Generated wallet address: $address"

# Step 6: Fund Wallet using Nigiri Faucet
nigiri faucet --liquid "$address"
nigiri faucet --liquid "$address"
nigiri faucet --liquid "$address"

echo "Script execution completed successfully."