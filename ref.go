package ref

import (
	"log"
	"time"

	jsoniter "github.com/json-iterator/go"
)

func Now() *time.Time { return Time(time.Now()) }

func Time(v time.Time) *time.Time { return &v }
func String(v string) *string     { return &v }
func Int(v int) *int              { return &v }
func Uint(v uint) *uint           { return &v }
func Bool(v bool) *bool           { return &v }
func Bytes(v []byte) *[]byte      { return &v }

func MustJSON(v interface{}) *[]byte {
	data, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(v)
	if err != nil {
		log.Panicln(err)
	}
	return &data
}
