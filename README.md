# 🔎 Oculusinfo
Oculusinfo is a lightweight cross-platform CLI tool written in Go for network reconnaissance tasks such as WHOIS lookup, DNS resolution, and port scanning.

It is designed to provide fast, structured, and readable network information directly from the terminal.

---

## ⚙️ Features

### 🧾 WHOIS lookup
Get basic IP/domain information (geolocation, ISP, etc.)

### 🌐 DNS lookup (dig-like)
Resolve common DNS records:
- A / AAAA
- MX
- NS
- TXT

### 🔓 Port scanner
Fast TCP port scanning using concurrency

---

## 🚀 Installation

### Build from source

```bash
git clone https://github.com/yourname/oculusinfo.git
cd oculusinfo
go build -o oculusinfo ./cmd
