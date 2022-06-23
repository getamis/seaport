// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Conduit

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

// ConduitMetaData contains all meta data concerning the Conduit contract.
var ConduitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BadReturnValueFromERC20OnTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChannelClosed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"identifiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"ERC1155BatchTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Invalid1155BatchTransferEncoding\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidController\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidERC721TransferAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidItemType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingItemAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NoContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenTransferGenericFailure\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"channel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"open\",\"type\":\"bool\"}],\"name\":\"ChannelUpdated\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumConduitItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structConduitTransfer[]\",\"name\":\"transfers\",\"type\":\"tuple[]\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structConduitBatch1155Transfer[]\",\"name\":\"batchTransfers\",\"type\":\"tuple[]\"}],\"name\":\"executeBatch1155\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumConduitItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structConduitTransfer[]\",\"name\":\"standardTransfers\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structConduitBatch1155Transfer[]\",\"name\":\"batchTransfers\",\"type\":\"tuple[]\"}],\"name\":\"executeWithBatch1155\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"channel\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"updateChannel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5033608052608051610a8261003060003960006102100152610a826000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80634ce34aa214610051578063899e104c146100815780638df25d9214610094578063c4e8fcb5146100a7575b600080fd5b61006461005f36600461085c565b6100bc565b6040516001600160e01b0319909116815260200160405180910390f35b61006461008f3660046108e3565b610136565b6100646100a236600461094f565b6101bc565b6100ba6100b53660046109a1565b610205565b005b3360009081526020819052604081205460ff166100ec57604051636821b7df60e01b815260040160405180910390fd5b8160005b81811015610125573685858381811061010b5761010b6109dd565b905060c00201905061011c816102af565b506001016100f0565b50632671a55160e11b949350505050565b3360009081526020819052604081205460ff1661016657604051636821b7df60e01b815260040160405180910390fd5b8360005b8181101561019f5736878783818110610185576101856109dd565b905060c002019050610196816102af565b5060010161016a565b506101aa848461041b565b50632267841360e21b95945050505050565b3360009081526020819052604081205460ff166101ec57604051636821b7df60e01b815260040160405180910390fd5b6101f6838361041b565b506346f92ec960e11b92915050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461024e576040516336abb4df60e11b815260040160405180910390fd5b6001600160a01b03821660008181526020818152604091829020805460ff19168515159081179091558251938452908301527fae63067d43ac07563b7eb8db6595635fc77f1578a2a5ea06ba91b63e2afa37e2910160405180910390a15050565b60016102be6020830183610a09565b60038111156102cf576102cf6109f3565b03610314576103116102e76040830160208401610a31565b6102f76060840160408501610a31565b6103076080850160608601610a31565b8460a00135610565565b50565b60026103236020830183610a09565b6003811115610334576103346109f3565b0361039b578060a0013560011461035e5760405163efcc00b160e01b815260040160405180910390fd5b6103116103716040830160208401610a31565b6103816060840160408501610a31565b6103916080850160608601610a31565b846080013561066b565b60036103aa6020830183610a09565b60038111156103bb576103bb6109f3565b03610402576103116103d36040830160208401610a31565b6103e36060840160408501610a31565b6103f36080850160608601610a31565b84608001358560a0013561072c565b604051631e4cbc7f60e21b815260040160405180910390fd5b808280631759616b60e11b60205260005b83811015610558578235820160208401935060806020820160243760a0810135604081026040018060a00160a452600081610104015260408202610104018160a0850160c4376020830260c00191508160808501351460a0606086013514168285013584141615925082156104ac57633ae8821360e21b60005260046000fd5b923592833b6104ca57632f8aeb3960e11b6000528360045260246000fd5b6000808260206000885af1925082610549573d156105245760203d04915060208104826003028184111561050c57818403600302610200838002868002030401015b5a602082011015610521573d6000803e3d6000fd5b50505b6357e222f160e11b600052836004526080604452608451602001608452602481016000fd5b5050505060018101905061042c565b5050505060806040525050565b6040516323b872dd60e01b600052836004528260245281604452602060006064600080895af1803d15601f3d116001600051141617163d1515811661065b5780873b15151661065b57806106465781610625573d156105ff5760203d046020840481600302818311156105e657818303600302610200838002858002030401015b5a6020820110156105fb573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005286600452856024528460445260006064528360845260a46000fd5b639889192360e01b6000528660045285602452846044528360645260846000fd5b632f8aeb3960e11b6000528660045260246000fd5b5050604052505060006060525050565b833b61068657632f8aeb3960e11b6000528360045260246000fd5b6040516323b872dd60e01b6000528360045282602452816044526000806064600080895af18061071d573d156106f75760203d046020830481600302818311156106de57818303600302610200838002858002030401015b5a6020820110156106f3573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005285600452846024528360445282606452600160845260a46000fd5b50604052505060006060525050565b843b61074757632f8aeb3960e11b6000528460045260246000fd5b60405160805160a05160c051637921219560e11b6000528760045286602452856044528460645260a0608452600060a45260008060c46000808d5af1806107f4573d156107cf5760203d046020860481600302818311156107b657818303600302610200838002858002030401015b5a6020820110156107cb573d6000803e3d6000fd5b5050505b63f486bc8760e01b600052896004528860245287604452866064528560845260a46000fd5b5060809290925260a05260c05260405250506000606052505050565b60008083601f84011261082257600080fd5b50813567ffffffffffffffff81111561083a57600080fd5b60208301915083602060c08302850101111561085557600080fd5b9250929050565b6000806020838503121561086f57600080fd5b823567ffffffffffffffff81111561088657600080fd5b61089285828601610810565b90969095509350505050565b60008083601f8401126108b057600080fd5b50813567ffffffffffffffff8111156108c857600080fd5b6020830191508360208260051b850101111561085557600080fd5b600080600080604085870312156108f957600080fd5b843567ffffffffffffffff8082111561091157600080fd5b61091d88838901610810565b9096509450602087013591508082111561093657600080fd5b506109438782880161089e565b95989497509550505050565b6000806020838503121561096257600080fd5b823567ffffffffffffffff81111561097957600080fd5b6108928582860161089e565b80356001600160a01b038116811461099c57600080fd5b919050565b600080604083850312156109b457600080fd5b6109bd83610985565b9150602083013580151581146109d257600080fd5b809150509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b600060208284031215610a1b57600080fd5b813560048110610a2a57600080fd5b9392505050565b600060208284031215610a4357600080fd5b610a2a8261098556fea2646970667358221220a42ce1840bb9e65dd7de3905229b790ddc477c9436d75acc17f58f88068bdf9764736f6c634300080d0033",
}

// ConduitABI is the input ABI used to generate the binding from.
// Deprecated: Use ConduitMetaData.ABI instead.
var ConduitABI = ConduitMetaData.ABI

// ConduitBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConduitMetaData.Bin instead.
var ConduitBin = ConduitMetaData.Bin

// DeployConduit deploys a new Ethereum contract, binding an instance of Conduit to it.
func DeployConduit(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Conduit, error) {
	parsed, err := ConduitMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConduitBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Conduit{ConduitCaller: ConduitCaller{contract: contract}, ConduitTransactor: ConduitTransactor{contract: contract}, ConduitFilterer: ConduitFilterer{contract: contract}}, nil
}

// Conduit is an auto generated Go binding around an Ethereum contract.
type Conduit struct {
	ConduitCaller     // Read-only binding to the contract
	ConduitTransactor // Write-only binding to the contract
	ConduitFilterer   // Log filterer for contract events
}

// ConduitCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConduitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConduitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConduitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConduitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConduitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConduitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConduitSession struct {
	Contract     *Conduit          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConduitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConduitCallerSession struct {
	Contract *ConduitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ConduitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConduitTransactorSession struct {
	Contract     *ConduitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ConduitRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConduitRaw struct {
	Contract *Conduit // Generic contract binding to access the raw methods on
}

// ConduitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConduitCallerRaw struct {
	Contract *ConduitCaller // Generic read-only contract binding to access the raw methods on
}

// ConduitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConduitTransactorRaw struct {
	Contract *ConduitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConduit creates a new instance of Conduit, bound to a specific deployed contract.
func NewConduit(address common.Address, backend bind.ContractBackend) (*Conduit, error) {
	contract, err := bindConduit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Conduit{ConduitCaller: ConduitCaller{contract: contract}, ConduitTransactor: ConduitTransactor{contract: contract}, ConduitFilterer: ConduitFilterer{contract: contract}}, nil
}

// NewConduitCaller creates a new read-only instance of Conduit, bound to a specific deployed contract.
func NewConduitCaller(address common.Address, caller bind.ContractCaller) (*ConduitCaller, error) {
	contract, err := bindConduit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConduitCaller{contract: contract}, nil
}

// NewConduitTransactor creates a new write-only instance of Conduit, bound to a specific deployed contract.
func NewConduitTransactor(address common.Address, transactor bind.ContractTransactor) (*ConduitTransactor, error) {
	contract, err := bindConduit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConduitTransactor{contract: contract}, nil
}

// NewConduitFilterer creates a new log filterer instance of Conduit, bound to a specific deployed contract.
func NewConduitFilterer(address common.Address, filterer bind.ContractFilterer) (*ConduitFilterer, error) {
	contract, err := bindConduit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConduitFilterer{contract: contract}, nil
}

// bindConduit binds a generic wrapper to an already deployed contract.
func bindConduit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConduitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Conduit *ConduitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Conduit.Contract.ConduitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Conduit *ConduitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Conduit.Contract.ConduitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Conduit *ConduitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Conduit.Contract.ConduitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Conduit *ConduitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Conduit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Conduit *ConduitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Conduit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Conduit *ConduitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Conduit.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0x4ce34aa2.
//
// Solidity: function execute((uint8,address,address,address,uint256,uint256)[] transfers) returns(bytes4 magicValue)
func (_Conduit *ConduitTransactor) Execute(opts *bind.TransactOpts, transfers []ConduitTransfer) (*types.Transaction, error) {
	return _Conduit.contract.Transact(opts, "execute", transfers)
}

// Execute is a paid mutator transaction binding the contract method 0x4ce34aa2.
//
// Solidity: function execute((uint8,address,address,address,uint256,uint256)[] transfers) returns(bytes4 magicValue)
func (_Conduit *ConduitSession) Execute(transfers []ConduitTransfer) (*types.Transaction, error) {
	return _Conduit.Contract.Execute(&_Conduit.TransactOpts, transfers)
}

// Execute is a paid mutator transaction binding the contract method 0x4ce34aa2.
//
// Solidity: function execute((uint8,address,address,address,uint256,uint256)[] transfers) returns(bytes4 magicValue)
func (_Conduit *ConduitTransactorSession) Execute(transfers []ConduitTransfer) (*types.Transaction, error) {
	return _Conduit.Contract.Execute(&_Conduit.TransactOpts, transfers)
}

// ExecuteBatch1155 is a paid mutator transaction binding the contract method 0x8df25d92.
//
// Solidity: function executeBatch1155((address,address,address,uint256[],uint256[])[] batchTransfers) returns(bytes4 magicValue)
func (_Conduit *ConduitTransactor) ExecuteBatch1155(opts *bind.TransactOpts, batchTransfers []ConduitBatch1155Transfer) (*types.Transaction, error) {
	return _Conduit.contract.Transact(opts, "executeBatch1155", batchTransfers)
}

// ExecuteBatch1155 is a paid mutator transaction binding the contract method 0x8df25d92.
//
// Solidity: function executeBatch1155((address,address,address,uint256[],uint256[])[] batchTransfers) returns(bytes4 magicValue)
func (_Conduit *ConduitSession) ExecuteBatch1155(batchTransfers []ConduitBatch1155Transfer) (*types.Transaction, error) {
	return _Conduit.Contract.ExecuteBatch1155(&_Conduit.TransactOpts, batchTransfers)
}

// ExecuteBatch1155 is a paid mutator transaction binding the contract method 0x8df25d92.
//
// Solidity: function executeBatch1155((address,address,address,uint256[],uint256[])[] batchTransfers) returns(bytes4 magicValue)
func (_Conduit *ConduitTransactorSession) ExecuteBatch1155(batchTransfers []ConduitBatch1155Transfer) (*types.Transaction, error) {
	return _Conduit.Contract.ExecuteBatch1155(&_Conduit.TransactOpts, batchTransfers)
}

// ExecuteWithBatch1155 is a paid mutator transaction binding the contract method 0x899e104c.
//
// Solidity: function executeWithBatch1155((uint8,address,address,address,uint256,uint256)[] standardTransfers, (address,address,address,uint256[],uint256[])[] batchTransfers) returns(bytes4 magicValue)
func (_Conduit *ConduitTransactor) ExecuteWithBatch1155(opts *bind.TransactOpts, standardTransfers []ConduitTransfer, batchTransfers []ConduitBatch1155Transfer) (*types.Transaction, error) {
	return _Conduit.contract.Transact(opts, "executeWithBatch1155", standardTransfers, batchTransfers)
}

// ExecuteWithBatch1155 is a paid mutator transaction binding the contract method 0x899e104c.
//
// Solidity: function executeWithBatch1155((uint8,address,address,address,uint256,uint256)[] standardTransfers, (address,address,address,uint256[],uint256[])[] batchTransfers) returns(bytes4 magicValue)
func (_Conduit *ConduitSession) ExecuteWithBatch1155(standardTransfers []ConduitTransfer, batchTransfers []ConduitBatch1155Transfer) (*types.Transaction, error) {
	return _Conduit.Contract.ExecuteWithBatch1155(&_Conduit.TransactOpts, standardTransfers, batchTransfers)
}

// ExecuteWithBatch1155 is a paid mutator transaction binding the contract method 0x899e104c.
//
// Solidity: function executeWithBatch1155((uint8,address,address,address,uint256,uint256)[] standardTransfers, (address,address,address,uint256[],uint256[])[] batchTransfers) returns(bytes4 magicValue)
func (_Conduit *ConduitTransactorSession) ExecuteWithBatch1155(standardTransfers []ConduitTransfer, batchTransfers []ConduitBatch1155Transfer) (*types.Transaction, error) {
	return _Conduit.Contract.ExecuteWithBatch1155(&_Conduit.TransactOpts, standardTransfers, batchTransfers)
}

// UpdateChannel is a paid mutator transaction binding the contract method 0xc4e8fcb5.
//
// Solidity: function updateChannel(address channel, bool isOpen) returns()
func (_Conduit *ConduitTransactor) UpdateChannel(opts *bind.TransactOpts, channel common.Address, isOpen bool) (*types.Transaction, error) {
	return _Conduit.contract.Transact(opts, "updateChannel", channel, isOpen)
}

// UpdateChannel is a paid mutator transaction binding the contract method 0xc4e8fcb5.
//
// Solidity: function updateChannel(address channel, bool isOpen) returns()
func (_Conduit *ConduitSession) UpdateChannel(channel common.Address, isOpen bool) (*types.Transaction, error) {
	return _Conduit.Contract.UpdateChannel(&_Conduit.TransactOpts, channel, isOpen)
}

// UpdateChannel is a paid mutator transaction binding the contract method 0xc4e8fcb5.
//
// Solidity: function updateChannel(address channel, bool isOpen) returns()
func (_Conduit *ConduitTransactorSession) UpdateChannel(channel common.Address, isOpen bool) (*types.Transaction, error) {
	return _Conduit.Contract.UpdateChannel(&_Conduit.TransactOpts, channel, isOpen)
}

// ConduitChannelUpdatedIterator is returned from FilterChannelUpdated and is used to iterate over the raw logs and unpacked data for ChannelUpdated events raised by the Conduit contract.
type ConduitChannelUpdatedIterator struct {
	Event *ConduitChannelUpdated // Event containing the contract specifics and raw log

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
func (it *ConduitChannelUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConduitChannelUpdated)
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
		it.Event = new(ConduitChannelUpdated)
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
func (it *ConduitChannelUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConduitChannelUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConduitChannelUpdated represents a ChannelUpdated event raised by the Conduit contract.
type ConduitChannelUpdated struct {
	Channel common.Address
	Open    bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChannelUpdated is a free log retrieval operation binding the contract event 0xae63067d43ac07563b7eb8db6595635fc77f1578a2a5ea06ba91b63e2afa37e2.
//
// Solidity: event ChannelUpdated(address channel, bool open)
func (_Conduit *ConduitFilterer) FilterChannelUpdated(opts *bind.FilterOpts) (*ConduitChannelUpdatedIterator, error) {

	logs, sub, err := _Conduit.contract.FilterLogs(opts, "ChannelUpdated")
	if err != nil {
		return nil, err
	}
	return &ConduitChannelUpdatedIterator{contract: _Conduit.contract, event: "ChannelUpdated", logs: logs, sub: sub}, nil
}

// WatchChannelUpdated is a free log subscription operation binding the contract event 0xae63067d43ac07563b7eb8db6595635fc77f1578a2a5ea06ba91b63e2afa37e2.
//
// Solidity: event ChannelUpdated(address channel, bool open)
func (_Conduit *ConduitFilterer) WatchChannelUpdated(opts *bind.WatchOpts, sink chan<- *ConduitChannelUpdated) (event.Subscription, error) {

	logs, sub, err := _Conduit.contract.WatchLogs(opts, "ChannelUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConduitChannelUpdated)
				if err := _Conduit.contract.UnpackLog(event, "ChannelUpdated", log); err != nil {
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
func (_Conduit *ConduitFilterer) ParseChannelUpdated(log types.Log) (*ConduitChannelUpdated, error) {
	event := new(ConduitChannelUpdated)
	if err := _Conduit.contract.UnpackLog(event, "ChannelUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
