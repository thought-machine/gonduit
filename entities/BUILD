go_library(
    name = "entities",
    srcs = glob(
        ["*.go"],
        exclude = ["*_test.go"],
    ),
    visibility = [
        "//:all",
        "//requests:all",
        "//responses:all",
    ],
    deps = [
        "//constants",
        "//util",
    ],
)

go_test(
    name = "entities_test",
    srcs = glob(["*_test.go"]),
    external = True,
    deps = [
        ":entities",
        "//third_party/go:ginkgo",
        "//third_party/go:gomega",
    ],
)
