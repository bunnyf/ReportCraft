package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"

	"github.com/genrep/internal/core"
)

// PluginManager handles loading and managing plugins
type PluginManager struct {
	pluginDir  string
	plugins    map[string]core.Plugin
}

// NewPluginManager creates a new plugin manager
func NewPluginManager(pluginDir string) *PluginManager {
	return &PluginManager{
		pluginDir: pluginDir,
		plugins:   make(map[string]core.Plugin),
	}
}

// LoadPlugins loads all plugins from the plugin directory
func (pm *PluginManager) LoadPlugins() error {
	// Create the plugin directory if it doesn't exist
	if _, err := os.Stat(pm.pluginDir); os.IsNotExist(err) {
		if err := os.MkdirAll(pm.pluginDir, 0755); err != nil {
			return fmt.Errorf("failed to create plugin directory: %w", err)
		}
	}

	// Walk the plugin directory and load all .so files
	return filepath.Walk(pm.pluginDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and non-.so files
		if info.IsDir() || filepath.Ext(path) != ".so" {
			return nil
		}

		// Load the plugin
		return pm.loadPlugin(path)
	})
}

// loadPlugin loads a single plugin from the specified path
func (pm *PluginManager) loadPlugin(path string) error {
	// Open the plugin
	p, err := plugin.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open plugin %s: %w", path, err)
	}

	// Look up the Plugin symbol
	pluginSymbol, err := p.Lookup("Plugin")
	if err != nil {
		return fmt.Errorf("plugin %s does not export 'Plugin' symbol: %w", path, err)
	}

	// Convert the symbol to a Plugin instance
	plugin, ok := pluginSymbol.(core.Plugin)
	if !ok {
		return fmt.Errorf("plugin %s does not implement the Plugin interface", path)
	}

	// Register the plugin
	pm.plugins[plugin.Name()] = plugin

	return nil
}

// GetPluginsByType returns all plugins of the specified type
func (pm *PluginManager) GetPluginsByType(pluginType string) []core.Plugin {
	var plugins []core.Plugin
	for _, p := range pm.plugins {
		if p.Type() == pluginType {
			plugins = append(plugins, p)
		}
	}
	return plugins
}

// GetPlugin returns the plugin with the specified name
func (pm *PluginManager) GetPlugin(name string) (core.Plugin, bool) {
	plugin, ok := pm.plugins[name]
	return plugin, ok
}

// RegisterPlugins registers all plugins with the report engine
func (pm *PluginManager) RegisterPlugins(engine *core.ReportEngine) error {
	// Register report generators
	for _, p := range pm.GetPluginsByType("report") {
		generator, ok := p.Instance().(core.ReportGenerator)
		if !ok {
			return fmt.Errorf("plugin %s does not implement ReportGenerator interface", p.Name())
		}
		engine.RegisterReportGenerator(p.Name(), generator)
	}

	// Register data sources
	for _, p := range pm.GetPluginsByType("dataSource") {
		factory, ok := p.Instance().(func() core.DataSource)
		if !ok {
			return fmt.Errorf("plugin %s does not provide a valid DataSource factory", p.Name())
		}
		engine.RegisterDataSource(p.Name(), factory)
	}

	// Register output formatters
	for _, p := range pm.GetPluginsByType("formatter") {
		formatter, ok := p.Instance().(core.OutputFormatter)
		if !ok {
			return fmt.Errorf("plugin %s does not implement OutputFormatter interface", p.Name())
		}
		engine.RegisterOutputFormatter(p.Name(), formatter)
	}

	return nil
}
