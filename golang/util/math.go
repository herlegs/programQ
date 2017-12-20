package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/myteksi/go/commons/util/ioclose"
)

const (
	defaultRandomInt64Limit   = int64(10000)
	defaultRandomInt32Limit   = int32(10000)
	defaultRandomStringLength = int(6)
)

func MinInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func MinInts(a int, others ...int) int {
	result := a
	for _, other := range others {
		if other < result {
			result = other
		}
	}
	return result
}

func MaxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func MinInt64(a, b int64) int64 {
	return int64(math.Min(float64(a), float64(b)))
}

func MaxInt64(a, b int64) int64 {
	return int64(math.Max(float64(a), float64(b)))
}

//SprintObj Sprint any struct object
func SprintObj(obj interface{}) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("err while marshaling:", err)
	}
	return string(bytes)
}

// GetRandomInt64 : helper function to get random int64
// len(args) == 0: get random between(inclusive) 1 ~ defaultRandomInt64Limit
// len(args) == 1: get random between(inclusive) 1 ~ args[0] (can ge negative)
// len(args) >= 2: get random between(inclusive) args[0] ~ args[1]
func GetRandomInt64(args ...int64) int64 {
	rand.Seed(time.Now().UTC().UnixNano())
	var minV int64
	var maxV int64
	if len(args) == 0 {
		maxV = defaultRandomInt64Limit
	} else if len(args) == 1 {
		minV = int64(math.Min(0, float64(args[0])))
		maxV = int64(math.Max(0, float64(args[0])))
	} else {
		minV = int64(math.Min(float64(args[0]), float64(args[1])))
		maxV = int64(math.Max(float64(args[0]), float64(args[1])))
	}

	return rand.Int63n(maxV-minV) + minV + 1
}

// GetRandomInt32 : helper function to get random int32
// len(args) == 0: get random between(inclusive) 1 ~ defaultRandomInt64Limit
// len(args) == 1: get random between(inclusive) 1 ~ args[0] (can ge negative)
// len(args) >= 2: get random between(inclusive) args[0] ~ args[1] (inclusive)
func GetRandomInt32(args ...int32) int32 {
	rand.Seed(time.Now().UTC().UnixNano())
	var minV int32
	var maxV int32
	if len(args) == 0 {
		maxV = defaultRandomInt32Limit
	} else if len(args) == 1 {
		minV = int32(math.Min(0, float64(args[0])))
		maxV = int32(math.Max(0, float64(args[0])))
	} else {
		minV = int32(math.Min(float64(args[0]), float64(args[1])))
		maxV = int32(math.Max(float64(args[0]), float64(args[1])))
	}

	return rand.Int31n(maxV-minV) + minV + 1
}

// GetRandomString returns random string of given length
func GetRandomString(lengths ...int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	var length int
	if len(lengths) == 0 {
		length = defaultRandomStringLength
	} else {
		length = lengths[0]
	}
	wordDict := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	return GetRandomStringWithDict(length, wordDict)
}

// GetRandomStringWithDict returns random string of given length with custom dictionary
func GetRandomStringWithDict(length int, wordDict string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	l := len(wordDict)
	b := bytes.NewBuffer(nil)
	for i := 0; i < length; i++ {
		pos := rand.Intn(l)
		_, _ = b.WriteString(string(wordDict[pos]))
	}
	return b.String()
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}

// RoundTime will round a time to certain accuracy
func RoundTime(t time.Time, d time.Duration) time.Time {
	sec := t.Unix()
	nanoOffset := t.Nanosecond()

	div := float64(nanoOffset) / float64(d)
	div = round(div)

	roundedNanoOffset := int64(div) * int64(d)

	return time.Unix(sec, roundedNanoOffset).UTC()
}

// LoadJSONFromFile read json file and parse it to object
func LoadJSONFromFile(filePath string, obj interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer ioclose.Safely(file)

	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(obj); err != nil {
		return err
	}
	return nil
}

func AbsInt64(i int64) int64 {
	return int64(math.Abs(float64(i)))
}

// GetLenOfInt64 returns number of digits of an integer (including sign)
func GetLenOfInt64(i int64) int {
	str := strconv.FormatInt(i, 10)
	return len(str)
}
