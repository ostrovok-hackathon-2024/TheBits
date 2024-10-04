//package main
//
//import (
//	"context"
//	"encoding/json"
//)
//
//type Producer interface {
//	Produce(ctx context.Context, data interface{}) (err error)
//}
//
//type Consumer interface {
//	Consumer()
//}
//
//// --
//
//type producer struct {
//	redisClient any
//	topic       string
//}
//
//func NewProducer(redisClient any, topic string) Producer {
//	return &producer{
//		topic: topic,
//	}
//}
//
//func (p *producer) Produce(ctx context.Context, data interface{}) (err error) {
//	key := "gen-uuid" // uuid.NewString()
//
//	//json.Marshal(data)
//
//	p.redisClient.HSet(ctx, key, data)
//	p.redisClient.LPush(ctx, p.topic, key)
//
//	return
//}
//
