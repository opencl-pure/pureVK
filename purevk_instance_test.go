package pureVK

import (
    "runtime"
    "testing"
    "unsafe"
)

func TestVkCreateInstance(t *testing.T) {
    runtime.LockOSThread()
    defer runtime.UnlockOSThread()

    err := LoadVulkanLibrary()
    if err != nil {
        t.Fatalf("Failed to load Vulkan library: %v", err)
    }
    defer UnloadVulkanLibrary()

    appName := []byte("TestApp\x00")
    engineName := []byte("NoEngine\x00")
    defer func() {
        runtime.KeepAlive(appName)
        runtime.KeepAlive(engineName)
    }()

    var appInfo VkApplicationInfo
    appInfo.SType = VK_STRUCTURE_TYPE_APPLICATION_INFO
    appInfo.PApplicationName = (*byte)(unsafe.Pointer(&appName[0]))
    appInfo.ApplicationVersion = VK_MAKE_VERSION(1, 0, 0)
    appInfo.PEngineName = (*byte)(unsafe.Pointer(&engineName[0]))
    appInfo.EngineVersion = VK_MAKE_VERSION(1, 0, 0)
    appInfo.ApiVersion = VK_API_VERSION_1_0

    var createInfo VkInstanceCreateInfo
    createInfo.SType = VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO
    createInfo.PApplicationInfo = &appInfo

    var instance VulkanInstance
    result, err := VkCreateInstance(&createInfo, nil, &instance)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v", err)
    }

    VkDestroyInstance(instance, nil)
}