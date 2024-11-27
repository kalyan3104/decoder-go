import requests
import json

# Function to read addresses from a file
def read_addresses_from_file(file_path):
    with open(file_path, 'r') as file:
        return [line.strip() for line in file.readlines()]

# Function to check if the address has an owner and print it
def check_account_owner(addr):
    url = f"https://api.elrond.com/accounts/{addr}"
    
    # Make the request
    response = requests.get(url)
    
    # Check if the request was successful
    if response.status_code == 200:
        data = response.json()
        
        # Print the entire response for debugging purposes
        print(json.dumps(data, indent=4))
        
        # Assuming the 'owner' field is present in the response, adjust as necessary
        owner = data.get('owner', 'N/A')  # Default to 'N/A' if no owner is found
        
        # Check for address with owner information
        if owner != 'N/A':
            print(f"Address: {addr} (Shard 1)\nOwner: {owner}\n")
        else:
            print(f"Address: {addr} (Shard 1)\nOwner: N/A\n")
    else:
        print(f"Failed to query address {addr}: {response.status_code} - {response.text}")

# Read addresses from the file
addresses = read_addresses_from_file('addresses.txt')

# Iterate over each address and check for owners
for addr in addresses:
    check_account_owner(addr)
