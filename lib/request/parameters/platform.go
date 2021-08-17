package parameters

func (params Query) SetNintendoSwitchPlatform() {
	params.setPlatform("HAC")
}

func (params Query) SetNintendoNew2DSXLPlatform() {
	params.setPlatform("KTR")
}

func (params Query) SetNintendoNew3DSPlatform() {
	params.setPlatform("KTR")
}

func (params Query) SetNintendoNew3DSXLPlatform() {
	params.setPlatform("KTR")
}

func (params Query) SetNintendo2DSPlatform() {
	params.setPlatform("CTR")
}

func (params Query) SetNintendo3DSPlatform() {
	params.setPlatform("CTR")
}

func (params Query) SetNinteno3DSXLPlatform() {
	params.setPlatform("CTR")
}

func (params Query) SetSmartDevicePlatform() {
	params.setPlatform("ZSD")
}

func (params Query) SetWiiUPlatform() {
	params.setPlatform("WUP")
}

func (params Query) SetWiiPlatform() {
	params.setPlatform("RVL")
}

func (params Query) SetDSiPlatform() {
	params.setPlatform("TWL")
}

func (params Query) SetDSPlatform() {
	params.setPlatform("NTR")
}

func (params Query) setPlatform(value string) {
	params.setRequestParameter("playable_on_txt", value, false)
}
