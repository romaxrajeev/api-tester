package structs

type ParameterData struct{
	Url string
}

type Request struct{
	Authorization string
	Method string
	Data string
	ExtraData ParameterData
}

type Response struct{
	StatusCode uint16
	Data string
}