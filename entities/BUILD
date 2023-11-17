go_library(
    name = "entities",
    srcs = [
        "applicationeditor.go",
        "common.go",
        "conduit.go",
        "differential.go",
        "diffusion.go",
        "harbormaster.go",
        "maniphest.go",
        "maniphest_task_transaction.go",
        "paste.go",
        "phid.go",
        "phriction.go",
        "projects.go",
        "remarkup.go",
        "repository.go",
        "resolvedref.go",
        "session.go",
        "user.go",
    ],
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
    srcs = [
        "applicationeditor_test.go",
        "entities_suite_test.go",
    ],
    external = True,
    deps = [
        ":entities",
        "//third_party/go:ginkgo",
        "//third_party/go:gomega",
    ],
)
