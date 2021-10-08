load("@bazel_skylib//:bzl_library.bzl", "bzl_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "bazel",
    srcs = ["env.go"],
    importpath = "github.com/dropbox/dbx_build_tools/build_tools/bazel",
    visibility = ["//visibility:public"],
)

bzl_library(
    name = "bazel_bzl",
    srcs = glob([
        "*.bzl",
    ]),
    deps = [
        "//build_tools/windows:windows_bzl",
    ],
)

exports_files([
    "runfiles.tmpl",
    "runfiles.bat.tmpl",
])