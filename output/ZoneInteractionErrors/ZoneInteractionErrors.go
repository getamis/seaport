// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ZoneInteractionErrors

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

// ZoneInteractionErrorsMetaData contains all meta data concerning the ZoneInteractionErrors contract.
var ZoneInteractionErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidRestrictedOrder\",\"type\":\"error\"}]",
}

// ZoneInteractionErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ZoneInteractionErrorsMetaData.ABI instead.
var ZoneInteractionErrorsABI = ZoneInteractionErrorsMetaData.ABI

// ZoneInteractionErrors is an auto generated Go binding around an Ethereum contract.
type ZoneInteractionErrors struct {
	ZoneInteractionErrorsCaller     // Read-only binding to the contract
	ZoneInteractionErrorsTransactor // Write-only binding to the contract
	ZoneInteractionErrorsFilterer   // Log filterer for contract events
}

// ZoneInteractionErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZoneInteractionErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZoneInteractionErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZoneInteractionErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZoneInteractionErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZoneInteractionErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZoneInteractionErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZoneInteractionErrorsSession struct {
	Contract     *ZoneInteractionErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ZoneInteractionErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZoneInteractionErrorsCallerSession struct {
	Contract *ZoneInteractionErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ZoneInteractionErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZoneInteractionErrorsTransactorSession struct {
	Contract     *ZoneInteractionErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ZoneInteractionErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZoneInteractionErrorsRaw struct {
	Contract *ZoneInteractionErrors // Generic contract binding to access the raw methods on
}

// ZoneInteractionErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZoneInteractionErrorsCallerRaw struct {
	Contract *ZoneInteractionErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// ZoneInteractionErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZoneInteractionErrorsTransactorRaw struct {
	Contract *ZoneInteractionErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZoneInteractionErrors creates a new instance of ZoneInteractionErrors, bound to a specific deployed contract.
func NewZoneInteractionErrors(address common.Address, backend bind.ContractBackend) (*ZoneInteractionErrors, error) {
	contract, err := bindZoneInteractionErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZoneInteractionErrors{ZoneInteractionErrorsCaller: ZoneInteractionErrorsCaller{contract: contract}, ZoneInteractionErrorsTransactor: ZoneInteractionErrorsTransactor{contract: contract}, ZoneInteractionErrorsFilterer: ZoneInteractionErrorsFilterer{contract: contract}}, nil
}

// NewZoneInteractionErrorsCaller creates a new read-only instance of ZoneInteractionErrors, bound to a specific deployed contract.
func NewZoneInteractionErrorsCaller(address common.Address, caller bind.ContractCaller) (*ZoneInteractionErrorsCaller, error) {
	contract, err := bindZoneInteractionErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZoneInteractionErrorsCaller{contract: contract}, nil
}

// NewZoneInteractionErrorsTransactor creates a new write-only instance of ZoneInteractionErrors, bound to a specific deployed contract.
func NewZoneInteractionErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ZoneInteractionErrorsTransactor, error) {
	contract, err := bindZoneInteractionErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZoneInteractionErrorsTransactor{contract: contract}, nil
}

// NewZoneInteractionErrorsFilterer creates a new log filterer instance of ZoneInteractionErrors, bound to a specific deployed contract.
func NewZoneInteractionErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ZoneInteractionErrorsFilterer, error) {
	contract, err := bindZoneInteractionErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZoneInteractionErrorsFilterer{contract: contract}, nil
}

// bindZoneInteractionErrors binds a generic wrapper to an already deployed contract.
func bindZoneInteractionErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZoneInteractionErrorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZoneInteractionErrors *ZoneInteractionErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZoneInteractionErrors.Contract.ZoneInteractionErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZoneInteractionErrors *ZoneInteractionErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZoneInteractionErrors.Contract.ZoneInteractionErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZoneInteractionErrors *ZoneInteractionErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZoneInteractionErrors.Contract.ZoneInteractionErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZoneInteractionErrors *ZoneInteractionErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZoneInteractionErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZoneInteractionErrors *ZoneInteractionErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZoneInteractionErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZoneInteractionErrors *ZoneInteractionErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZoneInteractionErrors.Contract.contract.Transact(opts, method, params...)
}
