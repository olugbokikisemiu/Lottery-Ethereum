pragma solidity >=0.5.2 <0.6.0;

contract lottery {
    address public manager;
    mapping(address => bool) playerMap;
    address[] public players;
    
    string private managerRestriction = "Managers are not allowed in competition";
    string private playerRestriction = "Only manager can select winner";
    string private joinRestriction = "Player already joined the lottery";
    
    constructor() public {
        manager = msg.sender;
    }
    
    function allPlayer() public view returns (address[] memory){
        return players;
    }
    
    function enter() public restrictManager payable {
        require(msg.value > 0.1 ether, "Insufficient amount send ");
        isPlayerExist();
        players.push(msg.sender);
    }
    
    function randomGenerator() private view returns (uint) {
        return uint(keccak256(abi.encodePacked(block.difficulty, now, players)));
    }
    
    function isPlayerExist() private hasJoined{
        playerMap[msg.sender] = true;
    }
    
    function selectWinner() public restrictPlayer returns (address){
        uint index = randomGenerator() % players.length;
        address payable selectedPlayer = address(uint160(players[index]));
        selectedPlayer.transfer(address(this).balance);
        players.length = 0;
        return selectedPlayer;
    }
    
    modifier restrictManager() {
       require(msg.sender != manager, managerRestriction);
        _;
    }
    
    modifier restrictPlayer() {
       require(msg.sender == manager, playerRestriction);
        _;
    }
    
    modifier hasJoined(){
        require(playerMap[msg.sender] == false, joinRestriction);
        _;
    }
}