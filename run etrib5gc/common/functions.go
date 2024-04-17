package common

import (
	"etrib5gc/sbi/models"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/free5gc/nas/security"
)

func WrapError(msg string, err error) error {
	return fmt.Errorf("%s: %s", msg, err.Error())
}

func BearerType(access models.AccessType) uint8 {
	if access == models.ACCESSTYPE__3_GPP_ACCESS {
		return security.Bearer3GPP
	} else if access == models.ACCESSTYPE_NON_3_GPP_ACCESS {
		return security.BearerNon3GPP
	} else {
		return security.OnlyOneBearer
	}
}

func IsSliceEqual(s1, s2 *models.Snssai) bool {
	if s1 == nil && s2 == nil {
		return true
	} else if s1 == nil || s2 == nil {
		return false
	}

	return s1.Sst == s2.Sst && strings.Compare(s1.Sd, s2.Sd) == 0
}

func IsPlmnIdEqual(id1, id2 *models.PlmnId) bool {
	if id1 == nil && id2 == nil {
		return true
	} else if id1 == nil || id2 == nil {
		return false
	}

	return strings.Compare(id1.Mnc, id2.Mnc) == 0 && strings.Compare(id1.Mcc, id2.Mcc) == 0
}
func ServingNetworkName(id *models.PlmnId) string {
	//return fmt.Sprintf("5G:mnc%03x.mcc%03x.3gppnetwork.org", id.Mnc, id.Mcc)
	if len(id.Mnc) == 2 {
		return fmt.Sprintf("5G:mnc0%s.mcc%s.3gppnetwork.org", id.Mnc, id.Mcc)
	} else {
		return fmt.Sprintf("5G:mnc%s.mcc%s.3gppnetwork.org", id.Mnc, id.Mcc)
	}
}

func Prob2Err(prob *models.ProblemDetails) error {
	return fmt.Errorf("%d: %s", prob.Status, prob.Detail)
}

func SmContextRef(supi string, sid uint32) string {
	return fmt.Sprintf("supi=%s-sid=%d", supi, sid)
}

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Warnf("Network not available: %+v", err)
		return net.ParseIP("127.0.0.1")
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}

func GetPublicIP() (ip net.IP) {
	req, err := http.Get("https://ipinfo.io/ip")
	if err != nil {
		log.Errorf("Public IP not discoverable: %+v", err)
		return
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Errorf("Public IP not discoverable: %+v", err)
	}
	ip = net.ParseIP(string(body))
	return
}

// create a map from a list
func MakeSet[T comparable](items []T) (m map[T]struct{}) {
	m = make(map[T]struct{})
	for _, t := range items {
		if _, ok := m[t]; !ok {
			m[t] = struct{}{}
		}
	}
	return
}

type Stringer interface {
	String() string
}

// create a string-indexed map from a list of Stringer items
func MakeStringSet[T Stringer](items []T) (m map[string]T) {
	m = make(map[string]T)
	for _, t := range items {
		if _, ok := m[t.String()]; !ok {
			m[t.String()] = t
		}
	}
	return
}

// make a list of unique item from a list
func UniqueArray[T comparable](items []T) (ret []T) {
	itemset := MakeSet[T](items)
	for i, _ := range itemset {
		ret = append(ret, i)
	}
	return
}

func L4Address(ip net.IP, port int) (addr string) {
	if len(ip) > 0 {
		addr = fmt.Sprintf("%s:%d", ip.String(), port)
	}
	return
}

// return an non-empty IP if success
func ParseL4Address(str string) (ip net.IP, port int) {
	toks := strings.Split(str, ":")
	if len(toks) == 2 {
		var err error
		if port, err = strconv.Atoi(toks[1]); err != nil {
			return
		}
		ip = net.ParseIP(toks[0])
	}
	return
}
