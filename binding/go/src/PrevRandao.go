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

// PrevRandaoMetaData contains all meta data concerning the PrevRandao contract.
var PrevRandaoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"blocknumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"05062247": "blocknumber()",
		"09bd5a60": "hash()",
		"d826f88f": "reset()",
		"853255cc": "sum()",
		"3fa4f245": "value()",
	},
	Bin: "0x6080604052348015600e575f5ffd5b506101218061001c5f395ff3fe6080604052348015600e575f5ffd5b5060043610604e575f3560e01c80630506224714605257806309bd5a6014606b5780633fa4f245146071578063853255cc146079578063d826f88f14607f575b5f5ffd5b60595f5481565b60405190815260200160405180910390f35b6059608c565b605960015481565b605960d1565b435f554460018190556059565b5f435f541460ca57435f55600154604080516020810192909252449082015260600160408051601f1981840301815291905280516020909101206001555b5060015490565b5f435f541460ca57435f556001805444019055506001549056fea26469706673582212201d5da2c87e86ee47d3b1cfb537d1f474bf55f18415cb7c0aa7884b704bebcfaf64736f6c634300081c0033",
}

// PrevRandaoABI is the input ABI used to generate the binding from.
// Deprecated: Use PrevRandaoMetaData.ABI instead.
var PrevRandaoABI = PrevRandaoMetaData.ABI

// Deprecated: Use PrevRandaoMetaData.Sigs instead.
// PrevRandaoFuncSigs maps the 4-byte function signature to its string representation.
var PrevRandaoFuncSigs = PrevRandaoMetaData.Sigs

// PrevRandaoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PrevRandaoMetaData.Bin instead.
var PrevRandaoBin = PrevRandaoMetaData.Bin

// DeployPrevRandao deploys a new Ethereum contract, binding an instance of PrevRandao to it.
func DeployPrevRandao(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PrevRandao, error) {
	parsed, err := PrevRandaoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PrevRandaoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrevRandao{PrevRandaoCaller: PrevRandaoCaller{contract: contract}, PrevRandaoTransactor: PrevRandaoTransactor{contract: contract}, PrevRandaoFilterer: PrevRandaoFilterer{contract: contract}}, nil
}

// PrevRandao is an auto generated Go binding around an Ethereum contract.
type PrevRandao struct {
	PrevRandaoCaller     // Read-only binding to the contract
	PrevRandaoTransactor // Write-only binding to the contract
	PrevRandaoFilterer   // Log filterer for contract events
}

// PrevRandaoCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrevRandaoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrevRandaoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrevRandaoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrevRandaoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrevRandaoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrevRandaoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrevRandaoSession struct {
	Contract     *PrevRandao       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrevRandaoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrevRandaoCallerSession struct {
	Contract *PrevRandaoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PrevRandaoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrevRandaoTransactorSession struct {
	Contract     *PrevRandaoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PrevRandaoRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrevRandaoRaw struct {
	Contract *PrevRandao // Generic contract binding to access the raw methods on
}

// PrevRandaoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrevRandaoCallerRaw struct {
	Contract *PrevRandaoCaller // Generic read-only contract binding to access the raw methods on
}

// PrevRandaoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrevRandaoTransactorRaw struct {
	Contract *PrevRandaoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrevRandao creates a new instance of PrevRandao, bound to a specific deployed contract.
func NewPrevRandao(address common.Address, backend bind.ContractBackend) (*PrevRandao, error) {
	contract, err := bindPrevRandao(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrevRandao{PrevRandaoCaller: PrevRandaoCaller{contract: contract}, PrevRandaoTransactor: PrevRandaoTransactor{contract: contract}, PrevRandaoFilterer: PrevRandaoFilterer{contract: contract}}, nil
}

// NewPrevRandaoCaller creates a new read-only instance of PrevRandao, bound to a specific deployed contract.
func NewPrevRandaoCaller(address common.Address, caller bind.ContractCaller) (*PrevRandaoCaller, error) {
	contract, err := bindPrevRandao(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrevRandaoCaller{contract: contract}, nil
}

// NewPrevRandaoTransactor creates a new write-only instance of PrevRandao, bound to a specific deployed contract.
func NewPrevRandaoTransactor(address common.Address, transactor bind.ContractTransactor) (*PrevRandaoTransactor, error) {
	contract, err := bindPrevRandao(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrevRandaoTransactor{contract: contract}, nil
}

// NewPrevRandaoFilterer creates a new log filterer instance of PrevRandao, bound to a specific deployed contract.
func NewPrevRandaoFilterer(address common.Address, filterer bind.ContractFilterer) (*PrevRandaoFilterer, error) {
	contract, err := bindPrevRandao(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrevRandaoFilterer{contract: contract}, nil
}

// bindPrevRandao binds a generic wrapper to an already deployed contract.
func bindPrevRandao(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PrevRandaoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrevRandao *PrevRandaoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrevRandao.Contract.PrevRandaoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrevRandao *PrevRandaoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrevRandao.Contract.PrevRandaoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrevRandao *PrevRandaoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrevRandao.Contract.PrevRandaoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrevRandao *PrevRandaoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrevRandao.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrevRandao *PrevRandaoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrevRandao.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrevRandao *PrevRandaoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrevRandao.Contract.contract.Transact(opts, method, params...)
}

// Blocknumber is a free data retrieval call binding the contract method 0x05062247.
//
// Solidity: function blocknumber() view returns(uint256)
func (_PrevRandao *PrevRandaoCaller) Blocknumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrevRandao.contract.Call(opts, &out, "blocknumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Blocknumber is a free data retrieval call binding the contract method 0x05062247.
//
// Solidity: function blocknumber() view returns(uint256)
func (_PrevRandao *PrevRandaoSession) Blocknumber() (*big.Int, error) {
	return _PrevRandao.Contract.Blocknumber(&_PrevRandao.CallOpts)
}

// Blocknumber is a free data retrieval call binding the contract method 0x05062247.
//
// Solidity: function blocknumber() view returns(uint256)
func (_PrevRandao *PrevRandaoCallerSession) Blocknumber() (*big.Int, error) {
	return _PrevRandao.Contract.Blocknumber(&_PrevRandao.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_PrevRandao *PrevRandaoCaller) Value(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrevRandao.contract.Call(opts, &out, "value")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_PrevRandao *PrevRandaoSession) Value() (*big.Int, error) {
	return _PrevRandao.Contract.Value(&_PrevRandao.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_PrevRandao *PrevRandaoCallerSession) Value() (*big.Int, error) {
	return _PrevRandao.Contract.Value(&_PrevRandao.CallOpts)
}

// Hash is a paid mutator transaction binding the contract method 0x09bd5a60.
//
// Solidity: function hash() returns(uint256)
func (_PrevRandao *PrevRandaoTransactor) Hash(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrevRandao.contract.Transact(opts, "hash")
}

// Hash is a paid mutator transaction binding the contract method 0x09bd5a60.
//
// Solidity: function hash() returns(uint256)
func (_PrevRandao *PrevRandaoSession) Hash() (*types.Transaction, error) {
	return _PrevRandao.Contract.Hash(&_PrevRandao.TransactOpts)
}

// Hash is a paid mutator transaction binding the contract method 0x09bd5a60.
//
// Solidity: function hash() returns(uint256)
func (_PrevRandao *PrevRandaoTransactorSession) Hash() (*types.Transaction, error) {
	return _PrevRandao.Contract.Hash(&_PrevRandao.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns(uint256)
func (_PrevRandao *PrevRandaoTransactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrevRandao.contract.Transact(opts, "reset")
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns(uint256)
func (_PrevRandao *PrevRandaoSession) Reset() (*types.Transaction, error) {
	return _PrevRandao.Contract.Reset(&_PrevRandao.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns(uint256)
func (_PrevRandao *PrevRandaoTransactorSession) Reset() (*types.Transaction, error) {
	return _PrevRandao.Contract.Reset(&_PrevRandao.TransactOpts)
}

// Sum is a paid mutator transaction binding the contract method 0x853255cc.
//
// Solidity: function sum() returns(uint256)
func (_PrevRandao *PrevRandaoTransactor) Sum(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrevRandao.contract.Transact(opts, "sum")
}

// Sum is a paid mutator transaction binding the contract method 0x853255cc.
//
// Solidity: function sum() returns(uint256)
func (_PrevRandao *PrevRandaoSession) Sum() (*types.Transaction, error) {
	return _PrevRandao.Contract.Sum(&_PrevRandao.TransactOpts)
}

// Sum is a paid mutator transaction binding the contract method 0x853255cc.
//
// Solidity: function sum() returns(uint256)
func (_PrevRandao *PrevRandaoTransactorSession) Sum() (*types.Transaction, error) {
	return _PrevRandao.Contract.Sum(&_PrevRandao.TransactOpts)
}
