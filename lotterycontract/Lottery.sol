pragma solidity >=0.5.2 <0.6.0;

contract lottery {
    address public manager;
    address[] public players;
    constructor() public {
        manager = msg.sender;
    }
    function enter() public payable {
        require(msg.value > 0.1 ether, "Insufficient amount send ");
        players.push(msg.sender);
    }
    function randomGenerator() private view returns (uint) {
        return uint(keccak256(abi.encodePacked(block.difficulty, block.number, players)));
    }
}