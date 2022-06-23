// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ImmutableCreate2FactoryInterface

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

// ImmutableCreate2FactoryInterfaceMetaData contains all meta data concerning the ImmutableCreate2FactoryInterface contract.
var ImmutableCreate2FactoryInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"}],\"name\":\"findCreate2Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"deploymentAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"initCodeHash\",\"type\":\"bytes32\"}],\"name\":\"findCreate2AddressViaHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"deploymentAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deploymentAddress\",\"type\":\"address\"}],\"name\":\"hasBeenDeployed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"initializationCode\",\"type\":\"bytes\"}],\"name\":\"safeCreate2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"deploymentAddress\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ImmutableCreate2FactoryInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use ImmutableCreate2FactoryInterfaceMetaData.ABI instead.
var ImmutableCreate2FactoryInterfaceABI = ImmutableCreate2FactoryInterfaceMetaData.ABI

// ImmutableCreate2FactoryInterface is an auto generated Go binding around an Ethereum contract.
type ImmutableCreate2FactoryInterface struct {
	ImmutableCreate2FactoryInterfaceCaller     // Read-only binding to the contract
	ImmutableCreate2FactoryInterfaceTransactor // Write-only binding to the contract
	ImmutableCreate2FactoryInterfaceFilterer   // Log filterer for contract events
}

// ImmutableCreate2FactoryInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ImmutableCreate2FactoryInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImmutableCreate2FactoryInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ImmutableCreate2FactoryInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImmutableCreate2FactoryInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ImmutableCreate2FactoryInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImmutableCreate2FactoryInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ImmutableCreate2FactoryInterfaceSession struct {
	Contract     *ImmutableCreate2FactoryInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ImmutableCreate2FactoryInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ImmutableCreate2FactoryInterfaceCallerSession struct {
	Contract *ImmutableCreate2FactoryInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                           // Call options to use throughout this session
}

// ImmutableCreate2FactoryInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ImmutableCreate2FactoryInterfaceTransactorSession struct {
	Contract     *ImmutableCreate2FactoryInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// ImmutableCreate2FactoryInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ImmutableCreate2FactoryInterfaceRaw struct {
	Contract *ImmutableCreate2FactoryInterface // Generic contract binding to access the raw methods on
}

// ImmutableCreate2FactoryInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ImmutableCreate2FactoryInterfaceCallerRaw struct {
	Contract *ImmutableCreate2FactoryInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ImmutableCreate2FactoryInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ImmutableCreate2FactoryInterfaceTransactorRaw struct {
	Contract *ImmutableCreate2FactoryInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewImmutableCreate2FactoryInterface creates a new instance of ImmutableCreate2FactoryInterface, bound to a specific deployed contract.
func NewImmutableCreate2FactoryInterface(address common.Address, backend bind.ContractBackend) (*ImmutableCreate2FactoryInterface, error) {
	contract, err := bindImmutableCreate2FactoryInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ImmutableCreate2FactoryInterface{ImmutableCreate2FactoryInterfaceCaller: ImmutableCreate2FactoryInterfaceCaller{contract: contract}, ImmutableCreate2FactoryInterfaceTransactor: ImmutableCreate2FactoryInterfaceTransactor{contract: contract}, ImmutableCreate2FactoryInterfaceFilterer: ImmutableCreate2FactoryInterfaceFilterer{contract: contract}}, nil
}

// NewImmutableCreate2FactoryInterfaceCaller creates a new read-only instance of ImmutableCreate2FactoryInterface, bound to a specific deployed contract.
func NewImmutableCreate2FactoryInterfaceCaller(address common.Address, caller bind.ContractCaller) (*ImmutableCreate2FactoryInterfaceCaller, error) {
	contract, err := bindImmutableCreate2FactoryInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ImmutableCreate2FactoryInterfaceCaller{contract: contract}, nil
}

// NewImmutableCreate2FactoryInterfaceTransactor creates a new write-only instance of ImmutableCreate2FactoryInterface, bound to a specific deployed contract.
func NewImmutableCreate2FactoryInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ImmutableCreate2FactoryInterfaceTransactor, error) {
	contract, err := bindImmutableCreate2FactoryInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ImmutableCreate2FactoryInterfaceTransactor{contract: contract}, nil
}

// NewImmutableCreate2FactoryInterfaceFilterer creates a new log filterer instance of ImmutableCreate2FactoryInterface, bound to a specific deployed contract.
func NewImmutableCreate2FactoryInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ImmutableCreate2FactoryInterfaceFilterer, error) {
	contract, err := bindImmutableCreate2FactoryInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ImmutableCreate2FactoryInterfaceFilterer{contract: contract}, nil
}

// bindImmutableCreate2FactoryInterface binds a generic wrapper to an already deployed contract.
func bindImmutableCreate2FactoryInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ImmutableCreate2FactoryInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ImmutableCreate2FactoryInterface.Contract.ImmutableCreate2FactoryInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ImmutableCreate2FactoryInterface.Contract.ImmutableCreate2FactoryInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ImmutableCreate2FactoryInterface.Contract.ImmutableCreate2FactoryInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ImmutableCreate2FactoryInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ImmutableCreate2FactoryInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ImmutableCreate2FactoryInterface.Contract.contract.Transact(opts, method, params...)
}

// FindCreate2Address is a free data retrieval call binding the contract method 0x85cf97ab.
//
// Solidity: function findCreate2Address(bytes32 salt, bytes initCode) view returns(address deploymentAddress)
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceCaller) FindCreate2Address(opts *bind.CallOpts, salt [32]byte, initCode []byte) (common.Address, error) {
	var out []interface{}
	err := _ImmutableCreate2FactoryInterface.contract.Call(opts, &out, "findCreate2Address", salt, initCode)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindCreate2Address is a free data retrieval call binding the contract method 0x85cf97ab.
//
// Solidity: function findCreate2Address(bytes32 salt, bytes initCode) view returns(address deploymentAddress)
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceSession) FindCreate2Address(salt [32]byte, initCode []byte) (common.Address, error) {
	return _ImmutableCreate2FactoryInterface.Contract.FindCreate2Address(&_ImmutableCreate2FactoryInterface.CallOpts, salt, initCode)
}

// FindCreate2Address is a free data retrieval call binding the contract method 0x85cf97ab.
//
// Solidity: function findCreate2Address(bytes32 salt, bytes initCode) view returns(address deploymentAddress)
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceCallerSession) FindCreate2Address(salt [32]byte, initCode []byte) (common.Address, error) {
	return _ImmutableCreate2FactoryInterface.Contract.FindCreate2Address(&_ImmutableCreate2FactoryInterface.CallOpts, salt, initCode)
}

// FindCreate2AddressViaHash is a free data retrieval call binding the contract method 0xa49a7c90.
//
// Solidity: function findCreate2AddressViaHash(bytes32 salt, bytes32 initCodeHash) view returns(address deploymentAddress)
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceCaller) FindCreate2AddressViaHash(opts *bind.CallOpts, salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ImmutableCreate2FactoryInterface.contract.Call(opts, &out, "findCreate2AddressViaHash", salt, initCodeHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindCreate2AddressViaHash is a free data retrieval call binding the contract method 0xa49a7c90.
//
// Solidity: function findCreate2AddressViaHash(bytes32 salt, bytes32 initCodeHash) view returns(address deploymentAddress)
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceSession) FindCreate2AddressViaHash(salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	return _ImmutableCreate2FactoryInterface.Contract.FindCreate2AddressViaHash(&_ImmutableCreate2FactoryInterface.CallOpts, salt, initCodeHash)
}

// FindCreate2AddressViaHash is a free data retrieval call binding the contract method 0xa49a7c90.
//
// Solidity: function findCreate2AddressViaHash(bytes32 salt, bytes32 initCodeHash) view returns(address deploymentAddress)
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceCallerSession) FindCreate2AddressViaHash(salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	return _ImmutableCreate2FactoryInterface.Contract.FindCreate2AddressViaHash(&_ImmutableCreate2FactoryInterface.CallOpts, salt, initCodeHash)
}

// HasBeenDeployed is a free data retrieval call binding the contract method 0x08508b8f.
//
// Solidity: function hasBeenDeployed(address deploymentAddress) view returns(bool)
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceCaller) HasBeenDeployed(opts *bind.CallOpts, deploymentAddress common.Address) (bool, error) {
	var out []interface{}
	err := _ImmutableCreate2FactoryInterface.contract.Call(opts, &out, "hasBeenDeployed", deploymentAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasBeenDeployed is a free data retrieval call binding the contract method 0x08508b8f.
//
// Solidity: function hasBeenDeployed(address deploymentAddress) view returns(bool)
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceSession) HasBeenDeployed(deploymentAddress common.Address) (bool, error) {
	return _ImmutableCreate2FactoryInterface.Contract.HasBeenDeployed(&_ImmutableCreate2FactoryInterface.CallOpts, deploymentAddress)
}

// HasBeenDeployed is a free data retrieval call binding the contract method 0x08508b8f.
//
// Solidity: function hasBeenDeployed(address deploymentAddress) view returns(bool)
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceCallerSession) HasBeenDeployed(deploymentAddress common.Address) (bool, error) {
	return _ImmutableCreate2FactoryInterface.Contract.HasBeenDeployed(&_ImmutableCreate2FactoryInterface.CallOpts, deploymentAddress)
}

// SafeCreate2 is a paid mutator transaction binding the contract method 0x64e03087.
//
// Solidity: function safeCreate2(bytes32 salt, bytes initializationCode) payable returns(address deploymentAddress)
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceTransactor) SafeCreate2(opts *bind.TransactOpts, salt [32]byte, initializationCode []byte) (*types.Transaction, error) {
	return _ImmutableCreate2FactoryInterface.contract.Transact(opts, "safeCreate2", salt, initializationCode)
}

// SafeCreate2 is a paid mutator transaction binding the contract method 0x64e03087.
//
// Solidity: function safeCreate2(bytes32 salt, bytes initializationCode) payable returns(address deploymentAddress)
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceSession) SafeCreate2(salt [32]byte, initializationCode []byte) (*types.Transaction, error) {
	return _ImmutableCreate2FactoryInterface.Contract.SafeCreate2(&_ImmutableCreate2FactoryInterface.TransactOpts, salt, initializationCode)
}

// SafeCreate2 is a paid mutator transaction binding the contract method 0x64e03087.
//
// Solidity: function safeCreate2(bytes32 salt, bytes initializationCode) payable returns(address deploymentAddress)
func (_ImmutableCreate2FactoryInterface *ImmutableCreate2FactoryInterfaceTransactorSession) SafeCreate2(salt [32]byte, initializationCode []byte) (*types.Transaction, error) {
	return _ImmutableCreate2FactoryInterface.Contract.SafeCreate2(&_ImmutableCreate2FactoryInterface.TransactOpts, salt, initializationCode)
}
