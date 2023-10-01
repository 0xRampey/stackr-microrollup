// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Settlement {

    event Deposit(address indexed depositor, uint256 amount);
    event BatchSubmitted(bytes32 indexed batchHash);
    
    mapping(address => uint256) public balances;
    
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
        require(checkFraudProof(header.TxRoot, txList), "Invalid batch");
        // Calculate batch hash and append to batches
        bytes32 batch_hash = keccak256(abi.encode(header.PrevHash, header.StateRoot));
        batches.push(batch_hash);
        // emit event!
        emit BatchSubmitted(batch_hash);
    }
    
    // Pseudo Fraud Proof Function
    function checkFraudProof(bytes32 txRoot, Tx[] memory txList) private pure returns (bool) {
        // TODO: Compute merkle tree of txs and compare with txRoot
        // bytes32[] memory leaves = new bytes32[](txList.length);
        // for (uint256 i = 0; i < txList.length; i++) {
        //     leaves[i] = keccak256(abi.encode(txList[i].Signature));
        // }
        // bytes32 root = computeMerkleRoot(leaves);
        return true;
    }

    function computeMerkleRoot(bytes32[] memory leaves) private pure returns (bytes32) {
        uint256 n = leaves.length;
        bytes32[] memory nodes = new bytes32[](n * 2);
        for (uint256 i = 0; i < n; i++) {
            nodes[n + i] = leaves[i];
        }
        for (uint256 i = n - 1; i > 0; i--) {
            nodes[i] = keccak256(abi.encodePacked(nodes[i * 2], nodes[i * 2 + 1]));
        }
        return nodes[1];
    }
    }
