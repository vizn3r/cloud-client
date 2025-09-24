# Cloud Storage Thingy (WIP - no official releases yet)

*"What are you working on?" - classmate*  
*"Just... a cloud storage... thingy" - vizn3r*

## The Vision

A self-hosted cloud storage solution with privacy-first design, featuring both traditional cloud storage and innovative P2P file sharing capabilities.

<p align="center">
  <img src="https://raw.githubusercontent.com/vizn3r/cloud/refs/heads/main/documentation/server-client.png">
  </br>
  <em>Client-Server relations</em>
  </br>
  </br>
  <img src="https://raw.githubusercontent.com/vizn3r/cloud/refs/heads/main/documentation/file-encryption.png">
  </br>
  <em>Client-Server E2E architecture</em>
  </br>
  </br>
</p>

## Core Features

### Security & Privacy First
- **End-to-End Encryption**: Files are encrypted client-side before upload
- **Zero-Knowledge Architecture**: Server never sees your file contents
- **Hybrid Security Model**: Choose between maximum security or convenience features
- **User-Controlled Keys**: You control your encryption keys, not us

### Traditional Cloud Storage (Client 1)
- Upload, download, and manage files through web interface
- File organization with folders and metadata
- Secure file sharing with time-limited access
- REST API for integration with external tools

### Smart File Sharing (Client 4)  
- **Public/Private Shares**: Create shareable links with expiration dates
- **ShareX Integration**: Direct upload support for screenshot tools
- **Rich Embeds**: Preview files on Discord, Twitter, Facebook
- **Access Control**: Set download limits and view permissions

### P2P File Transfer (Clients 2 & 3) *[Planned]*
- **Direct Device-to-Device**: Share files without uploading to server
- **Server as Orchestrator**: Facilitates P2P connections between your devices
- **Real-time Sync**: Live file synchronization across connected devices
- **Offline-First**: Access files even when server is unreachable

## Current Implementation Status

### Completed
- **Backend Server**: Go-based REST API with Fiber framework
- **Database Layer**: SQLite with user management and file ownership
- **Authentication**: Secure session-based auth with bcrypt password hashing
- **File Storage**: Local filesystem with metadata management
- **File Sharing**: Temporary share links with expiration
- **Docker Support**: Containerized deployment with BuildKit optimization

### In Progress
- **Web Client**: React-based responsive file management interface
- **Encryption Layer**: Client-side E2E encryption implementation
- **Advanced Sharing**: Enhanced embed support and access controls

### Planned
- **Desktop App**: Native application for P2P features
- **Mobile Apps**: iOS and Android clients
- **P2P Protocol**: Direct device-to-device file transfer
- **Sync Engine**: Real-time file synchronization
- **Plugin System**: Extensions for popular tools and services

## Architecture

### Backend Stack
- **Language**: Go 1.25
- **Framework**: Fiber v3 (high-performance web framework)
- **Database**: SQLite (with plans for PostgreSQL support)
- **Storage**: Local filesystem with atomic operations
- **Deployment**: Docker with multi-stage builds

### Security Design
```
User Password
├── Content KEK (End-to-End encryption)
│   └── File CEK (encrypts actual file content)
└── Metadata KEK (Server-side features)
    ├── Metadata encryption (filenames, dates)
    └── Thumbnail encryption (preview images)
```

### API Endpoints
- `POST /v1/user/register` - User registration
- `POST /v1/user/login` - Authentication
- `GET /v1/user/files` - List user files
- `POST /v1/file/` - File upload
- `GET /v1/file/:id` - File download
- `POST /v1/share/:id` - Create share link
- `GET /v1/share/:id` - Access shared file

## Getting Started

### Quick Start with Docker
```bash
# Clone the repository
git clone https://github.com/vizn3r/cloud.git
cd cloud

# Build and start services
./build.sh

# Access the application
# Web interface: http://localhost:8818
# API server: http://localhost:8808
```

### Manual Setup
```bash
# Backend setup
cd server
go mod download
go run main.go

# Configure settings
./generate-config.sh
```

## Configuration

The server uses `server.json` for configuration:
```json
{
  "webClient": {
    "host": "http://localhost",
    "port": 8008
  },
  "port": 8080
}
```

## Why Another Cloud Storage?

### Privacy Concerns
- Big Tech has access to your files
- Data mining and content scanning
- Vendor lock-in and service discontinuation
- Limited control over your own data

### Our Approach
- **Self-Hosted**: You control your data and infrastructure
- **Open Source**: Transparent, auditable code
- **Privacy-First**: Zero-knowledge encryption by design
- **Flexible**: Choose your security vs. convenience balance

## Contributing

This is a personal learning project, but contributions are welcome!

### Development Setup
1. Fork the repository
2. Set up the development environment
3. Make your changes
4. Submit a pull request

### Areas for Contribution
- Frontend development (Svelte/TypeScript)
- Mobile app development
- P2P protocol implementation
- Documentation and testing
- Security auditing

## Roadmap

### Phase 1: Core Functionality - Completed
- Basic file upload/download
- User authentication
- File sharing

### Phase 2: Security & UX - In Progress
- End-to-end encryption
- Responsive web interface  
- Enhanced sharing features

### Phase 3: Advanced Features - Planned
- P2P file transfer
- Desktop/mobile apps
- Real-time synchronization
- Plugin ecosystem

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Disclaimer

This is a work-in-progress personal project. Use at your own risk in production environments. Security features are still under development and have not been professionally audited.

---

*© 2025 Simon "vizn3r" Vizner - Built with love and lots of coffee*
