package server

import (
	"html/template"
	"io/fs"
	"net/http"
	"regexp"
	"strings"
)

func loadAndAddToRoot(
	funcMap template.FuncMap,
	rootTemplate *template.Template,
	embedFS fs.FS,
	pattern string,
) error {
	pattern = strings.ReplaceAll(pattern, ".", "\\.")
	pattern = strings.ReplaceAll(pattern, "*", ".*")

	err := fs.WalkDir(embedFS, ".", func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if matched, _ := regexp.MatchString(pattern, path); !d.IsDir() && matched {
			data, readErr := fs.ReadFile(embedFS, path)
			if readErr != nil {
				return readErr
			}
			t := rootTemplate.New(path).Funcs(funcMap)
			if _, parseErr := t.Parse(string(data)); parseErr != nil {
				return parseErr
			}
		}
		return nil
	})
	return err
}

func (s *Server) loadHTMLFromEmbedFS(embedFS fs.FS, subFSPath string, paths ...string) error {
	subFS, err := fs.Sub(embedFS, subFSPath)
	if err != nil {
		return err
	}

	root := template.New("")

	for _, path := range paths {
		_ = template.Must(
			root,
			loadAndAddToRoot(
				s.router.FuncMap,
				root,
				subFS,
				path),
		)
	}

	for _, t := range root.Templates() {
		s.Log.Info().Str("tmpl", t.Name()).Msg("root template")
	}

	s.router.SetHTMLTemplate(root)

	// prepare subfs
	assetFS, _ := fs.Sub(embedFS, "web/assets")
	s.router.StaticFS("/assets", http.FS(assetFS))

	return nil
}
