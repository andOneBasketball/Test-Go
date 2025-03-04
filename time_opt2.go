package main

import (
	"reflect"
	"time"

	"go.uber.org/zap"
)

func main() {
	for {
		select {
		case message := <-config.UserEventChan:
			userEventLog, ok := message.(*sqlmodel.UserEventLog)
			if !ok {
				zap.S().Warnf("Invalid message type: %T, message: %#v", reflect.TypeOf(message), message)
				continue
			}
			buffer = append(buffer, userEventLog)
			if len(buffer) >= config.UserEventBatchSize {
				dao.AddBatchUserEventLog(ctx, buffer)
				buffer = buffer[:0]
			}
		case <-time.After(time.Second * 1):
			if len(buffer) > 0 {
				buffer = buffer[:0]
				dao.AddBatchUserEventLog(ctx, buffer)
			}
		}
	}
}
