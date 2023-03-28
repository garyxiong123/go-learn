// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zecreylegend

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

// StorageStoredBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type StorageStoredBlockInfo struct {
	BlockSize                    uint16
	BlockNumber                  uint32
	PriorityOperations           uint64
	PendingOnchainOperationsHash [32]byte
	Timestamp                    *big.Int
	StateRoot                    [32]byte
	Commitment                   [32]byte
}

// ZecreyLegendCommitBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type ZecreyLegendCommitBlockInfo struct {
	NewStateRoot      [32]byte
	PublicData        []byte
	Timestamp         *big.Int
	PublicDataOffsets []uint32
	BlockNumber       uint32
	BlockSize         uint16
}

// ZecreyLegendPairInfo is an auto generated low-level Go binding around an user-defined struct.
type ZecreyLegendPairInfo struct {
	TokenA               common.Address
	TokenB               common.Address
	FeeRate              uint16
	TreasuryAccountIndex uint32
	TreasuryRate         uint16
}

// ZecreyLegendVerifyAndExecuteBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type ZecreyLegendVerifyAndExecuteBlockInfo struct {
	BlockHeader              StorageStoredBlockInfo
	PendingOnchainOpsPubData [][]byte
}

// ZecreylegendMetaData contains all meta data concerning the Zecreylegend contract.
var ZecreylegendMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"BlockCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"BlockVerification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalBlocksVerified\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalBlocksCommitted\",\"type\":\"uint32\"}],\"name\":\"BlocksRevert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"asset0Id\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"asset1Id\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeRate\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"treasuryAccountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"treasuryRate\",\"type\":\"uint16\"}],\"name\":\"CreateTokenPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accountName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"zkbnbBlockNumber\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accountName\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"DepositCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accountNameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nftContentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"creatorTreasuryRate\",\"type\":\"uint16\"}],\"name\":\"DepositNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DesertMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"zkbnbBlockId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"FullExitCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"NewDefaultNFTFactory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_creatorAccountNameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"_collectionId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_factoryAddress\",\"type\":\"address\"}],\"name\":\"NewNFTFactory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serialId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumTxTypes.TxType\",\"name\":\"txType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationBlock\",\"type\":\"uint256\"}],\"name\":\"NewPriorityRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNoticePeriod\",\"type\":\"uint256\"}],\"name\":\"NoticePeriodChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"zkbnbPubKeyX\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"zkbnbPubKeyY\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"}],\"name\":\"RegisterZNS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeRate\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"treasuryAccountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"treasuryRate\",\"type\":\"uint16\"}],\"name\":\"UpdateTokenPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftL1Address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftL1TokenId\",\"type\":\"uint256\"}],\"name\":\"WithdrawNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"nftIndex\",\"type\":\"uint40\"}],\"name\":\"WithdrawalNFTPending\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_ACCOUNT_INDEX\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_AMOUNT_OF_REGISTERED_ASSETS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_DEPOSIT_AMOUNT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FUNGIBLE_ASSET_ID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NFT_INDEX\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECURITY_COUNCIL_MEMBERS_NUMBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHORTEST_UPGRADE_NOTICE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPECIAL_ACCOUNT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPECIAL_ACCOUNT_ID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TX_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_NOTICE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateDesertMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.StoredBlockInfo\",\"name\":\"_lastCommittedBlockData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32[]\",\"name\":\"publicDataOffsets\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"}],\"internalType\":\"structZecreyLegend.CommitBlockInfo[]\",\"name\":\"_newBlocksData\",\"type\":\"tuple[]\"}],\"name\":\"commitBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cutUpgradeNoticePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultNFTFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint104\",\"name\":\"_amount\",\"type\":\"uint104\"},{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"}],\"name\":\"depositBEP20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"}],\"name\":\"depositBNB\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_nftL1Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nftL1TokenId\",\"type\":\"uint256\"}],\"name\":\"depositNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"desertMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstPriorityRequestId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"accountNameHash\",\"type\":\"bytes32\"}],\"name\":\"getAddressByAccountNameHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_creatorAccountNameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_collectionId\",\"type\":\"uint32\"}],\"name\":\"getNFTFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNoticePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_assetAddr\",\"type\":\"address\"}],\"name\":\"getPendingBalance\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getZNSNamePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializationParameters\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isReadyForUpgrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"isRegisteredZNSName\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"nftFactories\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_zkbnbPubKeyX\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_zkbnbPubKeyY\",\"type\":\"bytes32\"}],\"name\":\"registerZNS\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"requestFullExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"_nftIndex\",\"type\":\"uint32\"}],\"name\":\"requestFullExitNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.StoredBlockInfo[]\",\"name\":\"_blocksToRevert\",\"type\":\"tuple[]\"}],\"name\":\"revertBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractNFTFactory\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"setDefaultNFTFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"storedBlockHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBlocksCommitted\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBlocksVerified\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalOpenPriorityRequests\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTokenPairs\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_maxAmount\",\"type\":\"uint128\"}],\"name\":\"transferERC20\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"withdrawnAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"feeRate\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"treasuryAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"treasuryRate\",\"type\":\"uint16\"}],\"internalType\":\"structZecreyLegend.PairInfo\",\"name\":\"_pairInfo\",\"type\":\"tuple\"}],\"name\":\"updatePairRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVerifierAddress\",\"type\":\"address\"}],\"name\":\"updateZkBNBVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"upgradeParameters\",\"type\":\"bytes\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeCanceled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeFinishes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeNoticePeriodStarted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradePreparationStarted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.StoredBlockInfo\",\"name\":\"blockHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"pendingOnchainOpsPubData\",\"type\":\"bytes[]\"}],\"internalType\":\"structZecreyLegend.VerifyAndExecuteBlockInfo[]\",\"name\":\"_blocks\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_proofs\",\"type\":\"uint256[]\"}],\"name\":\"verifyAndExecuteBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"}],\"name\":\"withdrawPendingBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_nftIndex\",\"type\":\"uint40\"}],\"name\":\"withdrawPendingNFTBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZecreylegendABI is the input ABI used to generate the binding from.
// Deprecated: Use ZecreylegendMetaData.ABI instead.
var ZecreylegendABI = ZecreylegendMetaData.ABI

// Zecreylegend is an auto generated Go binding around an Ethereum contract.
type Zecreylegend struct {
	ZecreylegendCaller     // Read-only binding to the contract
	ZecreylegendTransactor // Write-only binding to the contract
	ZecreylegendFilterer   // Log filterer for contract events
}

// ZecreylegendCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZecreylegendCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZecreylegendTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZecreylegendTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZecreylegendFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZecreylegendFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZecreylegendSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZecreylegendSession struct {
	Contract     *Zecreylegend     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZecreylegendCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZecreylegendCallerSession struct {
	Contract *ZecreylegendCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ZecreylegendTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZecreylegendTransactorSession struct {
	Contract     *ZecreylegendTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZecreylegendRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZecreylegendRaw struct {
	Contract *Zecreylegend // Generic contract binding to access the raw methods on
}

// ZecreylegendCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZecreylegendCallerRaw struct {
	Contract *ZecreylegendCaller // Generic read-only contract binding to access the raw methods on
}

// ZecreylegendTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZecreylegendTransactorRaw struct {
	Contract *ZecreylegendTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZecreylegend creates a new instance of Zecreylegend, bound to a specific deployed contract.
func NewZecreylegend(address common.Address, backend bind.ContractBackend) (*Zecreylegend, error) {
	contract, err := bindZecreylegend(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zecreylegend{ZecreylegendCaller: ZecreylegendCaller{contract: contract}, ZecreylegendTransactor: ZecreylegendTransactor{contract: contract}, ZecreylegendFilterer: ZecreylegendFilterer{contract: contract}}, nil
}

// NewZecreylegendCaller creates a new read-only instance of Zecreylegend, bound to a specific deployed contract.
func NewZecreylegendCaller(address common.Address, caller bind.ContractCaller) (*ZecreylegendCaller, error) {
	contract, err := bindZecreylegend(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZecreylegendCaller{contract: contract}, nil
}

// NewZecreylegendTransactor creates a new write-only instance of Zecreylegend, bound to a specific deployed contract.
func NewZecreylegendTransactor(address common.Address, transactor bind.ContractTransactor) (*ZecreylegendTransactor, error) {
	contract, err := bindZecreylegend(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZecreylegendTransactor{contract: contract}, nil
}

// NewZecreylegendFilterer creates a new log filterer instance of Zecreylegend, bound to a specific deployed contract.
func NewZecreylegendFilterer(address common.Address, filterer bind.ContractFilterer) (*ZecreylegendFilterer, error) {
	contract, err := bindZecreylegend(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZecreylegendFilterer{contract: contract}, nil
}

// bindZecreylegend binds a generic wrapper to an already deployed contract.
func bindZecreylegend(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZecreylegendABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zecreylegend *ZecreylegendRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zecreylegend.Contract.ZecreylegendCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zecreylegend *ZecreylegendRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zecreylegend.Contract.ZecreylegendTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zecreylegend *ZecreylegendRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zecreylegend.Contract.ZecreylegendTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zecreylegend *ZecreylegendCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zecreylegend.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zecreylegend *ZecreylegendTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zecreylegend.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zecreylegend *ZecreylegendTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zecreylegend.Contract.contract.Transact(opts, method, params...)
}

// MAXACCOUNTINDEX is a free data retrieval call binding the contract method 0x437545f9.
//
// Solidity: function MAX_ACCOUNT_INDEX() view returns(uint32)
func (_Zecreylegend *ZecreylegendCaller) MAXACCOUNTINDEX(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "MAX_ACCOUNT_INDEX")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXACCOUNTINDEX is a free data retrieval call binding the contract method 0x437545f9.
//
// Solidity: function MAX_ACCOUNT_INDEX() view returns(uint32)
func (_Zecreylegend *ZecreylegendSession) MAXACCOUNTINDEX() (uint32, error) {
	return _Zecreylegend.Contract.MAXACCOUNTINDEX(&_Zecreylegend.CallOpts)
}

// MAXACCOUNTINDEX is a free data retrieval call binding the contract method 0x437545f9.
//
// Solidity: function MAX_ACCOUNT_INDEX() view returns(uint32)
func (_Zecreylegend *ZecreylegendCallerSession) MAXACCOUNTINDEX() (uint32, error) {
	return _Zecreylegend.Contract.MAXACCOUNTINDEX(&_Zecreylegend.CallOpts)
}

// MAXAMOUNTOFREGISTEREDASSETS is a free data retrieval call binding the contract method 0x0d360b7f.
//
// Solidity: function MAX_AMOUNT_OF_REGISTERED_ASSETS() view returns(uint16)
func (_Zecreylegend *ZecreylegendCaller) MAXAMOUNTOFREGISTEREDASSETS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "MAX_AMOUNT_OF_REGISTERED_ASSETS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXAMOUNTOFREGISTEREDASSETS is a free data retrieval call binding the contract method 0x0d360b7f.
//
// Solidity: function MAX_AMOUNT_OF_REGISTERED_ASSETS() view returns(uint16)
func (_Zecreylegend *ZecreylegendSession) MAXAMOUNTOFREGISTEREDASSETS() (uint16, error) {
	return _Zecreylegend.Contract.MAXAMOUNTOFREGISTEREDASSETS(&_Zecreylegend.CallOpts)
}

// MAXAMOUNTOFREGISTEREDASSETS is a free data retrieval call binding the contract method 0x0d360b7f.
//
// Solidity: function MAX_AMOUNT_OF_REGISTERED_ASSETS() view returns(uint16)
func (_Zecreylegend *ZecreylegendCallerSession) MAXAMOUNTOFREGISTEREDASSETS() (uint16, error) {
	return _Zecreylegend.Contract.MAXAMOUNTOFREGISTEREDASSETS(&_Zecreylegend.CallOpts)
}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(uint128)
func (_Zecreylegend *ZecreylegendCaller) MAXDEPOSITAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "MAX_DEPOSIT_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(uint128)
func (_Zecreylegend *ZecreylegendSession) MAXDEPOSITAMOUNT() (*big.Int, error) {
	return _Zecreylegend.Contract.MAXDEPOSITAMOUNT(&_Zecreylegend.CallOpts)
}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(uint128)
func (_Zecreylegend *ZecreylegendCallerSession) MAXDEPOSITAMOUNT() (*big.Int, error) {
	return _Zecreylegend.Contract.MAXDEPOSITAMOUNT(&_Zecreylegend.CallOpts)
}

// MAXFUNGIBLEASSETID is a free data retrieval call binding the contract method 0x437da02f.
//
// Solidity: function MAX_FUNGIBLE_ASSET_ID() view returns(uint32)
func (_Zecreylegend *ZecreylegendCaller) MAXFUNGIBLEASSETID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "MAX_FUNGIBLE_ASSET_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXFUNGIBLEASSETID is a free data retrieval call binding the contract method 0x437da02f.
//
// Solidity: function MAX_FUNGIBLE_ASSET_ID() view returns(uint32)
func (_Zecreylegend *ZecreylegendSession) MAXFUNGIBLEASSETID() (uint32, error) {
	return _Zecreylegend.Contract.MAXFUNGIBLEASSETID(&_Zecreylegend.CallOpts)
}

// MAXFUNGIBLEASSETID is a free data retrieval call binding the contract method 0x437da02f.
//
// Solidity: function MAX_FUNGIBLE_ASSET_ID() view returns(uint32)
func (_Zecreylegend *ZecreylegendCallerSession) MAXFUNGIBLEASSETID() (uint32, error) {
	return _Zecreylegend.Contract.MAXFUNGIBLEASSETID(&_Zecreylegend.CallOpts)
}

// MAXNFTINDEX is a free data retrieval call binding the contract method 0x14791ad2.
//
// Solidity: function MAX_NFT_INDEX() view returns(uint40)
func (_Zecreylegend *ZecreylegendCaller) MAXNFTINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "MAX_NFT_INDEX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNFTINDEX is a free data retrieval call binding the contract method 0x14791ad2.
//
// Solidity: function MAX_NFT_INDEX() view returns(uint40)
func (_Zecreylegend *ZecreylegendSession) MAXNFTINDEX() (*big.Int, error) {
	return _Zecreylegend.Contract.MAXNFTINDEX(&_Zecreylegend.CallOpts)
}

// MAXNFTINDEX is a free data retrieval call binding the contract method 0x14791ad2.
//
// Solidity: function MAX_NFT_INDEX() view returns(uint40)
func (_Zecreylegend *ZecreylegendCallerSession) MAXNFTINDEX() (*big.Int, error) {
	return _Zecreylegend.Contract.MAXNFTINDEX(&_Zecreylegend.CallOpts)
}

// SECURITYCOUNCILMEMBERSNUMBER is a free data retrieval call binding the contract method 0x4a51a71f.
//
// Solidity: function SECURITY_COUNCIL_MEMBERS_NUMBER() view returns(uint256)
func (_Zecreylegend *ZecreylegendCaller) SECURITYCOUNCILMEMBERSNUMBER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "SECURITY_COUNCIL_MEMBERS_NUMBER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECURITYCOUNCILMEMBERSNUMBER is a free data retrieval call binding the contract method 0x4a51a71f.
//
// Solidity: function SECURITY_COUNCIL_MEMBERS_NUMBER() view returns(uint256)
func (_Zecreylegend *ZecreylegendSession) SECURITYCOUNCILMEMBERSNUMBER() (*big.Int, error) {
	return _Zecreylegend.Contract.SECURITYCOUNCILMEMBERSNUMBER(&_Zecreylegend.CallOpts)
}

// SECURITYCOUNCILMEMBERSNUMBER is a free data retrieval call binding the contract method 0x4a51a71f.
//
// Solidity: function SECURITY_COUNCIL_MEMBERS_NUMBER() view returns(uint256)
func (_Zecreylegend *ZecreylegendCallerSession) SECURITYCOUNCILMEMBERSNUMBER() (*big.Int, error) {
	return _Zecreylegend.Contract.SECURITYCOUNCILMEMBERSNUMBER(&_Zecreylegend.CallOpts)
}

// SHORTESTUPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0x85053581.
//
// Solidity: function SHORTEST_UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Zecreylegend *ZecreylegendCaller) SHORTESTUPGRADENOTICEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "SHORTEST_UPGRADE_NOTICE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHORTESTUPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0x85053581.
//
// Solidity: function SHORTEST_UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Zecreylegend *ZecreylegendSession) SHORTESTUPGRADENOTICEPERIOD() (*big.Int, error) {
	return _Zecreylegend.Contract.SHORTESTUPGRADENOTICEPERIOD(&_Zecreylegend.CallOpts)
}

// SHORTESTUPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0x85053581.
//
// Solidity: function SHORTEST_UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Zecreylegend *ZecreylegendCallerSession) SHORTESTUPGRADENOTICEPERIOD() (*big.Int, error) {
	return _Zecreylegend.Contract.SHORTESTUPGRADENOTICEPERIOD(&_Zecreylegend.CallOpts)
}

// SPECIALACCOUNTADDRESS is a free data retrieval call binding the contract method 0x7ea399c1.
//
// Solidity: function SPECIAL_ACCOUNT_ADDRESS() view returns(address)
func (_Zecreylegend *ZecreylegendCaller) SPECIALACCOUNTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "SPECIAL_ACCOUNT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SPECIALACCOUNTADDRESS is a free data retrieval call binding the contract method 0x7ea399c1.
//
// Solidity: function SPECIAL_ACCOUNT_ADDRESS() view returns(address)
func (_Zecreylegend *ZecreylegendSession) SPECIALACCOUNTADDRESS() (common.Address, error) {
	return _Zecreylegend.Contract.SPECIALACCOUNTADDRESS(&_Zecreylegend.CallOpts)
}

// SPECIALACCOUNTADDRESS is a free data retrieval call binding the contract method 0x7ea399c1.
//
// Solidity: function SPECIAL_ACCOUNT_ADDRESS() view returns(address)
func (_Zecreylegend *ZecreylegendCallerSession) SPECIALACCOUNTADDRESS() (common.Address, error) {
	return _Zecreylegend.Contract.SPECIALACCOUNTADDRESS(&_Zecreylegend.CallOpts)
}

// SPECIALACCOUNTID is a free data retrieval call binding the contract method 0x4242d5b3.
//
// Solidity: function SPECIAL_ACCOUNT_ID() view returns(uint32)
func (_Zecreylegend *ZecreylegendCaller) SPECIALACCOUNTID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "SPECIAL_ACCOUNT_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SPECIALACCOUNTID is a free data retrieval call binding the contract method 0x4242d5b3.
//
// Solidity: function SPECIAL_ACCOUNT_ID() view returns(uint32)
func (_Zecreylegend *ZecreylegendSession) SPECIALACCOUNTID() (uint32, error) {
	return _Zecreylegend.Contract.SPECIALACCOUNTID(&_Zecreylegend.CallOpts)
}

// SPECIALACCOUNTID is a free data retrieval call binding the contract method 0x4242d5b3.
//
// Solidity: function SPECIAL_ACCOUNT_ID() view returns(uint32)
func (_Zecreylegend *ZecreylegendCallerSession) SPECIALACCOUNTID() (uint32, error) {
	return _Zecreylegend.Contract.SPECIALACCOUNTID(&_Zecreylegend.CallOpts)
}

// TXSIZE is a free data retrieval call binding the contract method 0xe6e3c012.
//
// Solidity: function TX_SIZE() view returns(uint256)
func (_Zecreylegend *ZecreylegendCaller) TXSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "TX_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TXSIZE is a free data retrieval call binding the contract method 0xe6e3c012.
//
// Solidity: function TX_SIZE() view returns(uint256)
func (_Zecreylegend *ZecreylegendSession) TXSIZE() (*big.Int, error) {
	return _Zecreylegend.Contract.TXSIZE(&_Zecreylegend.CallOpts)
}

// TXSIZE is a free data retrieval call binding the contract method 0xe6e3c012.
//
// Solidity: function TX_SIZE() view returns(uint256)
func (_Zecreylegend *ZecreylegendCallerSession) TXSIZE() (*big.Int, error) {
	return _Zecreylegend.Contract.TXSIZE(&_Zecreylegend.CallOpts)
}

// UPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0xcc375fb7.
//
// Solidity: function UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Zecreylegend *ZecreylegendCaller) UPGRADENOTICEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "UPGRADE_NOTICE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0xcc375fb7.
//
// Solidity: function UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Zecreylegend *ZecreylegendSession) UPGRADENOTICEPERIOD() (*big.Int, error) {
	return _Zecreylegend.Contract.UPGRADENOTICEPERIOD(&_Zecreylegend.CallOpts)
}

// UPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0xcc375fb7.
//
// Solidity: function UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Zecreylegend *ZecreylegendCallerSession) UPGRADENOTICEPERIOD() (*big.Int, error) {
	return _Zecreylegend.Contract.UPGRADENOTICEPERIOD(&_Zecreylegend.CallOpts)
}

// WITHDRAWALGASLIMIT is a free data retrieval call binding the contract method 0xc701f955.
//
// Solidity: function WITHDRAWAL_GAS_LIMIT() view returns(uint256)
func (_Zecreylegend *ZecreylegendCaller) WITHDRAWALGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "WITHDRAWAL_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWALGASLIMIT is a free data retrieval call binding the contract method 0xc701f955.
//
// Solidity: function WITHDRAWAL_GAS_LIMIT() view returns(uint256)
func (_Zecreylegend *ZecreylegendSession) WITHDRAWALGASLIMIT() (*big.Int, error) {
	return _Zecreylegend.Contract.WITHDRAWALGASLIMIT(&_Zecreylegend.CallOpts)
}

// WITHDRAWALGASLIMIT is a free data retrieval call binding the contract method 0xc701f955.
//
// Solidity: function WITHDRAWAL_GAS_LIMIT() view returns(uint256)
func (_Zecreylegend *ZecreylegendCallerSession) WITHDRAWALGASLIMIT() (*big.Int, error) {
	return _Zecreylegend.Contract.WITHDRAWALGASLIMIT(&_Zecreylegend.CallOpts)
}

// DefaultNFTFactory is a free data retrieval call binding the contract method 0x940f19c0.
//
// Solidity: function defaultNFTFactory() view returns(address)
func (_Zecreylegend *ZecreylegendCaller) DefaultNFTFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "defaultNFTFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultNFTFactory is a free data retrieval call binding the contract method 0x940f19c0.
//
// Solidity: function defaultNFTFactory() view returns(address)
func (_Zecreylegend *ZecreylegendSession) DefaultNFTFactory() (common.Address, error) {
	return _Zecreylegend.Contract.DefaultNFTFactory(&_Zecreylegend.CallOpts)
}

// DefaultNFTFactory is a free data retrieval call binding the contract method 0x940f19c0.
//
// Solidity: function defaultNFTFactory() view returns(address)
func (_Zecreylegend *ZecreylegendCallerSession) DefaultNFTFactory() (common.Address, error) {
	return _Zecreylegend.Contract.DefaultNFTFactory(&_Zecreylegend.CallOpts)
}

// DesertMode is a free data retrieval call binding the contract method 0x02cfb563.
//
// Solidity: function desertMode() view returns(bool)
func (_Zecreylegend *ZecreylegendCaller) DesertMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "desertMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DesertMode is a free data retrieval call binding the contract method 0x02cfb563.
//
// Solidity: function desertMode() view returns(bool)
func (_Zecreylegend *ZecreylegendSession) DesertMode() (bool, error) {
	return _Zecreylegend.Contract.DesertMode(&_Zecreylegend.CallOpts)
}

// DesertMode is a free data retrieval call binding the contract method 0x02cfb563.
//
// Solidity: function desertMode() view returns(bool)
func (_Zecreylegend *ZecreylegendCallerSession) DesertMode() (bool, error) {
	return _Zecreylegend.Contract.DesertMode(&_Zecreylegend.CallOpts)
}

// FirstPriorityRequestId is a free data retrieval call binding the contract method 0x67708dae.
//
// Solidity: function firstPriorityRequestId() view returns(uint64)
func (_Zecreylegend *ZecreylegendCaller) FirstPriorityRequestId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "firstPriorityRequestId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FirstPriorityRequestId is a free data retrieval call binding the contract method 0x67708dae.
//
// Solidity: function firstPriorityRequestId() view returns(uint64)
func (_Zecreylegend *ZecreylegendSession) FirstPriorityRequestId() (uint64, error) {
	return _Zecreylegend.Contract.FirstPriorityRequestId(&_Zecreylegend.CallOpts)
}

// FirstPriorityRequestId is a free data retrieval call binding the contract method 0x67708dae.
//
// Solidity: function firstPriorityRequestId() view returns(uint64)
func (_Zecreylegend *ZecreylegendCallerSession) FirstPriorityRequestId() (uint64, error) {
	return _Zecreylegend.Contract.FirstPriorityRequestId(&_Zecreylegend.CallOpts)
}

// GetAddressByAccountNameHash is a free data retrieval call binding the contract method 0x0b8f1c0c.
//
// Solidity: function getAddressByAccountNameHash(bytes32 accountNameHash) view returns(address)
func (_Zecreylegend *ZecreylegendCaller) GetAddressByAccountNameHash(opts *bind.CallOpts, accountNameHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "getAddressByAccountNameHash", accountNameHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressByAccountNameHash is a free data retrieval call binding the contract method 0x0b8f1c0c.
//
// Solidity: function getAddressByAccountNameHash(bytes32 accountNameHash) view returns(address)
func (_Zecreylegend *ZecreylegendSession) GetAddressByAccountNameHash(accountNameHash [32]byte) (common.Address, error) {
	return _Zecreylegend.Contract.GetAddressByAccountNameHash(&_Zecreylegend.CallOpts, accountNameHash)
}

// GetAddressByAccountNameHash is a free data retrieval call binding the contract method 0x0b8f1c0c.
//
// Solidity: function getAddressByAccountNameHash(bytes32 accountNameHash) view returns(address)
func (_Zecreylegend *ZecreylegendCallerSession) GetAddressByAccountNameHash(accountNameHash [32]byte) (common.Address, error) {
	return _Zecreylegend.Contract.GetAddressByAccountNameHash(&_Zecreylegend.CallOpts, accountNameHash)
}

// GetNFTFactory is a free data retrieval call binding the contract method 0x76b4da60.
//
// Solidity: function getNFTFactory(bytes32 _creatorAccountNameHash, uint32 _collectionId) view returns(address)
func (_Zecreylegend *ZecreylegendCaller) GetNFTFactory(opts *bind.CallOpts, _creatorAccountNameHash [32]byte, _collectionId uint32) (common.Address, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "getNFTFactory", _creatorAccountNameHash, _collectionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNFTFactory is a free data retrieval call binding the contract method 0x76b4da60.
//
// Solidity: function getNFTFactory(bytes32 _creatorAccountNameHash, uint32 _collectionId) view returns(address)
func (_Zecreylegend *ZecreylegendSession) GetNFTFactory(_creatorAccountNameHash [32]byte, _collectionId uint32) (common.Address, error) {
	return _Zecreylegend.Contract.GetNFTFactory(&_Zecreylegend.CallOpts, _creatorAccountNameHash, _collectionId)
}

// GetNFTFactory is a free data retrieval call binding the contract method 0x76b4da60.
//
// Solidity: function getNFTFactory(bytes32 _creatorAccountNameHash, uint32 _collectionId) view returns(address)
func (_Zecreylegend *ZecreylegendCallerSession) GetNFTFactory(_creatorAccountNameHash [32]byte, _collectionId uint32) (common.Address, error) {
	return _Zecreylegend.Contract.GetNFTFactory(&_Zecreylegend.CallOpts, _creatorAccountNameHash, _collectionId)
}

// GetNoticePeriod is a free data retrieval call binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() pure returns(uint256)
func (_Zecreylegend *ZecreylegendCaller) GetNoticePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "getNoticePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNoticePeriod is a free data retrieval call binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() pure returns(uint256)
func (_Zecreylegend *ZecreylegendSession) GetNoticePeriod() (*big.Int, error) {
	return _Zecreylegend.Contract.GetNoticePeriod(&_Zecreylegend.CallOpts)
}

// GetNoticePeriod is a free data retrieval call binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() pure returns(uint256)
func (_Zecreylegend *ZecreylegendCallerSession) GetNoticePeriod() (*big.Int, error) {
	return _Zecreylegend.Contract.GetNoticePeriod(&_Zecreylegend.CallOpts)
}

// GetPendingBalance is a free data retrieval call binding the contract method 0x5aca41f6.
//
// Solidity: function getPendingBalance(address _address, address _assetAddr) view returns(uint128)
func (_Zecreylegend *ZecreylegendCaller) GetPendingBalance(opts *bind.CallOpts, _address common.Address, _assetAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "getPendingBalance", _address, _assetAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingBalance is a free data retrieval call binding the contract method 0x5aca41f6.
//
// Solidity: function getPendingBalance(address _address, address _assetAddr) view returns(uint128)
func (_Zecreylegend *ZecreylegendSession) GetPendingBalance(_address common.Address, _assetAddr common.Address) (*big.Int, error) {
	return _Zecreylegend.Contract.GetPendingBalance(&_Zecreylegend.CallOpts, _address, _assetAddr)
}

// GetPendingBalance is a free data retrieval call binding the contract method 0x5aca41f6.
//
// Solidity: function getPendingBalance(address _address, address _assetAddr) view returns(uint128)
func (_Zecreylegend *ZecreylegendCallerSession) GetPendingBalance(_address common.Address, _assetAddr common.Address) (*big.Int, error) {
	return _Zecreylegend.Contract.GetPendingBalance(&_Zecreylegend.CallOpts, _address, _assetAddr)
}

// GetZNSNamePrice is a free data retrieval call binding the contract method 0x1c6b30e1.
//
// Solidity: function getZNSNamePrice(string name) view returns(uint256)
func (_Zecreylegend *ZecreylegendCaller) GetZNSNamePrice(opts *bind.CallOpts, name string) (*big.Int, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "getZNSNamePrice", name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetZNSNamePrice is a free data retrieval call binding the contract method 0x1c6b30e1.
//
// Solidity: function getZNSNamePrice(string name) view returns(uint256)
func (_Zecreylegend *ZecreylegendSession) GetZNSNamePrice(name string) (*big.Int, error) {
	return _Zecreylegend.Contract.GetZNSNamePrice(&_Zecreylegend.CallOpts, name)
}

// GetZNSNamePrice is a free data retrieval call binding the contract method 0x1c6b30e1.
//
// Solidity: function getZNSNamePrice(string name) view returns(uint256)
func (_Zecreylegend *ZecreylegendCallerSession) GetZNSNamePrice(name string) (*big.Int, error) {
	return _Zecreylegend.Contract.GetZNSNamePrice(&_Zecreylegend.CallOpts, name)
}

// IsReadyForUpgrade is a free data retrieval call binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() view returns(bool)
func (_Zecreylegend *ZecreylegendCaller) IsReadyForUpgrade(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "isReadyForUpgrade")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReadyForUpgrade is a free data retrieval call binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() view returns(bool)
func (_Zecreylegend *ZecreylegendSession) IsReadyForUpgrade() (bool, error) {
	return _Zecreylegend.Contract.IsReadyForUpgrade(&_Zecreylegend.CallOpts)
}

// IsReadyForUpgrade is a free data retrieval call binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() view returns(bool)
func (_Zecreylegend *ZecreylegendCallerSession) IsReadyForUpgrade() (bool, error) {
	return _Zecreylegend.Contract.IsReadyForUpgrade(&_Zecreylegend.CallOpts)
}

// IsRegisteredZNSName is a free data retrieval call binding the contract method 0xa1af3e5c.
//
// Solidity: function isRegisteredZNSName(string _name) view returns(bool)
func (_Zecreylegend *ZecreylegendCaller) IsRegisteredZNSName(opts *bind.CallOpts, _name string) (bool, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "isRegisteredZNSName", _name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegisteredZNSName is a free data retrieval call binding the contract method 0xa1af3e5c.
//
// Solidity: function isRegisteredZNSName(string _name) view returns(bool)
func (_Zecreylegend *ZecreylegendSession) IsRegisteredZNSName(_name string) (bool, error) {
	return _Zecreylegend.Contract.IsRegisteredZNSName(&_Zecreylegend.CallOpts, _name)
}

// IsRegisteredZNSName is a free data retrieval call binding the contract method 0xa1af3e5c.
//
// Solidity: function isRegisteredZNSName(string _name) view returns(bool)
func (_Zecreylegend *ZecreylegendCallerSession) IsRegisteredZNSName(_name string) (bool, error) {
	return _Zecreylegend.Contract.IsRegisteredZNSName(&_Zecreylegend.CallOpts, _name)
}

// NftFactories is a free data retrieval call binding the contract method 0x16335e06.
//
// Solidity: function nftFactories(bytes32 , uint32 ) view returns(address)
func (_Zecreylegend *ZecreylegendCaller) NftFactories(opts *bind.CallOpts, arg0 [32]byte, arg1 uint32) (common.Address, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "nftFactories", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftFactories is a free data retrieval call binding the contract method 0x16335e06.
//
// Solidity: function nftFactories(bytes32 , uint32 ) view returns(address)
func (_Zecreylegend *ZecreylegendSession) NftFactories(arg0 [32]byte, arg1 uint32) (common.Address, error) {
	return _Zecreylegend.Contract.NftFactories(&_Zecreylegend.CallOpts, arg0, arg1)
}

// NftFactories is a free data retrieval call binding the contract method 0x16335e06.
//
// Solidity: function nftFactories(bytes32 , uint32 ) view returns(address)
func (_Zecreylegend *ZecreylegendCallerSession) NftFactories(arg0 [32]byte, arg1 uint32) (common.Address, error) {
	return _Zecreylegend.Contract.NftFactories(&_Zecreylegend.CallOpts, arg0, arg1)
}

// StateRoot is a free data retrieval call binding the contract method 0x9588eca2.
//
// Solidity: function stateRoot() view returns(bytes32)
func (_Zecreylegend *ZecreylegendCaller) StateRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "stateRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateRoot is a free data retrieval call binding the contract method 0x9588eca2.
//
// Solidity: function stateRoot() view returns(bytes32)
func (_Zecreylegend *ZecreylegendSession) StateRoot() ([32]byte, error) {
	return _Zecreylegend.Contract.StateRoot(&_Zecreylegend.CallOpts)
}

// StateRoot is a free data retrieval call binding the contract method 0x9588eca2.
//
// Solidity: function stateRoot() view returns(bytes32)
func (_Zecreylegend *ZecreylegendCallerSession) StateRoot() ([32]byte, error) {
	return _Zecreylegend.Contract.StateRoot(&_Zecreylegend.CallOpts)
}

// StoredBlockHashes is a free data retrieval call binding the contract method 0x9ba0d146.
//
// Solidity: function storedBlockHashes(uint32 ) view returns(bytes32)
func (_Zecreylegend *ZecreylegendCaller) StoredBlockHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "storedBlockHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StoredBlockHashes is a free data retrieval call binding the contract method 0x9ba0d146.
//
// Solidity: function storedBlockHashes(uint32 ) view returns(bytes32)
func (_Zecreylegend *ZecreylegendSession) StoredBlockHashes(arg0 uint32) ([32]byte, error) {
	return _Zecreylegend.Contract.StoredBlockHashes(&_Zecreylegend.CallOpts, arg0)
}

// StoredBlockHashes is a free data retrieval call binding the contract method 0x9ba0d146.
//
// Solidity: function storedBlockHashes(uint32 ) view returns(bytes32)
func (_Zecreylegend *ZecreylegendCallerSession) StoredBlockHashes(arg0 uint32) ([32]byte, error) {
	return _Zecreylegend.Contract.StoredBlockHashes(&_Zecreylegend.CallOpts, arg0)
}

// TotalBlocksCommitted is a free data retrieval call binding the contract method 0xfaf4d8cb.
//
// Solidity: function totalBlocksCommitted() view returns(uint32)
func (_Zecreylegend *ZecreylegendCaller) TotalBlocksCommitted(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "totalBlocksCommitted")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TotalBlocksCommitted is a free data retrieval call binding the contract method 0xfaf4d8cb.
//
// Solidity: function totalBlocksCommitted() view returns(uint32)
func (_Zecreylegend *ZecreylegendSession) TotalBlocksCommitted() (uint32, error) {
	return _Zecreylegend.Contract.TotalBlocksCommitted(&_Zecreylegend.CallOpts)
}

// TotalBlocksCommitted is a free data retrieval call binding the contract method 0xfaf4d8cb.
//
// Solidity: function totalBlocksCommitted() view returns(uint32)
func (_Zecreylegend *ZecreylegendCallerSession) TotalBlocksCommitted() (uint32, error) {
	return _Zecreylegend.Contract.TotalBlocksCommitted(&_Zecreylegend.CallOpts)
}

// TotalBlocksVerified is a free data retrieval call binding the contract method 0x2d24006c.
//
// Solidity: function totalBlocksVerified() view returns(uint32)
func (_Zecreylegend *ZecreylegendCaller) TotalBlocksVerified(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "totalBlocksVerified")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TotalBlocksVerified is a free data retrieval call binding the contract method 0x2d24006c.
//
// Solidity: function totalBlocksVerified() view returns(uint32)
func (_Zecreylegend *ZecreylegendSession) TotalBlocksVerified() (uint32, error) {
	return _Zecreylegend.Contract.TotalBlocksVerified(&_Zecreylegend.CallOpts)
}

// TotalBlocksVerified is a free data retrieval call binding the contract method 0x2d24006c.
//
// Solidity: function totalBlocksVerified() view returns(uint32)
func (_Zecreylegend *ZecreylegendCallerSession) TotalBlocksVerified() (uint32, error) {
	return _Zecreylegend.Contract.TotalBlocksVerified(&_Zecreylegend.CallOpts)
}

// TotalOpenPriorityRequests is a free data retrieval call binding the contract method 0xc57b22be.
//
// Solidity: function totalOpenPriorityRequests() view returns(uint64)
func (_Zecreylegend *ZecreylegendCaller) TotalOpenPriorityRequests(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "totalOpenPriorityRequests")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalOpenPriorityRequests is a free data retrieval call binding the contract method 0xc57b22be.
//
// Solidity: function totalOpenPriorityRequests() view returns(uint64)
func (_Zecreylegend *ZecreylegendSession) TotalOpenPriorityRequests() (uint64, error) {
	return _Zecreylegend.Contract.TotalOpenPriorityRequests(&_Zecreylegend.CallOpts)
}

// TotalOpenPriorityRequests is a free data retrieval call binding the contract method 0xc57b22be.
//
// Solidity: function totalOpenPriorityRequests() view returns(uint64)
func (_Zecreylegend *ZecreylegendCallerSession) TotalOpenPriorityRequests() (uint64, error) {
	return _Zecreylegend.Contract.TotalOpenPriorityRequests(&_Zecreylegend.CallOpts)
}

// TotalTokenPairs is a free data retrieval call binding the contract method 0xd0ad718d.
//
// Solidity: function totalTokenPairs() view returns(uint16)
func (_Zecreylegend *ZecreylegendCaller) TotalTokenPairs(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Zecreylegend.contract.Call(opts, &out, "totalTokenPairs")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TotalTokenPairs is a free data retrieval call binding the contract method 0xd0ad718d.
//
// Solidity: function totalTokenPairs() view returns(uint16)
func (_Zecreylegend *ZecreylegendSession) TotalTokenPairs() (uint16, error) {
	return _Zecreylegend.Contract.TotalTokenPairs(&_Zecreylegend.CallOpts)
}

// TotalTokenPairs is a free data retrieval call binding the contract method 0xd0ad718d.
//
// Solidity: function totalTokenPairs() view returns(uint16)
func (_Zecreylegend *ZecreylegendCallerSession) TotalTokenPairs() (uint16, error) {
	return _Zecreylegend.Contract.TotalTokenPairs(&_Zecreylegend.CallOpts)
}

// ActivateDesertMode is a paid mutator transaction binding the contract method 0x22b22256.
//
// Solidity: function activateDesertMode() returns(bool)
func (_Zecreylegend *ZecreylegendTransactor) ActivateDesertMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "activateDesertMode")
}

// ActivateDesertMode is a paid mutator transaction binding the contract method 0x22b22256.
//
// Solidity: function activateDesertMode() returns(bool)
func (_Zecreylegend *ZecreylegendSession) ActivateDesertMode() (*types.Transaction, error) {
	return _Zecreylegend.Contract.ActivateDesertMode(&_Zecreylegend.TransactOpts)
}

// ActivateDesertMode is a paid mutator transaction binding the contract method 0x22b22256.
//
// Solidity: function activateDesertMode() returns(bool)
func (_Zecreylegend *ZecreylegendTransactorSession) ActivateDesertMode() (*types.Transaction, error) {
	return _Zecreylegend.Contract.ActivateDesertMode(&_Zecreylegend.TransactOpts)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x2a54f972.
//
// Solidity: function commitBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,uint32[],uint32,uint16)[] _newBlocksData) returns()
func (_Zecreylegend *ZecreylegendTransactor) CommitBlocks(opts *bind.TransactOpts, _lastCommittedBlockData StorageStoredBlockInfo, _newBlocksData []ZecreyLegendCommitBlockInfo) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "commitBlocks", _lastCommittedBlockData, _newBlocksData)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x2a54f972.
//
// Solidity: function commitBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,uint32[],uint32,uint16)[] _newBlocksData) returns()
func (_Zecreylegend *ZecreylegendSession) CommitBlocks(_lastCommittedBlockData StorageStoredBlockInfo, _newBlocksData []ZecreyLegendCommitBlockInfo) (*types.Transaction, error) {
	return _Zecreylegend.Contract.CommitBlocks(&_Zecreylegend.TransactOpts, _lastCommittedBlockData, _newBlocksData)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x2a54f972.
//
// Solidity: function commitBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,uint32[],uint32,uint16)[] _newBlocksData) returns()
func (_Zecreylegend *ZecreylegendTransactorSession) CommitBlocks(_lastCommittedBlockData StorageStoredBlockInfo, _newBlocksData []ZecreyLegendCommitBlockInfo) (*types.Transaction, error) {
	return _Zecreylegend.Contract.CommitBlocks(&_Zecreylegend.TransactOpts, _lastCommittedBlockData, _newBlocksData)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _tokenA, address _tokenB) returns()
func (_Zecreylegend *ZecreylegendTransactor) CreatePair(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "createPair", _tokenA, _tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _tokenA, address _tokenB) returns()
func (_Zecreylegend *ZecreylegendSession) CreatePair(_tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _Zecreylegend.Contract.CreatePair(&_Zecreylegend.TransactOpts, _tokenA, _tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _tokenA, address _tokenB) returns()
func (_Zecreylegend *ZecreylegendTransactorSession) CreatePair(_tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _Zecreylegend.Contract.CreatePair(&_Zecreylegend.TransactOpts, _tokenA, _tokenB)
}

// CutUpgradeNoticePeriod is a paid mutator transaction binding the contract method 0x3e71e1e7.
//
// Solidity: function cutUpgradeNoticePeriod() returns()
func (_Zecreylegend *ZecreylegendTransactor) CutUpgradeNoticePeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "cutUpgradeNoticePeriod")
}

// CutUpgradeNoticePeriod is a paid mutator transaction binding the contract method 0x3e71e1e7.
//
// Solidity: function cutUpgradeNoticePeriod() returns()
func (_Zecreylegend *ZecreylegendSession) CutUpgradeNoticePeriod() (*types.Transaction, error) {
	return _Zecreylegend.Contract.CutUpgradeNoticePeriod(&_Zecreylegend.TransactOpts)
}

// CutUpgradeNoticePeriod is a paid mutator transaction binding the contract method 0x3e71e1e7.
//
// Solidity: function cutUpgradeNoticePeriod() returns()
func (_Zecreylegend *ZecreylegendTransactorSession) CutUpgradeNoticePeriod() (*types.Transaction, error) {
	return _Zecreylegend.Contract.CutUpgradeNoticePeriod(&_Zecreylegend.TransactOpts)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x1caf5d25.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, string _accountName) returns()
func (_Zecreylegend *ZecreylegendTransactor) DepositBEP20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _accountName string) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "depositBEP20", _token, _amount, _accountName)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x1caf5d25.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, string _accountName) returns()
func (_Zecreylegend *ZecreylegendSession) DepositBEP20(_token common.Address, _amount *big.Int, _accountName string) (*types.Transaction, error) {
	return _Zecreylegend.Contract.DepositBEP20(&_Zecreylegend.TransactOpts, _token, _amount, _accountName)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x1caf5d25.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, string _accountName) returns()
func (_Zecreylegend *ZecreylegendTransactorSession) DepositBEP20(_token common.Address, _amount *big.Int, _accountName string) (*types.Transaction, error) {
	return _Zecreylegend.Contract.DepositBEP20(&_Zecreylegend.TransactOpts, _token, _amount, _accountName)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x51344683.
//
// Solidity: function depositBNB(string _accountName) payable returns()
func (_Zecreylegend *ZecreylegendTransactor) DepositBNB(opts *bind.TransactOpts, _accountName string) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "depositBNB", _accountName)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x51344683.
//
// Solidity: function depositBNB(string _accountName) payable returns()
func (_Zecreylegend *ZecreylegendSession) DepositBNB(_accountName string) (*types.Transaction, error) {
	return _Zecreylegend.Contract.DepositBNB(&_Zecreylegend.TransactOpts, _accountName)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x51344683.
//
// Solidity: function depositBNB(string _accountName) payable returns()
func (_Zecreylegend *ZecreylegendTransactorSession) DepositBNB(_accountName string) (*types.Transaction, error) {
	return _Zecreylegend.Contract.DepositBNB(&_Zecreylegend.TransactOpts, _accountName)
}

// DepositNft is a paid mutator transaction binding the contract method 0xfb99514b.
//
// Solidity: function depositNft(string _accountName, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_Zecreylegend *ZecreylegendTransactor) DepositNft(opts *bind.TransactOpts, _accountName string, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "depositNft", _accountName, _nftL1Address, _nftL1TokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xfb99514b.
//
// Solidity: function depositNft(string _accountName, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_Zecreylegend *ZecreylegendSession) DepositNft(_accountName string, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _Zecreylegend.Contract.DepositNft(&_Zecreylegend.TransactOpts, _accountName, _nftL1Address, _nftL1TokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xfb99514b.
//
// Solidity: function depositNft(string _accountName, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_Zecreylegend *ZecreylegendTransactorSession) DepositNft(_accountName string, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _Zecreylegend.Contract.DepositNft(&_Zecreylegend.TransactOpts, _accountName, _nftL1Address, _nftL1TokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_Zecreylegend *ZecreylegendTransactor) Initialize(opts *bind.TransactOpts, initializationParameters []byte) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "initialize", initializationParameters)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_Zecreylegend *ZecreylegendSession) Initialize(initializationParameters []byte) (*types.Transaction, error) {
	return _Zecreylegend.Contract.Initialize(&_Zecreylegend.TransactOpts, initializationParameters)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_Zecreylegend *ZecreylegendTransactorSession) Initialize(initializationParameters []byte) (*types.Transaction, error) {
	return _Zecreylegend.Contract.Initialize(&_Zecreylegend.TransactOpts, initializationParameters)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Zecreylegend *ZecreylegendTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Zecreylegend *ZecreylegendSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Zecreylegend.Contract.OnERC721Received(&_Zecreylegend.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Zecreylegend *ZecreylegendTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Zecreylegend.Contract.OnERC721Received(&_Zecreylegend.TransactOpts, operator, from, tokenId, data)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0x3fdeb67d.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _zkbnbPubKeyX, bytes32 _zkbnbPubKeyY) payable returns()
func (_Zecreylegend *ZecreylegendTransactor) RegisterZNS(opts *bind.TransactOpts, _name string, _owner common.Address, _zkbnbPubKeyX [32]byte, _zkbnbPubKeyY [32]byte) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "registerZNS", _name, _owner, _zkbnbPubKeyX, _zkbnbPubKeyY)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0x3fdeb67d.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _zkbnbPubKeyX, bytes32 _zkbnbPubKeyY) payable returns()
func (_Zecreylegend *ZecreylegendSession) RegisterZNS(_name string, _owner common.Address, _zkbnbPubKeyX [32]byte, _zkbnbPubKeyY [32]byte) (*types.Transaction, error) {
	return _Zecreylegend.Contract.RegisterZNS(&_Zecreylegend.TransactOpts, _name, _owner, _zkbnbPubKeyX, _zkbnbPubKeyY)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0x3fdeb67d.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _zkbnbPubKeyX, bytes32 _zkbnbPubKeyY) payable returns()
func (_Zecreylegend *ZecreylegendTransactorSession) RegisterZNS(_name string, _owner common.Address, _zkbnbPubKeyX [32]byte, _zkbnbPubKeyY [32]byte) (*types.Transaction, error) {
	return _Zecreylegend.Contract.RegisterZNS(&_Zecreylegend.TransactOpts, _name, _owner, _zkbnbPubKeyX, _zkbnbPubKeyY)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x8da64c7f.
//
// Solidity: function requestFullExit(string _accountName, address _asset) returns()
func (_Zecreylegend *ZecreylegendTransactor) RequestFullExit(opts *bind.TransactOpts, _accountName string, _asset common.Address) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "requestFullExit", _accountName, _asset)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x8da64c7f.
//
// Solidity: function requestFullExit(string _accountName, address _asset) returns()
func (_Zecreylegend *ZecreylegendSession) RequestFullExit(_accountName string, _asset common.Address) (*types.Transaction, error) {
	return _Zecreylegend.Contract.RequestFullExit(&_Zecreylegend.TransactOpts, _accountName, _asset)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x8da64c7f.
//
// Solidity: function requestFullExit(string _accountName, address _asset) returns()
func (_Zecreylegend *ZecreylegendTransactorSession) RequestFullExit(_accountName string, _asset common.Address) (*types.Transaction, error) {
	return _Zecreylegend.Contract.RequestFullExit(&_Zecreylegend.TransactOpts, _accountName, _asset)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x1bd24317.
//
// Solidity: function requestFullExitNft(string _accountName, uint32 _nftIndex) returns()
func (_Zecreylegend *ZecreylegendTransactor) RequestFullExitNft(opts *bind.TransactOpts, _accountName string, _nftIndex uint32) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "requestFullExitNft", _accountName, _nftIndex)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x1bd24317.
//
// Solidity: function requestFullExitNft(string _accountName, uint32 _nftIndex) returns()
func (_Zecreylegend *ZecreylegendSession) RequestFullExitNft(_accountName string, _nftIndex uint32) (*types.Transaction, error) {
	return _Zecreylegend.Contract.RequestFullExitNft(&_Zecreylegend.TransactOpts, _accountName, _nftIndex)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x1bd24317.
//
// Solidity: function requestFullExitNft(string _accountName, uint32 _nftIndex) returns()
func (_Zecreylegend *ZecreylegendTransactorSession) RequestFullExitNft(_accountName string, _nftIndex uint32) (*types.Transaction, error) {
	return _Zecreylegend.Contract.RequestFullExitNft(&_Zecreylegend.TransactOpts, _accountName, _nftIndex)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0x97a2eabe.
//
// Solidity: function revertBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _blocksToRevert) returns()
func (_Zecreylegend *ZecreylegendTransactor) RevertBlocks(opts *bind.TransactOpts, _blocksToRevert []StorageStoredBlockInfo) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "revertBlocks", _blocksToRevert)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0x97a2eabe.
//
// Solidity: function revertBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _blocksToRevert) returns()
func (_Zecreylegend *ZecreylegendSession) RevertBlocks(_blocksToRevert []StorageStoredBlockInfo) (*types.Transaction, error) {
	return _Zecreylegend.Contract.RevertBlocks(&_Zecreylegend.TransactOpts, _blocksToRevert)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0x97a2eabe.
//
// Solidity: function revertBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _blocksToRevert) returns()
func (_Zecreylegend *ZecreylegendTransactorSession) RevertBlocks(_blocksToRevert []StorageStoredBlockInfo) (*types.Transaction, error) {
	return _Zecreylegend.Contract.RevertBlocks(&_Zecreylegend.TransactOpts, _blocksToRevert)
}

// SetDefaultNFTFactory is a paid mutator transaction binding the contract method 0xce09e20d.
//
// Solidity: function setDefaultNFTFactory(address _factory) returns()
func (_Zecreylegend *ZecreylegendTransactor) SetDefaultNFTFactory(opts *bind.TransactOpts, _factory common.Address) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "setDefaultNFTFactory", _factory)
}

// SetDefaultNFTFactory is a paid mutator transaction binding the contract method 0xce09e20d.
//
// Solidity: function setDefaultNFTFactory(address _factory) returns()
func (_Zecreylegend *ZecreylegendSession) SetDefaultNFTFactory(_factory common.Address) (*types.Transaction, error) {
	return _Zecreylegend.Contract.SetDefaultNFTFactory(&_Zecreylegend.TransactOpts, _factory)
}

// SetDefaultNFTFactory is a paid mutator transaction binding the contract method 0xce09e20d.
//
// Solidity: function setDefaultNFTFactory(address _factory) returns()
func (_Zecreylegend *ZecreylegendTransactorSession) SetDefaultNFTFactory(_factory common.Address) (*types.Transaction, error) {
	return _Zecreylegend.Contract.SetDefaultNFTFactory(&_Zecreylegend.TransactOpts, _factory)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x68809a23.
//
// Solidity: function transferERC20(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128 withdrawnAmount)
func (_Zecreylegend *ZecreylegendTransactor) TransferERC20(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "transferERC20", _token, _to, _amount, _maxAmount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x68809a23.
//
// Solidity: function transferERC20(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128 withdrawnAmount)
func (_Zecreylegend *ZecreylegendSession) TransferERC20(_token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _Zecreylegend.Contract.TransferERC20(&_Zecreylegend.TransactOpts, _token, _to, _amount, _maxAmount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x68809a23.
//
// Solidity: function transferERC20(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128 withdrawnAmount)
func (_Zecreylegend *ZecreylegendTransactorSession) TransferERC20(_token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _Zecreylegend.Contract.TransferERC20(&_Zecreylegend.TransactOpts, _token, _to, _amount, _maxAmount)
}

// UpdatePairRate is a paid mutator transaction binding the contract method 0x13a05e23.
//
// Solidity: function updatePairRate((address,address,uint16,uint32,uint16) _pairInfo) returns()
func (_Zecreylegend *ZecreylegendTransactor) UpdatePairRate(opts *bind.TransactOpts, _pairInfo ZecreyLegendPairInfo) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "updatePairRate", _pairInfo)
}

// UpdatePairRate is a paid mutator transaction binding the contract method 0x13a05e23.
//
// Solidity: function updatePairRate((address,address,uint16,uint32,uint16) _pairInfo) returns()
func (_Zecreylegend *ZecreylegendSession) UpdatePairRate(_pairInfo ZecreyLegendPairInfo) (*types.Transaction, error) {
	return _Zecreylegend.Contract.UpdatePairRate(&_Zecreylegend.TransactOpts, _pairInfo)
}

// UpdatePairRate is a paid mutator transaction binding the contract method 0x13a05e23.
//
// Solidity: function updatePairRate((address,address,uint16,uint32,uint16) _pairInfo) returns()
func (_Zecreylegend *ZecreylegendTransactorSession) UpdatePairRate(_pairInfo ZecreyLegendPairInfo) (*types.Transaction, error) {
	return _Zecreylegend.Contract.UpdatePairRate(&_Zecreylegend.TransactOpts, _pairInfo)
}

// UpdateZkBNBVerifier is a paid mutator transaction binding the contract method 0x6f88b156.
//
// Solidity: function updateZkBNBVerifier(address _newVerifierAddress) returns()
func (_Zecreylegend *ZecreylegendTransactor) UpdateZkBNBVerifier(opts *bind.TransactOpts, _newVerifierAddress common.Address) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "updateZkBNBVerifier", _newVerifierAddress)
}

// UpdateZkBNBVerifier is a paid mutator transaction binding the contract method 0x6f88b156.
//
// Solidity: function updateZkBNBVerifier(address _newVerifierAddress) returns()
func (_Zecreylegend *ZecreylegendSession) UpdateZkBNBVerifier(_newVerifierAddress common.Address) (*types.Transaction, error) {
	return _Zecreylegend.Contract.UpdateZkBNBVerifier(&_Zecreylegend.TransactOpts, _newVerifierAddress)
}

// UpdateZkBNBVerifier is a paid mutator transaction binding the contract method 0x6f88b156.
//
// Solidity: function updateZkBNBVerifier(address _newVerifierAddress) returns()
func (_Zecreylegend *ZecreylegendTransactorSession) UpdateZkBNBVerifier(_newVerifierAddress common.Address) (*types.Transaction, error) {
	return _Zecreylegend.Contract.UpdateZkBNBVerifier(&_Zecreylegend.TransactOpts, _newVerifierAddress)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_Zecreylegend *ZecreylegendTransactor) Upgrade(opts *bind.TransactOpts, upgradeParameters []byte) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "upgrade", upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_Zecreylegend *ZecreylegendSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _Zecreylegend.Contract.Upgrade(&_Zecreylegend.TransactOpts, upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_Zecreylegend *ZecreylegendTransactorSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _Zecreylegend.Contract.Upgrade(&_Zecreylegend.TransactOpts, upgradeParameters)
}

// UpgradeCanceled is a paid mutator transaction binding the contract method 0x871b8ff1.
//
// Solidity: function upgradeCanceled() returns()
func (_Zecreylegend *ZecreylegendTransactor) UpgradeCanceled(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "upgradeCanceled")
}

// UpgradeCanceled is a paid mutator transaction binding the contract method 0x871b8ff1.
//
// Solidity: function upgradeCanceled() returns()
func (_Zecreylegend *ZecreylegendSession) UpgradeCanceled() (*types.Transaction, error) {
	return _Zecreylegend.Contract.UpgradeCanceled(&_Zecreylegend.TransactOpts)
}

// UpgradeCanceled is a paid mutator transaction binding the contract method 0x871b8ff1.
//
// Solidity: function upgradeCanceled() returns()
func (_Zecreylegend *ZecreylegendTransactorSession) UpgradeCanceled() (*types.Transaction, error) {
	return _Zecreylegend.Contract.UpgradeCanceled(&_Zecreylegend.TransactOpts)
}

// UpgradeFinishes is a paid mutator transaction binding the contract method 0xb269b9ae.
//
// Solidity: function upgradeFinishes() returns()
func (_Zecreylegend *ZecreylegendTransactor) UpgradeFinishes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "upgradeFinishes")
}

// UpgradeFinishes is a paid mutator transaction binding the contract method 0xb269b9ae.
//
// Solidity: function upgradeFinishes() returns()
func (_Zecreylegend *ZecreylegendSession) UpgradeFinishes() (*types.Transaction, error) {
	return _Zecreylegend.Contract.UpgradeFinishes(&_Zecreylegend.TransactOpts)
}

// UpgradeFinishes is a paid mutator transaction binding the contract method 0xb269b9ae.
//
// Solidity: function upgradeFinishes() returns()
func (_Zecreylegend *ZecreylegendTransactorSession) UpgradeFinishes() (*types.Transaction, error) {
	return _Zecreylegend.Contract.UpgradeFinishes(&_Zecreylegend.TransactOpts)
}

// UpgradeNoticePeriodStarted is a paid mutator transaction binding the contract method 0x3b154b73.
//
// Solidity: function upgradeNoticePeriodStarted() returns()
func (_Zecreylegend *ZecreylegendTransactor) UpgradeNoticePeriodStarted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "upgradeNoticePeriodStarted")
}

// UpgradeNoticePeriodStarted is a paid mutator transaction binding the contract method 0x3b154b73.
//
// Solidity: function upgradeNoticePeriodStarted() returns()
func (_Zecreylegend *ZecreylegendSession) UpgradeNoticePeriodStarted() (*types.Transaction, error) {
	return _Zecreylegend.Contract.UpgradeNoticePeriodStarted(&_Zecreylegend.TransactOpts)
}

// UpgradeNoticePeriodStarted is a paid mutator transaction binding the contract method 0x3b154b73.
//
// Solidity: function upgradeNoticePeriodStarted() returns()
func (_Zecreylegend *ZecreylegendTransactorSession) UpgradeNoticePeriodStarted() (*types.Transaction, error) {
	return _Zecreylegend.Contract.UpgradeNoticePeriodStarted(&_Zecreylegend.TransactOpts)
}

// UpgradePreparationStarted is a paid mutator transaction binding the contract method 0x78b91e70.
//
// Solidity: function upgradePreparationStarted() returns()
func (_Zecreylegend *ZecreylegendTransactor) UpgradePreparationStarted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "upgradePreparationStarted")
}

// UpgradePreparationStarted is a paid mutator transaction binding the contract method 0x78b91e70.
//
// Solidity: function upgradePreparationStarted() returns()
func (_Zecreylegend *ZecreylegendSession) UpgradePreparationStarted() (*types.Transaction, error) {
	return _Zecreylegend.Contract.UpgradePreparationStarted(&_Zecreylegend.TransactOpts)
}

// UpgradePreparationStarted is a paid mutator transaction binding the contract method 0x78b91e70.
//
// Solidity: function upgradePreparationStarted() returns()
func (_Zecreylegend *ZecreylegendTransactorSession) UpgradePreparationStarted() (*types.Transaction, error) {
	return _Zecreylegend.Contract.UpgradePreparationStarted(&_Zecreylegend.TransactOpts)
}

// VerifyAndExecuteBlocks is a paid mutator transaction binding the contract method 0xf9ea2e71.
//
// Solidity: function verifyAndExecuteBlocks(((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocks, uint256[] _proofs) returns()
func (_Zecreylegend *ZecreylegendTransactor) VerifyAndExecuteBlocks(opts *bind.TransactOpts, _blocks []ZecreyLegendVerifyAndExecuteBlockInfo, _proofs []*big.Int) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "verifyAndExecuteBlocks", _blocks, _proofs)
}

// VerifyAndExecuteBlocks is a paid mutator transaction binding the contract method 0xf9ea2e71.
//
// Solidity: function verifyAndExecuteBlocks(((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocks, uint256[] _proofs) returns()
func (_Zecreylegend *ZecreylegendSession) VerifyAndExecuteBlocks(_blocks []ZecreyLegendVerifyAndExecuteBlockInfo, _proofs []*big.Int) (*types.Transaction, error) {
	return _Zecreylegend.Contract.VerifyAndExecuteBlocks(&_Zecreylegend.TransactOpts, _blocks, _proofs)
}

// VerifyAndExecuteBlocks is a paid mutator transaction binding the contract method 0xf9ea2e71.
//
// Solidity: function verifyAndExecuteBlocks(((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocks, uint256[] _proofs) returns()
func (_Zecreylegend *ZecreylegendTransactorSession) VerifyAndExecuteBlocks(_blocks []ZecreyLegendVerifyAndExecuteBlockInfo, _proofs []*big.Int) (*types.Transaction, error) {
	return _Zecreylegend.Contract.VerifyAndExecuteBlocks(&_Zecreylegend.TransactOpts, _blocks, _proofs)
}

// WithdrawPendingBalance is a paid mutator transaction binding the contract method 0xd514da50.
//
// Solidity: function withdrawPendingBalance(address _owner, address _token, uint128 _amount) returns()
func (_Zecreylegend *ZecreylegendTransactor) WithdrawPendingBalance(opts *bind.TransactOpts, _owner common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "withdrawPendingBalance", _owner, _token, _amount)
}

// WithdrawPendingBalance is a paid mutator transaction binding the contract method 0xd514da50.
//
// Solidity: function withdrawPendingBalance(address _owner, address _token, uint128 _amount) returns()
func (_Zecreylegend *ZecreylegendSession) WithdrawPendingBalance(_owner common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zecreylegend.Contract.WithdrawPendingBalance(&_Zecreylegend.TransactOpts, _owner, _token, _amount)
}

// WithdrawPendingBalance is a paid mutator transaction binding the contract method 0xd514da50.
//
// Solidity: function withdrawPendingBalance(address _owner, address _token, uint128 _amount) returns()
func (_Zecreylegend *ZecreylegendTransactorSession) WithdrawPendingBalance(_owner common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zecreylegend.Contract.WithdrawPendingBalance(&_Zecreylegend.TransactOpts, _owner, _token, _amount)
}

// WithdrawPendingNFTBalance is a paid mutator transaction binding the contract method 0x7ce1017d.
//
// Solidity: function withdrawPendingNFTBalance(uint40 _nftIndex) returns()
func (_Zecreylegend *ZecreylegendTransactor) WithdrawPendingNFTBalance(opts *bind.TransactOpts, _nftIndex *big.Int) (*types.Transaction, error) {
	return _Zecreylegend.contract.Transact(opts, "withdrawPendingNFTBalance", _nftIndex)
}

// WithdrawPendingNFTBalance is a paid mutator transaction binding the contract method 0x7ce1017d.
//
// Solidity: function withdrawPendingNFTBalance(uint40 _nftIndex) returns()
func (_Zecreylegend *ZecreylegendSession) WithdrawPendingNFTBalance(_nftIndex *big.Int) (*types.Transaction, error) {
	return _Zecreylegend.Contract.WithdrawPendingNFTBalance(&_Zecreylegend.TransactOpts, _nftIndex)
}

// WithdrawPendingNFTBalance is a paid mutator transaction binding the contract method 0x7ce1017d.
//
// Solidity: function withdrawPendingNFTBalance(uint40 _nftIndex) returns()
func (_Zecreylegend *ZecreylegendTransactorSession) WithdrawPendingNFTBalance(_nftIndex *big.Int) (*types.Transaction, error) {
	return _Zecreylegend.Contract.WithdrawPendingNFTBalance(&_Zecreylegend.TransactOpts, _nftIndex)
}

// ZecreylegendBlockCommitIterator is returned from FilterBlockCommit and is used to iterate over the raw logs and unpacked data for BlockCommit events raised by the Zecreylegend contract.
type ZecreylegendBlockCommitIterator struct {
	Event *ZecreylegendBlockCommit // Event containing the contract specifics and raw log

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
func (it *ZecreylegendBlockCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendBlockCommit)
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
		it.Event = new(ZecreylegendBlockCommit)
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
func (it *ZecreylegendBlockCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendBlockCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendBlockCommit represents a BlockCommit event raised by the Zecreylegend contract.
type ZecreylegendBlockCommit struct {
	BlockNumber uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockCommit is a free log retrieval operation binding the contract event 0x81a92942d0f9c33b897a438384c9c3d88be397776138efa3ba1a4fc8b6268424.
//
// Solidity: event BlockCommit(uint32 blockNumber)
func (_Zecreylegend *ZecreylegendFilterer) FilterBlockCommit(opts *bind.FilterOpts) (*ZecreylegendBlockCommitIterator, error) {

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "BlockCommit")
	if err != nil {
		return nil, err
	}
	return &ZecreylegendBlockCommitIterator{contract: _Zecreylegend.contract, event: "BlockCommit", logs: logs, sub: sub}, nil
}

// WatchBlockCommit is a free log subscription operation binding the contract event 0x81a92942d0f9c33b897a438384c9c3d88be397776138efa3ba1a4fc8b6268424.
//
// Solidity: event BlockCommit(uint32 blockNumber)
func (_Zecreylegend *ZecreylegendFilterer) WatchBlockCommit(opts *bind.WatchOpts, sink chan<- *ZecreylegendBlockCommit) (event.Subscription, error) {

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "BlockCommit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendBlockCommit)
				if err := _Zecreylegend.contract.UnpackLog(event, "BlockCommit", log); err != nil {
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

// ParseBlockCommit is a log parse operation binding the contract event 0x81a92942d0f9c33b897a438384c9c3d88be397776138efa3ba1a4fc8b6268424.
//
// Solidity: event BlockCommit(uint32 blockNumber)
func (_Zecreylegend *ZecreylegendFilterer) ParseBlockCommit(log types.Log) (*ZecreylegendBlockCommit, error) {
	event := new(ZecreylegendBlockCommit)
	if err := _Zecreylegend.contract.UnpackLog(event, "BlockCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreylegendBlockVerificationIterator is returned from FilterBlockVerification and is used to iterate over the raw logs and unpacked data for BlockVerification events raised by the Zecreylegend contract.
type ZecreylegendBlockVerificationIterator struct {
	Event *ZecreylegendBlockVerification // Event containing the contract specifics and raw log

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
func (it *ZecreylegendBlockVerificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendBlockVerification)
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
		it.Event = new(ZecreylegendBlockVerification)
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
func (it *ZecreylegendBlockVerificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendBlockVerificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendBlockVerification represents a BlockVerification event raised by the Zecreylegend contract.
type ZecreylegendBlockVerification struct {
	BlockNumber uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockVerification is a free log retrieval operation binding the contract event 0x0cdbd8bd7813095001c5fe7917bd69d834dc01db7c1dfcf52ca135bd20384413.
//
// Solidity: event BlockVerification(uint32 blockNumber)
func (_Zecreylegend *ZecreylegendFilterer) FilterBlockVerification(opts *bind.FilterOpts) (*ZecreylegendBlockVerificationIterator, error) {

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "BlockVerification")
	if err != nil {
		return nil, err
	}
	return &ZecreylegendBlockVerificationIterator{contract: _Zecreylegend.contract, event: "BlockVerification", logs: logs, sub: sub}, nil
}

// WatchBlockVerification is a free log subscription operation binding the contract event 0x0cdbd8bd7813095001c5fe7917bd69d834dc01db7c1dfcf52ca135bd20384413.
//
// Solidity: event BlockVerification(uint32 blockNumber)
func (_Zecreylegend *ZecreylegendFilterer) WatchBlockVerification(opts *bind.WatchOpts, sink chan<- *ZecreylegendBlockVerification) (event.Subscription, error) {

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "BlockVerification")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendBlockVerification)
				if err := _Zecreylegend.contract.UnpackLog(event, "BlockVerification", log); err != nil {
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

// ParseBlockVerification is a log parse operation binding the contract event 0x0cdbd8bd7813095001c5fe7917bd69d834dc01db7c1dfcf52ca135bd20384413.
//
// Solidity: event BlockVerification(uint32 blockNumber)
func (_Zecreylegend *ZecreylegendFilterer) ParseBlockVerification(log types.Log) (*ZecreylegendBlockVerification, error) {
	event := new(ZecreylegendBlockVerification)
	if err := _Zecreylegend.contract.UnpackLog(event, "BlockVerification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreylegendBlocksRevertIterator is returned from FilterBlocksRevert and is used to iterate over the raw logs and unpacked data for BlocksRevert events raised by the Zecreylegend contract.
type ZecreylegendBlocksRevertIterator struct {
	Event *ZecreylegendBlocksRevert // Event containing the contract specifics and raw log

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
func (it *ZecreylegendBlocksRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendBlocksRevert)
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
		it.Event = new(ZecreylegendBlocksRevert)
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
func (it *ZecreylegendBlocksRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendBlocksRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendBlocksRevert represents a BlocksRevert event raised by the Zecreylegend contract.
type ZecreylegendBlocksRevert struct {
	TotalBlocksVerified  uint32
	TotalBlocksCommitted uint32
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBlocksRevert is a free log retrieval operation binding the contract event 0x6f3a8259cce1ea2680115053d21c971aa1764295a45850f520525f2bfdf3c9d3.
//
// Solidity: event BlocksRevert(uint32 totalBlocksVerified, uint32 totalBlocksCommitted)
func (_Zecreylegend *ZecreylegendFilterer) FilterBlocksRevert(opts *bind.FilterOpts) (*ZecreylegendBlocksRevertIterator, error) {

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "BlocksRevert")
	if err != nil {
		return nil, err
	}
	return &ZecreylegendBlocksRevertIterator{contract: _Zecreylegend.contract, event: "BlocksRevert", logs: logs, sub: sub}, nil
}

// WatchBlocksRevert is a free log subscription operation binding the contract event 0x6f3a8259cce1ea2680115053d21c971aa1764295a45850f520525f2bfdf3c9d3.
//
// Solidity: event BlocksRevert(uint32 totalBlocksVerified, uint32 totalBlocksCommitted)
func (_Zecreylegend *ZecreylegendFilterer) WatchBlocksRevert(opts *bind.WatchOpts, sink chan<- *ZecreylegendBlocksRevert) (event.Subscription, error) {

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "BlocksRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendBlocksRevert)
				if err := _Zecreylegend.contract.UnpackLog(event, "BlocksRevert", log); err != nil {
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

// ParseBlocksRevert is a log parse operation binding the contract event 0x6f3a8259cce1ea2680115053d21c971aa1764295a45850f520525f2bfdf3c9d3.
//
// Solidity: event BlocksRevert(uint32 totalBlocksVerified, uint32 totalBlocksCommitted)
func (_Zecreylegend *ZecreylegendFilterer) ParseBlocksRevert(log types.Log) (*ZecreylegendBlocksRevert, error) {
	event := new(ZecreylegendBlocksRevert)
	if err := _Zecreylegend.contract.UnpackLog(event, "BlocksRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreylegendCreateTokenPairIterator is returned from FilterCreateTokenPair and is used to iterate over the raw logs and unpacked data for CreateTokenPair events raised by the Zecreylegend contract.
type ZecreylegendCreateTokenPairIterator struct {
	Event *ZecreylegendCreateTokenPair // Event containing the contract specifics and raw log

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
func (it *ZecreylegendCreateTokenPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendCreateTokenPair)
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
		it.Event = new(ZecreylegendCreateTokenPair)
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
func (it *ZecreylegendCreateTokenPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendCreateTokenPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendCreateTokenPair represents a CreateTokenPair event raised by the Zecreylegend contract.
type ZecreylegendCreateTokenPair struct {
	PairIndex            uint16
	Asset0Id             uint16
	Asset1Id             uint16
	FeeRate              uint16
	TreasuryAccountIndex uint32
	TreasuryRate         uint16
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterCreateTokenPair is a free log retrieval operation binding the contract event 0x82e21fd4ced46b510ee1bf5a34eea3dc34aa1f5460f53a25dd22828cde30e087.
//
// Solidity: event CreateTokenPair(uint16 pairIndex, uint16 asset0Id, uint16 asset1Id, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_Zecreylegend *ZecreylegendFilterer) FilterCreateTokenPair(opts *bind.FilterOpts) (*ZecreylegendCreateTokenPairIterator, error) {

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "CreateTokenPair")
	if err != nil {
		return nil, err
	}
	return &ZecreylegendCreateTokenPairIterator{contract: _Zecreylegend.contract, event: "CreateTokenPair", logs: logs, sub: sub}, nil
}

// WatchCreateTokenPair is a free log subscription operation binding the contract event 0x82e21fd4ced46b510ee1bf5a34eea3dc34aa1f5460f53a25dd22828cde30e087.
//
// Solidity: event CreateTokenPair(uint16 pairIndex, uint16 asset0Id, uint16 asset1Id, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_Zecreylegend *ZecreylegendFilterer) WatchCreateTokenPair(opts *bind.WatchOpts, sink chan<- *ZecreylegendCreateTokenPair) (event.Subscription, error) {

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "CreateTokenPair")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendCreateTokenPair)
				if err := _Zecreylegend.contract.UnpackLog(event, "CreateTokenPair", log); err != nil {
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

// ParseCreateTokenPair is a log parse operation binding the contract event 0x82e21fd4ced46b510ee1bf5a34eea3dc34aa1f5460f53a25dd22828cde30e087.
//
// Solidity: event CreateTokenPair(uint16 pairIndex, uint16 asset0Id, uint16 asset1Id, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_Zecreylegend *ZecreylegendFilterer) ParseCreateTokenPair(log types.Log) (*ZecreylegendCreateTokenPair, error) {
	event := new(ZecreylegendCreateTokenPair)
	if err := _Zecreylegend.contract.UnpackLog(event, "CreateTokenPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreylegendDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Zecreylegend contract.
type ZecreylegendDepositIterator struct {
	Event *ZecreylegendDeposit // Event containing the contract specifics and raw log

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
func (it *ZecreylegendDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendDeposit)
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
		it.Event = new(ZecreylegendDeposit)
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
func (it *ZecreylegendDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendDeposit represents a Deposit event raised by the Zecreylegend contract.
type ZecreylegendDeposit struct {
	AssetId     uint16
	AccountName [32]byte
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xaa46d46658f805449ac7eaaf11d3eaebb6f508930128d276dfefcc8dc02a13c7.
//
// Solidity: event Deposit(uint16 assetId, bytes32 accountName, uint128 amount)
func (_Zecreylegend *ZecreylegendFilterer) FilterDeposit(opts *bind.FilterOpts) (*ZecreylegendDepositIterator, error) {

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &ZecreylegendDepositIterator{contract: _Zecreylegend.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xaa46d46658f805449ac7eaaf11d3eaebb6f508930128d276dfefcc8dc02a13c7.
//
// Solidity: event Deposit(uint16 assetId, bytes32 accountName, uint128 amount)
func (_Zecreylegend *ZecreylegendFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ZecreylegendDeposit) (event.Subscription, error) {

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendDeposit)
				if err := _Zecreylegend.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xaa46d46658f805449ac7eaaf11d3eaebb6f508930128d276dfefcc8dc02a13c7.
//
// Solidity: event Deposit(uint16 assetId, bytes32 accountName, uint128 amount)
func (_Zecreylegend *ZecreylegendFilterer) ParseDeposit(log types.Log) (*ZecreylegendDeposit, error) {
	event := new(ZecreylegendDeposit)
	if err := _Zecreylegend.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreylegendDepositCommitIterator is returned from FilterDepositCommit and is used to iterate over the raw logs and unpacked data for DepositCommit events raised by the Zecreylegend contract.
type ZecreylegendDepositCommitIterator struct {
	Event *ZecreylegendDepositCommit // Event containing the contract specifics and raw log

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
func (it *ZecreylegendDepositCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendDepositCommit)
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
		it.Event = new(ZecreylegendDepositCommit)
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
func (it *ZecreylegendDepositCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendDepositCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendDepositCommit represents a DepositCommit event raised by the Zecreylegend contract.
type ZecreylegendDepositCommit struct {
	ZkBNBBlockNumber uint32
	AccountIndex     uint32
	AccountName      [32]byte
	AssetId          uint16
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDepositCommit is a free log retrieval operation binding the contract event 0x9b3e77f05cee61edb7f3e0e85d71a29e89cb18cb4236a8b8cc3e4d66640d8f71.
//
// Solidity: event DepositCommit(uint32 indexed zkbnbBlockNumber, uint32 indexed accountIndex, bytes32 accountName, uint16 indexed assetId, uint128 amount)
func (_Zecreylegend *ZecreylegendFilterer) FilterDepositCommit(opts *bind.FilterOpts, zkbnbBlockNumber []uint32, accountIndex []uint32, assetId []uint16) (*ZecreylegendDepositCommitIterator, error) {

	var zkbnbBlockNumberRule []interface{}
	for _, zkbnbBlockNumberItem := range zkbnbBlockNumber {
		zkbnbBlockNumberRule = append(zkbnbBlockNumberRule, zkbnbBlockNumberItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "DepositCommit", zkbnbBlockNumberRule, accountIndexRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &ZecreylegendDepositCommitIterator{contract: _Zecreylegend.contract, event: "DepositCommit", logs: logs, sub: sub}, nil
}

// WatchDepositCommit is a free log subscription operation binding the contract event 0x9b3e77f05cee61edb7f3e0e85d71a29e89cb18cb4236a8b8cc3e4d66640d8f71.
//
// Solidity: event DepositCommit(uint32 indexed zkbnbBlockNumber, uint32 indexed accountIndex, bytes32 accountName, uint16 indexed assetId, uint128 amount)
func (_Zecreylegend *ZecreylegendFilterer) WatchDepositCommit(opts *bind.WatchOpts, sink chan<- *ZecreylegendDepositCommit, zkbnbBlockNumber []uint32, accountIndex []uint32, assetId []uint16) (event.Subscription, error) {

	var zkbnbBlockNumberRule []interface{}
	for _, zkbnbBlockNumberItem := range zkbnbBlockNumber {
		zkbnbBlockNumberRule = append(zkbnbBlockNumberRule, zkbnbBlockNumberItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "DepositCommit", zkbnbBlockNumberRule, accountIndexRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendDepositCommit)
				if err := _Zecreylegend.contract.UnpackLog(event, "DepositCommit", log); err != nil {
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

// ParseDepositCommit is a log parse operation binding the contract event 0x9b3e77f05cee61edb7f3e0e85d71a29e89cb18cb4236a8b8cc3e4d66640d8f71.
//
// Solidity: event DepositCommit(uint32 indexed zkbnbBlockNumber, uint32 indexed accountIndex, bytes32 accountName, uint16 indexed assetId, uint128 amount)
func (_Zecreylegend *ZecreylegendFilterer) ParseDepositCommit(log types.Log) (*ZecreylegendDepositCommit, error) {
	event := new(ZecreylegendDepositCommit)
	if err := _Zecreylegend.contract.UnpackLog(event, "DepositCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreylegendDepositNftIterator is returned from FilterDepositNft and is used to iterate over the raw logs and unpacked data for DepositNft events raised by the Zecreylegend contract.
type ZecreylegendDepositNftIterator struct {
	Event *ZecreylegendDepositNft // Event containing the contract specifics and raw log

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
func (it *ZecreylegendDepositNftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendDepositNft)
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
		it.Event = new(ZecreylegendDepositNft)
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
func (it *ZecreylegendDepositNftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendDepositNftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendDepositNft represents a DepositNft event raised by the Zecreylegend contract.
type ZecreylegendDepositNft struct {
	AccountNameHash     [32]byte
	NftContentHash      [32]byte
	TokenAddress        common.Address
	NftTokenId          *big.Int
	CreatorTreasuryRate uint16
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDepositNft is a free log retrieval operation binding the contract event 0x42f9fbd77faf066fde386943aaafcc841fba6b755dc5da7939a046cbc493c24f.
//
// Solidity: event DepositNft(bytes32 accountNameHash, bytes32 nftContentHash, address tokenAddress, uint256 nftTokenId, uint16 creatorTreasuryRate)
func (_Zecreylegend *ZecreylegendFilterer) FilterDepositNft(opts *bind.FilterOpts) (*ZecreylegendDepositNftIterator, error) {

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "DepositNft")
	if err != nil {
		return nil, err
	}
	return &ZecreylegendDepositNftIterator{contract: _Zecreylegend.contract, event: "DepositNft", logs: logs, sub: sub}, nil
}

// WatchDepositNft is a free log subscription operation binding the contract event 0x42f9fbd77faf066fde386943aaafcc841fba6b755dc5da7939a046cbc493c24f.
//
// Solidity: event DepositNft(bytes32 accountNameHash, bytes32 nftContentHash, address tokenAddress, uint256 nftTokenId, uint16 creatorTreasuryRate)
func (_Zecreylegend *ZecreylegendFilterer) WatchDepositNft(opts *bind.WatchOpts, sink chan<- *ZecreylegendDepositNft) (event.Subscription, error) {

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "DepositNft")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendDepositNft)
				if err := _Zecreylegend.contract.UnpackLog(event, "DepositNft", log); err != nil {
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

// ParseDepositNft is a log parse operation binding the contract event 0x42f9fbd77faf066fde386943aaafcc841fba6b755dc5da7939a046cbc493c24f.
//
// Solidity: event DepositNft(bytes32 accountNameHash, bytes32 nftContentHash, address tokenAddress, uint256 nftTokenId, uint16 creatorTreasuryRate)
func (_Zecreylegend *ZecreylegendFilterer) ParseDepositNft(log types.Log) (*ZecreylegendDepositNft, error) {
	event := new(ZecreylegendDepositNft)
	if err := _Zecreylegend.contract.UnpackLog(event, "DepositNft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreylegendDesertModeIterator is returned from FilterDesertMode and is used to iterate over the raw logs and unpacked data for DesertMode events raised by the Zecreylegend contract.
type ZecreylegendDesertModeIterator struct {
	Event *ZecreylegendDesertMode // Event containing the contract specifics and raw log

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
func (it *ZecreylegendDesertModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendDesertMode)
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
		it.Event = new(ZecreylegendDesertMode)
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
func (it *ZecreylegendDesertModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendDesertModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendDesertMode represents a DesertMode event raised by the Zecreylegend contract.
type ZecreylegendDesertMode struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDesertMode is a free log retrieval operation binding the contract event 0x9f7e400a81dddbf1c18b1c37f82aa303d166295ca4b577eb2a7c23d4b704ba89.
//
// Solidity: event DesertMode()
func (_Zecreylegend *ZecreylegendFilterer) FilterDesertMode(opts *bind.FilterOpts) (*ZecreylegendDesertModeIterator, error) {

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "DesertMode")
	if err != nil {
		return nil, err
	}
	return &ZecreylegendDesertModeIterator{contract: _Zecreylegend.contract, event: "DesertMode", logs: logs, sub: sub}, nil
}

// WatchDesertMode is a free log subscription operation binding the contract event 0x9f7e400a81dddbf1c18b1c37f82aa303d166295ca4b577eb2a7c23d4b704ba89.
//
// Solidity: event DesertMode()
func (_Zecreylegend *ZecreylegendFilterer) WatchDesertMode(opts *bind.WatchOpts, sink chan<- *ZecreylegendDesertMode) (event.Subscription, error) {

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "DesertMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendDesertMode)
				if err := _Zecreylegend.contract.UnpackLog(event, "DesertMode", log); err != nil {
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

// ParseDesertMode is a log parse operation binding the contract event 0x9f7e400a81dddbf1c18b1c37f82aa303d166295ca4b577eb2a7c23d4b704ba89.
//
// Solidity: event DesertMode()
func (_Zecreylegend *ZecreylegendFilterer) ParseDesertMode(log types.Log) (*ZecreylegendDesertMode, error) {
	event := new(ZecreylegendDesertMode)
	if err := _Zecreylegend.contract.UnpackLog(event, "DesertMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreylegendFullExitCommitIterator is returned from FilterFullExitCommit and is used to iterate over the raw logs and unpacked data for FullExitCommit events raised by the Zecreylegend contract.
type ZecreylegendFullExitCommitIterator struct {
	Event *ZecreylegendFullExitCommit // Event containing the contract specifics and raw log

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
func (it *ZecreylegendFullExitCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendFullExitCommit)
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
		it.Event = new(ZecreylegendFullExitCommit)
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
func (it *ZecreylegendFullExitCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendFullExitCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendFullExitCommit represents a FullExitCommit event raised by the Zecreylegend contract.
type ZecreylegendFullExitCommit struct {
	ZkBNBBlockId uint32
	AccountId    uint32
	Owner        common.Address
	TokenId      uint16
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFullExitCommit is a free log retrieval operation binding the contract event 0x66fc63d751ecbefca61d4e2e7c534e4f29c61aed8ece23ed635277a7ea6f9bc4.
//
// Solidity: event FullExitCommit(uint32 indexed zkbnbBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_Zecreylegend *ZecreylegendFilterer) FilterFullExitCommit(opts *bind.FilterOpts, zkbnbBlockId []uint32, accountId []uint32, tokenId []uint16) (*ZecreylegendFullExitCommitIterator, error) {

	var zkbnbBlockIdRule []interface{}
	for _, zkbnbBlockIdItem := range zkbnbBlockId {
		zkbnbBlockIdRule = append(zkbnbBlockIdRule, zkbnbBlockIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "FullExitCommit", zkbnbBlockIdRule, accountIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZecreylegendFullExitCommitIterator{contract: _Zecreylegend.contract, event: "FullExitCommit", logs: logs, sub: sub}, nil
}

// WatchFullExitCommit is a free log subscription operation binding the contract event 0x66fc63d751ecbefca61d4e2e7c534e4f29c61aed8ece23ed635277a7ea6f9bc4.
//
// Solidity: event FullExitCommit(uint32 indexed zkbnbBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_Zecreylegend *ZecreylegendFilterer) WatchFullExitCommit(opts *bind.WatchOpts, sink chan<- *ZecreylegendFullExitCommit, zkbnbBlockId []uint32, accountId []uint32, tokenId []uint16) (event.Subscription, error) {

	var zkbnbBlockIdRule []interface{}
	for _, zkbnbBlockIdItem := range zkbnbBlockId {
		zkbnbBlockIdRule = append(zkbnbBlockIdRule, zkbnbBlockIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "FullExitCommit", zkbnbBlockIdRule, accountIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendFullExitCommit)
				if err := _Zecreylegend.contract.UnpackLog(event, "FullExitCommit", log); err != nil {
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

// ParseFullExitCommit is a log parse operation binding the contract event 0x66fc63d751ecbefca61d4e2e7c534e4f29c61aed8ece23ed635277a7ea6f9bc4.
//
// Solidity: event FullExitCommit(uint32 indexed zkbnbBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_Zecreylegend *ZecreylegendFilterer) ParseFullExitCommit(log types.Log) (*ZecreylegendFullExitCommit, error) {
	event := new(ZecreylegendFullExitCommit)
	if err := _Zecreylegend.contract.UnpackLog(event, "FullExitCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreylegendNewDefaultNFTFactoryIterator is returned from FilterNewDefaultNFTFactory and is used to iterate over the raw logs and unpacked data for NewDefaultNFTFactory events raised by the Zecreylegend contract.
type ZecreylegendNewDefaultNFTFactoryIterator struct {
	Event *ZecreylegendNewDefaultNFTFactory // Event containing the contract specifics and raw log

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
func (it *ZecreylegendNewDefaultNFTFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendNewDefaultNFTFactory)
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
		it.Event = new(ZecreylegendNewDefaultNFTFactory)
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
func (it *ZecreylegendNewDefaultNFTFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendNewDefaultNFTFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendNewDefaultNFTFactory represents a NewDefaultNFTFactory event raised by the Zecreylegend contract.
type ZecreylegendNewDefaultNFTFactory struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewDefaultNFTFactory is a free log retrieval operation binding the contract event 0xadf2ec6c092b7101c958de6b55caf4818f26a946e94fa556cbe3b3f29ff9f2ee.
//
// Solidity: event NewDefaultNFTFactory(address indexed factory)
func (_Zecreylegend *ZecreylegendFilterer) FilterNewDefaultNFTFactory(opts *bind.FilterOpts, factory []common.Address) (*ZecreylegendNewDefaultNFTFactoryIterator, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "NewDefaultNFTFactory", factoryRule)
	if err != nil {
		return nil, err
	}
	return &ZecreylegendNewDefaultNFTFactoryIterator{contract: _Zecreylegend.contract, event: "NewDefaultNFTFactory", logs: logs, sub: sub}, nil
}

// WatchNewDefaultNFTFactory is a free log subscription operation binding the contract event 0xadf2ec6c092b7101c958de6b55caf4818f26a946e94fa556cbe3b3f29ff9f2ee.
//
// Solidity: event NewDefaultNFTFactory(address indexed factory)
func (_Zecreylegend *ZecreylegendFilterer) WatchNewDefaultNFTFactory(opts *bind.WatchOpts, sink chan<- *ZecreylegendNewDefaultNFTFactory, factory []common.Address) (event.Subscription, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "NewDefaultNFTFactory", factoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendNewDefaultNFTFactory)
				if err := _Zecreylegend.contract.UnpackLog(event, "NewDefaultNFTFactory", log); err != nil {
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

// ParseNewDefaultNFTFactory is a log parse operation binding the contract event 0xadf2ec6c092b7101c958de6b55caf4818f26a946e94fa556cbe3b3f29ff9f2ee.
//
// Solidity: event NewDefaultNFTFactory(address indexed factory)
func (_Zecreylegend *ZecreylegendFilterer) ParseNewDefaultNFTFactory(log types.Log) (*ZecreylegendNewDefaultNFTFactory, error) {
	event := new(ZecreylegendNewDefaultNFTFactory)
	if err := _Zecreylegend.contract.UnpackLog(event, "NewDefaultNFTFactory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreylegendNewNFTFactoryIterator is returned from FilterNewNFTFactory and is used to iterate over the raw logs and unpacked data for NewNFTFactory events raised by the Zecreylegend contract.
type ZecreylegendNewNFTFactoryIterator struct {
	Event *ZecreylegendNewNFTFactory // Event containing the contract specifics and raw log

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
func (it *ZecreylegendNewNFTFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendNewNFTFactory)
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
		it.Event = new(ZecreylegendNewNFTFactory)
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
func (it *ZecreylegendNewNFTFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendNewNFTFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendNewNFTFactory represents a NewNFTFactory event raised by the Zecreylegend contract.
type ZecreylegendNewNFTFactory struct {
	CreatorAccountNameHash [32]byte
	CollectionId           uint32
	FactoryAddress         common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewNFTFactory is a free log retrieval operation binding the contract event 0xaa31b08b26b4240f7d2837e3fb6359824b509a4e4c5a5766af80f63c319add7e.
//
// Solidity: event NewNFTFactory(bytes32 indexed _creatorAccountNameHash, uint32 _collectionId, address _factoryAddress)
func (_Zecreylegend *ZecreylegendFilterer) FilterNewNFTFactory(opts *bind.FilterOpts, _creatorAccountNameHash [][32]byte) (*ZecreylegendNewNFTFactoryIterator, error) {

	var _creatorAccountNameHashRule []interface{}
	for _, _creatorAccountNameHashItem := range _creatorAccountNameHash {
		_creatorAccountNameHashRule = append(_creatorAccountNameHashRule, _creatorAccountNameHashItem)
	}

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "NewNFTFactory", _creatorAccountNameHashRule)
	if err != nil {
		return nil, err
	}
	return &ZecreylegendNewNFTFactoryIterator{contract: _Zecreylegend.contract, event: "NewNFTFactory", logs: logs, sub: sub}, nil
}

// WatchNewNFTFactory is a free log subscription operation binding the contract event 0xaa31b08b26b4240f7d2837e3fb6359824b509a4e4c5a5766af80f63c319add7e.
//
// Solidity: event NewNFTFactory(bytes32 indexed _creatorAccountNameHash, uint32 _collectionId, address _factoryAddress)
func (_Zecreylegend *ZecreylegendFilterer) WatchNewNFTFactory(opts *bind.WatchOpts, sink chan<- *ZecreylegendNewNFTFactory, _creatorAccountNameHash [][32]byte) (event.Subscription, error) {

	var _creatorAccountNameHashRule []interface{}
	for _, _creatorAccountNameHashItem := range _creatorAccountNameHash {
		_creatorAccountNameHashRule = append(_creatorAccountNameHashRule, _creatorAccountNameHashItem)
	}

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "NewNFTFactory", _creatorAccountNameHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendNewNFTFactory)
				if err := _Zecreylegend.contract.UnpackLog(event, "NewNFTFactory", log); err != nil {
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

// ParseNewNFTFactory is a log parse operation binding the contract event 0xaa31b08b26b4240f7d2837e3fb6359824b509a4e4c5a5766af80f63c319add7e.
//
// Solidity: event NewNFTFactory(bytes32 indexed _creatorAccountNameHash, uint32 _collectionId, address _factoryAddress)
func (_Zecreylegend *ZecreylegendFilterer) ParseNewNFTFactory(log types.Log) (*ZecreylegendNewNFTFactory, error) {
	event := new(ZecreylegendNewNFTFactory)
	if err := _Zecreylegend.contract.UnpackLog(event, "NewNFTFactory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreylegendNewPriorityRequestIterator is returned from FilterNewPriorityRequest and is used to iterate over the raw logs and unpacked data for NewPriorityRequest events raised by the Zecreylegend contract.
type ZecreylegendNewPriorityRequestIterator struct {
	Event *ZecreylegendNewPriorityRequest // Event containing the contract specifics and raw log

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
func (it *ZecreylegendNewPriorityRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendNewPriorityRequest)
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
		it.Event = new(ZecreylegendNewPriorityRequest)
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
func (it *ZecreylegendNewPriorityRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendNewPriorityRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendNewPriorityRequest represents a NewPriorityRequest event raised by the Zecreylegend contract.
type ZecreylegendNewPriorityRequest struct {
	Sender          common.Address
	SerialId        uint64
	TxType          uint8
	PubData         []byte
	ExpirationBlock *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPriorityRequest is a free log retrieval operation binding the contract event 0xd0943372c08b438a88d4b39d77216901079eda9ca59d45349841c099083b6830.
//
// Solidity: event NewPriorityRequest(address sender, uint64 serialId, uint8 txType, bytes pubData, uint256 expirationBlock)
func (_Zecreylegend *ZecreylegendFilterer) FilterNewPriorityRequest(opts *bind.FilterOpts) (*ZecreylegendNewPriorityRequestIterator, error) {

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return &ZecreylegendNewPriorityRequestIterator{contract: _Zecreylegend.contract, event: "NewPriorityRequest", logs: logs, sub: sub}, nil
}

// WatchNewPriorityRequest is a free log subscription operation binding the contract event 0xd0943372c08b438a88d4b39d77216901079eda9ca59d45349841c099083b6830.
//
// Solidity: event NewPriorityRequest(address sender, uint64 serialId, uint8 txType, bytes pubData, uint256 expirationBlock)
func (_Zecreylegend *ZecreylegendFilterer) WatchNewPriorityRequest(opts *bind.WatchOpts, sink chan<- *ZecreylegendNewPriorityRequest) (event.Subscription, error) {

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendNewPriorityRequest)
				if err := _Zecreylegend.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
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

// ParseNewPriorityRequest is a log parse operation binding the contract event 0xd0943372c08b438a88d4b39d77216901079eda9ca59d45349841c099083b6830.
//
// Solidity: event NewPriorityRequest(address sender, uint64 serialId, uint8 txType, bytes pubData, uint256 expirationBlock)
func (_Zecreylegend *ZecreylegendFilterer) ParseNewPriorityRequest(log types.Log) (*ZecreylegendNewPriorityRequest, error) {
	event := new(ZecreylegendNewPriorityRequest)
	if err := _Zecreylegend.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreylegendNoticePeriodChangeIterator is returned from FilterNoticePeriodChange and is used to iterate over the raw logs and unpacked data for NoticePeriodChange events raised by the Zecreylegend contract.
type ZecreylegendNoticePeriodChangeIterator struct {
	Event *ZecreylegendNoticePeriodChange // Event containing the contract specifics and raw log

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
func (it *ZecreylegendNoticePeriodChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendNoticePeriodChange)
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
		it.Event = new(ZecreylegendNoticePeriodChange)
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
func (it *ZecreylegendNoticePeriodChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendNoticePeriodChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendNoticePeriodChange represents a NoticePeriodChange event raised by the Zecreylegend contract.
type ZecreylegendNoticePeriodChange struct {
	NewNoticePeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNoticePeriodChange is a free log retrieval operation binding the contract event 0xf2b18f8abbd8a0d0c1fb8245146eedf5304887b12f6395b548ca238e054a1483.
//
// Solidity: event NoticePeriodChange(uint256 newNoticePeriod)
func (_Zecreylegend *ZecreylegendFilterer) FilterNoticePeriodChange(opts *bind.FilterOpts) (*ZecreylegendNoticePeriodChangeIterator, error) {

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "NoticePeriodChange")
	if err != nil {
		return nil, err
	}
	return &ZecreylegendNoticePeriodChangeIterator{contract: _Zecreylegend.contract, event: "NoticePeriodChange", logs: logs, sub: sub}, nil
}

// WatchNoticePeriodChange is a free log subscription operation binding the contract event 0xf2b18f8abbd8a0d0c1fb8245146eedf5304887b12f6395b548ca238e054a1483.
//
// Solidity: event NoticePeriodChange(uint256 newNoticePeriod)
func (_Zecreylegend *ZecreylegendFilterer) WatchNoticePeriodChange(opts *bind.WatchOpts, sink chan<- *ZecreylegendNoticePeriodChange) (event.Subscription, error) {

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "NoticePeriodChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendNoticePeriodChange)
				if err := _Zecreylegend.contract.UnpackLog(event, "NoticePeriodChange", log); err != nil {
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

// ParseNoticePeriodChange is a log parse operation binding the contract event 0xf2b18f8abbd8a0d0c1fb8245146eedf5304887b12f6395b548ca238e054a1483.
//
// Solidity: event NoticePeriodChange(uint256 newNoticePeriod)
func (_Zecreylegend *ZecreylegendFilterer) ParseNoticePeriodChange(log types.Log) (*ZecreylegendNoticePeriodChange, error) {
	event := new(ZecreylegendNoticePeriodChange)
	if err := _Zecreylegend.contract.UnpackLog(event, "NoticePeriodChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreylegendRegisterZNSIterator is returned from FilterRegisterZNS and is used to iterate over the raw logs and unpacked data for RegisterZNS events raised by the Zecreylegend contract.
type ZecreylegendRegisterZNSIterator struct {
	Event *ZecreylegendRegisterZNS // Event containing the contract specifics and raw log

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
func (it *ZecreylegendRegisterZNSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendRegisterZNS)
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
		it.Event = new(ZecreylegendRegisterZNS)
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
func (it *ZecreylegendRegisterZNSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendRegisterZNSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendRegisterZNS represents a RegisterZNS event raised by the Zecreylegend contract.
type ZecreylegendRegisterZNS struct {
	Name         string
	NameHash     [32]byte
	Owner        common.Address
	ZkBNBPubKeyX [32]byte
	ZkBNBPubKeyY [32]byte
	AccountIndex uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegisterZNS is a free log retrieval operation binding the contract event 0xebe5b00e9e9052d45c93d72544961460aa0fa30e0142c07d7dd81025cf9dab6c.
//
// Solidity: event RegisterZNS(string name, bytes32 nameHash, address owner, bytes32 zkbnbPubKeyX, bytes32 zkbnbPubKeyY, uint32 accountIndex)
func (_Zecreylegend *ZecreylegendFilterer) FilterRegisterZNS(opts *bind.FilterOpts) (*ZecreylegendRegisterZNSIterator, error) {

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "RegisterZNS")
	if err != nil {
		return nil, err
	}
	return &ZecreylegendRegisterZNSIterator{contract: _Zecreylegend.contract, event: "RegisterZNS", logs: logs, sub: sub}, nil
}

// WatchRegisterZNS is a free log subscription operation binding the contract event 0xebe5b00e9e9052d45c93d72544961460aa0fa30e0142c07d7dd81025cf9dab6c.
//
// Solidity: event RegisterZNS(string name, bytes32 nameHash, address owner, bytes32 zkbnbPubKeyX, bytes32 zkbnbPubKeyY, uint32 accountIndex)
func (_Zecreylegend *ZecreylegendFilterer) WatchRegisterZNS(opts *bind.WatchOpts, sink chan<- *ZecreylegendRegisterZNS) (event.Subscription, error) {

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "RegisterZNS")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendRegisterZNS)
				if err := _Zecreylegend.contract.UnpackLog(event, "RegisterZNS", log); err != nil {
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

// ParseRegisterZNS is a log parse operation binding the contract event 0xebe5b00e9e9052d45c93d72544961460aa0fa30e0142c07d7dd81025cf9dab6c.
//
// Solidity: event RegisterZNS(string name, bytes32 nameHash, address owner, bytes32 zkbnbPubKeyX, bytes32 zkbnbPubKeyY, uint32 accountIndex)
func (_Zecreylegend *ZecreylegendFilterer) ParseRegisterZNS(log types.Log) (*ZecreylegendRegisterZNS, error) {
	event := new(ZecreylegendRegisterZNS)
	if err := _Zecreylegend.contract.UnpackLog(event, "RegisterZNS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreylegendUpdateTokenPairIterator is returned from FilterUpdateTokenPair and is used to iterate over the raw logs and unpacked data for UpdateTokenPair events raised by the Zecreylegend contract.
type ZecreylegendUpdateTokenPairIterator struct {
	Event *ZecreylegendUpdateTokenPair // Event containing the contract specifics and raw log

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
func (it *ZecreylegendUpdateTokenPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendUpdateTokenPair)
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
		it.Event = new(ZecreylegendUpdateTokenPair)
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
func (it *ZecreylegendUpdateTokenPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendUpdateTokenPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendUpdateTokenPair represents a UpdateTokenPair event raised by the Zecreylegend contract.
type ZecreylegendUpdateTokenPair struct {
	PairIndex            uint16
	FeeRate              uint16
	TreasuryAccountIndex uint32
	TreasuryRate         uint16
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenPair is a free log retrieval operation binding the contract event 0x8727712867f019c156b0c25f308498ea1524b635902129b0acd3e9ff2c96f1c8.
//
// Solidity: event UpdateTokenPair(uint16 pairIndex, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_Zecreylegend *ZecreylegendFilterer) FilterUpdateTokenPair(opts *bind.FilterOpts) (*ZecreylegendUpdateTokenPairIterator, error) {

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "UpdateTokenPair")
	if err != nil {
		return nil, err
	}
	return &ZecreylegendUpdateTokenPairIterator{contract: _Zecreylegend.contract, event: "UpdateTokenPair", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenPair is a free log subscription operation binding the contract event 0x8727712867f019c156b0c25f308498ea1524b635902129b0acd3e9ff2c96f1c8.
//
// Solidity: event UpdateTokenPair(uint16 pairIndex, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_Zecreylegend *ZecreylegendFilterer) WatchUpdateTokenPair(opts *bind.WatchOpts, sink chan<- *ZecreylegendUpdateTokenPair) (event.Subscription, error) {

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "UpdateTokenPair")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendUpdateTokenPair)
				if err := _Zecreylegend.contract.UnpackLog(event, "UpdateTokenPair", log); err != nil {
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

// ParseUpdateTokenPair is a log parse operation binding the contract event 0x8727712867f019c156b0c25f308498ea1524b635902129b0acd3e9ff2c96f1c8.
//
// Solidity: event UpdateTokenPair(uint16 pairIndex, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_Zecreylegend *ZecreylegendFilterer) ParseUpdateTokenPair(log types.Log) (*ZecreylegendUpdateTokenPair, error) {
	event := new(ZecreylegendUpdateTokenPair)
	if err := _Zecreylegend.contract.UnpackLog(event, "UpdateTokenPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreylegendWithdrawNftIterator is returned from FilterWithdrawNft and is used to iterate over the raw logs and unpacked data for WithdrawNft events raised by the Zecreylegend contract.
type ZecreylegendWithdrawNftIterator struct {
	Event *ZecreylegendWithdrawNft // Event containing the contract specifics and raw log

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
func (it *ZecreylegendWithdrawNftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendWithdrawNft)
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
		it.Event = new(ZecreylegendWithdrawNft)
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
func (it *ZecreylegendWithdrawNftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendWithdrawNftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendWithdrawNft represents a WithdrawNft event raised by the Zecreylegend contract.
type ZecreylegendWithdrawNft struct {
	AccountIndex uint32
	NftL1Address common.Address
	ToAddress    common.Address
	NftL1TokenId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawNft is a free log retrieval operation binding the contract event 0x001c1fcac2c66b39f6b0a825948e9f8892b2d4a17edc60aaf61ca474e08402ec.
//
// Solidity: event WithdrawNft(uint32 accountIndex, address nftL1Address, address toAddress, uint256 nftL1TokenId)
func (_Zecreylegend *ZecreylegendFilterer) FilterWithdrawNft(opts *bind.FilterOpts) (*ZecreylegendWithdrawNftIterator, error) {

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "WithdrawNft")
	if err != nil {
		return nil, err
	}
	return &ZecreylegendWithdrawNftIterator{contract: _Zecreylegend.contract, event: "WithdrawNft", logs: logs, sub: sub}, nil
}

// WatchWithdrawNft is a free log subscription operation binding the contract event 0x001c1fcac2c66b39f6b0a825948e9f8892b2d4a17edc60aaf61ca474e08402ec.
//
// Solidity: event WithdrawNft(uint32 accountIndex, address nftL1Address, address toAddress, uint256 nftL1TokenId)
func (_Zecreylegend *ZecreylegendFilterer) WatchWithdrawNft(opts *bind.WatchOpts, sink chan<- *ZecreylegendWithdrawNft) (event.Subscription, error) {

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "WithdrawNft")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendWithdrawNft)
				if err := _Zecreylegend.contract.UnpackLog(event, "WithdrawNft", log); err != nil {
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

// ParseWithdrawNft is a log parse operation binding the contract event 0x001c1fcac2c66b39f6b0a825948e9f8892b2d4a17edc60aaf61ca474e08402ec.
//
// Solidity: event WithdrawNft(uint32 accountIndex, address nftL1Address, address toAddress, uint256 nftL1TokenId)
func (_Zecreylegend *ZecreylegendFilterer) ParseWithdrawNft(log types.Log) (*ZecreylegendWithdrawNft, error) {
	event := new(ZecreylegendWithdrawNft)
	if err := _Zecreylegend.contract.UnpackLog(event, "WithdrawNft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreylegendWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Zecreylegend contract.
type ZecreylegendWithdrawalIterator struct {
	Event *ZecreylegendWithdrawal // Event containing the contract specifics and raw log

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
func (it *ZecreylegendWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendWithdrawal)
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
		it.Event = new(ZecreylegendWithdrawal)
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
func (it *ZecreylegendWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendWithdrawal represents a Withdrawal event raised by the Zecreylegend contract.
type ZecreylegendWithdrawal struct {
	AssetId uint16
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xf4bf32c167ee6e782944cd1db8174729b46adcd3bc732e282cc4a80793933154.
//
// Solidity: event Withdrawal(uint16 assetId, uint128 amount)
func (_Zecreylegend *ZecreylegendFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*ZecreylegendWithdrawalIterator, error) {

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &ZecreylegendWithdrawalIterator{contract: _Zecreylegend.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xf4bf32c167ee6e782944cd1db8174729b46adcd3bc732e282cc4a80793933154.
//
// Solidity: event Withdrawal(uint16 assetId, uint128 amount)
func (_Zecreylegend *ZecreylegendFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ZecreylegendWithdrawal) (event.Subscription, error) {

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendWithdrawal)
				if err := _Zecreylegend.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xf4bf32c167ee6e782944cd1db8174729b46adcd3bc732e282cc4a80793933154.
//
// Solidity: event Withdrawal(uint16 assetId, uint128 amount)
func (_Zecreylegend *ZecreylegendFilterer) ParseWithdrawal(log types.Log) (*ZecreylegendWithdrawal, error) {
	event := new(ZecreylegendWithdrawal)
	if err := _Zecreylegend.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreylegendWithdrawalNFTPendingIterator is returned from FilterWithdrawalNFTPending and is used to iterate over the raw logs and unpacked data for WithdrawalNFTPending events raised by the Zecreylegend contract.
type ZecreylegendWithdrawalNFTPendingIterator struct {
	Event *ZecreylegendWithdrawalNFTPending // Event containing the contract specifics and raw log

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
func (it *ZecreylegendWithdrawalNFTPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreylegendWithdrawalNFTPending)
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
		it.Event = new(ZecreylegendWithdrawalNFTPending)
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
func (it *ZecreylegendWithdrawalNFTPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreylegendWithdrawalNFTPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreylegendWithdrawalNFTPending represents a WithdrawalNFTPending event raised by the Zecreylegend contract.
type ZecreylegendWithdrawalNFTPending struct {
	NftIndex *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalNFTPending is a free log retrieval operation binding the contract event 0x71aea045e572384e5b53812f175c12adb074001c9c13d379730d15188d3729bc.
//
// Solidity: event WithdrawalNFTPending(uint40 indexed nftIndex)
func (_Zecreylegend *ZecreylegendFilterer) FilterWithdrawalNFTPending(opts *bind.FilterOpts, nftIndex []*big.Int) (*ZecreylegendWithdrawalNFTPendingIterator, error) {

	var nftIndexRule []interface{}
	for _, nftIndexItem := range nftIndex {
		nftIndexRule = append(nftIndexRule, nftIndexItem)
	}

	logs, sub, err := _Zecreylegend.contract.FilterLogs(opts, "WithdrawalNFTPending", nftIndexRule)
	if err != nil {
		return nil, err
	}
	return &ZecreylegendWithdrawalNFTPendingIterator{contract: _Zecreylegend.contract, event: "WithdrawalNFTPending", logs: logs, sub: sub}, nil
}

// WatchWithdrawalNFTPending is a free log subscription operation binding the contract event 0x71aea045e572384e5b53812f175c12adb074001c9c13d379730d15188d3729bc.
//
// Solidity: event WithdrawalNFTPending(uint40 indexed nftIndex)
func (_Zecreylegend *ZecreylegendFilterer) WatchWithdrawalNFTPending(opts *bind.WatchOpts, sink chan<- *ZecreylegendWithdrawalNFTPending, nftIndex []*big.Int) (event.Subscription, error) {

	var nftIndexRule []interface{}
	for _, nftIndexItem := range nftIndex {
		nftIndexRule = append(nftIndexRule, nftIndexItem)
	}

	logs, sub, err := _Zecreylegend.contract.WatchLogs(opts, "WithdrawalNFTPending", nftIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreylegendWithdrawalNFTPending)
				if err := _Zecreylegend.contract.UnpackLog(event, "WithdrawalNFTPending", log); err != nil {
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

// ParseWithdrawalNFTPending is a log parse operation binding the contract event 0x71aea045e572384e5b53812f175c12adb074001c9c13d379730d15188d3729bc.
//
// Solidity: event WithdrawalNFTPending(uint40 indexed nftIndex)
func (_Zecreylegend *ZecreylegendFilterer) ParseWithdrawalNFTPending(log types.Log) (*ZecreylegendWithdrawalNFTPending, error) {
	event := new(ZecreylegendWithdrawalNFTPending)
	if err := _Zecreylegend.contract.UnpackLog(event, "WithdrawalNFTPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
