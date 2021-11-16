//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";

contract Stock {
  string public name = "Company Stock Option";
  string public symbol = "CSO";
  uint public totalSupply = 1000000;
  address public minter;
  mapping(address => uint) balances;

  event Sent(address from, address to, uint amount);

  constructor() {
    minter = msg.sender;
    balances[minter] = totalSupply;
  }

  error InsufficientBalance(uint requested, uint available);

  function send(address receiver, uint amount) public {
       if (amount > balances[msg.sender])
           revert InsufficientBalance({
               requested: amount,
               available: balances[msg.sender]
           });

       balances[msg.sender] -= amount;
       balances[receiver] += amount;
       emit Sent(msg.sender, receiver, amount);
   }

  function stockQteOf(address account) external view returns (uint) {
    return balances[account];
  }

  function stockQte() external view returns (uint) {
    return balances[msg.sender];
  }

  function addresse() external view returns (address) {
    return msg.sender;
  }
}
