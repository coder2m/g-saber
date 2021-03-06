package xcache

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	nilError = errors.New("data is nil")
)

type (
	Handle interface {
		//
		// Create
		// @Description: 创建数据
		// @return []byte
		// @return error
		//
		Create(ctx context.Context) ([]byte, error)
		//
		// Expire
		// @Description: 设置时间过期时间
		// @return time.Duration
		//
		Expire() time.Duration
	}

	Cache interface {
		//
		// Del
		// @Description: 删除缓存
		// @param keys ...string
		// @return error
		//
		Del(keys ...string) error
		//
		// GetE
		// @Description: 获取缓存
		// @param key string
		// @return error
		//
		GetE(key string) ([]byte, error)

		Get(key string) []byte

		GetWithCreateE(key string, h Handle) ([]byte, error)
		//
		// GetWithCreate 获取缓存，如果没有就使用Handle创建并放回
		// @Description:
		// @param key string
		// @param v interface{}
		// @param h Handle
		// @return error
		//
		GetWithCreate(key string, h Handle) []byte
		//
		// Set
		// @Description: 设置缓存
		// @param key string
		// @param h Handle
		// @return error
		//
		Set(key string, h Handle) error
		//
		// IsNilError
		// @Description: 判断空数据 error
		// @param err error
		// @return bool
		//
		IsNilError(err error) bool
		//
		// IsExist
		// @Description: 判断存在
		// @param key string
		// @return bool
		//
		IsExist(key string) bool

		//
		// WithContext
		// @Description:
		// @param ctx context.Context
		// @return Cache
		//
		WithContext(ctx context.Context) Cache

		GetContext() context.Context
	}

	Basis struct {
		lock sync.Mutex
	}
)
