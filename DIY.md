# APIpsum

[![Go Version](https://img.shields.io/github/go-mod/go-version/gofiber/fiber)](https://golang.org/doc/go1.23) [![Fiber Version](https://img.shields.io/badge/fiber-v2.32.0-blue)](https://gofiber.io)

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [File Structure](#file-structure)
- [Development](#development)
- [Deployment](#deployment)
- [License](#license)

## Introduction

**APIpsum** is a RESTful API that generates JSON objects based on a provided schema. The API offers two primary endpoints: a GET request to verify availability and a POST request to generate data. It is built with [Fiber](https://gofiber.io/), a fast, lightweight web framework written in [Go](https://go.dev/).

## Requirements

- Go >= 1.16
- Node.js >= 14.x
- NPM >= 6.x (or Yarn/PNPM)
- TailwindCSS >= 3.x
- PostCSS >= 8.x

## Installation

### Clone this repository

```bash
git clone https://github.com/aabboudi/apipsum.git
cd apipsum
```

### Install Go dependencies

```bash
go mod tidy
```

### Install Node.js dependencies

```bash
cd ./static/
npm install
cd ..
```

### Run PostCSS

```bash
npm run watch:css
```

### Run the Go Fiber application

```bash
go run main.go
```

## File Structure

```
├── controllers/          # Business logic
│   └── controllers.go
├── docs/                 # Swagger API documentation
│   ├── doc.go
│   ├── docs.json
│   └── docs.yaml
├── middleware/           # Logging and validation mechanisms
│   ├── logger.go
│   └── validator.go
├── routes/               # API route definitions
│   └── routes.go
├── static/               # Static files
│   ├── css/
│   │   └── input.css
│   ├── package.json
│   ├── postcss.config.js
│   └── tailwind.config.js
├── utils/                # Utility functions and packages
│   ├── utils.go
│   └── letters/
│       └── letters.go
├── views/                # HTML templates
│   └── index.html
├── .air.toml             # Air live-reload configuration
├── .gitignore            # Git ignore file
├── DIY.md                # This file
├── go.mod                # Go module definition
├── go.sum                # Go module dependencies
└── main.go               # Application entry point
```
