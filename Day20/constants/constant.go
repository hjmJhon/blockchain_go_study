package constants

import "path/filepath"

//wallet
const VERSION = byte(0x00)
const ADDRESS_CHECKSUM_LEN = 4

const WALLETDIR = "." + string(filepath.Separator) + "wallets" + string(filepath.Separator)
const WALLET_SUFFIX = ".dat"
