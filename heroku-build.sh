#!/bin/sh

echo "📌 Running Prisma Generate..."
go run github.com/steebchen/prisma-client-go generate

echo "✅ Prisma Client Generated!"

# Ensure build script permissions
chmod +x app
