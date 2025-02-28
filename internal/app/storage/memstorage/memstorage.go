package memstorage

type MemStorage struct {
}

func New() *MemStorage {
	return &MemStorage{}
}

func (s *MemStorage) Close() error {
	return nil
}

func (s *MemStorage) Ping() error {
	return nil
}
