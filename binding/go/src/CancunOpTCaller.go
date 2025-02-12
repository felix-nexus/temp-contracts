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

// CancunOpTCallerMetaData contains all meta data concerning the CancunOpTCaller contract.
var CancunOpTCallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_op\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"op\",\"outputs\":[{\"internalType\":\"contractCancunOp\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Value\",\"type\":\"event\"}]",
	Sigs: map[string]string{
		"61461954": "execute()",
		"ec1f1f69": "op()",
		"3fa4f245": "value()",
	},
	Bin: "0x6080604052348015600e575f5ffd5b506040516101e23803806101e2833981016040819052602b91604e565b5f80546001600160a01b0319166001600160a01b03929092169190911790556079565b5f60208284031215605d575f5ffd5b81516001600160a01b03811681146072575f5ffd5b9392505050565b61015c806100865f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c80633fa4f24514610043578063614619541461005f578063ec1f1f6914610067575b5f5ffd5b61004c60015481565b6040519081526020015b60405180910390f35b61004c610091565b5f54610079906001600160a01b031681565b6040516001600160a01b039091168152602001610056565b5f5f5f9054906101000a90046001600160a01b03166001600160a01b0316633fa4f2456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100e1573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610105919061010f565b6001819055919050565b5f6020828403121561011f575f5ffd5b505191905056fea26469706673582212201c8d95ebc08ca63ee9dbd876ba92d25000e7158229c59247145b4d988ada6cc864736f6c634300081c0033",
}

// CancunOpTCallerABI is the input ABI used to generate the binding from.
// Deprecated: Use CancunOpTCallerMetaData.ABI instead.
var CancunOpTCallerABI = CancunOpTCallerMetaData.ABI

// Deprecated: Use CancunOpTCallerMetaData.Sigs instead.
// CancunOpTCallerFuncSigs maps the 4-byte function signature to its string representation.
var CancunOpTCallerFuncSigs = CancunOpTCallerMetaData.Sigs

// CancunOpTCallerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CancunOpTCallerMetaData.Bin instead.
var CancunOpTCallerBin = CancunOpTCallerMetaData.Bin

// DeployCancunOpTCaller deploys a new Ethereum contract, binding an instance of CancunOpTCaller to it.
func DeployCancunOpTCaller(auth *bind.TransactOpts, backend bind.ContractBackend, _op common.Address) (common.Address, *types.Transaction, *CancunOpTCaller, error) {
	parsed, err := CancunOpTCallerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CancunOpTCallerBin), backend, _op)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CancunOpTCaller{CancunOpTCallerCaller: CancunOpTCallerCaller{contract: contract}, CancunOpTCallerTransactor: CancunOpTCallerTransactor{contract: contract}, CancunOpTCallerFilterer: CancunOpTCallerFilterer{contract: contract}}, nil
}

// CancunOpTCaller is an auto generated Go binding around an Ethereum contract.
type CancunOpTCaller struct {
	CancunOpTCallerCaller     // Read-only binding to the contract
	CancunOpTCallerTransactor // Write-only binding to the contract
	CancunOpTCallerFilterer   // Log filterer for contract events
}

// CancunOpTCallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CancunOpTCallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CancunOpTCallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CancunOpTCallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CancunOpTCallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CancunOpTCallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CancunOpTCallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CancunOpTCallerSession struct {
	Contract     *CancunOpTCaller  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CancunOpTCallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CancunOpTCallerCallerSession struct {
	Contract *CancunOpTCallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CancunOpTCallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CancunOpTCallerTransactorSession struct {
	Contract     *CancunOpTCallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CancunOpTCallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CancunOpTCallerRaw struct {
	Contract *CancunOpTCaller // Generic contract binding to access the raw methods on
}

// CancunOpTCallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CancunOpTCallerCallerRaw struct {
	Contract *CancunOpTCallerCaller // Generic read-only contract binding to access the raw methods on
}

// CancunOpTCallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CancunOpTCallerTransactorRaw struct {
	Contract *CancunOpTCallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCancunOpTCaller creates a new instance of CancunOpTCaller, bound to a specific deployed contract.
func NewCancunOpTCaller(address common.Address, backend bind.ContractBackend) (*CancunOpTCaller, error) {
	contract, err := bindCancunOpTCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CancunOpTCaller{CancunOpTCallerCaller: CancunOpTCallerCaller{contract: contract}, CancunOpTCallerTransactor: CancunOpTCallerTransactor{contract: contract}, CancunOpTCallerFilterer: CancunOpTCallerFilterer{contract: contract}}, nil
}

// NewCancunOpTCallerCaller creates a new read-only instance of CancunOpTCaller, bound to a specific deployed contract.
func NewCancunOpTCallerCaller(address common.Address, caller bind.ContractCaller) (*CancunOpTCallerCaller, error) {
	contract, err := bindCancunOpTCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CancunOpTCallerCaller{contract: contract}, nil
}

// NewCancunOpTCallerTransactor creates a new write-only instance of CancunOpTCaller, bound to a specific deployed contract.
func NewCancunOpTCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*CancunOpTCallerTransactor, error) {
	contract, err := bindCancunOpTCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CancunOpTCallerTransactor{contract: contract}, nil
}

// NewCancunOpTCallerFilterer creates a new log filterer instance of CancunOpTCaller, bound to a specific deployed contract.
func NewCancunOpTCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*CancunOpTCallerFilterer, error) {
	contract, err := bindCancunOpTCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CancunOpTCallerFilterer{contract: contract}, nil
}

// bindCancunOpTCaller binds a generic wrapper to an already deployed contract.
func bindCancunOpTCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CancunOpTCallerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CancunOpTCaller *CancunOpTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CancunOpTCaller.Contract.CancunOpTCallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CancunOpTCaller *CancunOpTCallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CancunOpTCaller.Contract.CancunOpTCallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CancunOpTCaller *CancunOpTCallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CancunOpTCaller.Contract.CancunOpTCallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CancunOpTCaller *CancunOpTCallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CancunOpTCaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CancunOpTCaller *CancunOpTCallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CancunOpTCaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CancunOpTCaller *CancunOpTCallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CancunOpTCaller.Contract.contract.Transact(opts, method, params...)
}

// Op is a free data retrieval call binding the contract method 0xec1f1f69.
//
// Solidity: function op() view returns(address)
func (_CancunOpTCaller *CancunOpTCallerCaller) Op(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CancunOpTCaller.contract.Call(opts, &out, "op")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Op is a free data retrieval call binding the contract method 0xec1f1f69.
//
// Solidity: function op() view returns(address)
func (_CancunOpTCaller *CancunOpTCallerSession) Op() (common.Address, error) {
	return _CancunOpTCaller.Contract.Op(&_CancunOpTCaller.CallOpts)
}

// Op is a free data retrieval call binding the contract method 0xec1f1f69.
//
// Solidity: function op() view returns(address)
func (_CancunOpTCaller *CancunOpTCallerCallerSession) Op() (common.Address, error) {
	return _CancunOpTCaller.Contract.Op(&_CancunOpTCaller.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_CancunOpTCaller *CancunOpTCallerCaller) Value(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CancunOpTCaller.contract.Call(opts, &out, "value")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_CancunOpTCaller *CancunOpTCallerSession) Value() (*big.Int, error) {
	return _CancunOpTCaller.Contract.Value(&_CancunOpTCaller.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_CancunOpTCaller *CancunOpTCallerCallerSession) Value() (*big.Int, error) {
	return _CancunOpTCaller.Contract.Value(&_CancunOpTCaller.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns(uint256 _value)
func (_CancunOpTCaller *CancunOpTCallerTransactor) Execute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CancunOpTCaller.contract.Transact(opts, "execute")
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns(uint256 _value)
func (_CancunOpTCaller *CancunOpTCallerSession) Execute() (*types.Transaction, error) {
	return _CancunOpTCaller.Contract.Execute(&_CancunOpTCaller.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns(uint256 _value)
func (_CancunOpTCaller *CancunOpTCallerTransactorSession) Execute() (*types.Transaction, error) {
	return _CancunOpTCaller.Contract.Execute(&_CancunOpTCaller.TransactOpts)
}

// CancunOpTCallerValueIterator is returned from FilterValue and is used to iterate over the raw logs and unpacked data for Value events raised by the CancunOpTCaller contract.
type CancunOpTCallerValueIterator struct {
	Event *CancunOpTCallerValue // Event containing the contract specifics and raw log

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
func (it *CancunOpTCallerValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CancunOpTCallerValue)
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
		it.Event = new(CancunOpTCallerValue)
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
func (it *CancunOpTCallerValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CancunOpTCallerValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CancunOpTCallerValue represents a Value event raised by the CancunOpTCaller contract.
type CancunOpTCallerValue struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValue is a free log retrieval operation binding the contract event 0x2a27502c345a4cd966daa061d5537f54cd60d2d20b73680b3bf195c91e806a4b.
//
// Solidity: event Value(uint256 arg0)
func (_CancunOpTCaller *CancunOpTCallerFilterer) FilterValue(opts *bind.FilterOpts) (*CancunOpTCallerValueIterator, error) {

	logs, sub, err := _CancunOpTCaller.contract.FilterLogs(opts, "Value")
	if err != nil {
		return nil, err
	}
	return &CancunOpTCallerValueIterator{contract: _CancunOpTCaller.contract, event: "Value", logs: logs, sub: sub}, nil
}

// WatchValue is a free log subscription operation binding the contract event 0x2a27502c345a4cd966daa061d5537f54cd60d2d20b73680b3bf195c91e806a4b.
//
// Solidity: event Value(uint256 arg0)
func (_CancunOpTCaller *CancunOpTCallerFilterer) WatchValue(opts *bind.WatchOpts, sink chan<- *CancunOpTCallerValue) (event.Subscription, error) {

	logs, sub, err := _CancunOpTCaller.contract.WatchLogs(opts, "Value")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CancunOpTCallerValue)
				if err := _CancunOpTCaller.contract.UnpackLog(event, "Value", log); err != nil {
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

// ParseValue is a log parse operation binding the contract event 0x2a27502c345a4cd966daa061d5537f54cd60d2d20b73680b3bf195c91e806a4b.
//
// Solidity: event Value(uint256 arg0)
func (_CancunOpTCaller *CancunOpTCallerFilterer) ParseValue(log types.Log) (*CancunOpTCallerValue, error) {
	event := new(CancunOpTCallerValue)
	if err := _CancunOpTCaller.contract.UnpackLog(event, "Value", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
