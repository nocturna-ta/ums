package utils

import (
	"crypto/sha512"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/nocturna-ta/golib/router"
	"github.com/nocturna-ta/golib/utils/encryption"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/pkg/constants"
	"log"
	"regexp"
	"strings"
)

func CustomHash(str ...string) string {
	hash := sha512.New()
	hash.Write([]byte(strings.Join(str, "#")))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func PasswordHash(password, salt string) string {
	return CustomHash(password, salt)
}

func Encryption(text string) string {
	cfg := &config.MainConfig{}
	config.ReadConfig(cfg, "")

	enc, err := encryption.NewEncryption(cfg.Encryption.Key)
	if err != nil {
		log.Fatal(err)
	}

	res, err := enc.Encrypt(text)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func Decryption(text string) string {
	cfg := &config.MainConfig{}
	config.ReadConfig(cfg, "")

	enc, err := encryption.NewEncryption(cfg.Encryption.Key)
	if err != nil {
		log.Fatal(err)
	}

	res, err := enc.Decrypt(text)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func IsValidNIK(nik string) bool {
	nik = strings.TrimSpace(nik)
	if nik == constants.EmptyString {
		return false
	}

	if len(nik) != 16 {
		return false
	}

	matched, err := regexp.MatchString(`^\d{16}$`, nik)
	if err != nil {
		return false
	}

	return matched
}
func DecodeSignedTransaction(signedTx string) (*types.Transaction, error) {
	if signedTx[:2] == "0x" {
		signedTx = signedTx[2:]
	}
	rawTx := common.Hex2Bytes(signedTx)
	tx := new(types.Transaction)
	err := rlp.DecodeBytes(rawTx, tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func ExtractSenderAddressFromSignedTx(signedTx string) (common.Address, error) {
	tx, err := DecodeSignedTransaction(signedTx)
	if err != nil {
		return common.Address{}, err
	}

	signer := types.NewEIP155Signer(tx.ChainId())

	senderAddress, err := types.Sender(signer, tx)
	if err != nil {
		return common.Address{}, err
	}

	return senderAddress, nil
}
func ConvertToRouterCorsConfig(configCors *config.CorsConfig) *router.CorsConfig {
	return &router.CorsConfig{
		AllowOrigins:     configCors.AllowOrigins,
		AllowMethods:     configCors.AllowMethods,
		AllowHeaders:     configCors.AllowHeaders,
		AllowCredentials: configCors.AllowCredentials,
		ExposeHeaders:    configCors.ExposeHeaders,
		MaxAge:           configCors.MaxAge,
	}
}
