// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AmountDeriver

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
)

// AmountDeriverMetaData contains all meta data concerning the AmountDeriver contract.
var AmountDeriverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InexactFraction\",\"type\":\"error\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220e9fd663fd7346b0953bb9543d613f8d5c6b257b24811537763c43b3759ee785964736f6c634300080d0033",
}

// AmountDeriverABI is the input ABI used to generate the binding from.
// Deprecated: Use AmountDeriverMetaData.ABI instead.
var AmountDeriverABI = AmountDeriverMetaData.ABI

// AmountDeriverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AmountDeriverMetaData.Bin instead.
var AmountDeriverBin = AmountDeriverMetaData.Bin

// DeployAmountDeriver deploys a new Ethereum contract, binding an instance of AmountDeriver to it.
func DeployAmountDeriver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AmountDeriver, error) {
	parsed, err := AmountDeriverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AmountDeriverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AmountDeriver{AmountDeriverCaller: AmountDeriverCaller{contract: contract}, AmountDeriverTransactor: AmountDeriverTransactor{contract: contract}, AmountDeriverFilterer: AmountDeriverFilterer{contract: contract}}, nil
}

// AmountDeriver is an auto generated Go binding around an Ethereum contract.
type AmountDeriver struct {
	AmountDeriverCaller     // Read-only binding to the contract
	AmountDeriverTransactor // Write-only binding to the contract
	AmountDeriverFilterer   // Log filterer for contract events
}

// AmountDeriverCaller is an auto generated read-only Go binding around an Ethereum contract.
type AmountDeriverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmountDeriverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AmountDeriverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmountDeriverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AmountDeriverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmountDeriverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AmountDeriverSession struct {
	Contract     *AmountDeriver    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AmountDeriverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AmountDeriverCallerSession struct {
	Contract *AmountDeriverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AmountDeriverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AmountDeriverTransactorSession struct {
	Contract     *AmountDeriverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AmountDeriverRaw is an auto generated low-level Go binding around an Ethereum contract.
type AmountDeriverRaw struct {
	Contract *AmountDeriver // Generic contract binding to access the raw methods on
}

// AmountDeriverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AmountDeriverCallerRaw struct {
	Contract *AmountDeriverCaller // Generic read-only contract binding to access the raw methods on
}

// AmountDeriverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AmountDeriverTransactorRaw struct {
	Contract *AmountDeriverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAmountDeriver creates a new instance of AmountDeriver, bound to a specific deployed contract.
func NewAmountDeriver(address common.Address, backend bind.ContractBackend) (*AmountDeriver, error) {
	contract, err := bindAmountDeriver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AmountDeriver{AmountDeriverCaller: AmountDeriverCaller{contract: contract}, AmountDeriverTransactor: AmountDeriverTransactor{contract: contract}, AmountDeriverFilterer: AmountDeriverFilterer{contract: contract}}, nil
}

// NewAmountDeriverCaller creates a new read-only instance of AmountDeriver, bound to a specific deployed contract.
func NewAmountDeriverCaller(address common.Address, caller bind.ContractCaller) (*AmountDeriverCaller, error) {
	contract, err := bindAmountDeriver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AmountDeriverCaller{contract: contract}, nil
}

// NewAmountDeriverTransactor creates a new write-only instance of AmountDeriver, bound to a specific deployed contract.
func NewAmountDeriverTransactor(address common.Address, transactor bind.ContractTransactor) (*AmountDeriverTransactor, error) {
	contract, err := bindAmountDeriver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AmountDeriverTransactor{contract: contract}, nil
}

// NewAmountDeriverFilterer creates a new log filterer instance of AmountDeriver, bound to a specific deployed contract.
func NewAmountDeriverFilterer(address common.Address, filterer bind.ContractFilterer) (*AmountDeriverFilterer, error) {
	contract, err := bindAmountDeriver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AmountDeriverFilterer{contract: contract}, nil
}

// bindAmountDeriver binds a generic wrapper to an already deployed contract.
func bindAmountDeriver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AmountDeriverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AmountDeriver *AmountDeriverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AmountDeriver.Contract.AmountDeriverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AmountDeriver *AmountDeriverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AmountDeriver.Contract.AmountDeriverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AmountDeriver *AmountDeriverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AmountDeriver.Contract.AmountDeriverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AmountDeriver *AmountDeriverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AmountDeriver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AmountDeriver *AmountDeriverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AmountDeriver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AmountDeriver *AmountDeriverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AmountDeriver.Contract.contract.Transact(opts, method, params...)
}
