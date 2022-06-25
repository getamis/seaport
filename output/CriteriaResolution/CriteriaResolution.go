// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CriteriaResolution

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

// CriteriaResolutionMetaData contains all meta data concerning the CriteriaResolution contract.
var CriteriaResolutionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ConsiderationCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CriteriaNotEnabledForItem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnresolvedConsiderationCriteria\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnresolvedOfferCriteria\",\"type\":\"error\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea264697066735822122064b55239343c014b01b29bd345188ab53e72c59994a937271fbbcdc25fd0e91d64736f6c634300080d0033",
}

// CriteriaResolutionABI is the input ABI used to generate the binding from.
// Deprecated: Use CriteriaResolutionMetaData.ABI instead.
var CriteriaResolutionABI = CriteriaResolutionMetaData.ABI

// CriteriaResolutionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CriteriaResolutionMetaData.Bin instead.
var CriteriaResolutionBin = CriteriaResolutionMetaData.Bin

// DeployCriteriaResolution deploys a new Ethereum contract, binding an instance of CriteriaResolution to it.
func DeployCriteriaResolution(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CriteriaResolution, error) {
	parsed, err := CriteriaResolutionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CriteriaResolutionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CriteriaResolution{CriteriaResolutionCaller: CriteriaResolutionCaller{contract: contract}, CriteriaResolutionTransactor: CriteriaResolutionTransactor{contract: contract}, CriteriaResolutionFilterer: CriteriaResolutionFilterer{contract: contract}}, nil
}

// CriteriaResolution is an auto generated Go binding around an Ethereum contract.
type CriteriaResolution struct {
	CriteriaResolutionCaller     // Read-only binding to the contract
	CriteriaResolutionTransactor // Write-only binding to the contract
	CriteriaResolutionFilterer   // Log filterer for contract events
}

// CriteriaResolutionCaller is an auto generated read-only Go binding around an Ethereum contract.
type CriteriaResolutionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CriteriaResolutionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CriteriaResolutionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CriteriaResolutionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CriteriaResolutionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CriteriaResolutionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CriteriaResolutionSession struct {
	Contract     *CriteriaResolution // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CriteriaResolutionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CriteriaResolutionCallerSession struct {
	Contract *CriteriaResolutionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CriteriaResolutionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CriteriaResolutionTransactorSession struct {
	Contract     *CriteriaResolutionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CriteriaResolutionRaw is an auto generated low-level Go binding around an Ethereum contract.
type CriteriaResolutionRaw struct {
	Contract *CriteriaResolution // Generic contract binding to access the raw methods on
}

// CriteriaResolutionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CriteriaResolutionCallerRaw struct {
	Contract *CriteriaResolutionCaller // Generic read-only contract binding to access the raw methods on
}

// CriteriaResolutionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CriteriaResolutionTransactorRaw struct {
	Contract *CriteriaResolutionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCriteriaResolution creates a new instance of CriteriaResolution, bound to a specific deployed contract.
func NewCriteriaResolution(address common.Address, backend bind.ContractBackend) (*CriteriaResolution, error) {
	contract, err := bindCriteriaResolution(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CriteriaResolution{CriteriaResolutionCaller: CriteriaResolutionCaller{contract: contract}, CriteriaResolutionTransactor: CriteriaResolutionTransactor{contract: contract}, CriteriaResolutionFilterer: CriteriaResolutionFilterer{contract: contract}}, nil
}

// NewCriteriaResolutionCaller creates a new read-only instance of CriteriaResolution, bound to a specific deployed contract.
func NewCriteriaResolutionCaller(address common.Address, caller bind.ContractCaller) (*CriteriaResolutionCaller, error) {
	contract, err := bindCriteriaResolution(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CriteriaResolutionCaller{contract: contract}, nil
}

// NewCriteriaResolutionTransactor creates a new write-only instance of CriteriaResolution, bound to a specific deployed contract.
func NewCriteriaResolutionTransactor(address common.Address, transactor bind.ContractTransactor) (*CriteriaResolutionTransactor, error) {
	contract, err := bindCriteriaResolution(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CriteriaResolutionTransactor{contract: contract}, nil
}

// NewCriteriaResolutionFilterer creates a new log filterer instance of CriteriaResolution, bound to a specific deployed contract.
func NewCriteriaResolutionFilterer(address common.Address, filterer bind.ContractFilterer) (*CriteriaResolutionFilterer, error) {
	contract, err := bindCriteriaResolution(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CriteriaResolutionFilterer{contract: contract}, nil
}

// bindCriteriaResolution binds a generic wrapper to an already deployed contract.
func bindCriteriaResolution(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CriteriaResolutionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CriteriaResolution *CriteriaResolutionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CriteriaResolution.Contract.CriteriaResolutionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CriteriaResolution *CriteriaResolutionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CriteriaResolution.Contract.CriteriaResolutionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CriteriaResolution *CriteriaResolutionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CriteriaResolution.Contract.CriteriaResolutionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CriteriaResolution *CriteriaResolutionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CriteriaResolution.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CriteriaResolution *CriteriaResolutionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CriteriaResolution.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CriteriaResolution *CriteriaResolutionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CriteriaResolution.Contract.contract.Transact(opts, method, params...)
}
