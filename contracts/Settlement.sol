// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Settlement {

    event Deposit(address indexed depositor, uint256 amount);
    event BatchSubmitted(bytes32 indexed batchHash);
    event AppRegistered(address indexed app);
    
    mapping(address => uint256) public balances;
    mapping(address => bool) public registeredApps;
    
    function deposit() public payable {
        require(msg.value > 0, "Deposit amount must be greater than 0");
        
        // Update the balance of the sender
        balances[msg.sender] += msg.value;
        
        emit Deposit(msg.sender, msg.value);
    }
    
    // Function to allow users to withdraw their balances
    function withdraw(uint256 amount) public {
        require(balances[msg.sender] >= amount, "Insufficient balance");
        balances[msg.sender] -= amount;
        payable(msg.sender).transfer(amount);
    }

    bytes32[] public batches;
    address public aggregator; // This could easily be a list of aggregators. Keeping it simple for now.
    
    modifier onlyAggregator() {
        require(msg.sender == aggregator, "Only aggregator can call this function");
        _;
    }
    
    constructor() {
        aggregator = msg.sender; // Assigning the contract deployer as the aggregator
    }
    
    struct BatchHeader {
        bytes32 PrevHash;
        bytes32 StateRoot;
        bytes32 TxRoot;
    }

    struct Tx {
       bytes Signature;
    }

    function submitBatch(BatchHeader memory header, Tx[] memory txList) public onlyAggregator {
        // Pseudo code for fraud proof
        require(checkFraudProof(header.TxRoot, txList), "Invalid batch submitted");        
        bytes32 batch_hash = keccak256(abi.encode(header.PrevHash, header.StateRoot));
        batches.push(batch_hash);
        // emit event!
        emit BatchSubmitted(batch_hash);
    }
    
    // Pseudo Fraud Proof Function
    function checkFraudProof(bytes32 txRoot, Tx[] memory txList) private pure returns (bool) {
        // Pseudo code for fraud proof
        return true;
    }

    function registerApp() public {
        require(!registeredApps[msg.sender], "App already registered");
        // Add more constraints here
        registeredApps[msg.sender] = true;
        emit AppRegistered(msg.sender);
    }

    function isAppRegistered(address app) public view returns(bool) {
        return registeredApps[app];
    }
    }
