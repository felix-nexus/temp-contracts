// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std/Test.sol";
import {PrevRandao} from "../src/PrevRandao.sol";

contract PrevRandaoTest is Test {
    PrevRandao public c;

    // forge test -vv --fork-url $CROSS_DEV
    function setUp() external {
        c = new PrevRandao();
        uint256 value = c.value();
        console.log("deploy value", value);
    }

    function test_sum() external {
        for (uint256 i = 0; i < 10; i++) {
            vm.roll(block.number + 1);
            console.log("sum value", c.sum());
        }
    }

    function test_hash() external {
        for (uint256 i = 0; i < 10; i++) {
            vm.roll(block.number + 1);
            console.log("hash value", c.hash());
        }
    }
}
