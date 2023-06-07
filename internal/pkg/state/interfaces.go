package state

type State interface {
	Key() string
	Data() []byte
	MetaData() map[string]string
}
