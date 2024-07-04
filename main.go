package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type Process struct {
	PID         int
	Name        string
	CPUUsage    float64
	MemoryUsage float64
}

var (
	cpuThreshold    = flag.Float64("cpu", 20.0, "CPU usage threshold in percentage")
	memoryThreshold = flag.Float64("mem", 20.0, "Memory usage threshold in percentage")
	checkInterval   = flag.Int("interval", 10, "Interval between checks in seconds")
	logFileName     = flag.String("logFile", "log.txt", "Name of the log file")
	createLogFile   = flag.Bool("createLog", true, "Whether to create a log file")
)

var logger *log.Logger

func setupLogger() {
	var logOutput io.Writer = os.Stdout
	if *createLogFile {
		file, err := os.OpenFile(*logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		logOutput = io.MultiWriter(os.Stdout, file)
		defer file.Close()
	}
	logger = log.New(logOutput, "", log.LstdFlags)
}

func fetchProcesses() ([]Process, error) {
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(out.String(), "\n")
	var processes []Process
	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		if len(fields) < 10 {
			continue
		}
		pid, err := strconv.Atoi(fields[1])
		if err != nil {
			continue
		}
		cpuUsage, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			continue
		}
		memoryUsage, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			continue
		}
		processes = append(processes, Process{
			PID:         pid,
			Name:        fields[10],
			CPUUsage:    cpuUsage,
			MemoryUsage: memoryUsage,
		})
	}
	return processes, nil
}

func logHighUsageProcesses(processes []Process) {
	for _, process := range processes {
		if process.CPUUsage > *cpuThreshold || process.MemoryUsage > *memoryThreshold {
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			logger.Printf("[%s] High usage detected: PID=%d, Name=%s, CPU=%.2f%%, Memory=%.2f%%\n",
				currentTime, process.PID, process.Name, process.CPUUsage, process.MemoryUsage)
		}
	}
}

func main() {
	flag.Parse() // Parse the command-line flags
	setupLogger()

	// Setup channel for graceful shutdown
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-shutdownChan
		if *createLogFile {
			logger.Printf("Shutting down and logging to %s\n", *logFileName)
		} else {
			logger.Println("Shutting down")
		}
		os.Exit(0)
	}()

	for {
		processes, err := fetchProcesses()
		if err != nil {
			logger.Fatalf("Failed to fetch processes: %v", err)
		}
		logHighUsageProcesses(processes)
		time.Sleep(time.Duration(*checkInterval) * time.Second)
	}
}
