// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ZoneInteraction

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

// ZoneInteractionMetaData contains all meta data concerning the ZoneInteraction contract.
var ZoneInteractionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidRestrictedOrder\",\"type\":\"error\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea264697066735822122058ced0874bfe5c1efcab412bc8419b7da028afd238395a0cf81ef0050162ecc164736f6c634300080d0033",
}

// ZoneInteractionABI is the input ABI used to generate the binding from.
// Deprecated: Use ZoneInteractionMetaData.ABI instead.
var ZoneInteractionABI = ZoneInteractionMetaData.ABI

// ZoneInteractionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZoneInteractionMetaData.Bin instead.
var ZoneInteractionBin = ZoneInteractionMetaData.Bin

// DeployZoneInteraction deploys a new Ethereum contract, binding an instance of ZoneInteraction to it.
func DeployZoneInteraction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZoneInteraction, error) {
	parsed, err := ZoneInteractionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZoneInteractionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZoneInteraction{ZoneInteractionCaller: ZoneInteractionCaller{contract: contract}, ZoneInteractionTransactor: ZoneInteractionTransactor{contract: contract}, ZoneInteractionFilterer: ZoneInteractionFilterer{contract: contract}}, nil
}

// ZoneInteraction is an auto generated Go binding around an Ethereum contract.
type ZoneInteraction struct {
	ZoneInteractionCaller     // Read-only binding to the contract
	ZoneInteractionTransactor // Write-only binding to the contract
	ZoneInteractionFilterer   // Log filterer for contract events
}

// ZoneInteractionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZoneInteractionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZoneInteractionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZoneInteractionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZoneInteractionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZoneInteractionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZoneInteractionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZoneInteractionSession struct {
	Contract     *ZoneInteraction  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZoneInteractionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZoneInteractionCallerSession struct {
	Contract *ZoneInteractionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ZoneInteractionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZoneInteractionTransactorSession struct {
	Contract     *ZoneInteractionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ZoneInteractionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZoneInteractionRaw struct {
	Contract *ZoneInteraction // Generic contract binding to access the raw methods on
}

// ZoneInteractionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZoneInteractionCallerRaw struct {
	Contract *ZoneInteractionCaller // Generic read-only contract binding to access the raw methods on
}

// ZoneInteractionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZoneInteractionTransactorRaw struct {
	Contract *ZoneInteractionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZoneInteraction creates a new instance of ZoneInteraction, bound to a specific deployed contract.
func NewZoneInteraction(address common.Address, backend bind.ContractBackend) (*ZoneInteraction, error) {
	contract, err := bindZoneInteraction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZoneInteraction{ZoneInteractionCaller: ZoneInteractionCaller{contract: contract}, ZoneInteractionTransactor: ZoneInteractionTransactor{contract: contract}, ZoneInteractionFilterer: ZoneInteractionFilterer{contract: contract}}, nil
}

// NewZoneInteractionCaller creates a new read-only instance of ZoneInteraction, bound to a specific deployed contract.
func NewZoneInteractionCaller(address common.Address, caller bind.ContractCaller) (*ZoneInteractionCaller, error) {
	contract, err := bindZoneInteraction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZoneInteractionCaller{contract: contract}, nil
}

// NewZoneInteractionTransactor creates a new write-only instance of ZoneInteraction, bound to a specific deployed contract.
func NewZoneInteractionTransactor(address common.Address, transactor bind.ContractTransactor) (*ZoneInteractionTransactor, error) {
	contract, err := bindZoneInteraction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZoneInteractionTransactor{contract: contract}, nil
}

// NewZoneInteractionFilterer creates a new log filterer instance of ZoneInteraction, bound to a specific deployed contract.
func NewZoneInteractionFilterer(address common.Address, filterer bind.ContractFilterer) (*ZoneInteractionFilterer, error) {
	contract, err := bindZoneInteraction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZoneInteractionFilterer{contract: contract}, nil
}

// bindZoneInteraction binds a generic wrapper to an already deployed contract.
func bindZoneInteraction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZoneInteractionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZoneInteraction *ZoneInteractionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZoneInteraction.Contract.ZoneInteractionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZoneInteraction *ZoneInteractionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZoneInteraction.Contract.ZoneInteractionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZoneInteraction *ZoneInteractionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZoneInteraction.Contract.ZoneInteractionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZoneInteraction *ZoneInteractionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZoneInteraction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZoneInteraction *ZoneInteractionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZoneInteraction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZoneInteraction *ZoneInteractionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZoneInteraction.Contract.contract.Transact(opts, method, params...)
}
