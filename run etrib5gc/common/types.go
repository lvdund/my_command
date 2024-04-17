package common

import (
	"etrib5gc/sbi/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/free5gc/nas/security"
)

type NasSecAlgList struct {
	IntegrityOrder []byte `json:"intorder"`
	CipheringOrder []byte `json:"ciporder"`
}

var DefaultNasSecAlgs NasSecAlgList = NasSecAlgList{
	IntegrityOrder: []byte{
		//security.AlgIntegrity128NIA0,
		security.AlgIntegrity128NIA1,
		security.AlgIntegrity128NIA2,
		security.AlgIntegrity128NIA3,
	},
	CipheringOrder: []byte{
		//security.AlgCiphering128NEA0,
		security.AlgCiphering128NEA1,
		security.AlgCiphering128NEA2,
		security.AlgCiphering128NEA3,
	},
}

type PlmnItem struct {
	PlmnId models.PlmnId   `json:"plmnId"`
	Slices []models.Snssai `json:"slices"`
}

func (i *PlmnItem) String() (s string) {
	s = fmt.Sprintf("Plmn: %s", i.PlmnId.String())
	if len(i.Slices) > 0 {
		s = fmt.Sprintf("%s; Slices:", s)
		for k, slice := range i.Slices {
			if k == 0 {
				s = fmt.Sprintf("%s %s", s, slice.String())
			} else {
				s = fmt.Sprintf("%s, %s", s, slice.String())
			}
		}
	}
	return
}

type SupportedTai struct {
	Tac   string
	Plmns []PlmnItem
}
type SupportedTais []SupportedTai

func (st *SupportedTai) String() (s string) {
	s = fmt.Sprintf("TAC=%s", st.Tac)
	if len(st.Plmns) > 0 {
		s = fmt.Sprintf("%s; PLMNs=[", s)
		for k, item := range st.Plmns {
			if k == 0 {
				s = fmt.Sprintf("%s[%s]", s, item.String())
			} else {
				s = fmt.Sprintf("%s, [%s]", s, item.String())
			}
		}

		s = fmt.Sprintf("%s]", s)
	}
	return
}
func (l SupportedTais) String() (s string) {
	for _, st := range l {
		s = fmt.Sprintf("%s [%s]", s, st.String())
	}
	return
}

type AmfSet struct {
	Region uint8  `json:"region"`
	Set    uint16 `json:"set"`
}

func (s *AmfSet) String() string {
	return fmt.Sprintf("%d-%d", s.Region, s.Set)
}

func AmfInstanceId(pointer uint8) string {
	return fmt.Sprintf("%d", pointer)
}

func AmfSetFromString(s string) (id *AmfSet, err error) {
	tokens := strings.Split(s, "-")
	if len(tokens) != 2 {
		err = fmt.Errorf("Invalid AmfSet format: %s", s)
		return
	}
	value := AmfSet{}
	var v uint64
	if v, err = strconv.ParseUint(tokens[0], 10, 8); err != nil {
		return
	}
	value.Region = uint8(v)

	if v, err = strconv.ParseUint(tokens[1], 10, 16); err != nil {
		return
	}
	if v >= uint64(MAX_10_BITS) {
		err = fmt.Errorf("AmfSet must be 10bits")
		return
	}
	value.Set = uint16(v)
	id = &value
	return
}
func AmfId(amfset AmfSet, pointer uint8) string {
	return "112233"
}

/*
type AmfId struct {
	Region  uint8
	Set     uint16
	Pointer uint8
}

func (id *AmfId) SetHex(str string) (err error) {
	var buf [3]byte
	if err = loadHex(buf[:], str); err != nil {
		return
	}
	id.Region = buf[0]
	id.Set = uint16(buf[1])<<2 + (uint16(buf[2])&0x00c0)>>6
	id.Pointer = buf[2] & 0x3f
	return
}

func (id *AmfId) Bytes() (b []byte, err error) {
	var buf [3]byte
	if id.Set >= MAX_10_BITS {
		err = fmt.Errorf("AmfSet must be a 10-bit number")
		return
	}

	if id.Pointer >= MAX_6_BITS {
		err = fmt.Errorf("AmfPointer must be a 6-bit number")
		return
	}
	buf[0] = id.Region
	buf[1] = uint8(id.Set>>2) & 0xff
	buf[2] = uint8(id.Set&0x03) + id.Pointer&0x3f
	b = buf[:]
	return
}
func (id *AmfId) String() string {
	if b, err := id.Bytes(); err == nil {
		return hex.EncodeToString(b)
	}
	return ""
}
*/

func BitRate2kbps(bitrate string) uint64 {
	s := strings.Split(bitrate, " ")
	var kbps uint64

	var digit int

	if n, err := strconv.Atoi(s[0]); err != nil {
		return 0
	} else {
		digit = n
	}

	switch s[1] {
	case "bps":
		kbps = uint64(digit / 1000)
	case "Kbps":
		kbps = uint64(digit * 1)
	case "Mbps":
		kbps = uint64(digit * 1000)
	case "Gbps":
		kbps = uint64(digit * 1000000)
	case "Tbps":
		kbps = uint64(digit * 1000000000)
	}
	return kbps
}

/*
type Guami struct {
	PlmnId
	AmfId
}

func (id Guami) Bytes() (b []byte, err error) {
		var amfid _AmfId
		if err = amfid.Set(id.AmfId.Region, id.AmfId.Set, id.AmfId.Pointer); err != nil {
			return
		}
	var plmnid _PlmnId
	if err = plmnid.Set(id.PlmnId.Mcc, id.PlmnId.Mnc); err != nil {
		return
	}
	var guami _Guami
	//guami.Set(plmnid, amfid)
	b = guami[:]
	return
}

type Snssai struct {
	Sst uint8
	Sd  uint32
}

func (id Snssai) SetHex(str string) (err error) {
	var _id _Snssai
	if err = _id.SetHex(str); err != nil {
		return
	}
	id.Sst = _id.Sst()
	id.Sd = _id.Sd()
	return
}

func (id Snssai) Bytes() []byte {
	var _id _Snssai
	_id.Set(id.Sst, id.Sd)
	return _id[:]
}

type Guti struct {
	Guami
	Tmsi
}

func (id Guti) Bytes() (b []byte, err error) {
	if guami, err := id.Guami.Bytes(); err == nil {
		b = make([]byte, 10)
		copy(b, guami[:])
		copy(b[len(guami):], id.Tmsi[:])
	}
	return
}

func (id Guti) String() (str string) {
	if b, err := id.Bytes(); err == nil {
		str = hex.EncodeToString(b[:])
	}
	return
}
*/
