#!/usr/bin/env bash

pushd "gen/proto"
echo "📁 Generating protocol buffers from gen/proto..."
./generate.sh
popd

echo "🏁 Setup complete!"
