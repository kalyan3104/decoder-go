use solana_program::pubkey::Pubkey;
use std::time::{SystemTime, UNIX_EPOCH};
use sha2::{Digest, Sha256};

#[derive(Debug)]
struct Block {
    index: u64,
    timestamp: u64,
    data: String,
    previous_hash: String,
    hash: String,
}

impl Block {
    fn new(index: u64, data: String, previous_hash: String) -> Self {
        let timestamp = SystemTime::now()
            .duration_since(UNIX_EPOCH)
            .unwrap()
            .as_secs();

        let hash = Self::calculate_hash(index, timestamp, &data, &previous_hash);

        Block {
            index,
            timestamp,
            data,
            previous_hash,
            hash,
        }
    }

    fn calculate_hash(index: u64, timestamp: u64, data: &str, previous_hash: &str) -> String {
        let input = format!("{}{}{}{}", index, timestamp, data, previous_hash);
        let mut hasher = Sha256::new();
        hasher.update(input.as_bytes());
        format!("{:x}", hasher.finalize())
    }
}

#[derive(Debug)]
struct Blockchain {
    blocks: Vec<Block>,
}

impl Blockchain {
    fn new() -> Self {
        let genesis_block = Block::new(0, "Genesis Block".to_string(), "0".to_string());
        Blockchain {
            blocks: vec![genesis_block],
        }
    }

    fn add_block(&mut self, data: String) {
        let previous_block = self.blocks.last().unwrap();
        let new_block = Block::new(
            previous_block.index + 1,
            data,
            previous_block.hash.clone(),
        );
        self.blocks.push(new_block);
    }
}

fn main() {
    let mut blockchain = Blockchain::new();
    blockchain.add_block("Block 1: User A sent 5 SOL to User B".to_string());
    blockchain.add_block("Block 2: User C sent 3 SOL to User D".to_string());
    blockchain.add_block("Block 3: User E sent 2 SOL to User F".to_string());

    for block in &blockchain.blocks {
        println!("{:?}", block);
    }
}
