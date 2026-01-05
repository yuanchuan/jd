#!/bin/bash

# Detect shell and config file
if [ -n "$ZSH_VERSION" ]; then
    CONFIG="$HOME/.zshrc"
elif [ -n "$BASH_VERSION" ]; then
    CONFIG="$HOME/.bashrc"
else
    echo "Unsupported shell"
    exit 1
fi

# Install binary
echo "Installing jd binary..."
go install github.com/yuanchuan/jd@latest

# Get jd path
JD_PATH=$(which jd)

if [ -z "$JD_PATH" ]; then
    echo "Error: jd binary not found in PATH"
    exit 1
fi

# Add function if not exists
if ! grep -q "function jd" "$CONFIG" 2>/dev/null; then
    cat >> "$CONFIG" << EOF

# JD - Jump to directories quickly
function jd {
    local path
    path=\$(\$JD_PATH \$@) || return
    builtin cd "\$path"
}
EOF
    echo "Added jd function to $CONFIG"
else
    echo "jd function already exists in $CONFIG"
fi

echo ""
echo "Setup complete! Run 'source $CONFIG' or restart your shell to start using jd"