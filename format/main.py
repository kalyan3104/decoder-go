import requests
import json
import pandas as pd


def read_addresses_from_file(file_path):
    with open(file_path, 'r') as file:
        return [line.strip() for line in file.readlines()]


def check_account_owner(addr):
    url = f"https://api.elrond.com/accounts/{addr}"
    
    
    response = requests.get(url)
    
    
    if response.status_code == 200:
        data = response.json()
        
      
        owner = data.get('ownerAddress', 'N/A')  
        
        
        return {'Address': addr, 'Owner Address': owner}
    else:
        print(f"Failed to query address {addr}: {response.status_code} - {response.text}")
        return {'Address': addr, 'Owner Address': 'N/A'}


addresses = read_addresses_from_file('addresses.txt')


results = []


for addr in addresses:
    result = check_account_owner(addr)
    results.append(result)


df = pd.DataFrame(results)


df.to_excel('addresses_with_owners.xlsx', index=False)

print("Excel file has been created successfully!")
