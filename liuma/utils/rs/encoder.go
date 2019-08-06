package rs

import (
	"github.com/klauspost/reedsolomon"
	"liuma/exception/http_err"
)

type encoder struct {
	file    []byte
	enc 	reedsolomon.Encoder
	cache 	[]byte
}

func NewEncoder (file []byte) *encoder {
	enc, _ := reedsolomon.New(RsConfig.DataShards, RsConfig.ParityShards)
	return &encoder{file, enc, nil}
}

func (this *encoder) Encode () ([][]byte, interface{}){

	shards, err := this.enc.Split(this.file); if err != nil {
		return nil, http_err.StorageUnexpectedTermination(err)
	}
	err = this.enc.Encode(shards); if err != nil {
		return nil, http_err.StorageUnexpectedTermination(err)
	}
	return shards, nil
}