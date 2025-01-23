# repro_bazel_nvml_cgo
Repro the issue of bazel / ego / building with dependency on github.com/NVIDIA/go-nvml/pkg/nvml/

# Failure

```
$ bazel test --platforms=@rules_go//go/toolchain:linux_amd64_cgo //... --sandbox_debug
INFO: Analyzed 2 targets (0 packages loaded, 0 targets configured).
ERROR: /private/var/tmp/_bazel_allanc/a58c7937a47f267c72aa5981db16520f/external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/BUILD.bazel:3:11:
GoCompilePkg external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/nvml.a failed: (Exit 1):
sandbox-exec failed: error executing GoCompilePkg command
  (cd /private/var/tmp/_bazel_allanc/a58c7937a47f267c72aa5981db16520f/sandbox/darwin-sandbox/117/execroot/_main && \
  exec env - \
    CGO_ENABLED=0 \
    GOARCH=amd64 \
    GODEBUG='winsymlink=0' \
    GOEXPERIMENT=nocoverageredesign \
    GOOS=linux \
    GOPATH='' \
    GOROOT_FINAL=GOROOT \
    GOTOOLCHAIN=local \
    TMPDIR=/var/folders/qd/kfsthb811k93pvfddylzzcqh0000gn/T/ \
  /usr/bin/sandbox-exec \
-f /private/var/tmp/_bazel_allanc/a58c7937a47f267c72aa5981db16520f/sandbox/darwin-sandbox/117/sandbox.sb \
  /var/tmp/_bazel_allanc/install/2e917a78c62d9cd541741ddd561b410e/process-wrapper \
'--timeout=0' '--kill_delay=15' \
'--stats=/private/var/tmp/_bazel_allanc/a58c7937a47f267c72aa5981db16520f/sandbox/darwin-sandbox/117/stats.out' \
  bazel-out/darwin_arm64-opt-exec-ST-d57f47055a04/bin/external/rules_go++go_sdk+repro_bazel_nvml_cgo__download_0/builder_reset/builder \
  compilepkg \
-sdk external/rules_go++go_sdk+repro_bazel_nvml_cgo__download_0 \
-goroot bazel-out/darwin_arm64-fastbuild/bin/external/rules_go+/stdlib_ -installsuffix linux_amd64 \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/api.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/cgo_helpers.h \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/cgo_helpers_static.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/const.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/const_static.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/device.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/doc.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/dynamicLibrary_mock.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/event_set.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/gpm.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/init.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/lib.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/nvml.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/nvml.h \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/refcount.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/return.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/system.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/types_gen.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/unit.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/vgpu.go \
-src external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/zz_generated.api.go \
-arc 'github.com/NVIDIA/go-nvml/pkg/dl=github.com/NVIDIA/go-nvml/pkg/dl=bazel-out/darwin_arm64-fastbuild/bin/external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/dl/dl.x' \
-importpath github.com/NVIDIA/go-nvml/pkg/nvml \
-p github.com/NVIDIA/go-nvml/pkg/nvml \
-package_list bazel-out/darwin_arm64-opt-exec-ST-d57f47055a04/bin/external/rules_go++go_sdk+repro_bazel_nvml_cgo__download_0/packages.txt \
-embedroot '' \
-embedroot bazel-out/darwin_arm64-fastbuild/bin \
-embedlookupdir external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml \
-lo bazel-out/darwin_arm64-fastbuild/bin/external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/nvml.a \
-o bazel-out/darwin_arm64-fastbuild/bin/external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/nvml.x \
-gcflags '')
external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/zz_generated.api.go:677:24: undefined: Return
external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/zz_generated.api.go:678:21: undefined: Return
external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/zz_generated.api.go:679:22: undefined: EccCounterType
external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/zz_generated.api.go:679:38: undefined: Return
external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/zz_generated.api.go:680:33: undefined: Return
external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/zz_generated.api.go:910:71: undefined: Return
external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/zz_generated.api.go:911:111: undefined: Return
external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/zz_generated.api.go:912:12: undefined: Return
external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/zz_generated.api.go:926:12: undefined: Return
external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/zz_generated.api.go:927:34: undefined: Return
external/gazelle++go_deps+com_github_nvidia_go_nvml/pkg/nvml/zz_generated.api.go:912:12: too many errors
compilepkg: error running subcommand external/rules_go++go_sdk+repro_bazel_nvml_cgo__download_0/pkg/tool/darwin_arm64/compile: exit status 2
Use --verbose_failures to see the command lines of failed build steps.
INFO: Elapsed time: 0.197s, Critical Path: 0.06s
INFO: 10 processes: 8 internal, 2 darwin-sandbox.
ERROR: Build did NOT complete successfully
//:repro_bazel_nvml_cgo_test                                    FAILED TO BUILD

Executed 0 out of 1 test: 1 fails to build.
```

# NOTES

## CGO_ENABLED

I see `CGO_ENABLED=0` and that looks suspicious but adding `cgo = True` to the `go_test()` rule has no effect.

## Darwin -> Linux cross

This is a darwin/arm64 so it's cross-compiling, but compile fails with similar results for the
following values of --platform:
 - (nothing)
 - --platforms=@rules_go//go/toolchain:linux_amd64
 - --platforms=@rules_go//go/toolchain:linux_amd64_cgo
 - --platforms=@rules_go//go/toolchain:linux_arm64
 - --platforms=@rules_go//go/toolchain:linux_arm64_cgo
