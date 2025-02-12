// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std/Test.sol";
import {CancunOp, CancunOpTCaller} from "../src/CancunOp.sol";

contract CancunOpTest is Test {
    CancunOp public op;
    CancunOpTCaller public caller;

    function setUp() public {
        op = new CancunOp();
        caller = new CancunOpTCaller(address(op));
    }
}
