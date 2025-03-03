package pureVK

import (
	"runtime"
	"testing"
	"unsafe"
)

func TestVkCreateBuffer(t *testing.T) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	err := LoadVulkanLibrary()
	if err != nil {
		t.Fatalf("Failed to load Vulkan library: %v", err)
	}
	defer UnloadVulkanLibrary()

	appName := []byte("TestApp\x00")

	defer func ()  {
		runtime.KeepAlive(appName)
	}()

	var appInfo VkApplicationInfo
	appInfo.SType = VK_STRUCTURE_TYPE_APPLICATION_INFO
	appInfo.PApplicationName = (*byte)(unsafe.Pointer(&appName[0]))
	appInfo.ApiVersion = VK_API_VERSION_1_0

	var createInfo VkInstanceCreateInfo
	createInfo.SType = VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO
	createInfo.PApplicationInfo = &appInfo

	var instance VulkanInstance
	result, err := VkCreateInstance(&createInfo, nil, &instance)
	if err != nil || result != VK_SUCCESS {
		t.Fatalf("VkCreateInstance failed: %v", err)
	}
	defer VkDestroyInstance(instance, nil)

	var deviceCount uint32
	result, err = VkEnumerateDevices(instance, &deviceCount, nil)
	if err != nil || deviceCount == 0 || result != VK_SUCCESS {
		t.Fatalf("No Vulkan devices found: %v", err)
	}

	physicalDevices := make([]VulkanPhysicalDevice, deviceCount)
	result, err = VkEnumerateDevices(instance, &deviceCount, &physicalDevices[0])
	if err != nil  || result != VK_SUCCESS {
		t.Fatalf("VkEnumerateDevices failed: %v", err)
	}
	physicalDevice := physicalDevices[0]

	var queueCreateInfo VkDeviceQueueCreateInfo
	queueCreateInfo.SType = VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
	queueCreateInfo.QueueFamilyIndex = 0
	queueCreateInfo.QueueCount = 1
	var queuePriority float32 = 1.0
	queueCreateInfo.PQueuePriorities = &queuePriority

	var deviceCreateInfo VkDeviceCreateInfo
	deviceCreateInfo.SType = VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
	deviceCreateInfo.QueueCreateInfoCount = 1
	deviceCreateInfo.PQueueCreateInfos = unsafe.Pointer(&queueCreateInfo)

	var device VulkanDevice
	result, err = VkCreateDevice(physicalDevice, &deviceCreateInfo, nil, &device)
	if err != nil || result != VK_SUCCESS {
		t.Fatalf("VkCreateDevice failed: %v", err)
	}
	defer VkDestroyDevice(device, nil)

	var buffer VulkanBuffer
	var bufferInfo VkBufferCreateInfo
	bufferInfo.SType = VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
	bufferInfo.Size = 1024
	bufferInfo.Usage = VK_BUFFER_USAGE_VERTEX_BUFFER_BIT

	result, err = VkCreateBuffer(device, &bufferInfo, nil, &buffer)
	if err != nil || result != VK_SUCCESS {
		t.Fatalf("VkCreateBuffer failed: %v", err)
	}

	VkDestroyBuffer(device, buffer, nil)
}