// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TransferHelper

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

// TransferHelperItem is an auto generated low-level Go binding around an user-defined struct.
type TransferHelperItem struct {
	ItemType   uint8
	Token      common.Address
	Identifier *big.Int
	Amount     *big.Int
}

// TransferHelperMetaData contains all meta data concerning the TransferHelper contract.
var TransferHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduitController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BadReturnValueFromERC20OnTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"identifiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"ERC1155BatchTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Invalid1155BatchTransferEncoding\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidERC721TransferAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidItemType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingItemAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NoContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnusedItemParameters\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumConduitItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTransferHelperItem[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"name\":\"bulkTransfer\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b50604051610a15380380610a1583398101604081905261002f916100ab565b6000819050806001600160a01b0316630a96ad396040518163ffffffff1660e01b81526004016040805180830381865afa158015610071573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061009591906100db565b5060a0526001600160a01b0316608052506100ff565b6000602082840312156100bd57600080fd5b81516001600160a01b03811681146100d457600080fd5b9392505050565b600080604083850312156100ee57600080fd5b505080516020909101519092909150565b60805160a0516108f161012460003960006101e1015260006101b001526108f16000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806373b69d0314610030575b600080fd5b61004361003e3660046106cb565b610060565b6040516001600160e01b0319909116815260200160405180910390f35b6000838261018f5760005b81811015610189573687878381811061008657610086610757565b608002919091019150600090506100a06020830183610783565b60038111156100b1576100b161076d565b036100cf57604051631e4cbc7f60e21b815260040160405180910390fd5b60016100de6020830183610783565b60038111156100ef576100ef61076d565b036101185761011361010760408301602084016107ab565b338884606001356103fb565b610180565b60026101276020830183610783565b60038111156101385761013861076d565b0361015c5761011361015060408301602084016107ab565b33888460400135610504565b61018061016f60408301602084016107ab565b3388846040013585606001356105c8565b5060010161006b565b506103e9565b6040516001600160f81b031960208201526bffffffffffffffffffffffff197f000000000000000000000000000000000000000000000000000000000000000060601b166021820152603581018490527f000000000000000000000000000000000000000000000000000000000000000060558201526000906075016040516020818303038152906040528051906020012060001c905060008267ffffffffffffffff811115610241576102416107c6565b6040519080825280602002602001820160405280156102a157816020015b6040805160c08101825260008082526020808301829052928201819052606082018190526080820181905260a0820152825260001990920191018161025f5790505b50905060005b8381101561037557368989838181106102c2576102c2610757565b90506080020190506040518060c001604052808260000160208101906102e89190610783565b60038111156102f9576102f961076d565b815260200182602001602081019061031191906107ab565b6001600160a01b03168152602001336001600160a01b03168152602001896001600160a01b0316815260200182604001358152602001826060013581525083838151811061036157610361610757565b6020908102919091010152506001016102a7565b50604051632671a55160e11b81526001600160a01b03831690634ce34aa2906103a29084906004016107dc565b6020604051808303816000875af11580156103c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103e59190610891565b5050505b506373b69d0360e01b95945050505050565b6040516323b872dd60e01b600052836004528260245281604452602060006064600080895af1803d15601f3d116001600051141617163d151581166104f45780873b1515166104f457806104df57816104be573d15610498576020601f3d010460208404816003028183111561047f57818303600302610200838002858002030401015b5a602082011015610494573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005286600452856024528460445260006064528360845260a46000fd5b639889192360e01b6000528660045285602452846044528360645260846000fd5b632f8aeb3960e11b6000528660045260246000fd5b5050604052505060006060525050565b833b61051f57632f8aeb3960e11b6000528360045260246000fd5b6040516323b872dd60e01b6000528360045282602452816044526000806064600080895af1806105b9573d15610593576020601f3d010460208304816003028183111561057a57818303600302610200838002858002030401015b5a60208201101561058f573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005285600452846024528360445282606452600160845260a46000fd5b50604052505060006060525050565b843b6105e357632f8aeb3960e11b6000528460045260246000fd5b60405160805160a05160c051637921219560e11b6000528760045286602452856044528460645260a0608452600060a45260008060c46000808d5af180610693573d1561066e576020601f3d010460208604816003028183111561065557818303600302610200838002858002030401015b5a60208201101561066a573d6000803e3d6000fd5b5050505b63f486bc8760e01b600052896004528860245287604452866064528560845260a46000fd5b5060809290925260a05260c05260405250506000606052505050565b80356001600160a01b03811681146106c657600080fd5b919050565b600080600080606085870312156106e157600080fd5b843567ffffffffffffffff808211156106f957600080fd5b818701915087601f83011261070d57600080fd5b81358181111561071c57600080fd5b8860208260071b850101111561073157600080fd5b60209283019650945061074791870190506106af565b9396929550929360400135925050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b60006020828403121561079557600080fd5b8135600481106107a457600080fd5b9392505050565b6000602082840312156107bd57600080fd5b6107a4826106af565b634e487b7160e01b600052604160045260246000fd5b60208082528251828201819052600091906040908185019086840185805b838110156108835782518051600480821061082257634e487b7160e01b855260218152602485fd5b508652808801516001600160a01b039081168988015287820151168787015260608082015161085b828901826001600160a01b03169052565b50506080818101519087015260a0908101519086015260c090940193918601916001016107fa565b509298975050505050505050565b6000602082840312156108a357600080fd5b81516001600160e01b0319811681146107a457600080fdfea264697066735822122064b3e65306ea155b53fd1ad94a15c6369898bd51ebe2e97e87632715c5dacb8c64736f6c634300080d0033",
}

// TransferHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferHelperMetaData.ABI instead.
var TransferHelperABI = TransferHelperMetaData.ABI

// TransferHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransferHelperMetaData.Bin instead.
var TransferHelperBin = TransferHelperMetaData.Bin

// DeployTransferHelper deploys a new Ethereum contract, binding an instance of TransferHelper to it.
func DeployTransferHelper(auth *bind.TransactOpts, backend bind.ContractBackend, conduitController common.Address) (common.Address, *types.Transaction, *TransferHelper, error) {
	parsed, err := TransferHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransferHelperBin), backend, conduitController)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// TransferHelper is an auto generated Go binding around an Ethereum contract.
type TransferHelper struct {
	TransferHelperCaller     // Read-only binding to the contract
	TransferHelperTransactor // Write-only binding to the contract
	TransferHelperFilterer   // Log filterer for contract events
}

// TransferHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferHelperSession struct {
	Contract     *TransferHelper   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferHelperCallerSession struct {
	Contract *TransferHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TransferHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferHelperTransactorSession struct {
	Contract     *TransferHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TransferHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferHelperRaw struct {
	Contract *TransferHelper // Generic contract binding to access the raw methods on
}

// TransferHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferHelperCallerRaw struct {
	Contract *TransferHelperCaller // Generic read-only contract binding to access the raw methods on
}

// TransferHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferHelperTransactorRaw struct {
	Contract *TransferHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferHelper creates a new instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelper(address common.Address, backend bind.ContractBackend) (*TransferHelper, error) {
	contract, err := bindTransferHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// NewTransferHelperCaller creates a new read-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperCaller(address common.Address, caller bind.ContractCaller) (*TransferHelperCaller, error) {
	contract, err := bindTransferHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperCaller{contract: contract}, nil
}

// NewTransferHelperTransactor creates a new write-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferHelperTransactor, error) {
	contract, err := bindTransferHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperTransactor{contract: contract}, nil
}

// NewTransferHelperFilterer creates a new log filterer instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferHelperFilterer, error) {
	contract, err := bindTransferHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferHelperFilterer{contract: contract}, nil
}

// bindTransferHelper binds a generic wrapper to an already deployed contract.
func bindTransferHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.TransferHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transact(opts, method, params...)
}

// BulkTransfer is a paid mutator transaction binding the contract method 0x73b69d03.
//
// Solidity: function bulkTransfer((uint8,address,uint256,uint256)[] items, address recipient, bytes32 conduitKey) returns(bytes4 magicValue)
func (_TransferHelper *TransferHelperTransactor) BulkTransfer(opts *bind.TransactOpts, items []TransferHelperItem, recipient common.Address, conduitKey [32]byte) (*types.Transaction, error) {
	return _TransferHelper.contract.Transact(opts, "bulkTransfer", items, recipient, conduitKey)
}

// BulkTransfer is a paid mutator transaction binding the contract method 0x73b69d03.
//
// Solidity: function bulkTransfer((uint8,address,uint256,uint256)[] items, address recipient, bytes32 conduitKey) returns(bytes4 magicValue)
func (_TransferHelper *TransferHelperSession) BulkTransfer(items []TransferHelperItem, recipient common.Address, conduitKey [32]byte) (*types.Transaction, error) {
	return _TransferHelper.Contract.BulkTransfer(&_TransferHelper.TransactOpts, items, recipient, conduitKey)
}

// BulkTransfer is a paid mutator transaction binding the contract method 0x73b69d03.
//
// Solidity: function bulkTransfer((uint8,address,uint256,uint256)[] items, address recipient, bytes32 conduitKey) returns(bytes4 magicValue)
func (_TransferHelper *TransferHelperTransactorSession) BulkTransfer(items []TransferHelperItem, recipient common.Address, conduitKey [32]byte) (*types.Transaction, error) {
	return _TransferHelper.Contract.BulkTransfer(&_TransferHelper.TransactOpts, items, recipient, conduitKey)
}
