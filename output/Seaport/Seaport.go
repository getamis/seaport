// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Seaport

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

// AdditionalRecipient is an auto generated low-level Go binding around an user-defined struct.
type AdditionalRecipient struct {
	Amount    *big.Int
	Recipient common.Address
}

// AdvancedOrder is an auto generated low-level Go binding around an user-defined struct.
type AdvancedOrder struct {
	Parameters  OrderParameters
	Numerator   *big.Int
	Denominator *big.Int
	Signature   []byte
	ExtraData   []byte
}

// BasicOrderParameters is an auto generated low-level Go binding around an user-defined struct.
type BasicOrderParameters struct {
	ConsiderationToken                common.Address
	ConsiderationIdentifier           *big.Int
	ConsiderationAmount               *big.Int
	Offerer                           common.Address
	Zone                              common.Address
	OfferToken                        common.Address
	OfferIdentifier                   *big.Int
	OfferAmount                       *big.Int
	BasicOrderType                    uint8
	StartTime                         *big.Int
	EndTime                           *big.Int
	ZoneHash                          [32]byte
	Salt                              *big.Int
	OffererConduitKey                 [32]byte
	FulfillerConduitKey               [32]byte
	TotalOriginalAdditionalRecipients *big.Int
	AdditionalRecipients              []AdditionalRecipient
	Signature                         []byte
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

// Execution is an auto generated low-level Go binding around an user-defined struct.
type Execution struct {
	Item       ReceivedItem
	Offerer    common.Address
	ConduitKey [32]byte
}

// Fulfillment is an auto generated low-level Go binding around an user-defined struct.
type Fulfillment struct {
	OfferComponents         []FulfillmentComponent
	ConsiderationComponents []FulfillmentComponent
}

// FulfillmentComponent is an auto generated low-level Go binding around an user-defined struct.
type FulfillmentComponent struct {
	OrderIndex *big.Int
	ItemIndex  *big.Int
}

// OfferItem is an auto generated low-level Go binding around an user-defined struct.
type OfferItem struct {
	ItemType             uint8
	Token                common.Address
	IdentifierOrCriteria *big.Int
	StartAmount          *big.Int
	EndAmount            *big.Int
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Parameters OrderParameters
	Signature  []byte
}

// OrderComponents is an auto generated low-level Go binding around an user-defined struct.
type OrderComponents struct {
	Offerer       common.Address
	Zone          common.Address
	Offer         []OfferItem
	Consideration []ConsiderationItem
	OrderType     uint8
	StartTime     *big.Int
	EndTime       *big.Int
	ZoneHash      [32]byte
	Salt          *big.Int
	ConduitKey    [32]byte
	Nonce         *big.Int
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

// SeaportMetaData contains all meta data concerning the Seaport contract.
var SeaportMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduitController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadContractSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadFraction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BadReturnValueFromERC20OnTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"BadSignatureV\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConsiderationCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortfallAmount\",\"type\":\"uint256\"}],\"name\":\"ConsiderationNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CriteriaNotEnabledForItem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"identifiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"ERC1155BatchTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EtherTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InexactFraction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientEtherSupplied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBasicOrderParameterEncoding\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidCallToConduit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCanceller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidConduit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidERC721TransferAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFulfillmentComponentData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidRestrictedOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchedFulfillmentOfferAndConsiderationComponents\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"}],\"name\":\"MissingFulfillmentComponentOnAggregation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingItemAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingOriginalConsiderationItems\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NoContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoReentrantCalls\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSpecifiedOrdersAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferAndConsiderationRequiredOnFulfillment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderAlreadyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderIsCancelled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderPartiallyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PartialFillsNotEnabledForOrder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnresolvedConsiderationCriteria\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnresolvedOfferCriteria\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"}],\"name\":\"NonceIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fulfiller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSpentItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structReceivedItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"}],\"name\":\"OrderValidated\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder\",\"name\":\"advancedOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"criteriaResolvers\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"}],\"name\":\"fulfillAdvancedOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder[]\",\"name\":\"advancedOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"criteriaResolvers\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"offerFulfillments\",\"type\":\"tuple[][]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"considerationFulfillments\",\"type\":\"tuple[][]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maximumFulfilled\",\"type\":\"uint256\"}],\"name\":\"fulfillAvailableAdvancedOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"availableOrders\",\"type\":\"bool[]\"},{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"offerFulfillments\",\"type\":\"tuple[][]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"considerationFulfillments\",\"type\":\"tuple[][]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maximumFulfilled\",\"type\":\"uint256\"}],\"name\":\"fulfillAvailableOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"availableOrders\",\"type\":\"bool[]\"},{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"considerationToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"considerationIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"offerToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offerIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumBasicOrderType\",\"name\":\"basicOrderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offererConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalAdditionalRecipients\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structAdditionalRecipient[]\",\"name\":\"additionalRecipients\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBasicOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"}],\"name\":\"fulfillBasicOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"}],\"name\":\"fulfillOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"getOrderStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidated\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSize\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"information\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"conduitController\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder[]\",\"name\":\"advancedOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"criteriaResolvers\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"offerComponents\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"considerationComponents\",\"type\":\"tuple[]\"}],\"internalType\":\"structFulfillment[]\",\"name\":\"fulfillments\",\"type\":\"tuple[]\"}],\"name\":\"matchAdvancedOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"offerComponents\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"considerationComponents\",\"type\":\"tuple[]\"}],\"internalType\":\"structFulfillment[]\",\"name\":\"fulfillments\",\"type\":\"tuple[]\"}],\"name\":\"matchOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"contractName\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"validate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"validated\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101c06040523480156200001257600080fd5b5060405162005b6438038062005b648339810160408190526200003591620005c2565b808080808080808080806200004962000132565b610120526101005260e05260c081815260a0838152608085815246610140819052604080516020818101979097528082019890985260608801969096529086015230858201528351808603909101815293019091528151910120610160526001600160a01b03811661018081905260408051630a96ad3960e01b81528151630a96ad39926004808401939192918290030181865afa158015620000f0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001169190620005f4565b506101a052505060016000555062000684975050505050505050565b600080808080806200015e60408051808201909152600781526614d9585c1bdc9d60ca1b602082015290565b805160209182012060408051808201825260018152603160f81b90840152519097507fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc696506000916200026191016909ecccccae492e8cada560b31b81526e1d5a5b9d0e081a5d195b551e5c194b608a1b600a8201526d1859191c995cdcc81d1bdad95b8b60921b60198201527f75696e74323536206964656e7469666965724f7243726974657269612c00000060278201527f75696e74323536207374617274416d6f756e742c0000000000000000000000006044820152701d5a5b9d0c8d4d88195b99105b5bdd5b9d607a1b6058820152602960f81b6069820152606a0190565b60408051601f1981840301815282825271086dedce6d2c8cae4c2e8d2dedc92e8cada560731b60208401526e1d5a5b9d0e081a5d195b551e5c194b608a1b60328401526d1859191c995cdcc81d1bdad95b8b60921b60418401527f75696e74323536206964656e7469666965724f7243726974657269612c000000604f8401527f75696e74323536207374617274416d6f756e742c000000000000000000000000606c840152711d5a5b9d0c8d4d88195b99105b5bdd5b9d0b60721b6080840152701859191c995cdcc81c9958da5c1a595b9d607a1b6092840152602960f81b60a384018190528251808503608401815260a485019093526f09ee4c8cae486dedae0dedccadce8e6560831b60c48501526f1859191c995cdcc81bd999995c995c8b60821b60d48501526c1859191c995cdcc81e9bdb994b609a1b60e48501527113d999995c925d195b56d7481bd999995c8b60721b60f18501527f436f6e73696465726174696f6e4974656d5b5d20636f6e73696465726174696f610103850152611b8b60f21b6101238501526f1d5a5b9d0e081bdc99195c951e5c194b60821b610125850152711d5a5b9d0c8d4d881cdd185c9d151a5b594b60721b6101358501526f1d5a5b9d0c8d4d88195b99151a5b594b60821b61014785015270189e5d195ccccc881e9bdb9952185cda0b607a1b6101578501526c1d5a5b9d0c8d4d881cd85b1d0b609a1b6101688501527f6279746573333220636f6e647569744b65792c000000000000000000000000006101758501526c75696e74323536206e6f6e636560981b6101888501526101958401529250906000906101960160408051601f19818403018152908290526c08a92a06e626488dedac2d2dc5609b1b60208301526b1cdd1c9a5b99c81b985b594b60a21b602d8301526e1cdd1c9a5b99c81d995c9cda5bdb8b608a1b60398301526f1d5a5b9d0c8d4d8818da185a5b92590b60821b60488301527f6164647265737320766572696679696e67436f6e7472616374000000000000006058830152602960f81b6071830152915060720160408051601f19818403018152908290528051602091820120855186830120855186840120919a50985096506200059f918391859187910162000657565b604051602081830303815290604052805190602001209350505050909192939495565b600060208284031215620005d557600080fd5b81516001600160a01b0381168114620005ed57600080fd5b9392505050565b600080604083850312156200060857600080fd5b505080516020909101519092909150565b6000815160005b818110156200063c576020818501810151868301520162000620565b818111156200064c576000828601525b509290920192915050565b60006200067b620006746200066d848862000619565b8662000619565b8462000619565b95945050505050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516154486200071c600039600061353f015260008181610d6c015261351301526000612323015260006122530152600081816108a3015261253b01526000818161083201526123990152600081816107cb01526124b401526000612281015260006122cf015260006122a701526154486000f3fe6080604052600436106100e85760003560e01c8063a81744041161008a578063f47b774011610059578063f47b7740146102c7578063fb0f3ee1146102eb578063fb4c2af9146102fe578063fd9f1e101461031157600080fd5b8063a81744041461026d578063b3a34c4c14610280578063df7b0dac14610293578063ed98a574146102a657600080fd5b806355944a42116100c657806355944a42146101e8578063627cdcb91461020857806379df72bd1461021d578063881477321461023d57600080fd5b806306fdde03146100ed5780632d0335ab1461011857806346423aa714610146575b600080fd5b3480156100f957600080fd5b50610102610331565b60405161010f9190613fde565b60405180910390f35b34801561012457600080fd5b50610138610133366004614016565b610340565b60405190815260200161010f565b34801561015257600080fd5b506101c6610161366004614033565b6000908152600260209081526040918290208251608081018452905460ff8082161515808452610100830490911615159383018490526001600160781b036201000083048116958401869052600160881b909204909116606090920182905293919291565b604080519415158552921515602085015291830152606082015260800161010f565b6101fb6101f63660046145dd565b610360565b60405161010f919061473f565b34801561021457600080fd5b50610138610381565b34801561022957600080fd5b50610138610238366004614752565b61038b565b34801561024957600080fd5b5061025d61025836600461478d565b61051c565b604051901515815260200161010f565b6101fb61027b3660046147ce565b61052f565b61025d61028e366004614839565b6105ad565b61025d6102a1366004614882565b61061e565b6102b96102b43660046148f9565b61063c565b60405161010f9291906149a1565b3480156102d357600080fd5b506102dc6106c5565b60405161010f939291906149f0565b61025d6102f9366004614a23565b6106dd565b6102b961030c366004614a5e565b6106e8565b34801561031d57600080fd5b5061025d61032c36600461478d565b610716565b606061033b610722565b905090565b6001600160a01b0381166000908152600160205260408120545b92915050565b6060610377866103708688614b2e565b858561073b565b9695505050505050565b600061033b610756565b60408051610160810190915260009061035a90806103ac6020860186614016565b6001600160a01b031681526020018460200160208101906103cd9190614016565b6001600160a01b031681526020016103e86040860186614c62565b808060200260200160405190810160405280939291908181526020016000905b828210156104345761042560a08302860136819003810190614caa565b81526020019060010190610408565b505050918352505060200161044c6060860186614cc6565b808060200260200160405190810160405280939291908181526020016000905b828210156104985761048960c08302860136819003810190614d0e565b8152602001906001019061046c565b50505091835250506020016104b360a0860160808701614d2a565b60038111156104c4576104c4614671565b81526020018460a0013581526020018460c0013581526020018460e0013581526020018461010001358152602001846101200135815260200184806060019061050d9190614cc6565b909152506101408401356107b3565b600061052883836108f9565b9392505050565b60606105a261053e8686610ab5565b604080516000808252602082019092529061059a565b6105876040805160a081019091526000808252602082019081526020016000815260200160008152602001606081525090565b8152602001906001900390816105545790505b50858561073b565b90505b949350505050565b60006105286105bb84610b70565b6040805160008082526020820190925290610617565b6106046040805160a081019091526000808252602082019081526020016000815260200160008152602001606081525090565b8152602001906001900390816105d15790505b5084610c1b565b60006105a261062c86614d45565b6106368587614b2e565b84610c1b565b6060806106b461064c8b8b610ab5565b60408051600080825260208201909252906106a8565b6106956040805160a081019091526000808252602082019081526020016000815260200160008152602001606081525090565b8152602001906001900390816106625790505b508a8a8a8a8a8a610d0c565b915091509850989650505050505050565b60606000806106d2610d4b565b925092509250909192565b600061035a82610daa565b6060806107048b6106f98b8d614b2e565b8a8a8a8a8a8a610d0c565b91509150995099975050505050505050565b60006105288383611044565b606060206000526707536561706f727460275260606000f35b606061074b858560018851611245565b6105a285848461155c565b6000610760611687565b503360008181526001602081815260409283902080549092019182905591518181529092917f7ab0fc7de8910a6100b24df423c3d0835534506dca9473d30c3e7df51241b2cf910160405180910390a290565b610140820151604080519084015180516000939284927f000000000000000000000000000000000000000000000000000000000000000092602090910190845b81811015610820578251601f1901805186825260c0822086529052602093840193909201916001016107f3565b506020810260405120945050505060007f00000000000000000000000000000000000000000000000000000000000000009150604051602060608901510160005b8681101561088e578151601f1901805186825260e082208552905260209283019290910190600101610861565b505060408051602087029020601f198a0180517f00000000000000000000000000000000000000000000000000000000000000008252928b01805197815260608c018051938152610140909c019a8b5261018082209390915295909552939097525050925250919050565b6000610903611687565b60008083815b81811015610aa8573687878381811061092457610924614d51565b90506020028101906109369190614d67565b9050366109438280614d87565b90506109526020820182614016565b945061096561096082614d9e565b6116ac565b60008181526002602090815260408083208151608081018352905460ff808216151583526101008204161515938201939093526001600160781b03620100008404811692820192909252600160881b9092041660608201529197506109cf908890839060016116e7565b508051610a9a57610a2286886109e86020870187614daa565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506117a092505050565b600087815260026020908152604091829020805460ff19166001179055610a4d918401908401614016565b6001600160a01b0316866001600160a01b03167ffde361574a066b44b3b5fe98a87108b7565e327327954c4faeea56a4e6491a0a89604051610a9191815260200190565b60405180910390a35b836001019350505050610909565b5060019695505050505050565b606081806001600160401b03811115610ad057610ad061404c565b604051908082528060200260200182016040528015610b0957816020015b610af6613e9d565b815260200190600190039081610aee5790505b50915060005b81811015610b6857610b43858583818110610b2c57610b2c614d51565b9050602002810190610b3e9190614d67565b610b70565b838281518110610b5557610b55614d51565b6020908102919091010152600101610b0f565b505092915050565b610b78613e9d565b6040805160a0810190915280610b8e8480614d87565b610b9790614d9e565b815260200160016001600160781b0316815260200160016001600160781b03168152602001838060200190610bcc9190614daa565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509385525050604080516020818101909252928352909201525092915050565b6000610c256117f5565b60606000806000610c398888600187611804565b604080516001808252818301909252939650919450925060009190816020015b610c61613e9d565b815260200190600190039081610c595790505090508881600081518110610c8a57610c8a614d51565b6020026020010181905250610c9f8189611a8d565b600081600081518110610cb457610cb4614d51565b6020026020010151600001519050610cd48185858461012001518c611e0a565b610cf28582600001518360200151338560400151866060015161200e565b610cfc6001600055565b5060019998505050505050505050565b606080610d1c8a8a600086611245565b610d3a8a610d2a898b614e3e565b610d34888a614e3e565b87612073565b909b909a5098505050505050505050565b6060600080610d5861224f565b6040805160018082528183019092529193507f000000000000000000000000000000000000000000000000000000000000000092506020820181803683375050603160f81b6020830152509391925090565b600060046101243590810490600316600182113415811480610de657604051630a61be9f60e41b81523460048201526024015b60405180910390fd5b5060008060006003861160a08102602401359350600287146002881160028903020192506001830181026002861502880103915050610e29888684878786612345565b6000610e3b60808a0160608b01614016565b90506101c46003881160200201356000866005811115610e5d57610e5d614671565b03610ea957610e8883610e7660c08d0160a08e01614016565b84338e60c001358f60e0013587612685565b610ea460408b013583610e9f6102008e018e614f10565b612739565b610cfc565b6040805160208082528183019092526000916020820181803683370190505090506002896005811115610ede57610ede614671565b03610f3e57610f09610ef660c08d0160a08e01614016565b84338e60c001358f60e0013587876127fe565b610f393384610f1b60208f018f614016565b8e604001358f806102000190610f319190614f10565b60008861284b565b61102a565b6003896005811115610f5257610f52614671565b03610f7d57610f09610f6a60c08d0160a08e01614016565b84338e60c001358f60e0013587876128d2565b6004896005811115610f9157610f91614671565b03610fef57610fb9610fa660208d018d614016565b33858e602001358f6040013587876127fe565b610f3983338d60a0016020810190610fd19190614016565b8e60e001358f806102000190610fe79190614f10565b60018861284b565b611012610fff60208d018d614016565b33858e602001358f6040013587876128d2565b61102a83338d60a0016020810190610fd19190614016565b61103381612908565b505060019998505050505050505050565b600061104e611687565b60008083815b81811015610aa8573687878381811061106f5761106f614d51565b90506020028101906110819190614d87565b90506110906020820182614016565b94506110a26040820160208301614016565b9350336001600160a01b038616148015906110c65750336001600160a01b03851614155b156110e45760405163203b1cdd60e21b815260040160405180910390fd5b60006111d3604051806101600160405280886001600160a01b03168152602001876001600160a01b031681526020018480604001906111239190614c62565b808060200260200160405190810160405280939291908181526020016000905b8282101561116f5761116060a08302860136819003810190614caa565b81526020019060010190611143565b50505091835250506020016111876060860186614cc6565b808060200260200160405190810160405280939291908181526020016000905b82821015610498576111c460c08302860136819003810190614d0e565b815260200190600101906111a7565b60008181526002602052604090819020805461ffff1916610100179055519091506001600160a01b0380871691908816907f6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d906112339085815260200190565b60405180910390a35050600101611054565b61124d6117f5565b83516000816001600160401b038111156112695761126961404c565b604051908082528060200260200182016040528015611292578160200160208202803683370190505b5090506000815260005b828110156114b15760008782815181106112b8576112b8614d51565b60200260200101519050846000036112dd5760006020909101526001810182526114a9565b60008060006112ee848b8b89611804565b925092509250600185018652816000036113155750506000602090920191909152506114a9565b8286868151811061132857611328614d51565b6020908102919091010152835160a081015160c0820151604090920151600019909a019990918290039042839003908183039060005b81518110156113f557600082828151811061137b5761137b614d51565b6020026020010151905060006113968a8a8460800151612931565b905081608001518260600151036113b357606082018190526113c8565b6113c28a8a8460600151612931565b60608301525b6080820181905260608201516113e3908288888b6000612981565b6060909201919091525060010161135e565b5088516060015160005b815181101561149d57600082828151811061141c5761141c614d51565b6020026020010151905060006114378b8b8460800151612931565b905081608001518260600151036114545760608201819052611469565b6114638b8b8460600151612931565b60608301525b608082018190526060820151611484908289898c6001612981565b60608301525060a08101516080909101526001016113ff565b50505050505050505050505b60010161129c565b506114bc8686611a8d565b8315330260005b83811015611552576000801b8382815181106114e1576114e1614d51565b6020026020010151031561154a57600088828151811061150357611503614d51565b602002602001015160000151905061154884838151811061152657611526614d51565b602002602001015182600001518360200151868560400151866060015161200e565b505b6001016114c3565b5050505050505050565b606081806001600160401b038111156115775761157761404c565b6040519080825280602002602001820160405280156115b057816020015b61159d613ed1565b8152602001906001900390816115955790505b5091506000805b8281101561166557368686838181106115d2576115d2614d51565b90506020028101906115e49190614d67565b90506000611608896115f68480614f10565b6116036020870187614f10565b6129dc565b905080602001516001600160a01b03168160000151608001516001600160a01b03160361163a5760018401935061165b565b80868585038151811061164f5761164f614d51565b60200260200101819052505b50506001016115b7565b508015611673578083510383525b5061167e8583612ccd565b50509392505050565b6001600054146116aa57604051637fa8a98760e01b815260040160405180910390fd5b565b60006116c2826060015151836101400151612ee2565b81516001600160a01b031660009081526001602052604090205461035a9083906107b3565b600083602001511561171d57811561171557604051630694555d60e21b815260048101869052602401610ddd565b5060006105a5565b60408401516001600160781b0316156117955782156117525760405163ee9e0e6360e01b815260048101869052602401610ddd565b83606001516001600160781b031684604001516001600160781b031610611795578115611715576040516310fda3e160e01b815260048101869052602401610ddd565b506001949350505050565b336001600160a01b038416036117b557505050565b60006117e26117c261224f565b61190160f01b600090815260029190915260228581526042822091905290565b90506117ef848284612f03565b50505050565b6117fd611687565b6002600055565b600080600080876000015190506118248160a001518260c0015188613046565b611838575060009250829150819050611a83565b602088015160408901516001600160781b0391821691168082118061185b575081155b1561187957604051632d02959960e11b815260040160405180910390fd5b808210801561188d57506080830151600116155b156118ab5760405163a11b63ff60e01b815260040160405180910390fd5b6118b4836116ac565b95506118d68a8a89898760e00151886080015189600001518a6020015161308c565b60008681526002602090815260408083208151608081018352905460ff808216151583526101008204161515938201939093526001600160781b03620100008404811692820192909252600160881b9092041660608201529061193d90889083908c6116e7565b611952575060009450849350611a8392505050565b805161196b5761196b8460000151888d606001516117a0565b604081015160608201516001600160781b0391821691168015611a2f578360010361199b578094508093506119c7565b8381146119c7576119ac8483614f6f565b91506119b88186614f6f565b94506119c48185614f6f565b93505b836119d28684614f8e565b11156119de5781840394505b600089815260026020526040902080546001600160781b03868116600160881b026001600160881b03868a019290921662010000026001600160881b03199093169290921760011716179055611a78565b600089815260026020526040902080546001600160781b03868116600160881b026001600160881b0391891662010000026001600160881b031990931692909217600117161790555b509295509093505050505b9450945094915050565b8051825160005b82811015611cd4576000848281518110611ab057611ab0614d51565b60200260200101519050600081600001519050838110611ae3576040516321a561b160e21b815260040160405180910390fd5b868181518110611af557611af5614d51565b6020026020010151602001516001600160781b0316600003611b18575050611ccc565b6000878281518110611b2c57611b2c614d51565b60209081029190910101515160408401519091506000808086602001516001811115611b5a57611b5a614671565b03611bf457604084015180518410611b8557604051635fd9fc6760e11b815260040160405180910390fd5b6000818581518110611b9957611b99614d51565b602090810291909101015180516040820151909550935090506004841460030381816005811115611bcc57611bcc614671565b90816005811115611bdf57611bdf614671565b90525050606088015160409091015250611c85565b606084015180518410611c1a576040516330446bef60e11b815260040160405180910390fd5b6000818581518110611c2e57611c2e614d51565b602090810291909101015180516040820151909550935090506004841460030381816005811115611c6157611c61614671565b90816005811115611c7457611c74614671565b905250506060880151604090910152505b611c8f8260031090565b611cac57604051634a75b57b60e11b815260040160405180910390fd5b8015611cc557611cc5866060015182886080015161316d565b5050505050505b600101611a94565b5060005b81811015611e03576000858281518110611cf457611cf4614d51565b6020026020010151905080602001516001600160781b0316600003611d195750611dfb565b6000868381518110611d2d57611d2d614d51565b60200260200101516000015190506000816060015151905060005b81811015611da457611d7b83606001518281518110611d6957611d69614d51565b60200260200101516000015160031090565b15611d9c5760405160016202297360e61b0319815260040160405180910390fd5b600101611d48565b505060408101515160005b81811015611df657611dd083604001518281518110611d6957611d69614d51565b15611dee5760405163a6cfc67360e01b815260040160405180910390fd5b600101611daf565b505050505b600101611cd8565b5050505050565b60008560a001518660c00151611e209190614fa6565b905060008660a0015142611e349190614fa6565b90506000611e428284614fa6565b60408051602080825281830190925291925034916000916020820181803683370190505090506131dd60005b8b6040015151811015611f265760008c604001518281518110611e9357611e93614d51565b602002602001015190506000611eb8826060015183608001518f8f8c8c8f600061329a565b606083018190523360808401529050600082516005811115611edc57611edc614671565b03611f085785811115611f0257604051631a783b8d60e01b815260040160405180910390fd5b80860395505b611f1c828f600001518d888863ffffffff16565b5050600101611e6e565b506131dd905060005b8b6060015151811015611fe75760008c606001518281518110611f5457611f54614d51565b602002602001015190506000611f79826060015183608001518f8f8c8c8f600161329a565b6060830181905260a083015160808401529050600082516005811115611fa157611fa1614671565b03611fcd5785811115611fc757604051631a783b8d60e01b815260040160405180910390fd5b80860395505b611fdd82338c888863ffffffff16565b5050600101611f2f565b5050611ff281612908565b81156120025761200233836132e6565b50505050505050505050565b60608290506060829050856001600160a01b0316876001600160a01b03167f9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f318a8886866040516120619493929190614ff7565b60405180910390a35050505050505050565b8251825160609182916120868183614f8e565b6001600160401b0381111561209d5761209d61404c565b6040519080825280602002602001820160405280156120d657816020015b6120c3613ed1565b8152602001906001900390816120bb5790505b5092506000805b8381101561216f5760008982815181106120f9576120f9614d51565b6020026020010151905060006121128c6000848c61333a565b905080602001516001600160a01b03168160000151608001516001600160a01b03160361214457600184019350612165565b80878585038151811061215957612159614d51565b60200260200101819052505b50506001016120dd565b5060005b8281101561220757600088828151811061218f5761218f614d51565b6020026020010151905060006121a88c6001848c61333a565b905080602001516001600160a01b03168160000151608001516001600160a01b0316036121da576001840193506121fd565b80878588860103815181106121f1576121f1614d51565b60200260200101819052505b5050600101612173565b508015612215578084510384525b5082516000036122385760405163d5da9a1b60e01b815260040160405180910390fd5b6122428884612ccd565b9350505094509492505050565b60007f000000000000000000000000000000000000000000000000000000000000000046146123205761033b604080517f000000000000000000000000000000000000000000000000000000000000000060208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b507f000000000000000000000000000000000000000000000000000000000000000090565b61234d6117f5565b6123638661012001358761014001356001613046565b5061236c6133d5565b61239461237d610200880188614f10565b61238991506001614f8e565b876101e00135612ee2565b6000807f00000000000000000000000000000000000000000000000000000000000000009050806080528560a0526060602460c037604060646101203760e06080206101605261026435602081026102a0016001610264350181526020810190508781526080602460208301376101608760a0528660c052600060e05261020435925060005b83811015612465578060400261028401602081610100376040816101203760208301925060e0608020835260a08401935089845288602085015260408160608601375060010161241a565b60206001850102610160206060526102643593505b838110156124ab578060400261028401915060a083019250888352876020840152604082606085013760010161247a565b505050505060007f00000000000000000000000000000000000000000000000000000000000000009050806080528260a052606060c460c03760206101046101203760c0608020600052602060002060e052602061026435026102000160018152836020820152606060c4604083013750506084356001600160a01b0381166000908152600160205260408120547f000000000000000000000000000000000000000000000000000000000000000060808190529091506040608460a03760605161010052886101205260a061014461014037816101e0526101806080209350505050602061026435026101800181815233602082015260806040820152610120606082015260a061026435026101e00160a4356084357f9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f318385a35050600060605261262081886101600135888a606001602081019061260b9190614016565b61261b60a08d0160808e01614016565b613416565b61267c8161263460808a0160608b01614016565b6126426102208b018b614daa565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061346692505050565b50505050505050565b80156126e15760006040519050632671a55160e11b815260206004820152600160248201528760448201528660648201528560848201528460a48201528360c48201528260e48201526126db8282610104613508565b5061267c565b60028760058111156126f5576126f5614671565b0361272c578160011461271b5760405163efcc00b160e01b815260040160405180910390fd5b612727868686866135f8565b61267c565b61267c86868686866136b9565b348160005b818110156127ac573685858381811061275957612759614d51565b6040029190910191505080358481111561278657604051631a783b8d60e01b815260040160405180910390fd5b61279f6127996040840160208501614016565b826132e6565b909303925060010161273e565b50818611156127ce57604051631a783b8d60e01b815260040160405180910390fd5b6127d885876132e6565b858211156127ec576127ec338784036132e6565b6127f66001600055565b505050505050565b612808818361379d565b8161283a578260011461282e5760405163efcc00b160e01b815260040160405180910390fd5b612727878787876135f8565b61267c828260028a8a8a8a8a6137bc565b602082026101e403358360005b818110156128b9573687878381811061287357612873614d51565b60400291909101915050803586156128925761288f818b614fa6565b99505b6128af8b8e6128a76040860160208701614016565b84898b61383c565b5050600101612858565b506128c8888b8b8a868861383c565b6120026001600055565b6128db83613877565b6128e5818361379d565b816128f75761272787878787876136b9565b61267c828260038a8a8a8a8a6137bc565b80516040146129145750565b6000612921826020015190565b905061292d8183613898565b5050565b6000828403612941575080610528565b600083858409159050806129685760405163c63cf08960e01b815260040160405180910390fd5b60006129748685614f6f565b9490940495945050505050565b60008587146129d15760008215612999575060001983015b6000816129a6888a614f6f565b6129b0888c614f6f565b6129ba9190614f8e565b6129c49190614f8e565b8590049250610377915050565b509395945050505050565b6129e4613ed1565b8315806129ef575081155b15612a0d57604051634c74edb760e11b815260040160405180910390fd5b612a15613ed1565b612a72878585808060200260200160405190810160405280939291908181526020016000905b82821015612a6757612a5860408302860136819003810190615091565b81526020019060010190612a3b565b5050505050836138bc565b805160408051602080890282018101909252878152612ad1918a91908a908a90819060009085015b82821015612ac657612ab760408302860136819003810190615091565b81526020019060010190612a9a565b505050505085613a68565b80516005811115612ae457612ae4614671565b8351516005811115612af857612af8614671565b141580612b23575080602001516001600160a01b03168360000151602001516001600160a01b031614155b80612b3a5750806040015183600001516040015114155b15612b58576040516309cfb45560e01b815260040160405180910390fd5b82600001516060015181606001511115612c0f57600085856000818110612b8157612b81614d51565b905060400201803603810190612b979190615091565b90508360000151606001518260600151612bb19190614fa6565b89826000015181518110612bc757612bc7614d51565b60200260200101516000015160600151826020015181518110612bec57612bec614d51565b602090810291909101015160609081019190915284518101519083015250612ca1565b600087876000818110612c2457612c24614d51565b905060400201803603810190612c3a9190615091565b90508160600151846000015160600151612c549190614fa6565b89826000015181518110612c6a57612c6a614d51565b60200260200101516000015160400151826020015181518110612c8f57612c8f614d51565b60200260200101516060018181525050505b60608082015184519091015260809081015183516001600160a01b039091169101525095945050505050565b8151606090806001600160401b03811115612cea57612cea61404c565b604051908082528060200260200182016040528015612d13578160200160208202803683370190505b50915060005b81811015612df9576000858281518110612d3557612d35614d51565b6020026020010151905080602001516001600160781b0316600003612d5a5750612df1565b6001848381518110612d6e57612d6e614d51565b9115156020928302919091019091015280516060015160005b8151811015612ded576000828281518110612da457612da4614d51565b602002602001015160600151905080600014612de4576040516314bea84160e31b8152600481018690526024810183905260448101829052606401610ddd565b50600101612d87565b5050505b600101612d19565b5060408051602080825281830190925234916000919060208201818036833701905050905060005b8551811015612eb5576000868281518110612e3e57612e3e614d51565b60209081029190910101518051909150600081516005811115612e6357612e63614671565b03612e97578481606001511115612e8d57604051631a783b8d60e01b815260040160405180910390fd5b8060600151850394505b612eab8183602001518460400151876131dd565b5050600101612e21565b50612ebf81612908565b8115612ecf57612ecf33836132e6565b612ed96001600055565b50505092915050565b8082101561292d57604051632335530b60e11b815260040160405180910390fd5b60008060008351604003612f3457505050602081015160408201516001600160ff1b0381169060ff1c601b01612f9a565b8351604103612f8f5750505060208101516040820151606083015160001a601b8114801590612f6757508060ff16601c14155b15612f8a57604051630f801e8560e11b815260ff82166004820152602401610ddd565b612f9a565b6127f6868686613bf2565b6040805160008082526020820180845288905260ff841692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015612fee573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661302257604051638baa579f60e01b815260040160405180910390fd5b866001600160a01b0316816001600160a01b03161461267c5761267c878787613bf2565b6000428411806130565750428311155b1561308257811561307a576040516337bf561360e11b815260040160405180910390fd5b506000610528565b5060019392505050565b60018360038111156130a0576130a0614671565b1180156130b65750336001600160a01b03821614155b80156130cb5750336001600160a01b03831614155b15611552576080880151511580156130e257508651155b156130f8576130f381868487613c69565b611552565b600061315682633313157060e01b88338d8c8e60405160240161311f959493929190615265565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613cb5565b90506131628187613cca565b505050505050505050565b60008360208301602084510281015b808210156131b9578151808411801561319c5781600052846020526131a5565b84600052816020525b50506040600020925060208201915061317c565b505083149050806117ef576040516309bde33960e01b815260040160405180910390fd5b6000845160058111156131f2576131f2614671565b0361320e57613209846080015185606001516132e6565b6117ef565b60018451600581111561322357613223614671565b036132425761320984602001518486608001518760600151868661383c565b60028451600581111561325757613257614671565b0361327b5761320984602001518486608001518760400151886060015187876127fe565b6117ef84602001518486608001518760400151886060015187876128d2565b60008789036132b5576132ae87878a612931565b90506132da565b6132d76132c388888c612931565b6132ce89898c612931565b87878787612981565b90505b98975050505050505050565b6132ef81613877565b600080600080600085875af19050806133355761330a613d24565b60405163470c7c1d60e01b81526001600160a01b038416600482015260248101839052604401610ddd565b505050565b613342613ed1565b8251600003613366578360405163375c24c160e01b8152600401610ddd91906153eb565b600084600181111561337a5761337a614671565b0361338f5761338a858483613a68565b6133a8565b61339a8584836138bc565b336020820152604081018290525b8051606001516000036105a557602081015181516001600160a01b03909116608090910152949350505050565b610244356102606102643560400201146004356020146102243561024014161680613413576040516339f3e3fd60e01b815260040160405180910390fd5b50565b600183600381111561342a5761342a614671565b1180156134405750336001600160a01b03821614155b80156134555750336001600160a01b03831614155b15611e0357611e0381868487613c69565b6000838152600260209081526040918290208251608081018452905460ff808216151583526101008204161515928201929092526001600160781b03620100008304811693820193909352600160881b90910490911660608201526134ce84826001806116e7565b5080516134e0576134e08385846117a0565b5050506000908152600260205260409020710100000000000000000000000000000100019055565b6040805160ff60a01b7f000000000000000000000000000000000000000000000000000000000000000017600090815260208681527f000000000000000000000000000000000000000000000000000000000000000084526055600b20929093528080526001600160a01b039091169181848682865af19050806135b25761358e613d24565b60405163344f54f560e21b81526001600160a01b0383166004820152602401610ddd565b6000516001600160e01b03198116632671a55160e11b146127f657604051630e7ccd9360e11b8152600481018790526001600160a01b0384166024820152604401610ddd565b833b61361357632f8aeb3960e11b6000528360045260246000fd5b6040516323b872dd60e01b6000528360045282602452816044526000806064600080895af1806136aa573d156136845760203d0460208304816003028183111561366b57818303600302610200838002858002030401015b5a602082011015613680573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005285600452846024528360445282606452600160845260a46000fd5b50604052505060006060525050565b843b6136d457632f8aeb3960e11b6000528460045260246000fd5b60405160805160a05160c051637921219560e11b6000528760045286602452856044528460645260a0608452600060a45260008060c46000808d5af180613781573d1561375c5760203d0460208604816003028183111561374357818303600302610200838002858002030401015b5a602082011015613758573d6000803e3d6000fd5b5050505b63f486bc8760e01b600052896004528860245287604452866064528560845260a46000fd5b5060809290925260a05260c05260405250506000606052505050565b60006137aa836020015190565b90508181146133355761333583612908565b600060208851036137f75750604080885260208089018a9052632671a55160e11b918901919091526044880152600160648801819052613806565b50606487018051600101908190525b603c60c082028901038781528660208201528560408201528460608201528360808201528260a082015250505050505050505050565b61384583613877565b61384f818361379d565b816138655761386086868686613d69565b6127f6565b6127f68282600189898960008a6137bc565b806000036134135760405163246cf94560e21b815260040160405180910390fd5b6064810151604082019060c0026044016138b3848383613508565b50506020905250565b6138e8565b637fda727960e01b60005260206000fd5b634e487b7160e01b600052601160045260246000fd5b60208201805151845181106138ff576138ff6138c1565b6020810260208601015160608151015160208451015181518110613925576139256138c1565b6020810260208301015160008060208601511561394c575050606081018051600090915280155b885183518152602084015160208201526040840151604082015260a084015160808201526060812060208c51028c015b808b1015613a235760208b019a508a515199508d518a1061399f5761399f6138c1565b60208a0260208f010151985060208901511561397c57606089510151975060208b5101519650875187106139d5576139d56138c1565b602087026020890101519550606086018051860181511587821060011b1786179550809650506000815250606086208214608084015160a08801511416613a1e57613a1e6138c1565b61397c565b50506060018290528060018114613a415760028114613a5257613a5a565b63246cf94560e21b60005260206000fd5b613a5a6138d2565b505050505050505050505050565b6020820180515184518110613a7f57613a7f6138c1565b602081026020860101518051604081015160208551015181518110613aa657613aa66138c1565b60208102602083010151600080602087015115613acd575050606081018051600090915280155b8951336080820152835181526020840151602082015260408401516040820152865160208c015261012087015160408c015260608120905060208c51028c015b808b1015613bc15760208b019a508a515199508d518a10613b3057613b306138c1565b60208a0260208f0101519850602089015115613b0d57885197506040880151965060208b510151955086518610613b6957613b696138c1565b602086026020880101519450606085018051850181511586821060011b178517945080955050600081525060608520821460408d01516101208a01511460208e01518a51141616613bbc57613bbc6138c1565b613b0d565b50508160608b5101528060018114613a415760028103613be357613be36138d2565b50505050505050505050505050565b6000613c1384631626ba7e60e01b858560405160240161311f9291906153f9565b905080613c3b57613c22613d24565b604051634f7fb80d60e01b815260040160405180910390fd5b613c4b630b135d3f60e11b613e6f565b156117ef57604051632057875960e21b815260040160405180910390fd5b604051602481018490523360448201526001600160a01b038316606482015260848101829052600090613ca99086906303874c7760e21b9060a40161311f565b9050611e038185613cca565b6000806000835160208501865afa9392505050565b81613cf357613cd7613d24565b604051633ed4053f60e21b815260048101829052602401610ddd565b613d036303874c7760e21b613e6f565b1561292d57604051633ed4053f60e21b815260048101829052602401610ddd565b3d156116aa5760203d046020604051048160030281831115613d5457818303600302610200838002858002030401015b5a602082011015613335573d6000803e3d6000fd5b6040516323b872dd60e01b600052836004528260245281604452602060006064600080895af1803d15601f3d116001600051141617163d15158116613e5f5780873b151516613e5f5780613e4a5781613e29573d15613e035760203d04602084048160030281831115613dea57818303600302610200838002858002030401015b5a602082011015613dff573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005286600452856024528460445260006064528360845260a46000fd5b639889192360e01b6000528660045285602452846044528360645260846000fd5b632f8aeb3960e11b6000528660045260246000fd5b5050604052505060006060525050565b60008060203d03613e855760206000803e506000515b6001600160e01b031990811692169190911415919050565b6040518060a00160405280613eb0613f14565b81526000602082018190526040820152606080820181905260809091015290565b60408051610100810182526000606082018181526080830182905260a0830182905260c0830182905260e083018290528252602082018190529181019190915290565b60405180610160016040528060006001600160a01b0316815260200160006001600160a01b03168152602001606081526020016060815260200160006003811115613f6157613f61614671565b815260006020820181905260408201819052606082018190526080820181905260a0820181905260c09091015290565b6000815180845260005b81811015613fb757602081850181015186830182015201613f9b565b81811115613fc9576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006105286020830184613f91565b6001600160a01b038116811461341357600080fd5b803561401181613ff1565b919050565b60006020828403121561402857600080fd5b813561052881613ff1565b60006020828403121561404557600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b03811182821017156140845761408461404c565b60405290565b60405161016081016001600160401b03811182821017156140845761408461404c565b604051601f8201601f191681016001600160401b03811182821017156140d5576140d561404c565b604052919050565b60006001600160401b038211156140f6576140f661404c565b5060051b60200190565b80356006811061401157600080fd5b600060a0828403121561412157600080fd5b614129614062565b905061413482614100565b8152602082013561414481613ff1565b8060208301525060408201356040820152606082013560608201526080820135608082015292915050565b600082601f83011261418057600080fd5b81356020614195614190836140dd565b6140ad565b82815260a092830285018201928282019190878511156141b457600080fd5b8387015b858110156141d7576141ca898261410f565b84529284019281016141b8565b5090979650505050505050565b600060c082840312156141f657600080fd5b60405160c081018181106001600160401b03821117156142185761421861404c565b60405290508061422783614100565b8152602083013561423781613ff1565b8060208301525060408301356040820152606083013560608201526080830135608082015260a083013561426a81613ff1565b60a0919091015292915050565b600082601f83011261428857600080fd5b81356020614298614190836140dd565b82815260c092830285018201928282019190878511156142b757600080fd5b8387015b858110156141d7576142cd89826141e4565b84529284019281016142bb565b80356004811061401157600080fd5b600061016082840312156142fc57600080fd5b61430461408a565b905061430f82614006565b815261431d60208301614006565b602082015260408201356001600160401b038082111561433c57600080fd5b6143488583860161416f565b6040840152606084013591508082111561436157600080fd5b5061436e84828501614277565b606083015250614380608083016142da565b608082015260a082013560a082015260c082013560c082015260e082013560e082015261010080830135818301525061012080830135818301525061014080830135818301525092915050565b80356001600160781b038116811461401157600080fd5b600082601f8301126143f557600080fd5b81356001600160401b0381111561440e5761440e61404c565b614421601f8201601f19166020016140ad565b81815284602083860101111561443657600080fd5b816020850160208301376000918101602001919091529392505050565b600060a0828403121561446557600080fd5b61446d614062565b905081356001600160401b038082111561448657600080fd5b614492858386016142e9565b83526144a0602085016143cd565b60208401526144b1604085016143cd565b604084015260608401359150808211156144ca57600080fd5b6144d6858386016143e4565b606084015260808401359150808211156144ef57600080fd5b506144fc848285016143e4565b60808301525092915050565b600082601f83011261451957600080fd5b81356020614529614190836140dd565b82815260059290921b8401810191818101908684111561454857600080fd5b8286015b848110156145875780356001600160401b0381111561456b5760008081fd5b6145798986838b0101614453565b84525091830191830161454c565b509695505050505050565b60008083601f8401126145a457600080fd5b5081356001600160401b038111156145bb57600080fd5b6020830191508360208260051b85010111156145d657600080fd5b9250929050565b6000806000806000606086880312156145f557600080fd5b85356001600160401b038082111561460c57600080fd5b61461889838a01614508565b9650602088013591508082111561462e57600080fd5b61463a89838a01614592565b9096509450604088013591508082111561465357600080fd5b5061466088828901614592565b969995985093965092949392505050565b634e487b7160e01b600052602160045260246000fd5b6006811061469757614697614671565b9052565b6146a6828251614687565b6020818101516001600160a01b0390811691840191909152604080830151908401526060808301519084015260809182015116910152565b600081518084526020808501945080840160005b8381101561473457815161470788825161469b565b808401516001600160a01b031660a08901526040015160c088015260e090960195908201906001016146f2565b509495945050505050565b60208152600061052860208301846146de565b60006020828403121561476457600080fd5b81356001600160401b0381111561477a57600080fd5b8201610160818503121561052857600080fd5b600080602083850312156147a057600080fd5b82356001600160401b038111156147b657600080fd5b6147c285828601614592565b90969095509350505050565b600080600080604085870312156147e457600080fd5b84356001600160401b03808211156147fb57600080fd5b61480788838901614592565b9096509450602087013591508082111561482057600080fd5b5061482d87828801614592565b95989497509550505050565b6000806040838503121561484c57600080fd5b82356001600160401b0381111561486257600080fd5b83016040818603121561487457600080fd5b946020939093013593505050565b6000806000806060858703121561489857600080fd5b84356001600160401b03808211156148af57600080fd5b9086019060a082890312156148c357600080fd5b909450602086013590808211156148d957600080fd5b506148e687828801614592565b9598909750949560400135949350505050565b60008060008060008060008060a0898b03121561491557600080fd5b88356001600160401b038082111561492c57600080fd5b6149388c838d01614592565b909a50985060208b013591508082111561495157600080fd5b61495d8c838d01614592565b909850965060408b013591508082111561497657600080fd5b506149838b828c01614592565b999c989b509699959896976060870135966080013595509350505050565b604080825283519082018190526000906020906060840190828701845b828110156149dc5781511515845292840192908401906001016149be565b5050508381038285015261037781866146de565b606081526000614a036060830186613f91565b6020830194909452506001600160a01b0391909116604090910152919050565b600060208284031215614a3557600080fd5b81356001600160401b03811115614a4b57600080fd5b8201610240818503121561052857600080fd5b600080600080600080600080600060c08a8c031215614a7c57600080fd5b89356001600160401b0380821115614a9357600080fd5b614a9f8d838e01614508565b9a5060208c0135915080821115614ab557600080fd5b614ac18d838e01614592565b909a50985060408c0135915080821115614ada57600080fd5b614ae68d838e01614592565b909850965060608c0135915080821115614aff57600080fd5b50614b0c8c828d01614592565b9a9d999c50979a96999598959660808101359660a09091013595509350505050565b6000614b3c614190846140dd565b83815260208082019190600586811b860136811115614b5a57600080fd5b865b81811015614c555780356001600160401b0380821115614b7c5760008081fd5b818a01915060a08236031215614b925760008081fd5b614b9a614062565b823581528683013560028110614bb05760008081fd5b81880152604083810135908201526060808401359082015260808084013583811115614bdc5760008081fd5b939093019236601f850112614bf357600092508283fd5b83359250614c03614190846140dd565b83815292871b84018801928881019036851115614c205760008081fd5b948901945b84861015614c3e57853582529489019490890190614c25565b918301919091525088525050948301948301614b5c565b5092979650505050505050565b6000808335601e19843603018112614c7957600080fd5b8301803591506001600160401b03821115614c9357600080fd5b602001915060a0810236038213156145d657600080fd5b600060a08284031215614cbc57600080fd5b610528838361410f565b6000808335601e19843603018112614cdd57600080fd5b8301803591506001600160401b03821115614cf757600080fd5b602001915060c0810236038213156145d657600080fd5b600060c08284031215614d2057600080fd5b61052883836141e4565b600060208284031215614d3c57600080fd5b610528826142da565b600061035a3683614453565b634e487b7160e01b600052603260045260246000fd5b60008235603e19833603018112614d7d57600080fd5b9190910192915050565b6000823561015e19833603018112614d7d57600080fd5b600061035a36836142e9565b6000808335601e19843603018112614dc157600080fd5b8301803591506001600160401b03821115614ddb57600080fd5b6020019150368190038213156145d657600080fd5b600060408284031215614e0257600080fd5b604051604081018181106001600160401b0382111715614e2457614e2461404c565b604052823581526020928301359281019290925250919050565b6000614e4c614190846140dd565b80848252602080830192508560051b850136811115614e6a57600080fd5b855b81811015614f045780356001600160401b03811115614e8b5760008081fd5b870136601f820112614e9d5760008081fd5b8035614eab614190826140dd565b81815260069190911b82018501908581019036831115614ecb5760008081fd5b928601925b82841015614ef457614ee23685614df0565b82528682019150604084019350614ed0565b8852505050938201938201614e6c565b50919695505050505050565b6000808335601e19843603018112614f2757600080fd5b8301803591506001600160401b03821115614f4157600080fd5b6020019150600681901b36038213156145d657600080fd5b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615614f8957614f89614f59565b500290565b60008219821115614fa157614fa1614f59565b500190565b600082821015614fb857614fb8614f59565b500390565b600081518084526020808501945080840160005b8381101561473457614fe487835161469b565b60a0969096019590820190600101614fd1565b60006080808301878452602060018060a01b03808916828701526040848188015283895180865260a089019150848b01955060005b8181101561506d578651615041848251614687565b80870151861684880152848101518585015260609081015190840152958501959187019160010161502c565b50508781036060890152615081818a614fbd565b9c9b505050505050505050505050565b6000604082840312156150a357600080fd5b6105288383614df0565b600081518084526020808501945080840160005b838110156147345781516150d6888251614687565b838101516001600160a01b03168885015260408082015190890152606080820151908901526080908101519088015260a090960195908201906001016150c1565b600081518084526020808501945080840160005b83811015614734578151615140888251614687565b808401516001600160a01b0390811689860152604080830151908a0152606080830151908a0152608080830151908a015260a091820151169088015260c0909601959082019060010161512b565b6004811061469757614697614671565b600081518084526020808501945080840160005b83811015614734578151875295820195908201906001016151b2565b6002811061469757614697614671565b600082825180855260208086019550808260051b84010181860160005b848110156141d757601f19868403018952815160a08151855285820151615224878701826151ce565b5060408281015190860152606080830151908601526080918201519185018190526152518186018361519e565b9a86019a94505050908301906001016151fb565b85815260018060a01b038516602082015260a060408201526000610140855160a08085015261529f82850182516001600160a01b03169052565b60208101516101606152bb818701836001600160a01b03169052565b6040830151915080610180870152506152d86102a08601826150ad565b9050606082015161013f19868303016101a08701526152f78282615117565b915050608082015161530d6101c087018261518e565b5060a08201516101e086015260c082015161020086015260e082015161022086015261010080830151610240870152610120808401516102608801528484015161028088015260208a0151945061536f60c08801866001600160781b03169052565b60408a01516001600160781b031660e088015260608a0151878403609f19908101848a015290955093506153a38386613f91565b945060808a0151925083878603018188015250506153c18382613f91565b9250505082810360608401526153d7818661519e565b905082810360808401526132da81856151de565b6020810161035a82846151ce565b8281526040602082015260006105a56040830184613f9156fea26469706673582212208e752577dcf751ff7148a242455324782600ba827332881a3fb1f1b9cb84c45f64736f6c634300080d0033",
}

// SeaportABI is the input ABI used to generate the binding from.
// Deprecated: Use SeaportMetaData.ABI instead.
var SeaportABI = SeaportMetaData.ABI

// SeaportBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SeaportMetaData.Bin instead.
var SeaportBin = SeaportMetaData.Bin

// DeploySeaport deploys a new Ethereum contract, binding an instance of Seaport to it.
func DeploySeaport(auth *bind.TransactOpts, backend bind.ContractBackend, conduitController common.Address) (common.Address, *types.Transaction, *Seaport, error) {
	parsed, err := SeaportMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SeaportBin), backend, conduitController)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Seaport{SeaportCaller: SeaportCaller{contract: contract}, SeaportTransactor: SeaportTransactor{contract: contract}, SeaportFilterer: SeaportFilterer{contract: contract}}, nil
}

// Seaport is an auto generated Go binding around an Ethereum contract.
type Seaport struct {
	SeaportCaller     // Read-only binding to the contract
	SeaportTransactor // Write-only binding to the contract
	SeaportFilterer   // Log filterer for contract events
}

// SeaportCaller is an auto generated read-only Go binding around an Ethereum contract.
type SeaportCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeaportTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SeaportTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeaportFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SeaportFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeaportSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SeaportSession struct {
	Contract     *Seaport          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SeaportCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SeaportCallerSession struct {
	Contract *SeaportCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SeaportTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SeaportTransactorSession struct {
	Contract     *SeaportTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SeaportRaw is an auto generated low-level Go binding around an Ethereum contract.
type SeaportRaw struct {
	Contract *Seaport // Generic contract binding to access the raw methods on
}

// SeaportCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SeaportCallerRaw struct {
	Contract *SeaportCaller // Generic read-only contract binding to access the raw methods on
}

// SeaportTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SeaportTransactorRaw struct {
	Contract *SeaportTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSeaport creates a new instance of Seaport, bound to a specific deployed contract.
func NewSeaport(address common.Address, backend bind.ContractBackend) (*Seaport, error) {
	contract, err := bindSeaport(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Seaport{SeaportCaller: SeaportCaller{contract: contract}, SeaportTransactor: SeaportTransactor{contract: contract}, SeaportFilterer: SeaportFilterer{contract: contract}}, nil
}

// NewSeaportCaller creates a new read-only instance of Seaport, bound to a specific deployed contract.
func NewSeaportCaller(address common.Address, caller bind.ContractCaller) (*SeaportCaller, error) {
	contract, err := bindSeaport(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SeaportCaller{contract: contract}, nil
}

// NewSeaportTransactor creates a new write-only instance of Seaport, bound to a specific deployed contract.
func NewSeaportTransactor(address common.Address, transactor bind.ContractTransactor) (*SeaportTransactor, error) {
	contract, err := bindSeaport(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SeaportTransactor{contract: contract}, nil
}

// NewSeaportFilterer creates a new log filterer instance of Seaport, bound to a specific deployed contract.
func NewSeaportFilterer(address common.Address, filterer bind.ContractFilterer) (*SeaportFilterer, error) {
	contract, err := bindSeaport(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SeaportFilterer{contract: contract}, nil
}

// bindSeaport binds a generic wrapper to an already deployed contract.
func bindSeaport(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SeaportABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Seaport *SeaportRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Seaport.Contract.SeaportCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Seaport *SeaportRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seaport.Contract.SeaportTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Seaport *SeaportRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Seaport.Contract.SeaportTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Seaport *SeaportCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Seaport.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Seaport *SeaportTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seaport.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Seaport *SeaportTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Seaport.Contract.contract.Transact(opts, method, params...)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address offerer) view returns(uint256 nonce)
func (_Seaport *SeaportCaller) GetNonce(opts *bind.CallOpts, offerer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Seaport.contract.Call(opts, &out, "getNonce", offerer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address offerer) view returns(uint256 nonce)
func (_Seaport *SeaportSession) GetNonce(offerer common.Address) (*big.Int, error) {
	return _Seaport.Contract.GetNonce(&_Seaport.CallOpts, offerer)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address offerer) view returns(uint256 nonce)
func (_Seaport *SeaportCallerSession) GetNonce(offerer common.Address) (*big.Int, error) {
	return _Seaport.Contract.GetNonce(&_Seaport.CallOpts, offerer)
}

// GetOrderHash is a free data retrieval call binding the contract method 0x79df72bd.
//
// Solidity: function getOrderHash((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) order) view returns(bytes32 orderHash)
func (_Seaport *SeaportCaller) GetOrderHash(opts *bind.CallOpts, order OrderComponents) ([32]byte, error) {
	var out []interface{}
	err := _Seaport.contract.Call(opts, &out, "getOrderHash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOrderHash is a free data retrieval call binding the contract method 0x79df72bd.
//
// Solidity: function getOrderHash((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) order) view returns(bytes32 orderHash)
func (_Seaport *SeaportSession) GetOrderHash(order OrderComponents) ([32]byte, error) {
	return _Seaport.Contract.GetOrderHash(&_Seaport.CallOpts, order)
}

// GetOrderHash is a free data retrieval call binding the contract method 0x79df72bd.
//
// Solidity: function getOrderHash((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) order) view returns(bytes32 orderHash)
func (_Seaport *SeaportCallerSession) GetOrderHash(order OrderComponents) ([32]byte, error) {
	return _Seaport.Contract.GetOrderHash(&_Seaport.CallOpts, order)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns(bool isValidated, bool isCancelled, uint256 totalFilled, uint256 totalSize)
func (_Seaport *SeaportCaller) GetOrderStatus(opts *bind.CallOpts, orderHash [32]byte) (struct {
	IsValidated bool
	IsCancelled bool
	TotalFilled *big.Int
	TotalSize   *big.Int
}, error) {
	var out []interface{}
	err := _Seaport.contract.Call(opts, &out, "getOrderStatus", orderHash)

	outstruct := new(struct {
		IsValidated bool
		IsCancelled bool
		TotalFilled *big.Int
		TotalSize   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsValidated = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsCancelled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.TotalFilled = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalSize = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns(bool isValidated, bool isCancelled, uint256 totalFilled, uint256 totalSize)
func (_Seaport *SeaportSession) GetOrderStatus(orderHash [32]byte) (struct {
	IsValidated bool
	IsCancelled bool
	TotalFilled *big.Int
	TotalSize   *big.Int
}, error) {
	return _Seaport.Contract.GetOrderStatus(&_Seaport.CallOpts, orderHash)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns(bool isValidated, bool isCancelled, uint256 totalFilled, uint256 totalSize)
func (_Seaport *SeaportCallerSession) GetOrderStatus(orderHash [32]byte) (struct {
	IsValidated bool
	IsCancelled bool
	TotalFilled *big.Int
	TotalSize   *big.Int
}, error) {
	return _Seaport.Contract.GetOrderStatus(&_Seaport.CallOpts, orderHash)
}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() view returns(string version, bytes32 domainSeparator, address conduitController)
func (_Seaport *SeaportCaller) Information(opts *bind.CallOpts) (struct {
	Version           string
	DomainSeparator   [32]byte
	ConduitController common.Address
}, error) {
	var out []interface{}
	err := _Seaport.contract.Call(opts, &out, "information")

	outstruct := new(struct {
		Version           string
		DomainSeparator   [32]byte
		ConduitController common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Version = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.DomainSeparator = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ConduitController = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() view returns(string version, bytes32 domainSeparator, address conduitController)
func (_Seaport *SeaportSession) Information() (struct {
	Version           string
	DomainSeparator   [32]byte
	ConduitController common.Address
}, error) {
	return _Seaport.Contract.Information(&_Seaport.CallOpts)
}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() view returns(string version, bytes32 domainSeparator, address conduitController)
func (_Seaport *SeaportCallerSession) Information() (struct {
	Version           string
	DomainSeparator   [32]byte
	ConduitController common.Address
}, error) {
	return _Seaport.Contract.Information(&_Seaport.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string contractName)
func (_Seaport *SeaportCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Seaport.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string contractName)
func (_Seaport *SeaportSession) Name() (string, error) {
	return _Seaport.Contract.Name(&_Seaport.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string contractName)
func (_Seaport *SeaportCallerSession) Name() (string, error) {
	return _Seaport.Contract.Name(&_Seaport.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0xfd9f1e10.
//
// Solidity: function cancel((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256)[] orders) returns(bool cancelled)
func (_Seaport *SeaportTransactor) Cancel(opts *bind.TransactOpts, orders []OrderComponents) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "cancel", orders)
}

// Cancel is a paid mutator transaction binding the contract method 0xfd9f1e10.
//
// Solidity: function cancel((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256)[] orders) returns(bool cancelled)
func (_Seaport *SeaportSession) Cancel(orders []OrderComponents) (*types.Transaction, error) {
	return _Seaport.Contract.Cancel(&_Seaport.TransactOpts, orders)
}

// Cancel is a paid mutator transaction binding the contract method 0xfd9f1e10.
//
// Solidity: function cancel((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256)[] orders) returns(bool cancelled)
func (_Seaport *SeaportTransactorSession) Cancel(orders []OrderComponents) (*types.Transaction, error) {
	return _Seaport.Contract.Cancel(&_Seaport.TransactOpts, orders)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xdf7b0dac.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) advancedOrder, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Seaport *SeaportTransactor) FulfillAdvancedOrder(opts *bind.TransactOpts, advancedOrder AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "fulfillAdvancedOrder", advancedOrder, criteriaResolvers, fulfillerConduitKey)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xdf7b0dac.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) advancedOrder, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Seaport *SeaportSession) FulfillAdvancedOrder(advancedOrder AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillAdvancedOrder(&_Seaport.TransactOpts, advancedOrder, criteriaResolvers, fulfillerConduitKey)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xdf7b0dac.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) advancedOrder, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Seaport *SeaportTransactorSession) FulfillAdvancedOrder(advancedOrder AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillAdvancedOrder(&_Seaport.TransactOpts, advancedOrder, criteriaResolvers, fulfillerConduitKey)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0xfb4c2af9.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Seaport *SeaportTransactor) FulfillAvailableAdvancedOrders(opts *bind.TransactOpts, advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "fulfillAvailableAdvancedOrders", advancedOrders, criteriaResolvers, offerFulfillments, considerationFulfillments, fulfillerConduitKey, maximumFulfilled)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0xfb4c2af9.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Seaport *SeaportSession) FulfillAvailableAdvancedOrders(advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillAvailableAdvancedOrders(&_Seaport.TransactOpts, advancedOrders, criteriaResolvers, offerFulfillments, considerationFulfillments, fulfillerConduitKey, maximumFulfilled)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0xfb4c2af9.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Seaport *SeaportTransactorSession) FulfillAvailableAdvancedOrders(advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillAvailableAdvancedOrders(&_Seaport.TransactOpts, advancedOrders, criteriaResolvers, offerFulfillments, considerationFulfillments, fulfillerConduitKey, maximumFulfilled)
}

// FulfillAvailableOrders is a paid mutator transaction binding the contract method 0xed98a574.
//
// Solidity: function fulfillAvailableOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Seaport *SeaportTransactor) FulfillAvailableOrders(opts *bind.TransactOpts, orders []Order, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "fulfillAvailableOrders", orders, offerFulfillments, considerationFulfillments, fulfillerConduitKey, maximumFulfilled)
}

// FulfillAvailableOrders is a paid mutator transaction binding the contract method 0xed98a574.
//
// Solidity: function fulfillAvailableOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Seaport *SeaportSession) FulfillAvailableOrders(orders []Order, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillAvailableOrders(&_Seaport.TransactOpts, orders, offerFulfillments, considerationFulfillments, fulfillerConduitKey, maximumFulfilled)
}

// FulfillAvailableOrders is a paid mutator transaction binding the contract method 0xed98a574.
//
// Solidity: function fulfillAvailableOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Seaport *SeaportTransactorSession) FulfillAvailableOrders(orders []Order, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillAvailableOrders(&_Seaport.TransactOpts, orders, offerFulfillments, considerationFulfillments, fulfillerConduitKey, maximumFulfilled)
}

// FulfillBasicOrder is a paid mutator transaction binding the contract method 0xfb0f3ee1.
//
// Solidity: function fulfillBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Seaport *SeaportTransactor) FulfillBasicOrder(opts *bind.TransactOpts, parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "fulfillBasicOrder", parameters)
}

// FulfillBasicOrder is a paid mutator transaction binding the contract method 0xfb0f3ee1.
//
// Solidity: function fulfillBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Seaport *SeaportSession) FulfillBasicOrder(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillBasicOrder(&_Seaport.TransactOpts, parameters)
}

// FulfillBasicOrder is a paid mutator transaction binding the contract method 0xfb0f3ee1.
//
// Solidity: function fulfillBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Seaport *SeaportTransactorSession) FulfillBasicOrder(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillBasicOrder(&_Seaport.TransactOpts, parameters)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xb3a34c4c.
//
// Solidity: function fulfillOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) order, bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Seaport *SeaportTransactor) FulfillOrder(opts *bind.TransactOpts, order Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "fulfillOrder", order, fulfillerConduitKey)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xb3a34c4c.
//
// Solidity: function fulfillOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) order, bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Seaport *SeaportSession) FulfillOrder(order Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillOrder(&_Seaport.TransactOpts, order, fulfillerConduitKey)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xb3a34c4c.
//
// Solidity: function fulfillOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) order, bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Seaport *SeaportTransactorSession) FulfillOrder(order Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillOrder(&_Seaport.TransactOpts, order, fulfillerConduitKey)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns(uint256 newNonce)
func (_Seaport *SeaportTransactor) IncrementNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "incrementNonce")
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns(uint256 newNonce)
func (_Seaport *SeaportSession) IncrementNonce() (*types.Transaction, error) {
	return _Seaport.Contract.IncrementNonce(&_Seaport.TransactOpts)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns(uint256 newNonce)
func (_Seaport *SeaportTransactorSession) IncrementNonce() (*types.Transaction, error) {
	return _Seaport.Contract.IncrementNonce(&_Seaport.TransactOpts)
}

// MatchAdvancedOrders is a paid mutator transaction binding the contract method 0x55944a42.
//
// Solidity: function matchAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, ((uint256,uint256)[],(uint256,uint256)[])[] fulfillments) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Seaport *SeaportTransactor) MatchAdvancedOrders(opts *bind.TransactOpts, advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillments []Fulfillment) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "matchAdvancedOrders", advancedOrders, criteriaResolvers, fulfillments)
}

// MatchAdvancedOrders is a paid mutator transaction binding the contract method 0x55944a42.
//
// Solidity: function matchAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, ((uint256,uint256)[],(uint256,uint256)[])[] fulfillments) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Seaport *SeaportSession) MatchAdvancedOrders(advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillments []Fulfillment) (*types.Transaction, error) {
	return _Seaport.Contract.MatchAdvancedOrders(&_Seaport.TransactOpts, advancedOrders, criteriaResolvers, fulfillments)
}

// MatchAdvancedOrders is a paid mutator transaction binding the contract method 0x55944a42.
//
// Solidity: function matchAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, ((uint256,uint256)[],(uint256,uint256)[])[] fulfillments) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Seaport *SeaportTransactorSession) MatchAdvancedOrders(advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillments []Fulfillment) (*types.Transaction, error) {
	return _Seaport.Contract.MatchAdvancedOrders(&_Seaport.TransactOpts, advancedOrders, criteriaResolvers, fulfillments)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa8174404.
//
// Solidity: function matchOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders, ((uint256,uint256)[],(uint256,uint256)[])[] fulfillments) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Seaport *SeaportTransactor) MatchOrders(opts *bind.TransactOpts, orders []Order, fulfillments []Fulfillment) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "matchOrders", orders, fulfillments)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa8174404.
//
// Solidity: function matchOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders, ((uint256,uint256)[],(uint256,uint256)[])[] fulfillments) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Seaport *SeaportSession) MatchOrders(orders []Order, fulfillments []Fulfillment) (*types.Transaction, error) {
	return _Seaport.Contract.MatchOrders(&_Seaport.TransactOpts, orders, fulfillments)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa8174404.
//
// Solidity: function matchOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders, ((uint256,uint256)[],(uint256,uint256)[])[] fulfillments) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Seaport *SeaportTransactorSession) MatchOrders(orders []Order, fulfillments []Fulfillment) (*types.Transaction, error) {
	return _Seaport.Contract.MatchOrders(&_Seaport.TransactOpts, orders, fulfillments)
}

// Validate is a paid mutator transaction binding the contract method 0x88147732.
//
// Solidity: function validate(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders) returns(bool validated)
func (_Seaport *SeaportTransactor) Validate(opts *bind.TransactOpts, orders []Order) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "validate", orders)
}

// Validate is a paid mutator transaction binding the contract method 0x88147732.
//
// Solidity: function validate(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders) returns(bool validated)
func (_Seaport *SeaportSession) Validate(orders []Order) (*types.Transaction, error) {
	return _Seaport.Contract.Validate(&_Seaport.TransactOpts, orders)
}

// Validate is a paid mutator transaction binding the contract method 0x88147732.
//
// Solidity: function validate(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders) returns(bool validated)
func (_Seaport *SeaportTransactorSession) Validate(orders []Order) (*types.Transaction, error) {
	return _Seaport.Contract.Validate(&_Seaport.TransactOpts, orders)
}

// SeaportNonceIncrementedIterator is returned from FilterNonceIncremented and is used to iterate over the raw logs and unpacked data for NonceIncremented events raised by the Seaport contract.
type SeaportNonceIncrementedIterator struct {
	Event *SeaportNonceIncremented // Event containing the contract specifics and raw log

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
func (it *SeaportNonceIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeaportNonceIncremented)
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
		it.Event = new(SeaportNonceIncremented)
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
func (it *SeaportNonceIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeaportNonceIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeaportNonceIncremented represents a NonceIncremented event raised by the Seaport contract.
type SeaportNonceIncremented struct {
	NewNonce *big.Int
	Offerer  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceIncremented is a free log retrieval operation binding the contract event 0x7ab0fc7de8910a6100b24df423c3d0835534506dca9473d30c3e7df51241b2cf.
//
// Solidity: event NonceIncremented(uint256 newNonce, address indexed offerer)
func (_Seaport *SeaportFilterer) FilterNonceIncremented(opts *bind.FilterOpts, offerer []common.Address) (*SeaportNonceIncrementedIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Seaport.contract.FilterLogs(opts, "NonceIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return &SeaportNonceIncrementedIterator{contract: _Seaport.contract, event: "NonceIncremented", logs: logs, sub: sub}, nil
}

// WatchNonceIncremented is a free log subscription operation binding the contract event 0x7ab0fc7de8910a6100b24df423c3d0835534506dca9473d30c3e7df51241b2cf.
//
// Solidity: event NonceIncremented(uint256 newNonce, address indexed offerer)
func (_Seaport *SeaportFilterer) WatchNonceIncremented(opts *bind.WatchOpts, sink chan<- *SeaportNonceIncremented, offerer []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Seaport.contract.WatchLogs(opts, "NonceIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeaportNonceIncremented)
				if err := _Seaport.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
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

// ParseNonceIncremented is a log parse operation binding the contract event 0x7ab0fc7de8910a6100b24df423c3d0835534506dca9473d30c3e7df51241b2cf.
//
// Solidity: event NonceIncremented(uint256 newNonce, address indexed offerer)
func (_Seaport *SeaportFilterer) ParseNonceIncremented(log types.Log) (*SeaportNonceIncremented, error) {
	event := new(SeaportNonceIncremented)
	if err := _Seaport.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeaportOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Seaport contract.
type SeaportOrderCancelledIterator struct {
	Event *SeaportOrderCancelled // Event containing the contract specifics and raw log

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
func (it *SeaportOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeaportOrderCancelled)
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
		it.Event = new(SeaportOrderCancelled)
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
func (it *SeaportOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeaportOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeaportOrderCancelled represents a OrderCancelled event raised by the Seaport contract.
type SeaportOrderCancelled struct {
	OrderHash [32]byte
	Offerer   common.Address
	Zone      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Seaport *SeaportFilterer) FilterOrderCancelled(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*SeaportOrderCancelledIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Seaport.contract.FilterLogs(opts, "OrderCancelled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &SeaportOrderCancelledIterator{contract: _Seaport.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Seaport *SeaportFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *SeaportOrderCancelled, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Seaport.contract.WatchLogs(opts, "OrderCancelled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeaportOrderCancelled)
				if err := _Seaport.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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
func (_Seaport *SeaportFilterer) ParseOrderCancelled(log types.Log) (*SeaportOrderCancelled, error) {
	event := new(SeaportOrderCancelled)
	if err := _Seaport.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeaportOrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the Seaport contract.
type SeaportOrderFulfilledIterator struct {
	Event *SeaportOrderFulfilled // Event containing the contract specifics and raw log

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
func (it *SeaportOrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeaportOrderFulfilled)
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
		it.Event = new(SeaportOrderFulfilled)
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
func (it *SeaportOrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeaportOrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeaportOrderFulfilled represents a OrderFulfilled event raised by the Seaport contract.
type SeaportOrderFulfilled struct {
	OrderHash     [32]byte
	Offerer       common.Address
	Zone          common.Address
	Fulfiller     common.Address
	Offer         []SpentItem
	Consideration []ReceivedItem
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address fulfiller, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
func (_Seaport *SeaportFilterer) FilterOrderFulfilled(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*SeaportOrderFulfilledIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Seaport.contract.FilterLogs(opts, "OrderFulfilled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &SeaportOrderFulfilledIterator{contract: _Seaport.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address fulfiller, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
func (_Seaport *SeaportFilterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *SeaportOrderFulfilled, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Seaport.contract.WatchLogs(opts, "OrderFulfilled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeaportOrderFulfilled)
				if err := _Seaport.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
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
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address fulfiller, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
func (_Seaport *SeaportFilterer) ParseOrderFulfilled(log types.Log) (*SeaportOrderFulfilled, error) {
	event := new(SeaportOrderFulfilled)
	if err := _Seaport.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeaportOrderValidatedIterator is returned from FilterOrderValidated and is used to iterate over the raw logs and unpacked data for OrderValidated events raised by the Seaport contract.
type SeaportOrderValidatedIterator struct {
	Event *SeaportOrderValidated // Event containing the contract specifics and raw log

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
func (it *SeaportOrderValidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeaportOrderValidated)
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
		it.Event = new(SeaportOrderValidated)
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
func (it *SeaportOrderValidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeaportOrderValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeaportOrderValidated represents a OrderValidated event raised by the Seaport contract.
type SeaportOrderValidated struct {
	OrderHash [32]byte
	Offerer   common.Address
	Zone      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderValidated is a free log retrieval operation binding the contract event 0xfde361574a066b44b3b5fe98a87108b7565e327327954c4faeea56a4e6491a0a.
//
// Solidity: event OrderValidated(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Seaport *SeaportFilterer) FilterOrderValidated(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*SeaportOrderValidatedIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Seaport.contract.FilterLogs(opts, "OrderValidated", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &SeaportOrderValidatedIterator{contract: _Seaport.contract, event: "OrderValidated", logs: logs, sub: sub}, nil
}

// WatchOrderValidated is a free log subscription operation binding the contract event 0xfde361574a066b44b3b5fe98a87108b7565e327327954c4faeea56a4e6491a0a.
//
// Solidity: event OrderValidated(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Seaport *SeaportFilterer) WatchOrderValidated(opts *bind.WatchOpts, sink chan<- *SeaportOrderValidated, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Seaport.contract.WatchLogs(opts, "OrderValidated", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeaportOrderValidated)
				if err := _Seaport.contract.UnpackLog(event, "OrderValidated", log); err != nil {
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
func (_Seaport *SeaportFilterer) ParseOrderValidated(log types.Log) (*SeaportOrderValidated, error) {
	event := new(SeaportOrderValidated)
	if err := _Seaport.contract.UnpackLog(event, "OrderValidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
