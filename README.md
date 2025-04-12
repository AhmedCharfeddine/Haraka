# Haraka 🔠🪄

A fast and simple Arabic transliteration tool built with Go.

## Overview

Haraka is a (currently WIP) lightweight tool that helps convert Latin script into Arabic script. 

Based on Latin input, it provides real-time transliteration suggestions for Arabic words. Haraka enables you to type in Arabic with speed ⚡ 

## Who's this for?

If your main keyboard is in Latin and you:  

🔹 Type faster with Latin layouts  
🔹 Use a keyboard with missing Arabic labels  
🔹 Are too lazy to swap keyboard language

Then Haraka is for you ✨

## Build

The project assumes you have GO installed on your system. The build steps are:

1. Building the project binaries:
```
$ go build
```
2. Running the system-wide installation:
```
$ go intall
```
3. Testing the CLI tool:
```
$ haraka --help
$ haraka map salaam
سلام
```

## Contributions

You're always welcome to contribute by opening issues for suggestions or submitting pull requests!

## Roadmap

🔹 Support more arabic spelling rules (Tied Ta' `ة` and dha tuchel `ظ`)    
🔹 Add support for common Arabic words  
🔹 Implement a GUI version  
🔹 System-wide integration for real-time typing  
