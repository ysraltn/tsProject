package models

// InfoLog struct, bilgi amaçlı logları tutar
type InfoLog struct {
	Message     string
	UserName    string
	UserID      int
	Event       string
	ProductID   int
	ProductName string
}

// WarnLog struct, uyarı loglarını tutar
type WarnLog struct {
	Message string
	Warning string
	Context map[string]interface{}
}

// ErrorLog struct, hata loglarını tutar
type ErrorLog struct {
	Message string
	Error   string
	Context map[string]interface{}
}

type CycleInfoLog struct {
	Message   string
	Event     string
	ProductID int
	Year      int
	Month     int
	CycleNo   int
}
type InstitutionInfoLog struct {
	Message         string
	Event           string
	InstitutionName string
	City            string
}

type ProductInfoLog struct {
	Message       string
	Event         string
	ProductID     int
	ProductName   string
	ProductSerial string
	InstitutionID int
	Owner         string
	Status        string
	Responsible   string
}

type UserInfoLog struct {
	Message  string
	Event    string
	UserID   int
	UserName string
	Role     string
}

type ILogger interface {
	Info(log InfoLog)
	Warn(log WarnLog)
	Error(log ErrorLog)
	CycleInfoLog(log CycleInfoLog)
	InstitutionInfoLog(log InstitutionInfoLog)
	ProductInfoLog(log ProductInfoLog)
	UserInfoLog(log UserInfoLog)
}
