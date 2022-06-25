// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC1155BatchRecipient

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

// ERC1155BatchRecipientMetaData contains all meta data concerning the ERC1155BatchRecipient contract.
var ERC1155BatchRecipientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"UnexpectedBatchData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610273806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063bc197c8114610030575b600080fd5b61004361003e366004610116565b610060565b6040516001600160e01b0319909116815260200160405180910390f35b6000815160001461008457604051631de805c760e31b815260040160405180910390fd5b5063bc197c8160e01b979650505050505050565b80356001600160a01b03811681146100af57600080fd5b919050565b60008083601f8401126100c657600080fd5b50813567ffffffffffffffff8111156100de57600080fd5b6020830191508360208260051b85010111156100f957600080fd5b9250929050565b634e487b7160e01b600052604160045260246000fd5b600080600080600080600060a0888a03121561013157600080fd5b61013a88610098565b965061014860208901610098565b9550604088013567ffffffffffffffff8082111561016557600080fd5b6101718b838c016100b4565b909750955060608a013591508082111561018a57600080fd5b6101968b838c016100b4565b909550935060808a01359150808211156101af57600080fd5b818a0191508a601f8301126101c357600080fd5b8135818111156101d5576101d5610100565b604051601f8201601f19908116603f011681019083821181831017156101fd576101fd610100565b816040528281528d602084870101111561021657600080fd5b8260208601602083013760006020848301015280955050505050509295989194975092955056fea2646970667358221220fba32fbc4ac9a7cd68a6b52fac4216072500e408a9f5f5d06ee7eccee954e80b64736f6c634300080d0033",
}

// ERC1155BatchRecipientABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1155BatchRecipientMetaData.ABI instead.
var ERC1155BatchRecipientABI = ERC1155BatchRecipientMetaData.ABI

// ERC1155BatchRecipientBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC1155BatchRecipientMetaData.Bin instead.
var ERC1155BatchRecipientBin = ERC1155BatchRecipientMetaData.Bin

// DeployERC1155BatchRecipient deploys a new Ethereum contract, binding an instance of ERC1155BatchRecipient to it.
func DeployERC1155BatchRecipient(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC1155BatchRecipient, error) {
	parsed, err := ERC1155BatchRecipientMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC1155BatchRecipientBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC1155BatchRecipient{ERC1155BatchRecipientCaller: ERC1155BatchRecipientCaller{contract: contract}, ERC1155BatchRecipientTransactor: ERC1155BatchRecipientTransactor{contract: contract}, ERC1155BatchRecipientFilterer: ERC1155BatchRecipientFilterer{contract: contract}}, nil
}

// ERC1155BatchRecipient is an auto generated Go binding around an Ethereum contract.
type ERC1155BatchRecipient struct {
	ERC1155BatchRecipientCaller     // Read-only binding to the contract
	ERC1155BatchRecipientTransactor // Write-only binding to the contract
	ERC1155BatchRecipientFilterer   // Log filterer for contract events
}

// ERC1155BatchRecipientCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1155BatchRecipientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155BatchRecipientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1155BatchRecipientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155BatchRecipientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1155BatchRecipientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155BatchRecipientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1155BatchRecipientSession struct {
	Contract     *ERC1155BatchRecipient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC1155BatchRecipientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1155BatchRecipientCallerSession struct {
	Contract *ERC1155BatchRecipientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ERC1155BatchRecipientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1155BatchRecipientTransactorSession struct {
	Contract     *ERC1155BatchRecipientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ERC1155BatchRecipientRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1155BatchRecipientRaw struct {
	Contract *ERC1155BatchRecipient // Generic contract binding to access the raw methods on
}

// ERC1155BatchRecipientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1155BatchRecipientCallerRaw struct {
	Contract *ERC1155BatchRecipientCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1155BatchRecipientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1155BatchRecipientTransactorRaw struct {
	Contract *ERC1155BatchRecipientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155BatchRecipient creates a new instance of ERC1155BatchRecipient, bound to a specific deployed contract.
func NewERC1155BatchRecipient(address common.Address, backend bind.ContractBackend) (*ERC1155BatchRecipient, error) {
	contract, err := bindERC1155BatchRecipient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155BatchRecipient{ERC1155BatchRecipientCaller: ERC1155BatchRecipientCaller{contract: contract}, ERC1155BatchRecipientTransactor: ERC1155BatchRecipientTransactor{contract: contract}, ERC1155BatchRecipientFilterer: ERC1155BatchRecipientFilterer{contract: contract}}, nil
}

// NewERC1155BatchRecipientCaller creates a new read-only instance of ERC1155BatchRecipient, bound to a specific deployed contract.
func NewERC1155BatchRecipientCaller(address common.Address, caller bind.ContractCaller) (*ERC1155BatchRecipientCaller, error) {
	contract, err := bindERC1155BatchRecipient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155BatchRecipientCaller{contract: contract}, nil
}

// NewERC1155BatchRecipientTransactor creates a new write-only instance of ERC1155BatchRecipient, bound to a specific deployed contract.
func NewERC1155BatchRecipientTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1155BatchRecipientTransactor, error) {
	contract, err := bindERC1155BatchRecipient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155BatchRecipientTransactor{contract: contract}, nil
}

// NewERC1155BatchRecipientFilterer creates a new log filterer instance of ERC1155BatchRecipient, bound to a specific deployed contract.
func NewERC1155BatchRecipientFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1155BatchRecipientFilterer, error) {
	contract, err := bindERC1155BatchRecipient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155BatchRecipientFilterer{contract: contract}, nil
}

// bindERC1155BatchRecipient binds a generic wrapper to an already deployed contract.
func bindERC1155BatchRecipient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1155BatchRecipientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155BatchRecipient *ERC1155BatchRecipientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155BatchRecipient.Contract.ERC1155BatchRecipientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155BatchRecipient *ERC1155BatchRecipientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155BatchRecipient.Contract.ERC1155BatchRecipientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155BatchRecipient *ERC1155BatchRecipientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155BatchRecipient.Contract.ERC1155BatchRecipientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155BatchRecipient *ERC1155BatchRecipientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155BatchRecipient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155BatchRecipient *ERC1155BatchRecipientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155BatchRecipient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155BatchRecipient *ERC1155BatchRecipientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155BatchRecipient.Contract.contract.Transact(opts, method, params...)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes data) pure returns(bytes4)
func (_ERC1155BatchRecipient *ERC1155BatchRecipientCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, data []byte) ([4]byte, error) {
	var out []interface{}
	err := _ERC1155BatchRecipient.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes data) pure returns(bytes4)
func (_ERC1155BatchRecipient *ERC1155BatchRecipientSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, data []byte) ([4]byte, error) {
	return _ERC1155BatchRecipient.Contract.OnERC1155BatchReceived(&_ERC1155BatchRecipient.CallOpts, arg0, arg1, arg2, arg3, data)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes data) pure returns(bytes4)
func (_ERC1155BatchRecipient *ERC1155BatchRecipientCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, data []byte) ([4]byte, error) {
	return _ERC1155BatchRecipient.Contract.OnERC1155BatchReceived(&_ERC1155BatchRecipient.CallOpts, arg0, arg1, arg2, arg3, data)
}
