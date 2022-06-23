// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ZoneInterface

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

// AdvancedOrder is an auto generated low-level Go binding around an user-defined struct.
type AdvancedOrder struct {
	Parameters  OrderParameters
	Numerator   *big.Int
	Denominator *big.Int
	Signature   []byte
	ExtraData   []byte
}

// ConsiderationItem is an auto generated low-level Go binding around an user-defined struct.
type ConsiderationItem struct {
	ItemType             uint8
	Token                common.Address
	IdentifierOrCriteria *big.Int
	StartAmount          *big.Int
	EndAmount            *big.Int
	Recipient            common.Address
}

// CriteriaResolver is an auto generated low-level Go binding around an user-defined struct.
type CriteriaResolver struct {
	OrderIndex    *big.Int
	Side          uint8
	Index         *big.Int
	Identifier    *big.Int
	CriteriaProof [][32]byte
}

// OfferItem is an auto generated low-level Go binding around an user-defined struct.
type OfferItem struct {
	ItemType             uint8
	Token                common.Address
	IdentifierOrCriteria *big.Int
	StartAmount          *big.Int
	EndAmount            *big.Int
}

// OrderParameters is an auto generated low-level Go binding around an user-defined struct.
type OrderParameters struct {
	Offerer                         common.Address
	Zone                            common.Address
	Offer                           []OfferItem
	Consideration                   []ConsiderationItem
	OrderType                       uint8
	StartTime                       *big.Int
	EndTime                         *big.Int
	ZoneHash                        [32]byte
	Salt                            *big.Int
	ConduitKey                      [32]byte
	TotalOriginalConsiderationItems *big.Int
}

// ZoneInterfaceMetaData contains all meta data concerning the ZoneInterface contract.
var ZoneInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"}],\"name\":\"isValidOrder\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"validOrderMagicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"priorOrderHashes\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"criteriaResolvers\",\"type\":\"tuple[]\"}],\"name\":\"isValidOrderIncludingExtraData\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"validOrderMagicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ZoneInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use ZoneInterfaceMetaData.ABI instead.
var ZoneInterfaceABI = ZoneInterfaceMetaData.ABI

// ZoneInterface is an auto generated Go binding around an Ethereum contract.
type ZoneInterface struct {
	ZoneInterfaceCaller     // Read-only binding to the contract
	ZoneInterfaceTransactor // Write-only binding to the contract
	ZoneInterfaceFilterer   // Log filterer for contract events
}

// ZoneInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZoneInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZoneInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZoneInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZoneInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZoneInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZoneInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZoneInterfaceSession struct {
	Contract     *ZoneInterface    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZoneInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZoneInterfaceCallerSession struct {
	Contract *ZoneInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ZoneInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZoneInterfaceTransactorSession struct {
	Contract     *ZoneInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ZoneInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZoneInterfaceRaw struct {
	Contract *ZoneInterface // Generic contract binding to access the raw methods on
}

// ZoneInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZoneInterfaceCallerRaw struct {
	Contract *ZoneInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ZoneInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZoneInterfaceTransactorRaw struct {
	Contract *ZoneInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZoneInterface creates a new instance of ZoneInterface, bound to a specific deployed contract.
func NewZoneInterface(address common.Address, backend bind.ContractBackend) (*ZoneInterface, error) {
	contract, err := bindZoneInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZoneInterface{ZoneInterfaceCaller: ZoneInterfaceCaller{contract: contract}, ZoneInterfaceTransactor: ZoneInterfaceTransactor{contract: contract}, ZoneInterfaceFilterer: ZoneInterfaceFilterer{contract: contract}}, nil
}

// NewZoneInterfaceCaller creates a new read-only instance of ZoneInterface, bound to a specific deployed contract.
func NewZoneInterfaceCaller(address common.Address, caller bind.ContractCaller) (*ZoneInterfaceCaller, error) {
	contract, err := bindZoneInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZoneInterfaceCaller{contract: contract}, nil
}

// NewZoneInterfaceTransactor creates a new write-only instance of ZoneInterface, bound to a specific deployed contract.
func NewZoneInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ZoneInterfaceTransactor, error) {
	contract, err := bindZoneInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZoneInterfaceTransactor{contract: contract}, nil
}

// NewZoneInterfaceFilterer creates a new log filterer instance of ZoneInterface, bound to a specific deployed contract.
func NewZoneInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ZoneInterfaceFilterer, error) {
	contract, err := bindZoneInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZoneInterfaceFilterer{contract: contract}, nil
}

// bindZoneInterface binds a generic wrapper to an already deployed contract.
func bindZoneInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZoneInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZoneInterface *ZoneInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZoneInterface.Contract.ZoneInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZoneInterface *ZoneInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZoneInterface.Contract.ZoneInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZoneInterface *ZoneInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZoneInterface.Contract.ZoneInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZoneInterface *ZoneInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZoneInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZoneInterface *ZoneInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZoneInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZoneInterface *ZoneInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZoneInterface.Contract.contract.Transact(opts, method, params...)
}

// IsValidOrder is a free data retrieval call binding the contract method 0x0e1d31dc.
//
// Solidity: function isValidOrder(bytes32 orderHash, address caller, address offerer, bytes32 zoneHash) view returns(bytes4 validOrderMagicValue)
func (_ZoneInterface *ZoneInterfaceCaller) IsValidOrder(opts *bind.CallOpts, orderHash [32]byte, caller common.Address, offerer common.Address, zoneHash [32]byte) ([4]byte, error) {
	var out []interface{}
	err := _ZoneInterface.contract.Call(opts, &out, "isValidOrder", orderHash, caller, offerer, zoneHash)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidOrder is a free data retrieval call binding the contract method 0x0e1d31dc.
//
// Solidity: function isValidOrder(bytes32 orderHash, address caller, address offerer, bytes32 zoneHash) view returns(bytes4 validOrderMagicValue)
func (_ZoneInterface *ZoneInterfaceSession) IsValidOrder(orderHash [32]byte, caller common.Address, offerer common.Address, zoneHash [32]byte) ([4]byte, error) {
	return _ZoneInterface.Contract.IsValidOrder(&_ZoneInterface.CallOpts, orderHash, caller, offerer, zoneHash)
}

// IsValidOrder is a free data retrieval call binding the contract method 0x0e1d31dc.
//
// Solidity: function isValidOrder(bytes32 orderHash, address caller, address offerer, bytes32 zoneHash) view returns(bytes4 validOrderMagicValue)
func (_ZoneInterface *ZoneInterfaceCallerSession) IsValidOrder(orderHash [32]byte, caller common.Address, offerer common.Address, zoneHash [32]byte) ([4]byte, error) {
	return _ZoneInterface.Contract.IsValidOrder(&_ZoneInterface.CallOpts, orderHash, caller, offerer, zoneHash)
}

// IsValidOrderIncludingExtraData is a free data retrieval call binding the contract method 0x33131570.
//
// Solidity: function isValidOrderIncludingExtraData(bytes32 orderHash, address caller, ((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) order, bytes32[] priorOrderHashes, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers) view returns(bytes4 validOrderMagicValue)
func (_ZoneInterface *ZoneInterfaceCaller) IsValidOrderIncludingExtraData(opts *bind.CallOpts, orderHash [32]byte, caller common.Address, order AdvancedOrder, priorOrderHashes [][32]byte, criteriaResolvers []CriteriaResolver) ([4]byte, error) {
	var out []interface{}
	err := _ZoneInterface.contract.Call(opts, &out, "isValidOrderIncludingExtraData", orderHash, caller, order, priorOrderHashes, criteriaResolvers)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidOrderIncludingExtraData is a free data retrieval call binding the contract method 0x33131570.
//
// Solidity: function isValidOrderIncludingExtraData(bytes32 orderHash, address caller, ((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) order, bytes32[] priorOrderHashes, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers) view returns(bytes4 validOrderMagicValue)
func (_ZoneInterface *ZoneInterfaceSession) IsValidOrderIncludingExtraData(orderHash [32]byte, caller common.Address, order AdvancedOrder, priorOrderHashes [][32]byte, criteriaResolvers []CriteriaResolver) ([4]byte, error) {
	return _ZoneInterface.Contract.IsValidOrderIncludingExtraData(&_ZoneInterface.CallOpts, orderHash, caller, order, priorOrderHashes, criteriaResolvers)
}

// IsValidOrderIncludingExtraData is a free data retrieval call binding the contract method 0x33131570.
//
// Solidity: function isValidOrderIncludingExtraData(bytes32 orderHash, address caller, ((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) order, bytes32[] priorOrderHashes, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers) view returns(bytes4 validOrderMagicValue)
func (_ZoneInterface *ZoneInterfaceCallerSession) IsValidOrderIncludingExtraData(orderHash [32]byte, caller common.Address, order AdvancedOrder, priorOrderHashes [][32]byte, criteriaResolvers []CriteriaResolver) ([4]byte, error) {
	return _ZoneInterface.Contract.IsValidOrderIncludingExtraData(&_ZoneInterface.CallOpts, orderHash, caller, order, priorOrderHashes, criteriaResolvers)
}
