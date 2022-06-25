// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Assertions

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

// AssertionsMetaData contains all meta data concerning the Assertions contract.
var AssertionsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduitController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadFraction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BadReturnValueFromERC20OnTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortfallAmount\",\"type\":\"uint256\"}],\"name\":\"ConsiderationNotMet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"identifiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"ERC1155BatchTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EtherTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientEtherSupplied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Invalid1155BatchTransferEncoding\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBasicOrderParameterEncoding\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidCallToConduit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCanceller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidConduit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidERC721TransferAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNativeOfferItem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingItemAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingOriginalConsiderationItems\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NoContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoReentrantCalls\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSpecifiedOrdersAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderAlreadyFilled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderIsCancelled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderPartiallyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PartialFillsNotEnabledForOrder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnusedItemParameters\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCounter\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"}],\"name\":\"CounterIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSpentItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structReceivedItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"}],\"name\":\"OrderValidated\",\"type\":\"event\"}]",
	Bin: "0x6101c060405234801561001157600080fd5b506040516106fc3803806106fc833981016040819052610030916105ac565b808061003a610116565b610120526101005260e05260c081815260a0838152608085815246610140819052604080516020818101979097528082019890985260608801969096529086015230858201528351808603909101815293019091528151910120610160526001600160a01b03811661018081905260408051630a96ad3960e01b81528151630a96ad39926004808401939192918290030181865afa1580156100e0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061010491906105dc565b506101a0525050600160005550610662565b6000808080808061014760408051808201909152600d81526c21b7b739b4b232b930ba34b7b760991b602082015290565b80516020918201206040805180820182526003815262312e3160e81b90840152519097507f722c0e0c80487266e8c6a45e3a1a803aab23378a9c32e6ebe029d4fad7bfc965965060009161024b91016909ecccccae492e8cada560b31b81526e1d5a5b9d0e081a5d195b551e5c194b608a1b600a8201526d1859191c995cdcc81d1bdad95b8b60921b60198201527f75696e74323536206964656e7469666965724f7243726974657269612c00000060278201527f75696e74323536207374617274416d6f756e742c0000000000000000000000006044820152701d5a5b9d0c8d4d88195b99105b5bdd5b9d607a1b6058820152602960f81b6069820152606a0190565b60408051601f1981840301815282825271086dedce6d2c8cae4c2e8d2dedc92e8cada560731b60208401526e1d5a5b9d0e081a5d195b551e5c194b608a1b60328401526d1859191c995cdcc81d1bdad95b8b60921b60418401527f75696e74323536206964656e7469666965724f7243726974657269612c000000604f8401527f75696e74323536207374617274416d6f756e742c000000000000000000000000606c840152711d5a5b9d0c8d4d88195b99105b5bdd5b9d0b60721b6080840152701859191c995cdcc81c9958da5c1a595b9d607a1b6092840152602960f81b60a384018190528251808503608401815260a485019093526f09ee4c8cae486dedae0dedccadce8e6560831b60c48501526f1859191c995cdcc81bd999995c995c8b60821b60d48501526c1859191c995cdcc81e9bdb994b609a1b60e48501527113d999995c925d195b56d7481bd999995c8b60721b60f18501527f436f6e73696465726174696f6e4974656d5b5d20636f6e73696465726174696f610103850152611b8b60f21b6101238501526f1d5a5b9d0e081bdc99195c951e5c194b60821b610125850152711d5a5b9d0c8d4d881cdd185c9d151a5b594b60721b6101358501526f1d5a5b9d0c8d4d88195b99151a5b594b60821b61014785015270189e5d195ccccc881e9bdb9952185cda0b607a1b6101578501526c1d5a5b9d0c8d4d881cd85b1d0b609a1b6101688501527f6279746573333220636f6e647569744b65792c000000000000000000000000006101758501526e3ab4b73a191a9b1031b7bab73a32b960891b6101888501526101978401529250906000906101980160408051601f19818403018152908290526c08a92a06e626488dedac2d2dc5609b1b60208301526b1cdd1c9a5b99c81b985b594b60a21b602d8301526e1cdd1c9a5b99c81d995c9cda5bdb8b608a1b60398301526f1d5a5b9d0c8d4d8818da185a5b92590b60821b60488301527f6164647265737320766572696679696e67436f6e7472616374000000000000006058830152602960f81b6071830152915060720160408051601f19818403018152908290528051602091820120855186830120855186840120919a5098509650610589918391859187910161063b565b604051602081830303815290604052805190602001209350505050909192939495565b6000602082840312156105be57600080fd5b81516001600160a01b03811681146105d557600080fd5b9392505050565b600080604083850312156105ef57600080fd5b505080516020909101519092909150565b6000815160005b818110156106215760208185018101518683015201610607565b81811115610630576000828601525b509290920192915050565b600061065961065361064d8488610600565b86610600565b84610600565b95945050505050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a051603f6106bd60003960005050600050506000505060005050600050506000505060005050600050506000505060005050603f6000f3fe6080604052600080fdfea264697066735822122013a6016b27d5c3c3f8abfcd6ca91855766e97ef913c2a9193ebc63e79a368f0064736f6c634300080d0033",
}

// AssertionsABI is the input ABI used to generate the binding from.
// Deprecated: Use AssertionsMetaData.ABI instead.
var AssertionsABI = AssertionsMetaData.ABI

// AssertionsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AssertionsMetaData.Bin instead.
var AssertionsBin = AssertionsMetaData.Bin

// DeployAssertions deploys a new Ethereum contract, binding an instance of Assertions to it.
func DeployAssertions(auth *bind.TransactOpts, backend bind.ContractBackend, conduitController common.Address) (common.Address, *types.Transaction, *Assertions, error) {
	parsed, err := AssertionsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AssertionsBin), backend, conduitController)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Assertions{AssertionsCaller: AssertionsCaller{contract: contract}, AssertionsTransactor: AssertionsTransactor{contract: contract}, AssertionsFilterer: AssertionsFilterer{contract: contract}}, nil
}

// Assertions is an auto generated Go binding around an Ethereum contract.
type Assertions struct {
	AssertionsCaller     // Read-only binding to the contract
	AssertionsTransactor // Write-only binding to the contract
	AssertionsFilterer   // Log filterer for contract events
}

// AssertionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssertionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssertionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssertionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssertionsSession struct {
	Contract     *Assertions       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssertionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssertionsCallerSession struct {
	Contract *AssertionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AssertionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssertionsTransactorSession struct {
	Contract     *AssertionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AssertionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssertionsRaw struct {
	Contract *Assertions // Generic contract binding to access the raw methods on
}

// AssertionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssertionsCallerRaw struct {
	Contract *AssertionsCaller // Generic read-only contract binding to access the raw methods on
}

// AssertionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssertionsTransactorRaw struct {
	Contract *AssertionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssertions creates a new instance of Assertions, bound to a specific deployed contract.
func NewAssertions(address common.Address, backend bind.ContractBackend) (*Assertions, error) {
	contract, err := bindAssertions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Assertions{AssertionsCaller: AssertionsCaller{contract: contract}, AssertionsTransactor: AssertionsTransactor{contract: contract}, AssertionsFilterer: AssertionsFilterer{contract: contract}}, nil
}

// NewAssertionsCaller creates a new read-only instance of Assertions, bound to a specific deployed contract.
func NewAssertionsCaller(address common.Address, caller bind.ContractCaller) (*AssertionsCaller, error) {
	contract, err := bindAssertions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssertionsCaller{contract: contract}, nil
}

// NewAssertionsTransactor creates a new write-only instance of Assertions, bound to a specific deployed contract.
func NewAssertionsTransactor(address common.Address, transactor bind.ContractTransactor) (*AssertionsTransactor, error) {
	contract, err := bindAssertions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssertionsTransactor{contract: contract}, nil
}

// NewAssertionsFilterer creates a new log filterer instance of Assertions, bound to a specific deployed contract.
func NewAssertionsFilterer(address common.Address, filterer bind.ContractFilterer) (*AssertionsFilterer, error) {
	contract, err := bindAssertions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssertionsFilterer{contract: contract}, nil
}

// bindAssertions binds a generic wrapper to an already deployed contract.
func bindAssertions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssertionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Assertions *AssertionsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Assertions.Contract.AssertionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Assertions *AssertionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Assertions.Contract.AssertionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Assertions *AssertionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Assertions.Contract.AssertionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Assertions *AssertionsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Assertions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Assertions *AssertionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Assertions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Assertions *AssertionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Assertions.Contract.contract.Transact(opts, method, params...)
}

// AssertionsCounterIncrementedIterator is returned from FilterCounterIncremented and is used to iterate over the raw logs and unpacked data for CounterIncremented events raised by the Assertions contract.
type AssertionsCounterIncrementedIterator struct {
	Event *AssertionsCounterIncremented // Event containing the contract specifics and raw log

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
func (it *AssertionsCounterIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssertionsCounterIncremented)
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
		it.Event = new(AssertionsCounterIncremented)
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
func (it *AssertionsCounterIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssertionsCounterIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssertionsCounterIncremented represents a CounterIncremented event raised by the Assertions contract.
type AssertionsCounterIncremented struct {
	NewCounter *big.Int
	Offerer    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCounterIncremented is a free log retrieval operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_Assertions *AssertionsFilterer) FilterCounterIncremented(opts *bind.FilterOpts, offerer []common.Address) (*AssertionsCounterIncrementedIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Assertions.contract.FilterLogs(opts, "CounterIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return &AssertionsCounterIncrementedIterator{contract: _Assertions.contract, event: "CounterIncremented", logs: logs, sub: sub}, nil
}

// WatchCounterIncremented is a free log subscription operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_Assertions *AssertionsFilterer) WatchCounterIncremented(opts *bind.WatchOpts, sink chan<- *AssertionsCounterIncremented, offerer []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Assertions.contract.WatchLogs(opts, "CounterIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssertionsCounterIncremented)
				if err := _Assertions.contract.UnpackLog(event, "CounterIncremented", log); err != nil {
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
func (_Assertions *AssertionsFilterer) ParseCounterIncremented(log types.Log) (*AssertionsCounterIncremented, error) {
	event := new(AssertionsCounterIncremented)
	if err := _Assertions.contract.UnpackLog(event, "CounterIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssertionsOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Assertions contract.
type AssertionsOrderCancelledIterator struct {
	Event *AssertionsOrderCancelled // Event containing the contract specifics and raw log

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
func (it *AssertionsOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssertionsOrderCancelled)
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
		it.Event = new(AssertionsOrderCancelled)
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
func (it *AssertionsOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssertionsOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssertionsOrderCancelled represents a OrderCancelled event raised by the Assertions contract.
type AssertionsOrderCancelled struct {
	OrderHash [32]byte
	Offerer   common.Address
	Zone      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Assertions *AssertionsFilterer) FilterOrderCancelled(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*AssertionsOrderCancelledIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Assertions.contract.FilterLogs(opts, "OrderCancelled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &AssertionsOrderCancelledIterator{contract: _Assertions.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Assertions *AssertionsFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *AssertionsOrderCancelled, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Assertions.contract.WatchLogs(opts, "OrderCancelled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssertionsOrderCancelled)
				if err := _Assertions.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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
func (_Assertions *AssertionsFilterer) ParseOrderCancelled(log types.Log) (*AssertionsOrderCancelled, error) {
	event := new(AssertionsOrderCancelled)
	if err := _Assertions.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssertionsOrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the Assertions contract.
type AssertionsOrderFulfilledIterator struct {
	Event *AssertionsOrderFulfilled // Event containing the contract specifics and raw log

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
func (it *AssertionsOrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssertionsOrderFulfilled)
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
		it.Event = new(AssertionsOrderFulfilled)
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
func (it *AssertionsOrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssertionsOrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssertionsOrderFulfilled represents a OrderFulfilled event raised by the Assertions contract.
type AssertionsOrderFulfilled struct {
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
func (_Assertions *AssertionsFilterer) FilterOrderFulfilled(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*AssertionsOrderFulfilledIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Assertions.contract.FilterLogs(opts, "OrderFulfilled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &AssertionsOrderFulfilledIterator{contract: _Assertions.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address recipient, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
func (_Assertions *AssertionsFilterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *AssertionsOrderFulfilled, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Assertions.contract.WatchLogs(opts, "OrderFulfilled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssertionsOrderFulfilled)
				if err := _Assertions.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
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
func (_Assertions *AssertionsFilterer) ParseOrderFulfilled(log types.Log) (*AssertionsOrderFulfilled, error) {
	event := new(AssertionsOrderFulfilled)
	if err := _Assertions.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssertionsOrderValidatedIterator is returned from FilterOrderValidated and is used to iterate over the raw logs and unpacked data for OrderValidated events raised by the Assertions contract.
type AssertionsOrderValidatedIterator struct {
	Event *AssertionsOrderValidated // Event containing the contract specifics and raw log

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
func (it *AssertionsOrderValidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssertionsOrderValidated)
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
		it.Event = new(AssertionsOrderValidated)
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
func (it *AssertionsOrderValidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssertionsOrderValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssertionsOrderValidated represents a OrderValidated event raised by the Assertions contract.
type AssertionsOrderValidated struct {
	OrderHash [32]byte
	Offerer   common.Address
	Zone      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderValidated is a free log retrieval operation binding the contract event 0xfde361574a066b44b3b5fe98a87108b7565e327327954c4faeea56a4e6491a0a.
//
// Solidity: event OrderValidated(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Assertions *AssertionsFilterer) FilterOrderValidated(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*AssertionsOrderValidatedIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Assertions.contract.FilterLogs(opts, "OrderValidated", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &AssertionsOrderValidatedIterator{contract: _Assertions.contract, event: "OrderValidated", logs: logs, sub: sub}, nil
}

// WatchOrderValidated is a free log subscription operation binding the contract event 0xfde361574a066b44b3b5fe98a87108b7565e327327954c4faeea56a4e6491a0a.
//
// Solidity: event OrderValidated(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Assertions *AssertionsFilterer) WatchOrderValidated(opts *bind.WatchOpts, sink chan<- *AssertionsOrderValidated, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Assertions.contract.WatchLogs(opts, "OrderValidated", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssertionsOrderValidated)
				if err := _Assertions.contract.UnpackLog(event, "OrderValidated", log); err != nil {
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
func (_Assertions *AssertionsFilterer) ParseOrderValidated(log types.Log) (*AssertionsOrderValidated, error) {
	event := new(AssertionsOrderValidated)
	if err := _Assertions.contract.UnpackLog(event, "OrderValidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
