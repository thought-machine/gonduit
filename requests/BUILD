go_library(
    name = "requests",
    srcs = glob(
        ["*.go"],
        exclude = ["*_test.go"],
    ),
    visibility = [
        "//:all",
        "//core:all",
    ],
    deps = [
        "//constants",
        "//entities",
    ],
)
