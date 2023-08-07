package service

import (
	"cyberok/internal/model"
	"cyberok/internal/repository"
)

type WhoisItemService struct {
	repository repository.WhoisItem
}

func NewWhoisItemService(repository repository.WhoisItem) *WhoisItemService {
	return &WhoisItemService{repository: repository}
}

func (s *WhoisItemService) UpdateWhois(domain string, whois string) error {
	return s.UpdateWhois(domain, whois)
}

func (s *WhoisItemService) CreateWhois(domain string, whois string) (int64, error) {
	return s.repository.CreateWhois(domain, whois)
}

func (s *WhoisItemService) GetByDomains(domains []string) ([]model.Whois, error) {
	return s.repository.GetByDomains(domains)
}
