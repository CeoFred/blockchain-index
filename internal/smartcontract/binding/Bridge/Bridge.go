// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Bridge

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

// BridgeAssistGenericUpgradeableFulfillTx is an auto generated low-level Go binding around an user-defined struct.
type BridgeAssistGenericUpgradeableFulfillTx struct {
	Amount    *big.Int
	FromUser  string
	ToUser    common.Address
	FromChain string
	Nonce     *big.Int
}

// BridgeAssistGenericUpgradeableTransaction is an auto generated low-level Go binding around an user-defined struct.
type BridgeAssistGenericUpgradeableTransaction struct {
	Amount    *big.Int
	Timestamp *big.Int
	FromUser  common.Address
	ToUser    string
	FromChain string
	ToChain   string
	Nonce     *big.Int
	Block     *big.Int
}

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"}],\"name\":\"ChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"}],\"name\":\"ChainRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeFulfill\",\"type\":\"uint256\"}],\"name\":\"FeeFulfillSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeSend\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeFulfill\",\"type\":\"uint256\"}],\"name\":\"FeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeWallet\",\"type\":\"address\"}],\"name\":\"FeeWalletSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"fromUser\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toUser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fromChain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toChain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"}],\"name\":\"FulfilledTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limitPerSend\",\"type\":\"uint256\"}],\"name\":\"LimitPerSendSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RelayerConsensusThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromUser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"toUser\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fromChain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toChain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"}],\"name\":\"SentTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CURRENT_CHAIN\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CURRENT_CHAIN_B32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FULFILL_TX_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_RELAYERS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"chains\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"exchangeRatesFromPow\",\"type\":\"uint256[]\"}],\"name\":\"addChains\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"exchangeRateFrom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeFulfill\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeSend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"fromUser\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"toUser\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"fromChain\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structBridgeAssistGenericUpgradeable.FulfillTx\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"fulfill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fulfilledAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRelayers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserTransactions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromUser\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"toUser\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fromChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"toChain\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"internalType\":\"structBridgeAssistGenericUpgradeable.Transaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserTransactionsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getUserTransactionsSlice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromUser\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"toUser\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fromChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"toChain\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"internalType\":\"structBridgeAssistGenericUpgradeable.Transaction[]\",\"name\":\"transactions_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"limitPerSend_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeWallet_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeSend_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeFulfill_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"relayers_\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"relayerConsensusThreshold_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"}],\"name\":\"isSupportedChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitPerSend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayerConsensusThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"relayers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"chains\",\"type\":\"string[]\"}],\"name\":\"removeChains\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeFulfill_\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeWallet_\",\"type\":\"address\"}],\"name\":\"setFeeWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"setLimitPerSend\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"relayers_\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"relayerConsensusThreshold_\",\"type\":\"uint256\"}],\"name\":\"setRelayers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supportedChainList\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromUser\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"toUser\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fromChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"toChain\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// CURRENTCHAIN is a free data retrieval call binding the contract method 0x7233a666.
//
// Solidity: function CURRENT_CHAIN() view returns(string)
func (_Bridge *BridgeCaller) CURRENTCHAIN(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "CURRENT_CHAIN")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CURRENTCHAIN is a free data retrieval call binding the contract method 0x7233a666.
//
// Solidity: function CURRENT_CHAIN() view returns(string)
func (_Bridge *BridgeSession) CURRENTCHAIN() (string, error) {
	return _Bridge.Contract.CURRENTCHAIN(&_Bridge.CallOpts)
}

// CURRENTCHAIN is a free data retrieval call binding the contract method 0x7233a666.
//
// Solidity: function CURRENT_CHAIN() view returns(string)
func (_Bridge *BridgeCallerSession) CURRENTCHAIN() (string, error) {
	return _Bridge.Contract.CURRENTCHAIN(&_Bridge.CallOpts)
}

// CURRENTCHAINB32 is a free data retrieval call binding the contract method 0xac37b226.
//
// Solidity: function CURRENT_CHAIN_B32() view returns(bytes32)
func (_Bridge *BridgeCaller) CURRENTCHAINB32(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "CURRENT_CHAIN_B32")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CURRENTCHAINB32 is a free data retrieval call binding the contract method 0xac37b226.
//
// Solidity: function CURRENT_CHAIN_B32() view returns(bytes32)
func (_Bridge *BridgeSession) CURRENTCHAINB32() ([32]byte, error) {
	return _Bridge.Contract.CURRENTCHAINB32(&_Bridge.CallOpts)
}

// CURRENTCHAINB32 is a free data retrieval call binding the contract method 0xac37b226.
//
// Solidity: function CURRENT_CHAIN_B32() view returns(bytes32)
func (_Bridge *BridgeCallerSession) CURRENTCHAINB32() ([32]byte, error) {
	return _Bridge.Contract.CURRENTCHAINB32(&_Bridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.DEFAULTADMINROLE(&_Bridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.DEFAULTADMINROLE(&_Bridge.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Bridge *BridgeCaller) FEEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "FEE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Bridge *BridgeSession) FEEDENOMINATOR() (*big.Int, error) {
	return _Bridge.Contract.FEEDENOMINATOR(&_Bridge.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Bridge *BridgeCallerSession) FEEDENOMINATOR() (*big.Int, error) {
	return _Bridge.Contract.FEEDENOMINATOR(&_Bridge.CallOpts)
}

// FULFILLTXTYPEHASH is a free data retrieval call binding the contract method 0x41655326.
//
// Solidity: function FULFILL_TX_TYPEHASH() view returns(bytes32)
func (_Bridge *BridgeCaller) FULFILLTXTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "FULFILL_TX_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FULFILLTXTYPEHASH is a free data retrieval call binding the contract method 0x41655326.
//
// Solidity: function FULFILL_TX_TYPEHASH() view returns(bytes32)
func (_Bridge *BridgeSession) FULFILLTXTYPEHASH() ([32]byte, error) {
	return _Bridge.Contract.FULFILLTXTYPEHASH(&_Bridge.CallOpts)
}

// FULFILLTXTYPEHASH is a free data retrieval call binding the contract method 0x41655326.
//
// Solidity: function FULFILL_TX_TYPEHASH() view returns(bytes32)
func (_Bridge *BridgeCallerSession) FULFILLTXTYPEHASH() ([32]byte, error) {
	return _Bridge.Contract.FULFILLTXTYPEHASH(&_Bridge.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) MANAGERROLE() ([32]byte, error) {
	return _Bridge.Contract.MANAGERROLE(&_Bridge.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) MANAGERROLE() ([32]byte, error) {
	return _Bridge.Contract.MANAGERROLE(&_Bridge.CallOpts)
}

// MAXRELAYERS is a free data retrieval call binding the contract method 0x9debb3bd.
//
// Solidity: function MAX_RELAYERS() view returns(uint256)
func (_Bridge *BridgeCaller) MAXRELAYERS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "MAX_RELAYERS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXRELAYERS is a free data retrieval call binding the contract method 0x9debb3bd.
//
// Solidity: function MAX_RELAYERS() view returns(uint256)
func (_Bridge *BridgeSession) MAXRELAYERS() (*big.Int, error) {
	return _Bridge.Contract.MAXRELAYERS(&_Bridge.CallOpts)
}

// MAXRELAYERS is a free data retrieval call binding the contract method 0x9debb3bd.
//
// Solidity: function MAX_RELAYERS() view returns(uint256)
func (_Bridge *BridgeCallerSession) MAXRELAYERS() (*big.Int, error) {
	return _Bridge.Contract.MAXRELAYERS(&_Bridge.CallOpts)
}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_Bridge *BridgeCaller) TOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_Bridge *BridgeSession) TOKEN() (common.Address, error) {
	return _Bridge.Contract.TOKEN(&_Bridge.CallOpts)
}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_Bridge *BridgeCallerSession) TOKEN() (common.Address, error) {
	return _Bridge.Contract.TOKEN(&_Bridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Bridge *BridgeCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Bridge *BridgeSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Bridge.Contract.Eip712Domain(&_Bridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Bridge *BridgeCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Bridge.Contract.Eip712Domain(&_Bridge.CallOpts)
}

// ExchangeRateFrom is a free data retrieval call binding the contract method 0x97901c5a.
//
// Solidity: function exchangeRateFrom(bytes32 ) view returns(uint256)
func (_Bridge *BridgeCaller) ExchangeRateFrom(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "exchangeRateFrom", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateFrom is a free data retrieval call binding the contract method 0x97901c5a.
//
// Solidity: function exchangeRateFrom(bytes32 ) view returns(uint256)
func (_Bridge *BridgeSession) ExchangeRateFrom(arg0 [32]byte) (*big.Int, error) {
	return _Bridge.Contract.ExchangeRateFrom(&_Bridge.CallOpts, arg0)
}

// ExchangeRateFrom is a free data retrieval call binding the contract method 0x97901c5a.
//
// Solidity: function exchangeRateFrom(bytes32 ) view returns(uint256)
func (_Bridge *BridgeCallerSession) ExchangeRateFrom(arg0 [32]byte) (*big.Int, error) {
	return _Bridge.Contract.ExchangeRateFrom(&_Bridge.CallOpts, arg0)
}

// FeeFulfill is a free data retrieval call binding the contract method 0xb46c31a4.
//
// Solidity: function feeFulfill() view returns(uint256)
func (_Bridge *BridgeCaller) FeeFulfill(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "feeFulfill")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeFulfill is a free data retrieval call binding the contract method 0xb46c31a4.
//
// Solidity: function feeFulfill() view returns(uint256)
func (_Bridge *BridgeSession) FeeFulfill() (*big.Int, error) {
	return _Bridge.Contract.FeeFulfill(&_Bridge.CallOpts)
}

// FeeFulfill is a free data retrieval call binding the contract method 0xb46c31a4.
//
// Solidity: function feeFulfill() view returns(uint256)
func (_Bridge *BridgeCallerSession) FeeFulfill() (*big.Int, error) {
	return _Bridge.Contract.FeeFulfill(&_Bridge.CallOpts)
}

// FeeSend is a free data retrieval call binding the contract method 0xcfae7307.
//
// Solidity: function feeSend() view returns(uint256)
func (_Bridge *BridgeCaller) FeeSend(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "feeSend")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeSend is a free data retrieval call binding the contract method 0xcfae7307.
//
// Solidity: function feeSend() view returns(uint256)
func (_Bridge *BridgeSession) FeeSend() (*big.Int, error) {
	return _Bridge.Contract.FeeSend(&_Bridge.CallOpts)
}

// FeeSend is a free data retrieval call binding the contract method 0xcfae7307.
//
// Solidity: function feeSend() view returns(uint256)
func (_Bridge *BridgeCallerSession) FeeSend() (*big.Int, error) {
	return _Bridge.Contract.FeeSend(&_Bridge.CallOpts)
}

// FeeWallet is a free data retrieval call binding the contract method 0xf25f4b56.
//
// Solidity: function feeWallet() view returns(address)
func (_Bridge *BridgeCaller) FeeWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "feeWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeWallet is a free data retrieval call binding the contract method 0xf25f4b56.
//
// Solidity: function feeWallet() view returns(address)
func (_Bridge *BridgeSession) FeeWallet() (common.Address, error) {
	return _Bridge.Contract.FeeWallet(&_Bridge.CallOpts)
}

// FeeWallet is a free data retrieval call binding the contract method 0xf25f4b56.
//
// Solidity: function feeWallet() view returns(address)
func (_Bridge *BridgeCallerSession) FeeWallet() (common.Address, error) {
	return _Bridge.Contract.FeeWallet(&_Bridge.CallOpts)
}

// FulfilledAt is a free data retrieval call binding the contract method 0xf7b2bf68.
//
// Solidity: function fulfilledAt(string , string , uint256 ) view returns(uint256)
func (_Bridge *BridgeCaller) FulfilledAt(opts *bind.CallOpts, arg0 string, arg1 string, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "fulfilledAt", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FulfilledAt is a free data retrieval call binding the contract method 0xf7b2bf68.
//
// Solidity: function fulfilledAt(string , string , uint256 ) view returns(uint256)
func (_Bridge *BridgeSession) FulfilledAt(arg0 string, arg1 string, arg2 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.FulfilledAt(&_Bridge.CallOpts, arg0, arg1, arg2)
}

// FulfilledAt is a free data retrieval call binding the contract method 0xf7b2bf68.
//
// Solidity: function fulfilledAt(string , string , uint256 ) view returns(uint256)
func (_Bridge *BridgeCallerSession) FulfilledAt(arg0 string, arg1 string, arg2 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.FulfilledAt(&_Bridge.CallOpts, arg0, arg1, arg2)
}

// GetRelayers is a free data retrieval call binding the contract method 0x179ff4b2.
//
// Solidity: function getRelayers() view returns(address[])
func (_Bridge *BridgeCaller) GetRelayers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRelayers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRelayers is a free data retrieval call binding the contract method 0x179ff4b2.
//
// Solidity: function getRelayers() view returns(address[])
func (_Bridge *BridgeSession) GetRelayers() ([]common.Address, error) {
	return _Bridge.Contract.GetRelayers(&_Bridge.CallOpts)
}

// GetRelayers is a free data retrieval call binding the contract method 0x179ff4b2.
//
// Solidity: function getRelayers() view returns(address[])
func (_Bridge *BridgeCallerSession) GetRelayers() ([]common.Address, error) {
	return _Bridge.Contract.GetRelayers(&_Bridge.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridge.Contract.GetRoleAdmin(&_Bridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridge.Contract.GetRoleAdmin(&_Bridge.CallOpts, role)
}

// GetUserTransactions is a free data retrieval call binding the contract method 0xef925399.
//
// Solidity: function getUserTransactions(address user) view returns((uint256,uint256,address,string,string,string,uint256,uint256)[])
func (_Bridge *BridgeCaller) GetUserTransactions(opts *bind.CallOpts, user common.Address) ([]BridgeAssistGenericUpgradeableTransaction, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getUserTransactions", user)

	if err != nil {
		return *new([]BridgeAssistGenericUpgradeableTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new([]BridgeAssistGenericUpgradeableTransaction)).(*[]BridgeAssistGenericUpgradeableTransaction)

	return out0, err

}

// GetUserTransactions is a free data retrieval call binding the contract method 0xef925399.
//
// Solidity: function getUserTransactions(address user) view returns((uint256,uint256,address,string,string,string,uint256,uint256)[])
func (_Bridge *BridgeSession) GetUserTransactions(user common.Address) ([]BridgeAssistGenericUpgradeableTransaction, error) {
	return _Bridge.Contract.GetUserTransactions(&_Bridge.CallOpts, user)
}

// GetUserTransactions is a free data retrieval call binding the contract method 0xef925399.
//
// Solidity: function getUserTransactions(address user) view returns((uint256,uint256,address,string,string,string,uint256,uint256)[])
func (_Bridge *BridgeCallerSession) GetUserTransactions(user common.Address) ([]BridgeAssistGenericUpgradeableTransaction, error) {
	return _Bridge.Contract.GetUserTransactions(&_Bridge.CallOpts, user)
}

// GetUserTransactionsAmount is a free data retrieval call binding the contract method 0x2cf26701.
//
// Solidity: function getUserTransactionsAmount(address user) view returns(uint256)
func (_Bridge *BridgeCaller) GetUserTransactionsAmount(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getUserTransactionsAmount", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserTransactionsAmount is a free data retrieval call binding the contract method 0x2cf26701.
//
// Solidity: function getUserTransactionsAmount(address user) view returns(uint256)
func (_Bridge *BridgeSession) GetUserTransactionsAmount(user common.Address) (*big.Int, error) {
	return _Bridge.Contract.GetUserTransactionsAmount(&_Bridge.CallOpts, user)
}

// GetUserTransactionsAmount is a free data retrieval call binding the contract method 0x2cf26701.
//
// Solidity: function getUserTransactionsAmount(address user) view returns(uint256)
func (_Bridge *BridgeCallerSession) GetUserTransactionsAmount(user common.Address) (*big.Int, error) {
	return _Bridge.Contract.GetUserTransactionsAmount(&_Bridge.CallOpts, user)
}

// GetUserTransactionsSlice is a free data retrieval call binding the contract method 0x5f282ba4.
//
// Solidity: function getUserTransactionsSlice(address user_, uint256 offset_, uint256 limit_) view returns((uint256,uint256,address,string,string,string,uint256,uint256)[] transactions_)
func (_Bridge *BridgeCaller) GetUserTransactionsSlice(opts *bind.CallOpts, user_ common.Address, offset_ *big.Int, limit_ *big.Int) ([]BridgeAssistGenericUpgradeableTransaction, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getUserTransactionsSlice", user_, offset_, limit_)

	if err != nil {
		return *new([]BridgeAssistGenericUpgradeableTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new([]BridgeAssistGenericUpgradeableTransaction)).(*[]BridgeAssistGenericUpgradeableTransaction)

	return out0, err

}

// GetUserTransactionsSlice is a free data retrieval call binding the contract method 0x5f282ba4.
//
// Solidity: function getUserTransactionsSlice(address user_, uint256 offset_, uint256 limit_) view returns((uint256,uint256,address,string,string,string,uint256,uint256)[] transactions_)
func (_Bridge *BridgeSession) GetUserTransactionsSlice(user_ common.Address, offset_ *big.Int, limit_ *big.Int) ([]BridgeAssistGenericUpgradeableTransaction, error) {
	return _Bridge.Contract.GetUserTransactionsSlice(&_Bridge.CallOpts, user_, offset_, limit_)
}

// GetUserTransactionsSlice is a free data retrieval call binding the contract method 0x5f282ba4.
//
// Solidity: function getUserTransactionsSlice(address user_, uint256 offset_, uint256 limit_) view returns((uint256,uint256,address,string,string,string,uint256,uint256)[] transactions_)
func (_Bridge *BridgeCallerSession) GetUserTransactionsSlice(user_ common.Address, offset_ *big.Int, limit_ *big.Int) ([]BridgeAssistGenericUpgradeableTransaction, error) {
	return _Bridge.Contract.GetUserTransactionsSlice(&_Bridge.CallOpts, user_, offset_, limit_)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridge.Contract.HasRole(&_Bridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridge.Contract.HasRole(&_Bridge.CallOpts, role, account)
}

// IsSupportedChain is a free data retrieval call binding the contract method 0xb049cec7.
//
// Solidity: function isSupportedChain(string chain) view returns(bool)
func (_Bridge *BridgeCaller) IsSupportedChain(opts *bind.CallOpts, chain string) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isSupportedChain", chain)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedChain is a free data retrieval call binding the contract method 0xb049cec7.
//
// Solidity: function isSupportedChain(string chain) view returns(bool)
func (_Bridge *BridgeSession) IsSupportedChain(chain string) (bool, error) {
	return _Bridge.Contract.IsSupportedChain(&_Bridge.CallOpts, chain)
}

// IsSupportedChain is a free data retrieval call binding the contract method 0xb049cec7.
//
// Solidity: function isSupportedChain(string chain) view returns(bool)
func (_Bridge *BridgeCallerSession) IsSupportedChain(chain string) (bool, error) {
	return _Bridge.Contract.IsSupportedChain(&_Bridge.CallOpts, chain)
}

// LimitPerSend is a free data retrieval call binding the contract method 0xd3649d6c.
//
// Solidity: function limitPerSend() view returns(uint256)
func (_Bridge *BridgeCaller) LimitPerSend(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "limitPerSend")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitPerSend is a free data retrieval call binding the contract method 0xd3649d6c.
//
// Solidity: function limitPerSend() view returns(uint256)
func (_Bridge *BridgeSession) LimitPerSend() (*big.Int, error) {
	return _Bridge.Contract.LimitPerSend(&_Bridge.CallOpts)
}

// LimitPerSend is a free data retrieval call binding the contract method 0xd3649d6c.
//
// Solidity: function limitPerSend() view returns(uint256)
func (_Bridge *BridgeCallerSession) LimitPerSend() (*big.Int, error) {
	return _Bridge.Contract.LimitPerSend(&_Bridge.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Bridge *BridgeCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Bridge *BridgeSession) Nonce() (*big.Int, error) {
	return _Bridge.Contract.Nonce(&_Bridge.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Bridge *BridgeCallerSession) Nonce() (*big.Int, error) {
	return _Bridge.Contract.Nonce(&_Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCallerSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// RelayerConsensusThreshold is a free data retrieval call binding the contract method 0xfe62c28b.
//
// Solidity: function relayerConsensusThreshold() view returns(uint256)
func (_Bridge *BridgeCaller) RelayerConsensusThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "relayerConsensusThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayerConsensusThreshold is a free data retrieval call binding the contract method 0xfe62c28b.
//
// Solidity: function relayerConsensusThreshold() view returns(uint256)
func (_Bridge *BridgeSession) RelayerConsensusThreshold() (*big.Int, error) {
	return _Bridge.Contract.RelayerConsensusThreshold(&_Bridge.CallOpts)
}

// RelayerConsensusThreshold is a free data retrieval call binding the contract method 0xfe62c28b.
//
// Solidity: function relayerConsensusThreshold() view returns(uint256)
func (_Bridge *BridgeCallerSession) RelayerConsensusThreshold() (*big.Int, error) {
	return _Bridge.Contract.RelayerConsensusThreshold(&_Bridge.CallOpts)
}

// Relayers is a free data retrieval call binding the contract method 0x9a48e7f9.
//
// Solidity: function relayers(uint256 ) view returns(address)
func (_Bridge *BridgeCaller) Relayers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "relayers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relayers is a free data retrieval call binding the contract method 0x9a48e7f9.
//
// Solidity: function relayers(uint256 ) view returns(address)
func (_Bridge *BridgeSession) Relayers(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.Relayers(&_Bridge.CallOpts, arg0)
}

// Relayers is a free data retrieval call binding the contract method 0x9a48e7f9.
//
// Solidity: function relayers(uint256 ) view returns(address)
func (_Bridge *BridgeCallerSession) Relayers(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.Relayers(&_Bridge.CallOpts, arg0)
}

// RelayersLength is a free data retrieval call binding the contract method 0xae7cabbd.
//
// Solidity: function relayersLength() view returns(uint256)
func (_Bridge *BridgeCaller) RelayersLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "relayersLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayersLength is a free data retrieval call binding the contract method 0xae7cabbd.
//
// Solidity: function relayersLength() view returns(uint256)
func (_Bridge *BridgeSession) RelayersLength() (*big.Int, error) {
	return _Bridge.Contract.RelayersLength(&_Bridge.CallOpts)
}

// RelayersLength is a free data retrieval call binding the contract method 0xae7cabbd.
//
// Solidity: function relayersLength() view returns(uint256)
func (_Bridge *BridgeCallerSession) RelayersLength() (*big.Int, error) {
	return _Bridge.Contract.RelayersLength(&_Bridge.CallOpts)
}

// Send is a free data retrieval call binding the contract method 0xe0291592.
//
// Solidity: function send(uint256 , string , string ) pure returns()
func (_Bridge *BridgeCaller) Send(opts *bind.CallOpts, arg0 *big.Int, arg1 string, arg2 string) error {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "send", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// Send is a free data retrieval call binding the contract method 0xe0291592.
//
// Solidity: function send(uint256 , string , string ) pure returns()
func (_Bridge *BridgeSession) Send(arg0 *big.Int, arg1 string, arg2 string) error {
	return _Bridge.Contract.Send(&_Bridge.CallOpts, arg0, arg1, arg2)
}

// Send is a free data retrieval call binding the contract method 0xe0291592.
//
// Solidity: function send(uint256 , string , string ) pure returns()
func (_Bridge *BridgeCallerSession) Send(arg0 *big.Int, arg1 string, arg2 string) error {
	return _Bridge.Contract.Send(&_Bridge.CallOpts, arg0, arg1, arg2)
}

// SetFee is a free data retrieval call binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 , uint256 ) pure returns()
func (_Bridge *BridgeCaller) SetFee(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) error {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "setFee", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// SetFee is a free data retrieval call binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 , uint256 ) pure returns()
func (_Bridge *BridgeSession) SetFee(arg0 *big.Int, arg1 *big.Int) error {
	return _Bridge.Contract.SetFee(&_Bridge.CallOpts, arg0, arg1)
}

// SetFee is a free data retrieval call binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 , uint256 ) pure returns()
func (_Bridge *BridgeCallerSession) SetFee(arg0 *big.Int, arg1 *big.Int) error {
	return _Bridge.Contract.SetFee(&_Bridge.CallOpts, arg0, arg1)
}

// SetLimitPerSend is a free data retrieval call binding the contract method 0xa427b242.
//
// Solidity: function setLimitPerSend(uint256 ) pure returns()
func (_Bridge *BridgeCaller) SetLimitPerSend(opts *bind.CallOpts, arg0 *big.Int) error {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "setLimitPerSend", arg0)

	if err != nil {
		return err
	}

	return err

}

// SetLimitPerSend is a free data retrieval call binding the contract method 0xa427b242.
//
// Solidity: function setLimitPerSend(uint256 ) pure returns()
func (_Bridge *BridgeSession) SetLimitPerSend(arg0 *big.Int) error {
	return _Bridge.Contract.SetLimitPerSend(&_Bridge.CallOpts, arg0)
}

// SetLimitPerSend is a free data retrieval call binding the contract method 0xa427b242.
//
// Solidity: function setLimitPerSend(uint256 ) pure returns()
func (_Bridge *BridgeCallerSession) SetLimitPerSend(arg0 *big.Int) error {
	return _Bridge.Contract.SetLimitPerSend(&_Bridge.CallOpts, arg0)
}

// SupportedChainList is a free data retrieval call binding the contract method 0xe026faa7.
//
// Solidity: function supportedChainList() view returns(bytes32[])
func (_Bridge *BridgeCaller) SupportedChainList(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "supportedChainList")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// SupportedChainList is a free data retrieval call binding the contract method 0xe026faa7.
//
// Solidity: function supportedChainList() view returns(bytes32[])
func (_Bridge *BridgeSession) SupportedChainList() ([][32]byte, error) {
	return _Bridge.Contract.SupportedChainList(&_Bridge.CallOpts)
}

// SupportedChainList is a free data retrieval call binding the contract method 0xe026faa7.
//
// Solidity: function supportedChainList() view returns(bytes32[])
func (_Bridge *BridgeCallerSession) SupportedChainList() ([][32]byte, error) {
	return _Bridge.Contract.SupportedChainList(&_Bridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// Transactions is a free data retrieval call binding the contract method 0x14538128.
//
// Solidity: function transactions(address , uint256 ) view returns(uint256 amount, uint256 timestamp, address fromUser, string toUser, string fromChain, string toChain, uint256 nonce, uint256 block)
func (_Bridge *BridgeCaller) Transactions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Amount    *big.Int
	Timestamp *big.Int
	FromUser  common.Address
	ToUser    string
	FromChain string
	ToChain   string
	Nonce     *big.Int
	Block     *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "transactions", arg0, arg1)

	outstruct := new(struct {
		Amount    *big.Int
		Timestamp *big.Int
		FromUser  common.Address
		ToUser    string
		FromChain string
		ToChain   string
		Nonce     *big.Int
		Block     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FromUser = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ToUser = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.FromChain = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.ToChain = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Nonce = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x14538128.
//
// Solidity: function transactions(address , uint256 ) view returns(uint256 amount, uint256 timestamp, address fromUser, string toUser, string fromChain, string toChain, uint256 nonce, uint256 block)
func (_Bridge *BridgeSession) Transactions(arg0 common.Address, arg1 *big.Int) (struct {
	Amount    *big.Int
	Timestamp *big.Int
	FromUser  common.Address
	ToUser    string
	FromChain string
	ToChain   string
	Nonce     *big.Int
	Block     *big.Int
}, error) {
	return _Bridge.Contract.Transactions(&_Bridge.CallOpts, arg0, arg1)
}

// Transactions is a free data retrieval call binding the contract method 0x14538128.
//
// Solidity: function transactions(address , uint256 ) view returns(uint256 amount, uint256 timestamp, address fromUser, string toUser, string fromChain, string toChain, uint256 nonce, uint256 block)
func (_Bridge *BridgeCallerSession) Transactions(arg0 common.Address, arg1 *big.Int) (struct {
	Amount    *big.Int
	Timestamp *big.Int
	FromUser  common.Address
	ToUser    string
	FromChain string
	ToChain   string
	Nonce     *big.Int
	Block     *big.Int
}, error) {
	return _Bridge.Contract.Transactions(&_Bridge.CallOpts, arg0, arg1)
}

// AddChains is a paid mutator transaction binding the contract method 0x00f54e80.
//
// Solidity: function addChains(string[] chains, uint256[] exchangeRatesFromPow) returns()
func (_Bridge *BridgeTransactor) AddChains(opts *bind.TransactOpts, chains []string, exchangeRatesFromPow []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addChains", chains, exchangeRatesFromPow)
}

// AddChains is a paid mutator transaction binding the contract method 0x00f54e80.
//
// Solidity: function addChains(string[] chains, uint256[] exchangeRatesFromPow) returns()
func (_Bridge *BridgeSession) AddChains(chains []string, exchangeRatesFromPow []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddChains(&_Bridge.TransactOpts, chains, exchangeRatesFromPow)
}

// AddChains is a paid mutator transaction binding the contract method 0x00f54e80.
//
// Solidity: function addChains(string[] chains, uint256[] exchangeRatesFromPow) returns()
func (_Bridge *BridgeTransactorSession) AddChains(chains []string, exchangeRatesFromPow []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddChains(&_Bridge.TransactOpts, chains, exchangeRatesFromPow)
}

// Fulfill is a paid mutator transaction binding the contract method 0x984bd4b2.
//
// Solidity: function fulfill((uint256,string,address,string,uint256) transaction, bytes[] signatures) returns()
func (_Bridge *BridgeTransactor) Fulfill(opts *bind.TransactOpts, transaction BridgeAssistGenericUpgradeableFulfillTx, signatures [][]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "fulfill", transaction, signatures)
}

// Fulfill is a paid mutator transaction binding the contract method 0x984bd4b2.
//
// Solidity: function fulfill((uint256,string,address,string,uint256) transaction, bytes[] signatures) returns()
func (_Bridge *BridgeSession) Fulfill(transaction BridgeAssistGenericUpgradeableFulfillTx, signatures [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.Fulfill(&_Bridge.TransactOpts, transaction, signatures)
}

// Fulfill is a paid mutator transaction binding the contract method 0x984bd4b2.
//
// Solidity: function fulfill((uint256,string,address,string,uint256) transaction, bytes[] signatures) returns()
func (_Bridge *BridgeTransactorSession) Fulfill(transaction BridgeAssistGenericUpgradeableFulfillTx, signatures [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.Fulfill(&_Bridge.TransactOpts, transaction, signatures)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.GrantRole(&_Bridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.GrantRole(&_Bridge.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x680d82ea.
//
// Solidity: function initialize(address token_, uint256 limitPerSend_, address feeWallet_, uint256 feeSend_, uint256 feeFulfill_, address owner, address[] relayers_, uint256 relayerConsensusThreshold_) returns()
func (_Bridge *BridgeTransactor) Initialize(opts *bind.TransactOpts, token_ common.Address, limitPerSend_ *big.Int, feeWallet_ common.Address, feeSend_ *big.Int, feeFulfill_ *big.Int, owner common.Address, relayers_ []common.Address, relayerConsensusThreshold_ *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "initialize", token_, limitPerSend_, feeWallet_, feeSend_, feeFulfill_, owner, relayers_, relayerConsensusThreshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0x680d82ea.
//
// Solidity: function initialize(address token_, uint256 limitPerSend_, address feeWallet_, uint256 feeSend_, uint256 feeFulfill_, address owner, address[] relayers_, uint256 relayerConsensusThreshold_) returns()
func (_Bridge *BridgeSession) Initialize(token_ common.Address, limitPerSend_ *big.Int, feeWallet_ common.Address, feeSend_ *big.Int, feeFulfill_ *big.Int, owner common.Address, relayers_ []common.Address, relayerConsensusThreshold_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, token_, limitPerSend_, feeWallet_, feeSend_, feeFulfill_, owner, relayers_, relayerConsensusThreshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0x680d82ea.
//
// Solidity: function initialize(address token_, uint256 limitPerSend_, address feeWallet_, uint256 feeSend_, uint256 feeFulfill_, address owner, address[] relayers_, uint256 relayerConsensusThreshold_) returns()
func (_Bridge *BridgeTransactorSession) Initialize(token_ common.Address, limitPerSend_ *big.Int, feeWallet_ common.Address, feeSend_ *big.Int, feeFulfill_ *big.Int, owner common.Address, relayers_ []common.Address, relayerConsensusThreshold_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, token_, limitPerSend_, feeWallet_, feeSend_, feeFulfill_, owner, relayers_, relayerConsensusThreshold_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeSession) Pause() (*types.Transaction, error) {
	return _Bridge.Contract.Pause(&_Bridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _Bridge.Contract.Pause(&_Bridge.TransactOpts)
}

// RemoveChains is a paid mutator transaction binding the contract method 0x1e12ef29.
//
// Solidity: function removeChains(string[] chains) returns()
func (_Bridge *BridgeTransactor) RemoveChains(opts *bind.TransactOpts, chains []string) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "removeChains", chains)
}

// RemoveChains is a paid mutator transaction binding the contract method 0x1e12ef29.
//
// Solidity: function removeChains(string[] chains) returns()
func (_Bridge *BridgeSession) RemoveChains(chains []string) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveChains(&_Bridge.TransactOpts, chains)
}

// RemoveChains is a paid mutator transaction binding the contract method 0x1e12ef29.
//
// Solidity: function removeChains(string[] chains) returns()
func (_Bridge *BridgeTransactorSession) RemoveChains(chains []string) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveChains(&_Bridge.TransactOpts, chains)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceRole(&_Bridge.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceRole(&_Bridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRole(&_Bridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRole(&_Bridge.TransactOpts, role, account)
}

// SetFee0 is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 feeFulfill_) returns()
func (_Bridge *BridgeTransactor) SetFee0(opts *bind.TransactOpts, feeFulfill_ *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setFee0", feeFulfill_)
}

// SetFee0 is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 feeFulfill_) returns()
func (_Bridge *BridgeSession) SetFee0(feeFulfill_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetFee0(&_Bridge.TransactOpts, feeFulfill_)
}

// SetFee0 is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 feeFulfill_) returns()
func (_Bridge *BridgeTransactorSession) SetFee0(feeFulfill_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetFee0(&_Bridge.TransactOpts, feeFulfill_)
}

// SetFeeWallet is a paid mutator transaction binding the contract method 0x90d49b9d.
//
// Solidity: function setFeeWallet(address feeWallet_) returns()
func (_Bridge *BridgeTransactor) SetFeeWallet(opts *bind.TransactOpts, feeWallet_ common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setFeeWallet", feeWallet_)
}

// SetFeeWallet is a paid mutator transaction binding the contract method 0x90d49b9d.
//
// Solidity: function setFeeWallet(address feeWallet_) returns()
func (_Bridge *BridgeSession) SetFeeWallet(feeWallet_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetFeeWallet(&_Bridge.TransactOpts, feeWallet_)
}

// SetFeeWallet is a paid mutator transaction binding the contract method 0x90d49b9d.
//
// Solidity: function setFeeWallet(address feeWallet_) returns()
func (_Bridge *BridgeTransactorSession) SetFeeWallet(feeWallet_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetFeeWallet(&_Bridge.TransactOpts, feeWallet_)
}

// SetRelayers is a paid mutator transaction binding the contract method 0x2aaf5ed5.
//
// Solidity: function setRelayers(address[] relayers_, uint256 relayerConsensusThreshold_) returns()
func (_Bridge *BridgeTransactor) SetRelayers(opts *bind.TransactOpts, relayers_ []common.Address, relayerConsensusThreshold_ *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setRelayers", relayers_, relayerConsensusThreshold_)
}

// SetRelayers is a paid mutator transaction binding the contract method 0x2aaf5ed5.
//
// Solidity: function setRelayers(address[] relayers_, uint256 relayerConsensusThreshold_) returns()
func (_Bridge *BridgeSession) SetRelayers(relayers_ []common.Address, relayerConsensusThreshold_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetRelayers(&_Bridge.TransactOpts, relayers_, relayerConsensusThreshold_)
}

// SetRelayers is a paid mutator transaction binding the contract method 0x2aaf5ed5.
//
// Solidity: function setRelayers(address[] relayers_, uint256 relayerConsensusThreshold_) returns()
func (_Bridge *BridgeTransactorSession) SetRelayers(relayers_ []common.Address, relayerConsensusThreshold_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetRelayers(&_Bridge.TransactOpts, relayers_, relayerConsensusThreshold_)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeSession) Unpause() (*types.Transaction, error) {
	return _Bridge.Contract.Unpause(&_Bridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _Bridge.Contract.Unpause(&_Bridge.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token_, address to, uint256 amount) returns()
func (_Bridge *BridgeTransactor) Withdraw(opts *bind.TransactOpts, token_ common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdraw", token_, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token_, address to, uint256 amount) returns()
func (_Bridge *BridgeSession) Withdraw(token_ common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, token_, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token_, address to, uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) Withdraw(token_ common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, token_, to, amount)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address to, uint256 amount) returns()
func (_Bridge *BridgeTransactor) WithdrawNative(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawNative", to, amount)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address to, uint256 amount) returns()
func (_Bridge *BridgeSession) WithdrawNative(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawNative(&_Bridge.TransactOpts, to, amount)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address to, uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) WithdrawNative(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawNative(&_Bridge.TransactOpts, to, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// BridgeChainAddedIterator is returned from FilterChainAdded and is used to iterate over the raw logs and unpacked data for ChainAdded events raised by the Bridge contract.
type BridgeChainAddedIterator struct {
	Event *BridgeChainAdded // Event containing the contract specifics and raw log

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
func (it *BridgeChainAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeChainAdded)
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
		it.Event = new(BridgeChainAdded)
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
func (it *BridgeChainAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeChainAdded represents a ChainAdded event raised by the Bridge contract.
type BridgeChainAdded struct {
	Chain        common.Hash
	ExchangeRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterChainAdded is a free log retrieval operation binding the contract event 0x8398b0e4d6d15e6038c42e15bf0f0f5466dc36dfd49979c0599bc2eb0fb58302.
//
// Solidity: event ChainAdded(string indexed chain, uint256 exchangeRate)
func (_Bridge *BridgeFilterer) FilterChainAdded(opts *bind.FilterOpts, chain []string) (*BridgeChainAddedIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ChainAdded", chainRule)
	if err != nil {
		return nil, err
	}
	return &BridgeChainAddedIterator{contract: _Bridge.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

// WatchChainAdded is a free log subscription operation binding the contract event 0x8398b0e4d6d15e6038c42e15bf0f0f5466dc36dfd49979c0599bc2eb0fb58302.
//
// Solidity: event ChainAdded(string indexed chain, uint256 exchangeRate)
func (_Bridge *BridgeFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BridgeChainAdded, chain []string) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ChainAdded", chainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeChainAdded)
				if err := _Bridge.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

// ParseChainAdded is a log parse operation binding the contract event 0x8398b0e4d6d15e6038c42e15bf0f0f5466dc36dfd49979c0599bc2eb0fb58302.
//
// Solidity: event ChainAdded(string indexed chain, uint256 exchangeRate)
func (_Bridge *BridgeFilterer) ParseChainAdded(log types.Log) (*BridgeChainAdded, error) {
	event := new(BridgeChainAdded)
	if err := _Bridge.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeChainRemovedIterator is returned from FilterChainRemoved and is used to iterate over the raw logs and unpacked data for ChainRemoved events raised by the Bridge contract.
type BridgeChainRemovedIterator struct {
	Event *BridgeChainRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeChainRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeChainRemoved)
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
		it.Event = new(BridgeChainRemoved)
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
func (it *BridgeChainRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeChainRemoved represents a ChainRemoved event raised by the Bridge contract.
type BridgeChainRemoved struct {
	Chain common.Hash
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChainRemoved is a free log retrieval operation binding the contract event 0x9872686700ce3c111f0231d65c591f43d9592b5ef96649b3f7ee98c4133dbc62.
//
// Solidity: event ChainRemoved(string indexed chain)
func (_Bridge *BridgeFilterer) FilterChainRemoved(opts *bind.FilterOpts, chain []string) (*BridgeChainRemovedIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ChainRemoved", chainRule)
	if err != nil {
		return nil, err
	}
	return &BridgeChainRemovedIterator{contract: _Bridge.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

// WatchChainRemoved is a free log subscription operation binding the contract event 0x9872686700ce3c111f0231d65c591f43d9592b5ef96649b3f7ee98c4133dbc62.
//
// Solidity: event ChainRemoved(string indexed chain)
func (_Bridge *BridgeFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BridgeChainRemoved, chain []string) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ChainRemoved", chainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeChainRemoved)
				if err := _Bridge.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

// ParseChainRemoved is a log parse operation binding the contract event 0x9872686700ce3c111f0231d65c591f43d9592b5ef96649b3f7ee98c4133dbc62.
//
// Solidity: event ChainRemoved(string indexed chain)
func (_Bridge *BridgeFilterer) ParseChainRemoved(log types.Log) (*BridgeChainRemoved, error) {
	event := new(BridgeChainRemoved)
	if err := _Bridge.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Bridge contract.
type BridgeEIP712DomainChangedIterator struct {
	Event *BridgeEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *BridgeEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEIP712DomainChanged)
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
		it.Event = new(BridgeEIP712DomainChanged)
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
func (it *BridgeEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEIP712DomainChanged represents a EIP712DomainChanged event raised by the Bridge contract.
type BridgeEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Bridge *BridgeFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BridgeEIP712DomainChangedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BridgeEIP712DomainChangedIterator{contract: _Bridge.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Bridge *BridgeFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BridgeEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEIP712DomainChanged)
				if err := _Bridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Bridge *BridgeFilterer) ParseEIP712DomainChanged(log types.Log) (*BridgeEIP712DomainChanged, error) {
	event := new(BridgeEIP712DomainChanged)
	if err := _Bridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeFulfillSetIterator is returned from FilterFeeFulfillSet and is used to iterate over the raw logs and unpacked data for FeeFulfillSet events raised by the Bridge contract.
type BridgeFeeFulfillSetIterator struct {
	Event *BridgeFeeFulfillSet // Event containing the contract specifics and raw log

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
func (it *BridgeFeeFulfillSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeFulfillSet)
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
		it.Event = new(BridgeFeeFulfillSet)
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
func (it *BridgeFeeFulfillSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeFulfillSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeFulfillSet represents a FeeFulfillSet event raised by the Bridge contract.
type BridgeFeeFulfillSet struct {
	FeeFulfill *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeFulfillSet is a free log retrieval operation binding the contract event 0xfd6051a8a15da4aa1eda60d1ba20902908dddb289541ea6b0a8b257f53b2be4f.
//
// Solidity: event FeeFulfillSet(uint256 feeFulfill)
func (_Bridge *BridgeFilterer) FilterFeeFulfillSet(opts *bind.FilterOpts) (*BridgeFeeFulfillSetIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "FeeFulfillSet")
	if err != nil {
		return nil, err
	}
	return &BridgeFeeFulfillSetIterator{contract: _Bridge.contract, event: "FeeFulfillSet", logs: logs, sub: sub}, nil
}

// WatchFeeFulfillSet is a free log subscription operation binding the contract event 0xfd6051a8a15da4aa1eda60d1ba20902908dddb289541ea6b0a8b257f53b2be4f.
//
// Solidity: event FeeFulfillSet(uint256 feeFulfill)
func (_Bridge *BridgeFilterer) WatchFeeFulfillSet(opts *bind.WatchOpts, sink chan<- *BridgeFeeFulfillSet) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "FeeFulfillSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeFulfillSet)
				if err := _Bridge.contract.UnpackLog(event, "FeeFulfillSet", log); err != nil {
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

// ParseFeeFulfillSet is a log parse operation binding the contract event 0xfd6051a8a15da4aa1eda60d1ba20902908dddb289541ea6b0a8b257f53b2be4f.
//
// Solidity: event FeeFulfillSet(uint256 feeFulfill)
func (_Bridge *BridgeFilterer) ParseFeeFulfillSet(log types.Log) (*BridgeFeeFulfillSet, error) {
	event := new(BridgeFeeFulfillSet)
	if err := _Bridge.contract.UnpackLog(event, "FeeFulfillSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeSetIterator is returned from FilterFeeSet and is used to iterate over the raw logs and unpacked data for FeeSet events raised by the Bridge contract.
type BridgeFeeSetIterator struct {
	Event *BridgeFeeSet // Event containing the contract specifics and raw log

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
func (it *BridgeFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeSet)
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
		it.Event = new(BridgeFeeSet)
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
func (it *BridgeFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeSet represents a FeeSet event raised by the Bridge contract.
type BridgeFeeSet struct {
	FeeSend    *big.Int
	FeeFulfill *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeSet is a free log retrieval operation binding the contract event 0x74dbbbe280ef27b79a8a0c449d5ae2ba7a31849103241d0f98df70bbc9d03e37.
//
// Solidity: event FeeSet(uint256 feeSend, uint256 feeFulfill)
func (_Bridge *BridgeFilterer) FilterFeeSet(opts *bind.FilterOpts) (*BridgeFeeSetIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "FeeSet")
	if err != nil {
		return nil, err
	}
	return &BridgeFeeSetIterator{contract: _Bridge.contract, event: "FeeSet", logs: logs, sub: sub}, nil
}

// WatchFeeSet is a free log subscription operation binding the contract event 0x74dbbbe280ef27b79a8a0c449d5ae2ba7a31849103241d0f98df70bbc9d03e37.
//
// Solidity: event FeeSet(uint256 feeSend, uint256 feeFulfill)
func (_Bridge *BridgeFilterer) WatchFeeSet(opts *bind.WatchOpts, sink chan<- *BridgeFeeSet) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "FeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeSet)
				if err := _Bridge.contract.UnpackLog(event, "FeeSet", log); err != nil {
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

// ParseFeeSet is a log parse operation binding the contract event 0x74dbbbe280ef27b79a8a0c449d5ae2ba7a31849103241d0f98df70bbc9d03e37.
//
// Solidity: event FeeSet(uint256 feeSend, uint256 feeFulfill)
func (_Bridge *BridgeFilterer) ParseFeeSet(log types.Log) (*BridgeFeeSet, error) {
	event := new(BridgeFeeSet)
	if err := _Bridge.contract.UnpackLog(event, "FeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeWalletSetIterator is returned from FilterFeeWalletSet and is used to iterate over the raw logs and unpacked data for FeeWalletSet events raised by the Bridge contract.
type BridgeFeeWalletSetIterator struct {
	Event *BridgeFeeWalletSet // Event containing the contract specifics and raw log

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
func (it *BridgeFeeWalletSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeWalletSet)
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
		it.Event = new(BridgeFeeWalletSet)
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
func (it *BridgeFeeWalletSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeWalletSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeWalletSet represents a FeeWalletSet event raised by the Bridge contract.
type BridgeFeeWalletSet struct {
	FeeWallet common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeeWalletSet is a free log retrieval operation binding the contract event 0x0439162378682c6392215882cf21e973f49063ed3530a31d763835a2e9f6d17e.
//
// Solidity: event FeeWalletSet(address indexed feeWallet)
func (_Bridge *BridgeFilterer) FilterFeeWalletSet(opts *bind.FilterOpts, feeWallet []common.Address) (*BridgeFeeWalletSetIterator, error) {

	var feeWalletRule []interface{}
	for _, feeWalletItem := range feeWallet {
		feeWalletRule = append(feeWalletRule, feeWalletItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "FeeWalletSet", feeWalletRule)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeWalletSetIterator{contract: _Bridge.contract, event: "FeeWalletSet", logs: logs, sub: sub}, nil
}

// WatchFeeWalletSet is a free log subscription operation binding the contract event 0x0439162378682c6392215882cf21e973f49063ed3530a31d763835a2e9f6d17e.
//
// Solidity: event FeeWalletSet(address indexed feeWallet)
func (_Bridge *BridgeFilterer) WatchFeeWalletSet(opts *bind.WatchOpts, sink chan<- *BridgeFeeWalletSet, feeWallet []common.Address) (event.Subscription, error) {

	var feeWalletRule []interface{}
	for _, feeWalletItem := range feeWallet {
		feeWalletRule = append(feeWalletRule, feeWalletItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "FeeWalletSet", feeWalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeWalletSet)
				if err := _Bridge.contract.UnpackLog(event, "FeeWalletSet", log); err != nil {
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

// ParseFeeWalletSet is a log parse operation binding the contract event 0x0439162378682c6392215882cf21e973f49063ed3530a31d763835a2e9f6d17e.
//
// Solidity: event FeeWalletSet(address indexed feeWallet)
func (_Bridge *BridgeFilterer) ParseFeeWalletSet(log types.Log) (*BridgeFeeWalletSet, error) {
	event := new(BridgeFeeWalletSet)
	if err := _Bridge.contract.UnpackLog(event, "FeeWalletSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFulfilledTokensIterator is returned from FilterFulfilledTokens and is used to iterate over the raw logs and unpacked data for FulfilledTokens events raised by the Bridge contract.
type BridgeFulfilledTokensIterator struct {
	Event *BridgeFulfilledTokens // Event containing the contract specifics and raw log

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
func (it *BridgeFulfilledTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFulfilledTokens)
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
		it.Event = new(BridgeFulfilledTokens)
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
func (it *BridgeFulfilledTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFulfilledTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFulfilledTokens represents a FulfilledTokens event raised by the Bridge contract.
type BridgeFulfilledTokens struct {
	FromUser     common.Hash
	ToUser       common.Address
	FromChain    string
	ToChain      string
	Amount       *big.Int
	ExchangeRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFulfilledTokens is a free log retrieval operation binding the contract event 0x1a82954c00ba4231b18dbc7d7d5187028cfc140ac6efedeeb60397b8710a9d93.
//
// Solidity: event FulfilledTokens(string indexed fromUser, address indexed toUser, string fromChain, string toChain, uint256 amount, uint256 exchangeRate)
func (_Bridge *BridgeFilterer) FilterFulfilledTokens(opts *bind.FilterOpts, fromUser []string, toUser []common.Address) (*BridgeFulfilledTokensIterator, error) {

	var fromUserRule []interface{}
	for _, fromUserItem := range fromUser {
		fromUserRule = append(fromUserRule, fromUserItem)
	}
	var toUserRule []interface{}
	for _, toUserItem := range toUser {
		toUserRule = append(toUserRule, toUserItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "FulfilledTokens", fromUserRule, toUserRule)
	if err != nil {
		return nil, err
	}
	return &BridgeFulfilledTokensIterator{contract: _Bridge.contract, event: "FulfilledTokens", logs: logs, sub: sub}, nil
}

// WatchFulfilledTokens is a free log subscription operation binding the contract event 0x1a82954c00ba4231b18dbc7d7d5187028cfc140ac6efedeeb60397b8710a9d93.
//
// Solidity: event FulfilledTokens(string indexed fromUser, address indexed toUser, string fromChain, string toChain, uint256 amount, uint256 exchangeRate)
func (_Bridge *BridgeFilterer) WatchFulfilledTokens(opts *bind.WatchOpts, sink chan<- *BridgeFulfilledTokens, fromUser []string, toUser []common.Address) (event.Subscription, error) {

	var fromUserRule []interface{}
	for _, fromUserItem := range fromUser {
		fromUserRule = append(fromUserRule, fromUserItem)
	}
	var toUserRule []interface{}
	for _, toUserItem := range toUser {
		toUserRule = append(toUserRule, toUserItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "FulfilledTokens", fromUserRule, toUserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFulfilledTokens)
				if err := _Bridge.contract.UnpackLog(event, "FulfilledTokens", log); err != nil {
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

// ParseFulfilledTokens is a log parse operation binding the contract event 0x1a82954c00ba4231b18dbc7d7d5187028cfc140ac6efedeeb60397b8710a9d93.
//
// Solidity: event FulfilledTokens(string indexed fromUser, address indexed toUser, string fromChain, string toChain, uint256 amount, uint256 exchangeRate)
func (_Bridge *BridgeFilterer) ParseFulfilledTokens(log types.Log) (*BridgeFulfilledTokens, error) {
	event := new(BridgeFulfilledTokens)
	if err := _Bridge.contract.UnpackLog(event, "FulfilledTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bridge contract.
type BridgeInitializedIterator struct {
	Event *BridgeInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeInitialized)
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
		it.Event = new(BridgeInitialized)
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
func (it *BridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeInitialized represents a Initialized event raised by the Bridge contract.
type BridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridge *BridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeInitializedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeInitializedIterator{contract: _Bridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridge *BridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeInitialized)
				if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridge *BridgeFilterer) ParseInitialized(log types.Log) (*BridgeInitialized, error) {
	event := new(BridgeInitialized)
	if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLimitPerSendSetIterator is returned from FilterLimitPerSendSet and is used to iterate over the raw logs and unpacked data for LimitPerSendSet events raised by the Bridge contract.
type BridgeLimitPerSendSetIterator struct {
	Event *BridgeLimitPerSendSet // Event containing the contract specifics and raw log

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
func (it *BridgeLimitPerSendSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLimitPerSendSet)
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
		it.Event = new(BridgeLimitPerSendSet)
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
func (it *BridgeLimitPerSendSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLimitPerSendSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLimitPerSendSet represents a LimitPerSendSet event raised by the Bridge contract.
type BridgeLimitPerSendSet struct {
	LimitPerSend *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLimitPerSendSet is a free log retrieval operation binding the contract event 0xafed7b515412d2f3cb916d7ae1be64019a06ed7205d98e354ea26c4ab928c2be.
//
// Solidity: event LimitPerSendSet(uint256 limitPerSend)
func (_Bridge *BridgeFilterer) FilterLimitPerSendSet(opts *bind.FilterOpts) (*BridgeLimitPerSendSetIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "LimitPerSendSet")
	if err != nil {
		return nil, err
	}
	return &BridgeLimitPerSendSetIterator{contract: _Bridge.contract, event: "LimitPerSendSet", logs: logs, sub: sub}, nil
}

// WatchLimitPerSendSet is a free log subscription operation binding the contract event 0xafed7b515412d2f3cb916d7ae1be64019a06ed7205d98e354ea26c4ab928c2be.
//
// Solidity: event LimitPerSendSet(uint256 limitPerSend)
func (_Bridge *BridgeFilterer) WatchLimitPerSendSet(opts *bind.WatchOpts, sink chan<- *BridgeLimitPerSendSet) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "LimitPerSendSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLimitPerSendSet)
				if err := _Bridge.contract.UnpackLog(event, "LimitPerSendSet", log); err != nil {
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

// ParseLimitPerSendSet is a log parse operation binding the contract event 0xafed7b515412d2f3cb916d7ae1be64019a06ed7205d98e354ea26c4ab928c2be.
//
// Solidity: event LimitPerSendSet(uint256 limitPerSend)
func (_Bridge *BridgeFilterer) ParseLimitPerSendSet(log types.Log) (*BridgeLimitPerSendSet, error) {
	event := new(BridgeLimitPerSendSet)
	if err := _Bridge.contract.UnpackLog(event, "LimitPerSendSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bridge contract.
type BridgePausedIterator struct {
	Event *BridgePaused // Event containing the contract specifics and raw log

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
func (it *BridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePaused)
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
		it.Event = new(BridgePaused)
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
func (it *BridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePaused represents a Paused event raised by the Bridge contract.
type BridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgePausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgePausedIterator{contract: _Bridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgePaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePaused)
				if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) ParsePaused(log types.Log) (*BridgePaused, error) {
	event := new(BridgePaused)
	if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRelayerConsensusThresholdSetIterator is returned from FilterRelayerConsensusThresholdSet and is used to iterate over the raw logs and unpacked data for RelayerConsensusThresholdSet events raised by the Bridge contract.
type BridgeRelayerConsensusThresholdSetIterator struct {
	Event *BridgeRelayerConsensusThresholdSet // Event containing the contract specifics and raw log

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
func (it *BridgeRelayerConsensusThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRelayerConsensusThresholdSet)
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
		it.Event = new(BridgeRelayerConsensusThresholdSet)
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
func (it *BridgeRelayerConsensusThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRelayerConsensusThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRelayerConsensusThresholdSet represents a RelayerConsensusThresholdSet event raised by the Bridge contract.
type BridgeRelayerConsensusThresholdSet struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRelayerConsensusThresholdSet is a free log retrieval operation binding the contract event 0x9b9e4f688e164fdef0488fb51151fe99d23c096317060cd0a46b90e50efff9c1.
//
// Solidity: event RelayerConsensusThresholdSet(uint256 indexed value)
func (_Bridge *BridgeFilterer) FilterRelayerConsensusThresholdSet(opts *bind.FilterOpts, value []*big.Int) (*BridgeRelayerConsensusThresholdSetIterator, error) {

	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RelayerConsensusThresholdSet", valueRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRelayerConsensusThresholdSetIterator{contract: _Bridge.contract, event: "RelayerConsensusThresholdSet", logs: logs, sub: sub}, nil
}

// WatchRelayerConsensusThresholdSet is a free log subscription operation binding the contract event 0x9b9e4f688e164fdef0488fb51151fe99d23c096317060cd0a46b90e50efff9c1.
//
// Solidity: event RelayerConsensusThresholdSet(uint256 indexed value)
func (_Bridge *BridgeFilterer) WatchRelayerConsensusThresholdSet(opts *bind.WatchOpts, sink chan<- *BridgeRelayerConsensusThresholdSet, value []*big.Int) (event.Subscription, error) {

	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RelayerConsensusThresholdSet", valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRelayerConsensusThresholdSet)
				if err := _Bridge.contract.UnpackLog(event, "RelayerConsensusThresholdSet", log); err != nil {
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

// ParseRelayerConsensusThresholdSet is a log parse operation binding the contract event 0x9b9e4f688e164fdef0488fb51151fe99d23c096317060cd0a46b90e50efff9c1.
//
// Solidity: event RelayerConsensusThresholdSet(uint256 indexed value)
func (_Bridge *BridgeFilterer) ParseRelayerConsensusThresholdSet(log types.Log) (*BridgeRelayerConsensusThresholdSet, error) {
	event := new(BridgeRelayerConsensusThresholdSet)
	if err := _Bridge.contract.UnpackLog(event, "RelayerConsensusThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRelayerSetIterator is returned from FilterRelayerSet and is used to iterate over the raw logs and unpacked data for RelayerSet events raised by the Bridge contract.
type BridgeRelayerSetIterator struct {
	Event *BridgeRelayerSet // Event containing the contract specifics and raw log

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
func (it *BridgeRelayerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRelayerSet)
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
		it.Event = new(BridgeRelayerSet)
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
func (it *BridgeRelayerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRelayerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRelayerSet represents a RelayerSet event raised by the Bridge contract.
type BridgeRelayerSet struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerSet is a free log retrieval operation binding the contract event 0xb7041340e0c2a075059bf0488a71c767724be15dae8e737f8460007325e8d857.
//
// Solidity: event RelayerSet(address indexed relayer)
func (_Bridge *BridgeFilterer) FilterRelayerSet(opts *bind.FilterOpts, relayer []common.Address) (*BridgeRelayerSetIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RelayerSet", relayerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRelayerSetIterator{contract: _Bridge.contract, event: "RelayerSet", logs: logs, sub: sub}, nil
}

// WatchRelayerSet is a free log subscription operation binding the contract event 0xb7041340e0c2a075059bf0488a71c767724be15dae8e737f8460007325e8d857.
//
// Solidity: event RelayerSet(address indexed relayer)
func (_Bridge *BridgeFilterer) WatchRelayerSet(opts *bind.WatchOpts, sink chan<- *BridgeRelayerSet, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RelayerSet", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRelayerSet)
				if err := _Bridge.contract.UnpackLog(event, "RelayerSet", log); err != nil {
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

// ParseRelayerSet is a log parse operation binding the contract event 0xb7041340e0c2a075059bf0488a71c767724be15dae8e737f8460007325e8d857.
//
// Solidity: event RelayerSet(address indexed relayer)
func (_Bridge *BridgeFilterer) ParseRelayerSet(log types.Log) (*BridgeRelayerSet, error) {
	event := new(BridgeRelayerSet)
	if err := _Bridge.contract.UnpackLog(event, "RelayerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Bridge contract.
type BridgeRoleAdminChangedIterator struct {
	Event *BridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleAdminChanged)
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
		it.Event = new(BridgeRoleAdminChanged)
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
func (it *BridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleAdminChanged represents a RoleAdminChanged event raised by the Bridge contract.
type BridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridge *BridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BridgeRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleAdminChangedIterator{contract: _Bridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridge *BridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleAdminChanged)
				if err := _Bridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridge *BridgeFilterer) ParseRoleAdminChanged(log types.Log) (*BridgeRoleAdminChanged, error) {
	event := new(BridgeRoleAdminChanged)
	if err := _Bridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Bridge contract.
type BridgeRoleGrantedIterator struct {
	Event *BridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *BridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleGranted)
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
		it.Event = new(BridgeRoleGranted)
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
func (it *BridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleGranted represents a RoleGranted event raised by the Bridge contract.
type BridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleGrantedIterator{contract: _Bridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleGranted)
				if err := _Bridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) ParseRoleGranted(log types.Log) (*BridgeRoleGranted, error) {
	event := new(BridgeRoleGranted)
	if err := _Bridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Bridge contract.
type BridgeRoleRevokedIterator struct {
	Event *BridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleRevoked)
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
		it.Event = new(BridgeRoleRevoked)
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
func (it *BridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleRevoked represents a RoleRevoked event raised by the Bridge contract.
type BridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleRevokedIterator{contract: _Bridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleRevoked)
				if err := _Bridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) ParseRoleRevoked(log types.Log) (*BridgeRoleRevoked, error) {
	event := new(BridgeRoleRevoked)
	if err := _Bridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSentTokensIterator is returned from FilterSentTokens and is used to iterate over the raw logs and unpacked data for SentTokens events raised by the Bridge contract.
type BridgeSentTokensIterator struct {
	Event *BridgeSentTokens // Event containing the contract specifics and raw log

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
func (it *BridgeSentTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSentTokens)
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
		it.Event = new(BridgeSentTokens)
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
func (it *BridgeSentTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSentTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSentTokens represents a SentTokens event raised by the Bridge contract.
type BridgeSentTokens struct {
	FromUser     common.Address
	ToUser       common.Hash
	FromChain    string
	ToChain      string
	Amount       *big.Int
	ExchangeRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSentTokens is a free log retrieval operation binding the contract event 0x39ca9eb9de2c42145bdc9525fbfac365c3f8e2abdf97dd94e64ec1258365b9f7.
//
// Solidity: event SentTokens(address fromUser, string indexed toUser, string fromChain, string toChain, uint256 amount, uint256 exchangeRate)
func (_Bridge *BridgeFilterer) FilterSentTokens(opts *bind.FilterOpts, toUser []string) (*BridgeSentTokensIterator, error) {

	var toUserRule []interface{}
	for _, toUserItem := range toUser {
		toUserRule = append(toUserRule, toUserItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SentTokens", toUserRule)
	if err != nil {
		return nil, err
	}
	return &BridgeSentTokensIterator{contract: _Bridge.contract, event: "SentTokens", logs: logs, sub: sub}, nil
}

// WatchSentTokens is a free log subscription operation binding the contract event 0x39ca9eb9de2c42145bdc9525fbfac365c3f8e2abdf97dd94e64ec1258365b9f7.
//
// Solidity: event SentTokens(address fromUser, string indexed toUser, string fromChain, string toChain, uint256 amount, uint256 exchangeRate)
func (_Bridge *BridgeFilterer) WatchSentTokens(opts *bind.WatchOpts, sink chan<- *BridgeSentTokens, toUser []string) (event.Subscription, error) {

	var toUserRule []interface{}
	for _, toUserItem := range toUser {
		toUserRule = append(toUserRule, toUserItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SentTokens", toUserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSentTokens)
				if err := _Bridge.contract.UnpackLog(event, "SentTokens", log); err != nil {
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

// ParseSentTokens is a log parse operation binding the contract event 0x39ca9eb9de2c42145bdc9525fbfac365c3f8e2abdf97dd94e64ec1258365b9f7.
//
// Solidity: event SentTokens(address fromUser, string indexed toUser, string fromChain, string toChain, uint256 amount, uint256 exchangeRate)
func (_Bridge *BridgeFilterer) ParseSentTokens(log types.Log) (*BridgeSentTokens, error) {
	event := new(BridgeSentTokens)
	if err := _Bridge.contract.UnpackLog(event, "SentTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bridge contract.
type BridgeUnpausedIterator struct {
	Event *BridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnpaused)
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
		it.Event = new(BridgeUnpaused)
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
func (it *BridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnpaused represents a Unpaused event raised by the Bridge contract.
type BridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeUnpausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeUnpausedIterator{contract: _Bridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnpaused)
				if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) ParseUnpaused(log types.Log) (*BridgeUnpaused, error) {
	event := new(BridgeUnpaused)
	if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
