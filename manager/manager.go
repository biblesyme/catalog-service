package manager

import (
	"github.com/jinzhu/gorm"
	"github.com/rancher/catalog-service/model"
)

// TODO: move elsewhere
type CatalogConfig struct {
	URL    string
	Branch string
}

type Manager struct {
	cacheRoot string
	config    map[string]CatalogConfig
	db        *gorm.DB
}

func NewManager(cacheRoot string, config map[string]CatalogConfig, db *gorm.DB) *Manager {
	return &Manager{
		cacheRoot: cacheRoot,
		config:    config,
		db:        db,
	}
}

func (m *Manager) RefreshAll() error {
	catalogs, err := m.lookupCatalogs("")
	if err != nil {
		return err
	}
	for _, catalog := range catalogs {
		if err := m.refreshCatalog(catalog); err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) Refresh(environmentId string) error {
	catalogs, err := m.lookupCatalogs(environmentId)
	if err != nil {
		return err
	}
	for _, catalog := range catalogs {
		if err := m.refreshCatalog(catalog); err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) refreshCatalog(catalog model.Catalog) error {
	repoPath, commit, err := m.prepareRepoPath(catalog)
	if err != nil {
		return err
	}

	// Catalog is already up to date
	if commit == catalog.Commit {
		return nil
	}

	templates, versions, err := traverseFiles(repoPath)
	if err != nil {
		return err
	}

	return m.updateDb(catalog, templates, versions, commit)
}

// TODO: move elsewhere
type TemplateConfig struct {
	Name           string            `yaml:"name"`
	Category       string            `yaml:"category"`
	Description    string            `yaml:"description"`
	Version        string            `yaml:"version"`
	Maintainer     string            `yaml:"maintainer"`
	License        string            `yaml:"license"`
	ProjectURL     string            `yaml:"projectURL"`
	IsSystem       string            `yaml:"isSystem"`
	DefaultVersion string            `yaml:"version"`
	Labels         map[string]string `yaml:"version"`
}
