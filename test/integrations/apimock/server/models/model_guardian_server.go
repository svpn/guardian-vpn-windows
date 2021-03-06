/*
 * Firefox Guardian API
 *
 * API to manage Guardian accounts, devices and servers
 *
 * API version: 0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type GuardianServer struct {
	Countries []Country `json:"countries,omitempty"`
}

type Country struct {
	Name   string `json:"name,omitempty"`
	Code   string `json:"code,omitempty"`
	Cities []City `json:"cities,omitempty"`
}

type City struct {
	Name      string   `json:"name,omitempty"`
	Code      string   `json:"code,omitempty"`
	Latitude  float32  `json:"latitude,omitempty"`
	Longitude float32  `json:"longitude,omitempty"`
	Servers   []Server `json:"servers,omitempty"`
}

type Server struct {
	Hostname         string  `json:"hostname,omitempty"`
	Ipv4AddrIn       string  `json:"ipv4_addr_in,omitempty"`
	Weight           int     `json:"weight,omitempty"`
	IncludeInCountry bool    `json:"include_in_country,omitempty"`
	PublicKey        string  `json:"public_key,omitempty"`
	PortRanges       [][]int `json:"port_ranges,omitempty"`
	Ipv4Gateway      string  `json:"ipv4_gateway,omitempty"`
	Ipv6Gateway      string  `json:"ipv6_gateway,omitempty"`
}
