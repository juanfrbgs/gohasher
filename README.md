# GoHasher

A tool that allows you to calculate different types of **hashes** of a file.

## Hash Types

* md5
* sha1
* sha256
* sha512
* all - md5,sha1,sha256,sha512

## Usage

```bash
go build -o gohasher main.go && ./gohasher

--
or 
--

go run main.go
```

## Example

```bash
  ____       _   _           _               
 / ___| ___ | | | | __ _ ___| |__   ___ _ __ 
| |  _ / _ \| |_| |/ _| / __| '_ \ / _ \ '__|
| |_| | (_) |  _  | (_| \__ \ | | |  __/ |   
 \____|\___/|_| |_|\__,_|___/_| |_|\___|_|   

Enter the file path: README.md

== Options Menu ==
=======================
1 -> MD5
2 -> SHA1
3 -> SHA256
4 -> SHA512
5 -> All
=======================

Choose an option: 5

File: README.md

MD5    -> cf36775472ffd0c37ae0b1428fbdf98e
SHA1   -> d372c68657ffdc9fb355b42f7ff78a7f37dd9f4d
SHA256 -> b8dc0498d39963e4545509b52195b75223ae6810d0d4d8bf48d800d971129ddd
SHA512 -> 5b7444cf0ab11810df64ce4182f4629936567c1908176b8843a8f1b399376471b398dd5f9a8f6e1c3b5404a8096135370e1cf1a9109310e999905ca84ce474c7
```

