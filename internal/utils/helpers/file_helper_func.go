package helpers

import (
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
)

// BaseDir Return project base directory
func BaseDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Panicf("Failed to get current directory: %v", err)
	}
	return dir
}

// WebPath returns web base directory path
func WebPath() string {
	return filepath.Join(BaseDir(), "web")
}

// TemplatesPath  returns template base directory path
func TemplatesPath() string {
	return filepath.Join(WebPath(), "template")
}

// ViewsPath  returns views base directory path
func ViewsPath() string {
	return filepath.Join(TemplatesPath(), "views")
}

// RenderView render a view to every element satisfying io.Writer
func RenderView(viewName string, data interface{}, output io.Writer) {

	viewTemplateFile := filepath.Join(ViewsPath(), "home.gohtml")

	if FileDoesNotExists(viewTemplateFile) {
		log.Panicf("%s view template file not exists", viewName)
	}

	cssTemplatePath := filepath.Join(TemplatesPath(), "css.gohtml")
	jsTemplateFile := filepath.Join(TemplatesPath(), "js.gohtml")
	menuTemplateFile := filepath.Join(TemplatesPath(), "menu.gohtml")
	layoutTemplateFile := filepath.Join(TemplatesPath(), "layout.gohtml")
	viewTemplate, err := template.ParseFiles(layoutTemplateFile, cssTemplatePath, jsTemplateFile, menuTemplateFile, viewTemplateFile)

	PanicOnError(err)

	err = viewTemplate.Execute(os.Stdout, nil)

	PanicOnError(err)
}

// RenderViewToStdout render a view to os.Stdout
func RenderViewToStdout(viewName string, data interface{}) {
	RenderView(viewName, data, os.Stdout)
}

// FileDoesNotExists  check if file does not exist
func FileDoesNotExists(filename string) bool {
	_, err := os.Stat(filename)
	return os.IsNotExist(err)
}

// FileExists  check if file  exist
func FileExists(filename string) bool {
	return !FileDoesNotExists(filename)
}
