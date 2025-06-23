package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func main() {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

	e.GET("/", home)
	e.GET("/:title", title)
	e.GET("/:title/:chapter", chapter)
	e.Static("/assets/*", "views/assets")
	e.Static("/comics/*", "views/comics")
	e.Static("/images/*", "views/images")

	if err := e.Start(":100"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, fmt.Sprintf("%s.html", name), data)
}

func home(c echo.Context) error {
	// Comics Directory
	itemDirs, err := os.ReadDir("./views/comics")
	if err != nil {
		return c.Render(http.StatusBadRequest, "oops", nil)
	}

	sort.Slice(itemDirs, func(i, j int) bool {
		infoI, _ := itemDirs[i].Info()
		infoJ, _ := itemDirs[j].Info()
		return infoI.ModTime().Unix() < infoJ.ModTime().Unix()
	})

	var comics = make([]map[string]interface{}, 0)
	for _, dir := range itemDirs {
		if dir.IsDir() {
			comics = append(comics, map[string]interface{}{
				"icon": fmt.Sprintf("comics/%s/icon.png", url.PathEscape(dir.Name())),
				"name": dir.Name(),
				"url":  fmt.Sprintf("/%s", url.PathEscape(dir.Name())),
			})
		}
	}

	// Images Directory
	imageDirs, err := os.ReadDir("./views/images")
	if err != nil {
		return c.Render(http.StatusBadRequest, "oops", nil)
	}

	sort.Slice(imageDirs, func(i, j int) bool {
		infoI, _ := imageDirs[i].Info()
		infoJ, _ := imageDirs[j].Info()
		return infoI.ModTime().Unix() < infoJ.ModTime().Unix()
	})

	var images = make([]map[string]interface{}, 0)
	for _, dir := range imageDirs {
		if dir.IsDir() {
			// Get the first image file in the folder to use as icon
			imageFiles, err := os.ReadDir(fmt.Sprintf("./views/images/%s", dir.Name()))
			if err != nil {
				continue
			}

			var icon = ""
			for _, file := range imageFiles {
				if !file.IsDir() {
					// Use the first file as icon
					icon = fmt.Sprintf("images/%s/%s", url.PathEscape(dir.Name()), url.PathEscape(file.Name()))
					break
				}
			}

			images = append(images, map[string]interface{}{
				"icon": icon,
				"name": dir.Name(),
				"url":  fmt.Sprintf("/%s", url.PathEscape(dir.Name())),
			})
		}
	}

	return c.Render(http.StatusOK, "home", map[string]interface{}{
		"comics": comics,
		"images": images,
	})
}

func title(c echo.Context) error {
	mode := "comics"
	title := c.Param("title")
	itemDirs, err := os.ReadDir(fmt.Sprintf("./views/comics/%s", title))
	if err != nil {
		mode = "images"
		itemDirs, err = os.ReadDir(fmt.Sprintf("./views/images/%s", title))
		if err != nil {
			return c.Render(http.StatusBadRequest, "oops", nil)
		}
	}

	sort.Slice(itemDirs, func(i, j int) bool {
		infoI, _ := itemDirs[i].Info()
		infoJ, _ := itemDirs[j].Info()
		return infoI.ModTime().Unix() < infoJ.ModTime().Unix()
	})

	var icon = ""
	var items = make([]map[string]interface{}, 0)
	for _, dir := range itemDirs {
		if mode == "comics" {
			if dir.IsDir() {
				items = append(items, map[string]interface{}{
					"name": dir.Name(),
					"url":  fmt.Sprintf("/%s/%s", url.PathEscape(title), url.PathEscape(dir.Name())),
				})
			} else {
				icon = fmt.Sprintf("comics/%s/%s", url.PathEscape(title), url.PathEscape(dir.Name()))
			}
		} else {
			items = append(items, map[string]interface{}{
				"name": dir.Name(),
				"url":  fmt.Sprintf("images/%s/%s", url.PathEscape(title), url.PathEscape(dir.Name())),
			})
		}
	}

	return c.Render(http.StatusOK, "title", map[string]interface{}{
		"title": title,
		"icon":  icon,
		"back":  "/",
		"mode":  mode,
		"items": items,
	})
}

func chapter(c echo.Context) error {
	title := c.Param("title")
	chapter := c.Param("chapter")
	itemDirs, err := os.ReadDir(fmt.Sprintf("./views/comics/%s/%s", title, chapter))
	if err != nil {
		return c.Render(http.StatusBadRequest, "oops", nil)
	}

	sort.Slice(itemDirs, func(i, j int) bool {
		infoI, _ := itemDirs[i].Info()
		infoJ, _ := itemDirs[j].Info()
		return infoI.ModTime().Unix() < infoJ.ModTime().Unix()
	})

	var comics = make([]map[string]interface{}, 0)
	for _, dir := range itemDirs {
		if !dir.IsDir() {
			comics = append(comics, map[string]interface{}{
				"url": fmt.Sprintf("/comics/%s/%s/%s", url.PathEscape(title), url.PathEscape(chapter), url.PathEscape(dir.Name())),
			})
		}
	}

	return c.Render(http.StatusOK, "chapter", map[string]interface{}{
		"back":   fmt.Sprintf("/%s", url.PathEscape(title)),
		"comics": comics,
	})
}
