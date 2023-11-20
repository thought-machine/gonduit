go_library(
    name = "core",
    srcs = glob(
        ["*.go"],
        exclude = ["*_test.go"],
    ),
    visibility = ["//:all"],
    deps = [
        "//requests",
        "//third_party/go:typed",
    ],
)

go_test(
    name = "core_test",
    srcs = glob(["*_test.go"]),
    deps = [
        ":core",
        "//requests",
        "//responses",
        "//third_party/go:gin",
        "//third_party/go:testify",
    ],
)
