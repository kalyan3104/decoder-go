import aiohttp
import asyncio
import pandas as pd
import time

# Function to read addresses from a file
async def read_addresses_from_file(file_path):
    with open(file_path, 'r') as file:
        return [line.strip() for line in file.readlines()]

# Function to check account owner with retry logic
async def check_account_owner(session, addr, retries=3, delay=10):
    url = f"https://api.elrond.com/accounts/{addr}"
    
    for attempt in range(retries):
        try:
            async with session.get(url) as response:
                if response.status == 200:
                    data = await response.json()
                    owner = data.get('ownerAddress', 'N/A')
                    return {'Address': addr, 'Owner Address': owner}
                elif response.status == 429:
                    # Rate limited, wait before retrying
                    print(f"Rate limited for address {addr}, retrying in {delay} seconds...")
                    await asyncio.sleep(delay)
                    delay *= 2  # Exponential backoff
                else:
                    # Handle other errors
                    print(f"Failed to query address {addr}: {response.status} - {await response.text()}")
                    return {'Address': addr, 'Owner Address': 'N/A'}
        except Exception as e:
            print(f"Error occurred for address {addr}: {str(e)}")
            return {'Address': addr, 'Owner Address': 'N/A'}
    
    return {'Address': addr, 'Owner Address': 'N/A'}  # After retries, return failure

# Main function to process the addresses
async def main():
    addresses = await read_addresses_from_file('addresses.txt')
    
    # Use aiohttp ClientSession to manage requests
    async with aiohttp.ClientSession() as session:
        tasks = [check_account_owner(session, addr) for addr in addresses]
        results = await asyncio.gather(*tasks)
        
    # Save results to DataFrame
    df = pd.DataFrame(results)
    df.to_excel('addresses_with_owners.xlsx', index=False)
    print("Excel file has been created successfully!")

# Run the asyncio event loop
if __name__ == "__main__":
    asyncio.run(main())
