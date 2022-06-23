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

// ConsiderationMetaData contains all meta data concerning the Consideration contract.
var ConsiderationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduitController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadContractSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadFraction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BadReturnValueFromERC20OnTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"BadSignatureV\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConsiderationCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortfallAmount\",\"type\":\"uint256\"}],\"name\":\"ConsiderationNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CriteriaNotEnabledForItem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"identifiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"ERC1155BatchTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EtherTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InexactFraction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientEtherSupplied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBasicOrderParameterEncoding\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidCallToConduit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCanceller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidConduit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidERC721TransferAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFulfillmentComponentData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidRestrictedOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchedFulfillmentOfferAndConsiderationComponents\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"}],\"name\":\"MissingFulfillmentComponentOnAggregation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingItemAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingOriginalConsiderationItems\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NoContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoReentrantCalls\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSpecifiedOrdersAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferAndConsiderationRequiredOnFulfillment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderAlreadyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderIsCancelled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderPartiallyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PartialFillsNotEnabledForOrder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnresolvedConsiderationCriteria\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnresolvedOfferCriteria\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"}],\"name\":\"NonceIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fulfiller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSpentItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structReceivedItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"}],\"name\":\"OrderValidated\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder\",\"name\":\"advancedOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"criteriaResolvers\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"}],\"name\":\"fulfillAdvancedOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder[]\",\"name\":\"advancedOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"criteriaResolvers\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"offerFulfillments\",\"type\":\"tuple[][]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"considerationFulfillments\",\"type\":\"tuple[][]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maximumFulfilled\",\"type\":\"uint256\"}],\"name\":\"fulfillAvailableAdvancedOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"availableOrders\",\"type\":\"bool[]\"},{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"offerFulfillments\",\"type\":\"tuple[][]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"considerationFulfillments\",\"type\":\"tuple[][]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maximumFulfilled\",\"type\":\"uint256\"}],\"name\":\"fulfillAvailableOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"availableOrders\",\"type\":\"bool[]\"},{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"considerationToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"considerationIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"offerToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offerIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumBasicOrderType\",\"name\":\"basicOrderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offererConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalAdditionalRecipients\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structAdditionalRecipient[]\",\"name\":\"additionalRecipients\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBasicOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"}],\"name\":\"fulfillBasicOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"}],\"name\":\"fulfillOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"getOrderStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidated\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSize\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"information\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"conduitController\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder[]\",\"name\":\"advancedOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"criteriaResolvers\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"offerComponents\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"considerationComponents\",\"type\":\"tuple[]\"}],\"internalType\":\"structFulfillment[]\",\"name\":\"fulfillments\",\"type\":\"tuple[]\"}],\"name\":\"matchAdvancedOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"offerComponents\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"considerationComponents\",\"type\":\"tuple[]\"}],\"internalType\":\"structFulfillment[]\",\"name\":\"fulfillments\",\"type\":\"tuple[]\"}],\"name\":\"matchOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"contractName\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"validate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"validated\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101c06040523480156200001257600080fd5b5060405162005b6e38038062005b6e8339810160408190526200003591620005c6565b8080808080808080806200004862000130565b610120526101005260e05260c081815260a0838152608085815246610140819052604080516020818101979097528082019890985260608801969096529086015230858201528351808603909101815293019091528151910120610160526001600160a01b03811661018081905260408051630a96ad3960e01b81528151630a96ad39926004808401939192918290030181865afa158015620000ef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001159190620005f8565b506101a0525050600160005550620006889650505050505050565b600080808080806200016260408051808201909152600d81526c21b7b739b4b232b930ba34b7b760991b602082015290565b805160209182012060408051808201825260018152603160f81b90840152519097507fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc696506000916200026591016909ecccccae492e8cada560b31b81526e1d5a5b9d0e081a5d195b551e5c194b608a1b600a8201526d1859191c995cdcc81d1bdad95b8b60921b60198201527f75696e74323536206964656e7469666965724f7243726974657269612c00000060278201527f75696e74323536207374617274416d6f756e742c0000000000000000000000006044820152701d5a5b9d0c8d4d88195b99105b5bdd5b9d607a1b6058820152602960f81b6069820152606a0190565b60408051601f1981840301815282825271086dedce6d2c8cae4c2e8d2dedc92e8cada560731b60208401526e1d5a5b9d0e081a5d195b551e5c194b608a1b60328401526d1859191c995cdcc81d1bdad95b8b60921b60418401527f75696e74323536206964656e7469666965724f7243726974657269612c000000604f8401527f75696e74323536207374617274416d6f756e742c000000000000000000000000606c840152711d5a5b9d0c8d4d88195b99105b5bdd5b9d0b60721b6080840152701859191c995cdcc81c9958da5c1a595b9d607a1b6092840152602960f81b60a384018190528251808503608401815260a485019093526f09ee4c8cae486dedae0dedccadce8e6560831b60c48501526f1859191c995cdcc81bd999995c995c8b60821b60d48501526c1859191c995cdcc81e9bdb994b609a1b60e48501527113d999995c925d195b56d7481bd999995c8b60721b60f18501527f436f6e73696465726174696f6e4974656d5b5d20636f6e73696465726174696f610103850152611b8b60f21b6101238501526f1d5a5b9d0e081bdc99195c951e5c194b60821b610125850152711d5a5b9d0c8d4d881cdd185c9d151a5b594b60721b6101358501526f1d5a5b9d0c8d4d88195b99151a5b594b60821b61014785015270189e5d195ccccc881e9bdb9952185cda0b607a1b6101578501526c1d5a5b9d0c8d4d881cd85b1d0b609a1b6101688501527f6279746573333220636f6e647569744b65792c000000000000000000000000006101758501526c75696e74323536206e6f6e636560981b6101888501526101958401529250906000906101960160408051601f19818403018152908290526c08a92a06e626488dedac2d2dc5609b1b60208301526b1cdd1c9a5b99c81b985b594b60a21b602d8301526e1cdd1c9a5b99c81d995c9cda5bdb8b608a1b60398301526f1d5a5b9d0c8d4d8818da185a5b92590b60821b60488301527f6164647265737320766572696679696e67436f6e7472616374000000000000006058830152602960f81b6071830152915060720160408051601f19818403018152908290528051602091820120855186830120855186840120919a5098509650620005a391839185918791016200065b565b604051602081830303815290604052805190602001209350505050909192939495565b600060208284031215620005d957600080fd5b81516001600160a01b0381168114620005f157600080fd5b9392505050565b600080604083850312156200060c57600080fd5b505080516020909101519092909150565b6000815160005b8181101562000640576020818501810151868301520162000624565b8181111562000650576000828601525b509290920192915050565b60006200067f620006786200067184886200061d565b866200061d565b846200061d565b95945050505050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a05161544e620007206000396000613545015260008181610d72015261351901526000612329015260006122590152600081816108a90152612541015260008181610838015261239f0152600081816107d101526124ba01526000612287015260006122d5015260006122ad015261544e6000f3fe6080604052600436106100e85760003560e01c8063a81744041161008a578063f47b774011610059578063f47b7740146102c7578063fb0f3ee1146102eb578063fb4c2af9146102fe578063fd9f1e101461031157600080fd5b8063a81744041461026d578063b3a34c4c14610280578063df7b0dac14610293578063ed98a574146102a657600080fd5b806355944a42116100c657806355944a42146101e8578063627cdcb91461020857806379df72bd1461021d578063881477321461023d57600080fd5b806306fdde03146100ed5780632d0335ab1461011857806346423aa714610146575b600080fd5b3480156100f957600080fd5b50610102610331565b60405161010f9190613fe4565b60405180910390f35b34801561012457600080fd5b5061013861013336600461401c565b610340565b60405190815260200161010f565b34801561015257600080fd5b506101c6610161366004614039565b6000908152600260209081526040918290208251608081018452905460ff8082161515808452610100830490911615159383018490526001600160781b036201000083048116958401869052600160881b909204909116606090920182905293919291565b604080519415158552921515602085015291830152606082015260800161010f565b6101fb6101f63660046145e3565b610360565b60405161010f9190614745565b34801561021457600080fd5b50610138610381565b34801561022957600080fd5b50610138610238366004614758565b61038b565b34801561024957600080fd5b5061025d610258366004614793565b61051c565b604051901515815260200161010f565b6101fb61027b3660046147d4565b61052f565b61025d61028e36600461483f565b6105ad565b61025d6102a1366004614888565b61061e565b6102b96102b43660046148ff565b61063c565b60405161010f9291906149a7565b3480156102d357600080fd5b506102dc6106c5565b60405161010f939291906149f6565b61025d6102f9366004614a29565b6106dd565b6102b961030c366004614a64565b6106e8565b34801561031d57600080fd5b5061025d61032c366004614793565b610716565b606061033b610722565b905090565b6001600160a01b0381166000908152600160205260408120545b92915050565b6060610377866103708688614b34565b8585610741565b9695505050505050565b600061033b61075c565b60408051610160810190915260009061035a90806103ac602086018661401c565b6001600160a01b031681526020018460200160208101906103cd919061401c565b6001600160a01b031681526020016103e86040860186614c68565b808060200260200160405190810160405280939291908181526020016000905b828210156104345761042560a08302860136819003810190614cb0565b81526020019060010190610408565b505050918352505060200161044c6060860186614ccc565b808060200260200160405190810160405280939291908181526020016000905b828210156104985761048960c08302860136819003810190614d14565b8152602001906001019061046c565b50505091835250506020016104b360a0860160808701614d30565b60038111156104c4576104c4614677565b81526020018460a0013581526020018460c0013581526020018460e0013581526020018461010001358152602001846101200135815260200184806060019061050d9190614ccc565b909152506101408401356107b9565b600061052883836108ff565b9392505050565b60606105a261053e8686610abb565b604080516000808252602082019092529061059a565b6105876040805160a081019091526000808252602082019081526020016000815260200160008152602001606081525090565b8152602001906001900390816105545790505b508585610741565b90505b949350505050565b60006105286105bb84610b76565b6040805160008082526020820190925290610617565b6106046040805160a081019091526000808252602082019081526020016000815260200160008152602001606081525090565b8152602001906001900390816105d15790505b5084610c21565b60006105a261062c86614d4b565b6106368587614b34565b84610c21565b6060806106b461064c8b8b610abb565b60408051600080825260208201909252906106a8565b6106956040805160a081019091526000808252602082019081526020016000815260200160008152602001606081525090565b8152602001906001900390816106625790505b508a8a8a8a8a8a610d12565b915091509850989650505050505050565b60606000806106d2610d51565b925092509250909192565b600061035a82610db0565b6060806107048b6106f98b8d614b34565b8a8a8a8a8a8a610d12565b91509150995099975050505050505050565b6000610528838361104a565b606060206000526d0d436f6e73696465726174696f6e602d5260606000f35b606061075185856001885161124b565b6105a2858484611562565b600061076661168d565b503360008181526001602081815260409283902080549092019182905591518181529092917f7ab0fc7de8910a6100b24df423c3d0835534506dca9473d30c3e7df51241b2cf910160405180910390a290565b610140820151604080519084015180516000939284927f000000000000000000000000000000000000000000000000000000000000000092602090910190845b81811015610826578251601f1901805186825260c0822086529052602093840193909201916001016107f9565b506020810260405120945050505060007f00000000000000000000000000000000000000000000000000000000000000009150604051602060608901510160005b86811015610894578151601f1901805186825260e082208552905260209283019290910190600101610867565b505060408051602087029020601f198a0180517f00000000000000000000000000000000000000000000000000000000000000008252928b01805197815260608c018051938152610140909c019a8b5261018082209390915295909552939097525050925250919050565b600061090961168d565b60008083815b81811015610aae573687878381811061092a5761092a614d57565b905060200281019061093c9190614d6d565b9050366109498280614d8d565b9050610958602082018261401c565b945061096b61096682614da4565b6116b2565b60008181526002602090815260408083208151608081018352905460ff808216151583526101008204161515938201939093526001600160781b03620100008404811692820192909252600160881b9092041660608201529197506109d5908890839060016116ed565b508051610aa057610a2886886109ee6020870187614db0565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506117a692505050565b600087815260026020908152604091829020805460ff19166001179055610a5391840190840161401c565b6001600160a01b0316866001600160a01b03167ffde361574a066b44b3b5fe98a87108b7565e327327954c4faeea56a4e6491a0a89604051610a9791815260200190565b60405180910390a35b83600101935050505061090f565b5060019695505050505050565b606081806001600160401b03811115610ad657610ad6614052565b604051908082528060200260200182016040528015610b0f57816020015b610afc613ea3565b815260200190600190039081610af45790505b50915060005b81811015610b6e57610b49858583818110610b3257610b32614d57565b9050602002810190610b449190614d6d565b610b76565b838281518110610b5b57610b5b614d57565b6020908102919091010152600101610b15565b505092915050565b610b7e613ea3565b6040805160a0810190915280610b948480614d8d565b610b9d90614da4565b815260200160016001600160781b0316815260200160016001600160781b03168152602001838060200190610bd29190614db0565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509385525050604080516020818101909252928352909201525092915050565b6000610c2b6117fb565b60606000806000610c3f888860018761180a565b604080516001808252818301909252939650919450925060009190816020015b610c67613ea3565b815260200190600190039081610c5f5790505090508881600081518110610c9057610c90614d57565b6020026020010181905250610ca58189611a93565b600081600081518110610cba57610cba614d57565b6020026020010151600001519050610cda8185858461012001518c611e10565b610cf885826000015183602001513385604001518660600151612014565b610d026001600055565b5060019998505050505050505050565b606080610d228a8a60008661124b565b610d408a610d30898b614e44565b610d3a888a614e44565b87612079565b909b909a5098505050505050505050565b6060600080610d5e612255565b6040805160018082528183019092529193507f000000000000000000000000000000000000000000000000000000000000000092506020820181803683375050603160f81b6020830152509391925090565b600060046101243590810490600316600182113415811480610dec57604051630a61be9f60e41b81523460048201526024015b60405180910390fd5b5060008060006003861160a08102602401359350600287146002881160028903020192506001830181026002861502880103915050610e2f88868487878661234b565b6000610e4160808a0160608b0161401c565b90506101c46003881160200201356000866005811115610e6357610e63614677565b03610eaf57610e8e83610e7c60c08d0160a08e0161401c565b84338e60c001358f60e001358761268b565b610eaa60408b013583610ea56102008e018e614f16565b61273f565b610d02565b6040805160208082528183019092526000916020820181803683370190505090506002896005811115610ee457610ee4614677565b03610f4457610f0f610efc60c08d0160a08e0161401c565b84338e60c001358f60e001358787612804565b610f3f3384610f2160208f018f61401c565b8e604001358f806102000190610f379190614f16565b600088612851565b611030565b6003896005811115610f5857610f58614677565b03610f8357610f0f610f7060c08d0160a08e0161401c565b84338e60c001358f60e0013587876128d8565b6004896005811115610f9757610f97614677565b03610ff557610fbf610fac60208d018d61401c565b33858e602001358f604001358787612804565b610f3f83338d60a0016020810190610fd7919061401c565b8e60e001358f806102000190610fed9190614f16565b600188612851565b61101861100560208d018d61401c565b33858e602001358f6040013587876128d8565b61103083338d60a0016020810190610fd7919061401c565b6110398161290e565b505060019998505050505050505050565b600061105461168d565b60008083815b81811015610aae573687878381811061107557611075614d57565b90506020028101906110879190614d8d565b9050611096602082018261401c565b94506110a8604082016020830161401c565b9350336001600160a01b038616148015906110cc5750336001600160a01b03851614155b156110ea5760405163203b1cdd60e21b815260040160405180910390fd5b60006111d9604051806101600160405280886001600160a01b03168152602001876001600160a01b031681526020018480604001906111299190614c68565b808060200260200160405190810160405280939291908181526020016000905b828210156111755761116660a08302860136819003810190614cb0565b81526020019060010190611149565b505050918352505060200161118d6060860186614ccc565b808060200260200160405190810160405280939291908181526020016000905b82821015610498576111ca60c08302860136819003810190614d14565b815260200190600101906111ad565b60008181526002602052604090819020805461ffff1916610100179055519091506001600160a01b0380871691908816907f6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d906112399085815260200190565b60405180910390a3505060010161105a565b6112536117fb565b83516000816001600160401b0381111561126f5761126f614052565b604051908082528060200260200182016040528015611298578160200160208202803683370190505b5090506000815260005b828110156114b75760008782815181106112be576112be614d57565b60200260200101519050846000036112e35760006020909101526001810182526114af565b60008060006112f4848b8b8961180a565b9250925092506001850186528160000361131b5750506000602090920191909152506114af565b8286868151811061132e5761132e614d57565b6020908102919091010152835160a081015160c0820151604090920151600019909a019990918290039042839003908183039060005b81518110156113fb57600082828151811061138157611381614d57565b60200260200101519050600061139c8a8a8460800151612937565b905081608001518260600151036113b957606082018190526113ce565b6113c88a8a8460600151612937565b60608301525b6080820181905260608201516113e9908288888b6000612987565b60609092019190915250600101611364565b5088516060015160005b81518110156114a357600082828151811061142257611422614d57565b60200260200101519050600061143d8b8b8460800151612937565b9050816080015182606001510361145a576060820181905261146f565b6114698b8b8460600151612937565b60608301525b60808201819052606082015161148a908289898c6001612987565b60608301525060a0810151608090910152600101611405565b50505050505050505050505b6001016112a2565b506114c28686611a93565b8315330260005b83811015611558576000801b8382815181106114e7576114e7614d57565b6020026020010151031561155057600088828151811061150957611509614d57565b602002602001015160000151905061154e84838151811061152c5761152c614d57565b6020026020010151826000015183602001518685604001518660600151612014565b505b6001016114c9565b5050505050505050565b606081806001600160401b0381111561157d5761157d614052565b6040519080825280602002602001820160405280156115b657816020015b6115a3613ed7565b81526020019060019003908161159b5790505b5091506000805b8281101561166b57368686838181106115d8576115d8614d57565b90506020028101906115ea9190614d6d565b9050600061160e896115fc8480614f16565b6116096020870187614f16565b6129e2565b905080602001516001600160a01b03168160000151608001516001600160a01b03160361164057600184019350611661565b80868585038151811061165557611655614d57565b60200260200101819052505b50506001016115bd565b508015611679578083510383525b506116848583612cd3565b50509392505050565b6001600054146116b057604051637fa8a98760e01b815260040160405180910390fd5b565b60006116c8826060015151836101400151612ee8565b81516001600160a01b031660009081526001602052604090205461035a9083906107b9565b600083602001511561172357811561171b57604051630694555d60e21b815260048101869052602401610de3565b5060006105a5565b60408401516001600160781b03161561179b5782156117585760405163ee9e0e6360e01b815260048101869052602401610de3565b83606001516001600160781b031684604001516001600160781b03161061179b57811561171b576040516310fda3e160e01b815260048101869052602401610de3565b506001949350505050565b336001600160a01b038416036117bb57505050565b60006117e86117c8612255565b61190160f01b600090815260029190915260228581526042822091905290565b90506117f5848284612f09565b50505050565b61180361168d565b6002600055565b6000806000808760000151905061182a8160a001518260c001518861304c565b61183e575060009250829150819050611a89565b602088015160408901516001600160781b03918216911680821180611861575081155b1561187f57604051632d02959960e11b815260040160405180910390fd5b808210801561189357506080830151600116155b156118b15760405163a11b63ff60e01b815260040160405180910390fd5b6118ba836116b2565b95506118dc8a8a89898760e00151886080015189600001518a60200151613092565b60008681526002602090815260408083208151608081018352905460ff808216151583526101008204161515938201939093526001600160781b03620100008404811692820192909252600160881b9092041660608201529061194390889083908c6116ed565b611958575060009450849350611a8992505050565b8051611971576119718460000151888d606001516117a6565b604081015160608201516001600160781b0391821691168015611a3557836001036119a1578094508093506119cd565b8381146119cd576119b28483614f75565b91506119be8186614f75565b94506119ca8185614f75565b93505b836119d88684614f94565b11156119e45781840394505b600089815260026020526040902080546001600160781b03868116600160881b026001600160881b03868a019290921662010000026001600160881b03199093169290921760011716179055611a7e565b600089815260026020526040902080546001600160781b03868116600160881b026001600160881b0391891662010000026001600160881b031990931692909217600117161790555b509295509093505050505b9450945094915050565b8051825160005b82811015611cda576000848281518110611ab657611ab6614d57565b60200260200101519050600081600001519050838110611ae9576040516321a561b160e21b815260040160405180910390fd5b868181518110611afb57611afb614d57565b6020026020010151602001516001600160781b0316600003611b1e575050611cd2565b6000878281518110611b3257611b32614d57565b60209081029190910101515160408401519091506000808086602001516001811115611b6057611b60614677565b03611bfa57604084015180518410611b8b57604051635fd9fc6760e11b815260040160405180910390fd5b6000818581518110611b9f57611b9f614d57565b602090810291909101015180516040820151909550935090506004841460030381816005811115611bd257611bd2614677565b90816005811115611be557611be5614677565b90525050606088015160409091015250611c8b565b606084015180518410611c20576040516330446bef60e11b815260040160405180910390fd5b6000818581518110611c3457611c34614d57565b602090810291909101015180516040820151909550935090506004841460030381816005811115611c6757611c67614677565b90816005811115611c7a57611c7a614677565b905250506060880151604090910152505b611c958260031090565b611cb257604051634a75b57b60e11b815260040160405180910390fd5b8015611ccb57611ccb8660600151828860800151613173565b5050505050505b600101611a9a565b5060005b81811015611e09576000858281518110611cfa57611cfa614d57565b6020026020010151905080602001516001600160781b0316600003611d1f5750611e01565b6000868381518110611d3357611d33614d57565b60200260200101516000015190506000816060015151905060005b81811015611daa57611d8183606001518281518110611d6f57611d6f614d57565b60200260200101516000015160031090565b15611da25760405160016202297360e61b0319815260040160405180910390fd5b600101611d4e565b505060408101515160005b81811015611dfc57611dd683604001518281518110611d6f57611d6f614d57565b15611df45760405163a6cfc67360e01b815260040160405180910390fd5b600101611db5565b505050505b600101611cde565b5050505050565b60008560a001518660c00151611e269190614fac565b905060008660a0015142611e3a9190614fac565b90506000611e488284614fac565b60408051602080825281830190925291925034916000916020820181803683370190505090506131e360005b8b6040015151811015611f2c5760008c604001518281518110611e9957611e99614d57565b602002602001015190506000611ebe826060015183608001518f8f8c8c8f60006132a0565b606083018190523360808401529050600082516005811115611ee257611ee2614677565b03611f0e5785811115611f0857604051631a783b8d60e01b815260040160405180910390fd5b80860395505b611f22828f600001518d888863ffffffff16565b5050600101611e74565b506131e3905060005b8b6060015151811015611fed5760008c606001518281518110611f5a57611f5a614d57565b602002602001015190506000611f7f826060015183608001518f8f8c8c8f60016132a0565b6060830181905260a083015160808401529050600082516005811115611fa757611fa7614677565b03611fd35785811115611fcd57604051631a783b8d60e01b815260040160405180910390fd5b80860395505b611fe382338c888863ffffffff16565b5050600101611f35565b5050611ff88161290e565b81156120085761200833836132ec565b50505050505050505050565b60608290506060829050856001600160a01b0316876001600160a01b03167f9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f318a8886866040516120679493929190614ffd565b60405180910390a35050505050505050565b82518251606091829161208c8183614f94565b6001600160401b038111156120a3576120a3614052565b6040519080825280602002602001820160405280156120dc57816020015b6120c9613ed7565b8152602001906001900390816120c15790505b5092506000805b838110156121755760008982815181106120ff576120ff614d57565b6020026020010151905060006121188c6000848c613340565b905080602001516001600160a01b03168160000151608001516001600160a01b03160361214a5760018401935061216b565b80878585038151811061215f5761215f614d57565b60200260200101819052505b50506001016120e3565b5060005b8281101561220d57600088828151811061219557612195614d57565b6020026020010151905060006121ae8c6001848c613340565b905080602001516001600160a01b03168160000151608001516001600160a01b0316036121e057600184019350612203565b80878588860103815181106121f7576121f7614d57565b60200260200101819052505b5050600101612179565b50801561221b578084510384525b50825160000361223e5760405163d5da9a1b60e01b815260040160405180910390fd5b6122488884612cd3565b9350505094509492505050565b60007f000000000000000000000000000000000000000000000000000000000000000046146123265761033b604080517f000000000000000000000000000000000000000000000000000000000000000060208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b507f000000000000000000000000000000000000000000000000000000000000000090565b6123536117fb565b612369866101200135876101400135600161304c565b506123726133db565b61239a612383610200880188614f16565b61238f91506001614f94565b876101e00135612ee8565b6000807f00000000000000000000000000000000000000000000000000000000000000009050806080528560a0526060602460c037604060646101203760e06080206101605261026435602081026102a0016001610264350181526020810190508781526080602460208301376101608760a0528660c052600060e05261020435925060005b8381101561246b578060400261028401602081610100376040816101203760208301925060e0608020835260a084019350898452886020850152604081606086013750600101612420565b60206001850102610160206060526102643593505b838110156124b1578060400261028401915060a0830192508883528760208401526040826060850137600101612480565b505050505060007f00000000000000000000000000000000000000000000000000000000000000009050806080528260a052606060c460c03760206101046101203760c0608020600052602060002060e052602061026435026102000160018152836020820152606060c4604083013750506084356001600160a01b0381166000908152600160205260408120547f000000000000000000000000000000000000000000000000000000000000000060808190529091506040608460a03760605161010052886101205260a061014461014037816101e0526101806080209350505050602061026435026101800181815233602082015260806040820152610120606082015260a061026435026101e00160a4356084357f9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f318385a35050600060605261262681886101600135888a6060016020810190612611919061401c565b61262160a08d0160808e0161401c565b61341c565b6126828161263a60808a0160608b0161401c565b6126486102208b018b614db0565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061346c92505050565b50505050505050565b80156126e75760006040519050632671a55160e11b815260206004820152600160248201528760448201528660648201528560848201528460a48201528360c48201528260e48201526126e1828261010461350e565b50612682565b60028760058111156126fb576126fb614677565b0361273257816001146127215760405163efcc00b160e01b815260040160405180910390fd5b61272d868686866135fe565b612682565b61268286868686866136bf565b348160005b818110156127b2573685858381811061275f5761275f614d57565b6040029190910191505080358481111561278c57604051631a783b8d60e01b815260040160405180910390fd5b6127a561279f604084016020850161401c565b826132ec565b9093039250600101612744565b50818611156127d457604051631a783b8d60e01b815260040160405180910390fd5b6127de85876132ec565b858211156127f2576127f2338784036132ec565b6127fc6001600055565b505050505050565b61280e81836137a3565b8161284057826001146128345760405163efcc00b160e01b815260040160405180910390fd5b61272d878787876135fe565b612682828260028a8a8a8a8a6137c2565b602082026101e403358360005b818110156128bf573687878381811061287957612879614d57565b604002919091019150508035861561289857612895818b614fac565b99505b6128b58b8e6128ad604086016020870161401c565b84898b613842565b505060010161285e565b506128ce888b8b8a8688613842565b6120086001600055565b6128e18361387d565b6128eb81836137a3565b816128fd5761272d87878787876136bf565b612682828260038a8a8a8a8a6137c2565b805160401461291a5750565b6000612927826020015190565b9050612933818361389e565b5050565b6000828403612947575080610528565b6000838584091590508061296e5760405163c63cf08960e01b815260040160405180910390fd5b600061297a8685614f75565b9490940495945050505050565b60008587146129d7576000821561299f575060001983015b6000816129ac888a614f75565b6129b6888c614f75565b6129c09190614f94565b6129ca9190614f94565b8590049250610377915050565b509395945050505050565b6129ea613ed7565b8315806129f5575081155b15612a1357604051634c74edb760e11b815260040160405180910390fd5b612a1b613ed7565b612a78878585808060200260200160405190810160405280939291908181526020016000905b82821015612a6d57612a5e60408302860136819003810190615097565b81526020019060010190612a41565b5050505050836138c2565b805160408051602080890282018101909252878152612ad7918a91908a908a90819060009085015b82821015612acc57612abd60408302860136819003810190615097565b81526020019060010190612aa0565b505050505085613a6e565b80516005811115612aea57612aea614677565b8351516005811115612afe57612afe614677565b141580612b29575080602001516001600160a01b03168360000151602001516001600160a01b031614155b80612b405750806040015183600001516040015114155b15612b5e576040516309cfb45560e01b815260040160405180910390fd5b82600001516060015181606001511115612c1557600085856000818110612b8757612b87614d57565b905060400201803603810190612b9d9190615097565b90508360000151606001518260600151612bb79190614fac565b89826000015181518110612bcd57612bcd614d57565b60200260200101516000015160600151826020015181518110612bf257612bf2614d57565b602090810291909101015160609081019190915284518101519083015250612ca7565b600087876000818110612c2a57612c2a614d57565b905060400201803603810190612c409190615097565b90508160600151846000015160600151612c5a9190614fac565b89826000015181518110612c7057612c70614d57565b60200260200101516000015160400151826020015181518110612c9557612c95614d57565b60200260200101516060018181525050505b60608082015184519091015260809081015183516001600160a01b039091169101525095945050505050565b8151606090806001600160401b03811115612cf057612cf0614052565b604051908082528060200260200182016040528015612d19578160200160208202803683370190505b50915060005b81811015612dff576000858281518110612d3b57612d3b614d57565b6020026020010151905080602001516001600160781b0316600003612d605750612df7565b6001848381518110612d7457612d74614d57565b9115156020928302919091019091015280516060015160005b8151811015612df3576000828281518110612daa57612daa614d57565b602002602001015160600151905080600014612dea576040516314bea84160e31b8152600481018690526024810183905260448101829052606401610de3565b50600101612d8d565b5050505b600101612d1f565b5060408051602080825281830190925234916000919060208201818036833701905050905060005b8551811015612ebb576000868281518110612e4457612e44614d57565b60209081029190910101518051909150600081516005811115612e6957612e69614677565b03612e9d578481606001511115612e9357604051631a783b8d60e01b815260040160405180910390fd5b8060600151850394505b612eb18183602001518460400151876131e3565b5050600101612e27565b50612ec58161290e565b8115612ed557612ed533836132ec565b612edf6001600055565b50505092915050565b8082101561293357604051632335530b60e11b815260040160405180910390fd5b60008060008351604003612f3a57505050602081015160408201516001600160ff1b0381169060ff1c601b01612fa0565b8351604103612f955750505060208101516040820151606083015160001a601b8114801590612f6d57508060ff16601c14155b15612f9057604051630f801e8560e11b815260ff82166004820152602401610de3565b612fa0565b6127fc868686613bf8565b6040805160008082526020820180845288905260ff841692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015612ff4573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661302857604051638baa579f60e01b815260040160405180910390fd5b866001600160a01b0316816001600160a01b03161461268257612682878787613bf8565b60004284118061305c5750428311155b15613088578115613080576040516337bf561360e11b815260040160405180910390fd5b506000610528565b5060019392505050565b60018360038111156130a6576130a6614677565b1180156130bc5750336001600160a01b03821614155b80156130d15750336001600160a01b03831614155b15611558576080880151511580156130e857508651155b156130fe576130f981868487613c6f565b611558565b600061315c82633313157060e01b88338d8c8e60405160240161312595949392919061526b565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613cbb565b90506131688187613cd0565b505050505050505050565b60008360208301602084510281015b808210156131bf57815180841180156131a25781600052846020526131ab565b84600052816020525b505060406000209250602082019150613182565b505083149050806117f5576040516309bde33960e01b815260040160405180910390fd5b6000845160058111156131f8576131f8614677565b036132145761320f846080015185606001516132ec565b6117f5565b60018451600581111561322957613229614677565b036132485761320f846020015184866080015187606001518686613842565b60028451600581111561325d5761325d614677565b036132815761320f8460200151848660800151876040015188606001518787612804565b6117f584602001518486608001518760400151886060015187876128d8565b60008789036132bb576132b487878a612937565b90506132e0565b6132dd6132c988888c612937565b6132d489898c612937565b87878787612987565b90505b98975050505050505050565b6132f58161387d565b600080600080600085875af190508061333b57613310613d2a565b60405163470c7c1d60e01b81526001600160a01b038416600482015260248101839052604401610de3565b505050565b613348613ed7565b825160000361336c578360405163375c24c160e01b8152600401610de391906153f1565b600084600181111561338057613380614677565b0361339557613390858483613a6e565b6133ae565b6133a08584836138c2565b336020820152604081018290525b8051606001516000036105a557602081015181516001600160a01b03909116608090910152949350505050565b610244356102606102643560400201146004356020146102243561024014161680613419576040516339f3e3fd60e01b815260040160405180910390fd5b50565b600183600381111561343057613430614677565b1180156134465750336001600160a01b03821614155b801561345b5750336001600160a01b03831614155b15611e0957611e0981868487613c6f565b6000838152600260209081526040918290208251608081018452905460ff808216151583526101008204161515928201929092526001600160781b03620100008304811693820193909352600160881b90910490911660608201526134d484826001806116ed565b5080516134e6576134e68385846117a6565b5050506000908152600260205260409020710100000000000000000000000000000100019055565b6040805160ff60a01b7f000000000000000000000000000000000000000000000000000000000000000017600090815260208681527f000000000000000000000000000000000000000000000000000000000000000084526055600b20929093528080526001600160a01b039091169181848682865af19050806135b857613594613d2a565b60405163344f54f560e21b81526001600160a01b0383166004820152602401610de3565b6000516001600160e01b03198116632671a55160e11b146127fc57604051630e7ccd9360e11b8152600481018790526001600160a01b0384166024820152604401610de3565b833b61361957632f8aeb3960e11b6000528360045260246000fd5b6040516323b872dd60e01b6000528360045282602452816044526000806064600080895af1806136b0573d1561368a5760203d0460208304816003028183111561367157818303600302610200838002858002030401015b5a602082011015613686573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005285600452846024528360445282606452600160845260a46000fd5b50604052505060006060525050565b843b6136da57632f8aeb3960e11b6000528460045260246000fd5b60405160805160a05160c051637921219560e11b6000528760045286602452856044528460645260a0608452600060a45260008060c46000808d5af180613787573d156137625760203d0460208604816003028183111561374957818303600302610200838002858002030401015b5a60208201101561375e573d6000803e3d6000fd5b5050505b63f486bc8760e01b600052896004528860245287604452866064528560845260a46000fd5b5060809290925260a05260c05260405250506000606052505050565b60006137b0836020015190565b905081811461333b5761333b8361290e565b600060208851036137fd5750604080885260208089018a9052632671a55160e11b91890191909152604488015260016064880181905261380c565b50606487018051600101908190525b603c60c082028901038781528660208201528560408201528460608201528360808201528260a082015250505050505050505050565b61384b8361387d565b61385581836137a3565b8161386b5761386686868686613d6f565b6127fc565b6127fc8282600189898960008a6137c2565b806000036134195760405163246cf94560e21b815260040160405180910390fd5b6064810151604082019060c0026044016138b984838361350e565b50506020905250565b6138ee565b637fda727960e01b60005260206000fd5b634e487b7160e01b600052601160045260246000fd5b6020820180515184518110613905576139056138c7565b602081026020860101516060815101516020845101518151811061392b5761392b6138c7565b60208102602083010151600080602086015115613952575050606081018051600090915280155b885183518152602084015160208201526040840151604082015260a084015160808201526060812060208c51028c015b808b1015613a295760208b019a508a515199508d518a106139a5576139a56138c7565b60208a0260208f010151985060208901511561398257606089510151975060208b5101519650875187106139db576139db6138c7565b602087026020890101519550606086018051860181511587821060011b1786179550809650506000815250606086208214608084015160a08801511416613a2457613a246138c7565b613982565b50506060018290528060018114613a475760028114613a5857613a60565b63246cf94560e21b60005260206000fd5b613a606138d8565b505050505050505050505050565b6020820180515184518110613a8557613a856138c7565b602081026020860101518051604081015160208551015181518110613aac57613aac6138c7565b60208102602083010151600080602087015115613ad3575050606081018051600090915280155b8951336080820152835181526020840151602082015260408401516040820152865160208c015261012087015160408c015260608120905060208c51028c015b808b1015613bc75760208b019a508a515199508d518a10613b3657613b366138c7565b60208a0260208f0101519850602089015115613b1357885197506040880151965060208b510151955086518610613b6f57613b6f6138c7565b602086026020880101519450606085018051850181511586821060011b178517945080955050600081525060608520821460408d01516101208a01511460208e01518a51141616613bc257613bc26138c7565b613b13565b50508160608b5101528060018114613a475760028103613be957613be96138d8565b50505050505050505050505050565b6000613c1984631626ba7e60e01b85856040516024016131259291906153ff565b905080613c4157613c28613d2a565b604051634f7fb80d60e01b815260040160405180910390fd5b613c51630b135d3f60e11b613e75565b156117f557604051632057875960e21b815260040160405180910390fd5b604051602481018490523360448201526001600160a01b038316606482015260848101829052600090613caf9086906303874c7760e21b9060a401613125565b9050611e098185613cd0565b6000806000835160208501865afa9392505050565b81613cf957613cdd613d2a565b604051633ed4053f60e21b815260048101829052602401610de3565b613d096303874c7760e21b613e75565b1561293357604051633ed4053f60e21b815260048101829052602401610de3565b3d156116b05760203d046020604051048160030281831115613d5a57818303600302610200838002858002030401015b5a60208201101561333b573d6000803e3d6000fd5b6040516323b872dd60e01b600052836004528260245281604452602060006064600080895af1803d15601f3d116001600051141617163d15158116613e655780873b151516613e655780613e505781613e2f573d15613e095760203d04602084048160030281831115613df057818303600302610200838002858002030401015b5a602082011015613e05573d6000803e3d6000fd5b5050505b63f486bc8760e01b60005286600452856024528460445260006064528360845260a46000fd5b639889192360e01b6000528660045285602452846044528360645260846000fd5b632f8aeb3960e11b6000528660045260246000fd5b5050604052505060006060525050565b60008060203d03613e8b5760206000803e506000515b6001600160e01b031990811692169190911415919050565b6040518060a00160405280613eb6613f1a565b81526000602082018190526040820152606080820181905260809091015290565b60408051610100810182526000606082018181526080830182905260a0830182905260c0830182905260e083018290528252602082018190529181019190915290565b60405180610160016040528060006001600160a01b0316815260200160006001600160a01b03168152602001606081526020016060815260200160006003811115613f6757613f67614677565b815260006020820181905260408201819052606082018190526080820181905260a0820181905260c09091015290565b6000815180845260005b81811015613fbd57602081850181015186830182015201613fa1565b81811115613fcf576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006105286020830184613f97565b6001600160a01b038116811461341957600080fd5b803561401781613ff7565b919050565b60006020828403121561402e57600080fd5b813561052881613ff7565b60006020828403121561404b57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b038111828210171561408a5761408a614052565b60405290565b60405161016081016001600160401b038111828210171561408a5761408a614052565b604051601f8201601f191681016001600160401b03811182821017156140db576140db614052565b604052919050565b60006001600160401b038211156140fc576140fc614052565b5060051b60200190565b80356006811061401757600080fd5b600060a0828403121561412757600080fd5b61412f614068565b905061413a82614106565b8152602082013561414a81613ff7565b8060208301525060408201356040820152606082013560608201526080820135608082015292915050565b600082601f83011261418657600080fd5b8135602061419b614196836140e3565b6140b3565b82815260a092830285018201928282019190878511156141ba57600080fd5b8387015b858110156141dd576141d08982614115565b84529284019281016141be565b5090979650505050505050565b600060c082840312156141fc57600080fd5b60405160c081018181106001600160401b038211171561421e5761421e614052565b60405290508061422d83614106565b8152602083013561423d81613ff7565b8060208301525060408301356040820152606083013560608201526080830135608082015260a083013561427081613ff7565b60a0919091015292915050565b600082601f83011261428e57600080fd5b8135602061429e614196836140e3565b82815260c092830285018201928282019190878511156142bd57600080fd5b8387015b858110156141dd576142d389826141ea565b84529284019281016142c1565b80356004811061401757600080fd5b6000610160828403121561430257600080fd5b61430a614090565b90506143158261400c565b81526143236020830161400c565b602082015260408201356001600160401b038082111561434257600080fd5b61434e85838601614175565b6040840152606084013591508082111561436757600080fd5b506143748482850161427d565b606083015250614386608083016142e0565b608082015260a082013560a082015260c082013560c082015260e082013560e082015261010080830135818301525061012080830135818301525061014080830135818301525092915050565b80356001600160781b038116811461401757600080fd5b600082601f8301126143fb57600080fd5b81356001600160401b0381111561441457614414614052565b614427601f8201601f19166020016140b3565b81815284602083860101111561443c57600080fd5b816020850160208301376000918101602001919091529392505050565b600060a0828403121561446b57600080fd5b614473614068565b905081356001600160401b038082111561448c57600080fd5b614498858386016142ef565b83526144a6602085016143d3565b60208401526144b7604085016143d3565b604084015260608401359150808211156144d057600080fd5b6144dc858386016143ea565b606084015260808401359150808211156144f557600080fd5b50614502848285016143ea565b60808301525092915050565b600082601f83011261451f57600080fd5b8135602061452f614196836140e3565b82815260059290921b8401810191818101908684111561454e57600080fd5b8286015b8481101561458d5780356001600160401b038111156145715760008081fd5b61457f8986838b0101614459565b845250918301918301614552565b509695505050505050565b60008083601f8401126145aa57600080fd5b5081356001600160401b038111156145c157600080fd5b6020830191508360208260051b85010111156145dc57600080fd5b9250929050565b6000806000806000606086880312156145fb57600080fd5b85356001600160401b038082111561461257600080fd5b61461e89838a0161450e565b9650602088013591508082111561463457600080fd5b61464089838a01614598565b9096509450604088013591508082111561465957600080fd5b5061466688828901614598565b969995985093965092949392505050565b634e487b7160e01b600052602160045260246000fd5b6006811061469d5761469d614677565b9052565b6146ac82825161468d565b6020818101516001600160a01b0390811691840191909152604080830151908401526060808301519084015260809182015116910152565b600081518084526020808501945080840160005b8381101561473a57815161470d8882516146a1565b808401516001600160a01b031660a08901526040015160c088015260e090960195908201906001016146f8565b509495945050505050565b60208152600061052860208301846146e4565b60006020828403121561476a57600080fd5b81356001600160401b0381111561478057600080fd5b8201610160818503121561052857600080fd5b600080602083850312156147a657600080fd5b82356001600160401b038111156147bc57600080fd5b6147c885828601614598565b90969095509350505050565b600080600080604085870312156147ea57600080fd5b84356001600160401b038082111561480157600080fd5b61480d88838901614598565b9096509450602087013591508082111561482657600080fd5b5061483387828801614598565b95989497509550505050565b6000806040838503121561485257600080fd5b82356001600160401b0381111561486857600080fd5b83016040818603121561487a57600080fd5b946020939093013593505050565b6000806000806060858703121561489e57600080fd5b84356001600160401b03808211156148b557600080fd5b9086019060a082890312156148c957600080fd5b909450602086013590808211156148df57600080fd5b506148ec87828801614598565b9598909750949560400135949350505050565b60008060008060008060008060a0898b03121561491b57600080fd5b88356001600160401b038082111561493257600080fd5b61493e8c838d01614598565b909a50985060208b013591508082111561495757600080fd5b6149638c838d01614598565b909850965060408b013591508082111561497c57600080fd5b506149898b828c01614598565b999c989b509699959896976060870135966080013595509350505050565b604080825283519082018190526000906020906060840190828701845b828110156149e25781511515845292840192908401906001016149c4565b5050508381038285015261037781866146e4565b606081526000614a096060830186613f97565b6020830194909452506001600160a01b0391909116604090910152919050565b600060208284031215614a3b57600080fd5b81356001600160401b03811115614a5157600080fd5b8201610240818503121561052857600080fd5b600080600080600080600080600060c08a8c031215614a8257600080fd5b89356001600160401b0380821115614a9957600080fd5b614aa58d838e0161450e565b9a5060208c0135915080821115614abb57600080fd5b614ac78d838e01614598565b909a50985060408c0135915080821115614ae057600080fd5b614aec8d838e01614598565b909850965060608c0135915080821115614b0557600080fd5b50614b128c828d01614598565b9a9d999c50979a96999598959660808101359660a09091013595509350505050565b6000614b42614196846140e3565b83815260208082019190600586811b860136811115614b6057600080fd5b865b81811015614c5b5780356001600160401b0380821115614b825760008081fd5b818a01915060a08236031215614b985760008081fd5b614ba0614068565b823581528683013560028110614bb65760008081fd5b81880152604083810135908201526060808401359082015260808084013583811115614be25760008081fd5b939093019236601f850112614bf957600092508283fd5b83359250614c09614196846140e3565b83815292871b84018801928881019036851115614c265760008081fd5b948901945b84861015614c4457853582529489019490890190614c2b565b918301919091525088525050948301948301614b62565b5092979650505050505050565b6000808335601e19843603018112614c7f57600080fd5b8301803591506001600160401b03821115614c9957600080fd5b602001915060a0810236038213156145dc57600080fd5b600060a08284031215614cc257600080fd5b6105288383614115565b6000808335601e19843603018112614ce357600080fd5b8301803591506001600160401b03821115614cfd57600080fd5b602001915060c0810236038213156145dc57600080fd5b600060c08284031215614d2657600080fd5b61052883836141ea565b600060208284031215614d4257600080fd5b610528826142e0565b600061035a3683614459565b634e487b7160e01b600052603260045260246000fd5b60008235603e19833603018112614d8357600080fd5b9190910192915050565b6000823561015e19833603018112614d8357600080fd5b600061035a36836142ef565b6000808335601e19843603018112614dc757600080fd5b8301803591506001600160401b03821115614de157600080fd5b6020019150368190038213156145dc57600080fd5b600060408284031215614e0857600080fd5b604051604081018181106001600160401b0382111715614e2a57614e2a614052565b604052823581526020928301359281019290925250919050565b6000614e52614196846140e3565b80848252602080830192508560051b850136811115614e7057600080fd5b855b81811015614f0a5780356001600160401b03811115614e915760008081fd5b870136601f820112614ea35760008081fd5b8035614eb1614196826140e3565b81815260069190911b82018501908581019036831115614ed15760008081fd5b928601925b82841015614efa57614ee83685614df6565b82528682019150604084019350614ed6565b8852505050938201938201614e72565b50919695505050505050565b6000808335601e19843603018112614f2d57600080fd5b8301803591506001600160401b03821115614f4757600080fd5b6020019150600681901b36038213156145dc57600080fd5b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615614f8f57614f8f614f5f565b500290565b60008219821115614fa757614fa7614f5f565b500190565b600082821015614fbe57614fbe614f5f565b500390565b600081518084526020808501945080840160005b8381101561473a57614fea8783516146a1565b60a0969096019590820190600101614fd7565b60006080808301878452602060018060a01b03808916828701526040848188015283895180865260a089019150848b01955060005b8181101561507357865161504784825161468d565b808701518616848801528481015185850152606090810151908401529585019591870191600101615032565b50508781036060890152615087818a614fc3565b9c9b505050505050505050505050565b6000604082840312156150a957600080fd5b6105288383614df6565b600081518084526020808501945080840160005b8381101561473a5781516150dc88825161468d565b838101516001600160a01b03168885015260408082015190890152606080820151908901526080908101519088015260a090960195908201906001016150c7565b600081518084526020808501945080840160005b8381101561473a57815161514688825161468d565b808401516001600160a01b0390811689860152604080830151908a0152606080830151908a0152608080830151908a015260a091820151169088015260c09096019590820190600101615131565b6004811061469d5761469d614677565b600081518084526020808501945080840160005b8381101561473a578151875295820195908201906001016151b8565b6002811061469d5761469d614677565b600082825180855260208086019550808260051b84010181860160005b848110156141dd57601f19868403018952815160a0815185528582015161522a878701826151d4565b506040828101519086015260608083015190860152608091820151918501819052615257818601836151a4565b9a86019a9450505090830190600101615201565b85815260018060a01b038516602082015260a060408201526000610140855160a0808501526152a582850182516001600160a01b03169052565b60208101516101606152c1818701836001600160a01b03169052565b6040830151915080610180870152506152de6102a08601826150b3565b9050606082015161013f19868303016101a08701526152fd828261511d565b91505060808201516153136101c0870182615194565b5060a08201516101e086015260c082015161020086015260e082015161022086015261010080830151610240870152610120808401516102608801528484015161028088015260208a0151945061537560c08801866001600160781b03169052565b60408a01516001600160781b031660e088015260608a0151878403609f19908101848a015290955093506153a98386613f97565b945060808a0151925083878603018188015250506153c78382613f97565b9250505082810360608401526153dd81866151a4565b905082810360808401526132e081856151e4565b6020810161035a82846151d4565b8281526040602082015260006105a56040830184613f9756fea264697066735822122099edcfb1b7b7d00fbcbc358c85ccf9cf79af0017e6c4e391bdadbe0f5b94d25564736f6c634300080d0033",
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

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address offerer) view returns(uint256 nonce)
func (_Consideration *ConsiderationCaller) GetNonce(opts *bind.CallOpts, offerer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Consideration.contract.Call(opts, &out, "getNonce", offerer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address offerer) view returns(uint256 nonce)
func (_Consideration *ConsiderationSession) GetNonce(offerer common.Address) (*big.Int, error) {
	return _Consideration.Contract.GetNonce(&_Consideration.CallOpts, offerer)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address offerer) view returns(uint256 nonce)
func (_Consideration *ConsiderationCallerSession) GetNonce(offerer common.Address) (*big.Int, error) {
	return _Consideration.Contract.GetNonce(&_Consideration.CallOpts, offerer)
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

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xdf7b0dac.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) advancedOrder, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Consideration *ConsiderationTransactor) FulfillAdvancedOrder(opts *bind.TransactOpts, advancedOrder AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Consideration.contract.Transact(opts, "fulfillAdvancedOrder", advancedOrder, criteriaResolvers, fulfillerConduitKey)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xdf7b0dac.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) advancedOrder, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Consideration *ConsiderationSession) FulfillAdvancedOrder(advancedOrder AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Consideration.Contract.FulfillAdvancedOrder(&_Consideration.TransactOpts, advancedOrder, criteriaResolvers, fulfillerConduitKey)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xdf7b0dac.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) advancedOrder, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Consideration *ConsiderationTransactorSession) FulfillAdvancedOrder(advancedOrder AdvancedOrder, criteriaResolvers []CriteriaResolver, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Consideration.Contract.FulfillAdvancedOrder(&_Consideration.TransactOpts, advancedOrder, criteriaResolvers, fulfillerConduitKey)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0xfb4c2af9.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Consideration *ConsiderationTransactor) FulfillAvailableAdvancedOrders(opts *bind.TransactOpts, advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Consideration.contract.Transact(opts, "fulfillAvailableAdvancedOrders", advancedOrders, criteriaResolvers, offerFulfillments, considerationFulfillments, fulfillerConduitKey, maximumFulfilled)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0xfb4c2af9.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Consideration *ConsiderationSession) FulfillAvailableAdvancedOrders(advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Consideration.Contract.FulfillAvailableAdvancedOrders(&_Consideration.TransactOpts, advancedOrders, criteriaResolvers, offerFulfillments, considerationFulfillments, fulfillerConduitKey, maximumFulfilled)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0xfb4c2af9.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] advancedOrders, (uint256,uint8,uint256,uint256,bytes32[])[] criteriaResolvers, (uint256,uint256)[][] offerFulfillments, (uint256,uint256)[][] considerationFulfillments, bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[] availableOrders, ((uint8,address,uint256,uint256,address),address,bytes32)[] executions)
func (_Consideration *ConsiderationTransactorSession) FulfillAvailableAdvancedOrders(advancedOrders []AdvancedOrder, criteriaResolvers []CriteriaResolver, offerFulfillments [][]FulfillmentComponent, considerationFulfillments [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Consideration.Contract.FulfillAvailableAdvancedOrders(&_Consideration.TransactOpts, advancedOrders, criteriaResolvers, offerFulfillments, considerationFulfillments, fulfillerConduitKey, maximumFulfilled)
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

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns(uint256 newNonce)
func (_Consideration *ConsiderationTransactor) IncrementNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Consideration.contract.Transact(opts, "incrementNonce")
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns(uint256 newNonce)
func (_Consideration *ConsiderationSession) IncrementNonce() (*types.Transaction, error) {
	return _Consideration.Contract.IncrementNonce(&_Consideration.TransactOpts)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns(uint256 newNonce)
func (_Consideration *ConsiderationTransactorSession) IncrementNonce() (*types.Transaction, error) {
	return _Consideration.Contract.IncrementNonce(&_Consideration.TransactOpts)
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

// ConsiderationNonceIncrementedIterator is returned from FilterNonceIncremented and is used to iterate over the raw logs and unpacked data for NonceIncremented events raised by the Consideration contract.
type ConsiderationNonceIncrementedIterator struct {
	Event *ConsiderationNonceIncremented // Event containing the contract specifics and raw log

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
func (it *ConsiderationNonceIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsiderationNonceIncremented)
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
		it.Event = new(ConsiderationNonceIncremented)
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
func (it *ConsiderationNonceIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsiderationNonceIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsiderationNonceIncremented represents a NonceIncremented event raised by the Consideration contract.
type ConsiderationNonceIncremented struct {
	NewNonce *big.Int
	Offerer  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceIncremented is a free log retrieval operation binding the contract event 0x7ab0fc7de8910a6100b24df423c3d0835534506dca9473d30c3e7df51241b2cf.
//
// Solidity: event NonceIncremented(uint256 newNonce, address indexed offerer)
func (_Consideration *ConsiderationFilterer) FilterNonceIncremented(opts *bind.FilterOpts, offerer []common.Address) (*ConsiderationNonceIncrementedIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Consideration.contract.FilterLogs(opts, "NonceIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return &ConsiderationNonceIncrementedIterator{contract: _Consideration.contract, event: "NonceIncremented", logs: logs, sub: sub}, nil
}

// WatchNonceIncremented is a free log subscription operation binding the contract event 0x7ab0fc7de8910a6100b24df423c3d0835534506dca9473d30c3e7df51241b2cf.
//
// Solidity: event NonceIncremented(uint256 newNonce, address indexed offerer)
func (_Consideration *ConsiderationFilterer) WatchNonceIncremented(opts *bind.WatchOpts, sink chan<- *ConsiderationNonceIncremented, offerer []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Consideration.contract.WatchLogs(opts, "NonceIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsiderationNonceIncremented)
				if err := _Consideration.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
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
func (_Consideration *ConsiderationFilterer) ParseNonceIncremented(log types.Log) (*ConsiderationNonceIncremented, error) {
	event := new(ConsiderationNonceIncremented)
	if err := _Consideration.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
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
	Fulfiller     common.Address
	Offer         []SpentItem
	Consideration []ReceivedItem
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address fulfiller, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
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
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address fulfiller, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
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
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address fulfiller, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
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
