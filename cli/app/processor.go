package app

type Processor interface {
	Do()
	GetName() string
	GetApp() App
}
