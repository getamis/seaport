// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Reenterer

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

// ReentererMetaData contains all meta data concerning the Reenterer contract.
var ReentererMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"Reentered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"callData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"msgValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetToUse\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValueToUse\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callDataToUse\",\"type\":\"bytes\"}],\"name\":\"prepare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"target\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b506104f7806100206000396000f3fe6080604052600436106100435760003560e01c80634e417a98146100ff5780635157589b1461012a578063d4b839921461014c578063ddf363d71461018457600080fd5b366100fa576000805460015460405183926001600160a01b0316919061006b9060029061033c565b60006040518083038185875af1925050503d80600081146100a8576040519150601f19603f3d011682016040523d82523d6000602084013e6100ad565b606091505b5091509150816100c1573d6000803e3d6000fd5b7f52b1b52a16f834754c20ed40ad488c0fc799dea307f9bb9cde6413f0e356703e816040516100f091906103d7565b60405180910390a1005b600080fd5b34801561010b57600080fd5b506101146101a8565b60405161012191906103d7565b60405180910390f35b34801561013657600080fd5b5061014a61014536600461042c565b610236565b005b34801561015857600080fd5b5060005461016c906001600160a01b031681565b6040516001600160a01b039091168152602001610121565b34801561019057600080fd5b5061019a60015481565b604051908152602001610121565b600280546101b590610302565b80601f01602080910402602001604051908101604052809291908181526020018280546101e190610302565b801561022e5780601f106102035761010080835404028352916020019161022e565b820191906000526020600020905b81548152906001019060200180831161021157829003601f168201915b505050505081565b600080546001600160a01b0319166001600160a01b038616179055600183905561026260028383610269565b5050505050565b82805461027590610302565b90600052602060002090601f01602090048101928261029757600085556102dd565b82601f106102b05782800160ff198235161785556102dd565b828001600101855582156102dd579182015b828111156102dd5782358255916020019190600101906102c2565b506102e99291506102ed565b5090565b5b808211156102e957600081556001016102ee565b600181811c9082168061031657607f821691505b60208210810361033657634e487b7160e01b600052602260045260246000fd5b50919050565b600080835481600182811c91508083168061035857607f831692505b6020808410820361037757634e487b7160e01b86526022600452602486fd5b81801561038b576001811461039c576103c9565b60ff198616895284890196506103c9565b60008a81526020902060005b868110156103c15781548b8201529085019083016103a8565b505084890196505b509498975050505050505050565b600060208083528351808285015260005b81811015610404578581018301518582016040015282016103e8565b81811115610416576000604083870101525b50601f01601f1916929092016040019392505050565b6000806000806060858703121561044257600080fd5b84356001600160a01b038116811461045957600080fd5b935060208501359250604085013567ffffffffffffffff8082111561047d57600080fd5b818701915087601f83011261049157600080fd5b8135818111156104a057600080fd5b8860208285010111156104b257600080fd5b9598949750506020019450505056fea264697066735822122018dd705e73323976c54bdb98fe5b215830dc46842afd2a416c9c055d97f85f1964736f6c634300080d0033",
}

// ReentererABI is the input ABI used to generate the binding from.
// Deprecated: Use ReentererMetaData.ABI instead.
var ReentererABI = ReentererMetaData.ABI

// ReentererBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReentererMetaData.Bin instead.
var ReentererBin = ReentererMetaData.Bin

// DeployReenterer deploys a new Ethereum contract, binding an instance of Reenterer to it.
func DeployReenterer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Reenterer, error) {
	parsed, err := ReentererMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReentererBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Reenterer{ReentererCaller: ReentererCaller{contract: contract}, ReentererTransactor: ReentererTransactor{contract: contract}, ReentererFilterer: ReentererFilterer{contract: contract}}, nil
}

// Reenterer is an auto generated Go binding around an Ethereum contract.
type Reenterer struct {
	ReentererCaller     // Read-only binding to the contract
	ReentererTransactor // Write-only binding to the contract
	ReentererFilterer   // Log filterer for contract events
}

// ReentererCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentererCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentererTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentererTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentererFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentererFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentererSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentererSession struct {
	Contract     *Reenterer        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentererCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentererCallerSession struct {
	Contract *ReentererCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ReentererTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentererTransactorSession struct {
	Contract     *ReentererTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ReentererRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentererRaw struct {
	Contract *Reenterer // Generic contract binding to access the raw methods on
}

// ReentererCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentererCallerRaw struct {
	Contract *ReentererCaller // Generic read-only contract binding to access the raw methods on
}

// ReentererTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentererTransactorRaw struct {
	Contract *ReentererTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReenterer creates a new instance of Reenterer, bound to a specific deployed contract.
func NewReenterer(address common.Address, backend bind.ContractBackend) (*Reenterer, error) {
	contract, err := bindReenterer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reenterer{ReentererCaller: ReentererCaller{contract: contract}, ReentererTransactor: ReentererTransactor{contract: contract}, ReentererFilterer: ReentererFilterer{contract: contract}}, nil
}

// NewReentererCaller creates a new read-only instance of Reenterer, bound to a specific deployed contract.
func NewReentererCaller(address common.Address, caller bind.ContractCaller) (*ReentererCaller, error) {
	contract, err := bindReenterer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentererCaller{contract: contract}, nil
}

// NewReentererTransactor creates a new write-only instance of Reenterer, bound to a specific deployed contract.
func NewReentererTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentererTransactor, error) {
	contract, err := bindReenterer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentererTransactor{contract: contract}, nil
}

// NewReentererFilterer creates a new log filterer instance of Reenterer, bound to a specific deployed contract.
func NewReentererFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentererFilterer, error) {
	contract, err := bindReenterer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentererFilterer{contract: contract}, nil
}

// bindReenterer binds a generic wrapper to an already deployed contract.
func bindReenterer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentererABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reenterer *ReentererRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reenterer.Contract.ReentererCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reenterer *ReentererRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reenterer.Contract.ReentererTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reenterer *ReentererRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reenterer.Contract.ReentererTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reenterer *ReentererCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reenterer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reenterer *ReentererTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reenterer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reenterer *ReentererTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reenterer.Contract.contract.Transact(opts, method, params...)
}

// CallData is a free data retrieval call binding the contract method 0x4e417a98.
//
// Solidity: function callData() view returns(bytes)
func (_Reenterer *ReentererCaller) CallData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Reenterer.contract.Call(opts, &out, "callData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CallData is a free data retrieval call binding the contract method 0x4e417a98.
//
// Solidity: function callData() view returns(bytes)
func (_Reenterer *ReentererSession) CallData() ([]byte, error) {
	return _Reenterer.Contract.CallData(&_Reenterer.CallOpts)
}

// CallData is a free data retrieval call binding the contract method 0x4e417a98.
//
// Solidity: function callData() view returns(bytes)
func (_Reenterer *ReentererCallerSession) CallData() ([]byte, error) {
	return _Reenterer.Contract.CallData(&_Reenterer.CallOpts)
}

// MsgValue is a free data retrieval call binding the contract method 0xddf363d7.
//
// Solidity: function msgValue() view returns(uint256)
func (_Reenterer *ReentererCaller) MsgValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reenterer.contract.Call(opts, &out, "msgValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MsgValue is a free data retrieval call binding the contract method 0xddf363d7.
//
// Solidity: function msgValue() view returns(uint256)
func (_Reenterer *ReentererSession) MsgValue() (*big.Int, error) {
	return _Reenterer.Contract.MsgValue(&_Reenterer.CallOpts)
}

// MsgValue is a free data retrieval call binding the contract method 0xddf363d7.
//
// Solidity: function msgValue() view returns(uint256)
func (_Reenterer *ReentererCallerSession) MsgValue() (*big.Int, error) {
	return _Reenterer.Contract.MsgValue(&_Reenterer.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_Reenterer *ReentererCaller) Target(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reenterer.contract.Call(opts, &out, "target")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_Reenterer *ReentererSession) Target() (common.Address, error) {
	return _Reenterer.Contract.Target(&_Reenterer.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_Reenterer *ReentererCallerSession) Target() (common.Address, error) {
	return _Reenterer.Contract.Target(&_Reenterer.CallOpts)
}

// Prepare is a paid mutator transaction binding the contract method 0x5157589b.
//
// Solidity: function prepare(address targetToUse, uint256 msgValueToUse, bytes callDataToUse) returns()
func (_Reenterer *ReentererTransactor) Prepare(opts *bind.TransactOpts, targetToUse common.Address, msgValueToUse *big.Int, callDataToUse []byte) (*types.Transaction, error) {
	return _Reenterer.contract.Transact(opts, "prepare", targetToUse, msgValueToUse, callDataToUse)
}

// Prepare is a paid mutator transaction binding the contract method 0x5157589b.
//
// Solidity: function prepare(address targetToUse, uint256 msgValueToUse, bytes callDataToUse) returns()
func (_Reenterer *ReentererSession) Prepare(targetToUse common.Address, msgValueToUse *big.Int, callDataToUse []byte) (*types.Transaction, error) {
	return _Reenterer.Contract.Prepare(&_Reenterer.TransactOpts, targetToUse, msgValueToUse, callDataToUse)
}

// Prepare is a paid mutator transaction binding the contract method 0x5157589b.
//
// Solidity: function prepare(address targetToUse, uint256 msgValueToUse, bytes callDataToUse) returns()
func (_Reenterer *ReentererTransactorSession) Prepare(targetToUse common.Address, msgValueToUse *big.Int, callDataToUse []byte) (*types.Transaction, error) {
	return _Reenterer.Contract.Prepare(&_Reenterer.TransactOpts, targetToUse, msgValueToUse, callDataToUse)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Reenterer *ReentererTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reenterer.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Reenterer *ReentererSession) Receive() (*types.Transaction, error) {
	return _Reenterer.Contract.Receive(&_Reenterer.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Reenterer *ReentererTransactorSession) Receive() (*types.Transaction, error) {
	return _Reenterer.Contract.Receive(&_Reenterer.TransactOpts)
}

// ReentererReenteredIterator is returned from FilterReentered and is used to iterate over the raw logs and unpacked data for Reentered events raised by the Reenterer contract.
type ReentererReenteredIterator struct {
	Event *ReentererReentered // Event containing the contract specifics and raw log

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
func (it *ReentererReenteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReentererReentered)
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
		it.Event = new(ReentererReentered)
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
func (it *ReentererReenteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReentererReenteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReentererReentered represents a Reentered event raised by the Reenterer contract.
type ReentererReentered struct {
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReentered is a free log retrieval operation binding the contract event 0x52b1b52a16f834754c20ed40ad488c0fc799dea307f9bb9cde6413f0e356703e.
//
// Solidity: event Reentered(bytes returnData)
func (_Reenterer *ReentererFilterer) FilterReentered(opts *bind.FilterOpts) (*ReentererReenteredIterator, error) {

	logs, sub, err := _Reenterer.contract.FilterLogs(opts, "Reentered")
	if err != nil {
		return nil, err
	}
	return &ReentererReenteredIterator{contract: _Reenterer.contract, event: "Reentered", logs: logs, sub: sub}, nil
}

// WatchReentered is a free log subscription operation binding the contract event 0x52b1b52a16f834754c20ed40ad488c0fc799dea307f9bb9cde6413f0e356703e.
//
// Solidity: event Reentered(bytes returnData)
func (_Reenterer *ReentererFilterer) WatchReentered(opts *bind.WatchOpts, sink chan<- *ReentererReentered) (event.Subscription, error) {

	logs, sub, err := _Reenterer.contract.WatchLogs(opts, "Reentered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReentererReentered)
				if err := _Reenterer.contract.UnpackLog(event, "Reentered", log); err != nil {
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

// ParseReentered is a log parse operation binding the contract event 0x52b1b52a16f834754c20ed40ad488c0fc799dea307f9bb9cde6413f0e356703e.
//
// Solidity: event Reentered(bytes returnData)
func (_Reenterer *ReentererFilterer) ParseReentered(log types.Log) (*ReentererReentered, error) {
	event := new(ReentererReentered)
	if err := _Reenterer.contract.UnpackLog(event, "Reentered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
