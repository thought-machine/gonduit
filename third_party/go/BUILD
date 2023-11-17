subinclude("///go//build_defs:go")

package(default_visibility = ["PUBLIC"])

go_toolchain(
    name = "toolchain",
    version = "1.21.3",
)

go_module(
    name = "typed",
    module = "github.com/karlseguin/typed",
    version = "v1.1.7",
)

go_module(
    name = "spew",
    install = ["spew"],
    module = "github.com/davecgh/go-spew",
    test_only = True,
    version = "d8f796af33cc11cb798c1aaeb27a4ebc5099927d",
)

go_module(
    name = "difflib",
    install = ["difflib"],
    module = "github.com/pmezard/go-difflib",
    test_only = True,
    version = "5d4384ee4fb2527b0a1256a821ebfc92f91efefc",
)

go_module(
    name = "testify",
    install = [
        "assert",
        "require",
    ],
    module = "github.com/stretchr/testify",
    test_only = True,
    version = "v1.3.0",
    deps = [
        ":difflib",
        ":spew",
    ],
)

go_module(
    name = "ginkgo",
    install = [
        ".",
        "config",
        "internal/codelocation",
        "internal/containernode",
        "internal/leafnodes",
        "internal/spec",
        "internal/specrunner",
        "internal/failer",
        "internal/suite",
        "internal/spec_iterator",
        "internal/global",
        "internal/remote",
        "internal/testingtproxy",
        "internal/writer",
        "reporters",
        "reporters/stenographer",
        "reporters/stenographer/support/go-colorable",
        "types",
    ],
    module = "github.com/onsi/ginkgo",
    test_only = True,
    version = "v1.15.1",
    deps = [
        ":tail",
    ],
)

go_module(
    name = "gomega",
    install = [
        ".",
        "format",
        "internal/assertion",
        "internal/asyncassertion",
        "internal/oraclematcher",
        "internal/testingtsupport",
        "matchers",
        "matchers/support/goraph/bipartitegraph",
        "matchers/support/goraph/node",
        "matchers/support/goraph/edge",
        "matchers/support/goraph/util",
        "types",
    ],
    module = "github.com/onsi/gomega",
    test_only = True,
    version = "v1.11.0",
    deps = [
        ":x_net",
        ":yaml",
    ],
)

go_module(
    name = "tail",
    install = [
        ".",
        "...",
    ],
    module = "github.com/nxadm/tail",
    test_only = True,
    version = "v1.4.8",
    deps = [
        ":fsnotify",
        ":tomb",
    ],
)

go_module(
    name = "fsnotify",
    module = "github.com/fsnotify/fsnotify",
    test_only = True,
    version = "v1.4.9",
    deps = [
        ":x_sys",
    ],
)

go_module(
    name = "x_sys",
    install = [
        "internal/unsafeheader",
        "unix",
    ],
    module = "golang.org/x/sys",
    version = "0b90fd5c4c48ef7ae7ef6c393cb118d2e7969999",
)

go_module(
    name = "tomb",
    module = "gopkg.in/tomb.v1",
    version = "dd632973f1e7218eb1089048e0798ec9ae7dceb8",
)

go_module(
    name = "x_net",
    install = [
        "html",
        "html/atom",
        "html/charset",
    ],
    module = "golang.org/x/net",
    version = "d523dce5a7f4b994f7ed0531dbe44cd8fd803e26",
    deps = [
        ":x_text",
    ],
)

go_module(
    name = "x_text",
    install = [
        "encoding",
        "encoding/charmap",
        "encoding/htmlindex",
        "transform",
        "encoding/internal/identifier",
        "encoding/internal",
        "language",
        "encoding/japanese",
        "encoding/korean",
        "encoding/simplifiedchinese",
        "encoding/traditionalchinese",
        "encoding/unicode",
        "internal/language",
        "internal/language/compact",
        "internal/tag",
        "internal/utf8internal",
        "runes",
    ],
    module = "golang.org/x/text",
    version = "e3aa4adf54f644ca0cb35f1f1fb19b239c40ef04",
)

go_module(
    name = "yaml",
    module = "gopkg.in/yaml.v2",
    version = "v2.4.0",
)

go_module(
    name = "gin",
    install = [
        ".",
        "internal/bytesconv",
        "binding",
        "render",
        "internal/json",
    ],
    module = "github.com/gin-gonic/gin",
    test_only = True,
    version = "v1.6.3",
    deps = [
        ":go-codec",
        ":isatty",
        ":protobuf",
        ":sse",
        ":validator",
        ":yaml",
    ],
)

go_module(
    name = "validator",
    module = "github.com/go-playground/validator/v10",
    test_only = True,
    version = "v10.4.1",
    deps = [
        ":universal-translator",
        ":urn",
        ":x_crypto",
    ],
)

go_module(
    name = "universal-translator",
    module = "github.com/go-playground/universal-translator",
    test_only = True,
    version = "v0.17.0",
    deps = [
        ":locales",
    ],
)

go_module(
    name = "locales",
    install = [
        ".",
        "currency",
    ],
    module = "github.com/go-playground/locales",
    test_only = True,
    version = "v0.13.0",
)

go_module(
    name = "x_crypto",
    install = [
        "sha3",
    ],
    module = "golang.org/x/crypto",
    version = "e6e6c4f2bb5b5887c7f7dd52f01ea7b2fbeb297d",
)

go_module(
    name = "urn",
    module = "github.com/leodido/go-urn",
    test_only = True,
    version = "v1.2.1",
)

go_module(
    name = "go-codec",
    module = "github.com/ugorji/go/codec",
    version = "v1.2.4",
)

go_module(
    name = "protobuf",
    install = [
        "proto",
    ],
    module = "github.com/golang/protobuf",
    version = "v1.3.3",
)

go_module(
    name = "sse",
    module = "github.com/gin-contrib/sse",
    version = "v0.1.0",
)

go_module(
    name = "isatty",
    module = "github.com/mattn/go-isatty",
    version = "v0.0.12",
    deps = [
        ":x_sys",
    ],
)
