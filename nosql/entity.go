package nosql

type (
	Entity[ID any] interface {
		CollectionName() string
		GetID() ID
	}
)
