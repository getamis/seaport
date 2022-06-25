// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Consideration

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

// ConsiderationMetaData contains all meta data concerning the Consideration contract.
var ConsiderationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduitController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadContractSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadFraction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BadReturnValueFromERC20OnTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"BadSignatureV\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConsiderationCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortfallAmount\",\"type\":\"uint256\"}],\"name\":\"ConsiderationNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CriteriaNotEnabledForItem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"identifiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"ERC1155BatchTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EtherTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InexactFraction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientEtherSupplied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Invalid1155BatchTransferEncoding\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBasicOrderParameterEncoding\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidCallToConduit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCanceller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidConduit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidERC721TransferAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFulfillmentComponentData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNativeOfferItem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidRestrictedOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchedFulfillmentOfferAndConsiderationComponents\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"}],\"name\":\"MissingFulfillmentComponentOnAggregation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingItemAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingOriginalConsiderationItems\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NoContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoReentrantCalls\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSpecifiedOrdersAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferAndConsiderationRequiredOnFulfillment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderAlreadyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderIsCancelled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderPartiallyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PartialFillsNotEnabledForOrder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnresolvedConsiderationCriteria\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnresolvedOfferCriteria\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnusedItemParameters\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCounter\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"}],\"name\":\"CounterIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSpentItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structReceivedItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"}],\"name\":\"OrderValidated\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder\",\"name\":\"advancedOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"criteriaResolvers\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"fulfillAdvancedOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder[]\",\"name\":\"advancedOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"criteriaResolvers\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"offerFulfillments\",\"type\":\"tuple[][]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"considerationFulfillments\",\"type\":\"tuple[][]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maximumFulfilled\",\"type\":\"uint256\"}],\"name\":\"fulfillAvailableAdvancedOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"availableOrders\",\"type\":\"bool[]\"},{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"offerFulfillments\",\"type\":\"tuple[][]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"considerationFulfillments\",\"type\":\"tuple[][]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maximumFulfilled\",\"type\":\"uint256\"}],\"name\":\"fulfillAvailableOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"availableOrders\",\"type\":\"bool[]\"},{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"considerationToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"considerationIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"offerToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offerIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumBasicOrderType\",\"name\":\"basicOrderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offererConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalAdditionalRecipients\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structAdditionalRecipient[]\",\"name\":\"additionalRecipients\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBasicOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"}],\"name\":\"fulfillBasicOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"}],\"name\":\"fulfillOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"}],\"name\":\"getCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"getOrderStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidated\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSize\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newCounter\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"information\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"conduitController\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder[]\",\"name\":\"advancedOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"criteriaResolvers\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"offerComponents\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"considerationComponents\",\"type\":\"tuple[]\"}],\"internalType\":\"structFulfillment[]\",\"name\":\"fulfillments\",\"type\":\"tuple[]\"}],\"name\":\"matchAdvancedOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"offerComponents\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"considerationComponents\",\"type\":\"tuple[]\"}],\"internalType\":\"structFulfillment[]\",\"name\":\"fulfillments\",\"type\":\"tuple[]\"}],\"name\":\"matchOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"contractName\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"validate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"validated\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101c06040523480156200001257600080fd5b5060405162005ba538038062005ba58339810160408190526200003591620005ca565b8080808080808080806200004862000130565b610120526101005260e05260c081815260a0838152608085815246610140819052604080516020818101979097528082019890985260608801969096529086015230858201528351808603909101815293019091528151910120610160526001600160a01b03811661018081905260408051630a96ad3960e01b81528151630a96ad39926004808401939192918290030181865afa158015620000ef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001159190620005fc565b506101a05250506001600055506200068c9650505050505050565b600080808080806200016260408051808201909152600d81526c21b7b739b4b232b930ba34b7b760991b602082015290565b80516020918201206040805180820182526003815262312e3160e81b90840152519097507f722c0e0c80487266e8c6a45e3a1a803aab23378a9c32e6ebe029d4fad7bfc96596506000916200026791016909ecccccae492e8cada560b31b81526e1d5a5b9d0e081a5d195b551e5c194b608a1b600a8201526d1859191c995cdcc81d1bdad95b8b60921b60198201527f75696e74323536206964656e7469666965724f7243726974657269612c00000060278201527f75696e74323536207374617274416d6f756e742c0000000000000000000000006044820152701d5a5b9d0c8d4d88195b99105b5bdd5b9d607a1b6058820152602960f81b6069820152606a0190565b60408051601f1981840301815282825271086dedce6d2c8cae4c2e8d2dedc92e8cada560731b60208401526e1d5a5b9d0e081a5d195b551e5c194b608a1b60328401526d1859191c995cdcc81d1bdad95b8b60921b60418401527f75696e74323536206964656e7469666965724f7243726974657269612c000000604f8401527f75696e74323536207374617274416d6f756e742c000000000000000000000000606c840152711d5a5b9d0c8d4d88195b99105b5bdd5b9d0b60721b6080840152701859191c995cdcc81c9958da5c1a595b9d607a1b6092840152602960f81b60a384018190528251808503608401815260a485019093526f09ee4c8cae486dedae0dedccadce8e6560831b60c48501526f1859191c995cdcc81bd999995c995c8b60821b60d48501526c1859191c995cdcc81e9bdb994b609a1b60e48501527113d999995c925d195b56d7481bd999995c8b60721b60f18501527f436f6e73696465726174696f6e4974656d5b5d20636f6e73696465726174696f610103850152611b8b60f21b6101238501526f1d5a5b9d0e081bdc99195c951e5c194b60821b610125850152711d5a5b9d0c8d4d881cdd185c9d151a5b594b60721b6101358501526f1d5a5b9d0c8d4d88195b99151a5b594b60821b61014785015270189e5d195ccccc881e9bdb9952185cda0b607a1b6101578501526c1d5a5b9d0c8d4d881cd85b1d0b609a1b6101688501527f6279746573333220636f6e647569744b65792c000000000000000000000000006101758501526e3ab4b73a191a9b1031b7bab73a32b960891b6101888501526101978401529250906000906101980160408051601f19818403018152908290526c08a92a06e626488dedac2d2dc5609b1b60208301526b1cdd1c9a5b99c81b985b594b60a21b602d8301526e1cdd1c9a5b99c81d995c9cda5bdb8b608a1b60398301526f1d5a5b9d0c8d4d8818da185a5b92590b60821b60488301527f6164647265737320766572696679696e67436f6e7472616374000000000000006058830152602960f81b6071830152915060720160408051601f19818403018152908290528051602091820120855186830120855186840120919a5098509650620005a791839185918791016200065f565b604051602081830303815290604052805190602001209350505050909192939495565b600060208284031215620005dd57600080fd5b81516001600160a01b0381168114620005f557600080fd5b9392505050565b600080604083850312156200061057600080fd5b505080516020909101519092909150565b6000815160005b8181101562000644576020818501810151868301520162000628565b8181111562000654576000828601525b509290920192915050565b6000620006836200067c62000675848862000621565b8662000621565b8462000621565b95945050505050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516154816200072460003960006135ad015260008181610d27015261357c01526000612319015260006122490152600081816108b6015261252701526000818161084501526123850152600081816107de01526124a001526000612277015260006122c50152600061229d01526154816000f3fe6080604052600436106100e85760003560e01c8063a81744041161008a578063f07ec37311610059578063f07ec37314610290578063f47b7740146102b0578063fb0f3ee1146102d4578063fd9f1e10146102e757600080fd5b8063a817440414610244578063b3a34c4c14610257578063e7acab241461026a578063ed98a5741461027d57600080fd5b80635b34b966116100c65780635b34b966146101b057806379df72bd146101d357806387201b41146101f3578063881477321461021457600080fd5b806306fdde03146100ed57806346423aa71461011857806355944a4214610190575b600080fd5b3480156100f957600080fd5b50610102610307565b60405161010f9190613fe8565b60405180910390f35b34801561012457600080fd5b5061016e610133366004613ffb565b60009081526002602052604090205460ff808216926101008304909116916001600160781b03620100008204811692600160881b9092041690565b604080519415158552921515602085015291830152606082015260800161010f565b6101a361019e3660046145ca565b610316565b60405161010f919061472c565b3480156101bc57600080fd5b506101c5610339565b60405190815260200161010f565b3480156101df57600080fd5b506101c56101ee36600461473f565b610343565b61020661020136600461477a565b6104da565b60405161010f92919061485a565b34801561022057600080fd5b5061023461022f3660046148b3565b610520565b604051901515815260200161010f565b6101a36102523660046148f4565b610533565b61023461026536600461495f565b6105b1565b6102346102783660046149a8565b610623565b61020661028b366004614a37565b610657565b34801561029c57600080fd5b506101c56102ab366004614adf565b6106e1565b3480156102bc57600080fd5b506102c56106ff565b60405161010f93929190614afc565b6102346102e2366004614b2f565b610717565b3480156102f357600080fd5b506102346103023660046148b3565b610722565b606061031161072e565b905090565b606061032d866103268688614b6a565b858561074c565b90505b95945050505050565b6000610311610769565b6040805161016081019091526000906104d490806103646020860186614adf565b6001600160a01b031681526020018460200160208101906103859190614adf565b6001600160a01b031681526020016103a06040860186614c9e565b808060200260200160405190810160405280939291908181526020016000905b828210156103ec576103dd60a08302860136819003810190614ce6565b815260200190600101906103c0565b50505091835250506020016104046060860186614d02565b808060200260200160405190810160405280939291908181526020016000905b828210156104505761044160c08302860136819003810190614d4a565b81526020019060010190610424565b505050918352505060200161046b60a0860160808701614d66565b600381111561047c5761047c61465e565b81526020018460a0013581526020018460c0013581526020018460e001358152602001846101000135815260200184610120013581526020018480606001906104c59190614d02565b909152506101408401356107c6565b92915050565b60608061050d8c6104eb8c8e614b6a565b8b8b8b8b8b6001600160a01b038c1615610505578b610507565b335b8b61090c565b915091509a509a98505050505050505050565b600061052c838361094e565b9392505050565b60606105a66105428686610ab3565b604080516000808252602082019092529061059e565b61058b6040805160a081019091526000808252602082019081526020016000815260200160008152602001606081525090565b8152602001906001900390816105585790505b50858561074c565b90505b949350505050565b600061052c6105bf84610b6e565b604080516000808252602082019092529061061b565b6106086040805160a081019091526000808252602082019081526020016000815260200160008152602001606081525090565b8152602001906001900390816105d55790505b508433610c19565b600061032d61063187614d81565b61063b8688614b6a565b856001600160a01b038616156106515785610c19565b33610c19565b6060806106d06106678b8b610ab3565b60408051600080825260208201909252906106c3565b6106b06040805160a081019091526000808252602082019081526020016000815260200160008152602001606081525090565b81526020019060019003908161067d5790505b508a8a8a8a8a338b61090c565b915091509850989650505050505050565b6001600160a01b0381166000908152600160205260408120546104d4565b606060008061070c610d06565b925092509250909192565b60006104d482610d67565b600061052c8383611027565b6060602080526d0d436f6e73696465726174696f6e604d5260606020f35b606061075e858560018851600061122d565b6105a685848461156c565b6000610773611697565b503360008181526001602081815260409283902080549092019182905591518181529092917f721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f910160405180910390a290565b610140820151604080519084015180516000939284927f000000000000000000000000000000000000000000000000000000000000000092602090910190845b81811015610833578251601f1901805186825260c082208652905260209384019390920191600101610806565b506020810260405120945050505060007f00000000000000000000000000000000000000000000000000000000000000009150604051602060608901510160005b868110156108a1578151601f1901805186825260e082208552905260209283019290910190600101610874565b505060408051602087029020601f198a0180517f00000000000000000000000000000000000000000000000000000000000000008252928b01805197815260608c018051938152610140909c019a8b5261018082209390915295909552939097525050925250919050565b60608061091d8b8b6000868861122d565b61093c8b61092b8a8c614ddb565b610935898b614ddb565b88886116bc565b909c909b509950505050505050505050565b6000610958611697565b6000808084815b81811015610aa5573688888381811061097a5761097a614ead565b905060200281019061098c9190614ec3565b9050366109998280614ee3565b90506109a86020820182614adf565b94506109bb6109b682614efa565b61189c565b600081815260026020526040812098509096506109dd908790899060016118d7565b50865460ff16610a9b57610a3385876109f96020860186614f06565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061199c92505050565b865460ff19166001178755610a4e6040820160208301614adf565b6001600160a01b0316856001600160a01b03167ffde361574a066b44b3b5fe98a87108b7565e327327954c4faeea56a4e6491a0a88604051610a9291815260200190565b60405180910390a35b505060010161095f565b506001979650505050505050565b606081806001600160401b03811115610ace57610ace614014565b604051908082528060200260200182016040528015610b0757816020015b610af4613ea7565b815260200190600190039081610aec5790505b50915060005b81811015610b6657610b41858583818110610b2a57610b2a614ead565b9050602002810190610b3c9190614ec3565b610b6e565b838281518110610b5357610b53614ead565b6020908102919091010152600101610b0d565b505092915050565b610b76613ea7565b6040805160a0810190915280610b8c8480614ee3565b610b9590614efa565b815260200160016001600160781b0316815260200160016001600160781b03168152602001838060200190610bca9190614f06565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509385525050604080516020818101909252928352909201525092915050565b6000610c236119f1565b60606000806000610c378989600187611a00565b604080516001808252818301909252939650919450925060009190816020015b610c5f613ea7565b815260200190600190039081610c575790505090508981600081518110610c8857610c88614ead565b6020026020010181905250610c9d818a611cb9565b600081600081518110610cb257610cb2614ead565b6020026020010151600001519050610ccd8185858c8c612012565b610ceb85826000015183602001518b856040015186606001516121e0565b610cf56001600055565b5060019a9950505050505050505050565b6060600080610d13612245565b6040805160038082528183019092529193507f00000000000000000000000000000000000000000000000000000000000000009250602082018180368337505062312e3160e81b6020830152509391925090565b600061012435600281901c90600316600182113415811480610da357604051630a61be9f60e41b81523460048201526024015b60405180910390fd5b5060008060008060038711915060a082026024013593506002871460028811600289030201905060018101820260028615028801039250610de889878388888861233b565b506101c46020820201356000856005811115610e0657610e0661465e565b03610ea95760208901803590610e1c908b614adf565b6001600160a01b03161715610e4457604051636ab37ce760e01b815260040160405180910390fd5b610e7983610e5860c08c0160a08d01614adf565b610e6860808d0160608e01614adf565b338d60c001358e60e0013587612671565b610ea460408a0135610e9160808c0160608d01614adf565b610e9f6102008d018d614f4c565b612725565b61100e565b6040805160208082528183019092526000916020820181803683370190505090506002886005811115610ede57610ede61465e565b03610f1d57610f18610ef660c08c0160a08d01614adf565b610f0660808d0160608e01614adf565b338d60c001358e60e0013587876127e0565b610fe8565b6003886005811115610f3157610f3161465e565b03610f6b57610f18610f4960c08c0160a08d01614adf565b610f5960808d0160608e01614adf565b338d60c001358e60e00135878761282d565b6004886005811115610f7f57610f7f61465e565b03610fb657610f18610f9460208c018c614adf565b33610fa560808e0160608f01614adf565b8d602001358e6040013587876127e0565b610fe8610fc660208c018c614adf565b33610fd760808e0160608f01614adf565b8d602001358e60400135878761282d565b611003610ffb60808c0160608d01614adf565b8b8584612863565b61100c81612984565b505b6110186001600055565b50600198975050505050505050565b6000611031611697565b6000808084815b81811015610aa5573688888381811061105357611053614ead565b90506020028101906110659190614ee3565b90506110746020820182614adf565b94506110866040820160208301614adf565b9350336001600160a01b038616148015906110aa5750336001600160a01b03851614155b156110c85760405163203b1cdd60e21b815260040160405180910390fd5b60006111b7604051806101600160405280886001600160a01b03168152602001876001600160a01b031681526020018480604001906111079190614c9e565b808060200260200160405190810160405280939291908181526020016000905b828210156111535761114460a08302860136819003810190614ce6565b81526020019060010190611127565b505050918352505060200161116b6060860186614d02565b808060200260200160405190810160405280939291908181526020016000905b82821015610450576111a860c08302860136819003810190614d4a565b8152602001906001019061118b565b60008181526002602052604090819020805461ffff191661010017815590519098509091506001600160a01b0380871691908816907f6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d9061121b9085815260200190565b60405180910390a35050600101611038565b6112356119f1565b84516000816001600160401b0381111561125157611251614014565b60405190808252806020026020018201604052801561127a578160200160208202803683370190505b506000808252909150601d6045823560e01c061160011b905b838110156114a35760008982815181106112af576112af614ead565b60200260200101519050866000036112d457600060209091015260018101835261149b565b60008060006112e5848d8d8a611a00565b9250925092506001850187528160000361130c57505060006020909201919091525061149b565b8287868151811061131f5761131f614ead565b6020908102919091010152835160a081015160c08201516040909201518051600019909d019c91929160005b818110156113e757600083828151811061136757611367614ead565b602002602001015190508051158c179b506000611389898984608001516129ad565b905081608001518260600151036113a657606082018190526113bb565b6113b5898984606001516129ad565b60608301525b6080820181905260608201516113d59082898960006129ef565b6060909201919091525060010161134b565b50875160600151805160005b8181101561148f57600083828151811061140f5761140f614ead565b60200260200101519050600061142a8b8b84608001516129ad565b90508160800151826060015103611447576060820181905261145c565b6114568b8b84606001516129ad565b60608301525b60808201819052606082015161147690828b8b60016129ef565b60608301525060a08101516080909101526001016113f3565b50505050505050505050505b600101611293565b50806003036114c5576040516312d3f5a360e01b815260040160405180910390fd5b6114cf8888611cb9565b60005b83811015611561576000801b8382815181106114f0576114f0614ead565b6020026020010151031561155957600089828151811061151257611512614ead565b602002602001015160000151905061155784838151811061153557611535614ead565b60200260200101518260000151836020015189856040015186606001516121e0565b505b6001016114d2565b505050505050505050565b606081806001600160401b0381111561158757611587614014565b6040519080825280602002602001820160405280156115c057816020015b6115ad613edb565b8152602001906001900390816115a55790505b5091506000805b8281101561167557368686838181106115e2576115e2614ead565b90506020028101906115f49190614ec3565b90506000611618896116068480614f4c565b6116136020870187614f4c565b612a44565b905080602001516001600160a01b03168160000151608001516001600160a01b03160361164a5783600101935061166b565b80868585038151811061165f5761165f614ead565b60200260200101819052505b50506001016115c7565b508015611683578083510383525b5061168e8583612d29565b50509392505050565b6001600054146116ba57604051637fa8a98760e01b815260040160405180910390fd5b565b8351835160609182916116cf8183614fab565b6001600160401b038111156116e6576116e6614014565b60405190808252806020026020018201604052801561171f57816020015b61170c613edb565b8152602001906001900390816117045790505b5092506000805b838110156117b95760008a828151811061174257611742614ead565b60200260200101519050600061175c8d6000848d8d612f43565b905080602001516001600160a01b03168160000151608001516001600160a01b03160361178e578360010193506117af565b8087858503815181106117a3576117a3614ead565b60200260200101819052505b5050600101611726565b5060005b828110156118535760008982815181106117d9576117d9614ead565b6020026020010151905060006117f48d6001848d6000612f43565b905080602001516001600160a01b03168160000151608001516001600160a01b03160361182657836001019350611849565b808785888601038151811061183d5761183d614ead565b60200260200101819052505b50506001016117bd565b508015611861578084510384525b5082516000036118845760405163d5da9a1b60e01b815260040160405180910390fd5b61188e8984612d29565b935050509550959350505050565b60006118b2826060015151836101400151612fe8565b81516001600160a01b03166000908152600160205260409020546104d49083906107c6565b8254600090610100900460ff161561191357811561190b57604051630694555d60e21b815260048101869052602401610d9a565b5060006105a9565b83546201000090046001600160781b0316801561199057831561194c5760405163ee9e0e6360e01b815260048101879052602401610d9a565b8454600160881b90046001600160781b03168110611990578215611986576040516310fda3e160e01b815260048101879052602401610d9a565b60009150506105a9565b50600195945050505050565b336001600160a01b038416036119b157505050565b60006119de6119be612245565b61190160f01b600090815260029190915260228581526042822091905290565b90506119eb848284613009565b50505050565b6119f9611697565b6002600055565b60008060008087600001519050611a208160a001518260c001518861315a565b611a34575060009250829150819050611caf565b602088015160408901516001600160781b03918216911680821180611a57575081155b15611a7557604051632d02959960e11b815260040160405180910390fd5b8082108015611a8957506080830151600116155b15611aa75760405163a11b63ff60e01b815260040160405180910390fd5b611ab08361189c565b9550611ad28a8a89898760e00151886080015189600001518a602001516131a0565b600086815260026020526040812090611aef90889083908c6118d7565b611b04575060009450849350611caf92505050565b805460ff16611b2057611b208460000151888d6060015161199c565b80546001600160781b03620100008204811691600160881b9004168015611c695783600103611b5457809450809350611b80565b838114611b8057611b658483614fc3565b9150611b718186614fc3565b9450611b7d8185614fc3565b93505b83611b8b8684614fab565b1115611b975781840394505b611ba18583614fab565b91506001600160781b0384116001600160781b0383111715611c2a57611bdc565b60005b8215611bd657908290069190611bc5565b50919050565b611bef611be98584611bc2565b86611bc2565b80150194859004949384900493909104906001600160781b038083119085111715611c2a57634e487b7160e01b600052601160045260246000fd5b82546001600160781b03858116600160881b026001600160881b0391851662010000026001600160881b03199093169290921760011716178355611ca4565b82546001600160781b03858116600160881b026001600160881b0391881662010000026001600160881b031990931692909217600117161783555b509295509093505050505b9450945094915050565b8051825160005b82811015611f00576000848281518110611cdc57611cdc614ead565b60200260200101519050600081600001519050838110611d0f576040516321a561b160e21b815260040160405180910390fd5b868181518110611d2157611d21614ead565b6020026020010151602001516001600160781b0316600003611d44575050611ef8565b6000878281518110611d5857611d58614ead565b60209081029190910101515160408401519091506000808086602001516001811115611d8657611d8661465e565b03611e2057604084015180518410611db157604051635fd9fc6760e11b815260040160405180910390fd5b6000818581518110611dc557611dc5614ead565b602090810291909101015180516040820151909550935090506004841460030381816005811115611df857611df861465e565b90816005811115611e0b57611e0b61465e565b90525050606088015160409091015250611eb1565b606084015180518410611e46576040516330446bef60e11b815260040160405180910390fd5b6000818581518110611e5a57611e5a614ead565b602090810291909101015180516040820151909550935090506004841460030381816005811115611e8d57611e8d61465e565b90816005811115611ea057611ea061465e565b905250506060880151604090910152505b611ebb8260031090565b611ed857604051634a75b57b60e11b815260040160405180910390fd5b8015611ef157611ef18660600151828860800151613280565b5050505050505b600101611cc0565b5060005b8181101561200b576000858281518110611f2057611f20614ead565b6020026020010151905080602001516001600160781b0316600003611f455750612003565b805160608101515160005b81811015611fac57611f8383606001518281518110611f7157611f71614ead565b60200260200101516000015160031090565b15611fa45760405160016202297360e61b0319815260040160405180910390fd5b600101611f50565b505060408101515160005b81811015611ffe57611fd883604001518281518110611f7157611f71614ead565b15611ff65760405163a6cfc67360e01b815260040160405180910390fd5b600101611fb7565b505050505b600101611f04565b5050505050565b60a085015160c08601516040805160208082528183019092526000916020820181803683375050506040890151519091506132e39060005b818110156121005760008b60400151828151811061206a5761206a614ead565b60200260200101519050600060058111156120875761208761465e565b8151600581111561209a5761209a61465e565b036120b8576040516312d3f5a360e01b815260040160405180910390fd5b60006120d2826060015183608001518e8e8c8c60006133f5565b606083015250608081018890528b516101208d01516120f79183918863ffffffff8916565b5060010161204a565b50505060608801515134906132e39060005b818110156121c45760008c60600151828151811061213257612132614ead565b602002602001015190506000612156826060015183608001518f8f8d8d60016133f5565b6060830181905260a08301516080840152905060008251600581111561217e5761217e61465e565b036121aa57858111156121a457604051631a783b8d60e01b815260040160405180910390fd5b80860395505b6121ba82338d8a8963ffffffff16565b5050600101612112565b5050506121d082612984565b801561156157611561338261343f565b60608290506060829050856001600160a01b0316876001600160a01b03167f9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f318a888686604051612233949392919061501c565b60405180910390a35050505050505050565b60007f0000000000000000000000000000000000000000000000000000000000000000461461231657610311604080517f000000000000000000000000000000000000000000000000000000000000000060208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b507f000000000000000000000000000000000000000000000000000000000000000090565b6123436119f1565b612359866101200135876101400135600161315a565b50612362613493565b612380612373610200880188614f4c565b9050876101e00135612fe8565b6000807f00000000000000000000000000000000000000000000000000000000000000009050806080528560a0526060602460c037604060646101203760e06080206101605261026435602081026102a0016001610264350181526020810190508781526080602460208301376101608760a0528660c052600060e05261020435925060005b83811015612451578060400261028401602081610100376040816101203760208301925060e0608020835260a084019350898452886020850152604081606086013750600101612406565b60206001850102610160206060526102643593505b83811015612497578060400261028401915060a0830192508883528760208401526040826060850137600101612466565b505050505060007f00000000000000000000000000000000000000000000000000000000000000009050806080528260a052606060c460c03760206101046101203760c0608020600052602060002060e052602061026435026102000160018152836020820152606060c4604083013750506084356001600160a01b0381166000908152600160205260408120547f000000000000000000000000000000000000000000000000000000000000000060808190529091506040608460a03760605161010052886101205260a061014461014037816101e0526101806080209350505050602061026435026101800181815233602082015260806040820152610120606082015260a061026435026101e00160a4356084357f9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f318385a35050600060605261260c81886101600135888a60600160208101906125f79190614adf565b61260760a08d0160808e01614adf565b6134dc565b6126688161262060808a0160608b01614adf565b61262e6102208b018b614f06565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061352c92505050565b50505050505050565b80156126cd5760006040519050632671a55160e11b815260206004820152600160248201528760448201528660648201528560848201528460a48201528360c48201528260e48201526126c78282610104613576565b50612668565b60028760058111156126e1576126e161465e565b0361271857816001146127075760405163efcc00b160e01b815260040160405180910390fd5b6127138686868661366f565b612668565b6126688686868686613733565b348160005b81811015612798573685858381811061274557612745614ead565b6040029190910191505080358481111561277257604051631a783b8d60e01b815260040160405180910390fd5b61278b6127856040840160208501614adf565b8261343f565b909303925060010161272a565b50818611156127ba57604051631a783b8d60e01b815260040160405180910390fd5b6127c4858761343f565b858211156127d8576127d83387840361343f565b505050505050565b6127ea818361381a565b8161281c57826001146128105760405163efcc00b160e01b815260040160405180910390fd5b6127138787878761366f565b612668828260028a8a8a8a8a613839565b612836836138b9565b612840818361381a565b81612852576127138787878787613733565b612668828260038a8a8a8a8a613839565b6000806000806000861561289a5788945033935061288760c0890160a08a01614adf565b9250505060e086013560c08701356128bc565b3394508893506128ad6020890189614adf565b92505050604086013560208701355b80156128db57604051636ab37ce760e01b815260040160405180910390fd5b50602086026101e4033560006128f56102008a018a614f4c565b9050905060005b8181101561296957366129136102008c018c614f4c565b8381811061292357612923614ead565b6040029190910191505080358a156129425761293f81876150b6565b95505b61295f878a6129576040860160208701614adf565b84898f6138da565b50506001016128fc565b5061297884878786868c6138da565b50505050505050505050565b60408151146129905750565b600061299d826020015190565b90506129a98183613915565b5050565b60008284036129bd57508061052c565b82848309156129d75763c63cf08960e01b60005260046000fd5b60006129e38584614fc3565b93909304949350505050565b6000848614612a3a57838303428590038082036000612a0e838a614fc3565b612a18838c614fc3565b612a229190614fab565b90508584878303040181151502945050505050610330565b5092949350505050565b612a4c613edb565b831580612a57575081155b15612a7557604051634c74edb760e11b815260040160405180910390fd5b612a7d613edb565b612ada878585808060200260200160405190810160405280939291908181526020016000905b82821015612acf57612ac0604083028601368190038101906150cd565b81526020019060010190612aa3565b505050505083613939565b805160408051602080890282018101909252878152612b39918a91908a908a90819060009085015b82821015612b2e57612b1f604083028601368190038101906150cd565b81526020019060010190612b02565b505050505085613adc565b80516005811115612b4c57612b4c61465e565b8351516005811115612b6057612b6061465e565b141580612b8b575080602001516001600160a01b03168360000151602001516001600160a01b031614155b80612ba25750806040015183600001516040015114155b15612bc0576040516309cfb45560e01b815260040160405180910390fd5b82600001516060015181606001511115612c6e57600085856000818110612be957612be9614ead565b905060400201803603810190612bff91906150cd565b905083600001516060015182606001510389826000015181518110612c2657612c26614ead565b60200260200101516000015160600151826020015181518110612c4b57612c4b614ead565b602090810291909101015160609081019190915284518101519083015250612d09565b600087876000818110612c8357612c83614ead565b905060400201803603810190612c9991906150cd565b905081606001518460000151606001510389826000015181518110612cc057612cc0614ead565b60200260200101516000015160400151826020015181518110612ce557612ce5614ead565b60200260200101516060018181525050816060015184600001516060018181525050505b60809081015183516001600160a01b039091169101525095945050505050565b8151606090806001600160401b03811115612d4657612d46614014565b604051908082528060200260200182016040528015612d6f578160200160208202803683370190505b50915060005b81811015612e57576000858281518110612d9157612d91614ead565b6020026020010151905080602001516001600160781b0316600003612db65750612e4f565b6001848381518110612dca57612dca614ead565b91151560209283029190910190910152805160600151805160005b81811015612e4a576000838281518110612e0157612e01614ead565b602002602001015160600151905080600014612e41576040516314bea84160e31b8152600481018790526024810183905260448101829052606401610d9a565b50600101612de5565b505050505b600101612d75565b5060408051602080825281830190925234916000919060208201818036833701905050855190915060005b81811015612f15576000878281518110612e9e57612e9e614ead565b60209081029190910101518051909150600081516005811115612ec357612ec361465e565b03612ef7578581606001511115612eed57604051631a783b8d60e01b815260040160405180910390fd5b8060600151860395505b612f0b8183602001518460400151886132e3565b5050600101612e82565b50612f1f82612984565b8215612f2f57612f2f338461343f565b612f396001600055565b5050505092915050565b612f4b613edb565b8351600003612f6f578460405163375c24c160e01b8152600401610d9a91906150f9565b6000856001811115612f8357612f8361465e565b03612faa5780516001600160a01b038316608090910152612fa5868583613adc565b612fc3565b612fb5868583613939565b336020820152604081018390525b8051606001516000036103305760006020820181905281516080015295945050505050565b808210156129a957604051632335530b60e11b815260040160405180910390fd5b6000806000526000825160208403805182604103600060018211613077576040880151606089015160001a9650821561305557601b8160ff1c0196506001600160ff1b03811660408a01525b8689528985526020600060808760015afa508385528589526040890152506000515b89148915151695508590506131385760408252604486038051604088038051630b135d3f60e11b84528a82526020600060648901868f5afa9850881561312e57630b135d3f60e11b6000511461312e578b3b156130df57634f7fb80d60e01b60005260046000fd5b60018760410311156130fc57638baa579f60e01b60005260046000fd5b640101000000881a61311d57630f801e8560e11b6000528760045260246000fd5b632057875960e21b60005260046000fd5b8486529190925290525b50505050806119eb57613149613c6d565b634f7fb80d60e01b60005260046000fd5b60004284118061316a5750428311155b1561319657811561318e576040516337bf561360e11b815260040160405180910390fd5b50600061052c565b5060019392505050565b60018360038111156131b4576131b461465e565b1180156131ca5750336001600160a01b03821614155b80156131df5750336001600160a01b03831614155b15613276576080880151511580156131f657508651155b1561320c5761320781868487613cb5565b613276565b600061326a82633313157060e01b88338d8c8e6040516024016132339594939291906152b9565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613d01565b90506115618187613d16565b5050505050505050565b600083600052602060002060208301835160051b81015b808210156132bf57815180841160051b93845260209384185260406000209290910190613297565b505083149050806119eb576040516309bde33960e01b815260040160405180910390fd5b6000845160058111156132f8576132f861465e565b0361334657604084015160208501516001600160a01b0316171561332f57604051636ab37ce760e01b815260040160405180910390fd5b6133418460800151856060015161343f565b6119eb565b60018451600581111561335b5761335b61465e565b0361339d5760408401511561338357604051636ab37ce760e01b815260040160405180910390fd5b6133418460200151848660800151876060015186866138da565b6002845160058111156133b2576133b261465e565b036133d65761334184602001518486608001518760400151886060015187876127e0565b6119eb846020015184866080015187604001518860600151878761282d565b6000868803613410576134098686896129ad565b9050613434565b61343161341e87878b6129ad565b61342988888b6129ad565b8686866129ef565b90505b979650505050505050565b613448816138b9565b600080600080600085875af190508061348e57613463613c6d565b60405163470c7c1d60e01b81526001600160a01b038416600482015260248101839052604401610d9a565b505050565b60186101243510610244356102606102643560400201146004356020146102243561024014161616806134d9576040516339f3e3fd60e01b815260040160405180910390fd5b50565b60018360038111156134f0576134f061465e565b1180156135065750336001600160a01b03821614155b801561351b5750336001600160a01b03831614155b1561200b5761200b81868487613cb5565b600083815260026020526040902061354784826001806118d7565b50805460ff1661355c5761355c83858461199c565b710100000000000000000000000000000100019055505050565b604080517f000000000000000000000000000000000000000000000000000000000000000060ff60a01b17600090815260208690527f000000000000000000000000000000000000000000000000000000000000000083526055600b20919092526001600160a01b03169050600080600080526020600085876000875af1915060005190508161362c57613608613c6d565b60405163344f54f560e21b81526001600160a01b0384166004820152602401610d9a565b6001600160e01b03198116632671a55160e11b146127d857604051630e7ccd9360e11b8152600481018790526001600160a01b0384166024820152604401610d9a565b833b61368a57632f8aeb3960e11b6000528360045260246000fd5b6040516323b872dd60e01b6000528360045282602452816044526000806064600080895af180613724573d156136fe576020601f3d01046020830481600302818311156136e557818303600302610200838002858002030401015b5a6020820110156136fa573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005285600452846024528360445282606452600160845260a46000fd5b50604052505060006060525050565b843b61374e57632f8aeb3960e11b6000528460045260246000fd5b60405160805160a05160c051637921219560e11b6000528760045286602452856044528460645260a0608452600060a45260008060c46000808d5af1806137fe573d156137d9576020601f3d01046020860481600302818311156137c057818303600302610200838002858002030401015b5a6020820110156137d5573d6000803e3d6000fd5b5050505b63f486bc8760e01b600052896004528860245287604452866064528560845260a46000fd5b5060809290925260a05260c05260405250506000606052505050565b6000613827836020015190565b905081811461348e5761348e83612984565b600060208851036138745750604080885260208089018a9052632671a55160e11b918901919091526044880152600160648801819052613883565b50606487018051600101908190525b603c60c082028901038781528660208201528560408201528460608201528360808201528260a082015250505050505050505050565b806000036134d95760405163246cf94560e21b815260040160405180910390fd5b6138e3836138b9565b6138ed818361381a565b81613903576138fe86868686613d70565b6127d8565b6127d88282600189898960008a613839565b6064810151604082019060c002604401613930848383613576565b50506020905250565b613965565b637fda727960e01b60005260046000fd5b634e487b7160e01b600052601160045260246000fd5b602082018051518451811061397c5761397c61393e565b60208102602086010151606081510151602084510151815181106139a2576139a261393e565b602081026020830101516000806020860151156139c9575050606081018051600090915280155b885183518152602084015160208201526040840151604082015260a084015160808201526060812060208c51028c015b808b1015613aa05760208b019a508a515199508d518a10613a1c57613a1c61393e565b60208a0260208f01015198506020890151156139f957606089510151975060208b510151965087518710613a5257613a5261393e565b602087026020890101519550606086018051860181511587821060011b1786179550809650506000815250606086208214608084015160a08801511416613a9b57613a9b61393e565b6139f9565b50506060018290528015613acf5760018103613ac75763246cf94560e21b60005260046000fd5b613acf61394f565b5050505050505050505050565b6020820180515184518110613af357613af361393e565b602081026020860101518051604081015160208551015181518110613b1a57613b1a61393e565b60208102602083010151600080602087015115613b41575050606081018051600090915280155b8951835181526020840151602082015260408401516040820152865160208c015261012087015160408c015260608120905060208c51028c015b808b1015613c2f5760208b019a508a515199508d518a10613b9e57613b9e61393e565b60208a0260208f0101519850602089015115613b7b57885197506040880151965060208b510151955086518610613bd757613bd761393e565b602086026020880101519450606085018051850181511586821060011b178517945080955050600081525060608520821460408d01516101208a01511460208e01518a51141616613c2a57613c2a61393e565b613b7b565b50508160608b5101528015613c5f5760018103613c575763246cf94560e21b60005260046000fd5b613c5f61394f565b505050505050505050505050565b3d156116ba576020601f3d01046020604051048160030281831115613ca057818303600302610200838002858002030401015b5a60208201101561348e573d6000803e3d6000fd5b604051602481018490523360448201526001600160a01b038316606482015260848101829052600090613cf59086906303874c7760e21b9060a401613233565b905061200b8185613d16565b6000806000835160208501865afa9392505050565b81613d3f57613d23613c6d565b604051633ed4053f60e21b815260048101829052602401610d9a565b613d4f6303874c7760e21b613e79565b156129a957604051633ed4053f60e21b815260048101829052602401610d9a565b6040516323b872dd60e01b600052836004528260245281604452602060006064600080895af1803d15601f3d116001600051141617163d15158116613e695780873b151516613e695780613e545781613e33573d15613e0d576020601f3d0104602084048160030281831115613df457818303600302610200838002858002030401015b5a602082011015613e09573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005286600452856024528460445260006064528360845260a46000fd5b639889192360e01b6000528660045285602452846044528360645260846000fd5b632f8aeb3960e11b6000528660045260246000fd5b5050604052505060006060525050565b60008060203d03613e8f5760206000803e506000515b6001600160e01b031990811692169190911415919050565b6040518060a00160405280613eba613f1e565b81526000602082018190526040820152606080820181905260809091015290565b60408051610100810182526000606082018181526080830182905260a0830182905260c0830182905260e083018290528252602082018190529181019190915290565b60405180610160016040528060006001600160a01b0316815260200160006001600160a01b03168152602001606081526020016060815260200160006003811115613f6b57613f6b61465e565b815260006020820181905260408201819052606082018190526080820181905260a0820181905260c09091015290565b6000815180845260005b81811015613fc157602081850181015186830182015201613fa5565b81811115613fd3576000602083870101525b50601f01601f19169290920160200192915050565b60208152600061052c6020830184613f9b565b60006020828403121561400d57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b038111828210171561404c5761404c614014565b60405290565b60405161016081016001600160401b038111828210171561404c5761404c614014565b604051601f8201601f191681016001600160401b038111828210171561409d5761409d614014565b604052919050565b60006001600160401b038211156140be576140be614014565b5060051b60200190565b6001600160a01b03811681146134d957600080fd5b80356140e8816140c8565b919050565b8035600681106140e857600080fd5b600060a0828403121561410e57600080fd5b61411661402a565b9050614121826140ed565b81526020820135614131816140c8565b8060208301525060408201356040820152606082013560608201526080820135608082015292915050565b600082601f83011261416d57600080fd5b8135602061418261417d836140a5565b614075565b82815260a092830285018201928282019190878511156141a157600080fd5b8387015b858110156141c4576141b789826140fc565b84529284019281016141a5565b5090979650505050505050565b600060c082840312156141e357600080fd5b60405160c081018181106001600160401b038211171561420557614205614014565b604052905080614214836140ed565b81526020830135614224816140c8565b8060208301525060408301356040820152606083013560608201526080830135608082015260a0830135614257816140c8565b60a0919091015292915050565b600082601f83011261427557600080fd5b8135602061428561417d836140a5565b82815260c092830285018201928282019190878511156142a457600080fd5b8387015b858110156141c4576142ba89826141d1565b84529284019281016142a8565b8035600481106140e857600080fd5b600061016082840312156142e957600080fd5b6142f1614052565b90506142fc826140dd565b815261430a602083016140dd565b602082015260408201356001600160401b038082111561432957600080fd5b6143358583860161415c565b6040840152606084013591508082111561434e57600080fd5b5061435b84828501614264565b60608301525061436d608083016142c7565b608082015260a082013560a082015260c082013560c082015260e082013560e082015261010080830135818301525061012080830135818301525061014080830135818301525092915050565b80356001600160781b03811681146140e857600080fd5b600082601f8301126143e257600080fd5b81356001600160401b038111156143fb576143fb614014565b61440e601f8201601f1916602001614075565b81815284602083860101111561442357600080fd5b816020850160208301376000918101602001919091529392505050565b600060a0828403121561445257600080fd5b61445a61402a565b905081356001600160401b038082111561447357600080fd5b61447f858386016142d6565b835261448d602085016143ba565b602084015261449e604085016143ba565b604084015260608401359150808211156144b757600080fd5b6144c3858386016143d1565b606084015260808401359150808211156144dc57600080fd5b506144e9848285016143d1565b60808301525092915050565b600082601f83011261450657600080fd5b8135602061451661417d836140a5565b82815260059290921b8401810191818101908684111561453557600080fd5b8286015b848110156145745780356001600160401b038111156145585760008081fd5b6145668986838b0101614440565b845250918301918301614539565b509695505050505050565b60008083601f84011261459157600080fd5b5081356001600160401b038111156145a857600080fd5b6020830191508360208260051b85010111156145c357600080fd5b9250929050565b6000806000806000606086880312156145e257600080fd5b85356001600160401b03808211156145f957600080fd5b61460589838a016144f5565b9650602088013591508082111561461b57600080fd5b61462789838a0161457f565b9096509450604088013591508082111561464057600080fd5b5061464d8882890161457f565b969995985093965092949392505050565b634e487b7160e01b600052602160045260246000fd5b600681106146845761468461465e565b9052565b614693828251614674565b6020818101516001600160a01b0390811691840191909152604080830151908401526060808301519084015260809182015116910152565b600081518084526020808501945080840160005b838110156147215781516146f4888251614688565b808401516001600160a01b031660a08901526040015160c088015260e090960195908201906001016146df565b509495945050505050565b60208152600061052c60208301846146cb565b60006020828403121561475157600080fd5b81356001600160401b0381111561476757600080fd5b8201610160818503121561052c57600080fd5b60008060008060008060008060008060e08b8d03121561479957600080fd5b8a356001600160401b03808211156147b057600080fd5b6147bc8e838f016144f5565b9b5060208d01359150808211156147d257600080fd5b6147de8e838f0161457f565b909b50995060408d01359150808211156147f757600080fd5b6148038e838f0161457f565b909950975060608d013591508082111561481c57600080fd5b506148298d828e0161457f565b90965094505060808b0135925061484260a08c016140dd565b915060c08b013590509295989b9194979a5092959850565b604080825283519082018190526000906020906060840190828701845b82811015614895578151151584529284019290840190600101614877565b505050838103828501526148a981866146cb565b9695505050505050565b600080602083850312156148c657600080fd5b82356001600160401b038111156148dc57600080fd5b6148e88582860161457f565b90969095509350505050565b6000806000806040858703121561490a57600080fd5b84356001600160401b038082111561492157600080fd5b61492d8883890161457f565b9096509450602087013591508082111561494657600080fd5b506149538782880161457f565b95989497509550505050565b6000806040838503121561497257600080fd5b82356001600160401b0381111561498857600080fd5b83016040818603121561499a57600080fd5b946020939093013593505050565b6000806000806000608086880312156149c057600080fd5b85356001600160401b03808211156149d757600080fd5b9087019060a0828a0312156149eb57600080fd5b90955060208701359080821115614a0157600080fd5b50614a0e8882890161457f565b909550935050604086013591506060860135614a29816140c8565b809150509295509295909350565b60008060008060008060008060a0898b031215614a5357600080fd5b88356001600160401b0380821115614a6a57600080fd5b614a768c838d0161457f565b909a50985060208b0135915080821115614a8f57600080fd5b614a9b8c838d0161457f565b909850965060408b0135915080821115614ab457600080fd5b50614ac18b828c0161457f565b999c989b509699959896976060870135966080013595509350505050565b600060208284031215614af157600080fd5b813561052c816140c8565b606081526000614b0f6060830186613f9b565b6020830194909452506001600160a01b0391909116604090910152919050565b600060208284031215614b4157600080fd5b81356001600160401b03811115614b5757600080fd5b8201610240818503121561052c57600080fd5b6000614b7861417d846140a5565b83815260208082019190600586811b860136811115614b9657600080fd5b865b81811015614c915780356001600160401b0380821115614bb85760008081fd5b818a01915060a08236031215614bce5760008081fd5b614bd661402a565b823581528683013560028110614bec5760008081fd5b81880152604083810135908201526060808401359082015260808084013583811115614c185760008081fd5b939093019236601f850112614c2f57600092508283fd5b83359250614c3f61417d846140a5565b83815292871b84018801928881019036851115614c5c5760008081fd5b948901945b84861015614c7a57853582529489019490890190614c61565b918301919091525088525050948301948301614b98565b5092979650505050505050565b6000808335601e19843603018112614cb557600080fd5b8301803591506001600160401b03821115614ccf57600080fd5b602001915060a0810236038213156145c357600080fd5b600060a08284031215614cf857600080fd5b61052c83836140fc565b6000808335601e19843603018112614d1957600080fd5b8301803591506001600160401b03821115614d3357600080fd5b602001915060c0810236038213156145c357600080fd5b600060c08284031215614d5c57600080fd5b61052c83836141d1565b600060208284031215614d7857600080fd5b61052c826142c7565b60006104d43683614440565b600060408284031215614d9f57600080fd5b604051604081018181106001600160401b0382111715614dc157614dc1614014565b604052823581526020928301359281019290925250919050565b6000614de961417d846140a5565b80848252602080830192508560051b850136811115614e0757600080fd5b855b81811015614ea15780356001600160401b03811115614e285760008081fd5b870136601f820112614e3a5760008081fd5b8035614e4861417d826140a5565b81815260069190911b82018501908581019036831115614e685760008081fd5b928601925b82841015614e9157614e7f3685614d8d565b82528682019150604084019350614e6d565b8852505050938201938201614e09565b50919695505050505050565b634e487b7160e01b600052603260045260246000fd5b60008235603e19833603018112614ed957600080fd5b9190910192915050565b6000823561015e19833603018112614ed957600080fd5b60006104d436836142d6565b6000808335601e19843603018112614f1d57600080fd5b8301803591506001600160401b03821115614f3757600080fd5b6020019150368190038213156145c357600080fd5b6000808335601e19843603018112614f6357600080fd5b8301803591506001600160401b03821115614f7d57600080fd5b6020019150600681901b36038213156145c357600080fd5b634e487b7160e01b600052601160045260246000fd5b60008219821115614fbe57614fbe614f95565b500190565b6000816000190483118215151615614fdd57614fdd614f95565b500290565b600081518084526020808501945080840160005b8381101561472157615009878351614688565b60a0969096019590820190600101614ff6565b60006080808301878452602060018060a01b03808916828701526040848188015283895180865260a089019150848b01955060005b81811015615092578651615066848251614674565b808701518616848801528481015185850152606090810151908401529585019591870191600101615051565b505087810360608901526150a6818a614fe2565b9c9b505050505050505050505050565b6000828210156150c8576150c8614f95565b500390565b6000604082840312156150df57600080fd5b61052c8383614d8d565b600281106146845761468461465e565b602081016104d482846150e9565b600081518084526020808501945080840160005b83811015614721578151615130888251614674565b838101516001600160a01b03168885015260408082015190890152606080820151908901526080908101519088015260a0909601959082019060010161511b565b600081518084526020808501945080840160005b8381101561472157815161519a888251614674565b808401516001600160a01b0390811689860152604080830151908a0152606080830151908a0152608080830151908a015260a091820151169088015260c09096019590820190600101615185565b600481106146845761468461465e565b600081518084526020808501945080840160005b838110156147215781518752958201959082019060010161520c565b600081518084526020808501808196508360051b8101915082860160005b858110156152ac578284038952815160a0815186528682015161526b888801826150e9565b506040828101519087015260608083015190870152608091820151918601819052615298818701836151f8565b9a87019a9550505090840190600101615246565b5091979650505050505050565b85815260018060a01b038516602082015260a060408201526000610140855160a0808501526152f382850182516001600160a01b03169052565b602081015161016061530f818701836001600160a01b03169052565b60408301519150806101808701525061532c6102a0860182615107565b9050606082015161013f19868303016101a087015261534b8282615171565b91505060808201516153616101c08701826151e8565b5060a08201516101e086015260c082015161020086015260e082015161022086015261010080830151610240870152610120808401516102608801528484015161028088015260208a015194506153c360c08801866001600160781b03169052565b60408a01516001600160781b031660e088015260608a0151878403609f19908101848a015290955093506153f78386613f9b565b945060808a0151925083878603018188015250506154158382613f9b565b92505050828103606084015261542b81866151f8565b9050828103608084015261543f8185615228565b9897505050505050505056fea26469706673582212202a524af1b63ccd051ee58650b72f8274d1a5cf7d5fa9df66cecb9f412b7f55af64736f6c634300080d0033",
}

// ConsiderationABI is the input ABI used to generate the binding from.
// Deprecated: Use ConsiderationMetaData.ABI instead.
var ConsiderationABI = ConsiderationMetaData.ABI

// ConsiderationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConsiderationMetaData.Bin instead.
var ConsiderationBin = ConsiderationMetaData.Bin

// DeployConsideration deploys a new Ethereum contract, binding an instance of Consideration to it.
func DeployConsideration(auth *bind.TransactOpts, backend bind.ContractBackend, conduitController common.Address) (common.Address, *types.Transaction, *Consideration, error) {
	parsed, err := ConsiderationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConsiderationBin), backend, conduitController)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Consideration{ConsiderationCaller: ConsiderationCaller{contract: contract}, ConsiderationTransactor: ConsiderationTransactor{contract: contract}, ConsiderationFilterer: ConsiderationFilterer{contract: contract}}, nil
}

// Consideration is an auto generated Go binding around an Ethereum contract.
type Consideration struct {
	ConsiderationCaller     // Read-only binding to the contract
	ConsiderationTransactor // Write-only binding to the contract
	ConsiderationFilterer   // Log filterer for contract events
}

// ConsiderationCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConsiderationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsiderationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConsiderationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsiderationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConsiderationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsiderationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConsiderationSession struct {
	Contract     *Consideration    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConsiderationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConsiderationCallerSession struct {
	Contract *ConsiderationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ConsiderationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConsiderationTransactorSession struct {
	Contract     *ConsiderationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ConsiderationRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConsiderationRaw struct {
	Contract *Consideration // Generic contract binding to access the raw methods on
}

// ConsiderationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConsiderationCallerRaw struct {
	Contract *ConsiderationCaller // Generic read-only contract binding to access the raw methods on
}

// ConsiderationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConsiderationTransactorRaw struct {
	Contract *ConsiderationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConsideration creates a new instance of Consideration, bound to a specific deployed contract.
func NewConsideration(address common.Address, backend bind.ContractBackend) (*Consideration, error) {
	contract, err := bindConsideration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Consideration{ConsiderationCaller: ConsiderationCaller{contract: contract}, ConsiderationTransactor: ConsiderationTransactor{contract: contract}, ConsiderationFilterer: ConsiderationFilterer{contract: contract}}, nil
}

// NewConsiderationCaller creates a new read-only instance of Consideration, bound to a specific deployed contract.
func NewConsiderationCaller(address common.Address, caller bind.ContractCaller) (*ConsiderationCaller, error) {
	contract, err := bindConsideration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConsiderationCaller{contract: contract}, nil
}

// NewConsiderationTransactor creates a new write-only instance of Consideration, bound to a specific deployed contract.
func NewConsiderationTransactor(address common.Address, transactor bind.ContractTransactor) (*ConsiderationTransactor, error) {
	contract, err := bindConsideration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConsiderationTransactor{contract: contract}, nil
}

// NewConsiderationFilterer creates a new log filterer instance of Consideration, bound to a specific deployed contract.
func NewConsiderationFilterer(address common.Address, filterer bind.ContractFilterer) (*ConsiderationFilterer, error) {
	contract, err := bindConsideration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConsiderationFilterer{contract: contract}, nil
}

// bindConsideration binds a generic wrapper to an already deployed contract.
func bindConsideration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConsiderationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Consideration *ConsiderationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Consideration.Contract.ConsiderationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Consideration *ConsiderationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Consideration.Contract.ConsiderationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Consideration *ConsiderationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Consideration.Contract.ConsiderationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Consideration *ConsiderationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Consideration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Consideration *ConsiderationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Consideration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Consideration *ConsiderationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Consideration.Contract.contract.Transact(opts, method, params...)
}

// GetCounter is a free data retrieval call binding the contract method 0xf07ec373.
//
// Solidity: function getCounter(address offerer) view returns(uint256 counter)
func (_Consideration *ConsiderationCaller) GetCounter(opts *bind.CallOpts, offerer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Consideration.contract.Call(opts, &out, "getCounter", offerer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCounter is a free data retrieval call binding the contract method 0xf07ec373.
//
// Solidity: function getCounter(address offerer) view returns(uint256 counter)
func (_Consideration *ConsiderationSession) GetCounter(offerer common.Address) (*big.Int, error) {
	return _Consideration.Contract.GetCounter(&_Consideration.CallOpts, offerer)
}

// GetCounter is a free data retrieval call binding the contract method 0xf07ec373.
//
// Solidity: function getCounter(address offerer) view returns(uint256 counter)
func (_Consideration *ConsiderationCallerSession) GetCounter(offerer common.Address) (*big.Int, error) {
	return _Consideration.Contract.GetCounter(&_Consideration.CallOpts, offerer)
}

// GetOrderHash is a free data retrieval call binding the contract method 0x79df72bd.
//
// Solidity: function getOrderHash((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) order) view returns(bytes32 orderHash)
func (_Consideration *ConsiderationCaller) GetOrderHash(opts *bind.CallOpts, order OrderComponents) ([32]byte, error) {
	var out []interface{}
	err := _Consideration.contract.Call(opts, &out, "getOrderHash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOrderHash is a free data retrieval call binding the contract method 0x79df72bd.
//
// Solidity: function getOrderHash((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) order) view returns(bytes32 orderHash)
func (_Consideration *ConsiderationSession) GetOrderHash(order OrderComponents) ([32]byte, error) {
	return _Consideration.Contract.GetOrderHash(&_Consideration.CallOpts, order)
}

// GetOrderHash is a free data retrieval call binding the contract method 0x79df72bd.
//
// Solidity: function getOrderHash((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) order) view returns(bytes32 orderHash)
func (_Consideration *ConsiderationCallerSession) GetOrderHash(order OrderComponents) ([32]byte, error) {
	return _Consideration.Contract.GetOrderHash(&_Consideration.CallOpts, order)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns(bool isValidated, bool isCancelled, uint256 totalFilled, uint256 totalSize)
func (_Consideration *ConsiderationCaller) GetOrderStatus(opts *bind.CallOpts, orderHash [32]byte) (struct {
	IsValidated bool
	IsCancelled bool
	TotalFilled *big.Int
	TotalSize   *big.Int
}, error) {
	var out []interface{}
	err := _Consideration.contract.Call(opts, &out, "getOrderStatus", orderHash)

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
func (_Consideration *ConsiderationSession) GetOrderStatus(orderHash [32]byte) (struct {
	IsValidated bool
	IsCancelled bool
	TotalFilled *big.Int
	TotalSize   *big.Int
}, error) {
	return _Consideration.Contract.GetOrderStatus(&_Consideration.CallOpts, orderHash)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns(bool isValidated, bool isCancelled, uint256 totalFilled, uint256 totalSize)
func (_Consideration *ConsiderationCallerSession) GetOrderStatus(orderHash [32]byte) (struct {
	IsValidated bool
	IsCancelled bool
	TotalFilled *big.Int
	TotalSize   *big.Int
}, error) {
	return _Consideration.Contract.GetOrderStatus(&_Consideration.CallOpts, orderHash)
}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() view returns(string version, bytes32 domainSeparator, address conduitController)
func (_Consideration *ConsiderationCaller) Information(opts *bind.CallOpts) (struct {
	Version           string
	DomainSeparator   [32]byte
	ConduitController common.Address
}, error) {
	var out []interface{}
	err := _Consideration.contract.Call(opts, &out, "information")

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
func (_Consideration *ConsiderationSession) Information() (struct {
	Version           string
	DomainSeparator   [32]byte
	ConduitController common.Address
}, error) {
	return _Consideration.Contract.Information(&_Consideration.CallOpts)
}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() view returns(string version, bytes32 domainSeparator, address conduitController)
func (_Consideration *ConsiderationCallerSession) Information() (struct {
	Version           string
	DomainSeparator   [32]byte
	ConduitController common.Address
}, error) {
	return _Consideration.Contract.Information(&_Consideration.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string contractName)
func (_Consideration *ConsiderationCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Consideration.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string contractName)
func (_Consideration *ConsiderationSession) Name() (string, error) {
	return _Consideration.Contract.Name(&_Consideration.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string contractName)
func (_Consideration *ConsiderationCallerSession) Name() (string, error) {
	return _Consideration.Contract.Name(&_Consideration.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0xfd9f1e10.
//
// Solidity: function cancel((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256)[] orders) returns(bool cancelled)
func (_Consideration *ConsiderationTransactor) Cancel(opts *bind.TransactOpts, orders []OrderComponents) (*types.Transaction, error) {
	return _Consideration.contract.Transact(opts, "cancel", orders)
}

// Cancel is a paid mutator transaction binding the contract method 0xfd9f1e10.
//
// Solidity: function cancel((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256)[] orders) returns(bool cancelled)
func (_Consideration *ConsiderationSession) Cancel(orders []OrderComponents) (*types.Transaction, error) {
	return _Consideration.Contract.Cancel(&_Consideration.TransactOpts, orders)
}

// Cancel is a paid mutator transaction binding the contract method 0xfd9f1e10.
//
// Solidity: function cancel((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256)[] orders) returns(bool cancelled)
func (_Consideration *ConsiderationTransactorSession) Cancel(orders []OrderComponents) (*types.Transaction, error) {
	return _Consideration.Contract.Cancel(&_Consideration.TransactOpts, orders)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xe7acab24.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) advancedOrder, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, bytes32 fulfillerConduitKey, address recipient) payable returns(bool fulfilled)
func (_Consideration *ConsiderationTransactor) FulfillAdvancedOrder(opts *bind.TransactOpts, advancedOrder AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillerConduitKey [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Consideration.contract.Transact(opts, "fulfillAdvancedOrder", advancedOrder, criteriaResolvers, fulfillerConduitKey, recipient)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xe7acab24.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) advancedOrder, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, bytes32 fulfillerConduitKey, address recipient) payable returns(bool fulfilled)
func (_Consideration *ConsiderationSession) FulfillAdvancedOrder(advancedOrder AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillerConduitKey [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Consideration.Contract.FulfillAdvancedOrder(&_Consideration.TransactOpts, advancedOrder, criteriaResolvers, fulfillerConduitKey, recipient)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xe7acab24.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) advancedOrder, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, bytes32 fulfillerConduitKey, address recipient) payable returns(bool fulfilled)
func (_Consideration *ConsiderationTransactorSession) FulfillAdvancedOrder(advancedOrder AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillerConduitKey [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Consideration.Contract.FulfillAdvancedOrder(&_Consideration.TransactOpts, advancedOrder, criteriaResolvers, fulfillerConduitKey, recipient)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0x87201b41.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, address recipient, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Consideration *ConsiderationTransactor) FulfillAvailableAdvancedOrders(opts *bind.TransactOpts, advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, recipient common.Address, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Consideration.contract.Transact(opts, "fulfillAvailableAdvancedOrders", advancedOrders, criteriaResolvers, offerFulfillments, considerationFulfillments, fulfillerConduitKey, recipient, maximumFulfilled)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0x87201b41.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, address recipient, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Consideration *ConsiderationSession) FulfillAvailableAdvancedOrders(advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, recipient common.Address, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Consideration.Contract.FulfillAvailableAdvancedOrders(&_Consideration.TransactOpts, advancedOrders, criteriaResolvers, offerFulfillments, considerationFulfillments, fulfillerConduitKey, recipient, maximumFulfilled)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0x87201b41.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, address recipient, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Consideration *ConsiderationTransactorSession) FulfillAvailableAdvancedOrders(advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, recipient common.Address, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Consideration.Contract.FulfillAvailableAdvancedOrders(&_Consideration.TransactOpts, advancedOrders, criteriaResolvers, offerFulfillments, considerationFulfillments, fulfillerConduitKey, recipient, maximumFulfilled)
}

// FulfillAvailableOrders is a paid mutator transaction binding the contract method 0xed98a574.
//
// Solidity: function fulfillAvailableOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Consideration *ConsiderationTransactor) FulfillAvailableOrders(opts *bind.TransactOpts, orders []Order, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Consideration.contract.Transact(opts, "fulfillAvailableOrders", orders, offerFulfillments, considerationFulfillments, fulfillerConduitKey, maximumFulfilled)
}

// FulfillAvailableOrders is a paid mutator transaction binding the contract method 0xed98a574.
//
// Solidity: function fulfillAvailableOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Consideration *ConsiderationSession) FulfillAvailableOrders(orders []Order, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Consideration.Contract.FulfillAvailableOrders(&_Consideration.TransactOpts, orders, offerFulfillments, considerationFulfillments, fulfillerConduitKey, maximumFulfilled)
}

// FulfillAvailableOrders is a paid mutator transaction binding the contract method 0xed98a574.
//
// Solidity: function fulfillAvailableOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Consideration *ConsiderationTransactorSession) FulfillAvailableOrders(orders []Order, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Consideration.Contract.FulfillAvailableOrders(&_Consideration.TransactOpts, orders, offerFulfillments, considerationFulfillments, fulfillerConduitKey, maximumFulfilled)
}

// FulfillBasicOrder is a paid mutator transaction binding the contract method 0xfb0f3ee1.
//
// Solidity: function fulfillBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Consideration *ConsiderationTransactor) FulfillBasicOrder(opts *bind.TransactOpts, parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Consideration.contract.Transact(opts, "fulfillBasicOrder", parameters)
}

// FulfillBasicOrder is a paid mutator transaction binding the contract method 0xfb0f3ee1.
//
// Solidity: function fulfillBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Consideration *ConsiderationSession) FulfillBasicOrder(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Consideration.Contract.FulfillBasicOrder(&_Consideration.TransactOpts, parameters)
}

// FulfillBasicOrder is a paid mutator transaction binding the contract method 0xfb0f3ee1.
//
// Solidity: function fulfillBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Consideration *ConsiderationTransactorSession) FulfillBasicOrder(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Consideration.Contract.FulfillBasicOrder(&_Consideration.TransactOpts, parameters)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xb3a34c4c.
//
// Solidity: function fulfillOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) order, bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Consideration *ConsiderationTransactor) FulfillOrder(opts *bind.TransactOpts, order Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Consideration.contract.Transact(opts, "fulfillOrder", order, fulfillerConduitKey)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xb3a34c4c.
//
// Solidity: function fulfillOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) order, bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Consideration *ConsiderationSession) FulfillOrder(order Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Consideration.Contract.FulfillOrder(&_Consideration.TransactOpts, order, fulfillerConduitKey)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xb3a34c4c.
//
// Solidity: function fulfillOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) order, bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Consideration *ConsiderationTransactorSession) FulfillOrder(order Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Consideration.Contract.FulfillOrder(&_Consideration.TransactOpts, order, fulfillerConduitKey)
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns(uint256 newCounter)
func (_Consideration *ConsiderationTransactor) IncrementCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Consideration.contract.Transact(opts, "incrementCounter")
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns(uint256 newCounter)
func (_Consideration *ConsiderationSession) IncrementCounter() (*types.Transaction, error) {
	return _Consideration.Contract.IncrementCounter(&_Consideration.TransactOpts)
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns(uint256 newCounter)
func (_Consideration *ConsiderationTransactorSession) IncrementCounter() (*types.Transaction, error) {
	return _Consideration.Contract.IncrementCounter(&_Consideration.TransactOpts)
}

// MatchAdvancedOrders is a paid mutator transaction binding the contract method 0x55944a42.
//
// Solidity: function matchAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, ((uint256,uint256)[],(uint256,uint256)[])[] fulfillments) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Consideration *ConsiderationTransactor) MatchAdvancedOrders(opts *bind.TransactOpts, advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillments []Fulfillment) (*types.Transaction, error) {
	return _Consideration.contract.Transact(opts, "matchAdvancedOrders", advancedOrders, criteriaResolvers, fulfillments)
}

// MatchAdvancedOrders is a paid mutator transaction binding the contract method 0x55944a42.
//
// Solidity: function matchAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, ((uint256,uint256)[],(uint256,uint256)[])[] fulfillments) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Consideration *ConsiderationSession) MatchAdvancedOrders(advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillments []Fulfillment) (*types.Transaction, error) {
	return _Consideration.Contract.MatchAdvancedOrders(&_Consideration.TransactOpts, advancedOrders, criteriaResolvers, fulfillments)
}

// MatchAdvancedOrders is a paid mutator transaction binding the contract method 0x55944a42.
//
// Solidity: function matchAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, ((uint256,uint256)[],(uint256,uint256)[])[] fulfillments) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Consideration *ConsiderationTransactorSession) MatchAdvancedOrders(advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillments []Fulfillment) (*types.Transaction, error) {
	return _Consideration.Contract.MatchAdvancedOrders(&_Consideration.TransactOpts, advancedOrders, criteriaResolvers, fulfillments)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa8174404.
//
// Solidity: function matchOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders, ((uint256,uint256)[],(uint256,uint256)[])[] fulfillments) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Consideration *ConsiderationTransactor) MatchOrders(opts *bind.TransactOpts, orders []Order, fulfillments []Fulfillment) (*types.Transaction, error) {
	return _Consideration.contract.Transact(opts, "matchOrders", orders, fulfillments)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa8174404.
//
// Solidity: function matchOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders, ((uint256,uint256)[],(uint256,uint256)[])[] fulfillments) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Consideration *ConsiderationSession) MatchOrders(orders []Order, fulfillments []Fulfillment) (*types.Transaction, error) {
	return _Consideration.Contract.MatchOrders(&_Consideration.TransactOpts, orders, fulfillments)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa8174404.
//
// Solidity: function matchOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders, ((uint256,uint256)[],(uint256,uint256)[])[] fulfillments) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Consideration *ConsiderationTransactorSession) MatchOrders(orders []Order, fulfillments []Fulfillment) (*types.Transaction, error) {
	return _Consideration.Contract.MatchOrders(&_Consideration.TransactOpts, orders, fulfillments)
}

// Validate is a paid mutator transaction binding the contract method 0x88147732.
//
// Solidity: function validate(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders) returns(bool validated)
func (_Consideration *ConsiderationTransactor) Validate(opts *bind.TransactOpts, orders []Order) (*types.Transaction, error) {
	return _Consideration.contract.Transact(opts, "validate", orders)
}

// Validate is a paid mutator transaction binding the contract method 0x88147732.
//
// Solidity: function validate(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders) returns(bool validated)
func (_Consideration *ConsiderationSession) Validate(orders []Order) (*types.Transaction, error) {
	return _Consideration.Contract.Validate(&_Consideration.TransactOpts, orders)
}

// Validate is a paid mutator transaction binding the contract method 0x88147732.
//
// Solidity: function validate(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders) returns(bool validated)
func (_Consideration *ConsiderationTransactorSession) Validate(orders []Order) (*types.Transaction, error) {
	return _Consideration.Contract.Validate(&_Consideration.TransactOpts, orders)
}

// ConsiderationCounterIncrementedIterator is returned from FilterCounterIncremented and is used to iterate over the raw logs and unpacked data for CounterIncremented events raised by the Consideration contract.
type ConsiderationCounterIncrementedIterator struct {
	Event *ConsiderationCounterIncremented // Event containing the contract specifics and raw log

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
func (it *ConsiderationCounterIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsiderationCounterIncremented)
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
		it.Event = new(ConsiderationCounterIncremented)
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
func (it *ConsiderationCounterIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsiderationCounterIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsiderationCounterIncremented represents a CounterIncremented event raised by the Consideration contract.
type ConsiderationCounterIncremented struct {
	NewCounter *big.Int
	Offerer    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCounterIncremented is a free log retrieval operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_Consideration *ConsiderationFilterer) FilterCounterIncremented(opts *bind.FilterOpts, offerer []common.Address) (*ConsiderationCounterIncrementedIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Consideration.contract.FilterLogs(opts, "CounterIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return &ConsiderationCounterIncrementedIterator{contract: _Consideration.contract, event: "CounterIncremented", logs: logs, sub: sub}, nil
}

// WatchCounterIncremented is a free log subscription operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_Consideration *ConsiderationFilterer) WatchCounterIncremented(opts *bind.WatchOpts, sink chan<- *ConsiderationCounterIncremented, offerer []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Consideration.contract.WatchLogs(opts, "CounterIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsiderationCounterIncremented)
				if err := _Consideration.contract.UnpackLog(event, "CounterIncremented", log); err != nil {
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
func (_Consideration *ConsiderationFilterer) ParseCounterIncremented(log types.Log) (*ConsiderationCounterIncremented, error) {
	event := new(ConsiderationCounterIncremented)
	if err := _Consideration.contract.UnpackLog(event, "CounterIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConsiderationOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Consideration contract.
type ConsiderationOrderCancelledIterator struct {
	Event *ConsiderationOrderCancelled // Event containing the contract specifics and raw log

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
func (it *ConsiderationOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsiderationOrderCancelled)
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
		it.Event = new(ConsiderationOrderCancelled)
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
func (it *ConsiderationOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsiderationOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsiderationOrderCancelled represents a OrderCancelled event raised by the Consideration contract.
type ConsiderationOrderCancelled struct {
	OrderHash [32]byte
	Offerer   common.Address
	Zone      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Consideration *ConsiderationFilterer) FilterOrderCancelled(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*ConsiderationOrderCancelledIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Consideration.contract.FilterLogs(opts, "OrderCancelled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &ConsiderationOrderCancelledIterator{contract: _Consideration.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Consideration *ConsiderationFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *ConsiderationOrderCancelled, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Consideration.contract.WatchLogs(opts, "OrderCancelled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsiderationOrderCancelled)
				if err := _Consideration.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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
func (_Consideration *ConsiderationFilterer) ParseOrderCancelled(log types.Log) (*ConsiderationOrderCancelled, error) {
	event := new(ConsiderationOrderCancelled)
	if err := _Consideration.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConsiderationOrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the Consideration contract.
type ConsiderationOrderFulfilledIterator struct {
	Event *ConsiderationOrderFulfilled // Event containing the contract specifics and raw log

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
func (it *ConsiderationOrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsiderationOrderFulfilled)
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
		it.Event = new(ConsiderationOrderFulfilled)
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
func (it *ConsiderationOrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsiderationOrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsiderationOrderFulfilled represents a OrderFulfilled event raised by the Consideration contract.
type ConsiderationOrderFulfilled struct {
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
func (_Consideration *ConsiderationFilterer) FilterOrderFulfilled(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*ConsiderationOrderFulfilledIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Consideration.contract.FilterLogs(opts, "OrderFulfilled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &ConsiderationOrderFulfilledIterator{contract: _Consideration.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address recipient, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
func (_Consideration *ConsiderationFilterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *ConsiderationOrderFulfilled, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Consideration.contract.WatchLogs(opts, "OrderFulfilled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsiderationOrderFulfilled)
				if err := _Consideration.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
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
func (_Consideration *ConsiderationFilterer) ParseOrderFulfilled(log types.Log) (*ConsiderationOrderFulfilled, error) {
	event := new(ConsiderationOrderFulfilled)
	if err := _Consideration.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConsiderationOrderValidatedIterator is returned from FilterOrderValidated and is used to iterate over the raw logs and unpacked data for OrderValidated events raised by the Consideration contract.
type ConsiderationOrderValidatedIterator struct {
	Event *ConsiderationOrderValidated // Event containing the contract specifics and raw log

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
func (it *ConsiderationOrderValidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsiderationOrderValidated)
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
		it.Event = new(ConsiderationOrderValidated)
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
func (it *ConsiderationOrderValidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsiderationOrderValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsiderationOrderValidated represents a OrderValidated event raised by the Consideration contract.
type ConsiderationOrderValidated struct {
	OrderHash [32]byte
	Offerer   common.Address
	Zone      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderValidated is a free log retrieval operation binding the contract event 0xfde361574a066b44b3b5fe98a87108b7565e327327954c4faeea56a4e6491a0a.
//
// Solidity: event OrderValidated(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Consideration *ConsiderationFilterer) FilterOrderValidated(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*ConsiderationOrderValidatedIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Consideration.contract.FilterLogs(opts, "OrderValidated", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &ConsiderationOrderValidatedIterator{contract: _Consideration.contract, event: "OrderValidated", logs: logs, sub: sub}, nil
}

// WatchOrderValidated is a free log subscription operation binding the contract event 0xfde361574a066b44b3b5fe98a87108b7565e327327954c4faeea56a4e6491a0a.
//
// Solidity: event OrderValidated(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Consideration *ConsiderationFilterer) WatchOrderValidated(opts *bind.WatchOpts, sink chan<- *ConsiderationOrderValidated, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Consideration.contract.WatchLogs(opts, "OrderValidated", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsiderationOrderValidated)
				if err := _Consideration.contract.UnpackLog(event, "OrderValidated", log); err != nil {
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
func (_Consideration *ConsiderationFilterer) ParseOrderValidated(log types.Log) (*ConsiderationOrderValidated, error) {
	event := new(ConsiderationOrderValidated)
	if err := _Consideration.contract.UnpackLog(event, "OrderValidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
