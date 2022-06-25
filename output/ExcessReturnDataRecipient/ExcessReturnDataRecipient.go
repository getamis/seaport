// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ExcessReturnDataRecipient

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

// ExcessReturnDataRecipientMetaData contains all meta data concerning the ExcessReturnDataRecipient contract.
var ExcessReturnDataRecipientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magic\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"setRevertDataSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610529806100206000396000f3fe60806040526004361061002d5760003560e01c8063e998da54146100bd578063f23a6e61146100dd57600080fd5b366100b857600054156100b65760006100455a61011a565b6100509060366103a8565b61005c906104b06103c7565b905060006002610c00610092845a61007491906103df565b610080906108006103a8565b61008d90628fe8006103c7565b6102f1565b61009c91906103df565b6100a691906103f6565b9050602081028060008037806000fd5b005b600080fd5b3480156100c957600080fd5b506100b66100d8366004610418565b600055565b3480156100e957600080fd5b506100fd6100f836600461044d565b61033d565b6040516001600160e01b0319909116815260200160405180910390f35b604080517ff8f9cbfae6cc78fbefe7cdc3a1793dfcf4f0e8bbd8cec470b6a28a7a5a3e1efd81527ff5ecf1b3e9debc68e1d9cfabc5997135bfb7a7a3938b7b606b5b4b3f2f1f0ffe60208201527ff6e4ed9ff2d6b458eadcdf97bd91692de2d4da8fd2d0ac50c6ae9a8272523616818301527fc8c0b887b0a8a4489c948c7f847c6125746c645c544c444038302820181008ff60608201527ff7cae577eec2a03cf3bad76fb589591debb2dd67e0aa9834bea6925f6a4a2e0e60808201527fe39ed557db96902cd38ed14fad815115c786af479b7e8324736353433727170760a08201527fc976c13bb96e881cb166a933a55e490d9d56952b8d4e801485467d236242260660c08201527f753a6d1b65325d0c552a4d1345224105391a310b29122104190a11030902010060e0820152610100818101909252600160f81b7e818283848586878898a8b8c8d8e8f929395969799a9b9d9e9faaeb6bedeeff600160801b68010000000000000000640100000000620100006010600460026000198c0190810417908104179081041788810417908104179081041790810417908104176001010281900460ff03909101516633413c26506520662386f26fc10000600160ff1b90951190930291900401919091020490565b6000600382116001811461030a57801561033257610337565b829150600260018401045b8281101561032c5791506002828404830104610315565b50610337565b600191505b50919050565b60005463f23a6e6160e01b90156103885760006103595a61011a565b6103649060366103a8565b610370906104b06103c7565b905060006004610c00610092845a61007491906103df565b9695505050505050565b634e487b7160e01b600052601160045260246000fd5b60008160001904831182151516156103c2576103c2610392565b500290565b600082198211156103da576103da610392565b500190565b6000828210156103f1576103f1610392565b500390565b60008261041357634e487b7160e01b600052601260045260246000fd5b500490565b60006020828403121561042a57600080fd5b5035919050565b80356001600160a01b038116811461044857600080fd5b919050565b60008060008060008060a0878903121561046657600080fd5b61046f87610431565b955061047d60208801610431565b94506040870135935060608701359250608087013567ffffffffffffffff808211156104a857600080fd5b818901915089601f8301126104bc57600080fd5b8135818111156104cb57600080fd5b8a60208285010111156104dd57600080fd5b602083019450809350505050929550929550929556fea26469706673582212204c09612f22117e01e17c7cd2034a930756bd5481cc8b93ac2f6e5d66412b323864736f6c634300080d0033",
}

// ExcessReturnDataRecipientABI is the input ABI used to generate the binding from.
// Deprecated: Use ExcessReturnDataRecipientMetaData.ABI instead.
var ExcessReturnDataRecipientABI = ExcessReturnDataRecipientMetaData.ABI

// ExcessReturnDataRecipientBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExcessReturnDataRecipientMetaData.Bin instead.
var ExcessReturnDataRecipientBin = ExcessReturnDataRecipientMetaData.Bin

// DeployExcessReturnDataRecipient deploys a new Ethereum contract, binding an instance of ExcessReturnDataRecipient to it.
func DeployExcessReturnDataRecipient(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExcessReturnDataRecipient, error) {
	parsed, err := ExcessReturnDataRecipientMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExcessReturnDataRecipientBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExcessReturnDataRecipient{ExcessReturnDataRecipientCaller: ExcessReturnDataRecipientCaller{contract: contract}, ExcessReturnDataRecipientTransactor: ExcessReturnDataRecipientTransactor{contract: contract}, ExcessReturnDataRecipientFilterer: ExcessReturnDataRecipientFilterer{contract: contract}}, nil
}

// ExcessReturnDataRecipient is an auto generated Go binding around an Ethereum contract.
type ExcessReturnDataRecipient struct {
	ExcessReturnDataRecipientCaller     // Read-only binding to the contract
	ExcessReturnDataRecipientTransactor // Write-only binding to the contract
	ExcessReturnDataRecipientFilterer   // Log filterer for contract events
}

// ExcessReturnDataRecipientCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExcessReturnDataRecipientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExcessReturnDataRecipientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExcessReturnDataRecipientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExcessReturnDataRecipientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExcessReturnDataRecipientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExcessReturnDataRecipientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExcessReturnDataRecipientSession struct {
	Contract     *ExcessReturnDataRecipient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ExcessReturnDataRecipientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExcessReturnDataRecipientCallerSession struct {
	Contract *ExcessReturnDataRecipientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ExcessReturnDataRecipientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExcessReturnDataRecipientTransactorSession struct {
	Contract     *ExcessReturnDataRecipientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ExcessReturnDataRecipientRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExcessReturnDataRecipientRaw struct {
	Contract *ExcessReturnDataRecipient // Generic contract binding to access the raw methods on
}

// ExcessReturnDataRecipientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExcessReturnDataRecipientCallerRaw struct {
	Contract *ExcessReturnDataRecipientCaller // Generic read-only contract binding to access the raw methods on
}

// ExcessReturnDataRecipientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExcessReturnDataRecipientTransactorRaw struct {
	Contract *ExcessReturnDataRecipientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExcessReturnDataRecipient creates a new instance of ExcessReturnDataRecipient, bound to a specific deployed contract.
func NewExcessReturnDataRecipient(address common.Address, backend bind.ContractBackend) (*ExcessReturnDataRecipient, error) {
	contract, err := bindExcessReturnDataRecipient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExcessReturnDataRecipient{ExcessReturnDataRecipientCaller: ExcessReturnDataRecipientCaller{contract: contract}, ExcessReturnDataRecipientTransactor: ExcessReturnDataRecipientTransactor{contract: contract}, ExcessReturnDataRecipientFilterer: ExcessReturnDataRecipientFilterer{contract: contract}}, nil
}

// NewExcessReturnDataRecipientCaller creates a new read-only instance of ExcessReturnDataRecipient, bound to a specific deployed contract.
func NewExcessReturnDataRecipientCaller(address common.Address, caller bind.ContractCaller) (*ExcessReturnDataRecipientCaller, error) {
	contract, err := bindExcessReturnDataRecipient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExcessReturnDataRecipientCaller{contract: contract}, nil
}

// NewExcessReturnDataRecipientTransactor creates a new write-only instance of ExcessReturnDataRecipient, bound to a specific deployed contract.
func NewExcessReturnDataRecipientTransactor(address common.Address, transactor bind.ContractTransactor) (*ExcessReturnDataRecipientTransactor, error) {
	contract, err := bindExcessReturnDataRecipient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExcessReturnDataRecipientTransactor{contract: contract}, nil
}

// NewExcessReturnDataRecipientFilterer creates a new log filterer instance of ExcessReturnDataRecipient, bound to a specific deployed contract.
func NewExcessReturnDataRecipientFilterer(address common.Address, filterer bind.ContractFilterer) (*ExcessReturnDataRecipientFilterer, error) {
	contract, err := bindExcessReturnDataRecipient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExcessReturnDataRecipientFilterer{contract: contract}, nil
}

// bindExcessReturnDataRecipient binds a generic wrapper to an already deployed contract.
func bindExcessReturnDataRecipient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExcessReturnDataRecipientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExcessReturnDataRecipient *ExcessReturnDataRecipientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExcessReturnDataRecipient.Contract.ExcessReturnDataRecipientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExcessReturnDataRecipient *ExcessReturnDataRecipientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExcessReturnDataRecipient.Contract.ExcessReturnDataRecipientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExcessReturnDataRecipient *ExcessReturnDataRecipientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExcessReturnDataRecipient.Contract.ExcessReturnDataRecipientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExcessReturnDataRecipient *ExcessReturnDataRecipientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExcessReturnDataRecipient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExcessReturnDataRecipient *ExcessReturnDataRecipientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExcessReturnDataRecipient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExcessReturnDataRecipient *ExcessReturnDataRecipientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExcessReturnDataRecipient.Contract.contract.Transact(opts, method, params...)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4 magic)
func (_ExcessReturnDataRecipient *ExcessReturnDataRecipientCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _ExcessReturnDataRecipient.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4 magic)
func (_ExcessReturnDataRecipient *ExcessReturnDataRecipientSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _ExcessReturnDataRecipient.Contract.OnERC1155Received(&_ExcessReturnDataRecipient.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4 magic)
func (_ExcessReturnDataRecipient *ExcessReturnDataRecipientCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _ExcessReturnDataRecipient.Contract.OnERC1155Received(&_ExcessReturnDataRecipient.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SetRevertDataSize is a paid mutator transaction binding the contract method 0xe998da54.
//
// Solidity: function setRevertDataSize(uint256 size) returns()
func (_ExcessReturnDataRecipient *ExcessReturnDataRecipientTransactor) SetRevertDataSize(opts *bind.TransactOpts, size *big.Int) (*types.Transaction, error) {
	return _ExcessReturnDataRecipient.contract.Transact(opts, "setRevertDataSize", size)
}

// SetRevertDataSize is a paid mutator transaction binding the contract method 0xe998da54.
//
// Solidity: function setRevertDataSize(uint256 size) returns()
func (_ExcessReturnDataRecipient *ExcessReturnDataRecipientSession) SetRevertDataSize(size *big.Int) (*types.Transaction, error) {
	return _ExcessReturnDataRecipient.Contract.SetRevertDataSize(&_ExcessReturnDataRecipient.TransactOpts, size)
}

// SetRevertDataSize is a paid mutator transaction binding the contract method 0xe998da54.
//
// Solidity: function setRevertDataSize(uint256 size) returns()
func (_ExcessReturnDataRecipient *ExcessReturnDataRecipientTransactorSession) SetRevertDataSize(size *big.Int) (*types.Transaction, error) {
	return _ExcessReturnDataRecipient.Contract.SetRevertDataSize(&_ExcessReturnDataRecipient.TransactOpts, size)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ExcessReturnDataRecipient *ExcessReturnDataRecipientTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExcessReturnDataRecipient.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ExcessReturnDataRecipient *ExcessReturnDataRecipientSession) Receive() (*types.Transaction, error) {
	return _ExcessReturnDataRecipient.Contract.Receive(&_ExcessReturnDataRecipient.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ExcessReturnDataRecipient *ExcessReturnDataRecipientTransactorSession) Receive() (*types.Transaction, error) {
	return _ExcessReturnDataRecipient.Contract.Receive(&_ExcessReturnDataRecipient.TransactOpts)
}
