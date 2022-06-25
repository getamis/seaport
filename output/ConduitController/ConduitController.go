// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ConduitController

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

// ConduitControllerMetaData contains all meta data concerning the ConduitController contract.
var ConduitControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"CallerIsNotNewPotentialOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"CallerIsNotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"ChannelOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"ConduitAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCreator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newPotentialOwner\",\"type\":\"address\"}],\"name\":\"NewPotentialOwnerAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"NewPotentialOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoConduit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"NoPotentialOwnerCurrentlySet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"name\":\"NewConduit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPotentialOwner\",\"type\":\"address\"}],\"name\":\"PotentialOwnerUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"cancelOwnershipTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"createConduit\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"channelIndex\",\"type\":\"uint256\"}],\"name\":\"getChannel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"channel\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"channel\",\"type\":\"address\"}],\"name\":\"getChannelStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"getChannels\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"channels\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"name\":\"getConduit\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConduitCodeHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"creationCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"runtimeCodeHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"getPotentialOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"potentialOwner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"getTotalChannels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalChannels\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newPotentialOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"channel\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"updateChannel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b5060405161002060208201610088565b6020820181038252601f19601f82011660405250805190602001206080818152505060008060001b60405161005490610088565b8190604051809103906000f5905080158015610074573d6000803e3d6000fd5b506001600160a01b03163f60a05250610095565b610ae380611a3f83390190565b60805160a0516119696100d660003960008181610148015281816108370152610907015260008181610125015281816107df01526108c201526119696000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80636d4354211161008c5780637b37e561116100665780637b37e561146102495780638b9e028b1461025c578063906c87cc1461027c57806393790f441461028f57600080fd5b80636d435421146101f15780636e9bfd9f14610204578063794593bc1461023657600080fd5b806314afd79e116100c857806314afd79e1461018757806333bc85721461019a5780634e3f9580146101bd57806351710e45146101de57600080fd5b8063027cc764146100ef5780630a96ad391461011f57806313ad9cab14610172575b600080fd5b6101026100fd366004610c72565b6102a2565b6040516001600160a01b0390911681526020015b60405180910390f35b604080517f000000000000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000602082015201610116565b610185610180366004610c9c565b610343565b005b610102610195366004610ce8565b610549565b6101ad6101a8366004610d0a565b610576565b6040519015158152602001610116565b6101d06101cb366004610ce8565b6105b1565b604051908152602001610116565b6101856101ec366004610ce8565b6105db565b6101856101ff366004610d0a565b6106e0565b610217610212366004610d3d565b6107d3565b604080516001600160a01b039093168352901515602083015201610116565b610102610244366004610d56565b61085e565b610185610257366004610ce8565b610a37565b61026f61026a366004610ce8565b610adb565b6040516101169190610d79565b61010261028a366004610ce8565b610b5b565b6101d061029d366004610ce8565b610b88565b60006102ad83610bc4565b6001600160a01b0383166000908152602081905260409020600301548083106102f957604051636ceb340b60e01b81526001600160a01b03851660048201526024015b60405180910390fd5b6001600160a01b038416600090815260208190526040902060030180548490811061032657610326610dc6565b6000918252602090912001546001600160a01b0316949350505050565b61034c83610bfd565b60405163c4e8fcb560e01b81526001600160a01b038381166004830152821515602483015284169063c4e8fcb590604401600060405180830381600087803b15801561039757600080fd5b505af11580156103ab573d6000803e3d6000fd5b505050506001600160a01b038381166000908152602081815260408083209386168352600484019091529020548015158380156103e6575080155b15610436576003830180546001810182556000828152602080822090920180546001600160a01b0319166001600160a01b038a169081179091559254928152600486019091526040902055610541565b831580156104415750805b1561054157600383015460001983019060009061046090600190610ddc565b90508181146104ee57600085600301828154811061048057610480610dc6565b6000918252602090912001546003870180546001600160a01b0390921692508291859081106104b1576104b1610dc6565b600091825260208083209190910180546001600160a01b0319166001600160a01b0394851617905592909116815260048701909152604090208490555b8460030180548061050157610501610e01565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b0389168252600487019052604081205550505b505050505050565b600061055482610bc4565b506001600160a01b039081166000908152602081905260409020600101541690565b600061058183610bc4565b506001600160a01b0391821660009081526020818152604080832093909416825260049092019091522054151590565b60006105bc82610bc4565b506001600160a01b031660009081526020819052604090206003015490565b6105e481610bc4565b6001600160a01b0381811660009081526020819052604090206002015416331461062c576040516388c3a11560e01b81526001600160a01b03821660048201526024016102f0565b6040516000907f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da908290a26001600160a01b038082166000818152602081905260408082206002810180546001600160a01b031916905560010154905133949190911692917fc8894f26f396ce8c004245c8b7cd1b92103a6e4302fcbab883987149ac01b7ec91a46001600160a01b0316600090815260208190526040902060010180546001600160a01b03191633179055565b6106e982610bfd565b6001600160a01b03811661071b5760405163a388d26360e01b81526001600160a01b03831660048201526024016102f0565b6001600160a01b0380831660009081526020819052604090206002015481169082160361076e576040516365e0406560e11b81526001600160a01b038084166004830152821660248201526044016102f0565b6040516001600160a01b038216907f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da90600090a26001600160a01b03918216600090815260208190526040902060020180546001600160a01b03191691909216179055565b60008060ff60f81b30847f00000000000000000000000000000000000000000000000000000000000000006040516020016108119493929190610e17565b60408051601f198184030181529190528051602090910120936001600160a01b0385163f7f0000000000000000000000000000000000000000000000000000000000000000149350915050565b60006001600160a01b0382166108875760405163267eaa8160e21b815260040160405180910390fd5b606083901c33146108ab576040516332db94d160e21b815260040160405180910390fd5b6040516108ea906001600160f81b031990309086907f000000000000000000000000000000000000000000000000000000000000000090602001610e17565b6040516020818303038152906040528051906020012060001c90507f0000000000000000000000000000000000000000000000000000000000000000816001600160a01b03163f0361095a57604051633194665960e11b81526001600160a01b03821660048201526024016102f0565b8260405161096790610c4e565b8190604051809103906000f5905080158015610987573d6000803e3d6000fd5b50506001600160a01b03818116600081815260208181526040918290206001810180546001600160a01b03191695881695909517909455868455815192835282018690527f4397af6128d529b8ae0442f99db1296d5136062597a15bbc61c1b2a6431a7d15910160405180910390a16040516001600160a01b03808516916000918516907fc8894f26f396ce8c004245c8b7cd1b92103a6e4302fcbab883987149ac01b7ec908390a45092915050565b610a4081610bfd565b6001600160a01b0381811660009081526020819052604090206002015416610a86576040516335809b0b60e11b81526001600160a01b03821660048201526024016102f0565b6040516000907f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da908290a26001600160a01b0316600090815260208190526040902060020180546001600160a01b0319169055565b6060610ae682610bc4565b6001600160a01b0382166000908152602081815260409182902060030180548351818402810184019094528084529091830182828015610b4f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610b31575b50505050509050919050565b6000610b6682610bc4565b506001600160a01b039081166000908152602081905260409020600201541690565b6001600160a01b03811660009081526020819052604090205480610bbf576040516304ca820960e41b815260040160405180910390fd5b919050565b6001600160a01b038116600090815260208190526040902054610bfa576040516304ca820960e41b815260040160405180910390fd5b50565b610c0681610bc4565b6001600160a01b03818116600090815260208190526040902060010154163314610bfa5760405163d4ed9a1760e01b81526001600160a01b03821660048201526024016102f0565b610ae380610e5183390190565b80356001600160a01b0381168114610bbf57600080fd5b60008060408385031215610c8557600080fd5b610c8e83610c5b565b946020939093013593505050565b600080600060608486031215610cb157600080fd5b610cba84610c5b565b9250610cc860208501610c5b565b915060408401358015158114610cdd57600080fd5b809150509250925092565b600060208284031215610cfa57600080fd5b610d0382610c5b565b9392505050565b60008060408385031215610d1d57600080fd5b610d2683610c5b565b9150610d3460208401610c5b565b90509250929050565b600060208284031215610d4f57600080fd5b5035919050565b60008060408385031215610d6957600080fd5b82359150610d3460208401610c5b565b6020808252825182820181905260009190848201906040850190845b81811015610dba5783516001600160a01b031683529284019291840191600101610d95565b50909695505050505050565b634e487b7160e01b600052603260045260246000fd5b600082821015610dfc57634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052603160045260246000fd5b6001600160f81b031994909416845260609290921b6bffffffffffffffffffffffff19166001840152601583015260358201526055019056fe60a060405234801561001057600080fd5b5033608052608051610ab361003060003960006101e90152610ab36000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80634ce34aa214610051578063899e104c146100815780638df25d9214610094578063c4e8fcb5146100a7575b600080fd5b61006461005f36600461088d565b6100bc565b6040516001600160e01b0319909116815260200160405180910390f35b61006461008f366004610914565b61012b565b6100646100a2366004610980565b61019b565b6100ba6100b53660046109d2565b6101de565b005b60003360005260006020526040600020546100e6576349ed56f960e11b6000523360045260246000fd5b8160005b8181101561011a5761011285858381811061010757610107610a0e565b905060c002016102dc565b6001016100ea565b50632671a55160e11b949350505050565b6000336000526000602052604060002054610155576349ed56f960e11b6000523360045260246000fd5b8360005b8181101561017e5761017687878381811061010757610107610a0e565b600101610159565b506101898484610448565b50632267841360e21b95945050505050565b60003360005260006020526040600020546101c5576349ed56f960e11b6000523360045260246000fd5b6101cf8383610448565b506346f92ec960e11b92915050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610227576040516336abb4df60e11b815260040160405180910390fd5b6001600160a01b03821660009081526020819052604090205481151560ff90911615150361027f576040516349271a0f60e11b81526001600160a01b0383166004820152811515602482015260440160405180910390fd5b6001600160a01b03821660008181526020818152604091829020805460ff191685151590811790915591519182527fae63067d43ac07563b7eb8db6595635fc77f1578a2a5ea06ba91b63e2afa37e2910160405180910390a25050565b60016102eb6020830183610a3a565b60038111156102fc576102fc610a24565b036103415761033e6103146040830160208401610a62565b6103246060840160408501610a62565b6103346080850160608601610a62565b8460a0013561058d565b50565b60026103506020830183610a3a565b600381111561036157610361610a24565b036103c8578060a0013560011461038b5760405163efcc00b160e01b815260040160405180910390fd5b61033e61039e6040830160208401610a62565b6103ae6060840160408501610a62565b6103be6080850160608601610a62565b8460800135610696565b60036103d76020830183610a3a565b60038111156103e8576103e8610a24565b0361042f5761033e6104006040830160208401610a62565b6104106060840160408501610a62565b6104206080850160608601610a62565b84608001358560a0013561075a565b604051631e4cbc7f60e21b815260040160405180910390fd5b808280631759616b60e11b60205260005b8381101561058057823582018035803b61048257632f8aeb3960e11b6000528060045260246000fd5b60a08201356020810260c0018060808501351460a0606086013514168185013583141615905080156104bf57633ae8821360e21b60005260046000fd5b506020860195506080602084016024376040810260400190508060a00160a45260008160c401528060c4018160a0850160c4376000808260206000875af1935083610571573d1561054f576020601f3d0104915060208104826003028184111561053757818403600302610200838002868002030401015b5a60208201101561054c573d6000803e3d6000fd5b50505b6357e222f160e11b6000528260045260c0606452608451602001608452806000fd5b50505050600181019050610459565b5050505060806040525050565b6040516323b872dd60e01b600052836004528260245281604452602060006064600080895af1803d15601f3d116001600051141617163d151581166106865780873b15151661068657806106715781610650573d1561062a576020601f3d010460208404816003028183111561061157818303600302610200838002858002030401015b5a602082011015610626573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005286600452856024528460445260006064528360845260a46000fd5b639889192360e01b6000528660045285602452846044528360645260846000fd5b632f8aeb3960e11b6000528660045260246000fd5b5050604052505060006060525050565b833b6106b157632f8aeb3960e11b6000528360045260246000fd5b6040516323b872dd60e01b6000528360045282602452816044526000806064600080895af18061074b573d15610725576020601f3d010460208304816003028183111561070c57818303600302610200838002858002030401015b5a602082011015610721573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005285600452846024528360445282606452600160845260a46000fd5b50604052505060006060525050565b843b61077557632f8aeb3960e11b6000528460045260246000fd5b60405160805160a05160c051637921219560e11b6000528760045286602452856044528460645260a0608452600060a45260008060c46000808d5af180610825573d15610800576020601f3d01046020860481600302818311156107e757818303600302610200838002858002030401015b5a6020820110156107fc573d6000803e3d6000fd5b5050505b63f486bc8760e01b600052896004528860245287604452866064528560845260a46000fd5b5060809290925260a05260c05260405250506000606052505050565b60008083601f84011261085357600080fd5b50813567ffffffffffffffff81111561086b57600080fd5b60208301915083602060c08302850101111561088657600080fd5b9250929050565b600080602083850312156108a057600080fd5b823567ffffffffffffffff8111156108b757600080fd5b6108c385828601610841565b90969095509350505050565b60008083601f8401126108e157600080fd5b50813567ffffffffffffffff8111156108f957600080fd5b6020830191508360208260051b850101111561088657600080fd5b6000806000806040858703121561092a57600080fd5b843567ffffffffffffffff8082111561094257600080fd5b61094e88838901610841565b9096509450602087013591508082111561096757600080fd5b50610974878288016108cf565b95989497509550505050565b6000806020838503121561099357600080fd5b823567ffffffffffffffff8111156109aa57600080fd5b6108c3858286016108cf565b80356001600160a01b03811681146109cd57600080fd5b919050565b600080604083850312156109e557600080fd5b6109ee836109b6565b915060208301358015158114610a0357600080fd5b809150509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b600060208284031215610a4c57600080fd5b813560048110610a5b57600080fd5b9392505050565b600060208284031215610a7457600080fd5b610a5b826109b656fea264697066735822122060c22dd101dc32b272f2268d391e21d76db36256cb9c3c5d4b94b6e5bb0dcb7a64736f6c634300080d0033a264697066735822122083f840381e6b0638b7d0b1ce6cee00aae21ac99466b131b4d210f606ca6d674364736f6c634300080d003360a060405234801561001057600080fd5b5033608052608051610ab361003060003960006101e90152610ab36000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80634ce34aa214610051578063899e104c146100815780638df25d9214610094578063c4e8fcb5146100a7575b600080fd5b61006461005f36600461088d565b6100bc565b6040516001600160e01b0319909116815260200160405180910390f35b61006461008f366004610914565b61012b565b6100646100a2366004610980565b61019b565b6100ba6100b53660046109d2565b6101de565b005b60003360005260006020526040600020546100e6576349ed56f960e11b6000523360045260246000fd5b8160005b8181101561011a5761011285858381811061010757610107610a0e565b905060c002016102dc565b6001016100ea565b50632671a55160e11b949350505050565b6000336000526000602052604060002054610155576349ed56f960e11b6000523360045260246000fd5b8360005b8181101561017e5761017687878381811061010757610107610a0e565b600101610159565b506101898484610448565b50632267841360e21b95945050505050565b60003360005260006020526040600020546101c5576349ed56f960e11b6000523360045260246000fd5b6101cf8383610448565b506346f92ec960e11b92915050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610227576040516336abb4df60e11b815260040160405180910390fd5b6001600160a01b03821660009081526020819052604090205481151560ff90911615150361027f576040516349271a0f60e11b81526001600160a01b0383166004820152811515602482015260440160405180910390fd5b6001600160a01b03821660008181526020818152604091829020805460ff191685151590811790915591519182527fae63067d43ac07563b7eb8db6595635fc77f1578a2a5ea06ba91b63e2afa37e2910160405180910390a25050565b60016102eb6020830183610a3a565b60038111156102fc576102fc610a24565b036103415761033e6103146040830160208401610a62565b6103246060840160408501610a62565b6103346080850160608601610a62565b8460a0013561058d565b50565b60026103506020830183610a3a565b600381111561036157610361610a24565b036103c8578060a0013560011461038b5760405163efcc00b160e01b815260040160405180910390fd5b61033e61039e6040830160208401610a62565b6103ae6060840160408501610a62565b6103be6080850160608601610a62565b8460800135610696565b60036103d76020830183610a3a565b60038111156103e8576103e8610a24565b0361042f5761033e6104006040830160208401610a62565b6104106060840160408501610a62565b6104206080850160608601610a62565b84608001358560a0013561075a565b604051631e4cbc7f60e21b815260040160405180910390fd5b808280631759616b60e11b60205260005b8381101561058057823582018035803b61048257632f8aeb3960e11b6000528060045260246000fd5b60a08201356020810260c0018060808501351460a0606086013514168185013583141615905080156104bf57633ae8821360e21b60005260046000fd5b506020860195506080602084016024376040810260400190508060a00160a45260008160c401528060c4018160a0850160c4376000808260206000875af1935083610571573d1561054f576020601f3d0104915060208104826003028184111561053757818403600302610200838002868002030401015b5a60208201101561054c573d6000803e3d6000fd5b50505b6357e222f160e11b6000528260045260c0606452608451602001608452806000fd5b50505050600181019050610459565b5050505060806040525050565b6040516323b872dd60e01b600052836004528260245281604452602060006064600080895af1803d15601f3d116001600051141617163d151581166106865780873b15151661068657806106715781610650573d1561062a576020601f3d010460208404816003028183111561061157818303600302610200838002858002030401015b5a602082011015610626573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005286600452856024528460445260006064528360845260a46000fd5b639889192360e01b6000528660045285602452846044528360645260846000fd5b632f8aeb3960e11b6000528660045260246000fd5b5050604052505060006060525050565b833b6106b157632f8aeb3960e11b6000528360045260246000fd5b6040516323b872dd60e01b6000528360045282602452816044526000806064600080895af18061074b573d15610725576020601f3d010460208304816003028183111561070c57818303600302610200838002858002030401015b5a602082011015610721573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005285600452846024528360445282606452600160845260a46000fd5b50604052505060006060525050565b843b61077557632f8aeb3960e11b6000528460045260246000fd5b60405160805160a05160c051637921219560e11b6000528760045286602452856044528460645260a0608452600060a45260008060c46000808d5af180610825573d15610800576020601f3d01046020860481600302818311156107e757818303600302610200838002858002030401015b5a6020820110156107fc573d6000803e3d6000fd5b5050505b63f486bc8760e01b600052896004528860245287604452866064528560845260a46000fd5b5060809290925260a05260c05260405250506000606052505050565b60008083601f84011261085357600080fd5b50813567ffffffffffffffff81111561086b57600080fd5b60208301915083602060c08302850101111561088657600080fd5b9250929050565b600080602083850312156108a057600080fd5b823567ffffffffffffffff8111156108b757600080fd5b6108c385828601610841565b90969095509350505050565b60008083601f8401126108e157600080fd5b50813567ffffffffffffffff8111156108f957600080fd5b6020830191508360208260051b850101111561088657600080fd5b6000806000806040858703121561092a57600080fd5b843567ffffffffffffffff8082111561094257600080fd5b61094e88838901610841565b9096509450602087013591508082111561096757600080fd5b50610974878288016108cf565b95989497509550505050565b6000806020838503121561099357600080fd5b823567ffffffffffffffff8111156109aa57600080fd5b6108c3858286016108cf565b80356001600160a01b03811681146109cd57600080fd5b919050565b600080604083850312156109e557600080fd5b6109ee836109b6565b915060208301358015158114610a0357600080fd5b809150509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b600060208284031215610a4c57600080fd5b813560048110610a5b57600080fd5b9392505050565b600060208284031215610a7457600080fd5b610a5b826109b656fea264697066735822122060c22dd101dc32b272f2268d391e21d76db36256cb9c3c5d4b94b6e5bb0dcb7a64736f6c634300080d0033",
}

// ConduitControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use ConduitControllerMetaData.ABI instead.
var ConduitControllerABI = ConduitControllerMetaData.ABI

// ConduitControllerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConduitControllerMetaData.Bin instead.
var ConduitControllerBin = ConduitControllerMetaData.Bin

// DeployConduitController deploys a new Ethereum contract, binding an instance of ConduitController to it.
func DeployConduitController(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ConduitController, error) {
	parsed, err := ConduitControllerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConduitControllerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ConduitController{ConduitControllerCaller: ConduitControllerCaller{contract: contract}, ConduitControllerTransactor: ConduitControllerTransactor{contract: contract}, ConduitControllerFilterer: ConduitControllerFilterer{contract: contract}}, nil
}

// ConduitController is an auto generated Go binding around an Ethereum contract.
type ConduitController struct {
	ConduitControllerCaller     // Read-only binding to the contract
	ConduitControllerTransactor // Write-only binding to the contract
	ConduitControllerFilterer   // Log filterer for contract events
}

// ConduitControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConduitControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConduitControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConduitControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConduitControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConduitControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConduitControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConduitControllerSession struct {
	Contract     *ConduitController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ConduitControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConduitControllerCallerSession struct {
	Contract *ConduitControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ConduitControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConduitControllerTransactorSession struct {
	Contract     *ConduitControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ConduitControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConduitControllerRaw struct {
	Contract *ConduitController // Generic contract binding to access the raw methods on
}

// ConduitControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConduitControllerCallerRaw struct {
	Contract *ConduitControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ConduitControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConduitControllerTransactorRaw struct {
	Contract *ConduitControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConduitController creates a new instance of ConduitController, bound to a specific deployed contract.
func NewConduitController(address common.Address, backend bind.ContractBackend) (*ConduitController, error) {
	contract, err := bindConduitController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConduitController{ConduitControllerCaller: ConduitControllerCaller{contract: contract}, ConduitControllerTransactor: ConduitControllerTransactor{contract: contract}, ConduitControllerFilterer: ConduitControllerFilterer{contract: contract}}, nil
}

// NewConduitControllerCaller creates a new read-only instance of ConduitController, bound to a specific deployed contract.
func NewConduitControllerCaller(address common.Address, caller bind.ContractCaller) (*ConduitControllerCaller, error) {
	contract, err := bindConduitController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConduitControllerCaller{contract: contract}, nil
}

// NewConduitControllerTransactor creates a new write-only instance of ConduitController, bound to a specific deployed contract.
func NewConduitControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ConduitControllerTransactor, error) {
	contract, err := bindConduitController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConduitControllerTransactor{contract: contract}, nil
}

// NewConduitControllerFilterer creates a new log filterer instance of ConduitController, bound to a specific deployed contract.
func NewConduitControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ConduitControllerFilterer, error) {
	contract, err := bindConduitController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConduitControllerFilterer{contract: contract}, nil
}

// bindConduitController binds a generic wrapper to an already deployed contract.
func bindConduitController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConduitControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConduitController *ConduitControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConduitController.Contract.ConduitControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConduitController *ConduitControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConduitController.Contract.ConduitControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConduitController *ConduitControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConduitController.Contract.ConduitControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConduitController *ConduitControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConduitController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConduitController *ConduitControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConduitController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConduitController *ConduitControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConduitController.Contract.contract.Transact(opts, method, params...)
}

// GetChannel is a free data retrieval call binding the contract method 0x027cc764.
//
// Solidity: function getChannel(address conduit, uint256 channelIndex) view returns(address channel)
func (_ConduitController *ConduitControllerCaller) GetChannel(opts *bind.CallOpts, conduit common.Address, channelIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConduitController.contract.Call(opts, &out, "getChannel", conduit, channelIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetChannel is a free data retrieval call binding the contract method 0x027cc764.
//
// Solidity: function getChannel(address conduit, uint256 channelIndex) view returns(address channel)
func (_ConduitController *ConduitControllerSession) GetChannel(conduit common.Address, channelIndex *big.Int) (common.Address, error) {
	return _ConduitController.Contract.GetChannel(&_ConduitController.CallOpts, conduit, channelIndex)
}

// GetChannel is a free data retrieval call binding the contract method 0x027cc764.
//
// Solidity: function getChannel(address conduit, uint256 channelIndex) view returns(address channel)
func (_ConduitController *ConduitControllerCallerSession) GetChannel(conduit common.Address, channelIndex *big.Int) (common.Address, error) {
	return _ConduitController.Contract.GetChannel(&_ConduitController.CallOpts, conduit, channelIndex)
}

// GetChannelStatus is a free data retrieval call binding the contract method 0x33bc8572.
//
// Solidity: function getChannelStatus(address conduit, address channel) view returns(bool isOpen)
func (_ConduitController *ConduitControllerCaller) GetChannelStatus(opts *bind.CallOpts, conduit common.Address, channel common.Address) (bool, error) {
	var out []interface{}
	err := _ConduitController.contract.Call(opts, &out, "getChannelStatus", conduit, channel)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetChannelStatus is a free data retrieval call binding the contract method 0x33bc8572.
//
// Solidity: function getChannelStatus(address conduit, address channel) view returns(bool isOpen)
func (_ConduitController *ConduitControllerSession) GetChannelStatus(conduit common.Address, channel common.Address) (bool, error) {
	return _ConduitController.Contract.GetChannelStatus(&_ConduitController.CallOpts, conduit, channel)
}

// GetChannelStatus is a free data retrieval call binding the contract method 0x33bc8572.
//
// Solidity: function getChannelStatus(address conduit, address channel) view returns(bool isOpen)
func (_ConduitController *ConduitControllerCallerSession) GetChannelStatus(conduit common.Address, channel common.Address) (bool, error) {
	return _ConduitController.Contract.GetChannelStatus(&_ConduitController.CallOpts, conduit, channel)
}

// GetChannels is a free data retrieval call binding the contract method 0x8b9e028b.
//
// Solidity: function getChannels(address conduit) view returns(address[] channels)
func (_ConduitController *ConduitControllerCaller) GetChannels(opts *bind.CallOpts, conduit common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ConduitController.contract.Call(opts, &out, "getChannels", conduit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetChannels is a free data retrieval call binding the contract method 0x8b9e028b.
//
// Solidity: function getChannels(address conduit) view returns(address[] channels)
func (_ConduitController *ConduitControllerSession) GetChannels(conduit common.Address) ([]common.Address, error) {
	return _ConduitController.Contract.GetChannels(&_ConduitController.CallOpts, conduit)
}

// GetChannels is a free data retrieval call binding the contract method 0x8b9e028b.
//
// Solidity: function getChannels(address conduit) view returns(address[] channels)
func (_ConduitController *ConduitControllerCallerSession) GetChannels(conduit common.Address) ([]common.Address, error) {
	return _ConduitController.Contract.GetChannels(&_ConduitController.CallOpts, conduit)
}

// GetConduit is a free data retrieval call binding the contract method 0x6e9bfd9f.
//
// Solidity: function getConduit(bytes32 conduitKey) view returns(address conduit, bool exists)
func (_ConduitController *ConduitControllerCaller) GetConduit(opts *bind.CallOpts, conduitKey [32]byte) (struct {
	Conduit common.Address
	Exists  bool
}, error) {
	var out []interface{}
	err := _ConduitController.contract.Call(opts, &out, "getConduit", conduitKey)

	outstruct := new(struct {
		Conduit common.Address
		Exists  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Conduit = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Exists = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetConduit is a free data retrieval call binding the contract method 0x6e9bfd9f.
//
// Solidity: function getConduit(bytes32 conduitKey) view returns(address conduit, bool exists)
func (_ConduitController *ConduitControllerSession) GetConduit(conduitKey [32]byte) (struct {
	Conduit common.Address
	Exists  bool
}, error) {
	return _ConduitController.Contract.GetConduit(&_ConduitController.CallOpts, conduitKey)
}

// GetConduit is a free data retrieval call binding the contract method 0x6e9bfd9f.
//
// Solidity: function getConduit(bytes32 conduitKey) view returns(address conduit, bool exists)
func (_ConduitController *ConduitControllerCallerSession) GetConduit(conduitKey [32]byte) (struct {
	Conduit common.Address
	Exists  bool
}, error) {
	return _ConduitController.Contract.GetConduit(&_ConduitController.CallOpts, conduitKey)
}

// GetConduitCodeHashes is a free data retrieval call binding the contract method 0x0a96ad39.
//
// Solidity: function getConduitCodeHashes() view returns(bytes32 creationCodeHash, bytes32 runtimeCodeHash)
func (_ConduitController *ConduitControllerCaller) GetConduitCodeHashes(opts *bind.CallOpts) (struct {
	CreationCodeHash [32]byte
	RuntimeCodeHash  [32]byte
}, error) {
	var out []interface{}
	err := _ConduitController.contract.Call(opts, &out, "getConduitCodeHashes")

	outstruct := new(struct {
		CreationCodeHash [32]byte
		RuntimeCodeHash  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CreationCodeHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.RuntimeCodeHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetConduitCodeHashes is a free data retrieval call binding the contract method 0x0a96ad39.
//
// Solidity: function getConduitCodeHashes() view returns(bytes32 creationCodeHash, bytes32 runtimeCodeHash)
func (_ConduitController *ConduitControllerSession) GetConduitCodeHashes() (struct {
	CreationCodeHash [32]byte
	RuntimeCodeHash  [32]byte
}, error) {
	return _ConduitController.Contract.GetConduitCodeHashes(&_ConduitController.CallOpts)
}

// GetConduitCodeHashes is a free data retrieval call binding the contract method 0x0a96ad39.
//
// Solidity: function getConduitCodeHashes() view returns(bytes32 creationCodeHash, bytes32 runtimeCodeHash)
func (_ConduitController *ConduitControllerCallerSession) GetConduitCodeHashes() (struct {
	CreationCodeHash [32]byte
	RuntimeCodeHash  [32]byte
}, error) {
	return _ConduitController.Contract.GetConduitCodeHashes(&_ConduitController.CallOpts)
}

// GetKey is a free data retrieval call binding the contract method 0x93790f44.
//
// Solidity: function getKey(address conduit) view returns(bytes32 conduitKey)
func (_ConduitController *ConduitControllerCaller) GetKey(opts *bind.CallOpts, conduit common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ConduitController.contract.Call(opts, &out, "getKey", conduit)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKey is a free data retrieval call binding the contract method 0x93790f44.
//
// Solidity: function getKey(address conduit) view returns(bytes32 conduitKey)
func (_ConduitController *ConduitControllerSession) GetKey(conduit common.Address) ([32]byte, error) {
	return _ConduitController.Contract.GetKey(&_ConduitController.CallOpts, conduit)
}

// GetKey is a free data retrieval call binding the contract method 0x93790f44.
//
// Solidity: function getKey(address conduit) view returns(bytes32 conduitKey)
func (_ConduitController *ConduitControllerCallerSession) GetKey(conduit common.Address) ([32]byte, error) {
	return _ConduitController.Contract.GetKey(&_ConduitController.CallOpts, conduit)
}

// GetPotentialOwner is a free data retrieval call binding the contract method 0x906c87cc.
//
// Solidity: function getPotentialOwner(address conduit) view returns(address potentialOwner)
func (_ConduitController *ConduitControllerCaller) GetPotentialOwner(opts *bind.CallOpts, conduit common.Address) (common.Address, error) {
	var out []interface{}
	err := _ConduitController.contract.Call(opts, &out, "getPotentialOwner", conduit)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPotentialOwner is a free data retrieval call binding the contract method 0x906c87cc.
//
// Solidity: function getPotentialOwner(address conduit) view returns(address potentialOwner)
func (_ConduitController *ConduitControllerSession) GetPotentialOwner(conduit common.Address) (common.Address, error) {
	return _ConduitController.Contract.GetPotentialOwner(&_ConduitController.CallOpts, conduit)
}

// GetPotentialOwner is a free data retrieval call binding the contract method 0x906c87cc.
//
// Solidity: function getPotentialOwner(address conduit) view returns(address potentialOwner)
func (_ConduitController *ConduitControllerCallerSession) GetPotentialOwner(conduit common.Address) (common.Address, error) {
	return _ConduitController.Contract.GetPotentialOwner(&_ConduitController.CallOpts, conduit)
}

// GetTotalChannels is a free data retrieval call binding the contract method 0x4e3f9580.
//
// Solidity: function getTotalChannels(address conduit) view returns(uint256 totalChannels)
func (_ConduitController *ConduitControllerCaller) GetTotalChannels(opts *bind.CallOpts, conduit common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConduitController.contract.Call(opts, &out, "getTotalChannels", conduit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalChannels is a free data retrieval call binding the contract method 0x4e3f9580.
//
// Solidity: function getTotalChannels(address conduit) view returns(uint256 totalChannels)
func (_ConduitController *ConduitControllerSession) GetTotalChannels(conduit common.Address) (*big.Int, error) {
	return _ConduitController.Contract.GetTotalChannels(&_ConduitController.CallOpts, conduit)
}

// GetTotalChannels is a free data retrieval call binding the contract method 0x4e3f9580.
//
// Solidity: function getTotalChannels(address conduit) view returns(uint256 totalChannels)
func (_ConduitController *ConduitControllerCallerSession) GetTotalChannels(conduit common.Address) (*big.Int, error) {
	return _ConduitController.Contract.GetTotalChannels(&_ConduitController.CallOpts, conduit)
}

// OwnerOf is a free data retrieval call binding the contract method 0x14afd79e.
//
// Solidity: function ownerOf(address conduit) view returns(address owner)
func (_ConduitController *ConduitControllerCaller) OwnerOf(opts *bind.CallOpts, conduit common.Address) (common.Address, error) {
	var out []interface{}
	err := _ConduitController.contract.Call(opts, &out, "ownerOf", conduit)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x14afd79e.
//
// Solidity: function ownerOf(address conduit) view returns(address owner)
func (_ConduitController *ConduitControllerSession) OwnerOf(conduit common.Address) (common.Address, error) {
	return _ConduitController.Contract.OwnerOf(&_ConduitController.CallOpts, conduit)
}

// OwnerOf is a free data retrieval call binding the contract method 0x14afd79e.
//
// Solidity: function ownerOf(address conduit) view returns(address owner)
func (_ConduitController *ConduitControllerCallerSession) OwnerOf(conduit common.Address) (common.Address, error) {
	return _ConduitController.Contract.OwnerOf(&_ConduitController.CallOpts, conduit)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x51710e45.
//
// Solidity: function acceptOwnership(address conduit) returns()
func (_ConduitController *ConduitControllerTransactor) AcceptOwnership(opts *bind.TransactOpts, conduit common.Address) (*types.Transaction, error) {
	return _ConduitController.contract.Transact(opts, "acceptOwnership", conduit)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x51710e45.
//
// Solidity: function acceptOwnership(address conduit) returns()
func (_ConduitController *ConduitControllerSession) AcceptOwnership(conduit common.Address) (*types.Transaction, error) {
	return _ConduitController.Contract.AcceptOwnership(&_ConduitController.TransactOpts, conduit)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x51710e45.
//
// Solidity: function acceptOwnership(address conduit) returns()
func (_ConduitController *ConduitControllerTransactorSession) AcceptOwnership(conduit common.Address) (*types.Transaction, error) {
	return _ConduitController.Contract.AcceptOwnership(&_ConduitController.TransactOpts, conduit)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x7b37e561.
//
// Solidity: function cancelOwnershipTransfer(address conduit) returns()
func (_ConduitController *ConduitControllerTransactor) CancelOwnershipTransfer(opts *bind.TransactOpts, conduit common.Address) (*types.Transaction, error) {
	return _ConduitController.contract.Transact(opts, "cancelOwnershipTransfer", conduit)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x7b37e561.
//
// Solidity: function cancelOwnershipTransfer(address conduit) returns()
func (_ConduitController *ConduitControllerSession) CancelOwnershipTransfer(conduit common.Address) (*types.Transaction, error) {
	return _ConduitController.Contract.CancelOwnershipTransfer(&_ConduitController.TransactOpts, conduit)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x7b37e561.
//
// Solidity: function cancelOwnershipTransfer(address conduit) returns()
func (_ConduitController *ConduitControllerTransactorSession) CancelOwnershipTransfer(conduit common.Address) (*types.Transaction, error) {
	return _ConduitController.Contract.CancelOwnershipTransfer(&_ConduitController.TransactOpts, conduit)
}

// CreateConduit is a paid mutator transaction binding the contract method 0x794593bc.
//
// Solidity: function createConduit(bytes32 conduitKey, address initialOwner) returns(address conduit)
func (_ConduitController *ConduitControllerTransactor) CreateConduit(opts *bind.TransactOpts, conduitKey [32]byte, initialOwner common.Address) (*types.Transaction, error) {
	return _ConduitController.contract.Transact(opts, "createConduit", conduitKey, initialOwner)
}

// CreateConduit is a paid mutator transaction binding the contract method 0x794593bc.
//
// Solidity: function createConduit(bytes32 conduitKey, address initialOwner) returns(address conduit)
func (_ConduitController *ConduitControllerSession) CreateConduit(conduitKey [32]byte, initialOwner common.Address) (*types.Transaction, error) {
	return _ConduitController.Contract.CreateConduit(&_ConduitController.TransactOpts, conduitKey, initialOwner)
}

// CreateConduit is a paid mutator transaction binding the contract method 0x794593bc.
//
// Solidity: function createConduit(bytes32 conduitKey, address initialOwner) returns(address conduit)
func (_ConduitController *ConduitControllerTransactorSession) CreateConduit(conduitKey [32]byte, initialOwner common.Address) (*types.Transaction, error) {
	return _ConduitController.Contract.CreateConduit(&_ConduitController.TransactOpts, conduitKey, initialOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address conduit, address newPotentialOwner) returns()
func (_ConduitController *ConduitControllerTransactor) TransferOwnership(opts *bind.TransactOpts, conduit common.Address, newPotentialOwner common.Address) (*types.Transaction, error) {
	return _ConduitController.contract.Transact(opts, "transferOwnership", conduit, newPotentialOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address conduit, address newPotentialOwner) returns()
func (_ConduitController *ConduitControllerSession) TransferOwnership(conduit common.Address, newPotentialOwner common.Address) (*types.Transaction, error) {
	return _ConduitController.Contract.TransferOwnership(&_ConduitController.TransactOpts, conduit, newPotentialOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address conduit, address newPotentialOwner) returns()
func (_ConduitController *ConduitControllerTransactorSession) TransferOwnership(conduit common.Address, newPotentialOwner common.Address) (*types.Transaction, error) {
	return _ConduitController.Contract.TransferOwnership(&_ConduitController.TransactOpts, conduit, newPotentialOwner)
}

// UpdateChannel is a paid mutator transaction binding the contract method 0x13ad9cab.
//
// Solidity: function updateChannel(address conduit, address channel, bool isOpen) returns()
func (_ConduitController *ConduitControllerTransactor) UpdateChannel(opts *bind.TransactOpts, conduit common.Address, channel common.Address, isOpen bool) (*types.Transaction, error) {
	return _ConduitController.contract.Transact(opts, "updateChannel", conduit, channel, isOpen)
}

// UpdateChannel is a paid mutator transaction binding the contract method 0x13ad9cab.
//
// Solidity: function updateChannel(address conduit, address channel, bool isOpen) returns()
func (_ConduitController *ConduitControllerSession) UpdateChannel(conduit common.Address, channel common.Address, isOpen bool) (*types.Transaction, error) {
	return _ConduitController.Contract.UpdateChannel(&_ConduitController.TransactOpts, conduit, channel, isOpen)
}

// UpdateChannel is a paid mutator transaction binding the contract method 0x13ad9cab.
//
// Solidity: function updateChannel(address conduit, address channel, bool isOpen) returns()
func (_ConduitController *ConduitControllerTransactorSession) UpdateChannel(conduit common.Address, channel common.Address, isOpen bool) (*types.Transaction, error) {
	return _ConduitController.Contract.UpdateChannel(&_ConduitController.TransactOpts, conduit, channel, isOpen)
}

// ConduitControllerNewConduitIterator is returned from FilterNewConduit and is used to iterate over the raw logs and unpacked data for NewConduit events raised by the ConduitController contract.
type ConduitControllerNewConduitIterator struct {
	Event *ConduitControllerNewConduit // Event containing the contract specifics and raw log

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
func (it *ConduitControllerNewConduitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConduitControllerNewConduit)
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
		it.Event = new(ConduitControllerNewConduit)
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
func (it *ConduitControllerNewConduitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConduitControllerNewConduitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConduitControllerNewConduit represents a NewConduit event raised by the ConduitController contract.
type ConduitControllerNewConduit struct {
	Conduit    common.Address
	ConduitKey [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewConduit is a free log retrieval operation binding the contract event 0x4397af6128d529b8ae0442f99db1296d5136062597a15bbc61c1b2a6431a7d15.
//
// Solidity: event NewConduit(address conduit, bytes32 conduitKey)
func (_ConduitController *ConduitControllerFilterer) FilterNewConduit(opts *bind.FilterOpts) (*ConduitControllerNewConduitIterator, error) {

	logs, sub, err := _ConduitController.contract.FilterLogs(opts, "NewConduit")
	if err != nil {
		return nil, err
	}
	return &ConduitControllerNewConduitIterator{contract: _ConduitController.contract, event: "NewConduit", logs: logs, sub: sub}, nil
}

// WatchNewConduit is a free log subscription operation binding the contract event 0x4397af6128d529b8ae0442f99db1296d5136062597a15bbc61c1b2a6431a7d15.
//
// Solidity: event NewConduit(address conduit, bytes32 conduitKey)
func (_ConduitController *ConduitControllerFilterer) WatchNewConduit(opts *bind.WatchOpts, sink chan<- *ConduitControllerNewConduit) (event.Subscription, error) {

	logs, sub, err := _ConduitController.contract.WatchLogs(opts, "NewConduit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConduitControllerNewConduit)
				if err := _ConduitController.contract.UnpackLog(event, "NewConduit", log); err != nil {
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

// ParseNewConduit is a log parse operation binding the contract event 0x4397af6128d529b8ae0442f99db1296d5136062597a15bbc61c1b2a6431a7d15.
//
// Solidity: event NewConduit(address conduit, bytes32 conduitKey)
func (_ConduitController *ConduitControllerFilterer) ParseNewConduit(log types.Log) (*ConduitControllerNewConduit, error) {
	event := new(ConduitControllerNewConduit)
	if err := _ConduitController.contract.UnpackLog(event, "NewConduit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConduitControllerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ConduitController contract.
type ConduitControllerOwnershipTransferredIterator struct {
	Event *ConduitControllerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ConduitControllerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConduitControllerOwnershipTransferred)
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
		it.Event = new(ConduitControllerOwnershipTransferred)
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
func (it *ConduitControllerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConduitControllerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConduitControllerOwnershipTransferred represents a OwnershipTransferred event raised by the ConduitController contract.
type ConduitControllerOwnershipTransferred struct {
	Conduit       common.Address
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0xc8894f26f396ce8c004245c8b7cd1b92103a6e4302fcbab883987149ac01b7ec.
//
// Solidity: event OwnershipTransferred(address indexed conduit, address indexed previousOwner, address indexed newOwner)
func (_ConduitController *ConduitControllerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, conduit []common.Address, previousOwner []common.Address, newOwner []common.Address) (*ConduitControllerOwnershipTransferredIterator, error) {

	var conduitRule []interface{}
	for _, conduitItem := range conduit {
		conduitRule = append(conduitRule, conduitItem)
	}
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ConduitController.contract.FilterLogs(opts, "OwnershipTransferred", conduitRule, previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConduitControllerOwnershipTransferredIterator{contract: _ConduitController.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0xc8894f26f396ce8c004245c8b7cd1b92103a6e4302fcbab883987149ac01b7ec.
//
// Solidity: event OwnershipTransferred(address indexed conduit, address indexed previousOwner, address indexed newOwner)
func (_ConduitController *ConduitControllerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ConduitControllerOwnershipTransferred, conduit []common.Address, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var conduitRule []interface{}
	for _, conduitItem := range conduit {
		conduitRule = append(conduitRule, conduitItem)
	}
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ConduitController.contract.WatchLogs(opts, "OwnershipTransferred", conduitRule, previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConduitControllerOwnershipTransferred)
				if err := _ConduitController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0xc8894f26f396ce8c004245c8b7cd1b92103a6e4302fcbab883987149ac01b7ec.
//
// Solidity: event OwnershipTransferred(address indexed conduit, address indexed previousOwner, address indexed newOwner)
func (_ConduitController *ConduitControllerFilterer) ParseOwnershipTransferred(log types.Log) (*ConduitControllerOwnershipTransferred, error) {
	event := new(ConduitControllerOwnershipTransferred)
	if err := _ConduitController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConduitControllerPotentialOwnerUpdatedIterator is returned from FilterPotentialOwnerUpdated and is used to iterate over the raw logs and unpacked data for PotentialOwnerUpdated events raised by the ConduitController contract.
type ConduitControllerPotentialOwnerUpdatedIterator struct {
	Event *ConduitControllerPotentialOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *ConduitControllerPotentialOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConduitControllerPotentialOwnerUpdated)
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
		it.Event = new(ConduitControllerPotentialOwnerUpdated)
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
func (it *ConduitControllerPotentialOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConduitControllerPotentialOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConduitControllerPotentialOwnerUpdated represents a PotentialOwnerUpdated event raised by the ConduitController contract.
type ConduitControllerPotentialOwnerUpdated struct {
	NewPotentialOwner common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPotentialOwnerUpdated is a free log retrieval operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address indexed newPotentialOwner)
func (_ConduitController *ConduitControllerFilterer) FilterPotentialOwnerUpdated(opts *bind.FilterOpts, newPotentialOwner []common.Address) (*ConduitControllerPotentialOwnerUpdatedIterator, error) {

	var newPotentialOwnerRule []interface{}
	for _, newPotentialOwnerItem := range newPotentialOwner {
		newPotentialOwnerRule = append(newPotentialOwnerRule, newPotentialOwnerItem)
	}

	logs, sub, err := _ConduitController.contract.FilterLogs(opts, "PotentialOwnerUpdated", newPotentialOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConduitControllerPotentialOwnerUpdatedIterator{contract: _ConduitController.contract, event: "PotentialOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchPotentialOwnerUpdated is a free log subscription operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address indexed newPotentialOwner)
func (_ConduitController *ConduitControllerFilterer) WatchPotentialOwnerUpdated(opts *bind.WatchOpts, sink chan<- *ConduitControllerPotentialOwnerUpdated, newPotentialOwner []common.Address) (event.Subscription, error) {

	var newPotentialOwnerRule []interface{}
	for _, newPotentialOwnerItem := range newPotentialOwner {
		newPotentialOwnerRule = append(newPotentialOwnerRule, newPotentialOwnerItem)
	}

	logs, sub, err := _ConduitController.contract.WatchLogs(opts, "PotentialOwnerUpdated", newPotentialOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConduitControllerPotentialOwnerUpdated)
				if err := _ConduitController.contract.UnpackLog(event, "PotentialOwnerUpdated", log); err != nil {
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

// ParsePotentialOwnerUpdated is a log parse operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address indexed newPotentialOwner)
func (_ConduitController *ConduitControllerFilterer) ParsePotentialOwnerUpdated(log types.Log) (*ConduitControllerPotentialOwnerUpdated, error) {
	event := new(ConduitControllerPotentialOwnerUpdated)
	if err := _ConduitController.contract.UnpackLog(event, "PotentialOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
