# nsdg1000t

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/masa23/nsdg1000t)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

## 概要
このリポジトリは、[SONY NSD-G1000T](https://www.nuro.jp/pdf/device/manual_NSD-G1000T.pdf) の情報を取得するためのgo言語のライブラリです。

## 使い方

exampleディレクトリにサンプルコードがあります。

## 対応API 
以下のAPIの情報取得のみ対応しています。

- stat
  - connectivity
  - 24g_enable
  - 5g_enable
  - wps
- support
  - system
  - wan
  - lan
  - ethernet_switch
  - dhcp_lease
  - lan_statistics
  - wlan_statistics