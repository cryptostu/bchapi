package bchapi

const DNS = "https://bch-chain.api.btc.com/"

const BlockTxsUrl = DNS + "v3/block/%d/tx"
const BlockUrl = DNS + "v3/block/%d"
const AddressUrl = DNS + "v3/address/%s"
const AddressTxUrl = DNS + "v3/address/%s/tx?page=%d&pagesize=%d"
const AddressUnspentUrl = DNS + "v3/address/%s/unspent?page=%d&pagesize=%d"
const ToolsTxDecodeUrl = DNS + "v3/tools/tx-decode"
const ToolsTxPublishUrl = DNS + "v3/tools/tx-publish"
const ToolsVerifyMessageUrl = DNS + "v3/tools/verify-message"
const TxUrl = DNS + "v3/tx/%s"
const TxUnconfirmedUrl = DNS + "v3/tx/unconfirmed"
const TxUnconfirmedSummaryUrl = DNS + "v3/tx/unconfirmed/summary"
