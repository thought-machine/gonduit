go_library(
    name = "gonduit",
    srcs = glob(
        ["*.go"],
        exclude = ["*_test.go"],
    ),
    deps = [
        "//core",
        "//entities",
        "//requests",
        "//responses",
        "//util",
    ],
)

go_test(
    name = "gonduit_test",
    srcs = glob(
        ["*_test.go"],
    ),
    deps = [
        ":gonduit",
        "//core",
        "//responses",
        "//test/server",
        "//third_party/go:gin",
        "//third_party/go:testify",
    ],
)
