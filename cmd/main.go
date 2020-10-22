package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/studyzy/fabric-sdk-go/pkg/client/ledger"
	"github.com/studyzy/fabric-sdk-go/pkg/fabsdk"
)
func main() {
	queryChainInfo()
}
var (
	fabricSdk        *fabsdk.FabricSDK
userName string
	orgName string
)

func queryChainInfo() error {

	channelContext := fabricSdk.ChannelContext("mychannel", fabsdk.WithUser(userName), fabsdk.WithOrg(orgName))
	client, err := ledger.New(channelContext)
	if err != nil {
		return err
	}
	result, err := client.QueryInfo()
	if err != nil {
		log.Fatal(err.Error())
		return err
	} else {
		height := result.BCI.Height
		blockHash := hex.EncodeToString(result.BCI.CurrentBlockHash)
		if err != nil {
			return err
		}
		fmt.Println("block_number", height)
		fmt.Println("current_block_hash", blockHash)
	}
	return nil
}
