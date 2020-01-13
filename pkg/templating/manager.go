package templating

import (
	"errors"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// This template manager is heavily influenced by the following
// * https://gist.github.com/logrusorgru/abd846adb521a6fb39c7405f32fec0cf
// * https://github.com/asit-dhal/golang-template-layout/blob/master/src/templmanager/templatemanager.go

type TemplateManager struct {
	dir string           // root directory
	ext string           // extension
	devel bool             // reload every time
	funcs template.FuncMap // functions
	loadedAt time.Time        // loaded at (last loading time)
	layoutDir string // template layout directory
	templates map[string]*template.Template
}

// NewTemplateManager creates new TemplateManager and loads templates_oild. The dir argument is
// directory to load templates_oild from. The ext argument is extension of
// tempaltes. The devel (if true) turns the TemplateManager to reload templates_oild
// every Render if there is a change in the dir.
func NewTemplateManager(dir, layoutDir string, ext string, devel bool) (tmpl *TemplateManager, err error) {

	// get absolute path
	if dir, err = filepath.Abs(dir); err != nil {
		return
	}

	// get absolute path
	if layoutDir, err = filepath.Abs(layoutDir); err != nil {
		return
	}

	tmpl = new(TemplateManager)
	tmpl.dir = dir
	tmpl.ext = ext
	tmpl.devel = devel
	tmpl.layoutDir = layoutDir
    tmpl.templates = make(map[string]*template.Template)

	if err = tmpl.Load(); err != nil {
		tmpl = nil // drop for GC
	}

	return
}

// Dir returns absolute path to directory with views
func (t *TemplateManager) Dir() string {
	return t.dir
}

// Ext returns extension of views
func (t *TemplateManager) Ext() string {
	return t.ext
}

// Devel returns development pin
func (t *TemplateManager) Devel() bool {
	return t.devel
}

func (t *TemplateManager) Templates() []*template.Template {
	var temps []*template.Template
	for _, v := range t.templates {
		temps  = append(temps, v)
	}
	return temps
}

// Funcs sets template functions
func (t *TemplateManager) Funcs(funcMap template.FuncMap) {
	// TODO Implement this next
	//t.Template = t.Template.Funcs(funcMap)
	//t.funcs = funcMap
}

// Load or reload templates_oild
func (t *TemplateManager) Load() (err error) {

	// time point
	t.loadedAt = time.Now()

	layoutFiles, err := filepath.Glob(t.layoutDir + "/*.gohtml")

	if err != nil {
		return err
	}

	var walkFunc = func(path string, info os.FileInfo, err error) (_ error) {

		// handle walking error if any
		if err != nil {
			return err
		}

		// skip all except regular files
		// TODO (kostyarin): follow symlinks
		if !info.Mode().IsRegular() {
			return
		}

		// filter by extension
		if filepath.Ext(path) != t.ext {
			return
		}

		// get relative path
		var rel string
		if rel, err = filepath.Rel(t.dir, path); err != nil {
			return err
		}

		// Ignore files in the layout directory
		if filepath.Dir(path) == t.layoutDir {
			return
		}

		// name of a template is its relative path
		// without extension
		rel = strings.TrimSuffix(rel, t.ext)

		var (
			nt = template.New(rel)
			b  []byte
		)

		if b, err = ioutil.ReadFile(path); err != nil {
			return err
		}
		tmpl, err := nt.ParseFiles(layoutFiles...)
		if err != nil {
			return err
		}

		tmpl, err = nt.Parse(string(b))
		if err != nil {
			return err
		}

		t.templates[tmpl.Name()] = tmpl

		return err
	}

	if err = filepath.Walk(t.dir, walkFunc); err != nil {
		return
	}

	// necessary for reloading
	// TODO use something like this for the funcMap
	//if t.funcs != nil {
	//	root = root.Funcs(t.funcs)
	//}


	return
}

// IsModified lookups directory for changes to
// reload (or not to reload) templates_oild if development
// pin is true.
func (t *TemplateManager) IsModified() (yep bool, err error) {

	var errStop = errors.New("stop")

	var walkFunc = func(path string, info os.FileInfo, err error) (_ error) {

		// handle walking error if any
		if err != nil {
			return err
		}

		// skip all except regular files
		if !info.Mode().IsRegular() {
			return
		}

		// filter by extension
		if filepath.Ext(path) != t.ext {
			return
		}

		if yep = info.ModTime().After(t.loadedAt); yep == true {
			return errStop
		}

		return
	}

	// clear the errStop
	if err = filepath.Walk(t.dir, walkFunc); err == errStop {
		err = nil
	}

	return
}

func (t *TemplateManager) Render(w io.Writer, name string, data interface{}) (err error) {

	// if development
	if t.devel == true {

		// lookup directory for changes
		var modified bool
		if modified, err = t.IsModified(); err != nil {
			return
		}

		// reload
		if modified == true {
			if err = t.Load(); err != nil {
				return
			}
		}

	}

	tmpl, ok := t.templates[name]
	if !ok{
		return errors.New("template not found")
	}
	return tmpl.Execute(w, data)
}


