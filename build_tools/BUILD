load("//build_tools/sh:sh.bzl", "dbx_sh_binary")

package(default_visibility = ["//visibility:public"])

config_setting(
    name = "services-verbose",
    values = {
        "define": "svc.verbose=1",
    },
)

config_setting(
    name = "noservices-verbose",
    values = {
        "define": "svc.verbose=0",
    },
)

config_setting(
    name = "disable-service-deps",
    values = {
        "define": "dbx_service_deps=disabled",
    },
)

# Used internally by Changes to not create a version file. This is so that tests are remote-cachable
# Do not set this by hand
config_setting(
    name = "services-version-file",
    values = {
        "define": "svc.version-file=1",
    },
)

config_setting(
    name = "noservices-version-file",
    values = {
        "define": "svc.version-file=0",
    },
)

cc_binary(
    name = "pass_binary",
    srcs = ["exit0.c"],
)

cc_test(
    name = "pass",
    srcs = ["exit0.c"],
)
