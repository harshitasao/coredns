// generated by directives_generate.go; DO NOT EDIT
package dnsserver

// Directives are registered in the order they should be
// executed.
//
// Ordering is VERY important. Every middleware will
// feel the effects of all other middleware below
// (after) them during a request, but they must not
// care what middleware above them are doing.

var directives = []string{
	"root",
	"bind",
	"trace",
	"health",
	"pprof",
	"prometheus",
	"errors",
	"log",
	"chaos",
	"cache",
	"rewrite",
	"loadbalance",
	"dnssec",
	"file",
	"auto",
	"secondary",
	"etcd",
	"kubernetes",
	"reverse",
	"proxy",
	"httpprox",
	"whoami",
	"erratic",
}
