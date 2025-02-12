// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

contract CancunOp {
    event SetValue(uint256 value);

    function setValue(uint256 _value) public {
        assembly {
            tstore(0xff, _value)
        }
        emit SetValue(value());
    }

    function value() public view returns (uint256 _value) {
        assembly {
            _value := tload(0xff)
        }
    }

    function execute(CancunOpTCaller _caller, uint256 _value) external returns (uint256) {
        setValue(_value);
        return _caller.execute();
    }
}

contract CancunOpTCaller {
    event Value(uint256);

    CancunOp public op;
    uint256 public value;

    constructor(address _op) {
        op = CancunOp(_op);
    }

    function execute() external returns (uint256 _value) {
        _value = op.value();
        value = _value;
    }
}
