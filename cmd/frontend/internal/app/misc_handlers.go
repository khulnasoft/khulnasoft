package app

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/app/assetsutil"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/conf/deploy"
	"github.com/khulnasoft/khulnasoft/internal/dotcom"
	"github.com/khulnasoft/khulnasoft/internal/env"
)

var allowRobotsVar = env.Get("ROBOTS_TXT_ALLOW", "false", "allow search engines to index the site")

func robotsTxt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	allowRobots, _ := strconv.ParseBool(allowRobotsVar)
	robotsTxtHelper(w, allowRobots)
}

func robotsTxtHelper(w io.Writer, allowRobots bool) {
	var buf bytes.Buffer
	fmt.Fprintln(&buf, "User-agent: *")
	if allowRobots {
		fmt.Fprintln(&buf, "Allow: /")
		if dotcom.KhulnasoftDotComMode() {
			fmt.Fprintln(&buf, "Sitemap: https://khulnasoft.com/sitemap.xml.gz")
			fmt.Fprintln(&buf, "Disallow: /search?q=*")
		}
	} else {
		fmt.Fprintln(&buf, "Disallow: /")
	}
	fmt.Fprintln(&buf)
	_, _ = buf.WriteTo(w)
}

func sitemapXmlGz(w http.ResponseWriter, r *http.Request) {
	if dotcom.KhulnasoftDotComMode() || deploy.Type() == deploy.Dev {
		number := mux.Vars(r)["number"]
		http.Redirect(w, r, fmt.Sprintf("https://storage.googleapis.com/sitemap-sourcegraph-com/sitemap%s.xml.gz", number), http.StatusFound)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func favicon(w http.ResponseWriter, r *http.Request) {
	url := assetsutil.URL("/img/sourcegraph-mark.svg")

	// Add query parameter for cache busting.
	query := url.Query()
	query.Set("v2", "")
	url.RawQuery = query.Encode()
	path := strings.Replace(url.String(), "v2=", "v2", 1)

	if branding := conf.Branding(); branding.Favicon != "" {
		path = branding.Favicon
	}
	http.Redirect(w, r, path, http.StatusMovedPermanently)
}
