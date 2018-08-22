package trustsql

const (
	PUBKEY_DIGEST_LENGTH     = 90   // public key length
	PRVKEY_DIGEST_LENGTH     = 45   // private key length
	ADDR_DIGEST_LENGTH       = 35   // address length
	SIGN_DIGEST_LENGTH       = 98   // signature length
	KEY_DES3_DIGEST_LENGTH   = 24   // max size of key for DES3 encrypt
	KEY_AES128_DIGEST_LENGTH = 16   // max size of key for AES128 encrypt
	TRANSSQL_DIGEST_LENGTH   = 8192 // max size of trans sql for TrustSQL

	RANDOM_NUMBER_ALGORITHM          = "SHA1PRNG"
	RANDOM_NUMBER_ALGORITHM_PROVIDER = "SUN"
	MAXPRIVATEKEY                    = "00FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364140"
	INFO_SHARE_PUBKEY                = "BC8s/4qEAvVl4Sv0LwQOWJcVU6Q5hBd+7LlJeEivVmUbdtwP4RTfN8x/G+muMhN8SrweyyVVMIcIrnMWoFqGfIA="
)
