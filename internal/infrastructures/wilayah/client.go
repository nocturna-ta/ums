package wilayah

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/tracing"
	"net/http"
	"strings"
	"time"
)

type WilayahAPIClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewWilayahAPIClient() *WilayahAPIClient {
	return &WilayahAPIClient{
		baseURL: "https://wilayah.id/api",
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (w *WilayahAPIClient) GetAllProvinces(ctx context.Context) ([]ProvinceData, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WilayahAPIClient.GetAllProvinces")
	defer span.End()

	url := fmt.Sprintf("%s/provinces.json", w.baseURL)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"url":   url,
		}).ErrorWithCtx(ctx, "[WilayahAPIClient.GetAllProvinces] failed to create request")
		return nil, err
	}

	resp, err := w.httpClient.Do(req)
	if err != nil {
		log.WithFields(log.Fields{
			"status_code": resp.StatusCode,
			"url":         url,
		}).ErrorWithCtx(ctx, "[WilayahAPIClient.GetAllProvinces] Request failed")
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"status_code": resp.Status,
			"url":         url,
		}).ErrorWithCtx(ctx, "[WilayahAPIClient.GetAllProvinces] API returned non-200 status")
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var response ProvinceAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[WilayahAPIClient.GetAllProvinces] Failed to decode response")
		return nil, err
	}

	log.WithFields(log.Fields{
		"province_count": len(response.Data),
		"updated_at":     response.Meta.UpdatedAt,
	}).InfoWithCtx(ctx, "[WilayahAPIClient.GetAllProvinces] Successfully fetched provinces")

	return response.Data, nil
}

func (w *WilayahAPIClient) GetRegenciesByProvinceCode(ctx context.Context, provinceCode string) ([]RegencyData, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WilayahAPIClient.GetRegenciesByProvinceCode")
	defer span.End()

	url := fmt.Sprintf("%s/regencies/%s.json", w.baseURL, provinceCode)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"error":        err,
			"url":          url,
			"provinceCode": provinceCode,
		}).ErrorWithCtx(ctx, "[WilayahAPIClient.GetRegenciesByProvinceCode] Failed to create request")
		return nil, err
	}

	resp, err := w.httpClient.Do(req)
	if err != nil {
		log.WithFields(log.Fields{
			"error":        err,
			"url":          url,
			"provinceCode": provinceCode,
		}).ErrorWithCtx(ctx, "[WilayahAPIClient.GetRegenciesByProvinceCode] Request failed")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"status_code":  resp.StatusCode,
			"url":          url,
			"provinceCode": provinceCode,
		}).ErrorWithCtx(ctx, "[WilayahAPIClient.GetRegenciesByProvinceCode] API returned non-200 status")
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var response RegencyAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.WithFields(log.Fields{
			"error":        err,
			"provinceCode": provinceCode,
		}).ErrorWithCtx(ctx, "[WilayahAPIClient.GetRegenciesByProvinceCode] Failed to decode response")
		return nil, err
	}

	log.WithFields(log.Fields{
		"regency_count": len(response.Data),
		"province_code": provinceCode,
		"updated_at":    response.Meta.UpdatedAt,
	}).InfoWithCtx(ctx, "[WilayahAPIClient.GetRegenciesByProvinceCode] Successfully fetched regencies")

	return response.Data, nil
}

func (w *WilayahAPIClient) FindProvinceByName(ctx context.Context, provinceName string) (string, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WilayahAPIClient.FindProvinceByName")
	defer span.End()

	provinces, err := w.GetAllProvinces(ctx)
	if err != nil {
		return "", err
	}

	normalizedSearchName := normalizeRegionName(provinceName)

	for _, province := range provinces {
		normalizedProvinceName := normalizeRegionName(province.Name)
		if normalizedProvinceName == normalizedSearchName {
			log.WithFields(log.Fields{
				"provinceName": provinceName,
				"provinceCode": province.Code,
			}).InfoWithCtx(ctx, "[WilayahAPIClient.FindProvinceByName] Found province")
			return province.Code, nil
		}
	}

	log.WithFields(log.Fields{
		"provinceName": provinceName,
	}).WarnWithCtx(ctx, "[WilayahAPIClient.FindProvinceByName] Province not found")
	return "", fmt.Errorf("province not found: %s", provinceName)
}

func (w *WilayahAPIClient) IsRegencyInProvince(ctx context.Context, provinceName, regencyName string) (bool, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WilayahAPIClient.IsRegencyInProvince")
	defer span.End()

	provinceCode, err := w.FindProvinceByName(ctx, provinceName)
	if err != nil {
		return false, err
	}

	regencies, err := w.GetRegenciesByProvinceCode(ctx, provinceCode)
	if err != nil {
		return false, err
	}

	normalizedSearchName := normalizeRegionName(regencyName)

	for _, regency := range regencies {
		normalizedRegencyName := normalizeRegionName(regency.Name)
		if normalizedRegencyName == normalizedSearchName {
			log.WithFields(log.Fields{
				"provinceName": provinceName,
				"regencyName":  regencyName,
				"found":        true,
			}).InfoWithCtx(ctx, "[WilayahAPIClient.IsRegencyInProvince] Regency found in province")
			return true, nil
		}
	}

	log.WithFields(log.Fields{
		"provinceName": provinceName,
		"regencyName":  regencyName,
		"found":        false,
	}).InfoWithCtx(ctx, "[WilayahAPIClient.IsRegencyInProvince] Regency not found in province")
	return false, nil
}

func normalizeRegionName(name string) string {
	normalized := strings.TrimSpace(name)
	normalized = strings.ToLower(normalized)

	prefixes := []string{"kota ", "kabupaten ", "kab. ", "kab ", "kot. ", "kot ", "provinsi "}
	for _, prefix := range prefixes {
		if strings.HasPrefix(normalized, prefix) {
			normalized = strings.TrimPrefix(normalized, prefix)
			break
		}
	}

	return normalized
}

func (w *WilayahAPIClient) GetRegenciesByProvinceName(ctx context.Context, provinceName string) ([]RegencyData, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WilayahAPIClient.GetRegenciesByProvinceName")
	defer span.End()

	provinceCode, err := w.FindProvinceByName(ctx, provinceName)
	if err != nil {
		return nil, err
	}

	return w.GetRegenciesByProvinceCode(ctx, provinceCode)
}
