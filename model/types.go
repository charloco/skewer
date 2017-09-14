package model

import "time"
import "fmt"
import "encoding/json"
import sarama "gopkg.in/Shopify/sarama.v1"

//go:generate msgp

type Priority int
type Facility int
type Severity int
type Version int

type AuditSubMessage struct {
	Type uint16 `json:"type" msg:"type"`
	Data string `json:"data" msg:"data"`
}

type AuditMessageGroup struct {
	Seq       int                `json:"sequence" msg:"sequence"`
	AuditTime string             `json:"timestamp" msg:"timestamp"`
	Msgs      []*AuditSubMessage `json:"messages" msg:"messages"`
	UidMap    map[string]string  `json:"uid_map" msg:"uid_map"`
}

type SyslogMessage struct {
	Priority         Priority                     `json:"priority,string" msg:"priority"`
	Facility         Facility                     `json:"facility,string" msg:"facility"`
	Severity         Severity                     `json:"severity,string" msg:"severity"`
	Version          Version                      `json:"version,string" msg:"version"`
	TimeReported     time.Time                    `json:"timereported,omitempty" msg:"timereported"`
	TimeGenerated    time.Time                    `json:"timegenerated,omitempty" msg:"timegenerated"`
	Hostname         string                       `json:"hostname" msg:"hostname"`
	Appname          string                       `json:"appname" msg:"appname"`
	Procid           string                       `json:"procid,omitempty" msg:"procid"`
	Msgid            string                       `json:"msgid,omitempty" msg:"msgid"`
	Structured       string                       `json:"structured,omitempty" msg:"structured"`
	Message          string                       `json:"message" msg:"message"`
	AuditSubMessages []*AuditSubMessage           `json:"audit,omitempty" msg:"audit"`
	Properties       map[string]map[string]string `json:"properties,omitempty" msg:"properties"`
}

type ParsedMessage struct {
	Fields         *SyslogMessage `json:"fields" msg:"fields"`
	Client         string         `json:"client,omitempty" msg:"client"`
	LocalPort      int            `json:"local_port,string" msg:"local_port"`
	UnixSocketPath string         `json:"unix_socket_path,omitempty" msg:"unix_socket_path"`
}

type TcpUdpParsedMessage struct {
	Parsed *ParsedMessage `json:"parsed" msg:"parsed"`
	Uid    string         `json:"uid" msg:"uid"`
	ConfId string         `json:"conf_id" msg:"conf_id"`
}

type RelpParsedMessage struct {
	Parsed *ParsedMessage `json:"parsed" msg:"parsed"`
	Txnr   int            `json:"txnr" msg:"txnr"`
}

func (m *ParsedMessage) ToKafkaMessage(partitionKey string, topic string) (km *sarama.ProducerMessage, err error) {
	value, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	kafka_msg := sarama.ProducerMessage{
		Key:       sarama.StringEncoder(partitionKey),
		Value:     sarama.ByteEncoder(value),
		Topic:     topic,
		Timestamp: m.Fields.TimeReported,
	}
	return &kafka_msg, nil
}

var SyslogMessageFmt string = `Facility: %d
Severity: %d
Version: %d
TimeReported: %s
TimeGenerated: %s
Hostname: %s
Appname: %s
ProcID: %s
MsgID: %s
Structured: %s
Message: %s
AuditSubMessages: %s
Properties: %s`

func (m *SyslogMessage) String() string {
	props := ""
	b, err := json.Marshal(m.Properties)
	if err == nil {
		props = string(b)
	}
	subs := ""
	b, err = json.Marshal(m.AuditSubMessages)
	if err == nil {
		subs = string(b)
	}
	return fmt.Sprintf(
		SyslogMessageFmt,
		m.Facility,
		m.Severity,
		m.Version,
		m.TimeReported.Format(time.RFC3339),
		m.TimeGenerated.Format(time.RFC3339),
		m.Hostname,
		m.Appname,
		m.Procid,
		m.Msgid,
		m.Structured,
		m.Message,
		subs,
		props,
	)
}