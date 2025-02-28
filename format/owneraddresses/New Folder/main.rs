use std::collections::HashMap;
use std::sync::{Arc, Mutex};
use std::thread;
use std::time::{Duration, SystemTime, UNIX_EPOCH};
use sha2::{Digest, Sha256};

#[derive(Debug, Clone)]
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

    fn get_last_block(&self) -> &Block {
        self.blocks.last().unwrap()
    }
}

#[derive(Debug)]
struct Node {
    id: u64,
    blockchain: Arc<Mutex<Blockchain>>,
    peers: Arc<Mutex<HashMap<u64, Arc<Mutex<Blockchain>>>>>,
}

impl Node {
    fn new(id: u64) -> Self {
        let blockchain = Arc::new(Mutex::new(Blockchain::new()));
        let peers = Arc::new(Mutex::new(HashMap::new()));
        Node {
            id,
            blockchain,
            peers,
        }
    }

    fn add_peer(&self, peer_id: u64, peer_blockchain: Arc<Mutex<Blockchain>>) {
        self.peers.lock().unwrap().insert(peer_id, peer_blockchain);
    }

    fn mine_block(&self, data: String) {
        let mut blockchain = self.blockchain.lock().unwrap();
        blockchain.add_block(data.clone());
        println!("Node {} mined a new block: {}", self.id, data);

        // Share the new block with peers
        self.share_block();
    }

    fn share_block(&self) {
        let latest_block = self.blockchain.lock().unwrap().get_last_block().clone();
        for (peer_id, peer_chain) in self.peers.lock().unwrap().iter() {
            let mut peer_chain = peer_chain.lock().unwrap();
            peer_chain.add_block(latest_block.data.clone());
            println!(
                "Node {} shared block '{}' with Node {}",
                self.id, latest_block.data, peer_id
            );
        }
    }
}

fn main() {
    // Create nodes
    let node1 = Node::new(1);
    let node2 = Node::new(2);
    let node3 = Node::new(3);

    // Establish peer connections
    node1.add_peer(2, Arc::clone(&node2.blockchain));
    node1.add_peer(3, Arc::clone(&node3.blockchain));

    // Simulate mining and sharing data
    node1.mine_block("Transaction 1: Alice sends 5 SOL to Bob".to_string());
    thread::sleep(Duration::from_secs(2));

    node1.mine_block("Transaction 2: Bob sends 2 SOL to Charlie".to_string());
    thread::sleep(Duration::from_secs(2));

    // Print the blockchains for each node
    println!("\nNode 1 Blockchain: {:?}", node1.blockchain.lock().unwrap());
    println!("Node 2 Blockchain: {:?}", node2.blockchain.lock().unwrap());
    println!("Node 3 Blockchain: {:?}", node3.blockchain.lock().unwrap());
}
