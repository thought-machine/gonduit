subinclude("///go//build_defs:go")

package(default_visibility = ["PUBLIC"])

go_toolchain(
    name = "toolchain",
    version = "1.21.3",
)

go_repo(
    name = "x_sys",
    install = ["unix"],
    module = "golang.org/x/sys",
    version = "0b90fd5c4c48ef7ae7ef6c393cb118d2e7969999",
)

go_repo(
    name = "validator",
    install = ["."],
    module = "github.com/go-playground/validator/v10",
    version = "v10.4.1",
)

go_repo(
    name = "spew",
    install = ["spew"],
    module = "github.com/davecgh/go-spew",
    version = "d8f796af33cc11cb798c1aaeb27a4ebc5099927d",
)

go_repo(
    name = "testify",
    install = [
        "assert",
        "require",
    ],
    module = "github.com/stretchr/testify",
    version = "v1.3.0",
)

go_repo(
    name = "gomega",
    install = [
        ".",
        "format",
        "matchers",
        "matchers/support/goraph/bipartitegraph",
        "matchers/support/goraph/node",
        "matchers/support/goraph/edge",
        "matchers/support/goraph/util",
        "types",
    ],
    module = "github.com/onsi/gomega",
    version = "v1.11.0",
)

go_repo(
    name = "fsnotify",
    install = ["."],
    module = "github.com/fsnotify/fsnotify",
    version = "v1.4.9",
)

go_repo(
    name = "tail",
    install = [
        ".",
        "...",
    ],
    module = "github.com/nxadm/tail",
    version = "v1.4.8",
)

go_repo(
    name = "urn",
    install = ["."],
    module = "github.com/leodido/go-urn",
    version = "v1.2.1",
)

go_repo(
    name = "go-codec",
    install = ["."],
    module = "github.com/ugorji/go/codec",
    version = "v1.2.4",
)

go_repo(
    name = "difflib",
    install = ["difflib"],
    module = "github.com/pmezard/go-difflib",
    version = "5d4384ee4fb2527b0a1256a821ebfc92f91efefc",
)

go_repo(
    name = "ginkgo",
    install = [
        ".",
        "config",
        "reporters",
        "reporters/stenographer",
        "reporters/stenographer/support/go-colorable",
        "types",
    ],
    module = "github.com/onsi/ginkgo",
    version = "v1.15.1",
)

go_repo(
    name = "gin",
    install = [
        ".",
        "binding",
        "render",
    ],
    module = "github.com/gin-gonic/gin",
    version = "v1.6.3",
)

go_repo(
    name = "locales",
    install = [
        ".",
        "currency",
    ],
    module = "github.com/go-playground/locales",
    version = "v0.13.0",
)

go_repo(
    name = "yaml",
    install = ["."],
    module = "gopkg.in/yaml.v2",
    version = "v2.4.0",
)

go_repo(
    name = "universal-translator",
    install = ["."],
    module = "github.com/go-playground/universal-translator",
    version = "v0.17.0",
)

go_repo(
    name = "x_crypto",
    install = ["sha3"],
    module = "golang.org/x/crypto",
    version = "e6e6c4f2bb5b5887c7f7dd52f01ea7b2fbeb297d",
)

go_repo(
    name = "protobuf",
    install = ["proto"],
    module = "github.com/golang/protobuf",
    version = "v1.3.3",
)

go_repo(
    name = "typed",
    install = ["."],
    module = "github.com/karlseguin/typed",
    version = "v1.1.7",
)

go_repo(
    name = "tomb",
    install = ["."],
    module = "gopkg.in/tomb.v1",
    version = "dd632973f1e7218eb1089048e0798ec9ae7dceb8",
)

go_repo(
    name = "x_net",
    install = [
        "html",
        "html/atom",
        "html/charset",
    ],
    module = "golang.org/x/net",
    version = "d523dce5a7f4b994f7ed0531dbe44cd8fd803e26",
)

go_repo(
    name = "x_text",
    install = [
        "encoding",
        "encoding/charmap",
        "encoding/htmlindex",
        "transform",
        "language",
        "encoding/japanese",
        "encoding/korean",
        "encoding/simplifiedchinese",
        "encoding/traditionalchinese",
        "encoding/unicode",
        "runes",
    ],
    module = "golang.org/x/text",
    version = "e3aa4adf54f644ca0cb35f1f1fb19b239c40ef04",
)

go_repo(
    name = "sse",
    install = ["."],
    module = "github.com/gin-contrib/sse",
    version = "v0.1.0",
)

go_repo(
    name = "isatty",
    install = ["."],
    module = "github.com/mattn/go-isatty",
    version = "v0.0.12",
)
