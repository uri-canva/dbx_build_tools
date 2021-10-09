load("//build_tools/sh:sh.bzl", "dbx_sh_test")
load("@bazel_skylib//:bzl_library.bzl", "bzl_library")

package(default_visibility = ["//visibility:public"])

dbx_sh_test(
    name = "restart_test",
    srcs = ["restart_test.sh"],
    data = [
        "//build_tools/svcctl/cmd/svcctl",
    ],
    tags = ["manual"],  # not meant to be run on its own. needs services
)

bzl_library(
    name = "services_bzl",
    srcs = [
        "svc.bzl",
    ],
)
