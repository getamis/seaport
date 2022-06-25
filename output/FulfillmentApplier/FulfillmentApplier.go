// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FulfillmentApplier

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

// FulfillmentApplierMetaData contains all meta data concerning the FulfillmentApplier contract.
var FulfillmentApplierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidFulfillmentComponentData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchedFulfillmentOfferAndConsiderationComponents\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"}],\"name\":\"MissingFulfillmentComponentOnAggregation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferAndConsiderationRequiredOnFulfillment\",\"type\":\"error\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea26469706673582212204f64a4ea0c0428c087ff932ec81de00792621cd2104f1d2845394aca9ed6836064736f6c634300080d0033",
}

// FulfillmentApplierABI is the input ABI used to generate the binding from.
// Deprecated: Use FulfillmentApplierMetaData.ABI instead.
var FulfillmentApplierABI = FulfillmentApplierMetaData.ABI

// FulfillmentApplierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FulfillmentApplierMetaData.Bin instead.
var FulfillmentApplierBin = FulfillmentApplierMetaData.Bin

// DeployFulfillmentApplier deploys a new Ethereum contract, binding an instance of FulfillmentApplier to it.
func DeployFulfillmentApplier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FulfillmentApplier, error) {
	parsed, err := FulfillmentApplierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FulfillmentApplierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FulfillmentApplier{FulfillmentApplierCaller: FulfillmentApplierCaller{contract: contract}, FulfillmentApplierTransactor: FulfillmentApplierTransactor{contract: contract}, FulfillmentApplierFilterer: FulfillmentApplierFilterer{contract: contract}}, nil
}

// FulfillmentApplier is an auto generated Go binding around an Ethereum contract.
type FulfillmentApplier struct {
	FulfillmentApplierCaller     // Read-only binding to the contract
	FulfillmentApplierTransactor // Write-only binding to the contract
	FulfillmentApplierFilterer   // Log filterer for contract events
}

// FulfillmentApplierCaller is an auto generated read-only Go binding around an Ethereum contract.
type FulfillmentApplierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FulfillmentApplierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FulfillmentApplierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FulfillmentApplierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FulfillmentApplierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FulfillmentApplierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FulfillmentApplierSession struct {
	Contract     *FulfillmentApplier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FulfillmentApplierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FulfillmentApplierCallerSession struct {
	Contract *FulfillmentApplierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// FulfillmentApplierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FulfillmentApplierTransactorSession struct {
	Contract     *FulfillmentApplierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// FulfillmentApplierRaw is an auto generated low-level Go binding around an Ethereum contract.
type FulfillmentApplierRaw struct {
	Contract *FulfillmentApplier // Generic contract binding to access the raw methods on
}

// FulfillmentApplierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FulfillmentApplierCallerRaw struct {
	Contract *FulfillmentApplierCaller // Generic read-only contract binding to access the raw methods on
}

// FulfillmentApplierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FulfillmentApplierTransactorRaw struct {
	Contract *FulfillmentApplierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFulfillmentApplier creates a new instance of FulfillmentApplier, bound to a specific deployed contract.
func NewFulfillmentApplier(address common.Address, backend bind.ContractBackend) (*FulfillmentApplier, error) {
	contract, err := bindFulfillmentApplier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FulfillmentApplier{FulfillmentApplierCaller: FulfillmentApplierCaller{contract: contract}, FulfillmentApplierTransactor: FulfillmentApplierTransactor{contract: contract}, FulfillmentApplierFilterer: FulfillmentApplierFilterer{contract: contract}}, nil
}

// NewFulfillmentApplierCaller creates a new read-only instance of FulfillmentApplier, bound to a specific deployed contract.
func NewFulfillmentApplierCaller(address common.Address, caller bind.ContractCaller) (*FulfillmentApplierCaller, error) {
	contract, err := bindFulfillmentApplier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FulfillmentApplierCaller{contract: contract}, nil
}

// NewFulfillmentApplierTransactor creates a new write-only instance of FulfillmentApplier, bound to a specific deployed contract.
func NewFulfillmentApplierTransactor(address common.Address, transactor bind.ContractTransactor) (*FulfillmentApplierTransactor, error) {
	contract, err := bindFulfillmentApplier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FulfillmentApplierTransactor{contract: contract}, nil
}

// NewFulfillmentApplierFilterer creates a new log filterer instance of FulfillmentApplier, bound to a specific deployed contract.
func NewFulfillmentApplierFilterer(address common.Address, filterer bind.ContractFilterer) (*FulfillmentApplierFilterer, error) {
	contract, err := bindFulfillmentApplier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FulfillmentApplierFilterer{contract: contract}, nil
}

// bindFulfillmentApplier binds a generic wrapper to an already deployed contract.
func bindFulfillmentApplier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FulfillmentApplierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FulfillmentApplier *FulfillmentApplierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FulfillmentApplier.Contract.FulfillmentApplierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FulfillmentApplier *FulfillmentApplierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FulfillmentApplier.Contract.FulfillmentApplierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FulfillmentApplier *FulfillmentApplierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FulfillmentApplier.Contract.FulfillmentApplierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FulfillmentApplier *FulfillmentApplierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FulfillmentApplier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FulfillmentApplier *FulfillmentApplierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FulfillmentApplier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FulfillmentApplier *FulfillmentApplierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FulfillmentApplier.Contract.contract.Transact(opts, method, params...)
}
