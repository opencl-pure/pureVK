package pureVK

import (
    "errors"
    "os"
)

// Funkcia na načítanie SPIR-V kódu zo súboru
func ReadSPIRVFromFile(filename string) ([]uint32, error) {
    file, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    if len(file)%4 != 0 {
        return nil, errors.New("SPIR-V file size is not a multiple of 4")
    }

    shaderCode := make([]uint32, len(file)/4)
    for i := 0; i < len(shaderCode); i++ {
        shaderCode[i] = uint32(file[i*4]) | uint32(file[i*4+1])<<8 | uint32(file[i*4+2])<<16 | uint32(file[i*4+3])<<24
    }

    return shaderCode, nil
}