#!/bin/bash

set -x

sudo /usr/local/cargo/bin/cargo binstall --disable-telemetry --no-confirm --strategies crate-meta-data jj-cli

gh repo clone minusnine/dotfiles /workspaces/dotfiles

for dir in /workspaces/*; do
    jj -R $dir git init && \
    jj -$ $dir bookmark track master@origin; \
done