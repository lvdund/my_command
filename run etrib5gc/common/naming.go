package common

import (
	"etrib5gc/sbi/models"
	"fmt"
)

func UpfServiceName(plmnid *models.PlmnId) string {
	return fmt.Sprintf("upf.%s", plmnid.String())
}

func NsmServiceName(plmnid *models.PlmnId) string {
	return fmt.Sprintf("nsm.%s", plmnid.String())
}

func DamfServiceName(plmnid *models.PlmnId) string {
	return fmt.Sprintf("damf.%s", plmnid.String())
}

func AmfServiceName(plmnid *models.PlmnId, amfid string) string {
	if len(amfid) == 0 {
		return fmt.Sprintf("amf.%s", plmnid.String())
	} else {
		return fmt.Sprintf("amf.%s.%s", plmnid.String(), amfid)
	}
}

/*
	func SmfServiceName(plmnid *models.PlmnId, slice *models.Snssai, dnn string) string {
		return fmt.Sprintf("/smf/%s/%s/%s", dnn, plmnid.String(), slice.String())
	}
*/
func SmfServiceName(plmnid *models.PlmnId, slice *models.Snssai) string {
	return fmt.Sprintf("smf.%s.%s", plmnid.String(), slice.String())
}

func UdmServiceName(plmnid *models.PlmnId) string {
	return fmt.Sprintf("udm.%s", plmnid.String())
}

func PcfServiceName(plmnid *models.PlmnId) string {
	return fmt.Sprintf("pcf.%s", plmnid.String())
}

func AusfServiceName(plmnid *models.PlmnId) string {
	return fmt.Sprintf("ausf.%s", plmnid.String())
}

func UdrServiceName(plmnid *models.PlmnId) string {
	return fmt.Sprintf("udr.%s", plmnid.String())
}

func RanServiceName(plmnid *models.PlmnId, id string) string {
	return fmt.Sprintf("pran.%s.%s", plmnid.String(), id)
}
func UpmfServiceName(plmnid *models.PlmnId) string {
	return fmt.Sprintf("upmf.%s", plmnid.String())
}
