go_library(
    name = "responses",
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
        "//util",
    ],
)
