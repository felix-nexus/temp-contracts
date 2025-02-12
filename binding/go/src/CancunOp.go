// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CancunOpMetaData contains all meta data concerning the CancunOp contract.
var CancunOpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCancunOpTCaller\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SetValue\",\"type\":\"event\"}]",
	Sigs: map[string]string{
		"3b89bb86": "execute(address,uint256)",
		"55241077": "setValue(uint256)",
		"3fa4f245": "value()",
	},
	Bin: "0x6080604052348015600e575f5ffd5b506101cc8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c80633b89bb86146100435780633fa4f245146100685780635524107714610070575b5f5ffd5b610056610051366004610133565b610085565b60405190815260200160405180910390f35b60ff5c610056565b61008361007e366004610168565b6100f7565b005b5f61008f826100f7565b826001600160a01b031663614619546040518163ffffffff1660e01b81526004016020604051808303815f875af11580156100cc573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100f0919061017f565b9392505050565b8060ff5d7fed8b07065ef0737c0cfb1bf1e23ccc881d797ec9804f74230a360b84982938ab60ff5c60405190815260200160405180910390a150565b5f5f60408385031215610144575f5ffd5b82356001600160a01b038116811461015a575f5ffd5b946020939093013593505050565b5f60208284031215610178575f5ffd5b5035919050565b5f6020828403121561018f575f5ffd5b505191905056fea2646970667358221220688807393224d95a34ba12a216d3b9531b01bebb13bc58a2b4d1225568d9dca564736f6c634300081c0033",
}

// CancunOpABI is the input ABI used to generate the binding from.
// Deprecated: Use CancunOpMetaData.ABI instead.
var CancunOpABI = CancunOpMetaData.ABI

// Deprecated: Use CancunOpMetaData.Sigs instead.
// CancunOpFuncSigs maps the 4-byte function signature to its string representation.
var CancunOpFuncSigs = CancunOpMetaData.Sigs

// CancunOpBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CancunOpMetaData.Bin instead.
var CancunOpBin = CancunOpMetaData.Bin

// DeployCancunOp deploys a new Ethereum contract, binding an instance of CancunOp to it.
func DeployCancunOp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CancunOp, error) {
	parsed, err := CancunOpMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CancunOpBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CancunOp{CancunOpCaller: CancunOpCaller{contract: contract}, CancunOpTransactor: CancunOpTransactor{contract: contract}, CancunOpFilterer: CancunOpFilterer{contract: contract}}, nil
}

// CancunOp is an auto generated Go binding around an Ethereum contract.
type CancunOp struct {
	CancunOpCaller     // Read-only binding to the contract
	CancunOpTransactor // Write-only binding to the contract
	CancunOpFilterer   // Log filterer for contract events
}

// CancunOpCaller is an auto generated read-only Go binding around an Ethereum contract.
type CancunOpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CancunOpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CancunOpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CancunOpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CancunOpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CancunOpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CancunOpSession struct {
	Contract     *CancunOp         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CancunOpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CancunOpCallerSession struct {
	Contract *CancunOpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CancunOpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CancunOpTransactorSession struct {
	Contract     *CancunOpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CancunOpRaw is an auto generated low-level Go binding around an Ethereum contract.
type CancunOpRaw struct {
	Contract *CancunOp // Generic contract binding to access the raw methods on
}

// CancunOpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CancunOpCallerRaw struct {
	Contract *CancunOpCaller // Generic read-only contract binding to access the raw methods on
}

// CancunOpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CancunOpTransactorRaw struct {
	Contract *CancunOpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCancunOp creates a new instance of CancunOp, bound to a specific deployed contract.
func NewCancunOp(address common.Address, backend bind.ContractBackend) (*CancunOp, error) {
	contract, err := bindCancunOp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CancunOp{CancunOpCaller: CancunOpCaller{contract: contract}, CancunOpTransactor: CancunOpTransactor{contract: contract}, CancunOpFilterer: CancunOpFilterer{contract: contract}}, nil
}

// NewCancunOpCaller creates a new read-only instance of CancunOp, bound to a specific deployed contract.
func NewCancunOpCaller(address common.Address, caller bind.ContractCaller) (*CancunOpCaller, error) {
	contract, err := bindCancunOp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CancunOpCaller{contract: contract}, nil
}

// NewCancunOpTransactor creates a new write-only instance of CancunOp, bound to a specific deployed contract.
func NewCancunOpTransactor(address common.Address, transactor bind.ContractTransactor) (*CancunOpTransactor, error) {
	contract, err := bindCancunOp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CancunOpTransactor{contract: contract}, nil
}

// NewCancunOpFilterer creates a new log filterer instance of CancunOp, bound to a specific deployed contract.
func NewCancunOpFilterer(address common.Address, filterer bind.ContractFilterer) (*CancunOpFilterer, error) {
	contract, err := bindCancunOp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CancunOpFilterer{contract: contract}, nil
}

// bindCancunOp binds a generic wrapper to an already deployed contract.
func bindCancunOp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CancunOpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CancunOp *CancunOpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CancunOp.Contract.CancunOpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CancunOp *CancunOpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CancunOp.Contract.CancunOpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CancunOp *CancunOpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CancunOp.Contract.CancunOpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CancunOp *CancunOpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CancunOp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CancunOp *CancunOpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CancunOp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CancunOp *CancunOpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CancunOp.Contract.contract.Transact(opts, method, params...)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256 _value)
func (_CancunOp *CancunOpCaller) Value(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CancunOp.contract.Call(opts, &out, "value")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256 _value)
func (_CancunOp *CancunOpSession) Value() (*big.Int, error) {
	return _CancunOp.Contract.Value(&_CancunOp.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256 _value)
func (_CancunOp *CancunOpCallerSession) Value() (*big.Int, error) {
	return _CancunOp.Contract.Value(&_CancunOp.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x3b89bb86.
//
// Solidity: function execute(address _caller, uint256 _value) returns(uint256)
func (_CancunOp *CancunOpTransactor) Execute(opts *bind.TransactOpts, _caller common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CancunOp.contract.Transact(opts, "execute", _caller, _value)
}

// Execute is a paid mutator transaction binding the contract method 0x3b89bb86.
//
// Solidity: function execute(address _caller, uint256 _value) returns(uint256)
func (_CancunOp *CancunOpSession) Execute(_caller common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CancunOp.Contract.Execute(&_CancunOp.TransactOpts, _caller, _value)
}

// Execute is a paid mutator transaction binding the contract method 0x3b89bb86.
//
// Solidity: function execute(address _caller, uint256 _value) returns(uint256)
func (_CancunOp *CancunOpTransactorSession) Execute(_caller common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CancunOp.Contract.Execute(&_CancunOp.TransactOpts, _caller, _value)
}

// SetValue is a paid mutator transaction binding the contract method 0x55241077.
//
// Solidity: function setValue(uint256 _value) returns()
func (_CancunOp *CancunOpTransactor) SetValue(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _CancunOp.contract.Transact(opts, "setValue", _value)
}

// SetValue is a paid mutator transaction binding the contract method 0x55241077.
//
// Solidity: function setValue(uint256 _value) returns()
func (_CancunOp *CancunOpSession) SetValue(_value *big.Int) (*types.Transaction, error) {
	return _CancunOp.Contract.SetValue(&_CancunOp.TransactOpts, _value)
}

// SetValue is a paid mutator transaction binding the contract method 0x55241077.
//
// Solidity: function setValue(uint256 _value) returns()
func (_CancunOp *CancunOpTransactorSession) SetValue(_value *big.Int) (*types.Transaction, error) {
	return _CancunOp.Contract.SetValue(&_CancunOp.TransactOpts, _value)
}

// CancunOpSetValueIterator is returned from FilterSetValue and is used to iterate over the raw logs and unpacked data for SetValue events raised by the CancunOp contract.
type CancunOpSetValueIterator struct {
	Event *CancunOpSetValue // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CancunOpSetValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CancunOpSetValue)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CancunOpSetValue)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CancunOpSetValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CancunOpSetValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CancunOpSetValue represents a SetValue event raised by the CancunOp contract.
type CancunOpSetValue struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetValue is a free log retrieval operation binding the contract event 0xed8b07065ef0737c0cfb1bf1e23ccc881d797ec9804f74230a360b84982938ab.
//
// Solidity: event SetValue(uint256 value)
func (_CancunOp *CancunOpFilterer) FilterSetValue(opts *bind.FilterOpts) (*CancunOpSetValueIterator, error) {

	logs, sub, err := _CancunOp.contract.FilterLogs(opts, "SetValue")
	if err != nil {
		return nil, err
	}
	return &CancunOpSetValueIterator{contract: _CancunOp.contract, event: "SetValue", logs: logs, sub: sub}, nil
}

// WatchSetValue is a free log subscription operation binding the contract event 0xed8b07065ef0737c0cfb1bf1e23ccc881d797ec9804f74230a360b84982938ab.
//
// Solidity: event SetValue(uint256 value)
func (_CancunOp *CancunOpFilterer) WatchSetValue(opts *bind.WatchOpts, sink chan<- *CancunOpSetValue) (event.Subscription, error) {

	logs, sub, err := _CancunOp.contract.WatchLogs(opts, "SetValue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CancunOpSetValue)
				if err := _CancunOp.contract.UnpackLog(event, "SetValue", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetValue is a log parse operation binding the contract event 0xed8b07065ef0737c0cfb1bf1e23ccc881d797ec9804f74230a360b84982938ab.
//
// Solidity: event SetValue(uint256 value)
func (_CancunOp *CancunOpFilterer) ParseSetValue(log types.Log) (*CancunOpSetValue, error) {
	event := new(CancunOpSetValue)
	if err := _CancunOp.contract.UnpackLog(event, "SetValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
