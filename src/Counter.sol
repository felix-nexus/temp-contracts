// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

contract Counter {
    uint256 public number;
    uint256 public valueRandao;

    function setNumber(uint256 newNumber) public {
        number = newNumber;
    }

    function increment() public {
        number++;
    }

    function setRandao() external {
        valueRandao = block.prevrandao;
    }
}
