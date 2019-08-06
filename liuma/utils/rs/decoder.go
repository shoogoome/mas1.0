package rs

import (
	"bytes"
	"github.com/klauspost/reedsolomon"
	"liuma/exception/http_err"
	"liuma/models"
	"liuma/utils"
)

type decoder struct {
	shards    [][]byte
	enc 	reedsolomon.Encoder
	serverIP  []string
}

func NewDecoder (shards [][]byte, serverIP []string) *decoder {
	enc, _ := reedsolomon.New(RsConfig.DataShards, RsConfig.ParityShards)
	return &decoder{shards, enc, serverIP}
}

func (this *decoder) Decode (token string) ([]byte, interface{}){
	// 检查数据分片是否正常，不正常则尝试修复数据
	ok, err := this.enc.Verify(this.shards); if !ok {

		unHealth := this.healthExamination()
		err = this.enc.Reconstruct(this.shards)

		if err != nil {
			return nil, http_err.DamageToRawData()
		} else {
			// 修复分片数据后将新分片传输到指定服务
			// 因获取时shards顺序已与serverIP顺序对应 所以可以直接按序操作
			var statusMap = make(chan models.ShardsStatus, RsConfig.AllShards)
			for _, index := range unHealth {
				go utils.SendFileShard(
					this.serverIP[index],
					this.shards[index],
					index,
					token,
					statusMap,
					)
			}
		}
		ok, err = this.enc.Verify(this.shards)
		if !ok {
			return nil, http_err.DamageToRawData()
		}
	}
	var dd bytes.Buffer
	err = this.enc.Join(&dd, this.shards, len(this.shards[0]) * RsConfig.DataShards); if err != nil {
		return nil, http_err.DamageToRawData()
	}
	return dd.Bytes(), nil
}

// 因为实际使用上分片数量有限 没必要做算法优化
func (this decoder) max () int {
	length := len(this.shards[0])
	for _, shard := range this.shards {
		if length < len(shard) {
			length = len(shard)
		}
	}
	return length
}

func (this decoder) healthExamination () []int {

	unHealth := make([]int, 0, RsConfig.AllShards)
	maxLength := this.max()

	for index, shard := range this.shards {
		if shard == nil || len(shard) < maxLength {
			unHealth = append(unHealth, index)
		}
	}
	return unHealth
}





