package ca

import (
	"crypto/x509"
	"encoding/asn1"
	"net"
	"net/url"
)

// Certificate 安全起见，当前只返回证书ID
type Certificate struct {
	ID string `json:"certificate"`
}

type DistinguishedName struct {
	Type  asn1.ObjectIdentifier
	Value []interface{}
}

type Extension struct {
	ID       asn1.ObjectIdentifier
	Critical bool
	Value    []byte
}

type CertificateSigningRequest struct {
	Version int

	SubjectCountry            []string
	SubjectOrganization       []string
	SubjectOrganizationalUnit []string
	SubjectLocality           []string
	SubjectProvince           []string
	SubjectStreetAddress      []string
	SubjectPostalCode         []string
	SubjectSerialNumber       string
	SubjectCommonName         string
	SubjectExtraNames         []DistinguishedName

	PubicKeyAlg        x509.PublicKeyAlgorithm
	SignatureAlgorithm x509.SignatureAlgorithm

	// 以下四个属性构成证书中的SANs
	DNSNames       []string
	EmailAddresses []string
	IPAddresses    []net.IP
	URIs           []url.URL
	Extensions     []Extension
}
