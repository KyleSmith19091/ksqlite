package storage

type Engine interface {
	Get(key uint64) (string, error)
}
