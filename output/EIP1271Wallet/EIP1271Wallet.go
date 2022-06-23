// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EIP1271Wallet

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

// EIP1271WalletMetaData contains all meta data concerning the EIP1271Wallet contract.
var EIP1271WalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"contractERC20ApprovalInterface\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractNFTApprovalInterface\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"approveNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"digestApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"registerDigest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"showMessage\",\"type\":\"bool\"}],\"name\":\"revertWithMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"setValid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"showRevertMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161083838038061083883398101604081905261002f9161005d565b6001600160a01b031660805260008054600160ff19918216811790925560028054909116909117905561008d565b60006020828403121561006f57600080fd5b81516001600160a01b038116811461008657600080fd5b9392505050565b60805161077b6100bd6000396000818161015e01528181610302015281816103b70152610482015261077b6000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80638da5cb5b116100665780638da5cb5b1461015957806397fe801414610198578063a8e5e4aa146101b5578063bb5d40eb146101c8578063de77d5e0146101d557600080fd5b80631626ba7e146100a357806329e282e8146100d457806346a35674146101045780636c64edee1461012557806384e6d7e714610146575b600080fd5b6100b66100b136600461056b565b6101f8565b6040516001600160e01b031990911681526020015b60405180910390f35b6101026100e2366004610637565b600091825260016020526040909120805460ff1916911515919091179055565b005b610102610112366004610667565b6000805460ff1916911515919091179055565b610102610133366004610667565b6002805460ff1916911515919091179055565b6101026101543660046106a0565b6103ac565b6101807f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100cb565b6000546101a59060ff1681565b60405190151581526020016100cb565b6101026101c33660046106ce565b610477565b6002546101a59060ff1681565b6101a56101e336600461070f565b60016020526000908152604090205460ff1681565b60008281526001602052604081205460ff161561021d5750630b135d3f60e11b6103a6565b815160411461022b57600080fd5b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561026a57600080fd5b8060ff16601b1415801561028257508060ff16601c14155b1561028c57600080fd5b6040805160008082526020820180845289905260ff841692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa1580156102e0573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661030057600080fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b03161461037e5760005460ff161561009e5760405162461bcd60e51b815260206004820152600a6024820152692120a21029a4a3a722a960b11b60448201526064015b60405180910390fd5b60025460ff16610396576001600160e01b031961039f565b630b135d3f60e11b5b9450505050505b92915050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104115760405162461bcd60e51b815260206004820152600a60248201526927b7363c9037bbb732b960b11b6044820152606401610375565b60405163a22cb46560e01b81526001600160a01b0382811660048301526001602483015283169063a22cb46590604401600060405180830381600087803b15801561045b57600080fd5b505af115801561046f573d6000803e3d6000fd5b505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104dc5760405162461bcd60e51b815260206004820152600a60248201526927b7363c9037bbb732b960b11b6044820152606401610375565b60405163095ea7b360e01b81526001600160a01b0383811660048301526024820183905284169063095ea7b3906044016020604051808303816000875af115801561052b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061054f9190610728565b50505050565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561057e57600080fd5b82359150602083013567ffffffffffffffff8082111561059d57600080fd5b818501915085601f8301126105b157600080fd5b8135818111156105c3576105c3610555565b604051601f8201601f19908116603f011681019083821181831017156105eb576105eb610555565b8160405282815288602084870101111561060457600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b801515811461063457600080fd5b50565b6000806040838503121561064a57600080fd5b82359150602083013561065c81610626565b809150509250929050565b60006020828403121561067957600080fd5b813561068481610626565b9392505050565b6001600160a01b038116811461063457600080fd5b600080604083850312156106b357600080fd5b82356106be8161068b565b9150602083013561065c8161068b565b6000806000606084860312156106e357600080fd5b83356106ee8161068b565b925060208401356106fe8161068b565b929592945050506040919091013590565b60006020828403121561072157600080fd5b5035919050565b60006020828403121561073a57600080fd5b81516106848161062656fea2646970667358221220361e69914a46ffd7a45993a3e7457def6786117e580b070dac8274d88bca263964736f6c634300080d0033",
}

// EIP1271WalletABI is the input ABI used to generate the binding from.
// Deprecated: Use EIP1271WalletMetaData.ABI instead.
var EIP1271WalletABI = EIP1271WalletMetaData.ABI

// EIP1271WalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EIP1271WalletMetaData.Bin instead.
var EIP1271WalletBin = EIP1271WalletMetaData.Bin

// DeployEIP1271Wallet deploys a new Ethereum contract, binding an instance of EIP1271Wallet to it.
func DeployEIP1271Wallet(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *EIP1271Wallet, error) {
	parsed, err := EIP1271WalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EIP1271WalletBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EIP1271Wallet{EIP1271WalletCaller: EIP1271WalletCaller{contract: contract}, EIP1271WalletTransactor: EIP1271WalletTransactor{contract: contract}, EIP1271WalletFilterer: EIP1271WalletFilterer{contract: contract}}, nil
}

// EIP1271Wallet is an auto generated Go binding around an Ethereum contract.
type EIP1271Wallet struct {
	EIP1271WalletCaller     // Read-only binding to the contract
	EIP1271WalletTransactor // Write-only binding to the contract
	EIP1271WalletFilterer   // Log filterer for contract events
}

// EIP1271WalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type EIP1271WalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP1271WalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EIP1271WalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP1271WalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EIP1271WalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP1271WalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EIP1271WalletSession struct {
	Contract     *EIP1271Wallet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIP1271WalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EIP1271WalletCallerSession struct {
	Contract *EIP1271WalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EIP1271WalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EIP1271WalletTransactorSession struct {
	Contract     *EIP1271WalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EIP1271WalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type EIP1271WalletRaw struct {
	Contract *EIP1271Wallet // Generic contract binding to access the raw methods on
}

// EIP1271WalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EIP1271WalletCallerRaw struct {
	Contract *EIP1271WalletCaller // Generic read-only contract binding to access the raw methods on
}

// EIP1271WalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EIP1271WalletTransactorRaw struct {
	Contract *EIP1271WalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEIP1271Wallet creates a new instance of EIP1271Wallet, bound to a specific deployed contract.
func NewEIP1271Wallet(address common.Address, backend bind.ContractBackend) (*EIP1271Wallet, error) {
	contract, err := bindEIP1271Wallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EIP1271Wallet{EIP1271WalletCaller: EIP1271WalletCaller{contract: contract}, EIP1271WalletTransactor: EIP1271WalletTransactor{contract: contract}, EIP1271WalletFilterer: EIP1271WalletFilterer{contract: contract}}, nil
}

// NewEIP1271WalletCaller creates a new read-only instance of EIP1271Wallet, bound to a specific deployed contract.
func NewEIP1271WalletCaller(address common.Address, caller bind.ContractCaller) (*EIP1271WalletCaller, error) {
	contract, err := bindEIP1271Wallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EIP1271WalletCaller{contract: contract}, nil
}

// NewEIP1271WalletTransactor creates a new write-only instance of EIP1271Wallet, bound to a specific deployed contract.
func NewEIP1271WalletTransactor(address common.Address, transactor bind.ContractTransactor) (*EIP1271WalletTransactor, error) {
	contract, err := bindEIP1271Wallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EIP1271WalletTransactor{contract: contract}, nil
}

// NewEIP1271WalletFilterer creates a new log filterer instance of EIP1271Wallet, bound to a specific deployed contract.
func NewEIP1271WalletFilterer(address common.Address, filterer bind.ContractFilterer) (*EIP1271WalletFilterer, error) {
	contract, err := bindEIP1271Wallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EIP1271WalletFilterer{contract: contract}, nil
}

// bindEIP1271Wallet binds a generic wrapper to an already deployed contract.
func bindEIP1271Wallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EIP1271WalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP1271Wallet *EIP1271WalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP1271Wallet.Contract.EIP1271WalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP1271Wallet *EIP1271WalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP1271Wallet.Contract.EIP1271WalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP1271Wallet *EIP1271WalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP1271Wallet.Contract.EIP1271WalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP1271Wallet *EIP1271WalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP1271Wallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP1271Wallet *EIP1271WalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP1271Wallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP1271Wallet *EIP1271WalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP1271Wallet.Contract.contract.Transact(opts, method, params...)
}

// DigestApproved is a free data retrieval call binding the contract method 0xde77d5e0.
//
// Solidity: function digestApproved(bytes32 ) view returns(bool)
func (_EIP1271Wallet *EIP1271WalletCaller) DigestApproved(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _EIP1271Wallet.contract.Call(opts, &out, "digestApproved", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DigestApproved is a free data retrieval call binding the contract method 0xde77d5e0.
//
// Solidity: function digestApproved(bytes32 ) view returns(bool)
func (_EIP1271Wallet *EIP1271WalletSession) DigestApproved(arg0 [32]byte) (bool, error) {
	return _EIP1271Wallet.Contract.DigestApproved(&_EIP1271Wallet.CallOpts, arg0)
}

// DigestApproved is a free data retrieval call binding the contract method 0xde77d5e0.
//
// Solidity: function digestApproved(bytes32 ) view returns(bool)
func (_EIP1271Wallet *EIP1271WalletCallerSession) DigestApproved(arg0 [32]byte) (bool, error) {
	return _EIP1271Wallet.Contract.DigestApproved(&_EIP1271Wallet.CallOpts, arg0)
}

// IsValid is a free data retrieval call binding the contract method 0xbb5d40eb.
//
// Solidity: function isValid() view returns(bool)
func (_EIP1271Wallet *EIP1271WalletCaller) IsValid(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EIP1271Wallet.contract.Call(opts, &out, "isValid")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValid is a free data retrieval call binding the contract method 0xbb5d40eb.
//
// Solidity: function isValid() view returns(bool)
func (_EIP1271Wallet *EIP1271WalletSession) IsValid() (bool, error) {
	return _EIP1271Wallet.Contract.IsValid(&_EIP1271Wallet.CallOpts)
}

// IsValid is a free data retrieval call binding the contract method 0xbb5d40eb.
//
// Solidity: function isValid() view returns(bool)
func (_EIP1271Wallet *EIP1271WalletCallerSession) IsValid() (bool, error) {
	return _EIP1271Wallet.Contract.IsValid(&_EIP1271Wallet.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 digest, bytes signature) view returns(bytes4)
func (_EIP1271Wallet *EIP1271WalletCaller) IsValidSignature(opts *bind.CallOpts, digest [32]byte, signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _EIP1271Wallet.contract.Call(opts, &out, "isValidSignature", digest, signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 digest, bytes signature) view returns(bytes4)
func (_EIP1271Wallet *EIP1271WalletSession) IsValidSignature(digest [32]byte, signature []byte) ([4]byte, error) {
	return _EIP1271Wallet.Contract.IsValidSignature(&_EIP1271Wallet.CallOpts, digest, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 digest, bytes signature) view returns(bytes4)
func (_EIP1271Wallet *EIP1271WalletCallerSession) IsValidSignature(digest [32]byte, signature []byte) ([4]byte, error) {
	return _EIP1271Wallet.Contract.IsValidSignature(&_EIP1271Wallet.CallOpts, digest, signature)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EIP1271Wallet *EIP1271WalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EIP1271Wallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EIP1271Wallet *EIP1271WalletSession) Owner() (common.Address, error) {
	return _EIP1271Wallet.Contract.Owner(&_EIP1271Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EIP1271Wallet *EIP1271WalletCallerSession) Owner() (common.Address, error) {
	return _EIP1271Wallet.Contract.Owner(&_EIP1271Wallet.CallOpts)
}

// ShowRevertMessage is a free data retrieval call binding the contract method 0x97fe8014.
//
// Solidity: function showRevertMessage() view returns(bool)
func (_EIP1271Wallet *EIP1271WalletCaller) ShowRevertMessage(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EIP1271Wallet.contract.Call(opts, &out, "showRevertMessage")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShowRevertMessage is a free data retrieval call binding the contract method 0x97fe8014.
//
// Solidity: function showRevertMessage() view returns(bool)
func (_EIP1271Wallet *EIP1271WalletSession) ShowRevertMessage() (bool, error) {
	return _EIP1271Wallet.Contract.ShowRevertMessage(&_EIP1271Wallet.CallOpts)
}

// ShowRevertMessage is a free data retrieval call binding the contract method 0x97fe8014.
//
// Solidity: function showRevertMessage() view returns(bool)
func (_EIP1271Wallet *EIP1271WalletCallerSession) ShowRevertMessage() (bool, error) {
	return _EIP1271Wallet.Contract.ShowRevertMessage(&_EIP1271Wallet.CallOpts)
}

// ApproveERC20 is a paid mutator transaction binding the contract method 0xa8e5e4aa.
//
// Solidity: function approveERC20(address token, address operator, uint256 amount) returns()
func (_EIP1271Wallet *EIP1271WalletTransactor) ApproveERC20(opts *bind.TransactOpts, token common.Address, operator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP1271Wallet.contract.Transact(opts, "approveERC20", token, operator, amount)
}

// ApproveERC20 is a paid mutator transaction binding the contract method 0xa8e5e4aa.
//
// Solidity: function approveERC20(address token, address operator, uint256 amount) returns()
func (_EIP1271Wallet *EIP1271WalletSession) ApproveERC20(token common.Address, operator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP1271Wallet.Contract.ApproveERC20(&_EIP1271Wallet.TransactOpts, token, operator, amount)
}

// ApproveERC20 is a paid mutator transaction binding the contract method 0xa8e5e4aa.
//
// Solidity: function approveERC20(address token, address operator, uint256 amount) returns()
func (_EIP1271Wallet *EIP1271WalletTransactorSession) ApproveERC20(token common.Address, operator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP1271Wallet.Contract.ApproveERC20(&_EIP1271Wallet.TransactOpts, token, operator, amount)
}

// ApproveNFT is a paid mutator transaction binding the contract method 0x84e6d7e7.
//
// Solidity: function approveNFT(address token, address operator) returns()
func (_EIP1271Wallet *EIP1271WalletTransactor) ApproveNFT(opts *bind.TransactOpts, token common.Address, operator common.Address) (*types.Transaction, error) {
	return _EIP1271Wallet.contract.Transact(opts, "approveNFT", token, operator)
}

// ApproveNFT is a paid mutator transaction binding the contract method 0x84e6d7e7.
//
// Solidity: function approveNFT(address token, address operator) returns()
func (_EIP1271Wallet *EIP1271WalletSession) ApproveNFT(token common.Address, operator common.Address) (*types.Transaction, error) {
	return _EIP1271Wallet.Contract.ApproveNFT(&_EIP1271Wallet.TransactOpts, token, operator)
}

// ApproveNFT is a paid mutator transaction binding the contract method 0x84e6d7e7.
//
// Solidity: function approveNFT(address token, address operator) returns()
func (_EIP1271Wallet *EIP1271WalletTransactorSession) ApproveNFT(token common.Address, operator common.Address) (*types.Transaction, error) {
	return _EIP1271Wallet.Contract.ApproveNFT(&_EIP1271Wallet.TransactOpts, token, operator)
}

// RegisterDigest is a paid mutator transaction binding the contract method 0x29e282e8.
//
// Solidity: function registerDigest(bytes32 digest, bool approved) returns()
func (_EIP1271Wallet *EIP1271WalletTransactor) RegisterDigest(opts *bind.TransactOpts, digest [32]byte, approved bool) (*types.Transaction, error) {
	return _EIP1271Wallet.contract.Transact(opts, "registerDigest", digest, approved)
}

// RegisterDigest is a paid mutator transaction binding the contract method 0x29e282e8.
//
// Solidity: function registerDigest(bytes32 digest, bool approved) returns()
func (_EIP1271Wallet *EIP1271WalletSession) RegisterDigest(digest [32]byte, approved bool) (*types.Transaction, error) {
	return _EIP1271Wallet.Contract.RegisterDigest(&_EIP1271Wallet.TransactOpts, digest, approved)
}

// RegisterDigest is a paid mutator transaction binding the contract method 0x29e282e8.
//
// Solidity: function registerDigest(bytes32 digest, bool approved) returns()
func (_EIP1271Wallet *EIP1271WalletTransactorSession) RegisterDigest(digest [32]byte, approved bool) (*types.Transaction, error) {
	return _EIP1271Wallet.Contract.RegisterDigest(&_EIP1271Wallet.TransactOpts, digest, approved)
}

// RevertWithMessage is a paid mutator transaction binding the contract method 0x46a35674.
//
// Solidity: function revertWithMessage(bool showMessage) returns()
func (_EIP1271Wallet *EIP1271WalletTransactor) RevertWithMessage(opts *bind.TransactOpts, showMessage bool) (*types.Transaction, error) {
	return _EIP1271Wallet.contract.Transact(opts, "revertWithMessage", showMessage)
}

// RevertWithMessage is a paid mutator transaction binding the contract method 0x46a35674.
//
// Solidity: function revertWithMessage(bool showMessage) returns()
func (_EIP1271Wallet *EIP1271WalletSession) RevertWithMessage(showMessage bool) (*types.Transaction, error) {
	return _EIP1271Wallet.Contract.RevertWithMessage(&_EIP1271Wallet.TransactOpts, showMessage)
}

// RevertWithMessage is a paid mutator transaction binding the contract method 0x46a35674.
//
// Solidity: function revertWithMessage(bool showMessage) returns()
func (_EIP1271Wallet *EIP1271WalletTransactorSession) RevertWithMessage(showMessage bool) (*types.Transaction, error) {
	return _EIP1271Wallet.Contract.RevertWithMessage(&_EIP1271Wallet.TransactOpts, showMessage)
}

// SetValid is a paid mutator transaction binding the contract method 0x6c64edee.
//
// Solidity: function setValid(bool valid) returns()
func (_EIP1271Wallet *EIP1271WalletTransactor) SetValid(opts *bind.TransactOpts, valid bool) (*types.Transaction, error) {
	return _EIP1271Wallet.contract.Transact(opts, "setValid", valid)
}

// SetValid is a paid mutator transaction binding the contract method 0x6c64edee.
//
// Solidity: function setValid(bool valid) returns()
func (_EIP1271Wallet *EIP1271WalletSession) SetValid(valid bool) (*types.Transaction, error) {
	return _EIP1271Wallet.Contract.SetValid(&_EIP1271Wallet.TransactOpts, valid)
}

// SetValid is a paid mutator transaction binding the contract method 0x6c64edee.
//
// Solidity: function setValid(bool valid) returns()
func (_EIP1271Wallet *EIP1271WalletTransactorSession) SetValid(valid bool) (*types.Transaction, error) {
	return _EIP1271Wallet.Contract.SetValid(&_EIP1271Wallet.TransactOpts, valid)
}
