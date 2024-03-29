// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package oapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RX32/bNhD+VwhuDxugxHKSdZmKAkt/rDPapsHcAQMCY7hIZ5mNRLLkKbFn6H8fSEqW",
	"bcnLS4MG2FMs3vF4933fHZk1T1WplURJlidrbtMFluB/vjFGGfdDG6XRkEC/nKoM3d+5MiUQT7iQdHrC",
	"I04rjeETczS8jniJ1kLuvRujJSNkzus64ga/VMJgxpPrELPzn22CqZvPmJKLdYXUz0Vk+5k8OxvMREI5",
	"lEbECfKH0xMZb0IcyCwkQ1j6H98bnPOEfzfqoB01uI5cGQ4ZWE6C+ziONzHBGFj50y2mlRG0mrptodYb",
	"BIPmoqLFhie3Jyx3RS+ItDshVepWYOsuJE+apbaUhE/fTKeTj5fdXtDiHboE6ogLOVeBbkmQeuyxBFHw",
	"hN8qAqOOFcEt/IpLKHWBx6kq3bEZ2tQITUK5Ex00pIw7sxApSutJaI7/MPnkGRBU+GzuIc/RsK09d2hs",
	"CDQ+jo9j5600StCCJ/zUL0VcAy08QCPdEJEHqeymUghLzHv4IAbc+iTjCX8vLF0FgwYDJRIay5Pr/Qi/",
	"q3tWglwxTzQjxQxSZSQDYkoiI1Ei+6GEJRvH8Y88CqB/qdCsOswLUQriUdNng31UwlKUVbkrjY2W65nT",
	"ptVK2qCLkzhueULpCwetC5H6Akefrct9vXXeA+q0gf7d0i+Yhhwz5gXK1LwFcoGQebDW/K+jS1wO4H7B",
	"CiFvHVq0QCZxST6WC9KVsY3HfiuGbOZQFfTV6gyjbaDQSuJSY0qYMWx8Im6rsgSzaqTCoCja+glypxTu",
	"P2d1xLWyAxikBoHQbeqJ75U3NfJzMwctvVTZ6mtSGgrtBhqZCuueisb9vENymeu7s2B/XPT/lFDRQhnx",
	"Tzj0G9PezGA/Cran7/XMNWGnigATg4bgPVHUUZhMo7VG+ltk9cERlaOfUOxmxfyNs6uU6ULdXyG9XE2y",
	"hybVpwUykbkWcz3nQoZhZQTeYTuY3Nzs5lLIje/L5L8a85EH0RBPH98FLZ49viQuFbHfVCWfghA3SpvI",
	"uWJzZRgwqzEVc5EeFp3BTBgMV/eg3jYO+1L7ozPsMHwan/SjBJB27oL3KqDSd95YtoXVPCGax4tNRqPt",
	"R0XU3ZCVEd1jZeeC+L+Pp+3X3v542mKzVYlF6x9WQSjtl2erQMI+b2GdtZ77gnntzdONdWgu9Dp5RzJT",
	"pKNX4X3a892yDcumeci+eM4+wPLoIscXR+PnB5TyVFr59T6ifW6iw9fEISLeIn1zFsYnp2c/Pfv5/Jf4",
	"iVPwdgfIgd7wXWbu2ku2MkUzo5LRqFApFAtlKTmPz2PuWq6JsB74J2jnorW8jvrotnk0fm0a9az+NwAA",
	"//+WppodpA8AAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
