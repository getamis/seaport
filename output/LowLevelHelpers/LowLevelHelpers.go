// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LowLevelHelpers

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

// LowLevelHelpersMetaData contains all meta data concerning the LowLevelHelpers contract.
var LowLevelHelpersMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220908019618b5c8f332d1672731b394c6bdab76f4dc3c2175837e2dc422a04888a64736f6c634300080d0033",
}

// LowLevelHelpersABI is the input ABI used to generate the binding from.
// Deprecated: Use LowLevelHelpersMetaData.ABI instead.
var LowLevelHelpersABI = LowLevelHelpersMetaData.ABI

// LowLevelHelpersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LowLevelHelpersMetaData.Bin instead.
var LowLevelHelpersBin = LowLevelHelpersMetaData.Bin

// DeployLowLevelHelpers deploys a new Ethereum contract, binding an instance of LowLevelHelpers to it.
func DeployLowLevelHelpers(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LowLevelHelpers, error) {
	parsed, err := LowLevelHelpersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LowLevelHelpersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LowLevelHelpers{LowLevelHelpersCaller: LowLevelHelpersCaller{contract: contract}, LowLevelHelpersTransactor: LowLevelHelpersTransactor{contract: contract}, LowLevelHelpersFilterer: LowLevelHelpersFilterer{contract: contract}}, nil
}

// LowLevelHelpers is an auto generated Go binding around an Ethereum contract.
type LowLevelHelpers struct {
	LowLevelHelpersCaller     // Read-only binding to the contract
	LowLevelHelpersTransactor // Write-only binding to the contract
	LowLevelHelpersFilterer   // Log filterer for contract events
}

// LowLevelHelpersCaller is an auto generated read-only Go binding around an Ethereum contract.
type LowLevelHelpersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LowLevelHelpersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LowLevelHelpersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LowLevelHelpersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LowLevelHelpersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LowLevelHelpersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LowLevelHelpersSession struct {
	Contract     *LowLevelHelpers  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LowLevelHelpersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LowLevelHelpersCallerSession struct {
	Contract *LowLevelHelpersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// LowLevelHelpersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LowLevelHelpersTransactorSession struct {
	Contract     *LowLevelHelpersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// LowLevelHelpersRaw is an auto generated low-level Go binding around an Ethereum contract.
type LowLevelHelpersRaw struct {
	Contract *LowLevelHelpers // Generic contract binding to access the raw methods on
}

// LowLevelHelpersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LowLevelHelpersCallerRaw struct {
	Contract *LowLevelHelpersCaller // Generic read-only contract binding to access the raw methods on
}

// LowLevelHelpersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LowLevelHelpersTransactorRaw struct {
	Contract *LowLevelHelpersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLowLevelHelpers creates a new instance of LowLevelHelpers, bound to a specific deployed contract.
func NewLowLevelHelpers(address common.Address, backend bind.ContractBackend) (*LowLevelHelpers, error) {
	contract, err := bindLowLevelHelpers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LowLevelHelpers{LowLevelHelpersCaller: LowLevelHelpersCaller{contract: contract}, LowLevelHelpersTransactor: LowLevelHelpersTransactor{contract: contract}, LowLevelHelpersFilterer: LowLevelHelpersFilterer{contract: contract}}, nil
}

// NewLowLevelHelpersCaller creates a new read-only instance of LowLevelHelpers, bound to a specific deployed contract.
func NewLowLevelHelpersCaller(address common.Address, caller bind.ContractCaller) (*LowLevelHelpersCaller, error) {
	contract, err := bindLowLevelHelpers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LowLevelHelpersCaller{contract: contract}, nil
}

// NewLowLevelHelpersTransactor creates a new write-only instance of LowLevelHelpers, bound to a specific deployed contract.
func NewLowLevelHelpersTransactor(address common.Address, transactor bind.ContractTransactor) (*LowLevelHelpersTransactor, error) {
	contract, err := bindLowLevelHelpers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LowLevelHelpersTransactor{contract: contract}, nil
}

// NewLowLevelHelpersFilterer creates a new log filterer instance of LowLevelHelpers, bound to a specific deployed contract.
func NewLowLevelHelpersFilterer(address common.Address, filterer bind.ContractFilterer) (*LowLevelHelpersFilterer, error) {
	contract, err := bindLowLevelHelpers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LowLevelHelpersFilterer{contract: contract}, nil
}

// bindLowLevelHelpers binds a generic wrapper to an already deployed contract.
func bindLowLevelHelpers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LowLevelHelpersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LowLevelHelpers *LowLevelHelpersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LowLevelHelpers.Contract.LowLevelHelpersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LowLevelHelpers *LowLevelHelpersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LowLevelHelpers.Contract.LowLevelHelpersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LowLevelHelpers *LowLevelHelpersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LowLevelHelpers.Contract.LowLevelHelpersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LowLevelHelpers *LowLevelHelpersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LowLevelHelpers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LowLevelHelpers *LowLevelHelpersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LowLevelHelpers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LowLevelHelpers *LowLevelHelpersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LowLevelHelpers.Contract.contract.Transact(opts, method, params...)
}
