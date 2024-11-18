package main

import (
	"fmt"
	"regexp"
)


func ExtractMoa1qqqAddresses(input string) []string {
	
	re := regexp.MustCompile(`(moa1qqq[0-9a-z]+)`)

	
	matches := re.FindAllString(input, -1)
	return matches
}

func main() {

	input := `
/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/claimRewards/tx-claim-rewards.json
  7,16:   "receiver": "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllsxxctf0",

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/delegators/delegator-after-delegate.json
  3,16:   "contract": "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhllllss2qdpm",

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/delegators/delegator-after-revert.json
  3,16:   "contract": "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhllllss2qdpm",

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/delegators/delegator-after-un-delegate-1.json
  3,16:   "contract": "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhllllss2qdpm",

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/delegators/delegator-after-un-delegate-2.json
  3,16:   "contract": "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhllllss2qdpm",

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/delegators/delegator-after-withdraw.json
  3,16:   "contract": "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhllllss2qdpm",

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/multiTransferWithScCallAndErrorSignaledBySC/transaction-after-execution-of-scr-dst-shard.json
  23,6:     "moa1qqqqqqqqqqqqqpgqt6ltx52ukss9d2qag2k67at28a36xc9lkp2swhu7e9",
  24,6:     "moa1qqqqqqqqqqqqqpgqt6ltx52ukss9d2qag2k67at28a36xc9lkp2swhu7e9"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/multiTransferWithScCallAndErrorSignaledBySC/transaction-executed-on-source.json
  23,6:     "moa1qqqqqqqqqqqqqpgqt6ltx52ukss9d2qag2k67at28a36xc9lkp2swhu7e9",
  24,6:     "moa1qqqqqqqqqqqqqpgqt6ltx52ukss9d2qag2k67at28a36xc9lkp2swhu7e9"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/nftTransferCrossShard/op-nft-transfer-sc-call-after-refund.json
  33,6:     "moa1qqqqqqqqqqqqqpgq78y09lw93f3udvsplshdv2vk957l5vl70n4sv89j36"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/nftTransferCrossShard/tx-complete-with-status.json
  24,6:     "moa1qqqqqqqqqqqqqpgq57szwud2quysucrlq2e97ntdysdl7v4ejz3q7f4au9"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/nftTransferCrossShard/tx-nft-transfer-failed-on-dst.json
  33,6:     "moa1qqqqqqqqqqqqqpgq57szwud2quysucrlq2e97ntdysdl7v4ejz3q7f4au9"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/nftTransferCrossShard/tx-nft-transfer-sc-call-after-refund.json
  33,6:     "moa1qqqqqqqqqqqqqpgq78y09lw93f3udvsplshdv2vk957l5vl70n4sv89j36"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/nftTransferCrossShard/tx-nft-transfer-sc-call-source.json
  24,6:     "moa1qqqqqqqqqqqqqpgq78y09lw93f3udvsplshdv2vk957l5vl70n4sv89j36"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/nftTransferCrossShard/tx-nft-transfer-success-source-second.json
  20,6:     "moa1qqqqqqqqqqqqqpgq57szwud2quysucrlq2e97ntdysdl7v4ejz3q7f4au9"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/nftTransferCrossShard/tx-nft-transfer.json
  25,6:     "moa1qqqqqqqqqqqqqpgq57szwud2quysucrlq2e97ntdysdl7v4ejz3q7f4au9"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/nftTransferCrossShardWithScCall/cross-shard-transfer-with-sc-call.json
  33,6:     "moa1qqqqqqqqqqqqqpgq57szwud2quysucrlq2e97ntdysdl7v4ejz3q7f4au9"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/relayedTx/relayed-tx-after-refund.json
  24,6:     "moa1qqqqqqqqqqqqqpgq3dswlnnlkfd3gqrcv3dhzgnvh8ryf27g5rfs5q4ukq"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/relayedTx/relayed-tx-source.json
  24,6:     "moa1qqqqqqqqqqqqqpgq3dswlnnlkfd3gqrcv3dhzgnvh8ryf27g5rfs5q4ukq"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/scCallIntraShard/claim-rewards.json
  7,16:   "receiver": "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqfhlllls4mdmxn",

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/scCallIntraShard/sc-call-fail.json
  7,16:   "receiver": "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqfhlllls4mdmxn",

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/scCallWithIssueDcdt/scr-with-callback-executed-on-destination-shard.json
  8,14:   "sender": "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls29jpxv",
  9,16:   "receiver": "moa1qqqqqqqqqqqqqpgqahumqen35dr9k4rmcnd70mqt5t4mt7ey4nwqr2akec",

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/scCallWithIssueDcdt/scr-with-callback-executed-on-source.json
  8,14:   "sender": "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls29jpxv",
  9,16:   "receiver": "moa1qqqqqqqqqqqqqpgqahumqen35dr9k4rmcnd70mqt5t4mt7ey4nwqr2akec",

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/scCallWithIssueDcdt/scr-with-issue-executed-on-destination-shard.json
  8,14:   "sender": "moa1qqqqqqqqqqqqqpgqahumqen35dr9k4rmcnd70mqt5t4mt7ey4nwqr2akec",
  9,16:   "receiver": "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls29jpxv",

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/scCallWithIssueDcdt/scr-with-issue-executed-on-source-shard.json
  8,14:   "sender": "moa1qqqqqqqqqqqqqpgqahumqen35dr9k4rmcnd70mqt5t4mt7ey4nwqr2akec",
  9,16:   "receiver": "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls29jpxv",

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/scCallWithIssueDcdt/tx-after-execution-of-callback-on-destination-shard.json
  7,16:   "receiver": "moa1qqqqqqqqqqqqqpgqahumqen35dr9k4rmcnd70mqt5t4mt7ey4nwqr2akec",

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/scCallWithIssueDcdt/tx-after-execution-on-source-shard.json
  7,16:   "receiver": "moa1qqqqqqqqqqqqqpgqahumqen35dr9k4rmcnd70mqt5t4mt7ey4nwqr2akec",

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/scCallWithIssueDcdt/tx-in-op-index-execution-of-callback-on-destination-shard.json
  7,16:   "receiver": "moa1qqqqqqqqqqqqqpgqahumqen35dr9k4rmcnd70mqt5t4mt7ey4nwqr2akec",

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/testdata/scDeploy/tx-sc-deploy.json
  7,16:   "receiver": "moa1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhsx6tv",

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/claimRewards_test.go
  55,22: 	addressReceiver := "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllsxxctf0"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/delegators_test.go
  49,30: 					Address: decodeAddress("moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhllllss2qdpm"),
  80,30: 					Address: decodeAddress("moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhllllss2qdpm"),
  110,30: 					Address: decodeAddress("moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhllllss2qdpm"),
  150,30: 					Address: decodeAddress("moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhllllss2qdpm"),

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/multiTransferWithScCallAndErrorSignaledBySC_test.go
  53,15: 	address2 := "moa1qqqqqqqqqqqqqpgqa0fsfshnff4n76jhcye6k7uvd7qacsq42jpsvzkct6"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/nftTransferCrossShard_test.go
  55,15: 	address2 := "moa1qqqqqqqqqqqqqpgq78y09lw93f3udvsplshdv2vk957l5vl70n4sv89j36"
  190,15: 	address2 := "moa1qqqqqqqqqqqqqpgq57szwud2quysucrlq2e97ntdysdl7v4ejz3q7f4au9"
  322,15: 	address2 := "moa1qqqqqqqqqqqqqpgq57szwud2quysucrlq2e97ntdysdl7v4ejz3q7f4au9"
  473,15: 	address2 := "moa1qqqqqqqqqqqqqpgq57szwud2quysucrlq2e97ntdysdl7v4ejz3q7f4au9"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/nftTransferCrossWithScCall_test.go
  52,15: 	address2 := "moa1qqqqqqqqqqqqqpgq57szwud2quysucrlq2e97ntdysdl7v4ejz3q7f4au9"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/relayedTx_test.go
  54,15: 	address3 := "moa1qqqqqqqqqqqqqpgq3dswlnnlkfd3gqrcv3dhzgnvh8ryf27g5rfs5q4ukq"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/scCallIntraShard_test.go
  47,15: 	address2 := "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqfhlllls4mdmxn"
  142,15: 	address2 := "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqfhlllls4mdmxn"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/scCallWithIssueDct_test.go
  54,22: 	contractAddress := "moa1qqqqqqqqqqqqqpgqahumqen35dr9k4rmcnd70mqt5t4mt7ey4nwqr2akec"
  65,19: 	dcdtSystemSC := "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls29jpxv"

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/integrationtests/scDeploy_test.go
  48,28: 		RcvAddr:  decodeAddress("moa1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhsx6tv"),
  78,44: 							Topics:     [][]byte{decodeAddress("moa1qqqqqqqqqqqqqpgq4t2tqxpst9a6qttpak8cz8wvz6a0nses5qfq58uvt5"), decodeAddress("moa12m3x8jp6dl027pj5f2nw6ght2cyhhjfrs86cdwsa8xn83r375qfqwk8z6l"), []byte("codeHash")},
  99,18: 	ids = []string{"moa1qqqqqqqqqqqqqpgq4t2tqxpst9a6qttpak8cz8wvz6a0nses5qfq58uvt5"}
  120,44: 							Topics:     [][]byte{decodeAddress("moa1qqqqqqqqqqqqqpgq4t2tqxpst9a6qttpak8cz8wvz6a0nses5qfq58uvt5"), decodeAddress("moa12m3x8jp6dl027pj5f2nw6ght2cyhhjfrs86cdwsa8xn83r375qfqwk8z6l"), []byte("secondCodeHash")},
  131,18: 	ids = []string{"moa1qqqqqqqqqqqqqpgq4t2tqxpst9a6qttpak8cz8wvz6a0nses5qfq58uvt5"}
  150,35: 							Address:    decodeAddress("moa1qqqqqqqqqqqqqpgq4t2tqxpst9a6qttpak8cz8wvz6a0nses5qfq58uvt5"),
  163,18: 	ids = []string{"moa1qqqqqqqqqqqqqpgq4t2tqxpst9a6qttpak8cz8wvz6a0nses5qfq58uvt5"}
  182,35: 							Address:    decodeAddress("moa1qqqqqqqqqqqqqpgq4t2tqxpst9a6qttpak8cz8wvz6a0nses5qfq58uvt5"),
  195,18: 	ids = []string{"moa1qqqqqqqqqqqqqpgq4t2tqxpst9a6qttpak8cz8wvz6a0nses5qfq58uvt5"}

/home/magnus/Desktop/addressesTesting/drt-chain-es-indexer-go/process/elasticproc/transactions/transactionsProcessor_test.go
  609,28: 				RcvAddr: decodeBech32("moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls29jpxv"),
  617,35: 				SndAddr:        decodeBech32("moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls29jpxv"),
  623,35: 				SndAddr:        decodeBech32("moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls29jpxv"),
  638,28: 				RcvAddr: decodeBech32("moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls29jpxv"),
  646,35: 				SndAddr:        decodeBech32("moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls29jpxv"),

/home/magnus/Desktop/addressesTesting/drt-chain-go/cmd/node/config/config.toml
  702,10:         "moa1qqqqqqqqqqqqqpgqr46jrxr6r2unaqh75ugd308dwx5vgnhwh47qx5lw9p",

/home/magnus/Desktop/addressesTesting/drt-chain-go/cmd/node/config/genesis.json
  9,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  20,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  31,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  42,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  53,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  64,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  75,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  86,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  97,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  108,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",

/home/magnus/Desktop/addressesTesting/drt-chain-go/cmd/node/config/nodesSetup.json
  14,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  20,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  26,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  32,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",

/home/magnus/Desktop/addressesTesting/drt-chain-go/examples/address_test.go
  54,61: 	require.Equal(t, core.MetachainShardId, computeShardID(t, "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls29jpxv", shardCoordinator))
  114,19: 	assert.Equal(t, "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqlllsz87dvw", stakingScAddress)
  115,19: 	assert.Equal(t, "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllsxxctf0", validatorScAddress)
  116,19: 	assert.Equal(t, "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls29jpxv", dcdtScAddress)
  117,19: 	assert.Equal(t, "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqrlllswy58rd", governanceScAddress)
  118,19: 	assert.Equal(t, "moa1qqqqqqqqqqqqqqqpqqqqqqqqqrllllllllllllllllllllllllls7zfxnx", jailingAddress)
  119,19: 	assert.Equal(t, "moa1qqqqqqqqqqqqqqqpqqqqqqqqlllllllllllllllllllllllllllswawjsh", endOfEpochAddress)
  120,19: 	assert.Equal(t, "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqylllsjrx4c2", delegationManagerScAddress)
  121,19: 	assert.Equal(t, "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq0llllsdwmvu2", firstDelegationScAddress)
  122,19: 	assert.Equal(t, "moa1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhsx6tv", contractDeployScAdress)

/home/magnus/Desktop/addressesTesting/drt-chain-go/integrationTests/chainSimulator/staking/stake/stakeAndUnStake_test.go
  94,10: 	rcv := "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllsxxctf0"
  228,10: 	rcv := "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllsxxctf0"
  2559,10: 	rcv := "moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllsxxctf0"

/home/magnus/Desktop/addressesTesting/drt-chain-go/integrationTests/factory/testdata/genesis.json
  8,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  18,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  28,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  38,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  48,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  58,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  68,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  78,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",
  88,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha",

/home/magnus/Desktop/addressesTesting/drt-chain-go/integrationTests/factory/testdata/nodesSetup.json
  13,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha"
  17,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha"
  21,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha"
  25,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha"
  29,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha"
  33,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha"
  37,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha"
  41,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha"
  45,19:       "address": "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha"

/home/magnus/Desktop/addressesTesting/drt-chain-go/integrationTests/vm/testInitializer.go
  86,23: const DNSV2Address = "moa1qqqqqqqqqqqqqpgqcy67yanvwpepqmerkq6m8pgav0tlvgwxjmdqc06e67"

/home/magnus/Desktop/addressesTesting/drt-chain-go/node/chainSimulator/components/testOnlyProcessingNode_test.go
  262,16: 	scAddress := "moa1qqqqqqqqqqqqqpgqrchxzx5uu8sv3ceg8nx8cxc0gesezure5awq7du8ha"

/home/magnus/Desktop/addressesTesting/drt-chain-go/node/chainSimulator/chainSimulator_test.go
  226,22: 	contractAddress := "moa1qqqqqqqqqqqqqpgqmzzm05jeav6d5qvna0q2pmcllelkz8xddz3sf2kd6y"
  307,22: 	contractAddress := "moa1qqqqqqqqqqqqqpgqmzzm05jeav6d5qvna0q2pmcllelkz8xddz3sf2kd6y"

/home/magnus/Desktop/addressesTesting/drt-chain-go/node/external/logs/logsConverter_test.go
  15,28: 	contractAddressBech32 := "moa1qqqqqqqqqqqqqpgqxwakt2g7u9atsnr03gqcgmhcv38pt7mkd94qhg3njm"

/home/magnus/Desktop/addressesTesting/drt-chain-go/node/external/transactionAPI/apiTransactionProcessor_test.go
  1100,20: 		Sender:         "moa1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhsx6tv",

/home/magnus/Desktop/addressesTesting/drt-chain-go/node/external/transactionAPI/gasUsedAndFeeProcessor_test.go
  214,24: 		Receivers: []string{"moa1qqqqqqqqqqqqqpgq6wegs2xkypfpync8mn2sa5cmpqjlvrhwz5nq9p8t5h"},

/home/magnus/Desktop/addressesTesting/drt-chain-go/node/node_test.go
  4363,44: 		isMigrated, err := n.IsDataTrieMigrated("moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllsxxctf0", api.AccountQueryOptions{})
  4383,44: 		isMigrated, err := n.IsDataTrieMigrated("moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllsxxctf0", api.AccountQueryOptions{})
  4410,44: 		isMigrated, err := n.IsDataTrieMigrated("moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllsxxctf0", api.AccountQueryOptions{})
  4437,44: 		isMigrated, err := n.IsDataTrieMigrated("moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllsxxctf0", api.AccountQueryOptions{})

/home/magnus/Desktop/addressesTesting/drt-chain-go/outport/process/transactionsfee/transactionsFeeProcessor_test.go
  88,27: 		RcvAddr:        []byte("moa1qqqqqqqqqqqqqpgq3dswlnnlkfd3gqrcv3dhzgnvh8ryf27g5rfs5q4ukq"),
  108,30: 					SndAddr:        []byte("moa1qqqqqqqqqqqqqpgq3dswlnnlkfd3gqrcv3dhzgnvh8ryf27g5rfs5q4ukq"),
  146,28: 			SndAddr:        []byte("moa1qqqqqqqqqqqqqpgq3dswlnnlkfd3gqrcv3dhzgnvh8ryf27g5rfs5q4ukq"),
  398,28: 			SndAddr:        []byte("moa1qqqqqqqqqqqqqpgq3dswlnnlkfd3gqrcv3dhzgnvh8ryf27g5rfs5q4ukq"),
  446,28: 			SndAddr:        []byte("moa1qqqqqqqqqqqqqpgq3dswlnnlkfd3gqrcv3dhzgnvh8ryf27g5rfs5q4ukq"),
  501,28: 			SndAddr:        []byte("moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls29jpxv"),
  538,21: 		RcvAddr:  []byte("moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls29jpxv"),
  546,28: 			SndAddr:        []byte("moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls29jpxv"),

/home/magnus/Desktop/addressesTesting/drt-chain-go/testscommon/integrationtests/stringers_test.go
  67,12: 		RcvAddr: moa1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhsx6tv

/home/magnus/Desktop/addressesTesting/drt-chain-scenario-go/scenario/expression/integrationTests/exprInterpreterReconstructor_test.go
  192,44: 	result, err := ei.InterpretString("bech32:moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllsxxctf0")
  195,27: 	require.Equal(t, "bech32:moa1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllsxxctf0", er.Reconstruct(result, mer.AddressHint))

/home/magnus/Desktop/addressesTesting/drt-chain-vm-common-go/parsers/dataField/parseSingleNFTTransfer_test.go
  17,40: var receiverSC, _ = pubKeyConv.Decode("moa1qqqqqqqqqqqqqpgqp699jngundfqw07d8jzkepucvpzush6k3wvqfqn6lk")
	`

	addresses := ExtractMoa1qqqAddresses(input)
	for _, address := range addresses {
		fmt.Println(address)
	}
}
