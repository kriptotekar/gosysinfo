package main

import (
    "fmt"
    "sort"

    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/disk"
    "github.com/shirou/gopsutil/v3/host"
    "github.com/shirou/gopsutil/v3/mem"
    "github.com/shirou/gopsutil/v3/net"
    "github.com/shirou/gopsutil/v3/process"
)

type procStat struct {
    pid  int32
    name string
    cpu  float64
    mem  uint64
}

func main() {
    // Host / OS info
    hi, _ := host.Info()
    fmt.Printf("OS: %s %s\n", hi.Platform, hi.PlatformVersion)
    fmt.Printf("Kernel: %s\n", hi.KernelVersion)
    fmt.Printf("Hostname: %s\n", hi.Hostname)

    // CPU info
    cores, _ := cpu.Counts(true)
    infos, _ := cpu.Info()
    fmt.Printf("\n--- CPU Information ---\n")
    fmt.Printf("CPU Cores: %d\n", cores)
    if len(infos) > 0 {
        fmt.Printf("CPU Model: %s\n", infos[0].ModelName)
    }

    // Memory
    vm, _ := mem.VirtualMemory()
    sm, _ := mem.SwapMemory()
    fmt.Printf("\n--- Memory Information ---\n")
    fmt.Printf("Total: %d KB\nUsed:  %d KB\n", vm.Total/1024, vm.Used/1024)
    fmt.Printf("Swap Total: %d KB\nSwap Used:  %d KB\n", sm.Total/1024, sm.Used/1024)

    // Network I/O
    fmt.Printf("\n--- Network Interfaces ---\n")
    nets, _ := net.IOCounters(true)
    if len(nets) == 0 {
        fmt.Println("No network interfaces found.")
    }
    for _, n := range nets {
        fmt.Printf("[%s] Sent: %d bytes, Recv: %d bytes\n",
            n.Name, n.BytesSent, n.BytesRecv)
    }

    // Disk info
    fmt.Printf("\n--- Disk Information ---\n")
    parts, _ := disk.Partitions(false)
    if len(parts) == 0 {
        fmt.Println("No disks found.")
    }
    for _, p := range parts {
        du, _ := disk.Usage(p.Mountpoint)
        fmt.Printf("[%s] %s | FS: %s | Total: %d bytes, Free: %d bytes\n",
            p.Device, p.Mountpoint, p.Fstype, du.Total, du.Free)
    }

    // Top 5 processes by CPU
    procs, _ := process.Processes()
    var stats []procStat
    for _, p := range procs {
        if cpuPct, err := p.CPUPercent(); err == nil {
            if mi, err := p.MemoryInfo(); err == nil {
                if name, err := p.Name(); err == nil {
                    stats = append(stats, procStat{p.Pid, name, cpuPct, mi.RSS})
                }
            }
        }
    }
    sort.Slice(stats, func(i, j int) bool { return stats[i].cpu > stats[j].cpu })
    fmt.Printf("\n--- Top 5 Processes (by CPU) ---\n")
    for i, s := range stats {
        if i >= 5 {
            break
        }
        fmt.Printf("PID: %d, Name: %s, CPU: %.1f%%, Mem: %d KB\n",
            s.pid, s.name, s.cpu, s.mem/1024)
    }
}