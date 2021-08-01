package test

import (
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"ir.safari.shortlink/model"
	"ir.safari.shortlink/repository"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	repo, redisMock, err := getRedisMock()
	assert.Nil(t, err)

	tests := []struct {
		name    string
		arg     *model.CachedUrl
		wantErr bool
	}{
		{
			name:    "Insert: Success",
			arg: &model.CachedUrl{Key: "sampleKey"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockEntityAsJson, mockEntityAsJsonErr := jsoniter.Marshal(tt.arg)
			assert.Nil(t, mockEntityAsJsonErr)
			key := time.Now().Unix()
			fkey := fmt.Sprintf("cached_url_keys_%d", key)
			redisMock.On("Set", fkey, string(mockEntityAsJson), 0 * time.Second).Return(redis.NewStatusResult("", nil))
			if err := repo.Insert(key, tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func getRedisMock() (repository.CachedUrlRedisRepository, *redismock.ClientMock, error) {
	miniRedis, miniRedisRunErr := miniredis.Run()
	if miniRedisRunErr != nil {
		return nil, nil, miniRedisRunErr
	}
	client := redis.NewClient(&redis.Options{
		Addr: miniRedis.Addr(),
	})
	if pingErr := client.Ping().Err(); pingErr != nil {
		return nil, nil, pingErr
	}

	redisMock := redismock.NewNiceMock(client)
	redisManager := repository.NewRedisManager(redisMock)
	cachedUrlRedisRepository := repository.NewCachedUrlRedisRepository(redisManager)

	return cachedUrlRedisRepository, redisMock, nil
}
