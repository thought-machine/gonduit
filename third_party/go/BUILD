package(default_visibility = ["PUBLIC"])

go_toolchain(
    name = "toolchain",
    version = "1.16.2",
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
    version = "d8f796af33cc11cb798c1aaeb27a4ebc5099927d",
    test_only = True,
)

go_module(
    name = "difflib",
    install = ["difflib"],
    module = "github.com/pmezard/go-difflib",
    version = "5d4384ee4fb2527b0a1256a821ebfc92f91efefc",
    test_only = True,
)

go_module(
    name = "testify",
    install = [
        "assert",
        "require",
    ],
    module = "github.com/stretchr/testify",
    version = "v1.3.0",
    deps = [
        ":difflib",
        ":spew",
    ],
    test_only = True,
)

go_module(
    name = "ginkgo",
    module = "github.com/onsi/ginkgo",
    version = "v1.15.1",
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
    test_only = True,
    deps = [
        ":tail",
    ],
)

go_module(
    name = "gomega",
    module = "github.com/onsi/gomega",
    version = "v1.11.0",
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
    test_only = True,
    deps = [
        ":x_net",
        ":yaml",
    ],
)

go_module(
    name = "tail",
    module = "github.com/nxadm/tail",
    version = "v1.4.8",
    install = [
        ".",
        "...",
    ],
    test_only = True,
    deps = [
        ":fsnotify",
        ":tomb",
    ],
)

go_module(
    name = "fsnotify",
    module = "github.com/fsnotify/fsnotify",
    version = "v1.4.9",
    test_only = True,
    deps = [
        ":x_sys",
    ],
)

go_module(
    name = "x_sys",
    module = "golang.org/x/sys",
    install = [
        "internal/unsafeheader",
        "unix",
    ],
    version = "0b90fd5c4c48ef7ae7ef6c393cb118d2e7969999",
)

go_module(
    name = "tomb",
    module = "gopkg.in/tomb.v1",
    version = "dd632973f1e7218eb1089048e0798ec9ae7dceb8",
)

go_module(
    name = "x_net",
    module = "golang.org/x/net",
    version = "d523dce5a7f4b994f7ed0531dbe44cd8fd803e26",
    install = [
        "html",
        "html/atom",
        "html/charset",
    ],
    deps = [
        ":x_text",
    ],
)

go_module(
    name = "x_text",
    module = "golang.org/x/text",
    version = "e3aa4adf54f644ca0cb35f1f1fb19b239c40ef04",
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
)

go_module(
    name = "yaml",
    module = "gopkg.in/yaml.v2",
    version = "v2.4.0",
)

go_module(
    name = "gin",
    module = "github.com/gin-gonic/gin",
    version = "v1.6.3",
    install = [
        ".",
        "internal/bytesconv",
        "binding",
        "render",
        "internal/json",
    ],
    test_only = True,
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
    version = "v10.4.1",
    test_only = True,
    deps = [
        ":universal-translator",
        ":urn",
        ":x_crypto",
    ],
)

go_module(
    name = "universal-translator",
    module = "github.com/go-playground/universal-translator",
    version = "v0.17.0",
    test_only = True,
    deps = [
        ":locales",
    ],
)

go_module(
    name = "locales",
    module = "github.com/go-playground/locales",
    install = [
        ".",
        "currency",
    ],
    version = "v0.13.0",
    test_only = True,
)

go_module(
    name = "x_crypto",
    module = "golang.org/x/crypto",
    version = "e6e6c4f2bb5b5887c7f7dd52f01ea7b2fbeb297d",
    install = [
        "sha3",
    ],
)

go_module(
    name = "urn",
    module = "github.com/leodido/go-urn",
    version = "v1.2.1",
    test_only = True,
)

go_module(
    name = "go-codec",
    module = "github.com/ugorji/go/codec",
    version = "v1.2.4",
)

go_module(
    name = "protobuf",
    module = "github.com/golang/protobuf",
    version = "v1.3.3",
    install = [
        "proto",
    ],
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
