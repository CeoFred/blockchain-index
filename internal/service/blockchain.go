package service

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

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
		// ContractABI:     parsedABI,
		// ContractAddress: common.HexToAddress(contractAddress),
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
