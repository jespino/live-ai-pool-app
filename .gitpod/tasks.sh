#!/bin/bash

# Install Air for hot reloading
if [ ! -f /usr/bin/air ]; then
  echo "Installing Air for hot reloading..."
  curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/bin
fi

# Install golangci-lint
if [ ! -f /usr/bin/golangci-lint ]; then
  echo "Installing golangci-lint..."
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/bin v1.54.2
fi

# Set permissions
chmod +x /usr/bin/air
chmod +x /usr/bin/golangci-lint

echo "Tools installation complete!"