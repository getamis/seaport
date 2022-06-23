// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EIP1271Interface

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

// EIP1271InterfaceMetaData contains all meta data concerning the EIP1271Interface contract.
var EIP1271InterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EIP1271InterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use EIP1271InterfaceMetaData.ABI instead.
var EIP1271InterfaceABI = EIP1271InterfaceMetaData.ABI

// EIP1271Interface is an auto generated Go binding around an Ethereum contract.
type EIP1271Interface struct {
	EIP1271InterfaceCaller     // Read-only binding to the contract
	EIP1271InterfaceTransactor // Write-only binding to the contract
	EIP1271InterfaceFilterer   // Log filterer for contract events
}

// EIP1271InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type EIP1271InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP1271InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EIP1271InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP1271InterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EIP1271InterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP1271InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EIP1271InterfaceSession struct {
	Contract     *EIP1271Interface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIP1271InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EIP1271InterfaceCallerSession struct {
	Contract *EIP1271InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// EIP1271InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EIP1271InterfaceTransactorSession struct {
	Contract     *EIP1271InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// EIP1271InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type EIP1271InterfaceRaw struct {
	Contract *EIP1271Interface // Generic contract binding to access the raw methods on
}

// EIP1271InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EIP1271InterfaceCallerRaw struct {
	Contract *EIP1271InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// EIP1271InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EIP1271InterfaceTransactorRaw struct {
	Contract *EIP1271InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEIP1271Interface creates a new instance of EIP1271Interface, bound to a specific deployed contract.
func NewEIP1271Interface(address common.Address, backend bind.ContractBackend) (*EIP1271Interface, error) {
	contract, err := bindEIP1271Interface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EIP1271Interface{EIP1271InterfaceCaller: EIP1271InterfaceCaller{contract: contract}, EIP1271InterfaceTransactor: EIP1271InterfaceTransactor{contract: contract}, EIP1271InterfaceFilterer: EIP1271InterfaceFilterer{contract: contract}}, nil
}

// NewEIP1271InterfaceCaller creates a new read-only instance of EIP1271Interface, bound to a specific deployed contract.
func NewEIP1271InterfaceCaller(address common.Address, caller bind.ContractCaller) (*EIP1271InterfaceCaller, error) {
	contract, err := bindEIP1271Interface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EIP1271InterfaceCaller{contract: contract}, nil
}

// NewEIP1271InterfaceTransactor creates a new write-only instance of EIP1271Interface, bound to a specific deployed contract.
func NewEIP1271InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*EIP1271InterfaceTransactor, error) {
	contract, err := bindEIP1271Interface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EIP1271InterfaceTransactor{contract: contract}, nil
}

// NewEIP1271InterfaceFilterer creates a new log filterer instance of EIP1271Interface, bound to a specific deployed contract.
func NewEIP1271InterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*EIP1271InterfaceFilterer, error) {
	contract, err := bindEIP1271Interface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EIP1271InterfaceFilterer{contract: contract}, nil
}

// bindEIP1271Interface binds a generic wrapper to an already deployed contract.
func bindEIP1271Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EIP1271InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP1271Interface *EIP1271InterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP1271Interface.Contract.EIP1271InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP1271Interface *EIP1271InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP1271Interface.Contract.EIP1271InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP1271Interface *EIP1271InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP1271Interface.Contract.EIP1271InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP1271Interface *EIP1271InterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP1271Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP1271Interface *EIP1271InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP1271Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP1271Interface *EIP1271InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP1271Interface.Contract.contract.Transact(opts, method, params...)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 digest, bytes signature) view returns(bytes4)
func (_EIP1271Interface *EIP1271InterfaceCaller) IsValidSignature(opts *bind.CallOpts, digest [32]byte, signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _EIP1271Interface.contract.Call(opts, &out, "isValidSignature", digest, signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 digest, bytes signature) view returns(bytes4)
func (_EIP1271Interface *EIP1271InterfaceSession) IsValidSignature(digest [32]byte, signature []byte) ([4]byte, error) {
	return _EIP1271Interface.Contract.IsValidSignature(&_EIP1271Interface.CallOpts, digest, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 digest, bytes signature) view returns(bytes4)
func (_EIP1271Interface *EIP1271InterfaceCallerSession) IsValidSignature(digest [32]byte, signature []byte) ([4]byte, error) {
	return _EIP1271Interface.Contract.IsValidSignature(&_EIP1271Interface.CallOpts, digest, signature)
}
