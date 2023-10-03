# Stackr-mvp

A quick and dirty MVP for Stackr.

# Setup

## Running the L1 contract

Install `anvil` from the foundry toolkit and:
```
anvil
```
Then deploy your contract with:
```
cd stackr-mvp
forge create --rpc-url http://127.0.0.1:8545 --private-key <private-key> contracts/Settlement.sol:Settlement
```

Get the contract address and Deposit funds with:
```
cast send --value 1ether --rpc-url http://127.0.0.1:8545 --private-key <key> <contract-addr> "deposit()"
```

## Running the Application

To run the Rollapp, use the following commands:
    
    ```bash
    cd stackr-mvp
    go get .
    go run .
    ```

To send a signed msg to the rollapp, use the following command:

    ```bash
    cd stackr-mvp/scripts
    go run sendTx.go
    ```

