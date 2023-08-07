package service

import (
	"cyberok/internal/model"
	"cyberok/internal/repository"
)

type FqdnItemService struct {
	repository repository.FqdnItem
}

func NewFqdnItemService(repository repository.FqdnItem) *FqdnItemService {
	return &FqdnItemService{repository: repository}
}

func (s *FqdnItemService) CreateFqdn(fqdn string, ip []string) (int64, error) {
	return s.repository.CreateFqdn(fqdn, ip)
}

func (s *FqdnItemService) UpdateFqdn(fqdn string, ip []string) (int64, error) {
	return s.repository.UpdateFqdn(fqdn, ip)
}

func (s *FqdnItemService) GetByIPs(ips []string) ([]model.Fqdn, error) {
	return s.repository.GetByIPs(ips)
}

func (s *FqdnItemService) GetByFQDNs(fqdns []string) ([]model.Fqdn, error) {
	return s.repository.GetByFQDNs(fqdns)
}

func (s *FqdnItemService) GetAll() ([]model.Fqdn, error) {
	return s.repository.GetAll()
}

func (s *FqdnItemService) TruncateIp() error {
	return s.TruncateIp()
}
