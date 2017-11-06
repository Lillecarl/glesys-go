package glesys

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/net/idna"
)

// DomainService provides functions to interact with domains
type DomainService struct {
	client clientInterface
}

// Domain represents a domain
type Domain struct {
	DomainName    string    `json:"domainname"`
	CreateTime    time.Time `json:"createtime"`
	DisplayName   string    `json:"displayname"`
	RecordCount   int       `json:"recordcount"`
	RegistrarInfo struct {
		State            string      `json:"state"`
		StateDescription string      `json:"statedescription"`
		Expire           string      `json:"expire"`
		AutoRenew        string      `json:"autorenew"`
		TLD              string      `json:"tld"`
		InvoiceNumber    interface{} `json:"invoicenumber"`
	} `json:"registrarinfo"`
}

// Domainrecord represents a DNS-record
type DomainRecord struct {
	Recordid   int    `json:"recordid"`
	Domainname string `json:"domainname"`
	Host       string `json:"host"`
	Type       string `json:"type"`
	Data       string `json:"data"`
	TTL        int    `json:"ttl"`
}

// AddDomainParams is used when creating a new domain
type AddDomainParams struct {
	DomainName        string `json:"domainname"`
	PrimaryNameServer string `json:"primarynameserver,omitempty"`
	ResponsiblePerson string `json:"responsibleperson,omitempty"`
	TTL               string `json:"ttl,omitempty"`
	Refresh           string `json:"refresh,omitempty"`
	Retry             string `json:"retry,omitempty"`
	Expire            string `json:"expire,omitempty"`
	Minimum           string `json:"minimum,omitempty"`
	CreateRecords     int    `json:"createrecords"`
}

// AddDomainRecordParams is used when creating a new domain
type AddDomainRecordParams struct {
	DomainName string `json:"domainname"`
	Host       string `json:"host"`
	Type       string `json:"type"`
	Data       string `json:"data"`
	TTL        int    `json:"ttl,omitempty"`
}

// EditDomainParams is used when updating an existing domain
type EditDomainParams struct {
	DomainName        string `json:"domainname"`
	PrimaryNameServer string `json:"primarynameserver"`
	ResponsiblePerson string `json:"responsibleperson"`
	TTL               string `json:"ttl"`
	Refresh           string `json:"refresh"`
	Retry             string `json:"retry"`
	Expire            string `json:"expire"`
	Minimum           string `json:"minimum"`
}

// UpdateDomainRecordParams is used when updating an existing DNS record
type UpdateDomainRecordParams struct {
	RecordId int    `json:"domainname"`
	Host     string `json:"primarynameserver"`
	Type     string `json:"responsibleperson"`
	Data     string `json:"ttl"`
	TTL      int    `json:"refresh"`
}

// Add adds a new domain
func (s *DomainService) AddDomain(context context.Context, params AddDomainParams) (*Domain, error) {
	params.DomainName, _ = idna.New().ToASCII(params.DomainName)
	data := struct {
		Response struct {
			Domain Domain
		}
	}{}
	err := s.client.post(context, "domain/add", &data, params)
	return &data.Response.Domain, err
}

// Add adds a new DNS record
func (s *DomainService) AddRecord(context context.Context, params AddDomainRecordParams) (*DomainRecord, error) {
	data := struct {
		Response struct {
			DomainRecord DomainRecord
		}
	}{}
	err := s.client.post(context, "domain/addrecord", &data, params)
	return &data.Response.DomainRecord, err
}

// Details returns detailed information about one domain
func (s *DomainService) Details(context context.Context, domainname string) (*Domain, error) {
	domainname, _ = idna.New().ToASCII(domainname)
	data := struct {
		Response struct {
			Domain Domain
		}
	}{}
	err := s.client.get(context, fmt.Sprintf("domain/details/domainname/%s", domainname), &data)
	return &data.Response.Domain, err
}

func (s *DomainService) DeleteDomain(context context.Context, domainname string) error {
	domainname, _ = idna.New().ToASCII(domainname)
	return s.client.post(context, "domain/delete", nil, struct {
		DomainName string `json:"domainname"`
	}{domainname})
}

func (s *DomainService) DeleteRecord(context context.Context, recordID int) error {
	return s.client.post(context, "domain/deleterecord", nil, struct {
		recordID int `json:"recordid"`
	}{recordID})
}

// Edit modifies a domain
func (s *DomainService) EditDomain(context context.Context, domainname string, params EditDomainParams) (*Domain, error) {
	domainname, _ = idna.New().ToASCII(domainname)
	params.DomainName, _ = idna.New().ToASCII(params.DomainName)
	data := struct {
		Response struct {
			Domain Domain
		}
	}{}
	err := s.client.post(context, "domain/edit", &data, struct {
		EditDomainParams
		DomainName string `json:"domainname"`
	}{params, domainname})
	return &data.Response.Domain, err
}

// Edit modifies a domain
func (s *DomainService) UpdateRecord(context context.Context, recordID int, params UpdateDomainRecordParams) (*DomainRecord, error) {
	data := struct {
		Response struct {
			DomainRecord DomainRecord
		}
	}{}
	err := s.client.post(context, "domain/updaterecord", &data, struct {
		UpdateDomainRecordParams
		recordID int `json:"recordid"`
	}{params, recordID})
	return &data.Response.DomainRecord, err
}

// List returns a list of domains available under your account
func (s *DomainService) List(context context.Context) (*[]Domain, error) {
	data := struct {
		Response struct {
			Domains []Domain
		}
	}{}

	err := s.client.get(context, "domain/list", &data)
	return &data.Response.Domains, err
}

func (s *DomainService) ListRecords(context context.Context, domainname string) (*[]DomainRecord, error) {
	data := struct {
		Response struct {
			DomainRecord []DomainRecord
		}
	}{}

	err := s.client.post(context, "domain/listrecords", &data, struct {
		DomainName string `json:"domainname"`
	}{domainname})

	return &data.Response.DomainRecord, err
}
