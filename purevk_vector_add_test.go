package pureVK

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"runtime"
	"testing"
	"unsafe"
)

func TestVectorAdd(t *testing.T) {
        // Inicializácia Vulkan
        instance, physicalDevice, device, queue, commandPool, _ := initializeVulkan(t)
        if instance == 0 || device == 0 {
                return
        }
        defer cleanupVulkan(instance, device, commandPool)

        // Načítanie SPIR-V shaderu
        shaderCode, err := loadSPIRVShader("vector_add.comp.spirv")
        if err != nil {
                t.Fatalf("Failed to load shader: %v", err)
        }

        // Vstupné a výstupné dáta
        vectorSize := 1024
        vectorA := make([]float32, vectorSize)
        vectorB := make([]float32, vectorSize)
        vectorResult := make([]float32, vectorSize)

        for i := 0; i < vectorSize; i++ {
                vectorA[i] = float32(i)
                vectorB[i] = float32(vectorSize - i)
        }

        // Vytvorenie bufferov
        bufferA, memoryA, err := createBufferAndMemory(device, physicalDevice, vectorA)
        if err != nil {
                t.Fatalf("Failed to create buffer A: %v", err)
        }
        defer VkDestroyBuffer(device, bufferA, nil)
        defer VkFreeMemory(device, memoryA, nil)

        bufferB, memoryB, err := createBufferAndMemory(device, physicalDevice, vectorB)
        if err != nil {
                t.Fatalf("Failed to create buffer B: %v", err)
        }
        defer VkDestroyBuffer(device, bufferB, nil)
        defer VkFreeMemory(device, memoryB, nil)

        bufferResult, memoryResult, err := createBufferAndMemory(device, physicalDevice, vectorResult)
        if err != nil {
                t.Fatalf("Failed to create buffer Result: %v", err)
        }
        defer VkDestroyBuffer(device, bufferResult, nil)
        defer VkFreeMemory(device, memoryResult, nil)

        // Vytvorenie shader modulu
        shaderModule, err := createShaderModule(device, shaderCode)
        if err != nil {
                t.Fatalf("Failed to create shader module: %v", err)
        }
        defer VkDestroyShaderModule(device, shaderModule, nil)

        // Descriptor set layout
        descriptorSetLayout, err := createDescriptorSetLayout(device)
        if err != nil {
                t.Fatalf("Failed to create descriptor set layout: %v", err)
        }
        defer VkDestroyDescriptorSetLayout(device, descriptorSetLayout, nil)

        // Pipeline layout
        pipelineLayout, err := createPipelineLayout(device, descriptorSetLayout)
        if err != nil {
                t.Fatalf("Failed to create pipeline layout: %v", err)
        }
        defer VkDestroyPipelineLayout(device, pipelineLayout, nil)

        // Pipeline
        pipeline, err := createComputePipeline(device, shaderModule, pipelineLayout)
        if err != nil {
                t.Fatalf("Failed to create pipeline: %v", err)
        }
        defer VkDestroyPipeline(device, pipeline, nil)

        // Descriptor pool
        descriptorPool, err := createDescriptorPool(device)
        if err != nil {
                t.Fatalf("Failed to create descriptor pool: %v", err)
        }
        defer VkDestroyDescriptorPool(device, descriptorPool, nil)

        // Descriptor sets
        descriptorSets, err := createDescriptorSets(device, descriptorPool, descriptorSetLayout, bufferA, bufferB, bufferResult)
        if err != nil {
                t.Fatalf("Failed to create descriptor sets: %v", err)
        }

        // Command buffer
        commandBuffer, err := createCommandBuffer(device, commandPool)
        if err != nil {
                t.Fatalf("Failed to create command buffer: %v", err)
        }

        // Nahrávanie príkazov do command bufferu
        err = recordCommandBuffer(commandBuffer, pipeline, pipelineLayout, descriptorSets, uint32(vectorSize/256))
        if err != nil {
                t.Fatalf("Failed to record command buffer: %v", err)
        }

        // Odoslanie command bufferu
        err = submitCommandBuffer(device, queue, commandBuffer)
        if err != nil {
                t.Fatalf("Failed to submit command buffer: %v", err)
        }

        // Kopírovanie výsledku z bufferu
        err = mapMemory(device, memoryResult, unsafe.Pointer(&vectorResult[0]), uint64(len(vectorResult)*4))
        if err != nil {
                t.Fatalf("Failed to map memory: %v", err)
        }

        // Overenie výsledku
        for i := 0; i < vectorSize; i++ {
                if vectorResult[i] != float32(vectorSize) {
                        t.Errorf("Result mismatch at index %d: expected %f, got %f", i, float32(vectorSize), vectorResult[i])
                }
        }

        t.Log("Vector addition successful")
        runtime.KeepAlive(appName)
}
var appName = []byte("TestApp\x00")
var mainName = []byte("main\x00")

func initializeVulkan(t *testing.T) (VulkanInstance, VulkanPhysicalDevice, VulkanDevice, VulkanQueue, VulkanCommandPool, error) {
    err := LoadVulkanLibrary()
    if err != nil {
            t.Fatalf("Failed to load Vulkan library: %v", err)
            return VulkanInstance(0), VulkanPhysicalDevice(0), VulkanDevice(0), VulkanQueue(0), VulkanCommandPool(0), err
    }

    appInfo := VkApplicationInfo{
            SType:              VK_STRUCTURE_TYPE_APPLICATION_INFO,
            PApplicationName:   (*byte)(unsafe.Pointer(&appName[0])),
            ApiVersion:         VK_MAKE_VERSION(1, 0, 0),
    }

    createInfo := VkInstanceCreateInfo{
            SType:              VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO,
            PApplicationInfo:   &appInfo,
    }

    var instance VulkanInstance
    result, err := VkCreateInstance(&createInfo, nil, &instance)
    if err != nil || result != VK_SUCCESS {
            t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
            return VulkanInstance(0), VulkanPhysicalDevice(0), VulkanDevice(0), VulkanQueue(0), VulkanCommandPool(0), err
    }

    var deviceCount uint32
    result, err = VkEnumeratePhysicalDevices(instance, &deviceCount, nil)
    if err != nil || deviceCount == 0 || result != VK_SUCCESS {
            t.Fatalf("VkEnumeratePhysicalDevices failed: %v, %s", err, HandleVkResult(result))
            return VulkanInstance(0), VulkanPhysicalDevice(0), VulkanDevice(0), VulkanQueue(0), VulkanCommandPool(0), err
    }

    physicalDevices := make([]VulkanPhysicalDevice, deviceCount)
    result, err = VkEnumeratePhysicalDevices(instance, &deviceCount, &physicalDevices[0])
    if err != nil || result != VK_SUCCESS {
            t.Fatalf("VkEnumeratePhysicalDevices failed: %v, %s", err, HandleVkResult(result))
            return VulkanInstance(0), VulkanPhysicalDevice(0), VulkanDevice(0), VulkanQueue(0), VulkanCommandPool(0), err
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
            return VulkanInstance(0), VulkanPhysicalDevice(0), VulkanDevice(0), VulkanQueue(0), VulkanCommandPool(0), err
    }

    var queue VulkanQueue
    VkGetDeviceQueue(device, 0, 0, &queue)

    commandPoolCreateInfo := VkCommandPoolCreateInfo{
            SType:            VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO,
            QueueFamilyIndex: 0,
    }

    var commandPool VulkanCommandPool
    result, err = VkCreateCommandPool(device, &commandPoolCreateInfo, nil, &commandPool)
    if err != nil || result != VK_SUCCESS {
            t.Fatalf("VkCreateCommandPool failed: %v, %s", err, HandleVkResult(result))
            return VulkanInstance(0), VulkanPhysicalDevice(0), VulkanDevice(0), VulkanQueue(0), VulkanCommandPool(0), err
    }

    return instance, physicalDevice, device, queue, commandPool, nil
}

func createBufferAndMemory(device VulkanDevice, physicalDevice VulkanPhysicalDevice, data []float32) (VulkanBuffer, VulkanDeviceMemory, error) {
    bufferCreateInfo := VkBufferCreateInfo{
            SType: VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO,
            Size:  uint64(len(data) * 4), // 4 bytes na float32
            Usage: VK_BUFFER_USAGE_STORAGE_BUFFER_BIT,
    }

    var buffer VulkanBuffer
    result, err := VkCreateBuffer(device, &bufferCreateInfo, nil, &buffer)
    if err != nil || result != VK_SUCCESS {
            return VulkanBuffer(0), VulkanDeviceMemory(0), fmt.Errorf("VkCreateBuffer failed: %v, %s", err, HandleVkResult(result))
    }

    var memoryRequirements VkMemoryRequirements
    VkGetBufferMemoryRequirements(device, buffer, &memoryRequirements)

    memoryTypeIndex, err := findMemoryType(physicalDevice, memoryRequirements.MemoryTypeBits, VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT|VK_MEMORY_PROPERTY_HOST_COHERENT_BIT)
    if err != nil {
            VkDestroyBuffer(device, buffer, nil)
            return VulkanBuffer(0), VulkanDeviceMemory(0), fmt.Errorf("findMemoryType failed: %v", err)
    }

    memoryAllocateInfo := VkMemoryAllocateInfo{
            SType:           VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO,
            AllocationSize:  memoryRequirements.Size,
            MemoryTypeIndex: memoryTypeIndex,
    }

    var memory VulkanDeviceMemory
    result, err = VkAllocateMemory(device, &memoryAllocateInfo, nil, &memory)
    if err != nil || result != VK_SUCCESS {
            VkDestroyBuffer(device, buffer, nil)
            return VulkanBuffer(0), VulkanDeviceMemory(0), fmt.Errorf("VkAllocateMemory failed: %v, %s", err, HandleVkResult(result))
    }

    result, err = VkBindBufferMemory(device, buffer, memory, 0)
    if err != nil || result != VK_SUCCESS {
            VkDestroyBuffer(device, buffer, nil)
            VkFreeMemory(device, memory, nil)
            return VulkanBuffer(0), VulkanDeviceMemory(0), fmt.Errorf("VkBindBufferMemory failed: %v, %s", err, HandleVkResult(result))
    }

    var dataPtr unsafe.Pointer
    result, err = VkMapMemory(device, memory, 0, memoryRequirements.Size, 0, &dataPtr)
    if err != nil || result != VK_SUCCESS {
            VkDestroyBuffer(device, buffer, nil)
            VkFreeMemory(device, memory, nil)
            return VulkanBuffer(0), VulkanDeviceMemory(0), fmt.Errorf("VkMapMemory failed: %v, %s", err, HandleVkResult(result))
    }

    dataBytes := unsafe.Slice((*byte)(dataPtr), len(data)*4)
    for i, val := range data {
            binary.LittleEndian.PutUint32(dataBytes[i*4:], math.Float32bits(val))
    }

    VkUnmapMemory(device, memory)

    return buffer, memory, nil
}

func cleanupVulkan(instance VulkanInstance, device VulkanDevice, commandPool VulkanCommandPool) {
    VkDestroyCommandPool(device, commandPool, nil)
    VkDestroyDevice(device, nil)
    VkDestroyInstance(instance, nil)
    UnloadVulkanLibrary()
}

func loadSPIRVShader(filename string) ([]byte, error) {
    file, err := os.ReadFile(filename)
    if err != nil {
            return nil, err
    }
    return file, nil
}

func createShaderModule(device VulkanDevice, shaderCode []byte) (VulkanShaderModule, error) {
    shaderModuleCreateInfo := VkShaderModuleCreateInfo{
            SType:    VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO,
            CodeSize: uintptr(len(shaderCode)),
            PCode:    (*uint32)(unsafe.Pointer(&shaderCode[0])),
    }

    var shaderModule VulkanShaderModule
    result, err := VkCreateShaderModule(device, &shaderModuleCreateInfo, nil, &shaderModule)
    if err != nil || result != VK_SUCCESS {
            return VulkanShaderModule(0), fmt.Errorf("VkCreateShaderModule failed: %v, %s", err, HandleVkResult(result))
    }
    return shaderModule, nil
}

func createDescriptorSetLayout(device VulkanDevice) (VulkanDescriptorSetLayout, error) {
    descriptorSetLayoutBinding := []VkDescriptorSetLayoutBinding{
            {
                    Binding:         0,
                    DescriptorType:  VK_DESCRIPTOR_TYPE_STORAGE_BUFFER,
                    DescriptorCount: 1,
                    StageFlags:      VK_SHADER_STAGE_COMPUTE_BIT,
            },
            {
                    Binding:         1,
                    DescriptorType:  VK_DESCRIPTOR_TYPE_STORAGE_BUFFER,
                    DescriptorCount: 1,
                    StageFlags:      VK_SHADER_STAGE_COMPUTE_BIT,
            },
            {
                    Binding:         2,
                    DescriptorType:  VK_DESCRIPTOR_TYPE_STORAGE_BUFFER,
                    DescriptorCount: 1,
                    StageFlags:      VK_SHADER_STAGE_COMPUTE_BIT,
            },
    }

    descriptorSetLayoutCreateInfo := VkDescriptorSetLayoutCreateInfo{
            SType:        VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO,
            BindingCount: uint32(len(descriptorSetLayoutBinding)),
            PBindings:    &descriptorSetLayoutBinding[0],
    }

    var descriptorSetLayout VulkanDescriptorSetLayout
    result, err := VkCreateDescriptorSetLayout(device, &descriptorSetLayoutCreateInfo, nil, &descriptorSetLayout)
    if err != nil || result != VK_SUCCESS {
            return VulkanDescriptorSetLayout(0), fmt.Errorf("VkCreateDescriptorSetLayout failed: %v, %s", err, HandleVkResult(result))
    }
    return descriptorSetLayout, nil
}

func createPipelineLayout(device VulkanDevice, descriptorSetLayout VulkanDescriptorSetLayout) (VulkanPipelineLayout, error) {
    pipelineLayoutCreateInfo := VkPipelineLayoutCreateInfo{
            SType:          VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO,
            SetLayoutCount: 1,
            PSetLayouts:    &descriptorSetLayout,
    }

    var pipelineLayout VulkanPipelineLayout
    result, err := VkCreatePipelineLayout(device, &pipelineLayoutCreateInfo, nil, &pipelineLayout)
    if err != nil || result != VK_SUCCESS {
            return VulkanPipelineLayout(0), fmt.Errorf("VkCreatePipelineLayout failed: %v, %s", err, HandleVkResult(result))
    }
    return pipelineLayout, nil
}


func createComputePipeline(device VulkanDevice, shaderModule VulkanShaderModule, pipelineLayout VulkanPipelineLayout) (VulkanPipeline, error) {
    stageCreateInfo := VkPipelineShaderStageCreateInfo{
            SType:  VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO,
            Stage:  VK_SHADER_STAGE_COMPUTE_BIT,
            Module: shaderModule,
            PName:  (*byte)(unsafe.Pointer(&(mainName)[0])),
    }

    computePipelineCreateInfo := VkComputePipelineCreateInfo{
            SType:  VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO,
            Stage:  stageCreateInfo,
            Layout: pipelineLayout,
    }

    var pipeline VulkanPipeline
    result, err := VkCreateComputePipelines(device, VulkanPipelineCache(0), 1, &computePipelineCreateInfo, nil, &pipeline)
    if err != nil || result != VK_SUCCESS {
            return VulkanPipeline(0), fmt.Errorf("VkCreateComputePipelines failed: %v, %s", err, HandleVkResult(result))
    }
    return pipeline, nil
}
func createDescriptorPool(device VulkanDevice) (VulkanDescriptorPool, error) {
    poolSizes := []VkDescriptorPoolSize{
            {
                    Type:            VK_DESCRIPTOR_TYPE_STORAGE_BUFFER,
                    DescriptorCount: 3,
            },
    }

    descriptorPoolCreateInfo := VkDescriptorPoolCreateInfo{
            SType:         VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO,
            MaxSets:       1,
            PoolSizeCount: uint32(len(poolSizes)),
            PPoolSizes:    &poolSizes[0],
    }

    var descriptorPool VulkanDescriptorPool
    result, err := VkCreateDescriptorPool(device, &descriptorPoolCreateInfo, nil, &descriptorPool)
    if err != nil || result != VK_SUCCESS {
            return VulkanDescriptorPool(0), fmt.Errorf("VkCreateDescriptorPool failed: %v, %s", err, HandleVkResult(result))
    }
    return descriptorPool, nil
}
func createDescriptorSets(device VulkanDevice, descriptorPool VulkanDescriptorPool, descriptorSetLayout VulkanDescriptorSetLayout, bufferA, bufferB, bufferResult VulkanBuffer) ([]VulkanDescriptorSet, error) {
    descriptorSetAllocateInfo := VkDescriptorSetAllocateInfo{
            SType:              VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO,
            DescriptorPool:     descriptorPool,
            DescriptorSetCount: 1,
            PSetLayouts:        &descriptorSetLayout,
    }

    var descriptorSet VulkanDescriptorSet
    result, err := VkAllocateDescriptorSets(device, &descriptorSetAllocateInfo, &descriptorSet)
    if err != nil || result != VK_SUCCESS {
            return nil, fmt.Errorf("VkAllocateDescriptorSets failed: %v, %s", err, HandleVkResult(result))
    }

    writeDescriptorSets := []VkWriteDescriptorSet{
            {
                    SType:           VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET,
                    DstSet:          descriptorSet,
                    DstBinding:      0,
                    DescriptorCount: 1,
                    DescriptorType:  VK_DESCRIPTOR_TYPE_STORAGE_BUFFER,
                    PBufferInfo: &VkDescriptorBufferInfo{
                            Buffer: bufferA,
                            Range:  VK_WHOLE_SIZE,
                    },
            },
            {
                    SType:           VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET,
                    DstSet:          descriptorSet,
                    DstBinding:      1,
                    DescriptorCount: 1,
                    DescriptorType:  VK_DESCRIPTOR_TYPE_STORAGE_BUFFER,
                    PBufferInfo: &VkDescriptorBufferInfo{
                            Buffer: bufferB,
                            Range:  VK_WHOLE_SIZE,
                    },
            },
            {
                    SType:           VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET,
                    DstSet:          descriptorSet,
                    DstBinding:      2,
                    DescriptorCount: 1,
                    DescriptorType:  VK_DESCRIPTOR_TYPE_STORAGE_BUFFER,
                    PBufferInfo: &VkDescriptorBufferInfo{
                            Buffer: bufferResult,
                            Range:  VK_WHOLE_SIZE,
                    },
            },
    }

    VkUpdateDescriptorSets(device, uint32(len(writeDescriptorSets)), &writeDescriptorSets[0], 0, nil)

    return []VulkanDescriptorSet{descriptorSet}, nil
}

func recordCommandBuffer(commandBuffer VulkanCommandBuffer, pipeline VulkanPipeline, pipelineLayout VulkanPipelineLayout, descriptorSets []VulkanDescriptorSet, groupCountX uint32) error {
    beginInfo := VkCommandBufferBeginInfo{
        SType: VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO,
        Flags: VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT,
    }

    result, err := VkBeginCommandBuffer(commandBuffer, &beginInfo)
    if err != nil || result != VK_SUCCESS {
        return fmt.Errorf("VkBeginCommandBuffer failed: %s", HandleVkResult(result))
    }

    // Overenie VkResult pre VkCmdBindPipeline
    result, err = VkCmdBindPipeline(commandBuffer, VK_PIPELINE_BIND_POINT_COMPUTE, pipeline)
    if err != nil || result < VK_SUCCESS {
        return fmt.Errorf("VkBeginCommandBuffer failed: %s", HandleVkResult(result))
    }

    // Overenie VkResult pre VkCmdBindDescriptorSets
    result, err = VkCmdBindDescriptorSets(commandBuffer, VK_PIPELINE_BIND_POINT_COMPUTE, pipelineLayout, 0, 1, &descriptorSets[0], 0, nil)
    if err != nil || result != VK_SUCCESS {
        return fmt.Errorf("VkBeginCommandBuffer failed: %s", HandleVkResult(result))
    }

    fmt.Printf("groupCountX: %d\n", groupCountX)

    // Overenie VkResult pre VkCmdDispatch a spracovanie chyby
    result, err = VkCmdDispatch(commandBuffer, groupCountX, 1, 1)
    if err != nil || result != VK_SUCCESS {
        if err != nil {
            return fmt.Errorf("VkCmdDispatch failed: %s, %v", HandleVkResult(result), err)
        }
        return fmt.Errorf("VkCmdDispatch failed: %s", HandleVkResult(result))
    }

    result, err = VkEndCommandBuffer(commandBuffer)
    if err != nil || result != VK_SUCCESS {
        return fmt.Errorf("VkEndCommandBuffer failed: %s", HandleVkResult(result))
    }

    return nil
}

func submitCommandBuffer(device VulkanDevice, queue VulkanQueue, commandBuffer VulkanCommandBuffer) error {
    submitInfo := VkSubmitInfo{
            SType:                VK_STRUCTURE_TYPE_SUBMIT_INFO,
            CommandBufferCount:   1,
            PCommandBuffers:      &commandBuffer,
    }

    result, err := VkQueueSubmit(queue, 1, &submitInfo, VulkanFence(0))
    if err != nil || result != VK_SUCCESS {
            return fmt.Errorf("VkQueueSubmit failed: %v, %s", err, HandleVkResult(result))
    }

    VkQueueWaitIdle(queue)

    return nil
}
func mapMemory(device VulkanDevice, memory VulkanDeviceMemory, data unsafe.Pointer, size uint64) error {
    var mappedData unsafe.Pointer
    result, err := VkMapMemory(device, memory, 0, size, 0, &mappedData)
    if err != nil || result != VK_SUCCESS {
            return fmt.Errorf("VkMapMemory failed: %v, %s", err, HandleVkResult(result))
    }

    src := unsafe.Slice((*byte)(mappedData), int(size))
    dst := unsafe.Slice((*byte)(data), int(size))
    copy(dst, src)

    VkUnmapMemory(device, memory)
    return nil
}

func createCommandBuffer(device VulkanDevice, commandPool VulkanCommandPool) (VulkanCommandBuffer, error) {
    allocateInfo := VkCommandBufferAllocateInfo{
        SType:              VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO,
        CommandPool:        commandPool,
        Level:              VK_COMMAND_BUFFER_LEVEL_PRIMARY,
        CommandBufferCount: 1,
    }

    var commandBuffer VulkanCommandBuffer
    result, err := VkAllocateCommandBuffers(device, &allocateInfo, &commandBuffer)
    if err != nil || result != VK_SUCCESS {
        return VulkanCommandBuffer(0), fmt.Errorf("VkAllocateCommandBuffers failed: %s", HandleVkResult(result))
    }

    return commandBuffer, nil
}