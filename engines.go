package websocket

import "github.com/centrifugal/centrifuge"

type Engine interface {
	Set(*centrifuge.Node) error
}

type MemoryEngine struct {
	MemoryBrokerConfig          centrifuge.MemoryBrokerConfig
	MemoryPresenceManagerConfig centrifuge.MemoryPresenceManagerConfig
}

func (e MemoryEngine) Set(node *centrifuge.Node) error {
	broker, err := centrifuge.NewMemoryBroker(node, e.MemoryBrokerConfig)
	if err != nil {
		return err
	}
	presenceManager, err := centrifuge.NewMemoryPresenceManager(node, e.MemoryPresenceManagerConfig)
	if err != nil {
		return err
	}
	node.SetBroker(broker)
	node.SetPresenceManager(presenceManager)
	return nil
}

type RedisEngine struct {
	RedisBrokerConfig          centrifuge.RedisBrokerConfig
	RedisPresenceManagerConfig centrifuge.RedisPresenceManagerConfig
	RedisShards                []centrifuge.RedisShardConfig
}

func (e RedisEngine) Set(node *centrifuge.Node) error {
	var redisShards []*centrifuge.RedisShard
	for i := 0; i < len(e.RedisShards); i++ {
		redisShard, err := centrifuge.NewRedisShard(node, e.RedisShards[i])
		if err != nil {
			return err
		}
		redisShards = append(redisShards, redisShard)
	}
	e.RedisBrokerConfig.Shards = redisShards
	broker, err := centrifuge.NewRedisBroker(node, e.RedisBrokerConfig)
	if err != nil {
		return err
	}
	e.RedisPresenceManagerConfig.Shards = redisShards
	presenceManager, err := centrifuge.NewRedisPresenceManager(node, e.RedisPresenceManagerConfig)
	if err != nil {
		return err
	}
	node.SetBroker(broker)
	node.SetPresenceManager(presenceManager)
	return nil
}
