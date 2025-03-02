#!/bin/bash
# Install Node.js dependencies needed for Prisma
go install github.com/steebchen/prisma-client-go@latest

# Generate Prisma client
go run github.com/steebchen/prisma-client-go generate


# Continue with normal Go build