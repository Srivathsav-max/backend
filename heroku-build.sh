#!/bin/sh

echo "🔧 Setting up build environment..."
export DEVELOPMENT=false

echo "📌 Running Prisma Generate..."
# Ensure Prisma engines are available
mkdir -p prisma/db

# Verify database configuration before running generate
echo "🔍 Verifying database configuration..."
if [ -z "$DATABASE_URL" ]; then
    echo "❌ Error: DATABASE_URL environment variable is not set"
    exit 1
fi

if [ -z "$DIRECT_URL" ]; then
    echo "❌ Error: DIRECT_URL environment variable is not set"
    exit 1
fi

echo "✅ Database configuration verified"

# Run Prisma generate
go run github.com/steebchen/prisma-client-go generate

echo "✅ Prisma Client Generated!"

# Ensure binary permissions
chmod +x bin/backend

echo "✅ Build setup complete!"
