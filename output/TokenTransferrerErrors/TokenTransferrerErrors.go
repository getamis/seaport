// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TokenTransferrerErrors

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

// TokenTransferrerErrorsMetaData contains all meta data concerning the TokenTransferrerErrors contract.
var TokenTransferrerErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BadReturnValueFromERC20OnTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"identifiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"ERC1155BatchTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidERC721TransferAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingItemAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NoContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenTransferGenericFailure\",\"type\":\"error\"}]",
}

// TokenTransferrerErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenTransferrerErrorsMetaData.ABI instead.
var TokenTransferrerErrorsABI = TokenTransferrerErrorsMetaData.ABI

// TokenTransferrerErrors is an auto generated Go binding around an Ethereum contract.
type TokenTransferrerErrors struct {
	TokenTransferrerErrorsCaller     // Read-only binding to the contract
	TokenTransferrerErrorsTransactor // Write-only binding to the contract
	TokenTransferrerErrorsFilterer   // Log filterer for contract events
}

// TokenTransferrerErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenTransferrerErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransferrerErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransferrerErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransferrerErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenTransferrerErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransferrerErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenTransferrerErrorsSession struct {
	Contract     *TokenTransferrerErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenTransferrerErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenTransferrerErrorsCallerSession struct {
	Contract *TokenTransferrerErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// TokenTransferrerErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransferrerErrorsTransactorSession struct {
	Contract     *TokenTransferrerErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// TokenTransferrerErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenTransferrerErrorsRaw struct {
	Contract *TokenTransferrerErrors // Generic contract binding to access the raw methods on
}

// TokenTransferrerErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenTransferrerErrorsCallerRaw struct {
	Contract *TokenTransferrerErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransferrerErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransferrerErrorsTransactorRaw struct {
	Contract *TokenTransferrerErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenTransferrerErrors creates a new instance of TokenTransferrerErrors, bound to a specific deployed contract.
func NewTokenTransferrerErrors(address common.Address, backend bind.ContractBackend) (*TokenTransferrerErrors, error) {
	contract, err := bindTokenTransferrerErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenTransferrerErrors{TokenTransferrerErrorsCaller: TokenTransferrerErrorsCaller{contract: contract}, TokenTransferrerErrorsTransactor: TokenTransferrerErrorsTransactor{contract: contract}, TokenTransferrerErrorsFilterer: TokenTransferrerErrorsFilterer{contract: contract}}, nil
}

// NewTokenTransferrerErrorsCaller creates a new read-only instance of TokenTransferrerErrors, bound to a specific deployed contract.
func NewTokenTransferrerErrorsCaller(address common.Address, caller bind.ContractCaller) (*TokenTransferrerErrorsCaller, error) {
	contract, err := bindTokenTransferrerErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransferrerErrorsCaller{contract: contract}, nil
}

// NewTokenTransferrerErrorsTransactor creates a new write-only instance of TokenTransferrerErrors, bound to a specific deployed contract.
func NewTokenTransferrerErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransferrerErrorsTransactor, error) {
	contract, err := bindTokenTransferrerErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransferrerErrorsTransactor{contract: contract}, nil
}

// NewTokenTransferrerErrorsFilterer creates a new log filterer instance of TokenTransferrerErrors, bound to a specific deployed contract.
func NewTokenTransferrerErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenTransferrerErrorsFilterer, error) {
	contract, err := bindTokenTransferrerErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenTransferrerErrorsFilterer{contract: contract}, nil
}

// bindTokenTransferrerErrors binds a generic wrapper to an already deployed contract.
func bindTokenTransferrerErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenTransferrerErrorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenTransferrerErrors *TokenTransferrerErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenTransferrerErrors.Contract.TokenTransferrerErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenTransferrerErrors *TokenTransferrerErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenTransferrerErrors.Contract.TokenTransferrerErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenTransferrerErrors *TokenTransferrerErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenTransferrerErrors.Contract.TokenTransferrerErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenTransferrerErrors *TokenTransferrerErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenTransferrerErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenTransferrerErrors *TokenTransferrerErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenTransferrerErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenTransferrerErrors *TokenTransferrerErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenTransferrerErrors.Contract.contract.Transact(opts, method, params...)
}
