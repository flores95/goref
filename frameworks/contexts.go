package frameworks

// common contextual interfaces
type (
	// Identifier defines how a class with an ID should provide access to it's ID
	Identifier interface {
		ID() ID
	}

	// Identifyable defines how a class with an ID should set it's ID
	Identifyable interface {
		SetID(ID)
	}

	// Namer defines how a class with a name should provide access to it's name
	Namer interface {
		Name() string
	}

	// Nameable defines how a class with a name should set it's name
	Nameable interface {
		SetName(string) 
	}

	// Typer defines how a class with a type should provide access to it's type
	Typer interface {
		Type() string
	}

	// Typeable defines how a class with a type should set it's type
	Typeable interface {
		SetType(string)
	}

	// Tagger defines how to access the tags of something than can be tagged
	Tagger interface {
		Tags() []Tag
		HasTag(Tag) bool
	}

	// Taggable defines how to add a tag to something that can be tagged
	Taggable interface {
		AddTag(Tag)
	}

	// Describer defines how a class with a description should provide access to it's description
	Describer interface {
		Description() string
	}

	// Describable defines how a class with a description should set it's description
	Describable interface {
		SetDescription(string)
	}
)

type (
	// Tag represents a way to categorize things
	Tag string

	// ID represents a common
	ID string
)