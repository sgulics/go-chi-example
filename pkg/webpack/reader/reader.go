package reader

import (
	"errors"

	"github.com/sgulics/go-chi-example/pkg/webpack/reader/manifest"
	"github.com/sgulics/go-chi-example/pkg/webpack/reader/stats"
)

// Read assets asset manifest
func Read(plugin, host, fsPath, webPath string, isDev bool) (map[string][]string, error) {
	//log.Println("read", plugin, isDev)
	if plugin == "stats" {
		return stats.Read(isDev, host, fsPath, webPath)
	} else if plugin == "manifest" {
		return manifest.Read(fsPath)
	} else {
		return map[string][]string{}, errors.New("go-webpack: bad plugin type")
	}

}
