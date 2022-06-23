// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SignatureVerificationErrors

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

// SignatureVerificationErrorsMetaData contains all meta data concerning the SignatureVerificationErrors contract.
var SignatureVerificationErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BadContractSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"BadSignatureV\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"}]",
}

// SignatureVerificationErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use SignatureVerificationErrorsMetaData.ABI instead.
var SignatureVerificationErrorsABI = SignatureVerificationErrorsMetaData.ABI

// SignatureVerificationErrors is an auto generated Go binding around an Ethereum contract.
type SignatureVerificationErrors struct {
	SignatureVerificationErrorsCaller     // Read-only binding to the contract
	SignatureVerificationErrorsTransactor // Write-only binding to the contract
	SignatureVerificationErrorsFilterer   // Log filterer for contract events
}

// SignatureVerificationErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignatureVerificationErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureVerificationErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignatureVerificationErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureVerificationErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignatureVerificationErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureVerificationErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignatureVerificationErrorsSession struct {
	Contract     *SignatureVerificationErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SignatureVerificationErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignatureVerificationErrorsCallerSession struct {
	Contract *SignatureVerificationErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// SignatureVerificationErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignatureVerificationErrorsTransactorSession struct {
	Contract     *SignatureVerificationErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// SignatureVerificationErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignatureVerificationErrorsRaw struct {
	Contract *SignatureVerificationErrors // Generic contract binding to access the raw methods on
}

// SignatureVerificationErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignatureVerificationErrorsCallerRaw struct {
	Contract *SignatureVerificationErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// SignatureVerificationErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignatureVerificationErrorsTransactorRaw struct {
	Contract *SignatureVerificationErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignatureVerificationErrors creates a new instance of SignatureVerificationErrors, bound to a specific deployed contract.
func NewSignatureVerificationErrors(address common.Address, backend bind.ContractBackend) (*SignatureVerificationErrors, error) {
	contract, err := bindSignatureVerificationErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignatureVerificationErrors{SignatureVerificationErrorsCaller: SignatureVerificationErrorsCaller{contract: contract}, SignatureVerificationErrorsTransactor: SignatureVerificationErrorsTransactor{contract: contract}, SignatureVerificationErrorsFilterer: SignatureVerificationErrorsFilterer{contract: contract}}, nil
}

// NewSignatureVerificationErrorsCaller creates a new read-only instance of SignatureVerificationErrors, bound to a specific deployed contract.
func NewSignatureVerificationErrorsCaller(address common.Address, caller bind.ContractCaller) (*SignatureVerificationErrorsCaller, error) {
	contract, err := bindSignatureVerificationErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureVerificationErrorsCaller{contract: contract}, nil
}

// NewSignatureVerificationErrorsTransactor creates a new write-only instance of SignatureVerificationErrors, bound to a specific deployed contract.
func NewSignatureVerificationErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*SignatureVerificationErrorsTransactor, error) {
	contract, err := bindSignatureVerificationErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureVerificationErrorsTransactor{contract: contract}, nil
}

// NewSignatureVerificationErrorsFilterer creates a new log filterer instance of SignatureVerificationErrors, bound to a specific deployed contract.
func NewSignatureVerificationErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*SignatureVerificationErrorsFilterer, error) {
	contract, err := bindSignatureVerificationErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignatureVerificationErrorsFilterer{contract: contract}, nil
}

// bindSignatureVerificationErrors binds a generic wrapper to an already deployed contract.
func bindSignatureVerificationErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SignatureVerificationErrorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureVerificationErrors *SignatureVerificationErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignatureVerificationErrors.Contract.SignatureVerificationErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureVerificationErrors *SignatureVerificationErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureVerificationErrors.Contract.SignatureVerificationErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureVerificationErrors *SignatureVerificationErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureVerificationErrors.Contract.SignatureVerificationErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureVerificationErrors *SignatureVerificationErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignatureVerificationErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureVerificationErrors *SignatureVerificationErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureVerificationErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureVerificationErrors *SignatureVerificationErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureVerificationErrors.Contract.contract.Transact(opts, method, params...)
}
