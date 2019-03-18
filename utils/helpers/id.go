package helpers

import (
	"github.com/satori/go.uuid"
	mathRand "math/rand"
	cryptoRand "crypto/rand"
	"time"
	"math/big"
	"strconv"
	"gin-frame/utils"
	"crypto/sha1"
	"fmt"
)

func Uuid4() (ret string) {
	ret = uuid.NewV4().String()
	return ret
}

func Uuid3(name string) (ret string) {
	var ns uuid.UUID
	copy(ns[:], utils.GetConfig("app::name").String())
	ret = uuid.NewV3(ns, name).String()
	return ret
}

const FullCharset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *mathRand.Rand = mathRand.New(
	mathRand.NewSource(time.Now().UnixNano()))

func RandomInt(length int) (ans int) {
	num, error := cryptoRand.Int(cryptoRand.Reader, big.NewInt(int64(length)))
	ans = int(num.Int64())
	if error != nil {
		ans = seededRand.Intn(length)
	}
	return ans
}

func RamdomStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[RandomInt(len(charset))]
	}
	return string(b)
}

func RandomString(length int, charset string) string {
	return RamdomStringWithCharset(length, charset)
}

func GenerateRandomId() (orderId string) {
	salt := RamdomStringWithCharset(50, FullCharset)
	orderId = Uuid3(strconv.FormatInt(time.Now().UnixNano(), 10) + salt)
	return orderId
}
func GenerateRandomHash() (hash string) {
	salt := RamdomStringWithCharset(50, FullCharset)
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(salt))

	return fmt.Sprintf("%x", Sha1Inst.Sum(nil))
}
