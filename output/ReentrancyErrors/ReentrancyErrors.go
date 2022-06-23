// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ReentrancyErrors

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

// ReentrancyErrorsMetaData contains all meta data concerning the ReentrancyErrors contract.
var ReentrancyErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"NoReentrantCalls\",\"type\":\"error\"}]",
}

// ReentrancyErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ReentrancyErrorsMetaData.ABI instead.
var ReentrancyErrorsABI = ReentrancyErrorsMetaData.ABI

// ReentrancyErrors is an auto generated Go binding around an Ethereum contract.
type ReentrancyErrors struct {
	ReentrancyErrorsCaller     // Read-only binding to the contract
	ReentrancyErrorsTransactor // Write-only binding to the contract
	ReentrancyErrorsFilterer   // Log filterer for contract events
}

// ReentrancyErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyErrorsSession struct {
	Contract     *ReentrancyErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyErrorsCallerSession struct {
	Contract *ReentrancyErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ReentrancyErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyErrorsTransactorSession struct {
	Contract     *ReentrancyErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ReentrancyErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyErrorsRaw struct {
	Contract *ReentrancyErrors // Generic contract binding to access the raw methods on
}

// ReentrancyErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyErrorsCallerRaw struct {
	Contract *ReentrancyErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyErrorsTransactorRaw struct {
	Contract *ReentrancyErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyErrors creates a new instance of ReentrancyErrors, bound to a specific deployed contract.
func NewReentrancyErrors(address common.Address, backend bind.ContractBackend) (*ReentrancyErrors, error) {
	contract, err := bindReentrancyErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyErrors{ReentrancyErrorsCaller: ReentrancyErrorsCaller{contract: contract}, ReentrancyErrorsTransactor: ReentrancyErrorsTransactor{contract: contract}, ReentrancyErrorsFilterer: ReentrancyErrorsFilterer{contract: contract}}, nil
}

// NewReentrancyErrorsCaller creates a new read-only instance of ReentrancyErrors, bound to a specific deployed contract.
func NewReentrancyErrorsCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyErrorsCaller, error) {
	contract, err := bindReentrancyErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyErrorsCaller{contract: contract}, nil
}

// NewReentrancyErrorsTransactor creates a new write-only instance of ReentrancyErrors, bound to a specific deployed contract.
func NewReentrancyErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyErrorsTransactor, error) {
	contract, err := bindReentrancyErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyErrorsTransactor{contract: contract}, nil
}

// NewReentrancyErrorsFilterer creates a new log filterer instance of ReentrancyErrors, bound to a specific deployed contract.
func NewReentrancyErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyErrorsFilterer, error) {
	contract, err := bindReentrancyErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyErrorsFilterer{contract: contract}, nil
}

// bindReentrancyErrors binds a generic wrapper to an already deployed contract.
func bindReentrancyErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyErrorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyErrors *ReentrancyErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyErrors.Contract.ReentrancyErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyErrors *ReentrancyErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyErrors.Contract.ReentrancyErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyErrors *ReentrancyErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyErrors.Contract.ReentrancyErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyErrors *ReentrancyErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyErrors *ReentrancyErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyErrors *ReentrancyErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyErrors.Contract.contract.Transact(opts, method, params...)
}
