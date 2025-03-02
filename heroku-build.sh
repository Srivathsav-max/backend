#!/bin/sh

echo "📌 Running Prisma Generate..."
go run github.com/steebchen/prisma-client-go generate

echo "✅ Prisma Client Generated! Running go install..."
go install -v -tags heroku .
