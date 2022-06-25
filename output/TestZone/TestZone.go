// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TestZone

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

// TestZoneMetaData contains all meta data concerning the TestZone contract.
var TestZoneMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"}],\"name\":\"isValidOrder\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"validOrderMagicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"priorOrderHashes\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"criteriaResolvers\",\"type\":\"tuple[]\"}],\"name\":\"isValidOrderIncludingExtraData\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"validOrderMagicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506103b9806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80630e1d31dc1461003b578063331315701461006b575b600080fd5b61004e6100493660046101d4565b61007e565b6040516001600160e01b0319909116815260200160405180910390f35b61004e610079366004610264565b610105565b600060001982016100ce5760405162461bcd60e51b8152602060048201526015602482015274526576657274206f6e207a6f6e652068617368203160581b60448201526064015b60405180910390fd5b60011982016100dc57600080fd5b60021982016100f3576001600160e01b03196100fc565b6303874c7760e21b5b95945050505050565b6000610114608087018761031b565b90506004036101655760405162461bcd60e51b815260206004820152601c60248201527f526576657274206f6e20657874726144617461206c656e67746820340000000060448201526064016100c5565b610172608087018761031b565b905060050361018057600080fd5b600361018c8780610362565b60e00135036101a3576001600160e01b03196101ac565b6303874c7760e21b5b98975050505050505050565b80356001600160a01b03811681146101cf57600080fd5b919050565b600080600080608085870312156101ea57600080fd5b843593506101fa602086016101b8565b9250610208604086016101b8565b9396929550929360600135925050565b60008083601f84011261022a57600080fd5b50813567ffffffffffffffff81111561024257600080fd5b6020830191508360208260051b850101111561025d57600080fd5b9250929050565b600080600080600080600060a0888a03121561027f57600080fd5b8735965061028f602089016101b8565b9550604088013567ffffffffffffffff808211156102ac57600080fd5b9089019060a0828c0312156102c057600080fd5b909550606089013590808211156102d657600080fd5b6102e28b838c01610218565b909650945060808a01359150808211156102fb57600080fd5b506103088a828b01610218565b989b979a50959850939692959293505050565b6000808335601e1984360301811261033257600080fd5b83018035915067ffffffffffffffff82111561034d57600080fd5b60200191503681900382131561025d57600080fd5b6000823561015e1983360301811261037957600080fd5b919091019291505056fea2646970667358221220a2f8597aa17d32909d24ccdd271fb0964a7d59b97598a202aeabdebfd66ec74364736f6c634300080d0033",
}

// TestZoneABI is the input ABI used to generate the binding from.
// Deprecated: Use TestZoneMetaData.ABI instead.
var TestZoneABI = TestZoneMetaData.ABI

// TestZoneBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestZoneMetaData.Bin instead.
var TestZoneBin = TestZoneMetaData.Bin

// DeployTestZone deploys a new Ethereum contract, binding an instance of TestZone to it.
func DeployTestZone(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestZone, error) {
	parsed, err := TestZoneMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestZoneBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestZone{TestZoneCaller: TestZoneCaller{contract: contract}, TestZoneTransactor: TestZoneTransactor{contract: contract}, TestZoneFilterer: TestZoneFilterer{contract: contract}}, nil
}

// TestZone is an auto generated Go binding around an Ethereum contract.
type TestZone struct {
	TestZoneCaller     // Read-only binding to the contract
	TestZoneTransactor // Write-only binding to the contract
	TestZoneFilterer   // Log filterer for contract events
}

// TestZoneCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestZoneCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestZoneTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestZoneTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestZoneFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestZoneFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestZoneSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestZoneSession struct {
	Contract     *TestZone         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestZoneCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestZoneCallerSession struct {
	Contract *TestZoneCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TestZoneTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestZoneTransactorSession struct {
	Contract     *TestZoneTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TestZoneRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestZoneRaw struct {
	Contract *TestZone // Generic contract binding to access the raw methods on
}

// TestZoneCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestZoneCallerRaw struct {
	Contract *TestZoneCaller // Generic read-only contract binding to access the raw methods on
}

// TestZoneTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestZoneTransactorRaw struct {
	Contract *TestZoneTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestZone creates a new instance of TestZone, bound to a specific deployed contract.
func NewTestZone(address common.Address, backend bind.ContractBackend) (*TestZone, error) {
	contract, err := bindTestZone(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestZone{TestZoneCaller: TestZoneCaller{contract: contract}, TestZoneTransactor: TestZoneTransactor{contract: contract}, TestZoneFilterer: TestZoneFilterer{contract: contract}}, nil
}

// NewTestZoneCaller creates a new read-only instance of TestZone, bound to a specific deployed contract.
func NewTestZoneCaller(address common.Address, caller bind.ContractCaller) (*TestZoneCaller, error) {
	contract, err := bindTestZone(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestZoneCaller{contract: contract}, nil
}

// NewTestZoneTransactor creates a new write-only instance of TestZone, bound to a specific deployed contract.
func NewTestZoneTransactor(address common.Address, transactor bind.ContractTransactor) (*TestZoneTransactor, error) {
	contract, err := bindTestZone(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestZoneTransactor{contract: contract}, nil
}

// NewTestZoneFilterer creates a new log filterer instance of TestZone, bound to a specific deployed contract.
func NewTestZoneFilterer(address common.Address, filterer bind.ContractFilterer) (*TestZoneFilterer, error) {
	contract, err := bindTestZone(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestZoneFilterer{contract: contract}, nil
}

// bindTestZone binds a generic wrapper to an already deployed contract.
func bindTestZone(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestZoneABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestZone *TestZoneRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestZone.Contract.TestZoneCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestZone *TestZoneRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestZone.Contract.TestZoneTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestZone *TestZoneRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestZone.Contract.TestZoneTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestZone *TestZoneCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestZone.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestZone *TestZoneTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestZone.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestZone *TestZoneTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestZone.Contract.contract.Transact(opts, method, params...)
}

// IsValidOrder is a free data retrieval call binding the contract method 0x0e1d31dc.
//
// Solidity: function isValidOrder(bytes32 orderHash, address caller, address offerer, bytes32 zoneHash) pure returns(bytes4 validOrderMagicValue)
func (_TestZone *TestZoneCaller) IsValidOrder(opts *bind.CallOpts, orderHash [32]byte, caller common.Address, offerer common.Address, zoneHash [32]byte) ([4]byte, error) {
	var out []interface{}
	err := _TestZone.contract.Call(opts, &out, "isValidOrder", orderHash, caller, offerer, zoneHash)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidOrder is a free data retrieval call binding the contract method 0x0e1d31dc.
//
// Solidity: function isValidOrder(bytes32 orderHash, address caller, address offerer, bytes32 zoneHash) pure returns(bytes4 validOrderMagicValue)
func (_TestZone *TestZoneSession) IsValidOrder(orderHash [32]byte, caller common.Address, offerer common.Address, zoneHash [32]byte) ([4]byte, error) {
	return _TestZone.Contract.IsValidOrder(&_TestZone.CallOpts, orderHash, caller, offerer, zoneHash)
}

// IsValidOrder is a free data retrieval call binding the contract method 0x0e1d31dc.
//
// Solidity: function isValidOrder(bytes32 orderHash, address caller, address offerer, bytes32 zoneHash) pure returns(bytes4 validOrderMagicValue)
func (_TestZone *TestZoneCallerSession) IsValidOrder(orderHash [32]byte, caller common.Address, offerer common.Address, zoneHash [32]byte) ([4]byte, error) {
	return _TestZone.Contract.IsValidOrder(&_TestZone.CallOpts, orderHash, caller, offerer, zoneHash)
}

// IsValidOrderIncludingExtraData is a free data retrieval call binding the contract method 0x33131570.
//
// Solidity: function isValidOrderIncludingExtraData(bytes32 orderHash, address caller, ((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) order, bytes32[] priorOrderHashes, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers) pure returns(bytes4 validOrderMagicValue)
func (_TestZone *TestZoneCaller) IsValidOrderIncludingExtraData(opts *bind.CallOpts, orderHash [32]byte, caller common.Address, order AdvancedOrder, priorOrderHashes [][32]byte, criteriaResolvers []CriteriaResolver) ([4]byte, error) {
	var out []interface{}
	err := _TestZone.contract.Call(opts, &out, "isValidOrderIncludingExtraData", orderHash, caller, order, priorOrderHashes, criteriaResolvers)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidOrderIncludingExtraData is a free data retrieval call binding the contract method 0x33131570.
//
// Solidity: function isValidOrderIncludingExtraData(bytes32 orderHash, address caller, ((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) order, bytes32[] priorOrderHashes, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers) pure returns(bytes4 validOrderMagicValue)
func (_TestZone *TestZoneSession) IsValidOrderIncludingExtraData(orderHash [32]byte, caller common.Address, order AdvancedOrder, priorOrderHashes [][32]byte, criteriaResolvers []CriteriaResolver) ([4]byte, error) {
	return _TestZone.Contract.IsValidOrderIncludingExtraData(&_TestZone.CallOpts, orderHash, caller, order, priorOrderHashes, criteriaResolvers)
}

// IsValidOrderIncludingExtraData is a free data retrieval call binding the contract method 0x33131570.
//
// Solidity: function isValidOrderIncludingExtraData(bytes32 orderHash, address caller, ((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) order, bytes32[] priorOrderHashes, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers) pure returns(bytes4 validOrderMagicValue)
func (_TestZone *TestZoneCallerSession) IsValidOrderIncludingExtraData(orderHash [32]byte, caller common.Address, order AdvancedOrder, priorOrderHashes [][32]byte, criteriaResolvers []CriteriaResolver) ([4]byte, error) {
	return _TestZone.Contract.IsValidOrderIncludingExtraData(&_TestZone.CallOpts, orderHash, caller, order, priorOrderHashes, criteriaResolvers)
}
