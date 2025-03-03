package pureVK

import (
    "runtime"
    "testing"
)

func TestLoadVulkanLibrary(t *testing.T) {
    runtime.LockOSThread()
    defer runtime.UnlockOSThread()

    err := LoadVulkanLibrary()
    if err != nil {
        t.Fatalf("Failed to load Vulkan library: %v", err)
    }
    UnloadVulkanLibrary()
}