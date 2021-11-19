//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";

contract Stock {
  string public functionCalled;
  string public name = "Company Stock Option";
  string public symbol = "CSO";
  uint public totalSupply = 1000000;
  uint public totalEther = 0;
  address public minter;
  mapping(address => uint) balances;
  address payable[] addressOwners;

  event Sent(address from, address to, uint amount);

  constructor() {
    minter = msg.sender;
    balances[minter] = totalSupply;
    addressOwners.push(payable(minter));
  }

  error InsufficientBalance(uint requested, uint available);

  function sendStockTo(address receiver, uint amount) public {
       if (amount > balances[msg.sender])
           revert InsufficientBalance({
               requested: amount,
               available: balances[msg.sender]
           });

      if (balances[receiver]  == 0) {
        addressOwners.push(payable(receiver));
      }

       balances[msg.sender] -= amount;
       balances[receiver] += amount;
       emit Sent(msg.sender, receiver, amount);
   }

   function getOwners() external view returns (address[] memory) {
    address[] memory ret = new address[](addressOwners.length);
    for (uint i = 0; i < addressOwners.length; i++) {
        ret[i] = addressOwners[i];
    }
     return ret;
   }

   function getTotalSupply() external view returns (uint) {
     return totalSupply;
   }

  function stockQteOf(address account) external view returns (uint) {
    return balances[account];
  }

  function stockQte() external view returns (uint) {
    return balances[msg.sender];
  }

  function nbrOwners() external view returns (uint) {
    return addressOwners.length;
  }

  function addresse() external view returns (address) {
    return msg.sender;
  }

  function sendEtherTo(address payable receiver, uint amountEther) public {
    receiver.transfer(amountEther);
  }

  function sendEther() external payable {
    functionCalled = 'sendEther';
  }

  function distribute() external {
    for (uint i=0; i < addressOwners.length; i++) {
      sendEtherTo(addressOwners[i], address(this).balance * balances[addressOwners[i]]/totalSupply);
    }
  }
}
