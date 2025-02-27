# Blockchain Network Simulation

## Introduction

Blockchain is a groundbreaking technology of the 21st century that continues to evolve, with much of its potential yet to be fully realized. At its core, blockchain operates as a distributed ledger where records are maintained across a network. Unlike traditional private databases, blockchain is inherently public—each participant holds a complete or partial copy of the ledger. New records are added only when the network reaches consensus, ensuring both transparency and security. This innovative framework has paved the way for cryptocurrencies and smart contracts.

## Network Architecture

In my implementation, I chose to simplify the network by introducing a level of centralization. In a fully decentralized blockchain, nodes typically discover peers using DNS seeds. However, to keep the simulation straightforward, I implemented a central node that all other nodes connect to. This design decision makes connectivity and data exchange simple while ensuring that every key blockchain element—from Merkle trees to mining, digital signing, and wallet creation—is robust and fully functional.

My setup consists of three nodes:

- **Central Node:** Acts as the hub for all node connections and data transfers.
- **Miner Node:** Collects new transactions in a memory pool and mines new blocks once a sufficient number of transactions have accumulated (for now 2tx, and reward 10 coins).
- **Wallet Node:** Used to send coins between wallets, storing a full copy of the blockchain (unlike SPV nodes).

## The Scenario

This project implements the following sequence of events:

1. **Blockchain Creation:** The central node creates the blockchain.
2. **Initial Synchronization:** A wallet node connects to the central node and downloads the blockchain.
3. **Additional Node Connection:** A miner node connects to the central node and downloads the blockchain.
4. **Transaction Initiation:** The wallet node creates a transaction.
5. **Transaction Reception:** Miner nodes receive the transaction and store it in their memory pool.
6. **Block Mining:** Once enough transactions accumulate, a miner starts mining a new block.
7. **Block Dissemination:** When the new block is mined, it is sent to the central node.
8. **Node Synchronization:** The wallet node synchronizes with the central node.
9. **Transaction Verification:** The user of the wallet node confirms that their payment was successfully processed.

## Result

1. Export Node IDs
   Node 3000 is the cntral node, Node 3001 is the wallet node, Node 3002 is the miner node
   ![Screenshot from 2025-02-27 16-05-33](https://github.com/user-attachments/assets/f49d3b1f-1804-4590-b65b-443068327ca5)

2. Create a wallet for the central node and create the the blockchain by creating the genesis block, at the end the wallet will have 10 coins as a reward.
   Save the genesis block to use it in other nodes, because the genesis block should be hardcoded, to start the network.
   in the image below we can see there are four 0's indicating that the block was mined.
   ![Screenshot from 2025-02-27 16-06-57](https://github.com/user-attachments/assets/50dc8442-e5f6-496a-a157-2d3d2af65199)

3. Create a wallet for Node 3001, and use Node 3000 to send it 10 coins,
   add the `-mine` flag because there are no miners yet and we want the sender to mine the block directly,
   then start the central node server till the end of the simulation.
   ![Screenshot from 2025-02-27 16-08-13](https://github.com/user-attachments/assets/c9f3c738-c3e2-461c-b97c-6b1d82d8d28e)

4. Initialize the blockchain for Node 3001 with the genesis block, then start the node, in the image below we can see that the central node 3000 recieved a request to send the blockchain to Node 3001.
   ![Screenshot from 2025-02-27 16-09-23](https://github.com/user-attachments/assets/f263c447-863a-4062-ac5c-385a7768323f)

5. Stop Node 3001, to check the balances. Now it has it's own copy of the blockchain.
   Central Node Wallet: mined two block so 20 and sent 10, so final result is 10.
   Wallet Node: recieved 10.
   ![Screenshot from 2025-02-27 16-10-29](https://github.com/user-attachments/assets/914d39cb-67da-4b5a-aff9-8be414ae6b93)

6. Initialize the blockchain for Node 3002, then create a wallet for it and, start Node 3002 with `-miner` flag to indicate, that it will taket transactions from the mempool and mine the new blocks.
   It will download the blockchain from the central node as shown below and wait for the mempool to fill at least two transactions.
   ![Screenshot from 2025-02-27 16-10-29](https://github.com/user-attachments/assets/36bd8eb3-0569-43c5-84e3-b92281fa5088)

7. Create two new wallets, in Node 3001, and send 3 coins for each one and notice how the mempool gets filled, and Node 3002 starts the mining process.
   ![Screenshot from 2025-02-27 16-15-37](https://github.com/user-attachments/assets/9e284ab8-107e-496b-aa2a-6f53aa602047)

8. Finally start the Node 3001 to synchronize it and download the last block, then check the blanced of the wallets.
   the two new wallets must have both 3, and the miner's wallet will have the reward 10 coins.
   ![Screenshot from 2025-02-27 16-18-26](https://github.com/user-attachments/assets/c21552cd-9f80-4876-bdcc-97a3fe60051f)

## Notes:
- there are other commands like `printchain`, `reindexutxo`, and `listaddresses` that I didn't cover here in the scenario.
- sending coins from a wallet can be done only if the sender's wallet private key was known by the node and in my implementation, the node that create a wallet save the private keys in wallet_300.dat file (for example for node 3000)

## Acknowledgements

Special thanks to [Ivan Kuznetsov](https://github.com/jeiwan) for his amazing tutorial.   
