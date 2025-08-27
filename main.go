package main

import (
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "fmt"
    "hash"
    "io"
    "log"
    "os"
    "path/filepath"
)

func main() {
    showBanner()
    filename := getFilePath()
    validateFile(filename)
    showOptions()

    var option int
    fmt.Print("Elige una opción: ")
    fmt.Scanf("%d", &option)
    fmt.Println("")

    switch option {
    case 1:
        calculateHash(filename, md5.New(), "MD5")
    case 2:
        calculateHash(filename, sha1.New(), "SHA1")
    case 3:
        calculateHash(filename, sha256.New(), "SHA256")
    case 4:
        calculateHash(filename, sha512.New(), "SHA512")
    case 5:
        calculateAllHashes(filename)
    default:
        fmt.Println("Opción inválida. Por favor, elige entre 1 y 5.")
    }
}

func showBanner() {
    banner := `
  ____       _   _           _               
 / ___| ___ | | | | __ _ ___| |__   ___ _ __ 
| |  _ / _ \| |_| |/ _| / __| '_ \ / _ \ '__|
| |_| | (_) |  _  | (_| \__ \ | | |  __/ |   
 \____|\___/|_| |_|\__,_|___/_| |_|\___|_|   
`
    fmt.Println(banner)
}

func getFilePath() string {
    var filename string
    fmt.Print("Ingresa la ruta del archivo: ")
    fmt.Scan(&filename)
    return filename
}

func validateFile(filename string) {
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        log.Fatalf("El archivo no existe: %s\n", filename)
    }
}

func showOptions() {
    fmt.Println("\n== Menú de Opciones ==")
    fmt.Println("=======================")
    fmt.Println("1 -> MD5")
    fmt.Println("2 -> SHA1")
    fmt.Println("3 -> SHA256")
    fmt.Println("4 -> SHA512")
    fmt.Println("5 -> Todos")
    fmt.Println("=======================\n")
}

func calculateHash(filename string, hasher hash.Hash, label string) {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    if _, err := io.Copy(hasher, file); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s -> %x (%s)\n", label, hasher.Sum(nil), filepath.Base(filename))
}

func calculateAllHashes(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Crear múltiples hashers
    md5Hash := md5.New()
    sha1Hash := sha1.New()
    sha256Hash := sha256.New()
    sha512Hash := sha512.New()

    // Crear un writer múltiple
    multiWriter := io.MultiWriter(md5Hash, sha1Hash, sha256Hash, sha512Hash)

    if _, err := io.Copy(multiWriter, file); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Archivo: %s\n\n", filepath.Base(filename))
    fmt.Printf("MD5    -> %x\n", md5Hash.Sum(nil))
    fmt.Printf("SHA1   -> %x\n", sha1Hash.Sum(nil))
    fmt.Printf("SHA256 -> %x\n", sha256Hash.Sum(nil))
    fmt.Printf("SHA512 -> %x\n", sha512Hash.Sum(nil))
}
