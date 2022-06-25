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
	Counter       *big.Int
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduitController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadContractSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadFraction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BadReturnValueFromERC20OnTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"BadSignatureV\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConsiderationCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortfallAmount\",\"type\":\"uint256\"}],\"name\":\"ConsiderationNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CriteriaNotEnabledForItem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"identifiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"ERC1155BatchTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EtherTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InexactFraction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientEtherSupplied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Invalid1155BatchTransferEncoding\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBasicOrderParameterEncoding\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidCallToConduit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCanceller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidConduit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidERC721TransferAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFulfillmentComponentData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNativeOfferItem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidRestrictedOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchedFulfillmentOfferAndConsiderationComponents\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"}],\"name\":\"MissingFulfillmentComponentOnAggregation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingItemAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingOriginalConsiderationItems\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NoContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoReentrantCalls\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSpecifiedOrdersAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferAndConsiderationRequiredOnFulfillment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderAlreadyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderIsCancelled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderPartiallyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PartialFillsNotEnabledForOrder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnresolvedConsiderationCriteria\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnresolvedOfferCriteria\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnusedItemParameters\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCounter\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"}],\"name\":\"CounterIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSpentItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structReceivedItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"}],\"name\":\"OrderValidated\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder\",\"name\":\"advancedOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"criteriaResolvers\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"fulfillAdvancedOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder[]\",\"name\":\"advancedOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"criteriaResolvers\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"offerFulfillments\",\"type\":\"tuple[][]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"considerationFulfillments\",\"type\":\"tuple[][]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maximumFulfilled\",\"type\":\"uint256\"}],\"name\":\"fulfillAvailableAdvancedOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"availableOrders\",\"type\":\"bool[]\"},{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"offerFulfillments\",\"type\":\"tuple[][]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"considerationFulfillments\",\"type\":\"tuple[][]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maximumFulfilled\",\"type\":\"uint256\"}],\"name\":\"fulfillAvailableOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"availableOrders\",\"type\":\"bool[]\"},{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"considerationToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"considerationIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"offerToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offerIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumBasicOrderType\",\"name\":\"basicOrderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offererConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalAdditionalRecipients\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structAdditionalRecipient[]\",\"name\":\"additionalRecipients\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBasicOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"}],\"name\":\"fulfillBasicOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"}],\"name\":\"fulfillOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"}],\"name\":\"getCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"getOrderStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidated\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSize\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newCounter\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"information\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"conduitController\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder[]\",\"name\":\"advancedOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"criteriaResolvers\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"offerComponents\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"considerationComponents\",\"type\":\"tuple[]\"}],\"internalType\":\"structFulfillment[]\",\"name\":\"fulfillments\",\"type\":\"tuple[]\"}],\"name\":\"matchAdvancedOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"offerComponents\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"considerationComponents\",\"type\":\"tuple[]\"}],\"internalType\":\"structFulfillment[]\",\"name\":\"fulfillments\",\"type\":\"tuple[]\"}],\"name\":\"matchOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"contractName\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"validate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"validated\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101c06040523480156200001257600080fd5b5060405162005b9b38038062005b9b8339810160408190526200003591620005c6565b808080808080808080806200004962000132565b610120526101005260e05260c081815260a0838152608085815246610140819052604080516020818101979097528082019890985260608801969096529086015230858201528351808603909101815293019091528151910120610160526001600160a01b03811661018081905260408051630a96ad3960e01b81528151630a96ad39926004808401939192918290030181865afa158015620000f0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001169190620005f8565b506101a052505060016000555062000688975050505050505050565b600080808080806200015e60408051808201909152600781526614d9585c1bdc9d60ca1b602082015290565b80516020918201206040805180820182526003815262312e3160e81b90840152519097507f722c0e0c80487266e8c6a45e3a1a803aab23378a9c32e6ebe029d4fad7bfc96596506000916200026391016909ecccccae492e8cada560b31b81526e1d5a5b9d0e081a5d195b551e5c194b608a1b600a8201526d1859191c995cdcc81d1bdad95b8b60921b60198201527f75696e74323536206964656e7469666965724f7243726974657269612c00000060278201527f75696e74323536207374617274416d6f756e742c0000000000000000000000006044820152701d5a5b9d0c8d4d88195b99105b5bdd5b9d607a1b6058820152602960f81b6069820152606a0190565b60408051601f1981840301815282825271086dedce6d2c8cae4c2e8d2dedc92e8cada560731b60208401526e1d5a5b9d0e081a5d195b551e5c194b608a1b60328401526d1859191c995cdcc81d1bdad95b8b60921b60418401527f75696e74323536206964656e7469666965724f7243726974657269612c000000604f8401527f75696e74323536207374617274416d6f756e742c000000000000000000000000606c840152711d5a5b9d0c8d4d88195b99105b5bdd5b9d0b60721b6080840152701859191c995cdcc81c9958da5c1a595b9d607a1b6092840152602960f81b60a384018190528251808503608401815260a485019093526f09ee4c8cae486dedae0dedccadce8e6560831b60c48501526f1859191c995cdcc81bd999995c995c8b60821b60d48501526c1859191c995cdcc81e9bdb994b609a1b60e48501527113d999995c925d195b56d7481bd999995c8b60721b60f18501527f436f6e73696465726174696f6e4974656d5b5d20636f6e73696465726174696f610103850152611b8b60f21b6101238501526f1d5a5b9d0e081bdc99195c951e5c194b60821b610125850152711d5a5b9d0c8d4d881cdd185c9d151a5b594b60721b6101358501526f1d5a5b9d0c8d4d88195b99151a5b594b60821b61014785015270189e5d195ccccc881e9bdb9952185cda0b607a1b6101578501526c1d5a5b9d0c8d4d881cd85b1d0b609a1b6101688501527f6279746573333220636f6e647569744b65792c000000000000000000000000006101758501526e3ab4b73a191a9b1031b7bab73a32b960891b6101888501526101978401529250906000906101980160408051601f19818403018152908290526c08a92a06e626488dedac2d2dc5609b1b60208301526b1cdd1c9a5b99c81b985b594b60a21b602d8301526e1cdd1c9a5b99c81d995c9cda5bdb8b608a1b60398301526f1d5a5b9d0c8d4d8818da185a5b92590b60821b60488301527f6164647265737320766572696679696e67436f6e7472616374000000000000006058830152602960f81b6071830152915060720160408051601f19818403018152908290528051602091820120855186830120855186840120919a5098509650620005a391839185918791016200065b565b604051602081830303815290604052805190602001209350505050909192939495565b600060208284031215620005d957600080fd5b81516001600160a01b0381168114620005f157600080fd5b9392505050565b600080604083850312156200060c57600080fd5b505080516020909101519092909150565b6000815160005b8181101562000640576020818501810151868301520162000624565b8181111562000650576000828601525b509290920192915050565b60006200067f620006786200067184886200061d565b866200061d565b846200061d565b95945050505050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a05161547b6200072060003960006135a7015260008181610d21015261357601526000612313015260006122430152600081816108b0015261252101526000818161083f015261237f0152600081816107d8015261249a01526000612271015260006122bf01526000612297015261547b6000f3fe6080604052600436106100e85760003560e01c8063a81744041161008a578063f07ec37311610059578063f07ec37314610290578063f47b7740146102b0578063fb0f3ee1146102d4578063fd9f1e10146102e757600080fd5b8063a817440414610244578063b3a34c4c14610257578063e7acab241461026a578063ed98a5741461027d57600080fd5b80635b34b966116100c65780635b34b966146101b057806379df72bd146101d357806387201b41146101f3578063881477321461021457600080fd5b806306fdde03146100ed57806346423aa71461011857806355944a4214610190575b600080fd5b3480156100f957600080fd5b50610102610307565b60405161010f9190613fe2565b60405180910390f35b34801561012457600080fd5b5061016e610133366004613ff5565b60009081526002602052604090205460ff808216926101008304909116916001600160781b03620100008204811692600160881b9092041690565b604080519415158552921515602085015291830152606082015260800161010f565b6101a361019e3660046145c4565b610316565b60405161010f9190614726565b3480156101bc57600080fd5b506101c5610339565b60405190815260200161010f565b3480156101df57600080fd5b506101c56101ee366004614739565b610343565b610206610201366004614774565b6104da565b60405161010f929190614854565b34801561022057600080fd5b5061023461022f3660046148ad565b610520565b604051901515815260200161010f565b6101a36102523660046148ee565b610533565b610234610265366004614959565b6105b1565b6102346102783660046149a2565b610623565b61020661028b366004614a31565b610657565b34801561029c57600080fd5b506101c56102ab366004614ad9565b6106e1565b3480156102bc57600080fd5b506102c56106ff565b60405161010f93929190614af6565b6102346102e2366004614b29565b610717565b3480156102f357600080fd5b506102346103023660046148ad565b610722565b606061031161072e565b905090565b606061032d866103268688614b64565b8585610746565b90505b95945050505050565b6000610311610763565b6040805161016081019091526000906104d490806103646020860186614ad9565b6001600160a01b031681526020018460200160208101906103859190614ad9565b6001600160a01b031681526020016103a06040860186614c98565b808060200260200160405190810160405280939291908181526020016000905b828210156103ec576103dd60a08302860136819003810190614ce0565b815260200190600101906103c0565b50505091835250506020016104046060860186614cfc565b808060200260200160405190810160405280939291908181526020016000905b828210156104505761044160c08302860136819003810190614d44565b81526020019060010190610424565b505050918352505060200161046b60a0860160808701614d60565b600381111561047c5761047c614658565b81526020018460a0013581526020018460c0013581526020018460e001358152602001846101000135815260200184610120013581526020018480606001906104c59190614cfc565b909152506101408401356107c0565b92915050565b60608061050d8c6104eb8c8e614b64565b8b8b8b8b8b6001600160a01b038c1615610505578b610507565b335b8b610906565b915091509a509a98505050505050505050565b600061052c8383610948565b9392505050565b60606105a66105428686610aad565b604080516000808252602082019092529061059e565b61058b6040805160a081019091526000808252602082019081526020016000815260200160008152602001606081525090565b8152602001906001900390816105585790505b508585610746565b90505b949350505050565b600061052c6105bf84610b68565b604080516000808252602082019092529061061b565b6106086040805160a081019091526000808252602082019081526020016000815260200160008152602001606081525090565b8152602001906001900390816105d55790505b508433610c13565b600061032d61063187614d7b565b61063b8688614b64565b856001600160a01b038616156106515785610c13565b33610c13565b6060806106d06106678b8b610aad565b60408051600080825260208201909252906106c3565b6106b06040805160a081019091526000808252602082019081526020016000815260200160008152602001606081525090565b81526020019060019003908161067d5790505b508a8a8a8a8a338b610906565b915091509850989650505050505050565b6001600160a01b0381166000908152600160205260408120546104d4565b606060008061070c610d00565b925092509250909192565b60006104d482610d61565b600061052c8383611021565b6060602080526707536561706f727460475260606020f35b60606107588585600188516000611227565b6105a6858484611566565b600061076d611691565b503360008181526001602081815260409283902080549092019182905591518181529092917f721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f910160405180910390a290565b610140820151604080519084015180516000939284927f000000000000000000000000000000000000000000000000000000000000000092602090910190845b8181101561082d578251601f1901805186825260c082208652905260209384019390920191600101610800565b506020810260405120945050505060007f00000000000000000000000000000000000000000000000000000000000000009150604051602060608901510160005b8681101561089b578151601f1901805186825260e08220855290526020928301929091019060010161086e565b505060408051602087029020601f198a0180517f00000000000000000000000000000000000000000000000000000000000000008252928b01805197815260608c018051938152610140909c019a8b5261018082209390915295909552939097525050925250919050565b6060806109178b8b60008688611227565b6109368b6109258a8c614dd5565b61092f898b614dd5565b88886116b6565b909c909b509950505050505050505050565b6000610952611691565b6000808084815b81811015610a9f573688888381811061097457610974614ea7565b90506020028101906109869190614ebd565b9050366109938280614edd565b90506109a26020820182614ad9565b94506109b56109b082614ef4565b611896565b600081815260026020526040812098509096506109d7908790899060016118d1565b50865460ff16610a9557610a2d85876109f36020860186614f00565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061199692505050565b865460ff19166001178755610a486040820160208301614ad9565b6001600160a01b0316856001600160a01b03167ffde361574a066b44b3b5fe98a87108b7565e327327954c4faeea56a4e6491a0a88604051610a8c91815260200190565b60405180910390a35b5050600101610959565b506001979650505050505050565b606081806001600160401b03811115610ac857610ac861400e565b604051908082528060200260200182016040528015610b0157816020015b610aee613ea1565b815260200190600190039081610ae65790505b50915060005b81811015610b6057610b3b858583818110610b2457610b24614ea7565b9050602002810190610b369190614ebd565b610b68565b838281518110610b4d57610b4d614ea7565b6020908102919091010152600101610b07565b505092915050565b610b70613ea1565b6040805160a0810190915280610b868480614edd565b610b8f90614ef4565b815260200160016001600160781b0316815260200160016001600160781b03168152602001838060200190610bc49190614f00565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509385525050604080516020818101909252928352909201525092915050565b6000610c1d6119eb565b60606000806000610c3189896001876119fa565b604080516001808252818301909252939650919450925060009190816020015b610c59613ea1565b815260200190600190039081610c515790505090508981600081518110610c8257610c82614ea7565b6020026020010181905250610c97818a611cb3565b600081600081518110610cac57610cac614ea7565b6020026020010151600001519050610cc78185858c8c61200c565b610ce585826000015183602001518b856040015186606001516121da565b610cef6001600055565b5060019a9950505050505050505050565b6060600080610d0d61223f565b6040805160038082528183019092529193507f00000000000000000000000000000000000000000000000000000000000000009250602082018180368337505062312e3160e81b6020830152509391925090565b600061012435600281901c90600316600182113415811480610d9d57604051630a61be9f60e41b81523460048201526024015b60405180910390fd5b5060008060008060038711915060a082026024013593506002871460028811600289030201905060018101820260028615028801039250610de2898783888888612335565b506101c46020820201356000856005811115610e0057610e00614658565b03610ea35760208901803590610e16908b614ad9565b6001600160a01b03161715610e3e57604051636ab37ce760e01b815260040160405180910390fd5b610e7383610e5260c08c0160a08d01614ad9565b610e6260808d0160608e01614ad9565b338d60c001358e60e001358761266b565b610e9e60408a0135610e8b60808c0160608d01614ad9565b610e996102008d018d614f46565b61271f565b611008565b6040805160208082528183019092526000916020820181803683370190505090506002886005811115610ed857610ed8614658565b03610f1757610f12610ef060c08c0160a08d01614ad9565b610f0060808d0160608e01614ad9565b338d60c001358e60e0013587876127da565b610fe2565b6003886005811115610f2b57610f2b614658565b03610f6557610f12610f4360c08c0160a08d01614ad9565b610f5360808d0160608e01614ad9565b338d60c001358e60e001358787612827565b6004886005811115610f7957610f79614658565b03610fb057610f12610f8e60208c018c614ad9565b33610f9f60808e0160608f01614ad9565b8d602001358e6040013587876127da565b610fe2610fc060208c018c614ad9565b33610fd160808e0160608f01614ad9565b8d602001358e604001358787612827565b610ffd610ff560808c0160608d01614ad9565b8b858461285d565b6110068161297e565b505b6110126001600055565b50600198975050505050505050565b600061102b611691565b6000808084815b81811015610a9f573688888381811061104d5761104d614ea7565b905060200281019061105f9190614edd565b905061106e6020820182614ad9565b94506110806040820160208301614ad9565b9350336001600160a01b038616148015906110a45750336001600160a01b03851614155b156110c25760405163203b1cdd60e21b815260040160405180910390fd5b60006111b1604051806101600160405280886001600160a01b03168152602001876001600160a01b031681526020018480604001906111019190614c98565b808060200260200160405190810160405280939291908181526020016000905b8282101561114d5761113e60a08302860136819003810190614ce0565b81526020019060010190611121565b50505091835250506020016111656060860186614cfc565b808060200260200160405190810160405280939291908181526020016000905b82821015610450576111a260c08302860136819003810190614d44565b81526020019060010190611185565b60008181526002602052604090819020805461ffff191661010017815590519098509091506001600160a01b0380871691908816907f6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d906112159085815260200190565b60405180910390a35050600101611032565b61122f6119eb565b84516000816001600160401b0381111561124b5761124b61400e565b604051908082528060200260200182016040528015611274578160200160208202803683370190505b506000808252909150601d6045823560e01c061160011b905b8381101561149d5760008982815181106112a9576112a9614ea7565b60200260200101519050866000036112ce576000602090910152600181018352611495565b60008060006112df848d8d8a6119fa565b92509250925060018501875281600003611306575050600060209092019190915250611495565b8287868151811061131957611319614ea7565b6020908102919091010152835160a081015160c08201516040909201518051600019909d019c91929160005b818110156113e157600083828151811061136157611361614ea7565b602002602001015190508051158c179b506000611383898984608001516129a7565b905081608001518260600151036113a057606082018190526113b5565b6113af898984606001516129a7565b60608301525b6080820181905260608201516113cf9082898960006129e9565b60609092019190915250600101611345565b50875160600151805160005b8181101561148957600083828151811061140957611409614ea7565b6020026020010151905060006114248b8b84608001516129a7565b905081608001518260600151036114415760608201819052611456565b6114508b8b84606001516129a7565b60608301525b60808201819052606082015161147090828b8b60016129e9565b60608301525060a08101516080909101526001016113ed565b50505050505050505050505b60010161128d565b50806003036114bf576040516312d3f5a360e01b815260040160405180910390fd5b6114c98888611cb3565b60005b8381101561155b576000801b8382815181106114ea576114ea614ea7565b6020026020010151031561155357600089828151811061150c5761150c614ea7565b602002602001015160000151905061155184838151811061152f5761152f614ea7565b60200260200101518260000151836020015189856040015186606001516121da565b505b6001016114cc565b505050505050505050565b606081806001600160401b038111156115815761158161400e565b6040519080825280602002602001820160405280156115ba57816020015b6115a7613ed5565b81526020019060019003908161159f5790505b5091506000805b8281101561166f57368686838181106115dc576115dc614ea7565b90506020028101906115ee9190614ebd565b90506000611612896116008480614f46565b61160d6020870187614f46565b612a3e565b905080602001516001600160a01b03168160000151608001516001600160a01b03160361164457836001019350611665565b80868585038151811061165957611659614ea7565b60200260200101819052505b50506001016115c1565b50801561167d578083510383525b506116888583612d23565b50509392505050565b6001600054146116b457604051637fa8a98760e01b815260040160405180910390fd5b565b8351835160609182916116c98183614fa5565b6001600160401b038111156116e0576116e061400e565b60405190808252806020026020018201604052801561171957816020015b611706613ed5565b8152602001906001900390816116fe5790505b5092506000805b838110156117b35760008a828151811061173c5761173c614ea7565b6020026020010151905060006117568d6000848d8d612f3d565b905080602001516001600160a01b03168160000151608001516001600160a01b031603611788578360010193506117a9565b80878585038151811061179d5761179d614ea7565b60200260200101819052505b5050600101611720565b5060005b8281101561184d5760008982815181106117d3576117d3614ea7565b6020026020010151905060006117ee8d6001848d6000612f3d565b905080602001516001600160a01b03168160000151608001516001600160a01b03160361182057836001019350611843565b808785888601038151811061183757611837614ea7565b60200260200101819052505b50506001016117b7565b50801561185b578084510384525b50825160000361187e5760405163d5da9a1b60e01b815260040160405180910390fd5b6118888984612d23565b935050509550959350505050565b60006118ac826060015151836101400151612fe2565b81516001600160a01b03166000908152600160205260409020546104d49083906107c0565b8254600090610100900460ff161561190d57811561190557604051630694555d60e21b815260048101869052602401610d94565b5060006105a9565b83546201000090046001600160781b0316801561198a5783156119465760405163ee9e0e6360e01b815260048101879052602401610d94565b8454600160881b90046001600160781b0316811061198a578215611980576040516310fda3e160e01b815260048101879052602401610d94565b60009150506105a9565b50600195945050505050565b336001600160a01b038416036119ab57505050565b60006119d86119b861223f565b61190160f01b600090815260029190915260228581526042822091905290565b90506119e5848284613003565b50505050565b6119f3611691565b6002600055565b60008060008087600001519050611a1a8160a001518260c0015188613154565b611a2e575060009250829150819050611ca9565b602088015160408901516001600160781b03918216911680821180611a51575081155b15611a6f57604051632d02959960e11b815260040160405180910390fd5b8082108015611a8357506080830151600116155b15611aa15760405163a11b63ff60e01b815260040160405180910390fd5b611aaa83611896565b9550611acc8a8a89898760e00151886080015189600001518a6020015161319a565b600086815260026020526040812090611ae990889083908c6118d1565b611afe575060009450849350611ca992505050565b805460ff16611b1a57611b1a8460000151888d60600151611996565b80546001600160781b03620100008204811691600160881b9004168015611c635783600103611b4e57809450809350611b7a565b838114611b7a57611b5f8483614fbd565b9150611b6b8186614fbd565b9450611b778185614fbd565b93505b83611b858684614fa5565b1115611b915781840394505b611b9b8583614fa5565b91506001600160781b0384116001600160781b0383111715611c2457611bd6565b60005b8215611bd057908290069190611bbf565b50919050565b611be9611be38584611bbc565b86611bbc565b80150194859004949384900493909104906001600160781b038083119085111715611c2457634e487b7160e01b600052601160045260246000fd5b82546001600160781b03858116600160881b026001600160881b0391851662010000026001600160881b03199093169290921760011716178355611c9e565b82546001600160781b03858116600160881b026001600160881b0391881662010000026001600160881b031990931692909217600117161783555b509295509093505050505b9450945094915050565b8051825160005b82811015611efa576000848281518110611cd657611cd6614ea7565b60200260200101519050600081600001519050838110611d09576040516321a561b160e21b815260040160405180910390fd5b868181518110611d1b57611d1b614ea7565b6020026020010151602001516001600160781b0316600003611d3e575050611ef2565b6000878281518110611d5257611d52614ea7565b60209081029190910101515160408401519091506000808086602001516001811115611d8057611d80614658565b03611e1a57604084015180518410611dab57604051635fd9fc6760e11b815260040160405180910390fd5b6000818581518110611dbf57611dbf614ea7565b602090810291909101015180516040820151909550935090506004841460030381816005811115611df257611df2614658565b90816005811115611e0557611e05614658565b90525050606088015160409091015250611eab565b606084015180518410611e40576040516330446bef60e11b815260040160405180910390fd5b6000818581518110611e5457611e54614ea7565b602090810291909101015180516040820151909550935090506004841460030381816005811115611e8757611e87614658565b90816005811115611e9a57611e9a614658565b905250506060880151604090910152505b611eb58260031090565b611ed257604051634a75b57b60e11b815260040160405180910390fd5b8015611eeb57611eeb866060015182886080015161327a565b5050505050505b600101611cba565b5060005b81811015612005576000858281518110611f1a57611f1a614ea7565b6020026020010151905080602001516001600160781b0316600003611f3f5750611ffd565b805160608101515160005b81811015611fa657611f7d83606001518281518110611f6b57611f6b614ea7565b60200260200101516000015160031090565b15611f9e5760405160016202297360e61b0319815260040160405180910390fd5b600101611f4a565b505060408101515160005b81811015611ff857611fd283604001518281518110611f6b57611f6b614ea7565b15611ff05760405163a6cfc67360e01b815260040160405180910390fd5b600101611fb1565b505050505b600101611efe565b5050505050565b60a085015160c08601516040805160208082528183019092526000916020820181803683375050506040890151519091506132dd9060005b818110156120fa5760008b60400151828151811061206457612064614ea7565b602002602001015190506000600581111561208157612081614658565b8151600581111561209457612094614658565b036120b2576040516312d3f5a360e01b815260040160405180910390fd5b60006120cc826060015183608001518e8e8c8c60006133ef565b606083015250608081018890528b516101208d01516120f19183918863ffffffff8916565b50600101612044565b50505060608801515134906132dd9060005b818110156121be5760008c60600151828151811061212c5761212c614ea7565b602002602001015190506000612150826060015183608001518f8f8d8d60016133ef565b6060830181905260a08301516080840152905060008251600581111561217857612178614658565b036121a4578581111561219e57604051631a783b8d60e01b815260040160405180910390fd5b80860395505b6121b482338d8a8963ffffffff16565b505060010161210c565b5050506121ca8261297e565b801561155b5761155b3382613439565b60608290506060829050856001600160a01b0316876001600160a01b03167f9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f318a88868660405161222d9493929190615016565b60405180910390a35050505050505050565b60007f0000000000000000000000000000000000000000000000000000000000000000461461231057610311604080517f000000000000000000000000000000000000000000000000000000000000000060208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b507f000000000000000000000000000000000000000000000000000000000000000090565b61233d6119eb565b6123538661012001358761014001356001613154565b5061235c61348d565b61237a61236d610200880188614f46565b9050876101e00135612fe2565b6000807f00000000000000000000000000000000000000000000000000000000000000009050806080528560a0526060602460c037604060646101203760e06080206101605261026435602081026102a0016001610264350181526020810190508781526080602460208301376101608760a0528660c052600060e05261020435925060005b8381101561244b578060400261028401602081610100376040816101203760208301925060e0608020835260a084019350898452886020850152604081606086013750600101612400565b60206001850102610160206060526102643593505b83811015612491578060400261028401915060a0830192508883528760208401526040826060850137600101612460565b505050505060007f00000000000000000000000000000000000000000000000000000000000000009050806080528260a052606060c460c03760206101046101203760c0608020600052602060002060e052602061026435026102000160018152836020820152606060c4604083013750506084356001600160a01b0381166000908152600160205260408120547f000000000000000000000000000000000000000000000000000000000000000060808190529091506040608460a03760605161010052886101205260a061014461014037816101e0526101806080209350505050602061026435026101800181815233602082015260806040820152610120606082015260a061026435026101e00160a4356084357f9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f318385a35050600060605261260681886101600135888a60600160208101906125f19190614ad9565b61260160a08d0160808e01614ad9565b6134d6565b6126628161261a60808a0160608b01614ad9565b6126286102208b018b614f00565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061352692505050565b50505050505050565b80156126c75760006040519050632671a55160e11b815260206004820152600160248201528760448201528660648201528560848201528460a48201528360c48201528260e48201526126c18282610104613570565b50612662565b60028760058111156126db576126db614658565b0361271257816001146127015760405163efcc00b160e01b815260040160405180910390fd5b61270d86868686613669565b612662565b612662868686868661372d565b348160005b81811015612792573685858381811061273f5761273f614ea7565b6040029190910191505080358481111561276c57604051631a783b8d60e01b815260040160405180910390fd5b61278561277f6040840160208501614ad9565b82613439565b9093039250600101612724565b50818611156127b457604051631a783b8d60e01b815260040160405180910390fd5b6127be8587613439565b858211156127d2576127d233878403613439565b505050505050565b6127e48183613814565b81612816578260011461280a5760405163efcc00b160e01b815260040160405180910390fd5b61270d87878787613669565b612662828260028a8a8a8a8a613833565b612830836138b3565b61283a8183613814565b8161284c5761270d878787878761372d565b612662828260038a8a8a8a8a613833565b600080600080600086156128945788945033935061288160c0890160a08a01614ad9565b9250505060e086013560c08701356128b6565b3394508893506128a76020890189614ad9565b92505050604086013560208701355b80156128d557604051636ab37ce760e01b815260040160405180910390fd5b50602086026101e4033560006128ef6102008a018a614f46565b9050905060005b81811015612963573661290d6102008c018c614f46565b8381811061291d5761291d614ea7565b6040029190910191505080358a1561293c5761293981876150b0565b95505b612959878a6129516040860160208701614ad9565b84898f6138d4565b50506001016128f6565b5061297284878786868c6138d4565b50505050505050505050565b604081511461298a5750565b6000612997826020015190565b90506129a3818361390f565b5050565b60008284036129b757508061052c565b82848309156129d15763c63cf08960e01b60005260046000fd5b60006129dd8584614fbd565b93909304949350505050565b6000848614612a3457838303428590038082036000612a08838a614fbd565b612a12838c614fbd565b612a1c9190614fa5565b90508584878303040181151502945050505050610330565b5092949350505050565b612a46613ed5565b831580612a51575081155b15612a6f57604051634c74edb760e11b815260040160405180910390fd5b612a77613ed5565b612ad4878585808060200260200160405190810160405280939291908181526020016000905b82821015612ac957612aba604083028601368190038101906150c7565b81526020019060010190612a9d565b505050505083613933565b805160408051602080890282018101909252878152612b33918a91908a908a90819060009085015b82821015612b2857612b19604083028601368190038101906150c7565b81526020019060010190612afc565b505050505085613ad6565b80516005811115612b4657612b46614658565b8351516005811115612b5a57612b5a614658565b141580612b85575080602001516001600160a01b03168360000151602001516001600160a01b031614155b80612b9c5750806040015183600001516040015114155b15612bba576040516309cfb45560e01b815260040160405180910390fd5b82600001516060015181606001511115612c6857600085856000818110612be357612be3614ea7565b905060400201803603810190612bf991906150c7565b905083600001516060015182606001510389826000015181518110612c2057612c20614ea7565b60200260200101516000015160600151826020015181518110612c4557612c45614ea7565b602090810291909101015160609081019190915284518101519083015250612d03565b600087876000818110612c7d57612c7d614ea7565b905060400201803603810190612c9391906150c7565b905081606001518460000151606001510389826000015181518110612cba57612cba614ea7565b60200260200101516000015160400151826020015181518110612cdf57612cdf614ea7565b60200260200101516060018181525050816060015184600001516060018181525050505b60809081015183516001600160a01b039091169101525095945050505050565b8151606090806001600160401b03811115612d4057612d4061400e565b604051908082528060200260200182016040528015612d69578160200160208202803683370190505b50915060005b81811015612e51576000858281518110612d8b57612d8b614ea7565b6020026020010151905080602001516001600160781b0316600003612db05750612e49565b6001848381518110612dc457612dc4614ea7565b91151560209283029190910190910152805160600151805160005b81811015612e44576000838281518110612dfb57612dfb614ea7565b602002602001015160600151905080600014612e3b576040516314bea84160e31b8152600481018790526024810183905260448101829052606401610d94565b50600101612ddf565b505050505b600101612d6f565b5060408051602080825281830190925234916000919060208201818036833701905050855190915060005b81811015612f0f576000878281518110612e9857612e98614ea7565b60209081029190910101518051909150600081516005811115612ebd57612ebd614658565b03612ef1578581606001511115612ee757604051631a783b8d60e01b815260040160405180910390fd5b8060600151860395505b612f058183602001518460400151886132dd565b5050600101612e7c565b50612f198261297e565b8215612f2957612f293384613439565b612f336001600055565b5050505092915050565b612f45613ed5565b8351600003612f69578460405163375c24c160e01b8152600401610d9491906150f3565b6000856001811115612f7d57612f7d614658565b03612fa45780516001600160a01b038316608090910152612f9f868583613ad6565b612fbd565b612faf868583613933565b336020820152604081018390525b8051606001516000036103305760006020820181905281516080015295945050505050565b808210156129a357604051632335530b60e11b815260040160405180910390fd5b6000806000526000825160208403805182604103600060018211613071576040880151606089015160001a9650821561304f57601b8160ff1c0196506001600160ff1b03811660408a01525b8689528985526020600060808760015afa508385528589526040890152506000515b89148915151695508590506131325760408252604486038051604088038051630b135d3f60e11b84528a82526020600060648901868f5afa9850881561312857630b135d3f60e11b60005114613128578b3b156130d957634f7fb80d60e01b60005260046000fd5b60018760410311156130f657638baa579f60e01b60005260046000fd5b640101000000881a61311757630f801e8560e11b6000528760045260246000fd5b632057875960e21b60005260046000fd5b8486529190925290525b50505050806119e557613143613c67565b634f7fb80d60e01b60005260046000fd5b6000428411806131645750428311155b15613190578115613188576040516337bf561360e11b815260040160405180910390fd5b50600061052c565b5060019392505050565b60018360038111156131ae576131ae614658565b1180156131c45750336001600160a01b03821614155b80156131d95750336001600160a01b03831614155b15613270576080880151511580156131f057508651155b156132065761320181868487613caf565b613270565b600061326482633313157060e01b88338d8c8e60405160240161322d9594939291906152b3565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613cfb565b905061155b8187613d10565b5050505050505050565b600083600052602060002060208301835160051b81015b808210156132b957815180841160051b93845260209384185260406000209290910190613291565b505083149050806119e5576040516309bde33960e01b815260040160405180910390fd5b6000845160058111156132f2576132f2614658565b0361334057604084015160208501516001600160a01b0316171561332957604051636ab37ce760e01b815260040160405180910390fd5b61333b84608001518560600151613439565b6119e5565b60018451600581111561335557613355614658565b036133975760408401511561337d57604051636ab37ce760e01b815260040160405180910390fd5b61333b8460200151848660800151876060015186866138d4565b6002845160058111156133ac576133ac614658565b036133d05761333b84602001518486608001518760400151886060015187876127da565b6119e58460200151848660800151876040015188606001518787612827565b600086880361340a576134038686896129a7565b905061342e565b61342b61341887878b6129a7565b61342388888b6129a7565b8686866129e9565b90505b979650505050505050565b613442816138b3565b600080600080600085875af19050806134885761345d613c67565b60405163470c7c1d60e01b81526001600160a01b038416600482015260248101839052604401610d94565b505050565b60186101243510610244356102606102643560400201146004356020146102243561024014161616806134d3576040516339f3e3fd60e01b815260040160405180910390fd5b50565b60018360038111156134ea576134ea614658565b1180156135005750336001600160a01b03821614155b80156135155750336001600160a01b03831614155b156120055761200581868487613caf565b600083815260026020526040902061354184826001806118d1565b50805460ff1661355657613556838584611996565b710100000000000000000000000000000100019055505050565b604080517f000000000000000000000000000000000000000000000000000000000000000060ff60a01b17600090815260208690527f000000000000000000000000000000000000000000000000000000000000000083526055600b20919092526001600160a01b03169050600080600080526020600085876000875af1915060005190508161362657613602613c67565b60405163344f54f560e21b81526001600160a01b0384166004820152602401610d94565b6001600160e01b03198116632671a55160e11b146127d257604051630e7ccd9360e11b8152600481018790526001600160a01b0384166024820152604401610d94565b833b61368457632f8aeb3960e11b6000528360045260246000fd5b6040516323b872dd60e01b6000528360045282602452816044526000806064600080895af18061371e573d156136f8576020601f3d01046020830481600302818311156136df57818303600302610200838002858002030401015b5a6020820110156136f4573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005285600452846024528360445282606452600160845260a46000fd5b50604052505060006060525050565b843b61374857632f8aeb3960e11b6000528460045260246000fd5b60405160805160a05160c051637921219560e11b6000528760045286602452856044528460645260a0608452600060a45260008060c46000808d5af1806137f8573d156137d3576020601f3d01046020860481600302818311156137ba57818303600302610200838002858002030401015b5a6020820110156137cf573d6000803e3d6000fd5b5050505b63f486bc8760e01b600052896004528860245287604452866064528560845260a46000fd5b5060809290925260a05260c05260405250506000606052505050565b6000613821836020015190565b9050818114613488576134888361297e565b6000602088510361386e5750604080885260208089018a9052632671a55160e11b91890191909152604488015260016064880181905261387d565b50606487018051600101908190525b603c60c082028901038781528660208201528560408201528460608201528360808201528260a082015250505050505050505050565b806000036134d35760405163246cf94560e21b815260040160405180910390fd5b6138dd836138b3565b6138e78183613814565b816138fd576138f886868686613d6a565b6127d2565b6127d28282600189898960008a613833565b6064810151604082019060c00260440161392a848383613570565b50506020905250565b61395f565b637fda727960e01b60005260046000fd5b634e487b7160e01b600052601160045260246000fd5b602082018051518451811061397657613976613938565b602081026020860101516060815101516020845101518151811061399c5761399c613938565b602081026020830101516000806020860151156139c3575050606081018051600090915280155b885183518152602084015160208201526040840151604082015260a084015160808201526060812060208c51028c015b808b1015613a9a5760208b019a508a515199508d518a10613a1657613a16613938565b60208a0260208f01015198506020890151156139f357606089510151975060208b510151965087518710613a4c57613a4c613938565b602087026020890101519550606086018051860181511587821060011b1786179550809650506000815250606086208214608084015160a08801511416613a9557613a95613938565b6139f3565b50506060018290528015613ac95760018103613ac15763246cf94560e21b60005260046000fd5b613ac9613949565b5050505050505050505050565b6020820180515184518110613aed57613aed613938565b602081026020860101518051604081015160208551015181518110613b1457613b14613938565b60208102602083010151600080602087015115613b3b575050606081018051600090915280155b8951835181526020840151602082015260408401516040820152865160208c015261012087015160408c015260608120905060208c51028c015b808b1015613c295760208b019a508a515199508d518a10613b9857613b98613938565b60208a0260208f0101519850602089015115613b7557885197506040880151965060208b510151955086518610613bd157613bd1613938565b602086026020880101519450606085018051850181511586821060011b178517945080955050600081525060608520821460408d01516101208a01511460208e01518a51141616613c2457613c24613938565b613b75565b50508160608b5101528015613c595760018103613c515763246cf94560e21b60005260046000fd5b613c59613949565b505050505050505050505050565b3d156116b4576020601f3d01046020604051048160030281831115613c9a57818303600302610200838002858002030401015b5a602082011015613488573d6000803e3d6000fd5b604051602481018490523360448201526001600160a01b038316606482015260848101829052600090613cef9086906303874c7760e21b9060a40161322d565b90506120058185613d10565b6000806000835160208501865afa9392505050565b81613d3957613d1d613c67565b604051633ed4053f60e21b815260048101829052602401610d94565b613d496303874c7760e21b613e73565b156129a357604051633ed4053f60e21b815260048101829052602401610d94565b6040516323b872dd60e01b600052836004528260245281604452602060006064600080895af1803d15601f3d116001600051141617163d15158116613e635780873b151516613e635780613e4e5781613e2d573d15613e07576020601f3d0104602084048160030281831115613dee57818303600302610200838002858002030401015b5a602082011015613e03573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005286600452856024528460445260006064528360845260a46000fd5b639889192360e01b6000528660045285602452846044528360645260846000fd5b632f8aeb3960e11b6000528660045260246000fd5b5050604052505060006060525050565b60008060203d03613e895760206000803e506000515b6001600160e01b031990811692169190911415919050565b6040518060a00160405280613eb4613f18565b81526000602082018190526040820152606080820181905260809091015290565b60408051610100810182526000606082018181526080830182905260a0830182905260c0830182905260e083018290528252602082018190529181019190915290565b60405180610160016040528060006001600160a01b0316815260200160006001600160a01b03168152602001606081526020016060815260200160006003811115613f6557613f65614658565b815260006020820181905260408201819052606082018190526080820181905260a0820181905260c09091015290565b6000815180845260005b81811015613fbb57602081850181015186830182015201613f9f565b81811115613fcd576000602083870101525b50601f01601f19169290920160200192915050565b60208152600061052c6020830184613f95565b60006020828403121561400757600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b03811182821017156140465761404661400e565b60405290565b60405161016081016001600160401b03811182821017156140465761404661400e565b604051601f8201601f191681016001600160401b03811182821017156140975761409761400e565b604052919050565b60006001600160401b038211156140b8576140b861400e565b5060051b60200190565b6001600160a01b03811681146134d357600080fd5b80356140e2816140c2565b919050565b8035600681106140e257600080fd5b600060a0828403121561410857600080fd5b614110614024565b905061411b826140e7565b8152602082013561412b816140c2565b8060208301525060408201356040820152606082013560608201526080820135608082015292915050565b600082601f83011261416757600080fd5b8135602061417c6141778361409f565b61406f565b82815260a0928302850182019282820191908785111561419b57600080fd5b8387015b858110156141be576141b189826140f6565b845292840192810161419f565b5090979650505050505050565b600060c082840312156141dd57600080fd5b60405160c081018181106001600160401b03821117156141ff576141ff61400e565b60405290508061420e836140e7565b8152602083013561421e816140c2565b8060208301525060408301356040820152606083013560608201526080830135608082015260a0830135614251816140c2565b60a0919091015292915050565b600082601f83011261426f57600080fd5b8135602061427f6141778361409f565b82815260c0928302850182019282820191908785111561429e57600080fd5b8387015b858110156141be576142b489826141cb565b84529284019281016142a2565b8035600481106140e257600080fd5b600061016082840312156142e357600080fd5b6142eb61404c565b90506142f6826140d7565b8152614304602083016140d7565b602082015260408201356001600160401b038082111561432357600080fd5b61432f85838601614156565b6040840152606084013591508082111561434857600080fd5b506143558482850161425e565b606083015250614367608083016142c1565b608082015260a082013560a082015260c082013560c082015260e082013560e082015261010080830135818301525061012080830135818301525061014080830135818301525092915050565b80356001600160781b03811681146140e257600080fd5b600082601f8301126143dc57600080fd5b81356001600160401b038111156143f5576143f561400e565b614408601f8201601f191660200161406f565b81815284602083860101111561441d57600080fd5b816020850160208301376000918101602001919091529392505050565b600060a0828403121561444c57600080fd5b614454614024565b905081356001600160401b038082111561446d57600080fd5b614479858386016142d0565b8352614487602085016143b4565b6020840152614498604085016143b4565b604084015260608401359150808211156144b157600080fd5b6144bd858386016143cb565b606084015260808401359150808211156144d657600080fd5b506144e3848285016143cb565b60808301525092915050565b600082601f83011261450057600080fd5b813560206145106141778361409f565b82815260059290921b8401810191818101908684111561452f57600080fd5b8286015b8481101561456e5780356001600160401b038111156145525760008081fd5b6145608986838b010161443a565b845250918301918301614533565b509695505050505050565b60008083601f84011261458b57600080fd5b5081356001600160401b038111156145a257600080fd5b6020830191508360208260051b85010111156145bd57600080fd5b9250929050565b6000806000806000606086880312156145dc57600080fd5b85356001600160401b03808211156145f357600080fd5b6145ff89838a016144ef565b9650602088013591508082111561461557600080fd5b61462189838a01614579565b9096509450604088013591508082111561463a57600080fd5b5061464788828901614579565b969995985093965092949392505050565b634e487b7160e01b600052602160045260246000fd5b6006811061467e5761467e614658565b9052565b61468d82825161466e565b6020818101516001600160a01b0390811691840191909152604080830151908401526060808301519084015260809182015116910152565b600081518084526020808501945080840160005b8381101561471b5781516146ee888251614682565b808401516001600160a01b031660a08901526040015160c088015260e090960195908201906001016146d9565b509495945050505050565b60208152600061052c60208301846146c5565b60006020828403121561474b57600080fd5b81356001600160401b0381111561476157600080fd5b8201610160818503121561052c57600080fd5b60008060008060008060008060008060e08b8d03121561479357600080fd5b8a356001600160401b03808211156147aa57600080fd5b6147b68e838f016144ef565b9b5060208d01359150808211156147cc57600080fd5b6147d88e838f01614579565b909b50995060408d01359150808211156147f157600080fd5b6147fd8e838f01614579565b909950975060608d013591508082111561481657600080fd5b506148238d828e01614579565b90965094505060808b0135925061483c60a08c016140d7565b915060c08b013590509295989b9194979a5092959850565b604080825283519082018190526000906020906060840190828701845b8281101561488f578151151584529284019290840190600101614871565b505050838103828501526148a381866146c5565b9695505050505050565b600080602083850312156148c057600080fd5b82356001600160401b038111156148d657600080fd5b6148e285828601614579565b90969095509350505050565b6000806000806040858703121561490457600080fd5b84356001600160401b038082111561491b57600080fd5b61492788838901614579565b9096509450602087013591508082111561494057600080fd5b5061494d87828801614579565b95989497509550505050565b6000806040838503121561496c57600080fd5b82356001600160401b0381111561498257600080fd5b83016040818603121561499457600080fd5b946020939093013593505050565b6000806000806000608086880312156149ba57600080fd5b85356001600160401b03808211156149d157600080fd5b9087019060a0828a0312156149e557600080fd5b909550602087013590808211156149fb57600080fd5b50614a0888828901614579565b909550935050604086013591506060860135614a23816140c2565b809150509295509295909350565b60008060008060008060008060a0898b031215614a4d57600080fd5b88356001600160401b0380821115614a6457600080fd5b614a708c838d01614579565b909a50985060208b0135915080821115614a8957600080fd5b614a958c838d01614579565b909850965060408b0135915080821115614aae57600080fd5b50614abb8b828c01614579565b999c989b509699959896976060870135966080013595509350505050565b600060208284031215614aeb57600080fd5b813561052c816140c2565b606081526000614b096060830186613f95565b6020830194909452506001600160a01b0391909116604090910152919050565b600060208284031215614b3b57600080fd5b81356001600160401b03811115614b5157600080fd5b8201610240818503121561052c57600080fd5b6000614b726141778461409f565b83815260208082019190600586811b860136811115614b9057600080fd5b865b81811015614c8b5780356001600160401b0380821115614bb25760008081fd5b818a01915060a08236031215614bc85760008081fd5b614bd0614024565b823581528683013560028110614be65760008081fd5b81880152604083810135908201526060808401359082015260808084013583811115614c125760008081fd5b939093019236601f850112614c2957600092508283fd5b83359250614c396141778461409f565b83815292871b84018801928881019036851115614c565760008081fd5b948901945b84861015614c7457853582529489019490890190614c5b565b918301919091525088525050948301948301614b92565b5092979650505050505050565b6000808335601e19843603018112614caf57600080fd5b8301803591506001600160401b03821115614cc957600080fd5b602001915060a0810236038213156145bd57600080fd5b600060a08284031215614cf257600080fd5b61052c83836140f6565b6000808335601e19843603018112614d1357600080fd5b8301803591506001600160401b03821115614d2d57600080fd5b602001915060c0810236038213156145bd57600080fd5b600060c08284031215614d5657600080fd5b61052c83836141cb565b600060208284031215614d7257600080fd5b61052c826142c1565b60006104d4368361443a565b600060408284031215614d9957600080fd5b604051604081018181106001600160401b0382111715614dbb57614dbb61400e565b604052823581526020928301359281019290925250919050565b6000614de36141778461409f565b80848252602080830192508560051b850136811115614e0157600080fd5b855b81811015614e9b5780356001600160401b03811115614e225760008081fd5b870136601f820112614e345760008081fd5b8035614e426141778261409f565b81815260069190911b82018501908581019036831115614e625760008081fd5b928601925b82841015614e8b57614e793685614d87565b82528682019150604084019350614e67565b8852505050938201938201614e03565b50919695505050505050565b634e487b7160e01b600052603260045260246000fd5b60008235603e19833603018112614ed357600080fd5b9190910192915050565b6000823561015e19833603018112614ed357600080fd5b60006104d436836142d0565b6000808335601e19843603018112614f1757600080fd5b8301803591506001600160401b03821115614f3157600080fd5b6020019150368190038213156145bd57600080fd5b6000808335601e19843603018112614f5d57600080fd5b8301803591506001600160401b03821115614f7757600080fd5b6020019150600681901b36038213156145bd57600080fd5b634e487b7160e01b600052601160045260246000fd5b60008219821115614fb857614fb8614f8f565b500190565b6000816000190483118215151615614fd757614fd7614f8f565b500290565b600081518084526020808501945080840160005b8381101561471b57615003878351614682565b60a0969096019590820190600101614ff0565b60006080808301878452602060018060a01b03808916828701526040848188015283895180865260a089019150848b01955060005b8181101561508c57865161506084825161466e565b80870151861684880152848101518585015260609081015190840152958501959187019160010161504b565b505087810360608901526150a0818a614fdc565b9c9b505050505050505050505050565b6000828210156150c2576150c2614f8f565b500390565b6000604082840312156150d957600080fd5b61052c8383614d87565b6002811061467e5761467e614658565b602081016104d482846150e3565b600081518084526020808501945080840160005b8381101561471b57815161512a88825161466e565b838101516001600160a01b03168885015260408082015190890152606080820151908901526080908101519088015260a09096019590820190600101615115565b600081518084526020808501945080840160005b8381101561471b57815161519488825161466e565b808401516001600160a01b0390811689860152604080830151908a0152606080830151908a0152608080830151908a015260a091820151169088015260c0909601959082019060010161517f565b6004811061467e5761467e614658565b600081518084526020808501945080840160005b8381101561471b57815187529582019590820190600101615206565b600081518084526020808501808196508360051b8101915082860160005b858110156152a6578284038952815160a08151865286820151615265888801826150e3565b506040828101519087015260608083015190870152608091820151918601819052615292818701836151f2565b9a87019a9550505090840190600101615240565b5091979650505050505050565b85815260018060a01b038516602082015260a060408201526000610140855160a0808501526152ed82850182516001600160a01b03169052565b6020810151610160615309818701836001600160a01b03169052565b6040830151915080610180870152506153266102a0860182615101565b9050606082015161013f19868303016101a0870152615345828261516b565b915050608082015161535b6101c08701826151e2565b5060a08201516101e086015260c082015161020086015260e082015161022086015261010080830151610240870152610120808401516102608801528484015161028088015260208a015194506153bd60c08801866001600160781b03169052565b60408a01516001600160781b031660e088015260608a0151878403609f19908101848a015290955093506153f18386613f95565b945060808a01519250838786030181880152505061540f8382613f95565b92505050828103606084015261542581866151f2565b905082810360808401526154398185615222565b9897505050505050505056fea2646970667358221220a843c280870a495a0a1fa864f6216677d1d791eaa4c8cf4ba0d4e22b551d2aef64736f6c634300080d0033",
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

// GetCounter is a free data retrieval call binding the contract method 0xf07ec373.
//
// Solidity: function getCounter(address offerer) view returns(uint256 counter)
func (_Seaport *SeaportCaller) GetCounter(opts *bind.CallOpts, offerer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Seaport.contract.Call(opts, &out, "getCounter", offerer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCounter is a free data retrieval call binding the contract method 0xf07ec373.
//
// Solidity: function getCounter(address offerer) view returns(uint256 counter)
func (_Seaport *SeaportSession) GetCounter(offerer common.Address) (*big.Int, error) {
	return _Seaport.Contract.GetCounter(&_Seaport.CallOpts, offerer)
}

// GetCounter is a free data retrieval call binding the contract method 0xf07ec373.
//
// Solidity: function getCounter(address offerer) view returns(uint256 counter)
func (_Seaport *SeaportCallerSession) GetCounter(offerer common.Address) (*big.Int, error) {
	return _Seaport.Contract.GetCounter(&_Seaport.CallOpts, offerer)
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

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xe7acab24.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) advancedOrder, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, bytes32 fulfillerConduitKey, address recipient) payable returns(bool fulfilled)
func (_Seaport *SeaportTransactor) FulfillAdvancedOrder(opts *bind.TransactOpts, advancedOrder AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillerConduitKey [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "fulfillAdvancedOrder", advancedOrder, criteriaResolvers, fulfillerConduitKey, recipient)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xe7acab24.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) advancedOrder, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, bytes32 fulfillerConduitKey, address recipient) payable returns(bool fulfilled)
func (_Seaport *SeaportSession) FulfillAdvancedOrder(advancedOrder AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillerConduitKey [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillAdvancedOrder(&_Seaport.TransactOpts, advancedOrder, criteriaResolvers, fulfillerConduitKey, recipient)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xe7acab24.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) advancedOrder, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, bytes32 fulfillerConduitKey, address recipient) payable returns(bool fulfilled)
func (_Seaport *SeaportTransactorSession) FulfillAdvancedOrder(advancedOrder AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillerConduitKey [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillAdvancedOrder(&_Seaport.TransactOpts, advancedOrder, criteriaResolvers, fulfillerConduitKey, recipient)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0x87201b41.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, address recipient, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Seaport *SeaportTransactor) FulfillAvailableAdvancedOrders(opts *bind.TransactOpts, advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, recipient common.Address, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "fulfillAvailableAdvancedOrders", advancedOrders, criteriaResolvers, offerFulfillments, considerationFulfillments, fulfillerConduitKey, recipient, maximumFulfilled)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0x87201b41.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, address recipient, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Seaport *SeaportSession) FulfillAvailableAdvancedOrders(advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, recipient common.Address, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillAvailableAdvancedOrders(&_Seaport.TransactOpts, advancedOrders, criteriaResolvers, offerFulfillments, considerationFulfillments, fulfillerConduitKey, recipient, maximumFulfilled)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0x87201b41.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, address recipient, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Seaport *SeaportTransactorSession) FulfillAvailableAdvancedOrders(advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, recipient common.Address, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillAvailableAdvancedOrders(&_Seaport.TransactOpts, advancedOrders, criteriaResolvers, offerFulfillments, considerationFulfillments, fulfillerConduitKey, recipient, maximumFulfilled)
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

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns(uint256 newCounter)
func (_Seaport *SeaportTransactor) IncrementCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "incrementCounter")
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns(uint256 newCounter)
func (_Seaport *SeaportSession) IncrementCounter() (*types.Transaction, error) {
	return _Seaport.Contract.IncrementCounter(&_Seaport.TransactOpts)
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns(uint256 newCounter)
func (_Seaport *SeaportTransactorSession) IncrementCounter() (*types.Transaction, error) {
	return _Seaport.Contract.IncrementCounter(&_Seaport.TransactOpts)
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

// SeaportCounterIncrementedIterator is returned from FilterCounterIncremented and is used to iterate over the raw logs and unpacked data for CounterIncremented events raised by the Seaport contract.
type SeaportCounterIncrementedIterator struct {
	Event *SeaportCounterIncremented // Event containing the contract specifics and raw log

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
func (it *SeaportCounterIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeaportCounterIncremented)
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
		it.Event = new(SeaportCounterIncremented)
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
func (it *SeaportCounterIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeaportCounterIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeaportCounterIncremented represents a CounterIncremented event raised by the Seaport contract.
type SeaportCounterIncremented struct {
	NewCounter *big.Int
	Offerer    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCounterIncremented is a free log retrieval operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_Seaport *SeaportFilterer) FilterCounterIncremented(opts *bind.FilterOpts, offerer []common.Address) (*SeaportCounterIncrementedIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Seaport.contract.FilterLogs(opts, "CounterIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return &SeaportCounterIncrementedIterator{contract: _Seaport.contract, event: "CounterIncremented", logs: logs, sub: sub}, nil
}

// WatchCounterIncremented is a free log subscription operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_Seaport *SeaportFilterer) WatchCounterIncremented(opts *bind.WatchOpts, sink chan<- *SeaportCounterIncremented, offerer []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Seaport.contract.WatchLogs(opts, "CounterIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeaportCounterIncremented)
				if err := _Seaport.contract.UnpackLog(event, "CounterIncremented", log); err != nil {
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
func (_Seaport *SeaportFilterer) ParseCounterIncremented(log types.Log) (*SeaportCounterIncremented, error) {
	event := new(SeaportCounterIncremented)
	if err := _Seaport.contract.UnpackLog(event, "CounterIncremented", log); err != nil {
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
	Recipient     common.Address
	Offer         []SpentItem
	Consideration []ReceivedItem
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address recipient, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
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
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address recipient, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
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
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address recipient, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
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
