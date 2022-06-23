// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CriteriaResolutionErrors

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

// CriteriaResolutionErrorsMetaData contains all meta data concerning the CriteriaResolutionErrors contract.
var CriteriaResolutionErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ConsiderationCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CriteriaNotEnabledForItem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnresolvedConsiderationCriteria\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnresolvedOfferCriteria\",\"type\":\"error\"}]",
}

// CriteriaResolutionErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use CriteriaResolutionErrorsMetaData.ABI instead.
var CriteriaResolutionErrorsABI = CriteriaResolutionErrorsMetaData.ABI

// CriteriaResolutionErrors is an auto generated Go binding around an Ethereum contract.
type CriteriaResolutionErrors struct {
	CriteriaResolutionErrorsCaller     // Read-only binding to the contract
	CriteriaResolutionErrorsTransactor // Write-only binding to the contract
	CriteriaResolutionErrorsFilterer   // Log filterer for contract events
}

// CriteriaResolutionErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type CriteriaResolutionErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CriteriaResolutionErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CriteriaResolutionErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CriteriaResolutionErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CriteriaResolutionErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CriteriaResolutionErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CriteriaResolutionErrorsSession struct {
	Contract     *CriteriaResolutionErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CriteriaResolutionErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CriteriaResolutionErrorsCallerSession struct {
	Contract *CriteriaResolutionErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// CriteriaResolutionErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CriteriaResolutionErrorsTransactorSession struct {
	Contract     *CriteriaResolutionErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// CriteriaResolutionErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type CriteriaResolutionErrorsRaw struct {
	Contract *CriteriaResolutionErrors // Generic contract binding to access the raw methods on
}

// CriteriaResolutionErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CriteriaResolutionErrorsCallerRaw struct {
	Contract *CriteriaResolutionErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// CriteriaResolutionErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CriteriaResolutionErrorsTransactorRaw struct {
	Contract *CriteriaResolutionErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCriteriaResolutionErrors creates a new instance of CriteriaResolutionErrors, bound to a specific deployed contract.
func NewCriteriaResolutionErrors(address common.Address, backend bind.ContractBackend) (*CriteriaResolutionErrors, error) {
	contract, err := bindCriteriaResolutionErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CriteriaResolutionErrors{CriteriaResolutionErrorsCaller: CriteriaResolutionErrorsCaller{contract: contract}, CriteriaResolutionErrorsTransactor: CriteriaResolutionErrorsTransactor{contract: contract}, CriteriaResolutionErrorsFilterer: CriteriaResolutionErrorsFilterer{contract: contract}}, nil
}

// NewCriteriaResolutionErrorsCaller creates a new read-only instance of CriteriaResolutionErrors, bound to a specific deployed contract.
func NewCriteriaResolutionErrorsCaller(address common.Address, caller bind.ContractCaller) (*CriteriaResolutionErrorsCaller, error) {
	contract, err := bindCriteriaResolutionErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CriteriaResolutionErrorsCaller{contract: contract}, nil
}

// NewCriteriaResolutionErrorsTransactor creates a new write-only instance of CriteriaResolutionErrors, bound to a specific deployed contract.
func NewCriteriaResolutionErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*CriteriaResolutionErrorsTransactor, error) {
	contract, err := bindCriteriaResolutionErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CriteriaResolutionErrorsTransactor{contract: contract}, nil
}

// NewCriteriaResolutionErrorsFilterer creates a new log filterer instance of CriteriaResolutionErrors, bound to a specific deployed contract.
func NewCriteriaResolutionErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*CriteriaResolutionErrorsFilterer, error) {
	contract, err := bindCriteriaResolutionErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CriteriaResolutionErrorsFilterer{contract: contract}, nil
}

// bindCriteriaResolutionErrors binds a generic wrapper to an already deployed contract.
func bindCriteriaResolutionErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CriteriaResolutionErrorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CriteriaResolutionErrors *CriteriaResolutionErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CriteriaResolutionErrors.Contract.CriteriaResolutionErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CriteriaResolutionErrors *CriteriaResolutionErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CriteriaResolutionErrors.Contract.CriteriaResolutionErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CriteriaResolutionErrors *CriteriaResolutionErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CriteriaResolutionErrors.Contract.CriteriaResolutionErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CriteriaResolutionErrors *CriteriaResolutionErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CriteriaResolutionErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CriteriaResolutionErrors *CriteriaResolutionErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CriteriaResolutionErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CriteriaResolutionErrors *CriteriaResolutionErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CriteriaResolutionErrors.Contract.contract.Transact(opts, method, params...)
}
