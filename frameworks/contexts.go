package frameworks

// Nameable defines how a class with a name should provide access to it's name
type Nameable interface {
	Name() string
}

// Typeable defines how a class with a type should provide access to it's type
type Typeable interface {
	Type() string
}

// Describable defines how a class with a description should provide access to it's description
type Describable interface {
	Describe() string
}
