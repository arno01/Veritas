# Veritas

A basic, centralized blockchain written in Golang. Commands are given through arguments while running the binary.

COMMANDS:

createwallet --> creates and returns a new wallet and address.

createblockchain -address [WALLET_ADDRESS] --> creates a new blockchain, stored in blockchain.db. The address argument contains the wallet address to which the reward for mining the first block is sent.

send -to [RECIEVER_WALLET_ADDRESS] -from [SENDER_WALLET_ADDRESS] -amount [AMOUNT_TO_SEND] --> sends the specified amount from sender the wallet address to the reciever wallet address. You must have the sender wallet in the wallet.dat file to send. The transaction will be added to a new block and will immediately be mined. 

getbalance -address [WALLET_ADDRESS] --> returns the amount stored in that address.

listaddresses --> lists all the wallet addresses you have stored in your wallet.dat file.

printblockchain --> prints the blockchain, block by block with included transactions.
