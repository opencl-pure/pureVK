package pureVK

import (
    "runtime"
    "testing"
    "unsafe"
)

func TestGPUMemoryAllocation(t *testing.T) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	err := LoadVulkanLibrary()
	if err != nil {
			t.Fatalf("Failed to load Vulkan library: %v", err)
	}
	defer UnloadVulkanLibrary()

	appName := []byte("TestApp\x00")

	defer func() {
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
			t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
	}
	defer VkDestroyInstance(instance, nil)

	var deviceCount uint32
	result, err = VkEnumeratePhysicalDevices(instance, &deviceCount, nil)
	if err != nil || deviceCount == 0 || result != VK_SUCCESS {
			t.Fatalf("VkEnumeratePhysicalDevices failed: %v, %s", err, HandleVkResult(result))
	}

	physicalDevices := make([]VulkanPhysicalDevice, deviceCount)
	result, err = VkEnumeratePhysicalDevices(instance, &deviceCount, &physicalDevices[0])
	if err != nil || result != VK_SUCCESS {
			t.Fatalf("VkEnumeratePhysicalDevices failed: %v, %s", err, HandleVkResult(result))
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
			t.Fatalf("VkCreateDevice failed: %v, %s", err, HandleVkResult(result))
	}
	defer VkDestroyDevice(device, nil)

	// Alokácia pamäte GPU
	var memoryRequirements VkMemoryRequirements
	bufferCreateInfo := VkBufferCreateInfo{
			SType: VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO,
			Size:  1024 * 1024 * 1024, // 1 GB
			Usage: VK_BUFFER_USAGE_STORAGE_BUFFER_BIT,
	}
	var buffer VulkanBuffer
	result, err = VkCreateBuffer(device, &bufferCreateInfo, nil, &buffer)
	if err != nil || result != VK_SUCCESS {
			t.Fatalf("VkCreateBuffer failed: %v, %s", err, HandleVkResult(result))
	}
	defer VkDestroyBuffer(device, buffer, nil)

	VkGetBufferMemoryRequirements(device, buffer, &memoryRequirements)
	memoryTypeIndex, err := findMemoryType(physicalDevice, memoryRequirements.MemoryTypeBits, VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT)
	if err != nil {
		t.Fatalf("findMemoryType failed: %v", err)
	}
	memoryAllocateInfo := VkMemoryAllocateInfo{
		SType:           VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO,
		AllocationSize:  memoryRequirements.Size,
		MemoryTypeIndex: memoryTypeIndex,
	}

	var memory VulkanDeviceMemory
	result, err = VkAllocateMemory(device, &memoryAllocateInfo, nil, &memory)
	if err != nil || result != VK_SUCCESS {
			t.Fatalf("VkAllocateMemory failed: %v, %s", err, HandleVkResult(result))
	}
	defer VkFreeMemory(device, memory, nil)

	t.Log("GPU memory allocation successful")
}