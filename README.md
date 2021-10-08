# Golang lib for camera capture

[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/camera.svg)](https://pkg.go.dev/github.com/bububa/camera)
[![Go](https://github.com/bububa/camera/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/camera/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/camera/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/camera/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/camera.svg)](https://github.com/bububa/camera)
[![GoReportCard](https://goreportcard.com/badge/github.com/bububa/camera)](https://goreportcard.com/report/github.com/bububa/camera)
[![GitHub license](https://img.shields.io/github/license/bububa/camera.svg)](https://github.com/bububa/camera/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/camera.svg)](https://GitHub.com/bububa/camera/releases/)

## Demo

![demo screen capture](https://github.com/bububa/camera/blob/main/demo.gif?raw=true)

## Install

go get -u github.com/bububa/camera

## Requirements

- [libjpeg-turbo](https://www.libjpeg-turbo.org/) (use `-tags jpeg` to build without `CGo`)
- On Linux/RPi native Go [V4L](https://github.com/korandiz/v4l) implementation is used to capture images.

## Build tags

- `cv4` - build with `OpenCV` 4.x ([gocv](https://gocv.io/x/gocv))
- `jpeg` - build with native Go `image/jpeg` instead of `libjpeg-turbo`
- `android` - build for Android
