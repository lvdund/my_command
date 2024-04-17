package common

import "etrib5gc/sbi/models"

type UpfDnInfo struct {
	//Dnn  string
	Cidr string
	Addr string
}

type UpfInfInfo struct {
	Addr string
	Mtu  uint32
}

type UpfInfo struct {
	Slices map[string]models.Snssai
	DNs    map[string]UpfDnInfo
	INFs   map[string]UpfInfInfo
	Active bool
}
