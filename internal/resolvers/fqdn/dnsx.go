package fqdn

import (
	"cyberok/internal/config"
	"fmt"
	miekgdns "github.com/miekg/dns"
	"github.com/projectdiscovery/dnsx/libs/dnsx"
	"math"
)

type Resolver struct {
	client *dnsx.DNSX
}

func NewDnsxResolver(config *config.DNSConfig) (*Resolver, error) {
	const op = "resolvers.dnsx.new"

	options := dnsx.Options{
		BaseResolvers:     config.BaseResolvers,
		MaxRetries:        config.MaxRetries,
		QuestionTypes:     []uint16{miekgdns.TypeA},
		TraceMaxRecursion: math.MaxUint16,
		Hostsfile:         true,
	}
	client, err := dnsx.New(options)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Resolver{client: client}, nil
}

func (r *Resolver) LookupIp(host string) ([]string, error) {
	result, err := r.client.Lookup(host)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	return result, nil
}
