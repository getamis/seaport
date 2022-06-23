// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ConduitInterface

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

// ConduitBatch1155Transfer is an auto generated low-level Go binding around an user-defined struct.
type ConduitBatch1155Transfer struct {
	Token   common.Address
	From    common.Address
	To      common.Address
	Ids     []*big.Int
	Amounts []*big.Int
}

// ConduitTransfer is an auto generated low-level Go binding around an user-defined struct.
type ConduitTransfer struct {
	ItemType   uint8
	Token      common.Address
	From       common.Address
	To         common.Address
	Identifier *big.Int
	Amount     *big.Int
}

// ConduitInterfaceMetaData contains all meta data concerning the ConduitInterface contract.
var ConduitInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ChannelClosed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Invalid1155BatchTransferEncoding\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidController\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidItemType\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"channel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"open\",\"type\":\"bool\"}],\"name\":\"ChannelUpdated\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumConduitItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structConduitTransfer[]\",\"name\":\"transfers\",\"type\":\"tuple[]\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structConduitBatch1155Transfer[]\",\"name\":\"batch1155Transfers\",\"type\":\"tuple[]\"}],\"name\":\"executeBatch1155\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumConduitItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structConduitTransfer[]\",\"name\":\"standardTransfers\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structConduitBatch1155Transfer[]\",\"name\":\"batch1155Transfers\",\"type\":\"tuple[]\"}],\"name\":\"executeWithBatch1155\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"channel\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"updateChannel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ConduitInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use ConduitInterfaceMetaData.ABI instead.
var ConduitInterfaceABI = ConduitInterfaceMetaData.ABI

// ConduitInterface is an auto generated Go binding around an Ethereum contract.
type ConduitInterface struct {
	ConduitInterfaceCaller     // Read-only binding to the contract
	ConduitInterfaceTransactor // Write-only binding to the contract
	ConduitInterfaceFilterer   // Log filterer for contract events
}

// ConduitInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConduitInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConduitInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConduitInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConduitInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConduitInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConduitInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConduitInterfaceSession struct {
	Contract     *ConduitInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConduitInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConduitInterfaceCallerSession struct {
	Contract *ConduitInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ConduitInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConduitInterfaceTransactorSession struct {
	Contract     *ConduitInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ConduitInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConduitInterfaceRaw struct {
	Contract *ConduitInterface // Generic contract binding to access the raw methods on
}

// ConduitInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConduitInterfaceCallerRaw struct {
	Contract *ConduitInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ConduitInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConduitInterfaceTransactorRaw struct {
	Contract *ConduitInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConduitInterface creates a new instance of ConduitInterface, bound to a specific deployed contract.
func NewConduitInterface(address common.Address, backend bind.ContractBackend) (*ConduitInterface, error) {
	contract, err := bindConduitInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConduitInterface{ConduitInterfaceCaller: ConduitInterfaceCaller{contract: contract}, ConduitInterfaceTransactor: ConduitInterfaceTransactor{contract: contract}, ConduitInterfaceFilterer: ConduitInterfaceFilterer{contract: contract}}, nil
}

// NewConduitInterfaceCaller creates a new read-only instance of ConduitInterface, bound to a specific deployed contract.
func NewConduitInterfaceCaller(address common.Address, caller bind.ContractCaller) (*ConduitInterfaceCaller, error) {
	contract, err := bindConduitInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConduitInterfaceCaller{contract: contract}, nil
}

// NewConduitInterfaceTransactor creates a new write-only instance of ConduitInterface, bound to a specific deployed contract.
func NewConduitInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ConduitInterfaceTransactor, error) {
	contract, err := bindConduitInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConduitInterfaceTransactor{contract: contract}, nil
}

// NewConduitInterfaceFilterer creates a new log filterer instance of ConduitInterface, bound to a specific deployed contract.
func NewConduitInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ConduitInterfaceFilterer, error) {
	contract, err := bindConduitInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConduitInterfaceFilterer{contract: contract}, nil
}

// bindConduitInterface binds a generic wrapper to an already deployed contract.
func bindConduitInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConduitInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConduitInterface *ConduitInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConduitInterface.Contract.ConduitInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConduitInterface *ConduitInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConduitInterface.Contract.ConduitInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConduitInterface *ConduitInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConduitInterface.Contract.ConduitInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConduitInterface *ConduitInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConduitInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConduitInterface *ConduitInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConduitInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConduitInterface *ConduitInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConduitInterface.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0x4ce34aa2.
//
// Solidity: function execute((uint8,address,address,address,uint256,uint256)[] transfers) returns(bytes4 magicValue)
func (_ConduitInterface *ConduitInterfaceTransactor) Execute(opts *bind.TransactOpts, transfers []ConduitTransfer) (*types.Transaction, error) {
	return _ConduitInterface.contract.Transact(opts, "execute", transfers)
}

// Execute is a paid mutator transaction binding the contract method 0x4ce34aa2.
//
// Solidity: function execute((uint8,address,address,address,uint256,uint256)[] transfers) returns(bytes4 magicValue)
func (_ConduitInterface *ConduitInterfaceSession) Execute(transfers []ConduitTransfer) (*types.Transaction, error) {
	return _ConduitInterface.Contract.Execute(&_ConduitInterface.TransactOpts, transfers)
}

// Execute is a paid mutator transaction binding the contract method 0x4ce34aa2.
//
// Solidity: function execute((uint8,address,address,address,uint256,uint256)[] transfers) returns(bytes4 magicValue)
func (_ConduitInterface *ConduitInterfaceTransactorSession) Execute(transfers []ConduitTransfer) (*types.Transaction, error) {
	return _ConduitInterface.Contract.Execute(&_ConduitInterface.TransactOpts, transfers)
}

// ExecuteBatch1155 is a paid mutator transaction binding the contract method 0x8df25d92.
//
// Solidity: function executeBatch1155((address,address,address,uint256[],uint256[])[] batch1155Transfers) returns(bytes4 magicValue)
func (_ConduitInterface *ConduitInterfaceTransactor) ExecuteBatch1155(opts *bind.TransactOpts, batch1155Transfers []ConduitBatch1155Transfer) (*types.Transaction, error) {
	return _ConduitInterface.contract.Transact(opts, "executeBatch1155", batch1155Transfers)
}

// ExecuteBatch1155 is a paid mutator transaction binding the contract method 0x8df25d92.
//
// Solidity: function executeBatch1155((address,address,address,uint256[],uint256[])[] batch1155Transfers) returns(bytes4 magicValue)
func (_ConduitInterface *ConduitInterfaceSession) ExecuteBatch1155(batch1155Transfers []ConduitBatch1155Transfer) (*types.Transaction, error) {
	return _ConduitInterface.Contract.ExecuteBatch1155(&_ConduitInterface.TransactOpts, batch1155Transfers)
}

// ExecuteBatch1155 is a paid mutator transaction binding the contract method 0x8df25d92.
//
// Solidity: function executeBatch1155((address,address,address,uint256[],uint256[])[] batch1155Transfers) returns(bytes4 magicValue)
func (_ConduitInterface *ConduitInterfaceTransactorSession) ExecuteBatch1155(batch1155Transfers []ConduitBatch1155Transfer) (*types.Transaction, error) {
	return _ConduitInterface.Contract.ExecuteBatch1155(&_ConduitInterface.TransactOpts, batch1155Transfers)
}

// ExecuteWithBatch1155 is a paid mutator transaction binding the contract method 0x899e104c.
//
// Solidity: function executeWithBatch1155((uint8,address,address,address,uint256,uint256)[] standardTransfers, (address,address,address,uint256[],uint256[])[] batch1155Transfers) returns(bytes4 magicValue)
func (_ConduitInterface *ConduitInterfaceTransactor) ExecuteWithBatch1155(opts *bind.TransactOpts, standardTransfers []ConduitTransfer, batch1155Transfers []ConduitBatch1155Transfer) (*types.Transaction, error) {
	return _ConduitInterface.contract.Transact(opts, "executeWithBatch1155", standardTransfers, batch1155Transfers)
}

// ExecuteWithBatch1155 is a paid mutator transaction binding the contract method 0x899e104c.
//
// Solidity: function executeWithBatch1155((uint8,address,address,address,uint256,uint256)[] standardTransfers, (address,address,address,uint256[],uint256[])[] batch1155Transfers) returns(bytes4 magicValue)
func (_ConduitInterface *ConduitInterfaceSession) ExecuteWithBatch1155(standardTransfers []ConduitTransfer, batch1155Transfers []ConduitBatch1155Transfer) (*types.Transaction, error) {
	return _ConduitInterface.Contract.ExecuteWithBatch1155(&_ConduitInterface.TransactOpts, standardTransfers, batch1155Transfers)
}

// ExecuteWithBatch1155 is a paid mutator transaction binding the contract method 0x899e104c.
//
// Solidity: function executeWithBatch1155((uint8,address,address,address,uint256,uint256)[] standardTransfers, (address,address,address,uint256[],uint256[])[] batch1155Transfers) returns(bytes4 magicValue)
func (_ConduitInterface *ConduitInterfaceTransactorSession) ExecuteWithBatch1155(standardTransfers []ConduitTransfer, batch1155Transfers []ConduitBatch1155Transfer) (*types.Transaction, error) {
	return _ConduitInterface.Contract.ExecuteWithBatch1155(&_ConduitInterface.TransactOpts, standardTransfers, batch1155Transfers)
}

// UpdateChannel is a paid mutator transaction binding the contract method 0xc4e8fcb5.
//
// Solidity: function updateChannel(address channel, bool isOpen) returns()
func (_ConduitInterface *ConduitInterfaceTransactor) UpdateChannel(opts *bind.TransactOpts, channel common.Address, isOpen bool) (*types.Transaction, error) {
	return _ConduitInterface.contract.Transact(opts, "updateChannel", channel, isOpen)
}

// UpdateChannel is a paid mutator transaction binding the contract method 0xc4e8fcb5.
//
// Solidity: function updateChannel(address channel, bool isOpen) returns()
func (_ConduitInterface *ConduitInterfaceSession) UpdateChannel(channel common.Address, isOpen bool) (*types.Transaction, error) {
	return _ConduitInterface.Contract.UpdateChannel(&_ConduitInterface.TransactOpts, channel, isOpen)
}

// UpdateChannel is a paid mutator transaction binding the contract method 0xc4e8fcb5.
//
// Solidity: function updateChannel(address channel, bool isOpen) returns()
func (_ConduitInterface *ConduitInterfaceTransactorSession) UpdateChannel(channel common.Address, isOpen bool) (*types.Transaction, error) {
	return _ConduitInterface.Contract.UpdateChannel(&_ConduitInterface.TransactOpts, channel, isOpen)
}

// ConduitInterfaceChannelUpdatedIterator is returned from FilterChannelUpdated and is used to iterate over the raw logs and unpacked data for ChannelUpdated events raised by the ConduitInterface contract.
type ConduitInterfaceChannelUpdatedIterator struct {
	Event *ConduitInterfaceChannelUpdated // Event containing the contract specifics and raw log

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
func (it *ConduitInterfaceChannelUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConduitInterfaceChannelUpdated)
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
		it.Event = new(ConduitInterfaceChannelUpdated)
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
func (it *ConduitInterfaceChannelUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConduitInterfaceChannelUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConduitInterfaceChannelUpdated represents a ChannelUpdated event raised by the ConduitInterface contract.
type ConduitInterfaceChannelUpdated struct {
	Channel common.Address
	Open    bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChannelUpdated is a free log retrieval operation binding the contract event 0xae63067d43ac07563b7eb8db6595635fc77f1578a2a5ea06ba91b63e2afa37e2.
//
// Solidity: event ChannelUpdated(address channel, bool open)
func (_ConduitInterface *ConduitInterfaceFilterer) FilterChannelUpdated(opts *bind.FilterOpts) (*ConduitInterfaceChannelUpdatedIterator, error) {

	logs, sub, err := _ConduitInterface.contract.FilterLogs(opts, "ChannelUpdated")
	if err != nil {
		return nil, err
	}
	return &ConduitInterfaceChannelUpdatedIterator{contract: _ConduitInterface.contract, event: "ChannelUpdated", logs: logs, sub: sub}, nil
}

// WatchChannelUpdated is a free log subscription operation binding the contract event 0xae63067d43ac07563b7eb8db6595635fc77f1578a2a5ea06ba91b63e2afa37e2.
//
// Solidity: event ChannelUpdated(address channel, bool open)
func (_ConduitInterface *ConduitInterfaceFilterer) WatchChannelUpdated(opts *bind.WatchOpts, sink chan<- *ConduitInterfaceChannelUpdated) (event.Subscription, error) {

	logs, sub, err := _ConduitInterface.contract.WatchLogs(opts, "ChannelUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConduitInterfaceChannelUpdated)
				if err := _ConduitInterface.contract.UnpackLog(event, "ChannelUpdated", log); err != nil {
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

// ParseChannelUpdated is a log parse operation binding the contract event 0xae63067d43ac07563b7eb8db6595635fc77f1578a2a5ea06ba91b63e2afa37e2.
//
// Solidity: event ChannelUpdated(address channel, bool open)
func (_ConduitInterface *ConduitInterfaceFilterer) ParseChannelUpdated(log types.Log) (*ConduitInterfaceChannelUpdated, error) {
	event := new(ConduitInterfaceChannelUpdated)
	if err := _ConduitInterface.contract.UnpackLog(event, "ChannelUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
