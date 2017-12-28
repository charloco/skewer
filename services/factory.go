package services

import (
	"fmt"
	"os"

	"github.com/inconshreveable/log15"
	"github.com/stephane-martin/skewer/conf"
	"github.com/stephane-martin/skewer/model"
	"github.com/stephane-martin/skewer/services/base"
	"github.com/stephane-martin/skewer/services/linux"
	"github.com/stephane-martin/skewer/services/network"
	"github.com/stephane-martin/skewer/sys/binder"
	"github.com/stephane-martin/skewer/sys/kring"
)

type Types int

const (
	TCP Types = iota
	UDP
	RELP
	Journal
	Store
	Accounting
	KafkaSource
	Configuration
	Graylog
)

var Names2Types map[string]Types = map[string]Types{
	"skewer-tcp":         TCP,
	"skewer-udp":         UDP,
	"skewer-relp":        RELP,
	"skewer-journal":     Journal,
	"skewer-store":       Store,
	"skewer-accounting":  Accounting,
	"skewer-kafkasource": KafkaSource,
	"skewer-conf":        Configuration,
	"skewer-graylog":     Graylog,
}

var Types2Names map[Types]string
var Types2ConfinedNames map[Types]string

var Handles []ServiceHandle
var HandlesMap map[ServiceHandle]uintptr

type HandleType uint8

const (
	BINDER HandleType = iota
	LOGGER
)

type ServiceHandle struct {
	Service string
	Type    HandleType
}

func init() {
	Types2Names = map[Types]string{}
	Types2ConfinedNames = map[Types]string{}
	for k, v := range Names2Types {
		Types2Names[v] = k
		Types2ConfinedNames[v] = "confined-" + k
	}

	Handles = []ServiceHandle{
		{"child", BINDER},
		{Types2Names[TCP], BINDER},
		{Types2Names[UDP], BINDER},
		{Types2Names[RELP], BINDER},
		{Types2Names[Graylog], BINDER},
		{"child", LOGGER},
		{Types2Names[TCP], LOGGER},
		{Types2Names[UDP], LOGGER},
		{Types2Names[RELP], LOGGER},
		{Types2Names[Journal], LOGGER},
		{Types2Names[Configuration], LOGGER},
		{Types2Names[Store], LOGGER},
		{Types2Names[Accounting], LOGGER},
		{Types2Names[KafkaSource], LOGGER},
		{Types2Names[Graylog], LOGGER},
	}

	HandlesMap = map[ServiceHandle]uintptr{}
	for i, h := range Handles {
		HandlesMap[h] = uintptr(i + 3)
	}
}

func LoggerHdl(typ Types) uintptr {
	return HandlesMap[ServiceHandle{Types2Names[typ], LOGGER}]
}

func BinderHdl(typ Types) uintptr {
	return HandlesMap[ServiceHandle{Types2Names[typ], BINDER}]
}

func ConfigureAndStartService(s Provider, c conf.BaseConfig) ([]model.ListenerInfo, error) {

	switch s := s.(type) {
	case *network.TcpServiceImpl:
		s.SetConf(c.TcpSource, c.Parsers, c.Main.InputQueueSize, c.Main.MaxInputMessageSize)
		return s.Start()
	case *network.UdpServiceImpl:
		s.SetConf(c.UdpSource, c.Parsers, c.Main.InputQueueSize)
		return s.Start()
	case *network.RelpService:
		s.SetConf(c.RelpSource, c.Parsers, c.Main.InputQueueSize)
		return s.Start()
	case *network.GraylogSvcImpl:
		s.SetConf(c.GraylogSource)
		return s.Start()
	case *linux.JournalService:
		s.SetConf(c.Journald)
		return s.Start()
	case *AccountingService:
		s.SetConf(c.Accounting)
		return s.Start()
	case *storeServiceImpl:
		return s.SetConfAndRestart(c)
	case *network.KafkaServiceImpl:
		s.SetConf(c.KafkaSource, c.Parsers, c.Main.InputQueueSize)
		return s.Start()
	default:
		return nil, fmt.Errorf("Unknown network service: %T", s)
	}

}

func ProviderFactory(t Types, confined bool, r kring.Ring, reporter base.Reporter, b *binder.BinderClientImpl, l log15.Logger, pipe *os.File) Provider {
	switch t {
	case TCP:
		return network.NewTcpService(reporter, confined, b, l)
	case UDP:
		return network.NewUdpService(reporter, b, l)
	case RELP:
		return network.NewRelpService(reporter, confined, b, l)
	case Graylog:
		return network.NewGraylogService(reporter, b, l)
	case Journal:
		svc, err := linux.NewJournalService(reporter, l)
		if err == nil {
			return svc
		} else {
			l.Error("Error creating the journal service", "error", err)
			return nil
		}
	case Accounting:
		svc, err := NewAccountingService(reporter, confined, l)
		if err == nil {
			return svc
		} else {
			l.Error("Error creating the accounting service", "error", err)
			return nil
		}
	case Store:
		return NewStoreService(confined, l, r, pipe)
	case KafkaSource:
		return network.NewKafkaService(reporter, confined, l)
	default:
		fmt.Fprintf(os.Stderr, "Unknown service type: %d\n", t)
		return nil
	}
}
