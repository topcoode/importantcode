package main

type AmfUe struct {
	RanUe map[AccessType]*RanUe
}

type AccessType string

// List of AccessType
const (
	AccessType__3_GPP_ACCESS    AccessType = "3GPP_ACCESS"
	AccessType_NON_3_GPP_ACCESS AccessType = "NON_3GPP_ACCESS"
)

func main() {

}
func (ue *AmfUe) CmConnect(anType AccessType) bool {
	if _, ok := ue.RanUe[anType]; !ok {
		return false
	}
	return true
}
