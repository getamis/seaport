// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ConduitControllerInterface

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

// ConduitControllerInterfaceMetaData contains all meta data concerning the ConduitControllerInterface contract.
var ConduitControllerInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"CallerIsNotNewPotentialOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"CallerIsNotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"ChannelOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"ConduitAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCreator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newPotentialOwner\",\"type\":\"address\"}],\"name\":\"NewPotentialOwnerAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"NewPotentialOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoConduit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"NoPotentialOwnerCurrentlySet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"name\":\"NewConduit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPotentialOwner\",\"type\":\"address\"}],\"name\":\"PotentialOwnerUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"cancelOwnershipTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"createConduit\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"channelIndex\",\"type\":\"uint256\"}],\"name\":\"getChannel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"channel\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"channel\",\"type\":\"address\"}],\"name\":\"getChannelStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"getChannels\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"channels\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"name\":\"getConduit\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConduitCodeHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"creationCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"runtimeCodeHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"getPotentialOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"potentialOwner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"getTotalChannels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalChannels\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newPotentialOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"channel\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"updateChannel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ConduitControllerInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use ConduitControllerInterfaceMetaData.ABI instead.
var ConduitControllerInterfaceABI = ConduitControllerInterfaceMetaData.ABI

// ConduitControllerInterface is an auto generated Go binding around an Ethereum contract.
type ConduitControllerInterface struct {
	ConduitControllerInterfaceCaller     // Read-only binding to the contract
	ConduitControllerInterfaceTransactor // Write-only binding to the contract
	ConduitControllerInterfaceFilterer   // Log filterer for contract events
}

// ConduitControllerInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConduitControllerInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConduitControllerInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConduitControllerInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConduitControllerInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConduitControllerInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConduitControllerInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConduitControllerInterfaceSession struct {
	Contract     *ConduitControllerInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ConduitControllerInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConduitControllerInterfaceCallerSession struct {
	Contract *ConduitControllerInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ConduitControllerInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConduitControllerInterfaceTransactorSession struct {
	Contract     *ConduitControllerInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ConduitControllerInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConduitControllerInterfaceRaw struct {
	Contract *ConduitControllerInterface // Generic contract binding to access the raw methods on
}

// ConduitControllerInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConduitControllerInterfaceCallerRaw struct {
	Contract *ConduitControllerInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ConduitControllerInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConduitControllerInterfaceTransactorRaw struct {
	Contract *ConduitControllerInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConduitControllerInterface creates a new instance of ConduitControllerInterface, bound to a specific deployed contract.
func NewConduitControllerInterface(address common.Address, backend bind.ContractBackend) (*ConduitControllerInterface, error) {
	contract, err := bindConduitControllerInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConduitControllerInterface{ConduitControllerInterfaceCaller: ConduitControllerInterfaceCaller{contract: contract}, ConduitControllerInterfaceTransactor: ConduitControllerInterfaceTransactor{contract: contract}, ConduitControllerInterfaceFilterer: ConduitControllerInterfaceFilterer{contract: contract}}, nil
}

// NewConduitControllerInterfaceCaller creates a new read-only instance of ConduitControllerInterface, bound to a specific deployed contract.
func NewConduitControllerInterfaceCaller(address common.Address, caller bind.ContractCaller) (*ConduitControllerInterfaceCaller, error) {
	contract, err := bindConduitControllerInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConduitControllerInterfaceCaller{contract: contract}, nil
}

// NewConduitControllerInterfaceTransactor creates a new write-only instance of ConduitControllerInterface, bound to a specific deployed contract.
func NewConduitControllerInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ConduitControllerInterfaceTransactor, error) {
	contract, err := bindConduitControllerInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConduitControllerInterfaceTransactor{contract: contract}, nil
}

// NewConduitControllerInterfaceFilterer creates a new log filterer instance of ConduitControllerInterface, bound to a specific deployed contract.
func NewConduitControllerInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ConduitControllerInterfaceFilterer, error) {
	contract, err := bindConduitControllerInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConduitControllerInterfaceFilterer{contract: contract}, nil
}

// bindConduitControllerInterface binds a generic wrapper to an already deployed contract.
func bindConduitControllerInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConduitControllerInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConduitControllerInterface *ConduitControllerInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConduitControllerInterface.Contract.ConduitControllerInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConduitControllerInterface *ConduitControllerInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConduitControllerInterface.Contract.ConduitControllerInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConduitControllerInterface *ConduitControllerInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConduitControllerInterface.Contract.ConduitControllerInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConduitControllerInterface *ConduitControllerInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConduitControllerInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConduitControllerInterface *ConduitControllerInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConduitControllerInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConduitControllerInterface *ConduitControllerInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConduitControllerInterface.Contract.contract.Transact(opts, method, params...)
}

// GetChannel is a free data retrieval call binding the contract method 0x027cc764.
//
// Solidity: function getChannel(address conduit, uint256 channelIndex) view returns(address channel)
func (_ConduitControllerInterface *ConduitControllerInterfaceCaller) GetChannel(opts *bind.CallOpts, conduit common.Address, channelIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConduitControllerInterface.contract.Call(opts, &out, "getChannel", conduit, channelIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetChannel is a free data retrieval call binding the contract method 0x027cc764.
//
// Solidity: function getChannel(address conduit, uint256 channelIndex) view returns(address channel)
func (_ConduitControllerInterface *ConduitControllerInterfaceSession) GetChannel(conduit common.Address, channelIndex *big.Int) (common.Address, error) {
	return _ConduitControllerInterface.Contract.GetChannel(&_ConduitControllerInterface.CallOpts, conduit, channelIndex)
}

// GetChannel is a free data retrieval call binding the contract method 0x027cc764.
//
// Solidity: function getChannel(address conduit, uint256 channelIndex) view returns(address channel)
func (_ConduitControllerInterface *ConduitControllerInterfaceCallerSession) GetChannel(conduit common.Address, channelIndex *big.Int) (common.Address, error) {
	return _ConduitControllerInterface.Contract.GetChannel(&_ConduitControllerInterface.CallOpts, conduit, channelIndex)
}

// GetChannelStatus is a free data retrieval call binding the contract method 0x33bc8572.
//
// Solidity: function getChannelStatus(address conduit, address channel) view returns(bool isOpen)
func (_ConduitControllerInterface *ConduitControllerInterfaceCaller) GetChannelStatus(opts *bind.CallOpts, conduit common.Address, channel common.Address) (bool, error) {
	var out []interface{}
	err := _ConduitControllerInterface.contract.Call(opts, &out, "getChannelStatus", conduit, channel)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetChannelStatus is a free data retrieval call binding the contract method 0x33bc8572.
//
// Solidity: function getChannelStatus(address conduit, address channel) view returns(bool isOpen)
func (_ConduitControllerInterface *ConduitControllerInterfaceSession) GetChannelStatus(conduit common.Address, channel common.Address) (bool, error) {
	return _ConduitControllerInterface.Contract.GetChannelStatus(&_ConduitControllerInterface.CallOpts, conduit, channel)
}

// GetChannelStatus is a free data retrieval call binding the contract method 0x33bc8572.
//
// Solidity: function getChannelStatus(address conduit, address channel) view returns(bool isOpen)
func (_ConduitControllerInterface *ConduitControllerInterfaceCallerSession) GetChannelStatus(conduit common.Address, channel common.Address) (bool, error) {
	return _ConduitControllerInterface.Contract.GetChannelStatus(&_ConduitControllerInterface.CallOpts, conduit, channel)
}

// GetChannels is a free data retrieval call binding the contract method 0x8b9e028b.
//
// Solidity: function getChannels(address conduit) view returns(address[] channels)
func (_ConduitControllerInterface *ConduitControllerInterfaceCaller) GetChannels(opts *bind.CallOpts, conduit common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ConduitControllerInterface.contract.Call(opts, &out, "getChannels", conduit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetChannels is a free data retrieval call binding the contract method 0x8b9e028b.
//
// Solidity: function getChannels(address conduit) view returns(address[] channels)
func (_ConduitControllerInterface *ConduitControllerInterfaceSession) GetChannels(conduit common.Address) ([]common.Address, error) {
	return _ConduitControllerInterface.Contract.GetChannels(&_ConduitControllerInterface.CallOpts, conduit)
}

// GetChannels is a free data retrieval call binding the contract method 0x8b9e028b.
//
// Solidity: function getChannels(address conduit) view returns(address[] channels)
func (_ConduitControllerInterface *ConduitControllerInterfaceCallerSession) GetChannels(conduit common.Address) ([]common.Address, error) {
	return _ConduitControllerInterface.Contract.GetChannels(&_ConduitControllerInterface.CallOpts, conduit)
}

// GetConduit is a free data retrieval call binding the contract method 0x6e9bfd9f.
//
// Solidity: function getConduit(bytes32 conduitKey) view returns(address conduit, bool exists)
func (_ConduitControllerInterface *ConduitControllerInterfaceCaller) GetConduit(opts *bind.CallOpts, conduitKey [32]byte) (struct {
	Conduit common.Address
	Exists  bool
}, error) {
	var out []interface{}
	err := _ConduitControllerInterface.contract.Call(opts, &out, "getConduit", conduitKey)

	outstruct := new(struct {
		Conduit common.Address
		Exists  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Conduit = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Exists = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetConduit is a free data retrieval call binding the contract method 0x6e9bfd9f.
//
// Solidity: function getConduit(bytes32 conduitKey) view returns(address conduit, bool exists)
func (_ConduitControllerInterface *ConduitControllerInterfaceSession) GetConduit(conduitKey [32]byte) (struct {
	Conduit common.Address
	Exists  bool
}, error) {
	return _ConduitControllerInterface.Contract.GetConduit(&_ConduitControllerInterface.CallOpts, conduitKey)
}

// GetConduit is a free data retrieval call binding the contract method 0x6e9bfd9f.
//
// Solidity: function getConduit(bytes32 conduitKey) view returns(address conduit, bool exists)
func (_ConduitControllerInterface *ConduitControllerInterfaceCallerSession) GetConduit(conduitKey [32]byte) (struct {
	Conduit common.Address
	Exists  bool
}, error) {
	return _ConduitControllerInterface.Contract.GetConduit(&_ConduitControllerInterface.CallOpts, conduitKey)
}

// GetConduitCodeHashes is a free data retrieval call binding the contract method 0x0a96ad39.
//
// Solidity: function getConduitCodeHashes() view returns(bytes32 creationCodeHash, bytes32 runtimeCodeHash)
func (_ConduitControllerInterface *ConduitControllerInterfaceCaller) GetConduitCodeHashes(opts *bind.CallOpts) (struct {
	CreationCodeHash [32]byte
	RuntimeCodeHash  [32]byte
}, error) {
	var out []interface{}
	err := _ConduitControllerInterface.contract.Call(opts, &out, "getConduitCodeHashes")

	outstruct := new(struct {
		CreationCodeHash [32]byte
		RuntimeCodeHash  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CreationCodeHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.RuntimeCodeHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetConduitCodeHashes is a free data retrieval call binding the contract method 0x0a96ad39.
//
// Solidity: function getConduitCodeHashes() view returns(bytes32 creationCodeHash, bytes32 runtimeCodeHash)
func (_ConduitControllerInterface *ConduitControllerInterfaceSession) GetConduitCodeHashes() (struct {
	CreationCodeHash [32]byte
	RuntimeCodeHash  [32]byte
}, error) {
	return _ConduitControllerInterface.Contract.GetConduitCodeHashes(&_ConduitControllerInterface.CallOpts)
}

// GetConduitCodeHashes is a free data retrieval call binding the contract method 0x0a96ad39.
//
// Solidity: function getConduitCodeHashes() view returns(bytes32 creationCodeHash, bytes32 runtimeCodeHash)
func (_ConduitControllerInterface *ConduitControllerInterfaceCallerSession) GetConduitCodeHashes() (struct {
	CreationCodeHash [32]byte
	RuntimeCodeHash  [32]byte
}, error) {
	return _ConduitControllerInterface.Contract.GetConduitCodeHashes(&_ConduitControllerInterface.CallOpts)
}

// GetKey is a free data retrieval call binding the contract method 0x93790f44.
//
// Solidity: function getKey(address conduit) view returns(bytes32 conduitKey)
func (_ConduitControllerInterface *ConduitControllerInterfaceCaller) GetKey(opts *bind.CallOpts, conduit common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ConduitControllerInterface.contract.Call(opts, &out, "getKey", conduit)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKey is a free data retrieval call binding the contract method 0x93790f44.
//
// Solidity: function getKey(address conduit) view returns(bytes32 conduitKey)
func (_ConduitControllerInterface *ConduitControllerInterfaceSession) GetKey(conduit common.Address) ([32]byte, error) {
	return _ConduitControllerInterface.Contract.GetKey(&_ConduitControllerInterface.CallOpts, conduit)
}

// GetKey is a free data retrieval call binding the contract method 0x93790f44.
//
// Solidity: function getKey(address conduit) view returns(bytes32 conduitKey)
func (_ConduitControllerInterface *ConduitControllerInterfaceCallerSession) GetKey(conduit common.Address) ([32]byte, error) {
	return _ConduitControllerInterface.Contract.GetKey(&_ConduitControllerInterface.CallOpts, conduit)
}

// GetPotentialOwner is a free data retrieval call binding the contract method 0x906c87cc.
//
// Solidity: function getPotentialOwner(address conduit) view returns(address potentialOwner)
func (_ConduitControllerInterface *ConduitControllerInterfaceCaller) GetPotentialOwner(opts *bind.CallOpts, conduit common.Address) (common.Address, error) {
	var out []interface{}
	err := _ConduitControllerInterface.contract.Call(opts, &out, "getPotentialOwner", conduit)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPotentialOwner is a free data retrieval call binding the contract method 0x906c87cc.
//
// Solidity: function getPotentialOwner(address conduit) view returns(address potentialOwner)
func (_ConduitControllerInterface *ConduitControllerInterfaceSession) GetPotentialOwner(conduit common.Address) (common.Address, error) {
	return _ConduitControllerInterface.Contract.GetPotentialOwner(&_ConduitControllerInterface.CallOpts, conduit)
}

// GetPotentialOwner is a free data retrieval call binding the contract method 0x906c87cc.
//
// Solidity: function getPotentialOwner(address conduit) view returns(address potentialOwner)
func (_ConduitControllerInterface *ConduitControllerInterfaceCallerSession) GetPotentialOwner(conduit common.Address) (common.Address, error) {
	return _ConduitControllerInterface.Contract.GetPotentialOwner(&_ConduitControllerInterface.CallOpts, conduit)
}

// GetTotalChannels is a free data retrieval call binding the contract method 0x4e3f9580.
//
// Solidity: function getTotalChannels(address conduit) view returns(uint256 totalChannels)
func (_ConduitControllerInterface *ConduitControllerInterfaceCaller) GetTotalChannels(opts *bind.CallOpts, conduit common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConduitControllerInterface.contract.Call(opts, &out, "getTotalChannels", conduit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalChannels is a free data retrieval call binding the contract method 0x4e3f9580.
//
// Solidity: function getTotalChannels(address conduit) view returns(uint256 totalChannels)
func (_ConduitControllerInterface *ConduitControllerInterfaceSession) GetTotalChannels(conduit common.Address) (*big.Int, error) {
	return _ConduitControllerInterface.Contract.GetTotalChannels(&_ConduitControllerInterface.CallOpts, conduit)
}

// GetTotalChannels is a free data retrieval call binding the contract method 0x4e3f9580.
//
// Solidity: function getTotalChannels(address conduit) view returns(uint256 totalChannels)
func (_ConduitControllerInterface *ConduitControllerInterfaceCallerSession) GetTotalChannels(conduit common.Address) (*big.Int, error) {
	return _ConduitControllerInterface.Contract.GetTotalChannels(&_ConduitControllerInterface.CallOpts, conduit)
}

// OwnerOf is a free data retrieval call binding the contract method 0x14afd79e.
//
// Solidity: function ownerOf(address conduit) view returns(address owner)
func (_ConduitControllerInterface *ConduitControllerInterfaceCaller) OwnerOf(opts *bind.CallOpts, conduit common.Address) (common.Address, error) {
	var out []interface{}
	err := _ConduitControllerInterface.contract.Call(opts, &out, "ownerOf", conduit)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x14afd79e.
//
// Solidity: function ownerOf(address conduit) view returns(address owner)
func (_ConduitControllerInterface *ConduitControllerInterfaceSession) OwnerOf(conduit common.Address) (common.Address, error) {
	return _ConduitControllerInterface.Contract.OwnerOf(&_ConduitControllerInterface.CallOpts, conduit)
}

// OwnerOf is a free data retrieval call binding the contract method 0x14afd79e.
//
// Solidity: function ownerOf(address conduit) view returns(address owner)
func (_ConduitControllerInterface *ConduitControllerInterfaceCallerSession) OwnerOf(conduit common.Address) (common.Address, error) {
	return _ConduitControllerInterface.Contract.OwnerOf(&_ConduitControllerInterface.CallOpts, conduit)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x51710e45.
//
// Solidity: function acceptOwnership(address conduit) returns()
func (_ConduitControllerInterface *ConduitControllerInterfaceTransactor) AcceptOwnership(opts *bind.TransactOpts, conduit common.Address) (*types.Transaction, error) {
	return _ConduitControllerInterface.contract.Transact(opts, "acceptOwnership", conduit)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x51710e45.
//
// Solidity: function acceptOwnership(address conduit) returns()
func (_ConduitControllerInterface *ConduitControllerInterfaceSession) AcceptOwnership(conduit common.Address) (*types.Transaction, error) {
	return _ConduitControllerInterface.Contract.AcceptOwnership(&_ConduitControllerInterface.TransactOpts, conduit)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x51710e45.
//
// Solidity: function acceptOwnership(address conduit) returns()
func (_ConduitControllerInterface *ConduitControllerInterfaceTransactorSession) AcceptOwnership(conduit common.Address) (*types.Transaction, error) {
	return _ConduitControllerInterface.Contract.AcceptOwnership(&_ConduitControllerInterface.TransactOpts, conduit)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x7b37e561.
//
// Solidity: function cancelOwnershipTransfer(address conduit) returns()
func (_ConduitControllerInterface *ConduitControllerInterfaceTransactor) CancelOwnershipTransfer(opts *bind.TransactOpts, conduit common.Address) (*types.Transaction, error) {
	return _ConduitControllerInterface.contract.Transact(opts, "cancelOwnershipTransfer", conduit)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x7b37e561.
//
// Solidity: function cancelOwnershipTransfer(address conduit) returns()
func (_ConduitControllerInterface *ConduitControllerInterfaceSession) CancelOwnershipTransfer(conduit common.Address) (*types.Transaction, error) {
	return _ConduitControllerInterface.Contract.CancelOwnershipTransfer(&_ConduitControllerInterface.TransactOpts, conduit)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x7b37e561.
//
// Solidity: function cancelOwnershipTransfer(address conduit) returns()
func (_ConduitControllerInterface *ConduitControllerInterfaceTransactorSession) CancelOwnershipTransfer(conduit common.Address) (*types.Transaction, error) {
	return _ConduitControllerInterface.Contract.CancelOwnershipTransfer(&_ConduitControllerInterface.TransactOpts, conduit)
}

// CreateConduit is a paid mutator transaction binding the contract method 0x794593bc.
//
// Solidity: function createConduit(bytes32 conduitKey, address initialOwner) returns(address conduit)
func (_ConduitControllerInterface *ConduitControllerInterfaceTransactor) CreateConduit(opts *bind.TransactOpts, conduitKey [32]byte, initialOwner common.Address) (*types.Transaction, error) {
	return _ConduitControllerInterface.contract.Transact(opts, "createConduit", conduitKey, initialOwner)
}

// CreateConduit is a paid mutator transaction binding the contract method 0x794593bc.
//
// Solidity: function createConduit(bytes32 conduitKey, address initialOwner) returns(address conduit)
func (_ConduitControllerInterface *ConduitControllerInterfaceSession) CreateConduit(conduitKey [32]byte, initialOwner common.Address) (*types.Transaction, error) {
	return _ConduitControllerInterface.Contract.CreateConduit(&_ConduitControllerInterface.TransactOpts, conduitKey, initialOwner)
}

// CreateConduit is a paid mutator transaction binding the contract method 0x794593bc.
//
// Solidity: function createConduit(bytes32 conduitKey, address initialOwner) returns(address conduit)
func (_ConduitControllerInterface *ConduitControllerInterfaceTransactorSession) CreateConduit(conduitKey [32]byte, initialOwner common.Address) (*types.Transaction, error) {
	return _ConduitControllerInterface.Contract.CreateConduit(&_ConduitControllerInterface.TransactOpts, conduitKey, initialOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address conduit, address newPotentialOwner) returns()
func (_ConduitControllerInterface *ConduitControllerInterfaceTransactor) TransferOwnership(opts *bind.TransactOpts, conduit common.Address, newPotentialOwner common.Address) (*types.Transaction, error) {
	return _ConduitControllerInterface.contract.Transact(opts, "transferOwnership", conduit, newPotentialOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address conduit, address newPotentialOwner) returns()
func (_ConduitControllerInterface *ConduitControllerInterfaceSession) TransferOwnership(conduit common.Address, newPotentialOwner common.Address) (*types.Transaction, error) {
	return _ConduitControllerInterface.Contract.TransferOwnership(&_ConduitControllerInterface.TransactOpts, conduit, newPotentialOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address conduit, address newPotentialOwner) returns()
func (_ConduitControllerInterface *ConduitControllerInterfaceTransactorSession) TransferOwnership(conduit common.Address, newPotentialOwner common.Address) (*types.Transaction, error) {
	return _ConduitControllerInterface.Contract.TransferOwnership(&_ConduitControllerInterface.TransactOpts, conduit, newPotentialOwner)
}

// UpdateChannel is a paid mutator transaction binding the contract method 0x13ad9cab.
//
// Solidity: function updateChannel(address conduit, address channel, bool isOpen) returns()
func (_ConduitControllerInterface *ConduitControllerInterfaceTransactor) UpdateChannel(opts *bind.TransactOpts, conduit common.Address, channel common.Address, isOpen bool) (*types.Transaction, error) {
	return _ConduitControllerInterface.contract.Transact(opts, "updateChannel", conduit, channel, isOpen)
}

// UpdateChannel is a paid mutator transaction binding the contract method 0x13ad9cab.
//
// Solidity: function updateChannel(address conduit, address channel, bool isOpen) returns()
func (_ConduitControllerInterface *ConduitControllerInterfaceSession) UpdateChannel(conduit common.Address, channel common.Address, isOpen bool) (*types.Transaction, error) {
	return _ConduitControllerInterface.Contract.UpdateChannel(&_ConduitControllerInterface.TransactOpts, conduit, channel, isOpen)
}

// UpdateChannel is a paid mutator transaction binding the contract method 0x13ad9cab.
//
// Solidity: function updateChannel(address conduit, address channel, bool isOpen) returns()
func (_ConduitControllerInterface *ConduitControllerInterfaceTransactorSession) UpdateChannel(conduit common.Address, channel common.Address, isOpen bool) (*types.Transaction, error) {
	return _ConduitControllerInterface.Contract.UpdateChannel(&_ConduitControllerInterface.TransactOpts, conduit, channel, isOpen)
}

// ConduitControllerInterfaceNewConduitIterator is returned from FilterNewConduit and is used to iterate over the raw logs and unpacked data for NewConduit events raised by the ConduitControllerInterface contract.
type ConduitControllerInterfaceNewConduitIterator struct {
	Event *ConduitControllerInterfaceNewConduit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConduitControllerInterfaceNewConduitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConduitControllerInterfaceNewConduit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConduitControllerInterfaceNewConduit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConduitControllerInterfaceNewConduitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConduitControllerInterfaceNewConduitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConduitControllerInterfaceNewConduit represents a NewConduit event raised by the ConduitControllerInterface contract.
type ConduitControllerInterfaceNewConduit struct {
	Conduit    common.Address
	ConduitKey [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewConduit is a free log retrieval operation binding the contract event 0x4397af6128d529b8ae0442f99db1296d5136062597a15bbc61c1b2a6431a7d15.
//
// Solidity: event NewConduit(address conduit, bytes32 conduitKey)
func (_ConduitControllerInterface *ConduitControllerInterfaceFilterer) FilterNewConduit(opts *bind.FilterOpts) (*ConduitControllerInterfaceNewConduitIterator, error) {

	logs, sub, err := _ConduitControllerInterface.contract.FilterLogs(opts, "NewConduit")
	if err != nil {
		return nil, err
	}
	return &ConduitControllerInterfaceNewConduitIterator{contract: _ConduitControllerInterface.contract, event: "NewConduit", logs: logs, sub: sub}, nil
}

// WatchNewConduit is a free log subscription operation binding the contract event 0x4397af6128d529b8ae0442f99db1296d5136062597a15bbc61c1b2a6431a7d15.
//
// Solidity: event NewConduit(address conduit, bytes32 conduitKey)
func (_ConduitControllerInterface *ConduitControllerInterfaceFilterer) WatchNewConduit(opts *bind.WatchOpts, sink chan<- *ConduitControllerInterfaceNewConduit) (event.Subscription, error) {

	logs, sub, err := _ConduitControllerInterface.contract.WatchLogs(opts, "NewConduit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConduitControllerInterfaceNewConduit)
				if err := _ConduitControllerInterface.contract.UnpackLog(event, "NewConduit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewConduit is a log parse operation binding the contract event 0x4397af6128d529b8ae0442f99db1296d5136062597a15bbc61c1b2a6431a7d15.
//
// Solidity: event NewConduit(address conduit, bytes32 conduitKey)
func (_ConduitControllerInterface *ConduitControllerInterfaceFilterer) ParseNewConduit(log types.Log) (*ConduitControllerInterfaceNewConduit, error) {
	event := new(ConduitControllerInterfaceNewConduit)
	if err := _ConduitControllerInterface.contract.UnpackLog(event, "NewConduit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConduitControllerInterfaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ConduitControllerInterface contract.
type ConduitControllerInterfaceOwnershipTransferredIterator struct {
	Event *ConduitControllerInterfaceOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConduitControllerInterfaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConduitControllerInterfaceOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConduitControllerInterfaceOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConduitControllerInterfaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConduitControllerInterfaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConduitControllerInterfaceOwnershipTransferred represents a OwnershipTransferred event raised by the ConduitControllerInterface contract.
type ConduitControllerInterfaceOwnershipTransferred struct {
	Conduit       common.Address
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0xc8894f26f396ce8c004245c8b7cd1b92103a6e4302fcbab883987149ac01b7ec.
//
// Solidity: event OwnershipTransferred(address indexed conduit, address indexed previousOwner, address indexed newOwner)
func (_ConduitControllerInterface *ConduitControllerInterfaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, conduit []common.Address, previousOwner []common.Address, newOwner []common.Address) (*ConduitControllerInterfaceOwnershipTransferredIterator, error) {

	var conduitRule []interface{}
	for _, conduitItem := range conduit {
		conduitRule = append(conduitRule, conduitItem)
	}
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ConduitControllerInterface.contract.FilterLogs(opts, "OwnershipTransferred", conduitRule, previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConduitControllerInterfaceOwnershipTransferredIterator{contract: _ConduitControllerInterface.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0xc8894f26f396ce8c004245c8b7cd1b92103a6e4302fcbab883987149ac01b7ec.
//
// Solidity: event OwnershipTransferred(address indexed conduit, address indexed previousOwner, address indexed newOwner)
func (_ConduitControllerInterface *ConduitControllerInterfaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ConduitControllerInterfaceOwnershipTransferred, conduit []common.Address, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var conduitRule []interface{}
	for _, conduitItem := range conduit {
		conduitRule = append(conduitRule, conduitItem)
	}
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ConduitControllerInterface.contract.WatchLogs(opts, "OwnershipTransferred", conduitRule, previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConduitControllerInterfaceOwnershipTransferred)
				if err := _ConduitControllerInterface.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0xc8894f26f396ce8c004245c8b7cd1b92103a6e4302fcbab883987149ac01b7ec.
//
// Solidity: event OwnershipTransferred(address indexed conduit, address indexed previousOwner, address indexed newOwner)
func (_ConduitControllerInterface *ConduitControllerInterfaceFilterer) ParseOwnershipTransferred(log types.Log) (*ConduitControllerInterfaceOwnershipTransferred, error) {
	event := new(ConduitControllerInterfaceOwnershipTransferred)
	if err := _ConduitControllerInterface.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConduitControllerInterfacePotentialOwnerUpdatedIterator is returned from FilterPotentialOwnerUpdated and is used to iterate over the raw logs and unpacked data for PotentialOwnerUpdated events raised by the ConduitControllerInterface contract.
type ConduitControllerInterfacePotentialOwnerUpdatedIterator struct {
	Event *ConduitControllerInterfacePotentialOwnerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConduitControllerInterfacePotentialOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConduitControllerInterfacePotentialOwnerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConduitControllerInterfacePotentialOwnerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConduitControllerInterfacePotentialOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConduitControllerInterfacePotentialOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConduitControllerInterfacePotentialOwnerUpdated represents a PotentialOwnerUpdated event raised by the ConduitControllerInterface contract.
type ConduitControllerInterfacePotentialOwnerUpdated struct {
	NewPotentialOwner common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPotentialOwnerUpdated is a free log retrieval operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address indexed newPotentialOwner)
func (_ConduitControllerInterface *ConduitControllerInterfaceFilterer) FilterPotentialOwnerUpdated(opts *bind.FilterOpts, newPotentialOwner []common.Address) (*ConduitControllerInterfacePotentialOwnerUpdatedIterator, error) {

	var newPotentialOwnerRule []interface{}
	for _, newPotentialOwnerItem := range newPotentialOwner {
		newPotentialOwnerRule = append(newPotentialOwnerRule, newPotentialOwnerItem)
	}

	logs, sub, err := _ConduitControllerInterface.contract.FilterLogs(opts, "PotentialOwnerUpdated", newPotentialOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConduitControllerInterfacePotentialOwnerUpdatedIterator{contract: _ConduitControllerInterface.contract, event: "PotentialOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchPotentialOwnerUpdated is a free log subscription operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address indexed newPotentialOwner)
func (_ConduitControllerInterface *ConduitControllerInterfaceFilterer) WatchPotentialOwnerUpdated(opts *bind.WatchOpts, sink chan<- *ConduitControllerInterfacePotentialOwnerUpdated, newPotentialOwner []common.Address) (event.Subscription, error) {

	var newPotentialOwnerRule []interface{}
	for _, newPotentialOwnerItem := range newPotentialOwner {
		newPotentialOwnerRule = append(newPotentialOwnerRule, newPotentialOwnerItem)
	}

	logs, sub, err := _ConduitControllerInterface.contract.WatchLogs(opts, "PotentialOwnerUpdated", newPotentialOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConduitControllerInterfacePotentialOwnerUpdated)
				if err := _ConduitControllerInterface.contract.UnpackLog(event, "PotentialOwnerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePotentialOwnerUpdated is a log parse operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address indexed newPotentialOwner)
func (_ConduitControllerInterface *ConduitControllerInterfaceFilterer) ParsePotentialOwnerUpdated(log types.Log) (*ConduitControllerInterfacePotentialOwnerUpdated, error) {
	event := new(ConduitControllerInterfacePotentialOwnerUpdated)
	if err := _ConduitControllerInterface.contract.UnpackLog(event, "PotentialOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
