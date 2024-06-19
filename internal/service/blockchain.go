package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"

	binding "github.com/CeoFred/gin-boilerplate/internal/smartcontract/binding"
	"github.com/gofrs/uuid"

	"github.com/CeoFred/gin-boilerplate/internal/models"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type LogTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

type LogApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
}

type BlockchainService struct {
	Client          *ethclient.Client
	ContractABI     abi.ABI
	ContractAddress common.Address
}

func NewBlockchainService(rpcURL string) (*BlockchainService, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	return &BlockchainService{
		Client: client,
	}, nil
}

func (b *BlockchainService) Close() {
	b.Client.Close()
}

func (b *BlockchainService) SetContract(address, contractABI string) error {
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return err
	}

	b.ContractABI = parsedABI
	b.ContractAddress = common.HexToAddress(address)

	return nil
}

func (b *BlockchainService) IsContract(contractAddress string) bool {
	address := common.HexToAddress(contractAddress)
	bytecode, err := b.Client.CodeAt(context.Background(), address, nil)
	if err != nil {
		return false
	}
	isContract := len(bytecode) > 0
	return isContract
}

func (b *BlockchainService) Subscribe(address string, logs chan types.Log) error {
	contractAddress := common.HexToAddress(address)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	sub, err := b.Client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	err = <-sub.Err()
	if err != nil {
		log.Printf("Error in subscription: %v", err)
	}

	log.Println("Subscription successful")

	return nil
}

func (b *BlockchainService) GetLatestBlockNumber() (*big.Int, error) {
	header, err := b.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}

func (b *BlockchainService) QueryLogs(contractAddress string, startBlock, endBlock uint) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
		FromBlock: big.NewInt(int64(startBlock)),
		ToBlock:   big.NewInt(int64(endBlock)),
	}

	logs, err := b.Client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to filter logs: %w", err)
	}

	return logs, nil
}

func (b *BlockchainService) ProcessERCTokenLogs(logs []types.Log, events []*models.ContractEvent) []*models.EventLog {

	eventLogs := []*models.EventLog{}

	eventStructMap := map[string]interface{}{
		"Transfer": new(LogTransfer),
		"Approval": new(LogApproval),
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(binding.ERC20ABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, elog := range logs {
		event_signature := elog.Topics[0].Hex()
		fmt.Println(event_signature)
		for _, contract_event := range events {

			if contract_event.Event.Signature == event_signature {

				t, ok := eventStructMap[contract_event.Event.Name]
				if !ok {
					log.Printf("Unknown event name: %s", contract_event.Event.Name)
					continue
				}

				err := contractAbi.UnpackIntoInterface(t, contract_event.Event.Name, elog.Data)
				if err != nil {
					log.Fatal(err)
				}

				var eventData interface{}

				switch eventType := t.(type) {
				case *LogTransfer:
					eventType.From = common.HexToAddress(elog.Topics[1].Hex())
					eventType.To = common.HexToAddress(elog.Topics[2].Hex())

					eventData = eventType
					fmt.Printf("Processed Transfer event: From=%s, To=%s, Value=%s\n\n", eventType.From.Hex(), eventType.To.Hex(), eventType.Value.String())
				case *LogApproval:
					eventType.Owner = common.HexToAddress(elog.Topics[1].Hex())
					eventType.Spender = common.HexToAddress(elog.Topics[2].Hex())
					eventData = eventType
					fmt.Printf("Processed Approval event: Owner=%s, Spender=%s, Value=%s\n\n", eventType.Owner.Hex(), eventType.Spender.Hex(), eventType.Value.String())
				default:
					log.Printf("Unhandled event type: %T", eventType)
				}

				tx, isPending, err := b.Client.TransactionByHash(context.Background(), elog.TxHash)
				if err != nil {
					log.Printf("Failed to fetch transaction by hash: %v", err)
					continue
				}

				if !isPending { //TODO: handle pending transactions
					from, err := b.Client.TransactionSender(context.Background(), tx, elog.BlockHash, elog.TxIndex)
					if err != nil {
						log.Printf("Failed to fetch transaction sender: %v", err)
						continue
					}
					fmt.Printf("Transaction initiated by: %s\n", from.Hex())
					//TODO: save user event/transaction
				}

				// PUSH TO EVENT LOGS
				ID, err := uuid.NewV7()
				if err != nil {
					log.Printf("Failed to generate event log uuid: %v", err)
					continue
				}
				blockNumber := elog.BlockNumber
				transactionHash := elog.TxHash.Hex()
				logIndex := elog.Index

				data := map[string]interface{}{
					"raw":       fmt.Sprintf("%x", elog.Data),
					"formatted": eventData,
				}

				dataJSON, err := json.Marshal(data)
				if err != nil {
					log.Printf("Failed to marshal event data: %v", err)
					continue
				}

				topicsJSON, err := json.Marshal(elog.Topics)
				if err != nil {
					log.Printf("Failed to marshal topics: %v", err)
					continue
				}

				log := &models.EventLog{
					ID:              ID,
					ContractAddress: contract_event.Contract.Address,
					ContractEventID: contract_event.ID,
					EventName:       contract_event.Event.Name,
					ContractID:      contract_event.Contract.ID,
					BlockNumber:     blockNumber,
					TransactionHash: transactionHash,
					LogIndex:        logIndex,
					Data:            dataJSON,
					Topics:          topicsJSON,
				}
				eventLogs = append(eventLogs, log)

			}

		}

	}
	return eventLogs
}

func (b *BlockchainService) GetTransactionByHash(hash string) {

}
