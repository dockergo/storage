package adapter

import "github.com/flyaways/storage/agent/util/log"

func (s *StorageAdapter) InitBucket(initdata *InitData) error {
	log.Warn("[%s InitBucket NotImplemented]", s.Name)
	Details()
	return nil
}

func (s *StorageAdapter) InitObject(initdata *InitData) error {
	log.Warn("[%s InitObject NotImplemented]", s.Name)
	Details()
	return nil
}
