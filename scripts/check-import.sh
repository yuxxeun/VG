#!/bin/bash

echo "🔍 Checking Go imports..."

# Check if all imports are valid
if go list ./... > /dev/null 2>&1; then
    echo "✅ All imports are valid"
else
    echo "❌ Import errors found"
    go list ./...
    exit 1
fi

# Check for circular dependencies
echo "🔄 Checking for circular dependencies..."
if go list -json ./... | jq -r '.Deps[]' | sort | uniq -d; then
    echo "⚠️  Potential circular dependencies detected"
else
    echo "✅ No circular dependencies"
fi

echo "✅ All checks passed!"