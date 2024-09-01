#!/usr/bin/env nix-shell
#!nix-shell -i python3 -p python3

"""
This place is not a place of honor...
no highly esteemed deed is commemorated here...
nothing valued is here.
"""

from pathlib import Path
import hashlib
import re
import shutil
import subprocess

PACKAGE = "github.com/zopieux/dumpsys-exporter"
SRC = "vendored/"
OUT = "proto"
TMP = "proto_out"

def run(to_compile, paths):
    opts = [f"--go_opt=M{e}={PACKAGE}/{OUT}/android/{str(Path(e).parent)}" for e in paths]
    r = subprocess.run([
        "protoc", "-I", SRC, f"--proto_path={SRC}", f"--go_out={TMP}"]
        + opts + to_compile, capture_output=True)
    if r.returncode == 1:
        stderr = r.stderr.decode()
        m_path = re.search(r'import path for "(.+?)"', stderr)
        if m_path:
            paths.append(m_path.group(1))
            return False
        raise RuntimeError(f"Unhandled: {stderr}")
    return True

if __name__ == '__main__':
    tmp_out = Path(TMP)
    shutil.rmtree(tmp_out, ignore_errors=True)
    tmp_out.mkdir()
    inner = tmp_out / PACKAGE / OUT

    to_compile = [f"{SRC}{p}" for p in [
        "frameworks/base/core/proto/android/content/component_name.proto",
        "frameworks/base/core/proto/android/content/intent.proto",
        "frameworks/base/core/proto/android/os/looper.proto",
        "frameworks/base/core/proto/android/os/message.proto",
        "frameworks/base/core/proto/android/os/messagequeue.proto",
        "frameworks/base/core/proto/android/os/patternmatcher.proto",
        "frameworks/base/core/proto/android/os/persistablebundle.proto",
        "frameworks/base/core/proto/android/os/powermanager.proto",
        "frameworks/base/core/proto/android/os/worksource.proto",
        "frameworks/base/core/proto/android/privacy.proto",
        "frameworks/base/core/proto/android/providers/settings.proto",
        "frameworks/base/core/proto/android/providers/settings/common.proto",
        "frameworks/base/core/proto/android/providers/settings/config.proto",
        "frameworks/base/core/proto/android/providers/settings/generation.proto",
        "frameworks/base/core/proto/android/providers/settings/global.proto",
        "frameworks/base/core/proto/android/providers/settings/secure.proto",
        "frameworks/base/core/proto/android/providers/settings/system.proto",
        "frameworks/base/core/proto/android/server/powermanagerservice.proto",
        "frameworks/base/core/proto/android/server/wirelesschargerdetector.proto",
        "frameworks/base/core/proto/android/service/procstats.proto",
        "frameworks/base/core/proto/android/util/common.proto",
        "frameworks/proto_logging/stats/enums/app/enums.proto",
        "frameworks/proto_logging/stats/enums/os/enums.proto",
        "frameworks/proto_logging/stats/enums/service/procstats_enum.proto",
        "frameworks/proto_logging/stats/enums/service/enums.proto",
        "frameworks/proto_logging/stats/enums/view/enums.proto",
    ]]
    paths = []
    while True:
        match run(to_compile, paths):
            case True:
                break
            case False:
                pass

    print(f"Generation success with {len(paths)} substitutions.")
    out = Path(OUT)
    print(f"Moving {inner} -> {out}")
    shutil.rmtree(out, ignore_errors=True)
    inner.rename(out)
    shutil.rmtree(tmp_out)
