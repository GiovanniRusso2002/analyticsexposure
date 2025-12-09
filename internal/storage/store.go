package storage

import (
	"errors"
	"sync"

	"github.com/GiovanniRusso2002/analyticsexposure/internal/models"
)

// Store interface defines the storage operations
type Store interface {
	CreateSubscription(afID, subscriptionID string, subsc *models.AnalyticsExposureSubsc) error
	GetSubscription(afID, subscriptionID string) (*models.AnalyticsExposureSubsc, error)
	GetAllSubscriptions(afID string) ([]*models.AnalyticsExposureSubsc, error)
	UpdateSubscription(afID, subscriptionID string, subsc *models.AnalyticsExposureSubsc) error
	DeleteSubscription(afID, subscriptionID string) error
	GetAnalyticsData(afID string, req *models.AnalyticsRequest) (*models.AnalyticsData, error)
}

// InMemoryStore is an in-memory implementation of Store
type InMemoryStore struct {
	subscriptions map[string]map[string]*models.AnalyticsExposureSubsc
	mu            sync.RWMutex
	analyticsData map[string]*models.AnalyticsData
}

// NewInMemoryStore creates a new in-memory store
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		subscriptions: make(map[string]map[string]*models.AnalyticsExposureSubsc),
		analyticsData: make(map[string]*models.AnalyticsData),
	}
}

// CreateSubscription creates a new subscription
func (s *InMemoryStore) CreateSubscription(afID, subscriptionID string, subsc *models.AnalyticsExposureSubsc) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.subscriptions[afID]; !exists {
		s.subscriptions[afID] = make(map[string]*models.AnalyticsExposureSubsc)
	}

	if _, exists := s.subscriptions[afID][subscriptionID]; exists {
		return errors.New("subscription already exists")
	}

	s.subscriptions[afID][subscriptionID] = subsc
	return nil
}

// GetSubscription retrieves a subscription
func (s *InMemoryStore) GetSubscription(afID, subscriptionID string) (*models.AnalyticsExposureSubsc, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if afSubs, exists := s.subscriptions[afID]; exists {
		if subsc, exists := afSubs[subscriptionID]; exists {
			return subsc, nil
		}
	}

	return nil, errors.New("subscription not found")
}

// GetAllSubscriptions retrieves all subscriptions for an AF
func (s *InMemoryStore) GetAllSubscriptions(afID string) ([]*models.AnalyticsExposureSubsc, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if afSubs, exists := s.subscriptions[afID]; exists {
		subs := make([]*models.AnalyticsExposureSubsc, 0, len(afSubs))
		for _, subsc := range afSubs {
			subs = append(subs, subsc)
		}
		return subs, nil
	}

	return []*models.AnalyticsExposureSubsc{}, nil
}

// UpdateSubscription updates an existing subscription
func (s *InMemoryStore) UpdateSubscription(afID, subscriptionID string, subsc *models.AnalyticsExposureSubsc) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if afSubs, exists := s.subscriptions[afID]; exists {
		if _, exists := afSubs[subscriptionID]; exists {
			afSubs[subscriptionID] = subsc
			return nil
		}
	}

	return errors.New("subscription not found")
}

// DeleteSubscription deletes a subscription
func (s *InMemoryStore) DeleteSubscription(afID, subscriptionID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if afSubs, exists := s.subscriptions[afID]; exists {
		if _, exists := afSubs[subscriptionID]; exists {
			delete(afSubs, subscriptionID)
			return nil
		}
	}

	return errors.New("subscription not found")
}

// GetAnalyticsData retrieves analytics data
func (s *InMemoryStore) GetAnalyticsData(afID string, req *models.AnalyticsRequest) (*models.AnalyticsData, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Return mock data based on the request
	data := &models.AnalyticsData{
		SuppFeat: &req.SuppFeat,
	}

	// Populate mock data based on analytics event type
	switch req.AnalyticsEvent {
	case models.UE_MOBILITY:
		data.UEMobilityInfos = []models.UEMobilityInfo{
			{
				Duration:     3600,
				LocationInfo: []string{"area1", "area2"},
			},
		}
	case models.NETWORK_PERFORMANCE:
		data.NetworkPerfInfos = []models.NetworkPerfInfo{
			{
				LocationArea:    "area1",
				NetworkPerfType: "THROUGHPUT",
				RelativeRatio:   0.95,
			},
		}
	}

	return data, nil
}
