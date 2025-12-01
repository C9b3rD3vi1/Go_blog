#!/bin/bash

# Go to project directory
cd /home/user/goapps/Go_blog

# Pull latest changes from GitHub
git pull origin main

# Build Go binary
go build -o Go_blog

# Stop old process (if running)
pkill Go_blog || true

# Start the app in background
nohup ./Go_blog > app.log 2>&1 &
