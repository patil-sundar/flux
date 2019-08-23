// DO NOT EDIT: This file is autogenerated via the builtin command.

package stdlib

import (
	ast "github.com/influxdata/flux/ast"
	experimental "github.com/influxdata/flux/stdlib/experimental"
	http "github.com/influxdata/flux/stdlib/http"
	monitor "github.com/influxdata/flux/stdlib/influxdata/influxdb/monitor"
	v1 "github.com/influxdata/flux/stdlib/influxdata/influxdb/v1"
	strings "github.com/influxdata/flux/stdlib/strings"
	chronograf "github.com/influxdata/flux/stdlib/testing/chronograf"
	kapacitor "github.com/influxdata/flux/stdlib/testing/kapacitor"
	pandas "github.com/influxdata/flux/stdlib/testing/pandas"
	testdata "github.com/influxdata/flux/stdlib/testing/testdata"
)

var FluxTestPackages = func() []*ast.Package {
	var pkgs []*ast.Package
	pkgs = append(pkgs, experimental.FluxTestPackages...)
	pkgs = append(pkgs, http.FluxTestPackages...)
	pkgs = append(pkgs, monitor.FluxTestPackages...)
	pkgs = append(pkgs, v1.FluxTestPackages...)
	pkgs = append(pkgs, strings.FluxTestPackages...)
	pkgs = append(pkgs, chronograf.FluxTestPackages...)
	pkgs = append(pkgs, kapacitor.FluxTestPackages...)
	pkgs = append(pkgs, pandas.FluxTestPackages...)
	pkgs = append(pkgs, testdata.FluxTestPackages...)
	return pkgs
}()
