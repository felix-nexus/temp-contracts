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

// CancunOpSelfDestructMetaData contains all meta data concerning the CancunOpSelfDestruct contract.
var CancunOpSelfDestructMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"destruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"2b68b9c6": "destruct()",
	},
	Bin: "0x6080604052348015600e575f5ffd5b50606780601a5f395ff3fe6080604052348015600e575f5ffd5b50600436106026575f3560e01c80632b68b9c614602a575b5f5ffd5b602f33ff5b00fea26469706673582212201920308db1853a8c0e6b8e09cd40b031687e5a85dcaa63a5727bb2a5bc18b94564736f6c634300081c0033",
}

// CancunOpSelfDestructABI is the input ABI used to generate the binding from.
// Deprecated: Use CancunOpSelfDestructMetaData.ABI instead.
var CancunOpSelfDestructABI = CancunOpSelfDestructMetaData.ABI

// Deprecated: Use CancunOpSelfDestructMetaData.Sigs instead.
// CancunOpSelfDestructFuncSigs maps the 4-byte function signature to its string representation.
var CancunOpSelfDestructFuncSigs = CancunOpSelfDestructMetaData.Sigs

// CancunOpSelfDestructBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CancunOpSelfDestructMetaData.Bin instead.
var CancunOpSelfDestructBin = CancunOpSelfDestructMetaData.Bin

// DeployCancunOpSelfDestruct deploys a new Ethereum contract, binding an instance of CancunOpSelfDestruct to it.
func DeployCancunOpSelfDestruct(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CancunOpSelfDestruct, error) {
	parsed, err := CancunOpSelfDestructMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CancunOpSelfDestructBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CancunOpSelfDestruct{CancunOpSelfDestructCaller: CancunOpSelfDestructCaller{contract: contract}, CancunOpSelfDestructTransactor: CancunOpSelfDestructTransactor{contract: contract}, CancunOpSelfDestructFilterer: CancunOpSelfDestructFilterer{contract: contract}}, nil
}

// CancunOpSelfDestruct is an auto generated Go binding around an Ethereum contract.
type CancunOpSelfDestruct struct {
	CancunOpSelfDestructCaller     // Read-only binding to the contract
	CancunOpSelfDestructTransactor // Write-only binding to the contract
	CancunOpSelfDestructFilterer   // Log filterer for contract events
}

// CancunOpSelfDestructCaller is an auto generated read-only Go binding around an Ethereum contract.
type CancunOpSelfDestructCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CancunOpSelfDestructTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CancunOpSelfDestructTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CancunOpSelfDestructFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CancunOpSelfDestructFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CancunOpSelfDestructSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CancunOpSelfDestructSession struct {
	Contract     *CancunOpSelfDestruct // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CancunOpSelfDestructCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CancunOpSelfDestructCallerSession struct {
	Contract *CancunOpSelfDestructCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CancunOpSelfDestructTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CancunOpSelfDestructTransactorSession struct {
	Contract     *CancunOpSelfDestructTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CancunOpSelfDestructRaw is an auto generated low-level Go binding around an Ethereum contract.
type CancunOpSelfDestructRaw struct {
	Contract *CancunOpSelfDestruct // Generic contract binding to access the raw methods on
}

// CancunOpSelfDestructCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CancunOpSelfDestructCallerRaw struct {
	Contract *CancunOpSelfDestructCaller // Generic read-only contract binding to access the raw methods on
}

// CancunOpSelfDestructTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CancunOpSelfDestructTransactorRaw struct {
	Contract *CancunOpSelfDestructTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCancunOpSelfDestruct creates a new instance of CancunOpSelfDestruct, bound to a specific deployed contract.
func NewCancunOpSelfDestruct(address common.Address, backend bind.ContractBackend) (*CancunOpSelfDestruct, error) {
	contract, err := bindCancunOpSelfDestruct(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CancunOpSelfDestruct{CancunOpSelfDestructCaller: CancunOpSelfDestructCaller{contract: contract}, CancunOpSelfDestructTransactor: CancunOpSelfDestructTransactor{contract: contract}, CancunOpSelfDestructFilterer: CancunOpSelfDestructFilterer{contract: contract}}, nil
}

// NewCancunOpSelfDestructCaller creates a new read-only instance of CancunOpSelfDestruct, bound to a specific deployed contract.
func NewCancunOpSelfDestructCaller(address common.Address, caller bind.ContractCaller) (*CancunOpSelfDestructCaller, error) {
	contract, err := bindCancunOpSelfDestruct(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CancunOpSelfDestructCaller{contract: contract}, nil
}

// NewCancunOpSelfDestructTransactor creates a new write-only instance of CancunOpSelfDestruct, bound to a specific deployed contract.
func NewCancunOpSelfDestructTransactor(address common.Address, transactor bind.ContractTransactor) (*CancunOpSelfDestructTransactor, error) {
	contract, err := bindCancunOpSelfDestruct(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CancunOpSelfDestructTransactor{contract: contract}, nil
}

// NewCancunOpSelfDestructFilterer creates a new log filterer instance of CancunOpSelfDestruct, bound to a specific deployed contract.
func NewCancunOpSelfDestructFilterer(address common.Address, filterer bind.ContractFilterer) (*CancunOpSelfDestructFilterer, error) {
	contract, err := bindCancunOpSelfDestruct(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CancunOpSelfDestructFilterer{contract: contract}, nil
}

// bindCancunOpSelfDestruct binds a generic wrapper to an already deployed contract.
func bindCancunOpSelfDestruct(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CancunOpSelfDestructMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CancunOpSelfDestruct *CancunOpSelfDestructRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CancunOpSelfDestruct.Contract.CancunOpSelfDestructCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CancunOpSelfDestruct *CancunOpSelfDestructRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CancunOpSelfDestruct.Contract.CancunOpSelfDestructTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CancunOpSelfDestruct *CancunOpSelfDestructRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CancunOpSelfDestruct.Contract.CancunOpSelfDestructTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CancunOpSelfDestruct *CancunOpSelfDestructCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CancunOpSelfDestruct.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CancunOpSelfDestruct *CancunOpSelfDestructTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CancunOpSelfDestruct.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CancunOpSelfDestruct *CancunOpSelfDestructTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CancunOpSelfDestruct.Contract.contract.Transact(opts, method, params...)
}

// Destruct is a paid mutator transaction binding the contract method 0x2b68b9c6.
//
// Solidity: function destruct() returns()
func (_CancunOpSelfDestruct *CancunOpSelfDestructTransactor) Destruct(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CancunOpSelfDestruct.contract.Transact(opts, "destruct")
}

// Destruct is a paid mutator transaction binding the contract method 0x2b68b9c6.
//
// Solidity: function destruct() returns()
func (_CancunOpSelfDestruct *CancunOpSelfDestructSession) Destruct() (*types.Transaction, error) {
	return _CancunOpSelfDestruct.Contract.Destruct(&_CancunOpSelfDestruct.TransactOpts)
}

// Destruct is a paid mutator transaction binding the contract method 0x2b68b9c6.
//
// Solidity: function destruct() returns()
func (_CancunOpSelfDestruct *CancunOpSelfDestructTransactorSession) Destruct() (*types.Transaction, error) {
	return _CancunOpSelfDestruct.Contract.Destruct(&_CancunOpSelfDestruct.TransactOpts)
}
