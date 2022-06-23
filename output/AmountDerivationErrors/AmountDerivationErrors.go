// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AmountDerivationErrors

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

// AmountDerivationErrorsMetaData contains all meta data concerning the AmountDerivationErrors contract.
var AmountDerivationErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InexactFraction\",\"type\":\"error\"}]",
}

// AmountDerivationErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use AmountDerivationErrorsMetaData.ABI instead.
var AmountDerivationErrorsABI = AmountDerivationErrorsMetaData.ABI

// AmountDerivationErrors is an auto generated Go binding around an Ethereum contract.
type AmountDerivationErrors struct {
	AmountDerivationErrorsCaller     // Read-only binding to the contract
	AmountDerivationErrorsTransactor // Write-only binding to the contract
	AmountDerivationErrorsFilterer   // Log filterer for contract events
}

// AmountDerivationErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AmountDerivationErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmountDerivationErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AmountDerivationErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmountDerivationErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AmountDerivationErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmountDerivationErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AmountDerivationErrorsSession struct {
	Contract     *AmountDerivationErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AmountDerivationErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AmountDerivationErrorsCallerSession struct {
	Contract *AmountDerivationErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// AmountDerivationErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AmountDerivationErrorsTransactorSession struct {
	Contract     *AmountDerivationErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// AmountDerivationErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AmountDerivationErrorsRaw struct {
	Contract *AmountDerivationErrors // Generic contract binding to access the raw methods on
}

// AmountDerivationErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AmountDerivationErrorsCallerRaw struct {
	Contract *AmountDerivationErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// AmountDerivationErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AmountDerivationErrorsTransactorRaw struct {
	Contract *AmountDerivationErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAmountDerivationErrors creates a new instance of AmountDerivationErrors, bound to a specific deployed contract.
func NewAmountDerivationErrors(address common.Address, backend bind.ContractBackend) (*AmountDerivationErrors, error) {
	contract, err := bindAmountDerivationErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AmountDerivationErrors{AmountDerivationErrorsCaller: AmountDerivationErrorsCaller{contract: contract}, AmountDerivationErrorsTransactor: AmountDerivationErrorsTransactor{contract: contract}, AmountDerivationErrorsFilterer: AmountDerivationErrorsFilterer{contract: contract}}, nil
}

// NewAmountDerivationErrorsCaller creates a new read-only instance of AmountDerivationErrors, bound to a specific deployed contract.
func NewAmountDerivationErrorsCaller(address common.Address, caller bind.ContractCaller) (*AmountDerivationErrorsCaller, error) {
	contract, err := bindAmountDerivationErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AmountDerivationErrorsCaller{contract: contract}, nil
}

// NewAmountDerivationErrorsTransactor creates a new write-only instance of AmountDerivationErrors, bound to a specific deployed contract.
func NewAmountDerivationErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*AmountDerivationErrorsTransactor, error) {
	contract, err := bindAmountDerivationErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AmountDerivationErrorsTransactor{contract: contract}, nil
}

// NewAmountDerivationErrorsFilterer creates a new log filterer instance of AmountDerivationErrors, bound to a specific deployed contract.
func NewAmountDerivationErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*AmountDerivationErrorsFilterer, error) {
	contract, err := bindAmountDerivationErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AmountDerivationErrorsFilterer{contract: contract}, nil
}

// bindAmountDerivationErrors binds a generic wrapper to an already deployed contract.
func bindAmountDerivationErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AmountDerivationErrorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AmountDerivationErrors *AmountDerivationErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AmountDerivationErrors.Contract.AmountDerivationErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AmountDerivationErrors *AmountDerivationErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AmountDerivationErrors.Contract.AmountDerivationErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AmountDerivationErrors *AmountDerivationErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AmountDerivationErrors.Contract.AmountDerivationErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AmountDerivationErrors *AmountDerivationErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AmountDerivationErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AmountDerivationErrors *AmountDerivationErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AmountDerivationErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AmountDerivationErrors *AmountDerivationErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AmountDerivationErrors.Contract.contract.Transact(opts, method, params...)
}
