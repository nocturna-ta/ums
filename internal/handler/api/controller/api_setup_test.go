package controller

import (
	"github.com/mitchellh/mapstructure"
	"github.com/nocturna-ta/golib/router"
	"reflect"
	"time"
)

var (
	server *router.FastRouter
)

func initServer(opts *Options) *router.FastRouter {
	server = New(opts).RegisterRoute()
	return server
}

func toTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data any) (any, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			return time.Parse(time.RFC3339, data.(string))
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
	}
}
