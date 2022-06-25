// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TestERC1155

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

// TestERC1155MetaData contains all meta data concerning the TestERC1155 contract.
var TestERC1155MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061107d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100925760003560e01c80632eb2c2d6116100665780632eb2c2d6146101415780634e1273f414610156578063a22cb46514610176578063e985e9c514610189578063f242432a146101b757600080fd5b8062fdd58e1461009757806301ffc9a7146100d25780630e89341c146100f5578063156e29f61461012e575b600080fd5b6100bf6100a53660046109ef565b600060208181529281526040808220909352908152205481565b6040519081526020015b60405180910390f35b6100e56100e0366004610a32565b6101ca565b60405190151581526020016100c9565b610121610103366004610a56565b5060408051808201909152600381526275726960e81b602082015290565b6040516100c99190610abc565b6100e561013c366004610acf565b61021c565b61015461014f366004610c48565b610243565b005b610169610164366004610cf2565b6104fa565b6040516100c99190610ded565b610154610184366004610e00565b610628565b6100e5610197366004610e3c565b600160209081526000928352604080842090915290825290205460ff1681565b6101546101c5366004610e6f565b610694565b60006301ffc9a760e01b6001600160e01b0319831614806101fb5750636cdb3d1360e11b6001600160e01b03198316145b8061021657506303a24d0760e21b6001600160e01b03198316145b92915050565b60006102398484846040518060200160405280600081525061088b565b5060019392505050565b82518251811461028c5760405162461bcd60e51b815260206004820152600f60248201526e0988a9c8ea890be9a92a69a82a8869608b1b60448201526064015b60405180910390fd5b336001600160a01b03871614806102c657506001600160a01b038616600090815260016020908152604080832033845290915290205460ff165b6103035760405162461bcd60e51b815260206004820152600e60248201526d1393d517d055551213d49256915160921b6044820152606401610283565b60005b818110156103d857600085828151811061032257610322610ed4565b60200260200101519050600085838151811061034057610340610ed4565b60200260200101519050806000808b6001600160a01b03166001600160a01b031681526020019081526020016000206000848152602001908152602001600020600082825461038f9190610f00565b90915550506001600160a01b038816600090815260208181526040808320858452909152812080548392906103c5908490610f17565b9091555050600190920191506103069050565b50846001600160a01b0316866001600160a01b0316336001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051610428929190610f2f565b60405180910390a46001600160a01b0385163b156104c95760405163bc197c8160e01b808252906001600160a01b0387169063bc197c81906104769033908b908a908a908a90600401610f5d565b6020604051808303816000875af1158015610495573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104b99190610fbb565b6001600160e01b031916146104d6565b6001600160a01b03851615155b6104f25760405162461bcd60e51b815260040161028390610fd8565b505050505050565b815181516060919081146105425760405162461bcd60e51b815260206004820152600f60248201526e0988a9c8ea890be9a92a69a82a8869608b1b6044820152606401610283565b835167ffffffffffffffff81111561055c5761055c610b02565b604051908082528060200260200182016040528015610585578160200160208202803683370190505b50915060005b81811015610620576000808683815181106105a8576105a8610ed4565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060008583815181106105e4576105e4610ed4565b602002602001015181526020019081526020016000205483828151811061060d5761060d610ed4565b602090810291909101015260010161058b565b505092915050565b3360008181526001602090815260408083206001600160a01b03871680855290835292819020805460ff191686151590811790915590519081529192917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b336001600160a01b03861614806106ce57506001600160a01b038516600090815260016020908152604080832033845290915290205460ff165b61070b5760405162461bcd60e51b815260206004820152600e60248201526d1393d517d055551213d49256915160921b6044820152606401610283565b6001600160a01b0385166000908152602081815260408083208684529091528120805484929061073c908490610f00565b90915550506001600160a01b03841660009081526020818152604080832086845290915281208054849290610772908490610f17565b909155505060408051848152602081018490526001600160a01b03808716929088169133917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a46001600160a01b0384163b1561085b5760405163f23a6e6160e01b808252906001600160a01b0386169063f23a6e61906108089033908a90899089908990600401611002565b6020604051808303816000875af1158015610827573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061084b9190610fbb565b6001600160e01b03191614610868565b6001600160a01b03841615155b6108845760405162461bcd60e51b815260040161028390610fd8565b5050505050565b6001600160a01b038416600090815260208181526040808320868452909152812080548492906108bc908490610f17565b909155505060408051848152602081018490526001600160a01b0386169160009133917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a46001600160a01b0384163b156109a45760405163f23a6e6160e01b808252906001600160a01b0386169063f23a6e6190610951903390600090899089908990600401611002565b6020604051808303816000875af1158015610970573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109949190610fbb565b6001600160e01b031916146109b1565b6001600160a01b03841615155b6109cd5760405162461bcd60e51b815260040161028390610fd8565b50505050565b80356001600160a01b03811681146109ea57600080fd5b919050565b60008060408385031215610a0257600080fd5b610a0b836109d3565b946020939093013593505050565b6001600160e01b031981168114610a2f57600080fd5b50565b600060208284031215610a4457600080fd5b8135610a4f81610a19565b9392505050565b600060208284031215610a6857600080fd5b5035919050565b6000815180845260005b81811015610a9557602081850181015186830182015201610a79565b81811115610aa7576000602083870101525b50601f01601f19169290920160200192915050565b602081526000610a4f6020830184610a6f565b600080600060608486031215610ae457600080fd5b610aed846109d3565b95602085013595506040909401359392505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610b4157610b41610b02565b604052919050565b600067ffffffffffffffff821115610b6357610b63610b02565b5060051b60200190565b600082601f830112610b7e57600080fd5b81356020610b93610b8e83610b49565b610b18565b82815260059290921b84018101918181019086841115610bb257600080fd5b8286015b84811015610bcd5780358352918301918301610bb6565b509695505050505050565b600082601f830112610be957600080fd5b813567ffffffffffffffff811115610c0357610c03610b02565b610c16601f8201601f1916602001610b18565b818152846020838601011115610c2b57600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600060a08688031215610c6057600080fd5b610c69866109d3565b9450610c77602087016109d3565b9350604086013567ffffffffffffffff80821115610c9457600080fd5b610ca089838a01610b6d565b94506060880135915080821115610cb657600080fd5b610cc289838a01610b6d565b93506080880135915080821115610cd857600080fd5b50610ce588828901610bd8565b9150509295509295909350565b60008060408385031215610d0557600080fd5b823567ffffffffffffffff80821115610d1d57600080fd5b818501915085601f830112610d3157600080fd5b81356020610d41610b8e83610b49565b82815260059290921b84018101918181019089841115610d6057600080fd5b948201945b83861015610d8557610d76866109d3565b82529482019490820190610d65565b96505086013592505080821115610d9b57600080fd5b50610da885828601610b6d565b9150509250929050565b600081518084526020808501945080840160005b83811015610de257815187529582019590820190600101610dc6565b509495945050505050565b602081526000610a4f6020830184610db2565b60008060408385031215610e1357600080fd5b610e1c836109d3565b915060208301358015158114610e3157600080fd5b809150509250929050565b60008060408385031215610e4f57600080fd5b610e58836109d3565b9150610e66602084016109d3565b90509250929050565b600080600080600060a08688031215610e8757600080fd5b610e90866109d3565b9450610e9e602087016109d3565b93506040860135925060608601359150608086013567ffffffffffffffff811115610ec857600080fd5b610ce588828901610bd8565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082821015610f1257610f12610eea565b500390565b60008219821115610f2a57610f2a610eea565b500190565b604081526000610f426040830185610db2565b8281036020840152610f548185610db2565b95945050505050565b6001600160a01b0386811682528516602082015260a060408201819052600090610f8990830186610db2565b8281036060840152610f9b8186610db2565b90508281036080840152610faf8185610a6f565b98975050505050505050565b600060208284031215610fcd57600080fd5b8151610a4f81610a19565b60208082526010908201526f155394d0519157d49150d2541251539560821b604082015260600190565b6001600160a01b03868116825285166020820152604081018490526060810183905260a06080820181905260009061103c90830184610a6f565b97965050505050505056fea2646970667358221220dcd50de16258174f95266ee96175b33f2dca953c34966bbc1be631ec67fb799f64736f6c634300080d0033",
}

// TestERC1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use TestERC1155MetaData.ABI instead.
var TestERC1155ABI = TestERC1155MetaData.ABI

// TestERC1155Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestERC1155MetaData.Bin instead.
var TestERC1155Bin = TestERC1155MetaData.Bin

// DeployTestERC1155 deploys a new Ethereum contract, binding an instance of TestERC1155 to it.
func DeployTestERC1155(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestERC1155, error) {
	parsed, err := TestERC1155MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestERC1155Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestERC1155{TestERC1155Caller: TestERC1155Caller{contract: contract}, TestERC1155Transactor: TestERC1155Transactor{contract: contract}, TestERC1155Filterer: TestERC1155Filterer{contract: contract}}, nil
}

// TestERC1155 is an auto generated Go binding around an Ethereum contract.
type TestERC1155 struct {
	TestERC1155Caller     // Read-only binding to the contract
	TestERC1155Transactor // Write-only binding to the contract
	TestERC1155Filterer   // Log filterer for contract events
}

// TestERC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type TestERC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestERC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TestERC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestERC1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestERC1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestERC1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestERC1155Session struct {
	Contract     *TestERC1155      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestERC1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestERC1155CallerSession struct {
	Contract *TestERC1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TestERC1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestERC1155TransactorSession struct {
	Contract     *TestERC1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestERC1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type TestERC1155Raw struct {
	Contract *TestERC1155 // Generic contract binding to access the raw methods on
}

// TestERC1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestERC1155CallerRaw struct {
	Contract *TestERC1155Caller // Generic read-only contract binding to access the raw methods on
}

// TestERC1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestERC1155TransactorRaw struct {
	Contract *TestERC1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTestERC1155 creates a new instance of TestERC1155, bound to a specific deployed contract.
func NewTestERC1155(address common.Address, backend bind.ContractBackend) (*TestERC1155, error) {
	contract, err := bindTestERC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestERC1155{TestERC1155Caller: TestERC1155Caller{contract: contract}, TestERC1155Transactor: TestERC1155Transactor{contract: contract}, TestERC1155Filterer: TestERC1155Filterer{contract: contract}}, nil
}

// NewTestERC1155Caller creates a new read-only instance of TestERC1155, bound to a specific deployed contract.
func NewTestERC1155Caller(address common.Address, caller bind.ContractCaller) (*TestERC1155Caller, error) {
	contract, err := bindTestERC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestERC1155Caller{contract: contract}, nil
}

// NewTestERC1155Transactor creates a new write-only instance of TestERC1155, bound to a specific deployed contract.
func NewTestERC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*TestERC1155Transactor, error) {
	contract, err := bindTestERC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestERC1155Transactor{contract: contract}, nil
}

// NewTestERC1155Filterer creates a new log filterer instance of TestERC1155, bound to a specific deployed contract.
func NewTestERC1155Filterer(address common.Address, filterer bind.ContractFilterer) (*TestERC1155Filterer, error) {
	contract, err := bindTestERC1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestERC1155Filterer{contract: contract}, nil
}

// bindTestERC1155 binds a generic wrapper to an already deployed contract.
func bindTestERC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestERC1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestERC1155 *TestERC1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestERC1155.Contract.TestERC1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestERC1155 *TestERC1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestERC1155.Contract.TestERC1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestERC1155 *TestERC1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestERC1155.Contract.TestERC1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestERC1155 *TestERC1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestERC1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestERC1155 *TestERC1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestERC1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestERC1155 *TestERC1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestERC1155.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address , uint256 ) view returns(uint256)
func (_TestERC1155 *TestERC1155Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestERC1155.contract.Call(opts, &out, "balanceOf", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address , uint256 ) view returns(uint256)
func (_TestERC1155 *TestERC1155Session) BalanceOf(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TestERC1155.Contract.BalanceOf(&_TestERC1155.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address , uint256 ) view returns(uint256)
func (_TestERC1155 *TestERC1155CallerSession) BalanceOf(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TestERC1155.Contract.BalanceOf(&_TestERC1155.CallOpts, arg0, arg1)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances)
func (_TestERC1155 *TestERC1155Caller) BalanceOfBatch(opts *bind.CallOpts, owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _TestERC1155.contract.Call(opts, &out, "balanceOfBatch", owners, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances)
func (_TestERC1155 *TestERC1155Session) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _TestERC1155.Contract.BalanceOfBatch(&_TestERC1155.CallOpts, owners, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances)
func (_TestERC1155 *TestERC1155CallerSession) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _TestERC1155.Contract.BalanceOfBatch(&_TestERC1155.CallOpts, owners, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_TestERC1155 *TestERC1155Caller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _TestERC1155.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_TestERC1155 *TestERC1155Session) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _TestERC1155.Contract.IsApprovedForAll(&_TestERC1155.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_TestERC1155 *TestERC1155CallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _TestERC1155.Contract.IsApprovedForAll(&_TestERC1155.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_TestERC1155 *TestERC1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TestERC1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_TestERC1155 *TestERC1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestERC1155.Contract.SupportsInterface(&_TestERC1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_TestERC1155 *TestERC1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestERC1155.Contract.SupportsInterface(&_TestERC1155.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) pure returns(string)
func (_TestERC1155 *TestERC1155Caller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _TestERC1155.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) pure returns(string)
func (_TestERC1155 *TestERC1155Session) Uri(arg0 *big.Int) (string, error) {
	return _TestERC1155.Contract.Uri(&_TestERC1155.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) pure returns(string)
func (_TestERC1155 *TestERC1155CallerSession) Uri(arg0 *big.Int) (string, error) {
	return _TestERC1155.Contract.Uri(&_TestERC1155.CallOpts, arg0)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 tokenId, uint256 amount) returns(bool)
func (_TestERC1155 *TestERC1155Transactor) Mint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TestERC1155.contract.Transact(opts, "mint", to, tokenId, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 tokenId, uint256 amount) returns(bool)
func (_TestERC1155 *TestERC1155Session) Mint(to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TestERC1155.Contract.Mint(&_TestERC1155.TransactOpts, to, tokenId, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 tokenId, uint256 amount) returns(bool)
func (_TestERC1155 *TestERC1155TransactorSession) Mint(to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TestERC1155.Contract.Mint(&_TestERC1155.TransactOpts, to, tokenId, amount)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_TestERC1155 *TestERC1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _TestERC1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_TestERC1155 *TestERC1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _TestERC1155.Contract.SafeBatchTransferFrom(&_TestERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_TestERC1155 *TestERC1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _TestERC1155.Contract.SafeBatchTransferFrom(&_TestERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_TestERC1155 *TestERC1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TestERC1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_TestERC1155 *TestERC1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TestERC1155.Contract.SafeTransferFrom(&_TestERC1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_TestERC1155 *TestERC1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TestERC1155.Contract.SafeTransferFrom(&_TestERC1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TestERC1155 *TestERC1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _TestERC1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TestERC1155 *TestERC1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TestERC1155.Contract.SetApprovalForAll(&_TestERC1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TestERC1155 *TestERC1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TestERC1155.Contract.SetApprovalForAll(&_TestERC1155.TransactOpts, operator, approved)
}

// TestERC1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the TestERC1155 contract.
type TestERC1155ApprovalForAllIterator struct {
	Event *TestERC1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TestERC1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestERC1155ApprovalForAll)
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
		it.Event = new(TestERC1155ApprovalForAll)
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
func (it *TestERC1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestERC1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestERC1155ApprovalForAll represents a ApprovalForAll event raised by the TestERC1155 contract.
type TestERC1155ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TestERC1155 *TestERC1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TestERC1155ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TestERC1155.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TestERC1155ApprovalForAllIterator{contract: _TestERC1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TestERC1155 *TestERC1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TestERC1155ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TestERC1155.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestERC1155ApprovalForAll)
				if err := _TestERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TestERC1155 *TestERC1155Filterer) ParseApprovalForAll(log types.Log) (*TestERC1155ApprovalForAll, error) {
	event := new(TestERC1155ApprovalForAll)
	if err := _TestERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestERC1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the TestERC1155 contract.
type TestERC1155TransferBatchIterator struct {
	Event *TestERC1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *TestERC1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestERC1155TransferBatch)
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
		it.Event = new(TestERC1155TransferBatch)
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
func (it *TestERC1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestERC1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestERC1155TransferBatch represents a TransferBatch event raised by the TestERC1155 contract.
type TestERC1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_TestERC1155 *TestERC1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*TestERC1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestERC1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestERC1155TransferBatchIterator{contract: _TestERC1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_TestERC1155 *TestERC1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *TestERC1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestERC1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestERC1155TransferBatch)
				if err := _TestERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_TestERC1155 *TestERC1155Filterer) ParseTransferBatch(log types.Log) (*TestERC1155TransferBatch, error) {
	event := new(TestERC1155TransferBatch)
	if err := _TestERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestERC1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the TestERC1155 contract.
type TestERC1155TransferSingleIterator struct {
	Event *TestERC1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *TestERC1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestERC1155TransferSingle)
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
		it.Event = new(TestERC1155TransferSingle)
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
func (it *TestERC1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestERC1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestERC1155TransferSingle represents a TransferSingle event raised by the TestERC1155 contract.
type TestERC1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_TestERC1155 *TestERC1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*TestERC1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestERC1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestERC1155TransferSingleIterator{contract: _TestERC1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_TestERC1155 *TestERC1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *TestERC1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestERC1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestERC1155TransferSingle)
				if err := _TestERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_TestERC1155 *TestERC1155Filterer) ParseTransferSingle(log types.Log) (*TestERC1155TransferSingle, error) {
	event := new(TestERC1155TransferSingle)
	if err := _TestERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestERC1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the TestERC1155 contract.
type TestERC1155URIIterator struct {
	Event *TestERC1155URI // Event containing the contract specifics and raw log

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
func (it *TestERC1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestERC1155URI)
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
		it.Event = new(TestERC1155URI)
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
func (it *TestERC1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestERC1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestERC1155URI represents a URI event raised by the TestERC1155 contract.
type TestERC1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_TestERC1155 *TestERC1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*TestERC1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TestERC1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &TestERC1155URIIterator{contract: _TestERC1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_TestERC1155 *TestERC1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *TestERC1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TestERC1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestERC1155URI)
				if err := _TestERC1155.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_TestERC1155 *TestERC1155Filterer) ParseURI(log types.Log) (*TestERC1155URI, error) {
	event := new(TestERC1155URI)
	if err := _TestERC1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
