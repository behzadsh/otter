package configs

import (
	"errors"
	"net/url"
	"strings"
)

var appVersion string

func init() {
	if appVersion == "" {
		appVersion = "dev"
	}
}

func GetAppVersion() string {
	return appVersion
}

func DevMode() bool {
	return strings.HasSuffix(appVersion, "beta") || strings.HasSuffix(appVersion, "dev")
}

type Config struct {
	BaseURL       string `yaml:"base-url"`
	TestCasesPath string `yaml:"test-cases-path"`

	parsedURL *url.URL
}

func NewConfig(baseURL, casesPath string) (*Config, error) {
	parsedURL, err := parseURL(baseURL)
	if err != nil {
		return nil, err
	}

	return &Config{
		BaseURL:       baseURL,
		TestCasesPath: casesPath,
		parsedURL:     parsedURL,
	}, nil
}

func parseURL(rawURL string) (*url.URL, error) {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return nil, err
	}

	if u.Scheme == "" || u.Host == "" {
		return nil, &url.Error{Op: "parse", URL: rawURL, Err: errors.New("invalid url")}
	}

	return u, nil
}
