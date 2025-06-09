package wilayah

import (
	"sync"
	"time"
)

type RegionCache struct {
	mu          sync.RWMutex
	provinces   map[string]string   // province name -> province code
	regencies   map[string][]string // province code -> regency names
	lastUpdated time.Time
	ttl         time.Duration
}

func NewRegionCache(ttl time.Duration) *RegionCache {
	return &RegionCache{
		provinces:   make(map[string]string),
		regencies:   make(map[string][]string),
		ttl:         ttl,
		lastUpdated: time.Time{},
	}
}

func (rc *RegionCache) IsExpired() bool {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	return time.Since(rc.lastUpdated) > rc.ttl
}

func (rc *RegionCache) GetProvinceCode(provinceName string) (string, bool) {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	code, exists := rc.provinces[provinceName]
	return code, exists
}

func (rc *RegionCache) GetRegencies(provinceCode string) ([]string, bool) {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	regencies, exists := rc.regencies[provinceCode]
	return regencies, exists
}

func (rc *RegionCache) SetProvinceMapping(provinceName, provinceCode string) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	rc.provinces[provinceName] = provinceCode
	rc.lastUpdated = time.Now()
}

func (rc *RegionCache) SetRegencies(provinceCode string, regencies []string) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	rc.regencies[provinceCode] = regencies
	rc.lastUpdated = time.Now()
}

func (rc *RegionCache) Clear() {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	rc.provinces = make(map[string]string)
	rc.regencies = make(map[string][]string)
	rc.lastUpdated = time.Time{}
}

func (rc *RegionCache) GetStats() map[string]interface{} {
	rc.mu.RLock()
	defer rc.mu.RUnlock()

	return map[string]interface{}{
		"provinces_count": len(rc.provinces),
		"regencies_count": len(rc.regencies),
		"last_updated":    rc.lastUpdated,
		"is_expired":      time.Since(rc.lastUpdated) > rc.ttl,
		"ttl_hours":       rc.ttl.Hours(),
	}
}
