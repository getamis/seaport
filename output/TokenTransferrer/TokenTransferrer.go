// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TokenTransferrer

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

// TokenTransferrerMetaData contains all meta data concerning the TokenTransferrer contract.
var TokenTransferrerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BadReturnValueFromERC20OnTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"identifiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"ERC1155BatchTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidERC721TransferAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingItemAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NoContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenTransferGenericFailure\",\"type\":\"error\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220074c37e07fd65d0f23dc6bc3414422f91f71d24239a6c799a906b8265960572564736f6c634300080d0033",
}

// TokenTransferrerABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenTransferrerMetaData.ABI instead.
var TokenTransferrerABI = TokenTransferrerMetaData.ABI

// TokenTransferrerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenTransferrerMetaData.Bin instead.
var TokenTransferrerBin = TokenTransferrerMetaData.Bin

// DeployTokenTransferrer deploys a new Ethereum contract, binding an instance of TokenTransferrer to it.
func DeployTokenTransferrer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenTransferrer, error) {
	parsed, err := TokenTransferrerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenTransferrerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenTransferrer{TokenTransferrerCaller: TokenTransferrerCaller{contract: contract}, TokenTransferrerTransactor: TokenTransferrerTransactor{contract: contract}, TokenTransferrerFilterer: TokenTransferrerFilterer{contract: contract}}, nil
}

// TokenTransferrer is an auto generated Go binding around an Ethereum contract.
type TokenTransferrer struct {
	TokenTransferrerCaller     // Read-only binding to the contract
	TokenTransferrerTransactor // Write-only binding to the contract
	TokenTransferrerFilterer   // Log filterer for contract events
}

// TokenTransferrerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenTransferrerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransferrerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransferrerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransferrerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenTransferrerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransferrerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenTransferrerSession struct {
	Contract     *TokenTransferrer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenTransferrerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenTransferrerCallerSession struct {
	Contract *TokenTransferrerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TokenTransferrerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransferrerTransactorSession struct {
	Contract     *TokenTransferrerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TokenTransferrerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenTransferrerRaw struct {
	Contract *TokenTransferrer // Generic contract binding to access the raw methods on
}

// TokenTransferrerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenTransferrerCallerRaw struct {
	Contract *TokenTransferrerCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransferrerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransferrerTransactorRaw struct {
	Contract *TokenTransferrerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenTransferrer creates a new instance of TokenTransferrer, bound to a specific deployed contract.
func NewTokenTransferrer(address common.Address, backend bind.ContractBackend) (*TokenTransferrer, error) {
	contract, err := bindTokenTransferrer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenTransferrer{TokenTransferrerCaller: TokenTransferrerCaller{contract: contract}, TokenTransferrerTransactor: TokenTransferrerTransactor{contract: contract}, TokenTransferrerFilterer: TokenTransferrerFilterer{contract: contract}}, nil
}

// NewTokenTransferrerCaller creates a new read-only instance of TokenTransferrer, bound to a specific deployed contract.
func NewTokenTransferrerCaller(address common.Address, caller bind.ContractCaller) (*TokenTransferrerCaller, error) {
	contract, err := bindTokenTransferrer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransferrerCaller{contract: contract}, nil
}

// NewTokenTransferrerTransactor creates a new write-only instance of TokenTransferrer, bound to a specific deployed contract.
func NewTokenTransferrerTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransferrerTransactor, error) {
	contract, err := bindTokenTransferrer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransferrerTransactor{contract: contract}, nil
}

// NewTokenTransferrerFilterer creates a new log filterer instance of TokenTransferrer, bound to a specific deployed contract.
func NewTokenTransferrerFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenTransferrerFilterer, error) {
	contract, err := bindTokenTransferrer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenTransferrerFilterer{contract: contract}, nil
}

// bindTokenTransferrer binds a generic wrapper to an already deployed contract.
func bindTokenTransferrer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenTransferrerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenTransferrer *TokenTransferrerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenTransferrer.Contract.TokenTransferrerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenTransferrer *TokenTransferrerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenTransferrer.Contract.TokenTransferrerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenTransferrer *TokenTransferrerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenTransferrer.Contract.TokenTransferrerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenTransferrer *TokenTransferrerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenTransferrer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenTransferrer *TokenTransferrerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenTransferrer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenTransferrer *TokenTransferrerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenTransferrer.Contract.contract.Transact(opts, method, params...)
}
