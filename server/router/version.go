// Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
// SPDX-License-Identifier: Apache-2.0

package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khulnasoft/khulnasoft/server/config"
)

type versionInfo struct {
	Revision string `json:"revision"`
}

func (rt *router) getVersion(c *gin.Context) {
	// this endpoint is most likely to be consumed by humans, so
	// we pretty print the output
	c.IndentedJSON(http.StatusOK, versionInfo{
		Revision: config.Revision,
	})
}
