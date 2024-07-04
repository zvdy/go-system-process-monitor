# 📊 Golang System Process Monitor

This project is a simple CLI tool that monitors system processes, logging any that exceed predefined CPU and memory usage thresholds. It's designed to help system administrators and developers keep an eye on resource-intensive processes.

## 🚀 Features

- **Process Monitoring:** Fetches running processes and their CPU and memory usage.
- **Threshold Alerts:** Logs processes that exceed CPU or memory usage thresholds.
- **Customizable Checks:** Allows setting custom thresholds for monitoring.
- **Graceful Shutdown:** Supports graceful shutdown by pressing `Ctrl+C`, ensuring all logs are properly saved.
- **Log to File:** Optionally logs all entries to a file upon shutdown.
- **Customizable Log File:** Allows customization of the log file name through command-line flags.

## 🛠️ Setup

To get started with the System Process Monitor, you'll need to have Go installed on your system. Follow these steps:

1. **Clone the repository:**

```bash
git clone https://github.com/zvdy/go-system-process-monitor.git
cd go-system-process-monitor
```

2. **Build the project:**

```bash
go build -o process-monitor
```

## 📝 How to Use

After building the project, you can run the monitor by executing the binary:

```bash
./process-monitor
```

### Customizing Thresholds, Check Interval, and Log File

You can customize the CPU and memory thresholds, the check interval, and the log file settings using command-line flags:

- **CPU Threshold** (`-cpu`): Set the CPU usage threshold in percentage.
- **Memory Threshold** (`-mem`): Set the memory usage threshold in percentage.
- **Check Interval** (`-interval`): Set the interval between checks in seconds.
- **Log File Name** (`-logFile`): Specify the name of the log file.
- **Create Log File** (`-createLog`): Decide whether to create a log file (`true` or `false`).

Example:

```bash
./process-monitor -cpu 5.0 -mem 5.0 -interval 5 -logFile "custom_log.txt" -createLog true
```

This command sets the CPU and memory thresholds to 5%, the check interval to 5 seconds, and specifies a custom log file name while enabling log file creation.

> Default values are 20% for CPU & memory, the interval is 10 seconds, the log file is named `log.txt`, and log file creation is enabled by default.

## 📚 Documentation

For more detailed information about the project and its implementation, refer to the source code comments.

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check [issues page](https://github.com/zvdy/go-system-process-monitor/issues) for open issues or to open a new issue.

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

## 📢 Acknowledgements

- This project is built using Go.
- Special thanks to the Go community for the comprehensive documentation and resources.

Happy Monitoring! 🎉