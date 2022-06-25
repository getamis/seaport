// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TransferHelperInterface

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

// TransferHelperItem is an auto generated low-level Go binding around an user-defined struct.
type TransferHelperItem struct {
	ItemType   uint8
	Token      common.Address
	Identifier *big.Int
	Amount     *big.Int
}

// TransferHelperInterfaceMetaData contains all meta data concerning the TransferHelperInterface contract.
var TransferHelperInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidItemType\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumConduitItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTransferHelperItem[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"name\":\"bulkTransfer\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TransferHelperInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferHelperInterfaceMetaData.ABI instead.
var TransferHelperInterfaceABI = TransferHelperInterfaceMetaData.ABI

// TransferHelperInterface is an auto generated Go binding around an Ethereum contract.
type TransferHelperInterface struct {
	TransferHelperInterfaceCaller     // Read-only binding to the contract
	TransferHelperInterfaceTransactor // Write-only binding to the contract
	TransferHelperInterfaceFilterer   // Log filterer for contract events
}

// TransferHelperInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferHelperInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferHelperInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferHelperInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferHelperInterfaceSession struct {
	Contract     *TransferHelperInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TransferHelperInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferHelperInterfaceCallerSession struct {
	Contract *TransferHelperInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// TransferHelperInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferHelperInterfaceTransactorSession struct {
	Contract     *TransferHelperInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// TransferHelperInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferHelperInterfaceRaw struct {
	Contract *TransferHelperInterface // Generic contract binding to access the raw methods on
}

// TransferHelperInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferHelperInterfaceCallerRaw struct {
	Contract *TransferHelperInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// TransferHelperInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferHelperInterfaceTransactorRaw struct {
	Contract *TransferHelperInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferHelperInterface creates a new instance of TransferHelperInterface, bound to a specific deployed contract.
func NewTransferHelperInterface(address common.Address, backend bind.ContractBackend) (*TransferHelperInterface, error) {
	contract, err := bindTransferHelperInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferHelperInterface{TransferHelperInterfaceCaller: TransferHelperInterfaceCaller{contract: contract}, TransferHelperInterfaceTransactor: TransferHelperInterfaceTransactor{contract: contract}, TransferHelperInterfaceFilterer: TransferHelperInterfaceFilterer{contract: contract}}, nil
}

// NewTransferHelperInterfaceCaller creates a new read-only instance of TransferHelperInterface, bound to a specific deployed contract.
func NewTransferHelperInterfaceCaller(address common.Address, caller bind.ContractCaller) (*TransferHelperInterfaceCaller, error) {
	contract, err := bindTransferHelperInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperInterfaceCaller{contract: contract}, nil
}

// NewTransferHelperInterfaceTransactor creates a new write-only instance of TransferHelperInterface, bound to a specific deployed contract.
func NewTransferHelperInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferHelperInterfaceTransactor, error) {
	contract, err := bindTransferHelperInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperInterfaceTransactor{contract: contract}, nil
}

// NewTransferHelperInterfaceFilterer creates a new log filterer instance of TransferHelperInterface, bound to a specific deployed contract.
func NewTransferHelperInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferHelperInterfaceFilterer, error) {
	contract, err := bindTransferHelperInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferHelperInterfaceFilterer{contract: contract}, nil
}

// bindTransferHelperInterface binds a generic wrapper to an already deployed contract.
func bindTransferHelperInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferHelperInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelperInterface *TransferHelperInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelperInterface.Contract.TransferHelperInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelperInterface *TransferHelperInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelperInterface.Contract.TransferHelperInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelperInterface *TransferHelperInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelperInterface.Contract.TransferHelperInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelperInterface *TransferHelperInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelperInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelperInterface *TransferHelperInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelperInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelperInterface *TransferHelperInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelperInterface.Contract.contract.Transact(opts, method, params...)
}

// BulkTransfer is a paid mutator transaction binding the contract method 0x73b69d03.
//
// Solidity: function bulkTransfer((uint8,address,uint256,uint256)[] items, address recipient, bytes32 conduitKey) returns(bytes4)
func (_TransferHelperInterface *TransferHelperInterfaceTransactor) BulkTransfer(opts *bind.TransactOpts, items []TransferHelperItem, recipient common.Address, conduitKey [32]byte) (*types.Transaction, error) {
	return _TransferHelperInterface.contract.Transact(opts, "bulkTransfer", items, recipient, conduitKey)
}

// BulkTransfer is a paid mutator transaction binding the contract method 0x73b69d03.
//
// Solidity: function bulkTransfer((uint8,address,uint256,uint256)[] items, address recipient, bytes32 conduitKey) returns(bytes4)
func (_TransferHelperInterface *TransferHelperInterfaceSession) BulkTransfer(items []TransferHelperItem, recipient common.Address, conduitKey [32]byte) (*types.Transaction, error) {
	return _TransferHelperInterface.Contract.BulkTransfer(&_TransferHelperInterface.TransactOpts, items, recipient, conduitKey)
}

// BulkTransfer is a paid mutator transaction binding the contract method 0x73b69d03.
//
// Solidity: function bulkTransfer((uint8,address,uint256,uint256)[] items, address recipient, bytes32 conduitKey) returns(bytes4)
func (_TransferHelperInterface *TransferHelperInterfaceTransactorSession) BulkTransfer(items []TransferHelperItem, recipient common.Address, conduitKey [32]byte) (*types.Transaction, error) {
	return _TransferHelperInterface.Contract.BulkTransfer(&_TransferHelperInterface.TransactOpts, items, recipient, conduitKey)
}
