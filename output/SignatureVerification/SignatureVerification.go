// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SignatureVerification

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

// SignatureVerificationMetaData contains all meta data concerning the SignatureVerification contract.
var SignatureVerificationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BadContractSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"BadSignatureV\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220b773836fb584543f5973c0154eea11c8dd98d379894ef68dbc6701524ec7d8dc64736f6c634300080d0033",
}

// SignatureVerificationABI is the input ABI used to generate the binding from.
// Deprecated: Use SignatureVerificationMetaData.ABI instead.
var SignatureVerificationABI = SignatureVerificationMetaData.ABI

// SignatureVerificationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SignatureVerificationMetaData.Bin instead.
var SignatureVerificationBin = SignatureVerificationMetaData.Bin

// DeploySignatureVerification deploys a new Ethereum contract, binding an instance of SignatureVerification to it.
func DeploySignatureVerification(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SignatureVerification, error) {
	parsed, err := SignatureVerificationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SignatureVerificationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SignatureVerification{SignatureVerificationCaller: SignatureVerificationCaller{contract: contract}, SignatureVerificationTransactor: SignatureVerificationTransactor{contract: contract}, SignatureVerificationFilterer: SignatureVerificationFilterer{contract: contract}}, nil
}

// SignatureVerification is an auto generated Go binding around an Ethereum contract.
type SignatureVerification struct {
	SignatureVerificationCaller     // Read-only binding to the contract
	SignatureVerificationTransactor // Write-only binding to the contract
	SignatureVerificationFilterer   // Log filterer for contract events
}

// SignatureVerificationCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignatureVerificationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureVerificationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignatureVerificationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureVerificationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignatureVerificationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureVerificationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignatureVerificationSession struct {
	Contract     *SignatureVerification // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SignatureVerificationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignatureVerificationCallerSession struct {
	Contract *SignatureVerificationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// SignatureVerificationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignatureVerificationTransactorSession struct {
	Contract     *SignatureVerificationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// SignatureVerificationRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignatureVerificationRaw struct {
	Contract *SignatureVerification // Generic contract binding to access the raw methods on
}

// SignatureVerificationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignatureVerificationCallerRaw struct {
	Contract *SignatureVerificationCaller // Generic read-only contract binding to access the raw methods on
}

// SignatureVerificationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignatureVerificationTransactorRaw struct {
	Contract *SignatureVerificationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignatureVerification creates a new instance of SignatureVerification, bound to a specific deployed contract.
func NewSignatureVerification(address common.Address, backend bind.ContractBackend) (*SignatureVerification, error) {
	contract, err := bindSignatureVerification(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignatureVerification{SignatureVerificationCaller: SignatureVerificationCaller{contract: contract}, SignatureVerificationTransactor: SignatureVerificationTransactor{contract: contract}, SignatureVerificationFilterer: SignatureVerificationFilterer{contract: contract}}, nil
}

// NewSignatureVerificationCaller creates a new read-only instance of SignatureVerification, bound to a specific deployed contract.
func NewSignatureVerificationCaller(address common.Address, caller bind.ContractCaller) (*SignatureVerificationCaller, error) {
	contract, err := bindSignatureVerification(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureVerificationCaller{contract: contract}, nil
}

// NewSignatureVerificationTransactor creates a new write-only instance of SignatureVerification, bound to a specific deployed contract.
func NewSignatureVerificationTransactor(address common.Address, transactor bind.ContractTransactor) (*SignatureVerificationTransactor, error) {
	contract, err := bindSignatureVerification(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureVerificationTransactor{contract: contract}, nil
}

// NewSignatureVerificationFilterer creates a new log filterer instance of SignatureVerification, bound to a specific deployed contract.
func NewSignatureVerificationFilterer(address common.Address, filterer bind.ContractFilterer) (*SignatureVerificationFilterer, error) {
	contract, err := bindSignatureVerification(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignatureVerificationFilterer{contract: contract}, nil
}

// bindSignatureVerification binds a generic wrapper to an already deployed contract.
func bindSignatureVerification(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SignatureVerificationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureVerification *SignatureVerificationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignatureVerification.Contract.SignatureVerificationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureVerification *SignatureVerificationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureVerification.Contract.SignatureVerificationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureVerification *SignatureVerificationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureVerification.Contract.SignatureVerificationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureVerification *SignatureVerificationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignatureVerification.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureVerification *SignatureVerificationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureVerification.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureVerification *SignatureVerificationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureVerification.Contract.contract.Transact(opts, method, params...)
}
