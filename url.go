package bchapi

const DNS = "https://bch-chain.api.btc.com/"

const BlockTxsUrl = DNS + "v3/block/%d/tx"
const BlockUrl = DNS + "v3/block/%d"
const AddressUrl = DNS + "v3/address/%s"
const AddressTxUrl = DNS + "v3/address/%s/tx?page=%d&pagesize=%d"
const AddressUnspentUrl = DNS + "v3/address/%s/unspent?page=%d&pagesize=%d"
