// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

contract PrevRandao {
    uint256 public blocknumber;
    uint256 public value;

    constructor() {}

    function sum() external returns (uint256) {
        if (blocknumber != block.number) {
            unchecked {
                blocknumber = block.number;
                value += block.prevrandao;
            }
        }
        return value;
    }

    function hash() external returns (uint256) {
        if (blocknumber != block.number) {
            unchecked {
                blocknumber = block.number;
                value = uint256(keccak256(abi.encodePacked(value, block.prevrandao)));
            }
        }
        return value;
    }

    function reset() external returns (uint256) {
        unchecked {
            blocknumber = block.number;
            value = block.prevrandao;
        }
        return value;
    }
}
