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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BadReturnValueFromERC20OnTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"channel\",\"type\":\"address\"}],\"name\":\"ChannelClosed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"channel\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"ChannelStatusAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"identifiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"ERC1155BatchTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Invalid1155BatchTransferEncoding\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidController\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidERC721TransferAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidItemType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingItemAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NoContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnusedItemParameters\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"channel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"open\",\"type\":\"bool\"}],\"name\":\"ChannelUpdated\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumConduitItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structConduitTransfer[]\",\"name\":\"transfers\",\"type\":\"tuple[]\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structConduitBatch1155Transfer[]\",\"name\":\"batchTransfers\",\"type\":\"tuple[]\"}],\"name\":\"executeBatch1155\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumConduitItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structConduitTransfer[]\",\"name\":\"standardTransfers\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structConduitBatch1155Transfer[]\",\"name\":\"batchTransfers\",\"type\":\"tuple[]\"}],\"name\":\"executeWithBatch1155\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"channel\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"updateChannel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5033608052608051610ab361003060003960006101e90152610ab36000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80634ce34aa214610051578063899e104c146100815780638df25d9214610094578063c4e8fcb5146100a7575b600080fd5b61006461005f36600461088d565b6100bc565b6040516001600160e01b0319909116815260200160405180910390f35b61006461008f366004610914565b61012b565b6100646100a2366004610980565b61019b565b6100ba6100b53660046109d2565b6101de565b005b60003360005260006020526040600020546100e6576349ed56f960e11b6000523360045260246000fd5b8160005b8181101561011a5761011285858381811061010757610107610a0e565b905060c002016102dc565b6001016100ea565b50632671a55160e11b949350505050565b6000336000526000602052604060002054610155576349ed56f960e11b6000523360045260246000fd5b8360005b8181101561017e5761017687878381811061010757610107610a0e565b600101610159565b506101898484610448565b50632267841360e21b95945050505050565b60003360005260006020526040600020546101c5576349ed56f960e11b6000523360045260246000fd5b6101cf8383610448565b506346f92ec960e11b92915050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610227576040516336abb4df60e11b815260040160405180910390fd5b6001600160a01b03821660009081526020819052604090205481151560ff90911615150361027f576040516349271a0f60e11b81526001600160a01b0383166004820152811515602482015260440160405180910390fd5b6001600160a01b03821660008181526020818152604091829020805460ff191685151590811790915591519182527fae63067d43ac07563b7eb8db6595635fc77f1578a2a5ea06ba91b63e2afa37e2910160405180910390a25050565b60016102eb6020830183610a3a565b60038111156102fc576102fc610a24565b036103415761033e6103146040830160208401610a62565b6103246060840160408501610a62565b6103346080850160608601610a62565b8460a0013561058d565b50565b60026103506020830183610a3a565b600381111561036157610361610a24565b036103c8578060a0013560011461038b5760405163efcc00b160e01b815260040160405180910390fd5b61033e61039e6040830160208401610a62565b6103ae6060840160408501610a62565b6103be6080850160608601610a62565b8460800135610696565b60036103d76020830183610a3a565b60038111156103e8576103e8610a24565b0361042f5761033e6104006040830160208401610a62565b6104106060840160408501610a62565b6104206080850160608601610a62565b84608001358560a0013561075a565b604051631e4cbc7f60e21b815260040160405180910390fd5b808280631759616b60e11b60205260005b8381101561058057823582018035803b61048257632f8aeb3960e11b6000528060045260246000fd5b60a08201356020810260c0018060808501351460a0606086013514168185013583141615905080156104bf57633ae8821360e21b60005260046000fd5b506020860195506080602084016024376040810260400190508060a00160a45260008160c401528060c4018160a0850160c4376000808260206000875af1935083610571573d1561054f576020601f3d0104915060208104826003028184111561053757818403600302610200838002868002030401015b5a60208201101561054c573d6000803e3d6000fd5b50505b6357e222f160e11b6000528260045260c0606452608451602001608452806000fd5b50505050600181019050610459565b5050505060806040525050565b6040516323b872dd60e01b600052836004528260245281604452602060006064600080895af1803d15601f3d116001600051141617163d151581166106865780873b15151661068657806106715781610650573d1561062a576020601f3d010460208404816003028183111561061157818303600302610200838002858002030401015b5a602082011015610626573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005286600452856024528460445260006064528360845260a46000fd5b639889192360e01b6000528660045285602452846044528360645260846000fd5b632f8aeb3960e11b6000528660045260246000fd5b5050604052505060006060525050565b833b6106b157632f8aeb3960e11b6000528360045260246000fd5b6040516323b872dd60e01b6000528360045282602452816044526000806064600080895af18061074b573d15610725576020601f3d010460208304816003028183111561070c57818303600302610200838002858002030401015b5a602082011015610721573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005285600452846024528360445282606452600160845260a46000fd5b50604052505060006060525050565b843b61077557632f8aeb3960e11b6000528460045260246000fd5b60405160805160a05160c051637921219560e11b6000528760045286602452856044528460645260a0608452600060a45260008060c46000808d5af180610825573d15610800576020601f3d01046020860481600302818311156107e757818303600302610200838002858002030401015b5a6020820110156107fc573d6000803e3d6000fd5b5050505b63f486bc8760e01b600052896004528860245287604452866064528560845260a46000fd5b5060809290925260a05260c05260405250506000606052505050565b60008083601f84011261085357600080fd5b50813567ffffffffffffffff81111561086b57600080fd5b60208301915083602060c08302850101111561088657600080fd5b9250929050565b600080602083850312156108a057600080fd5b823567ffffffffffffffff8111156108b757600080fd5b6108c385828601610841565b90969095509350505050565b60008083601f8401126108e157600080fd5b50813567ffffffffffffffff8111156108f957600080fd5b6020830191508360208260051b850101111561088657600080fd5b6000806000806040858703121561092a57600080fd5b843567ffffffffffffffff8082111561094257600080fd5b61094e88838901610841565b9096509450602087013591508082111561096757600080fd5b50610974878288016108cf565b95989497509550505050565b6000806020838503121561099357600080fd5b823567ffffffffffffffff8111156109aa57600080fd5b6108c3858286016108cf565b80356001600160a01b03811681146109cd57600080fd5b919050565b600080604083850312156109e557600080fd5b6109ee836109b6565b915060208301358015158114610a0357600080fd5b809150509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b600060208284031215610a4c57600080fd5b813560048110610a5b57600080fd5b9392505050565b600060208284031215610a7457600080fd5b610a5b826109b656fea264697066735822122060c22dd101dc32b272f2268d391e21d76db36256cb9c3c5d4b94b6e5bb0dcb7a64736f6c634300080d0033",
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
// Solidity: event ChannelUpdated(address indexed channel, bool open)
func (_Conduit *ConduitFilterer) FilterChannelUpdated(opts *bind.FilterOpts, channel []common.Address) (*ConduitChannelUpdatedIterator, error) {

	var channelRule []interface{}
	for _, channelItem := range channel {
		channelRule = append(channelRule, channelItem)
	}

	logs, sub, err := _Conduit.contract.FilterLogs(opts, "ChannelUpdated", channelRule)
	if err != nil {
		return nil, err
	}
	return &ConduitChannelUpdatedIterator{contract: _Conduit.contract, event: "ChannelUpdated", logs: logs, sub: sub}, nil
}

// WatchChannelUpdated is a free log subscription operation binding the contract event 0xae63067d43ac07563b7eb8db6595635fc77f1578a2a5ea06ba91b63e2afa37e2.
//
// Solidity: event ChannelUpdated(address indexed channel, bool open)
func (_Conduit *ConduitFilterer) WatchChannelUpdated(opts *bind.WatchOpts, sink chan<- *ConduitChannelUpdated, channel []common.Address) (event.Subscription, error) {

	var channelRule []interface{}
	for _, channelItem := range channel {
		channelRule = append(channelRule, channelItem)
	}

	logs, sub, err := _Conduit.contract.WatchLogs(opts, "ChannelUpdated", channelRule)
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
// Solidity: event ChannelUpdated(address indexed channel, bool open)
func (_Conduit *ConduitFilterer) ParseChannelUpdated(log types.Log) (*ConduitChannelUpdated, error) {
	event := new(ConduitChannelUpdated)
	if err := _Conduit.contract.UnpackLog(event, "ChannelUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
