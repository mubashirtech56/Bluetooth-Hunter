# Bluetooth Hunter

## System Architecture

### 1. Project Overview

Bluetooth Hunter is a Linux desktop application developed in Go. It provides a graphical interface for discovering, monitoring, and managing nearby Bluetooth devices.

The application uses the Linux BlueZ Bluetooth stack to communicate with Bluetooth hardware.

### 2. Technology Stack

| Component            | Technology  |
| -------------------- | ----------- |
| Programming Language | Go (Golang) |
| Desktop GUI          | Fyne        |
| Bluetooth Backend    | BlueZ       |
| System Communication | D-Bus       |
| Database             | SQLite      |
| Operating System     | Linux       |
| Build System         | Go Modules  |

### 3. High-Level Architecture

```text
+--------------------------------------+
|          Bluetooth Hunter            |
|                                      |
|  +--------------------------------+  |
|  |          Desktop GUI           |  |
|  |             Fyne               |  |
|  +--------------------------------+  |
|                  |                   |
|  +--------------------------------+  |
|  |        Application Layer       |  |
|  |                                |  |
|  | Device Manager                 |  |
|  | Scanner Manager                |  |
|  | Connection Manager             |  |
|  | Signal Monitor                 |  |
|  +--------------------------------+  |
|                  |                   |
|  +--------------------------------+  |
|  |       Bluetooth Service       |  |
|  |                                |  |
|  |          Go + D-Bus            |  |
|  +--------------------------------+  |
+--------------------------------------+
                   |
                   v
+--------------------------------------+
|             Linux BlueZ              |
|                                      |
|       Bluetooth System Service       |
+--------------------------------------+
                   |
                   v
+--------------------------------------+
|          Bluetooth Adapter           |
|                                      |
|       Nearby Bluetooth Devices       |
+--------------------------------------+
```

### 4. Core Modules

#### 4.1 Device Scanner

- Discover nearby Bluetooth devices.
- Detect Bluetooth Classic and BLE devices where supported.
- Retrieve available device names and addresses.
- Display device discovery results in real time.

#### 4.2 Device Manager

- Maintain a list of discovered devices.
- Display device names and MAC addresses when available.
- Show connection and pairing status.
- Identify newly discovered and previously detected devices.

#### 4.3 Connection Manager

- Display currently connected devices.
- Connect to supported devices.
- Disconnect devices belonging to the user.
- Monitor connection status changes.

#### 4.4 Signal Monitor

- Read RSSI (Received Signal Strength Indicator).
- Display signal strength.
- Monitor changes in signal strength over time.
- Provide an approximate proximity indicator.

#### 4.5 Dashboard

- Display discovered devices.
- Show connected and disconnected device counts.
- Provide scanning controls.
- Display device information and connection status.

#### 4.6 Notification System

- Notify when a new device is discovered.
- Notify when a monitored device disappears.
- Notify when a managed connection changes status.

### 5. Project Directory Structure

```text
bluetooth-hunter/
│
├── cmd/
│   └── bluetooth-hunter/
│       └── main.go
│
├── internal/
│   ├── bluetooth/
│   │   ├── adapter.go
│   │   ├── scanner.go
│   │   ├── device.go
│   │   └── connection.go
│   │
│   ├── manager/
│   │   ├── device_manager.go
│   │   └── connection_manager.go
│   │
│   ├── monitor/
│   │   └── signal_monitor.go
│   │
│   ├── storage/
│   │   └── database.go
│   │
│   └── notifications/
│       └── notification.go
│
├── internal/ui/
│   ├── dashboard.go
│   ├── device_list.go
│   ├── device_details.go
│   └── components.go
│
├── assets/
│   ├── icons/
│   └── images/
│
├── docs/
│   └── architecture.md
│
├── go.mod
├── go.sum
├── README.md
└── LICENSE
```

### 6. Bluetooth Communication

Bluetooth Hunter communicates with BlueZ through D-Bus.

Important BlueZ interfaces:

| Interface                       | Purpose                                      |
| ------------------------------- | -------------------------------------------- |
| org.bluez.Adapter1              | Bluetooth adapter management                 |
| org.bluez.Device1               | Device information and connection management |
| org.freedesktop.DBus.Properties | Read and monitor properties                  |
| org.bluez.Agent1                | Pairing and authentication handling          |

The application will use D-Bus methods and properties to interact with the local Bluetooth adapter.

### 7. Data Model

Each discovered device will have the following information where available:

```go
type BluetoothDevice struct {
    Name         string
    Address      string
    RSSI         int16
    Connected    bool
    Paired       bool
    Trusted      bool
    Discoverable bool
}
```

### 8. Application Workflow

1. Application starts.
2. Connect to the system D-Bus.
3. Detect available Bluetooth adapters.
4. Retrieve the adapter state.
5. Start device discovery.
6. Receive device discovery events.
7. Update the device list.
8. Display device information in the GUI.
9. Monitor device and connection status.
10. Allow users to manage their own Bluetooth connections.

### 9. Security and Limitations

- The application will only manage connections authorized by the local Bluetooth system.
- It will not attempt to interfere with third-party Bluetooth connections.
- It will not use Bluetooth jamming or radio interference.
- Device discovery depends on device discoverability and Bluetooth advertising behavior.
- RSSI does not provide an exact physical distance.
- Some Bluetooth devices may not expose complete device information.

### 10. Development Phases

| Phase   | Task                             |
| ------- | -------------------------------- |
| Phase 1 | Go environment and Fyne setup    |
| Phase 2 | BlueZ and D-Bus integration      |
| Phase 3 | Bluetooth device discovery       |
| Phase 4 | Device information management    |
| Phase 5 | Connection management            |
| Phase 6 | Signal monitoring                |
| Phase 7 | Desktop dashboard                |
| Phase 8 | Notifications and SQLite storage |
| Phase 9 | Testing and packaging            |

### 11. Expected Outcome

A lightweight Linux desktop application that provides real-time Bluetooth device discovery, device information, connection monitoring, and authorized connection management through an intuitive graphical interface.
