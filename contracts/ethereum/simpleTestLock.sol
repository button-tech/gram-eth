pragma solidity >=0.5.2;

contract Lock {

    mapping(address=>mapping(uint256=>address)) lockMap;
    mapping(address=>mapping(uint256=>address)) reverseLockMap;
    mapping(address=>mapping(uint256=>bytes32)) lockSecret;
    mapping(address=>mapping(uint256=>uint256)) lockedSum;
    mapping(address=>uint256) ownersIds;

    function lock(address forWho, bytes32 secret) public payable {
        lockMap[msg.sender][ownersIds[msg.sender]] = forWho;
        reverseLockMap[forWho][ownersIds[msg.sender]] = msg.sender;
        lockSecret[msg.sender][ownersIds[msg.sender]] = secret;
        lockedSum[msg.sender][ownersIds[msg.sender]] = msg.value;
        ownersIds[msg.sender] = ownersIds[msg.sender]+1;
    }

    function unLock(string memory secret, uint256 id) public {
        address ownerOfFunds = reverseLockMap[msg.sender][id];
        require(lockMap[ownerOfFunds][id] == msg.sender);
        require(lockSecret[ownerOfFunds][id] == getSha256(secret));
        uint256 sumToSend = lockedSum[msg.sender][id];
        lockedSum[msg.sender][id] = 0;
        msg.sender.transfer(sumToSend);
    }


    function getSha256(string memory str) public pure returns (bytes32) {
         bytes32 hash = sha256(abi.encodePacked(str));
         return hash;
    }

    function debug(string memory secret, uint256 id) public view returns(uint256) {
        address ownerOfFunds = reverseLockMap[msg.sender][id];
        require(lockMap[ownerOfFunds][id] == msg.sender);
        require(lockSecret[ownerOfFunds][id] == getSha256(secret));
        return lockedSum[msg.sender][id];
    }

    function reveal(string memory secret) public payable {
        msg.sender.transfer(msg.value);
        emit revealed(msg.sender, secret);
    }

    event revealed(address indexed who, string secret);

}