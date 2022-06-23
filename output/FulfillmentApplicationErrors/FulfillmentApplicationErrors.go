// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FulfillmentApplicationErrors

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

// FulfillmentApplicationErrorsMetaData contains all meta data concerning the FulfillmentApplicationErrors contract.
var FulfillmentApplicationErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidFulfillmentComponentData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchedFulfillmentOfferAndConsiderationComponents\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"}],\"name\":\"MissingFulfillmentComponentOnAggregation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferAndConsiderationRequiredOnFulfillment\",\"type\":\"error\"}]",
}

// FulfillmentApplicationErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use FulfillmentApplicationErrorsMetaData.ABI instead.
var FulfillmentApplicationErrorsABI = FulfillmentApplicationErrorsMetaData.ABI

// FulfillmentApplicationErrors is an auto generated Go binding around an Ethereum contract.
type FulfillmentApplicationErrors struct {
	FulfillmentApplicationErrorsCaller     // Read-only binding to the contract
	FulfillmentApplicationErrorsTransactor // Write-only binding to the contract
	FulfillmentApplicationErrorsFilterer   // Log filterer for contract events
}

// FulfillmentApplicationErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type FulfillmentApplicationErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FulfillmentApplicationErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FulfillmentApplicationErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FulfillmentApplicationErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FulfillmentApplicationErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FulfillmentApplicationErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FulfillmentApplicationErrorsSession struct {
	Contract     *FulfillmentApplicationErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// FulfillmentApplicationErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FulfillmentApplicationErrorsCallerSession struct {
	Contract *FulfillmentApplicationErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// FulfillmentApplicationErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FulfillmentApplicationErrorsTransactorSession struct {
	Contract     *FulfillmentApplicationErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// FulfillmentApplicationErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type FulfillmentApplicationErrorsRaw struct {
	Contract *FulfillmentApplicationErrors // Generic contract binding to access the raw methods on
}

// FulfillmentApplicationErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FulfillmentApplicationErrorsCallerRaw struct {
	Contract *FulfillmentApplicationErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// FulfillmentApplicationErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FulfillmentApplicationErrorsTransactorRaw struct {
	Contract *FulfillmentApplicationErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFulfillmentApplicationErrors creates a new instance of FulfillmentApplicationErrors, bound to a specific deployed contract.
func NewFulfillmentApplicationErrors(address common.Address, backend bind.ContractBackend) (*FulfillmentApplicationErrors, error) {
	contract, err := bindFulfillmentApplicationErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FulfillmentApplicationErrors{FulfillmentApplicationErrorsCaller: FulfillmentApplicationErrorsCaller{contract: contract}, FulfillmentApplicationErrorsTransactor: FulfillmentApplicationErrorsTransactor{contract: contract}, FulfillmentApplicationErrorsFilterer: FulfillmentApplicationErrorsFilterer{contract: contract}}, nil
}

// NewFulfillmentApplicationErrorsCaller creates a new read-only instance of FulfillmentApplicationErrors, bound to a specific deployed contract.
func NewFulfillmentApplicationErrorsCaller(address common.Address, caller bind.ContractCaller) (*FulfillmentApplicationErrorsCaller, error) {
	contract, err := bindFulfillmentApplicationErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FulfillmentApplicationErrorsCaller{contract: contract}, nil
}

// NewFulfillmentApplicationErrorsTransactor creates a new write-only instance of FulfillmentApplicationErrors, bound to a specific deployed contract.
func NewFulfillmentApplicationErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*FulfillmentApplicationErrorsTransactor, error) {
	contract, err := bindFulfillmentApplicationErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FulfillmentApplicationErrorsTransactor{contract: contract}, nil
}

// NewFulfillmentApplicationErrorsFilterer creates a new log filterer instance of FulfillmentApplicationErrors, bound to a specific deployed contract.
func NewFulfillmentApplicationErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*FulfillmentApplicationErrorsFilterer, error) {
	contract, err := bindFulfillmentApplicationErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FulfillmentApplicationErrorsFilterer{contract: contract}, nil
}

// bindFulfillmentApplicationErrors binds a generic wrapper to an already deployed contract.
func bindFulfillmentApplicationErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FulfillmentApplicationErrorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FulfillmentApplicationErrors *FulfillmentApplicationErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FulfillmentApplicationErrors.Contract.FulfillmentApplicationErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FulfillmentApplicationErrors *FulfillmentApplicationErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FulfillmentApplicationErrors.Contract.FulfillmentApplicationErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FulfillmentApplicationErrors *FulfillmentApplicationErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FulfillmentApplicationErrors.Contract.FulfillmentApplicationErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FulfillmentApplicationErrors *FulfillmentApplicationErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FulfillmentApplicationErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FulfillmentApplicationErrors *FulfillmentApplicationErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FulfillmentApplicationErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FulfillmentApplicationErrors *FulfillmentApplicationErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FulfillmentApplicationErrors.Contract.contract.Transact(opts, method, params...)
}
