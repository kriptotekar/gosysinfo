# gosysinfo

A simple cross‑platform system information reporter written in Go, using [gopsutil](https://github.com/shirou/gopsutil).  
Prints OS, kernel, hostname, CPU, memory, swap, network I/O, disk usage and top processes by CPU.

## Features

- Host / OS info (platform, version, kernel, hostname)  
- CPU info (cores, model name)  
- Memory & swap usage  
- Network interface I/O stats  
- Disk partitions & usage  
- Top 5 processes by CPU usage  

## Prerequisites

- Go 1.18 or newer  
- Git (to clone the repo)

## Installation

```bash
# Clone the repo
git clone https://github.com/kriptotekar/gosysinfo.git
cd gosysinfo

# Initialize and download dependencies
go mod tidy