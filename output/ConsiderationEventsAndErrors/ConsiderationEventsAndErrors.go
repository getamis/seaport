// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ConsiderationEventsAndErrors

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

// ReceivedItem is an auto generated low-level Go binding around an user-defined struct.
type ReceivedItem struct {
	ItemType   uint8
	Token      common.Address
	Identifier *big.Int
	Amount     *big.Int
	Recipient  common.Address
}

// SpentItem is an auto generated low-level Go binding around an user-defined struct.
type SpentItem struct {
	ItemType   uint8
	Token      common.Address
	Identifier *big.Int
	Amount     *big.Int
}

// ConsiderationEventsAndErrorsMetaData contains all meta data concerning the ConsiderationEventsAndErrors contract.
var ConsiderationEventsAndErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BadFraction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortfallAmount\",\"type\":\"uint256\"}],\"name\":\"ConsiderationNotMet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EtherTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientEtherSupplied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBasicOrderParameterEncoding\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidCallToConduit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCanceller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidConduit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNativeOfferItem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingOriginalConsiderationItems\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSpecifiedOrdersAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderAlreadyFilled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderIsCancelled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderPartiallyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PartialFillsNotEnabledForOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCounter\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"}],\"name\":\"CounterIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSpentItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structReceivedItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"}],\"name\":\"OrderValidated\",\"type\":\"event\"}]",
}

// ConsiderationEventsAndErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ConsiderationEventsAndErrorsMetaData.ABI instead.
var ConsiderationEventsAndErrorsABI = ConsiderationEventsAndErrorsMetaData.ABI

// ConsiderationEventsAndErrors is an auto generated Go binding around an Ethereum contract.
type ConsiderationEventsAndErrors struct {
	ConsiderationEventsAndErrorsCaller     // Read-only binding to the contract
	ConsiderationEventsAndErrorsTransactor // Write-only binding to the contract
	ConsiderationEventsAndErrorsFilterer   // Log filterer for contract events
}

// ConsiderationEventsAndErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConsiderationEventsAndErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsiderationEventsAndErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConsiderationEventsAndErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsiderationEventsAndErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConsiderationEventsAndErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsiderationEventsAndErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConsiderationEventsAndErrorsSession struct {
	Contract     *ConsiderationEventsAndErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ConsiderationEventsAndErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConsiderationEventsAndErrorsCallerSession struct {
	Contract *ConsiderationEventsAndErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// ConsiderationEventsAndErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConsiderationEventsAndErrorsTransactorSession struct {
	Contract     *ConsiderationEventsAndErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// ConsiderationEventsAndErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConsiderationEventsAndErrorsRaw struct {
	Contract *ConsiderationEventsAndErrors // Generic contract binding to access the raw methods on
}

// ConsiderationEventsAndErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConsiderationEventsAndErrorsCallerRaw struct {
	Contract *ConsiderationEventsAndErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// ConsiderationEventsAndErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConsiderationEventsAndErrorsTransactorRaw struct {
	Contract *ConsiderationEventsAndErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConsiderationEventsAndErrors creates a new instance of ConsiderationEventsAndErrors, bound to a specific deployed contract.
func NewConsiderationEventsAndErrors(address common.Address, backend bind.ContractBackend) (*ConsiderationEventsAndErrors, error) {
	contract, err := bindConsiderationEventsAndErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConsiderationEventsAndErrors{ConsiderationEventsAndErrorsCaller: ConsiderationEventsAndErrorsCaller{contract: contract}, ConsiderationEventsAndErrorsTransactor: ConsiderationEventsAndErrorsTransactor{contract: contract}, ConsiderationEventsAndErrorsFilterer: ConsiderationEventsAndErrorsFilterer{contract: contract}}, nil
}

// NewConsiderationEventsAndErrorsCaller creates a new read-only instance of ConsiderationEventsAndErrors, bound to a specific deployed contract.
func NewConsiderationEventsAndErrorsCaller(address common.Address, caller bind.ContractCaller) (*ConsiderationEventsAndErrorsCaller, error) {
	contract, err := bindConsiderationEventsAndErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConsiderationEventsAndErrorsCaller{contract: contract}, nil
}

// NewConsiderationEventsAndErrorsTransactor creates a new write-only instance of ConsiderationEventsAndErrors, bound to a specific deployed contract.
func NewConsiderationEventsAndErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ConsiderationEventsAndErrorsTransactor, error) {
	contract, err := bindConsiderationEventsAndErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConsiderationEventsAndErrorsTransactor{contract: contract}, nil
}

// NewConsiderationEventsAndErrorsFilterer creates a new log filterer instance of ConsiderationEventsAndErrors, bound to a specific deployed contract.
func NewConsiderationEventsAndErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ConsiderationEventsAndErrorsFilterer, error) {
	contract, err := bindConsiderationEventsAndErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConsiderationEventsAndErrorsFilterer{contract: contract}, nil
}

// bindConsiderationEventsAndErrors binds a generic wrapper to an already deployed contract.
func bindConsiderationEventsAndErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConsiderationEventsAndErrorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConsiderationEventsAndErrors.Contract.ConsiderationEventsAndErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsiderationEventsAndErrors.Contract.ConsiderationEventsAndErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsiderationEventsAndErrors.Contract.ConsiderationEventsAndErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConsiderationEventsAndErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsiderationEventsAndErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsiderationEventsAndErrors.Contract.contract.Transact(opts, method, params...)
}

// ConsiderationEventsAndErrorsCounterIncrementedIterator is returned from FilterCounterIncremented and is used to iterate over the raw logs and unpacked data for CounterIncremented events raised by the ConsiderationEventsAndErrors contract.
type ConsiderationEventsAndErrorsCounterIncrementedIterator struct {
	Event *ConsiderationEventsAndErrorsCounterIncremented // Event containing the contract specifics and raw log

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
func (it *ConsiderationEventsAndErrorsCounterIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsiderationEventsAndErrorsCounterIncremented)
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
		it.Event = new(ConsiderationEventsAndErrorsCounterIncremented)
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
func (it *ConsiderationEventsAndErrorsCounterIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsiderationEventsAndErrorsCounterIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsiderationEventsAndErrorsCounterIncremented represents a CounterIncremented event raised by the ConsiderationEventsAndErrors contract.
type ConsiderationEventsAndErrorsCounterIncremented struct {
	NewCounter *big.Int
	Offerer    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCounterIncremented is a free log retrieval operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsFilterer) FilterCounterIncremented(opts *bind.FilterOpts, offerer []common.Address) (*ConsiderationEventsAndErrorsCounterIncrementedIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _ConsiderationEventsAndErrors.contract.FilterLogs(opts, "CounterIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return &ConsiderationEventsAndErrorsCounterIncrementedIterator{contract: _ConsiderationEventsAndErrors.contract, event: "CounterIncremented", logs: logs, sub: sub}, nil
}

// WatchCounterIncremented is a free log subscription operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsFilterer) WatchCounterIncremented(opts *bind.WatchOpts, sink chan<- *ConsiderationEventsAndErrorsCounterIncremented, offerer []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _ConsiderationEventsAndErrors.contract.WatchLogs(opts, "CounterIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsiderationEventsAndErrorsCounterIncremented)
				if err := _ConsiderationEventsAndErrors.contract.UnpackLog(event, "CounterIncremented", log); err != nil {
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

// ParseCounterIncremented is a log parse operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsFilterer) ParseCounterIncremented(log types.Log) (*ConsiderationEventsAndErrorsCounterIncremented, error) {
	event := new(ConsiderationEventsAndErrorsCounterIncremented)
	if err := _ConsiderationEventsAndErrors.contract.UnpackLog(event, "CounterIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConsiderationEventsAndErrorsOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the ConsiderationEventsAndErrors contract.
type ConsiderationEventsAndErrorsOrderCancelledIterator struct {
	Event *ConsiderationEventsAndErrorsOrderCancelled // Event containing the contract specifics and raw log

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
func (it *ConsiderationEventsAndErrorsOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsiderationEventsAndErrorsOrderCancelled)
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
		it.Event = new(ConsiderationEventsAndErrorsOrderCancelled)
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
func (it *ConsiderationEventsAndErrorsOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsiderationEventsAndErrorsOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsiderationEventsAndErrorsOrderCancelled represents a OrderCancelled event raised by the ConsiderationEventsAndErrors contract.
type ConsiderationEventsAndErrorsOrderCancelled struct {
	OrderHash [32]byte
	Offerer   common.Address
	Zone      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsFilterer) FilterOrderCancelled(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*ConsiderationEventsAndErrorsOrderCancelledIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _ConsiderationEventsAndErrors.contract.FilterLogs(opts, "OrderCancelled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &ConsiderationEventsAndErrorsOrderCancelledIterator{contract: _ConsiderationEventsAndErrors.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *ConsiderationEventsAndErrorsOrderCancelled, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _ConsiderationEventsAndErrors.contract.WatchLogs(opts, "OrderCancelled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsiderationEventsAndErrorsOrderCancelled)
				if err := _ConsiderationEventsAndErrors.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsFilterer) ParseOrderCancelled(log types.Log) (*ConsiderationEventsAndErrorsOrderCancelled, error) {
	event := new(ConsiderationEventsAndErrorsOrderCancelled)
	if err := _ConsiderationEventsAndErrors.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConsiderationEventsAndErrorsOrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the ConsiderationEventsAndErrors contract.
type ConsiderationEventsAndErrorsOrderFulfilledIterator struct {
	Event *ConsiderationEventsAndErrorsOrderFulfilled // Event containing the contract specifics and raw log

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
func (it *ConsiderationEventsAndErrorsOrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsiderationEventsAndErrorsOrderFulfilled)
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
		it.Event = new(ConsiderationEventsAndErrorsOrderFulfilled)
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
func (it *ConsiderationEventsAndErrorsOrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsiderationEventsAndErrorsOrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsiderationEventsAndErrorsOrderFulfilled represents a OrderFulfilled event raised by the ConsiderationEventsAndErrors contract.
type ConsiderationEventsAndErrorsOrderFulfilled struct {
	OrderHash     [32]byte
	Offerer       common.Address
	Zone          common.Address
	Recipient     common.Address
	Offer         []SpentItem
	Consideration []ReceivedItem
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address recipient, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsFilterer) FilterOrderFulfilled(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*ConsiderationEventsAndErrorsOrderFulfilledIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _ConsiderationEventsAndErrors.contract.FilterLogs(opts, "OrderFulfilled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &ConsiderationEventsAndErrorsOrderFulfilledIterator{contract: _ConsiderationEventsAndErrors.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address recipient, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsFilterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *ConsiderationEventsAndErrorsOrderFulfilled, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _ConsiderationEventsAndErrors.contract.WatchLogs(opts, "OrderFulfilled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsiderationEventsAndErrorsOrderFulfilled)
				if err := _ConsiderationEventsAndErrors.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
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

// ParseOrderFulfilled is a log parse operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address recipient, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsFilterer) ParseOrderFulfilled(log types.Log) (*ConsiderationEventsAndErrorsOrderFulfilled, error) {
	event := new(ConsiderationEventsAndErrorsOrderFulfilled)
	if err := _ConsiderationEventsAndErrors.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConsiderationEventsAndErrorsOrderValidatedIterator is returned from FilterOrderValidated and is used to iterate over the raw logs and unpacked data for OrderValidated events raised by the ConsiderationEventsAndErrors contract.
type ConsiderationEventsAndErrorsOrderValidatedIterator struct {
	Event *ConsiderationEventsAndErrorsOrderValidated // Event containing the contract specifics and raw log

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
func (it *ConsiderationEventsAndErrorsOrderValidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsiderationEventsAndErrorsOrderValidated)
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
		it.Event = new(ConsiderationEventsAndErrorsOrderValidated)
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
func (it *ConsiderationEventsAndErrorsOrderValidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsiderationEventsAndErrorsOrderValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsiderationEventsAndErrorsOrderValidated represents a OrderValidated event raised by the ConsiderationEventsAndErrors contract.
type ConsiderationEventsAndErrorsOrderValidated struct {
	OrderHash [32]byte
	Offerer   common.Address
	Zone      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderValidated is a free log retrieval operation binding the contract event 0xfde361574a066b44b3b5fe98a87108b7565e327327954c4faeea56a4e6491a0a.
//
// Solidity: event OrderValidated(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsFilterer) FilterOrderValidated(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*ConsiderationEventsAndErrorsOrderValidatedIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _ConsiderationEventsAndErrors.contract.FilterLogs(opts, "OrderValidated", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &ConsiderationEventsAndErrorsOrderValidatedIterator{contract: _ConsiderationEventsAndErrors.contract, event: "OrderValidated", logs: logs, sub: sub}, nil
}

// WatchOrderValidated is a free log subscription operation binding the contract event 0xfde361574a066b44b3b5fe98a87108b7565e327327954c4faeea56a4e6491a0a.
//
// Solidity: event OrderValidated(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsFilterer) WatchOrderValidated(opts *bind.WatchOpts, sink chan<- *ConsiderationEventsAndErrorsOrderValidated, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _ConsiderationEventsAndErrors.contract.WatchLogs(opts, "OrderValidated", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsiderationEventsAndErrorsOrderValidated)
				if err := _ConsiderationEventsAndErrors.contract.UnpackLog(event, "OrderValidated", log); err != nil {
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

// ParseOrderValidated is a log parse operation binding the contract event 0xfde361574a066b44b3b5fe98a87108b7565e327327954c4faeea56a4e6491a0a.
//
// Solidity: event OrderValidated(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_ConsiderationEventsAndErrors *ConsiderationEventsAndErrorsFilterer) ParseOrderValidated(log types.Log) (*ConsiderationEventsAndErrorsOrderValidated, error) {
	event := new(ConsiderationEventsAndErrorsOrderValidated)
	if err := _ConsiderationEventsAndErrors.contract.UnpackLog(event, "OrderValidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
