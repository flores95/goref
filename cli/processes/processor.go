package processes

type Processor interface {
	Do()
	GetName() string
}
