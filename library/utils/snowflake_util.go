package utils

import (
"crypto/rand"
"github.com/go-kratos/kratos/pkg/log"
"math/big"
"time"
)

var (
	node      int64
	sequence  int64
	lastStamp int64
)

const (
	EPOCH         = 1288834974657	// 时间初始值
	NODE_BIT      = 10	// 机房节点占10位	1023 可区分机房与机器 各占5位
	NODE_MASK     = -1 ^ (-1 << NODE_BIT)
	SEQUENCE_BIT  = 12	// 毫秒内生成的需要占位 4095
	SEQUENCE_MASK = -1 ^ (-1 << SEQUENCE_BIT)
)

func GenerateSnowflake(tim, seqValue int64) int64 {
	if tim == 0 {
		tim = time.Now().UnixNano() / 1000000 // 精确到毫秒
	}
	if seqValue != 0 {
		sequence = seqValue
		if tim == lastStamp {
			for tim <= lastStamp {
				tim = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		// 下面是说假设在同一个毫秒内，又发送了一个请求生成一个id
		// 这个时候就得把seqence序号给递增1，最多就是4096
		if tim == lastStamp {
			// 这个意思是说一个毫秒内最多只能有4096个数字，无论你传递多少进来，
			//这个位运算保证始终就是在4096这个范围内，避免你自己传递个sequence超过了4096这个范围
			sequence = (sequence + 1) & SEQUENCE_MASK
			//当某一毫秒的时间，产生的id数 超过4095，系统会进入等待，直到下一毫秒，系统继续产生ID
			if sequence == 0 {
				// 如果当前（参数）毫秒数 小于 上一次毫秒数则继续生成
				for tim <= lastStamp {
					tim = time.Now().UnixNano() / 1000000
				}
			}
		} else {
			sequence = 0
		}
	}

	lastStamp = tim
	seq, _ := rand.Int(rand.Reader, big.NewInt(4095))
	seqInt64 := seq.Int64()

	id := (tim-EPOCH)<<(NODE_BIT+SEQUENCE_BIT) | node<<SEQUENCE_BIT | seqInt64
	log.Info("GenerateSnowflake [id:%v] [seq:%v] [tim:%v]", id, seqInt64, tim)

	return id
}

func Parse(id int64) int64 {
	return (id >> 22) + EPOCH
}
