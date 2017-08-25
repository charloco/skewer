package logging

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/inconshreveable/log15"
	"github.com/tinylib/msgp/msgp"
)

type RemoteLoggerHandler struct {
	remote  net.Conn
	msgChan chan *log15.Record
	ctx     context.Context
}

func NewRemoteLogger(ctx context.Context, remote net.Conn) log15.Logger {
	logger := log15.New()
	h := RemoteLoggerHandler{remote: remote, ctx: ctx}
	h.msgChan = make(chan *log15.Record, 1000)
	logger.SetHandler(&h)

	go func() {
		defer h.Close()
		var rbis Record
		var r *log15.Record
		var more bool
		rem := msgp.NewWriter(h.remote)
		for {
			select {
			case <-ctx.Done():
				return
			case r, more = <-h.msgChan:
				if more {
					rbis = Record{Time: r.Time, Lvl: int(r.Lvl), Msg: r.Msg}

					rbis.Ctx = map[string]string{}

					l := len(r.Ctx)
					var i int
					var ok bool
					label := ""
					val := ""

					for i < l {
						label, ok = r.Ctx[i].(string)
						if ok {
							i++
							if i < l {
								val = formatValue(r.Ctx[i])
								rbis.Ctx[label] = val
								i++
							}
						}

					}
					rbis.EncodeMsg(rem)
					rem.Flush()
				}
			}
		}
	}()

	return logger
}

func (h *RemoteLoggerHandler) Log(r *log15.Record) error {
	select {
	case h.msgChan <- r:
	case <-h.ctx.Done():
	}
	return nil
}

func (h *RemoteLoggerHandler) Close() error {
	return h.remote.Close()
}

const timeFormat = "2006-01-02T15:04:05-0700"

func formatShared(value interface{}) (result interface{}) {
	switch v := value.(type) {
	case time.Time:
		return v.Format(timeFormat)

	case error:
		return v.Error()

	case fmt.Stringer:
		return v.String()

	default:
		return v
	}
}

func formatValue(value interface{}) string {
	value = formatShared(value)
	switch v := value.(type) {
	case string:
		return v
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", value)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', 3, 64)
	case float64:
		return strconv.FormatFloat(v, 'f', 3, 64)
	default:
		return fmt.Sprintf("%+v", value)
	}
}
