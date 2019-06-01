package frameworks

type Nameable interface {
	Name() string
}

type Typeable interface {
	Type() string
}

type Describable interface {
	Describe() string
}
