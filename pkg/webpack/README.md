## go-webpack

[![GoDoc](https://godoc.org/github.com/go-webpack/webpack?status.svg)](https://godoc.org/github.com/go-webpack/webpack)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-webpack/webpack)](https://goreportcard.com/report/github.com/go-webpack/webpack)
[![Maintainability](https://api.codeclimate.com/v1/badges/93067e957c189653e7a3/maintainability)](https://codeclimate.com/github/go-webpack/webpack/maintainability)
[![Build Status](https://travis-ci.com/go-webpack/webpack.svg?branch=master)](https://travis-ci.com/go-webpack/webpack)

This module allows proper integration with webpack, with support for proper assets reloading in development and asset hashes for production caching.

This module is compatible with webpack 1.0 to 4.0

For webpack 4 use this example config file:
https://github.com/go-webpack/examples/blob/master/config/webpack.config.js
(the only thing required for go-webpack is ManifestPlugin - the rest is my personal example config)

## Examples

Are [here](https://github.com/go-webpack/examples)

## Changelog

#### Unreleased

#### Version 2.0.0
- Switch to go 1.11 modules
- Various refactoring
- Change default to stats plugin, to use StatsPlugin set it manually ```webpack.Plugin = "stats"```
- Update examples, add pongo2, gin and gin-render examples

#### Version 1.4.2
- Don't fail if manifest is not present in development mode

#### Version 1.4.1
- less noisy (remove some logs on verbose)

#### Version 1.4.0

- 2018-11-10 Fix globals usage and add some tests #4 thx @toshok
- 2018-11-10 tweak iris example #3 thx @vosmith
- assorted refactoring

#### Version 1.3.0

- 2018-09-12 Add ability to use multiple webpack manifests per app

#### Version 1.2

- 2018-07-17 Don't hardcode manifest path (thx @rodumani)
- 2018-07-17 Deprecate stats plugin default, specify plugin explicitly to continue using it.
- 2018-07-17 Pretty print manifest content

#### Version 1.1

- 2018-02-13 Move examples to [separate repo](https://github.com/go-webpack/examples)
- 2018-02-09 Refactor & cleanup code, add support for ManifestPlugin instead of outdated StatsPlugin (see new examples)
- 2017-08-09 Now reports if you didn't call webpack.Init() to set working mode properly

#### Version 1.0

- 2017-03-07 Initial version / extraction

#### Usage with QOR / Gin
##### main.go
```golang
import (
  ...
  "github.com/go-webpack/webpack"
)
func main() {
  is_dev := flag.Bool("dev", false, "development mode")
  flag.Parse()
  webpack.DevHost = "localhost:3808" // default
  webpack.Plugin = "manifest" // defaults to stats for compatability
  // webpack.IgnoreMissing = true // ignore assets not present in manifest
  webpack.Init(*is_dev)
  ...
}
```

##### controller.go (qor)
```golang
package controllers

import (
  "github.com/qor/render"
  "github.com/gin-gonic/gin"
  "github.com/go-webpack/webpack"
)

var Render *render.Render

func init() {
  Render = render.New()
}

func ViewHelpers() map[string]interface{} {
  return map[string]interface{}{"asset": webpack.AssetHelper}
}

func HomeIndex(ctx *gin.Context) {
  Render.Funcs(ViewHelpers()).Execute(
    "home_index",
    gin.H{},
    ctx.Request,
    ctx.Writer,
  )
}
```

##### alternate controller/route (gin / eztemplate)

```golang
import (
  "github.com/gin-gonic/gin"
  eztemplate "github.com/michelloworld/ez-gin-template"
)
r = gin.Default()
render := eztemplate.New()
render.TemplateFuncMap = template.FuncMap{
  "asset": webpack.AssetHelper,
}
r.HTMLRender = render.Init()
```

##### layouts/application.tmpl

```html
<!doctype html>
<html>
  <head>
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    {{ asset "vendor.css" }}
    {{ asset "application.css" }}
  </head>
  <body>
    <div class="page-wrap">
      {{render .Template}}
    </div>
    {{ asset "vendor.js" }}
    {{ asset "application.js" }}
  </body>
</html>
```

#### Usage with Iris

##### main.go

```golang
import (
  "flag"
  "github.com/go-webpack/webpack"
  iris "gopkg.in/kataras/iris.v6"
  "gopkg.in/kataras/iris.v6/adaptors/httprouter"
  "gopkg.in/kataras/iris.v6/adaptors/view"
)

func main() {
  is_dev := flag.Bool("dev", false, "development mode")
  flag.Parse()
  webpack.Plugin = "manifest"
  webpack.Init(*is_dev)
  view := view.HTML("./templates", ".html")
  view = view.Layout("layout.html")
  view = view.Funcs(map[string]interface{}{"asset": webpack.AssetHelper})
  app := iris.New()
  app.Adapt(view.Reload(*is_dev))

  app.Adapt(httprouter.New())

  ...

}
```

##### templates/layout.html
```html
<!DOCTYPE HTML>
<html lang="en" >
<head>
<meta charset="UTF-8">
<title></title>
{{ asset "vendor.css" }}
{{ asset "application.css" }}
</head>
<body>
{{ yield }}
{{ asset "vendor.js" }}
{{ asset "application.js" }}
```

#### Using multiple webpack manifests per app

Alternative usage if you need multiple webpack manifests (if you have separate js apps with separate webpack configs for separate portions of your app for example)

for qor/render, other would be similar

```
	LandingRender = render.New(&render.Config{
		ViewPaths:     []string{"app/views"},
		DefaultLayout: "landing",
		FuncMapMaker: func(*render.Render, *http.Request, http.ResponseWriter) template.FuncMap {
			return template.FuncMap{
				"asset": webpack.GetAssetHelper(webpack.BasicConfig(
					"localhost:7420", // dev server
					"./public/landing", // file path
					"landing", // web path
				)),
				//"json":  json.Marshal,
			}
		},
	})
```

#### Usage with other frameworks

- Configure webpack to serve manifest.json via ~~StatsPlugin~~ ManifestPlugin
- Call ```webpack.Plugin = "manifest"``` to set go-webpack to use ManifestPlugin, don't call to use old StatsPlugin
- Call ```webpack.Init()``` to set development or production mode.
- Add webpack.AssetHelper to your template functions.
- Call helper function with the name of your asset

Use webpack.config.js (and package.json) from this repo or create your own.

The only thing that must be present in your webpack config is ~~StatsPlugin~~ ManifestPlugin which is required to serve assets the proper way with hashes, etc.

Your compiled assets is expected to be at public/webpack and your webpack-dev-server at http://localhost:3808

When run with -dev flag, webpack asset manifest is loaded from http://localhost:3808/webpack/manifest.json, and updated automatically on every request. When running in production from public/webpack/manifest.json and is persistently cached in memory for performance reasons.

#### Running examples

Examples moved to separate repo [here](https://github.com/go-webpack/examples)

```
cd examples
yarn install # or npm install
# for development mode
yarn start
# or ./node_modules/.bin/webpack-dev-server --config webpack.config.js --hot --inline

# Or for production mode
yarn build
# or ./node_modules/.bin/webpack --config webpack.config.js --bail

# start backend
go get
go run iris/main.go -dev
go run qor/main.go -dev
```

If all is working, you will see a JS alert message.

#### Compiling assets for production

```
NODE_ENV=production ./node_modules/.bin/webpack --config webpack.config.js
```
Don't forget to set go-webpack to production mode (webpack.Init(false))

#### Additional settings

var DevHost = "localhost:3808" - webpack-dev-server host:port
var FsPath = "public/webpack" - filesystem path to public webpack dir
var WebPath = "webpack" - http path to public webpack dir
var Plugin = "stats" - webpack plugin to use, can be stats or manifest
var IgnoreMissing = true - ignore assets missing on manifest or fail on them
var Verbose = true - print error messages to console (even if error is ignored)

#### License

Copyright (c) 2017 glebtv

MIT License


