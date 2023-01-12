// Code generated by Thrift Compiler (0.17.0). DO NOT EDIT.

package impalaservice

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/n0dev/go-impala/services/status"
	"github.com/n0dev/go-impala/services/beeswax"
	"github.com/n0dev/go-impala/services/cli_service"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = errors.New
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

var _ = status.GoUnusedProtection__
var _ = beeswax.GoUnusedProtection__
var _ = cli_service.GoUnusedProtection__
var DEFAULT_QUERY_OPTIONS map[TImpalaQueryOptions]string

func init() {
DEFAULT_QUERY_OPTIONS = map[TImpalaQueryOptions]string{
  0: "false",
  9: "false",
  3: "0",
  11: "",
  10: "-1",
  2: "false",
  1: "0",
  7: "0",
  6: "0",
  4: "0",
  5: "0",
  8: "0",
}

}

