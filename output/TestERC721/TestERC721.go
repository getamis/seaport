// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TestERC721

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

// TestERC721MetaData contains all meta data concerning the TestERC721 contract.
var TestERC721MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060408051808201825260078152665465737437323160c81b60208083019182528351808501909452600684526554535437323160d01b9084015281519192916200005f916000916200007e565b508051620000759060019060208401906200007e565b50505062000160565b8280546200008c9062000124565b90600052602060002090601f016020900481019282620000b05760008555620000fb565b82601f10620000cb57805160ff1916838001178555620000fb565b82800160010185558215620000fb579182015b82811115620000fb578251825591602001919060010190620000de565b50620001099291506200010d565b5090565b5b808211156200010957600081556001016200010e565b600181811c908216806200013957607f821691505b6020821081036200015a57634e487b7160e01b600052602260045260246000fd5b50919050565b610d2180620001706000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80636352211e1161008c578063a22cb46511610066578063a22cb4651461021a578063b88d4fde1461022d578063c87b56dd14610240578063e985e9c51461027157600080fd5b80636352211e146101bb57806370a08231146101e457806395d89b411461021257600080fd5b8063095ea7b3116100c8578063095ea7b31461016d57806323b872dd1461018257806340c10f191461019557806342842e0e146101a857600080fd5b806301ffc9a7146100ef57806306fdde0314610117578063081812fc1461012c575b600080fd5b6101026100fd3660046109bc565b61029f565b60405190151581526020015b60405180910390f35b61011f6102f1565b60405161010e9190610a2d565b61015561013a366004610a40565b6004602052600090815260409020546001600160a01b031681565b6040516001600160a01b03909116815260200161010e565b61018061017b366004610a75565b61037f565b005b610180610190366004610a9f565b610466565b6101026101a3366004610a75565b61062d565b6101806101b6366004610a9f565b610642565b6101556101c9366004610a40565b6003602052600090815260409020546001600160a01b031681565b6102046101f2366004610adb565b60026020526000908152604090205481565b60405190815260200161010e565b61011f61073a565b610180610228366004610af6565b610747565b61018061023b366004610b48565b6107b3565b61011f61024e366004610a40565b50604080518082019091526008815267746f6b656e55524960c01b602082015290565b61010261027f366004610c24565b600560209081526000928352604080842090915290825290205460ff1681565b60006301ffc9a760e01b6001600160e01b0319831614806102d057506380ac58cd60e01b6001600160e01b03198316145b806102eb5750635b5e139f60e01b6001600160e01b03198316145b92915050565b600080546102fe90610c57565b80601f016020809104026020016040519081016040528092919081815260200182805461032a90610c57565b80156103775780601f1061034c57610100808354040283529160200191610377565b820191906000526020600020905b81548152906001019060200180831161035a57829003601f168201915b505050505081565b6000818152600360205260409020546001600160a01b0316338114806103c857506001600160a01b038116600090815260056020908152604080832033845290915290205460ff165b61040a5760405162461bcd60e51b815260206004820152600e60248201526d1393d517d055551213d49256915160921b60448201526064015b60405180910390fd5b60008281526004602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b6000818152600360205260409020546001600160a01b038481169116146104bc5760405162461bcd60e51b815260206004820152600a60248201526957524f4e475f46524f4d60b01b6044820152606401610401565b6001600160a01b0382166105065760405162461bcd60e51b81526020600482015260116024820152701253959053125117d49150d25412515395607a1b6044820152606401610401565b336001600160a01b038416148061053357506000818152600460205260409020546001600160a01b031633145b8061056157506001600160a01b038316600090815260056020908152604080832033845290915290205460ff165b61059e5760405162461bcd60e51b815260206004820152600e60248201526d1393d517d055551213d49256915160921b6044820152606401610401565b6001600160a01b0380841660008181526002602090815260408083208054600019019055938616808352848320805460010190558583526003825284832080546001600160a01b03199081168317909155600490925284832080549092169091559251849392917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b60006106398383610898565b50600192915050565b61064d838383610466565b6001600160a01b0382163b15806106f65750604051630a85bd0160e11b8082523360048301526001600160a01b03858116602484015260448301849052608060648401526000608484015290919084169063150b7a029060a4016020604051808303816000875af11580156106c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ea9190610c91565b6001600160e01b031916145b6107355760405162461bcd60e51b815260206004820152601060248201526f155394d0519157d49150d2541251539560821b6044820152606401610401565b505050565b600180546102fe90610c57565b3360008181526005602090815260408083206001600160a01b03871680855290835292819020805460ff191686151590811790915590519081529192917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b6107be848484610466565b6001600160a01b0383163b15806108535750604051630a85bd0160e11b808252906001600160a01b0385169063150b7a0290610804903390899088908890600401610cae565b6020604051808303816000875af1158015610823573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108479190610c91565b6001600160e01b031916145b6108925760405162461bcd60e51b815260206004820152601060248201526f155394d0519157d49150d2541251539560821b6044820152606401610401565b50505050565b6001600160a01b0382166108e25760405162461bcd60e51b81526020600482015260116024820152701253959053125117d49150d25412515395607a1b6044820152606401610401565b6000818152600360205260409020546001600160a01b0316156109385760405162461bcd60e51b815260206004820152600e60248201526d1053149150511657d3525395115160921b6044820152606401610401565b6001600160a01b038216600081815260026020908152604080832080546001019055848352600390915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b6001600160e01b0319811681146109b957600080fd5b50565b6000602082840312156109ce57600080fd5b81356109d9816109a3565b9392505050565b6000815180845260005b81811015610a06576020818501810151868301820152016109ea565b81811115610a18576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006109d960208301846109e0565b600060208284031215610a5257600080fd5b5035919050565b80356001600160a01b0381168114610a7057600080fd5b919050565b60008060408385031215610a8857600080fd5b610a9183610a59565b946020939093013593505050565b600080600060608486031215610ab457600080fd5b610abd84610a59565b9250610acb60208501610a59565b9150604084013590509250925092565b600060208284031215610aed57600080fd5b6109d982610a59565b60008060408385031215610b0957600080fd5b610b1283610a59565b915060208301358015158114610b2757600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b60008060008060808587031215610b5e57600080fd5b610b6785610a59565b9350610b7560208601610a59565b925060408501359150606085013567ffffffffffffffff80821115610b9957600080fd5b818701915087601f830112610bad57600080fd5b813581811115610bbf57610bbf610b32565b604051601f8201601f19908116603f01168101908382118183101715610be757610be7610b32565b816040528281528a6020848701011115610c0057600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b60008060408385031215610c3757600080fd5b610c4083610a59565b9150610c4e60208401610a59565b90509250929050565b600181811c90821680610c6b57607f821691505b602082108103610c8b57634e487b7160e01b600052602260045260246000fd5b50919050565b600060208284031215610ca357600080fd5b81516109d9816109a3565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090610ce1908301846109e0565b969550505050505056fea26469706673582212203c54dffcd9715d0138471e423cef2766ccee4d07a8f623f007778f1a56a7d25564736f6c634300080d0033",
}

// TestERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use TestERC721MetaData.ABI instead.
var TestERC721ABI = TestERC721MetaData.ABI

// TestERC721Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestERC721MetaData.Bin instead.
var TestERC721Bin = TestERC721MetaData.Bin

// DeployTestERC721 deploys a new Ethereum contract, binding an instance of TestERC721 to it.
func DeployTestERC721(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestERC721, error) {
	parsed, err := TestERC721MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestERC721Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestERC721{TestERC721Caller: TestERC721Caller{contract: contract}, TestERC721Transactor: TestERC721Transactor{contract: contract}, TestERC721Filterer: TestERC721Filterer{contract: contract}}, nil
}

// TestERC721 is an auto generated Go binding around an Ethereum contract.
type TestERC721 struct {
	TestERC721Caller     // Read-only binding to the contract
	TestERC721Transactor // Write-only binding to the contract
	TestERC721Filterer   // Log filterer for contract events
}

// TestERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type TestERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TestERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestERC721Session struct {
	Contract     *TestERC721       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestERC721CallerSession struct {
	Contract *TestERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TestERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestERC721TransactorSession struct {
	Contract     *TestERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TestERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type TestERC721Raw struct {
	Contract *TestERC721 // Generic contract binding to access the raw methods on
}

// TestERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestERC721CallerRaw struct {
	Contract *TestERC721Caller // Generic read-only contract binding to access the raw methods on
}

// TestERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestERC721TransactorRaw struct {
	Contract *TestERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTestERC721 creates a new instance of TestERC721, bound to a specific deployed contract.
func NewTestERC721(address common.Address, backend bind.ContractBackend) (*TestERC721, error) {
	contract, err := bindTestERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestERC721{TestERC721Caller: TestERC721Caller{contract: contract}, TestERC721Transactor: TestERC721Transactor{contract: contract}, TestERC721Filterer: TestERC721Filterer{contract: contract}}, nil
}

// NewTestERC721Caller creates a new read-only instance of TestERC721, bound to a specific deployed contract.
func NewTestERC721Caller(address common.Address, caller bind.ContractCaller) (*TestERC721Caller, error) {
	contract, err := bindTestERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestERC721Caller{contract: contract}, nil
}

// NewTestERC721Transactor creates a new write-only instance of TestERC721, bound to a specific deployed contract.
func NewTestERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*TestERC721Transactor, error) {
	contract, err := bindTestERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestERC721Transactor{contract: contract}, nil
}

// NewTestERC721Filterer creates a new log filterer instance of TestERC721, bound to a specific deployed contract.
func NewTestERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*TestERC721Filterer, error) {
	contract, err := bindTestERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestERC721Filterer{contract: contract}, nil
}

// bindTestERC721 binds a generic wrapper to an already deployed contract.
func bindTestERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestERC721 *TestERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestERC721.Contract.TestERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestERC721 *TestERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestERC721.Contract.TestERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestERC721 *TestERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestERC721.Contract.TestERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestERC721 *TestERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestERC721 *TestERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestERC721 *TestERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_TestERC721 *TestERC721Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestERC721.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_TestERC721 *TestERC721Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _TestERC721.Contract.BalanceOf(&_TestERC721.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_TestERC721 *TestERC721CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _TestERC721.Contract.BalanceOf(&_TestERC721.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_TestERC721 *TestERC721Caller) GetApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TestERC721.contract.Call(opts, &out, "getApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_TestERC721 *TestERC721Session) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _TestERC721.Contract.GetApproved(&_TestERC721.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_TestERC721 *TestERC721CallerSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _TestERC721.Contract.GetApproved(&_TestERC721.CallOpts, arg0)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_TestERC721 *TestERC721Caller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _TestERC721.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_TestERC721 *TestERC721Session) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _TestERC721.Contract.IsApprovedForAll(&_TestERC721.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_TestERC721 *TestERC721CallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _TestERC721.Contract.IsApprovedForAll(&_TestERC721.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestERC721 *TestERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestERC721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestERC721 *TestERC721Session) Name() (string, error) {
	return _TestERC721.Contract.Name(&_TestERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestERC721 *TestERC721CallerSession) Name() (string, error) {
	return _TestERC721.Contract.Name(&_TestERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 ) view returns(address)
func (_TestERC721 *TestERC721Caller) OwnerOf(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TestERC721.contract.Call(opts, &out, "ownerOf", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 ) view returns(address)
func (_TestERC721 *TestERC721Session) OwnerOf(arg0 *big.Int) (common.Address, error) {
	return _TestERC721.Contract.OwnerOf(&_TestERC721.CallOpts, arg0)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 ) view returns(address)
func (_TestERC721 *TestERC721CallerSession) OwnerOf(arg0 *big.Int) (common.Address, error) {
	return _TestERC721.Contract.OwnerOf(&_TestERC721.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_TestERC721 *TestERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TestERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_TestERC721 *TestERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestERC721.Contract.SupportsInterface(&_TestERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_TestERC721 *TestERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestERC721.Contract.SupportsInterface(&_TestERC721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestERC721 *TestERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestERC721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestERC721 *TestERC721Session) Symbol() (string, error) {
	return _TestERC721.Contract.Symbol(&_TestERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestERC721 *TestERC721CallerSession) Symbol() (string, error) {
	return _TestERC721.Contract.Symbol(&_TestERC721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 ) pure returns(string)
func (_TestERC721 *TestERC721Caller) TokenURI(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _TestERC721.contract.Call(opts, &out, "tokenURI", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 ) pure returns(string)
func (_TestERC721 *TestERC721Session) TokenURI(arg0 *big.Int) (string, error) {
	return _TestERC721.Contract.TokenURI(&_TestERC721.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 ) pure returns(string)
func (_TestERC721 *TestERC721CallerSession) TokenURI(arg0 *big.Int) (string, error) {
	return _TestERC721.Contract.TokenURI(&_TestERC721.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_TestERC721 *TestERC721Transactor) Approve(opts *bind.TransactOpts, spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _TestERC721.contract.Transact(opts, "approve", spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_TestERC721 *TestERC721Session) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _TestERC721.Contract.Approve(&_TestERC721.TransactOpts, spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_TestERC721 *TestERC721TransactorSession) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _TestERC721.Contract.Approve(&_TestERC721.TransactOpts, spender, id)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns(bool)
func (_TestERC721 *TestERC721Transactor) Mint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TestERC721.contract.Transact(opts, "mint", to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns(bool)
func (_TestERC721 *TestERC721Session) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TestERC721.Contract.Mint(&_TestERC721.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns(bool)
func (_TestERC721 *TestERC721TransactorSession) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TestERC721.Contract.Mint(&_TestERC721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_TestERC721 *TestERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _TestERC721.contract.Transact(opts, "safeTransferFrom", from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_TestERC721 *TestERC721Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _TestERC721.Contract.SafeTransferFrom(&_TestERC721.TransactOpts, from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_TestERC721 *TestERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _TestERC721.Contract.SafeTransferFrom(&_TestERC721.TransactOpts, from, to, id)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_TestERC721 *TestERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _TestERC721.contract.Transact(opts, "safeTransferFrom0", from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_TestERC721 *TestERC721Session) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _TestERC721.Contract.SafeTransferFrom0(&_TestERC721.TransactOpts, from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_TestERC721 *TestERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _TestERC721.Contract.SafeTransferFrom0(&_TestERC721.TransactOpts, from, to, id, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TestERC721 *TestERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _TestERC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TestERC721 *TestERC721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TestERC721.Contract.SetApprovalForAll(&_TestERC721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TestERC721 *TestERC721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TestERC721.Contract.SetApprovalForAll(&_TestERC721.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_TestERC721 *TestERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _TestERC721.contract.Transact(opts, "transferFrom", from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_TestERC721 *TestERC721Session) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _TestERC721.Contract.TransferFrom(&_TestERC721.TransactOpts, from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_TestERC721 *TestERC721TransactorSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _TestERC721.Contract.TransferFrom(&_TestERC721.TransactOpts, from, to, id)
}

// TestERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TestERC721 contract.
type TestERC721ApprovalIterator struct {
	Event *TestERC721Approval // Event containing the contract specifics and raw log

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
func (it *TestERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestERC721Approval)
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
		it.Event = new(TestERC721Approval)
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
func (it *TestERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestERC721Approval represents a Approval event raised by the TestERC721 contract.
type TestERC721Approval struct {
	Owner   common.Address
	Spender common.Address
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_TestERC721 *TestERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address, id []*big.Int) (*TestERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TestERC721.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return &TestERC721ApprovalIterator{contract: _TestERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_TestERC721 *TestERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TestERC721Approval, owner []common.Address, spender []common.Address, id []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TestERC721.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestERC721Approval)
				if err := _TestERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_TestERC721 *TestERC721Filterer) ParseApproval(log types.Log) (*TestERC721Approval, error) {
	event := new(TestERC721Approval)
	if err := _TestERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the TestERC721 contract.
type TestERC721ApprovalForAllIterator struct {
	Event *TestERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TestERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestERC721ApprovalForAll)
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
		it.Event = new(TestERC721ApprovalForAll)
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
func (it *TestERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestERC721ApprovalForAll represents a ApprovalForAll event raised by the TestERC721 contract.
type TestERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TestERC721 *TestERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TestERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TestERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TestERC721ApprovalForAllIterator{contract: _TestERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TestERC721 *TestERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TestERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TestERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestERC721ApprovalForAll)
				if err := _TestERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_TestERC721 *TestERC721Filterer) ParseApprovalForAll(log types.Log) (*TestERC721ApprovalForAll, error) {
	event := new(TestERC721ApprovalForAll)
	if err := _TestERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TestERC721 contract.
type TestERC721TransferIterator struct {
	Event *TestERC721Transfer // Event containing the contract specifics and raw log

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
func (it *TestERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestERC721Transfer)
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
		it.Event = new(TestERC721Transfer)
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
func (it *TestERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestERC721Transfer represents a Transfer event raised by the TestERC721 contract.
type TestERC721Transfer struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_TestERC721 *TestERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*TestERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TestERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &TestERC721TransferIterator{contract: _TestERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_TestERC721 *TestERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TestERC721Transfer, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TestERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestERC721Transfer)
				if err := _TestERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_TestERC721 *TestERC721Filterer) ParseTransfer(log types.Log) (*TestERC721Transfer, error) {
	event := new(TestERC721Transfer)
	if err := _TestERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
