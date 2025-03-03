package pureVK

import (
    "runtime"
    "testing"
    "unsafe"
)

func TestVectorAdd(t *testing.T) {
    runtime.LockOSThread()
    defer runtime.UnlockOSThread()

    err := LoadVulkanLibrary()
    if err != nil {
        t.Fatalf("Failed to load Vulkan library: %v", err)
    }
    defer UnloadVulkanLibrary()

    appName := []byte("TestApp\x00")
    mainName := []byte("main\x00")

    defer func() {
        runtime.KeepAlive(appName)
        runtime.KeepAlive(mainName)
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
    result, err = VkEnumerateDevices(instance, &deviceCount, nil)
    if err != nil || deviceCount == 0 || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }

    physicalDevices := make([]VulkanPhysicalDevice, deviceCount)
    result, err = VkEnumerateDevices(instance, &deviceCount, &physicalDevices[0])
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
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
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }
    defer VkDestroyDevice(device, nil)

    // 1. VkCreateBuffer (input buffer)
    var inputBuffer VulkanBuffer
    inputBufferCreateInfo := VkBufferCreateInfo{
        SType: VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO,
        Size:  uint64(unsafe.Sizeof(Vec3{}) * 2),
        Usage: VK_BUFFER_USAGE_STORAGE_BUFFER_BIT,
    }
    result, err = VkCreateBuffer(device, &inputBufferCreateInfo, nil, &inputBuffer)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }
    defer VkDestroyBuffer(device, inputBuffer, nil)

    // 2. VkCreateBuffer (output buffer)
    var outputBuffer VulkanBuffer
    outputBufferCreateInfo := VkBufferCreateInfo{
        SType: VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO,
        Size:  uint64(unsafe.Sizeof(Vec3{})),
        Usage: VK_BUFFER_USAGE_STORAGE_BUFFER_BIT,
    }
    result, err = VkCreateBuffer(device, &outputBufferCreateInfo, nil, &outputBuffer)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }
    defer VkDestroyBuffer(device, outputBuffer, nil)

    // 3. VkCreateCommandPool
    var commandPool VulkanCommandPool
    commandPoolCreateInfo := VkCommandPoolCreateInfo{
        SType:            VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO,
        QueueFamilyIndex: 0,
    }
    result, err = VkCreateCommandPool(device, &commandPoolCreateInfo, nil, &commandPool)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }
    defer VkDestroyCommandPool(device, commandPool, nil)

    // 4. VkAllocateCommandBuffers
    var commandBuffer VulkanCommandBuffer
    commandBufferAllocateInfo := VkCommandBufferAllocateInfo{
        SType:              VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO,
        CommandPool:        commandPool,
        Level:              VK_COMMAND_BUFFER_LEVEL_PRIMARY,
        CommandBufferCount: 1,
    }
    result, err = VkAllocateCommandBuffers(device, &commandBufferAllocateInfo, &commandBuffer)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }

    // 5. VkBeginCommandBuffer
    commandBufferBeginInfo := VkCommandBufferBeginInfo{
        SType: VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO,
    }
    result, err = VkBeginCommandBuffer(commandBuffer, &commandBufferBeginInfo)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }

    // 6. VkCreateShaderModule
    shaderCode, err := ReadSPIRVFromFile("vector_add.comp.spv")
    if err != nil {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }
    var shaderModule VulkanShaderModule
    shaderModuleCreateInfo := VkShaderModuleCreateInfo{
        SType:    VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO,
        CodeSize: uintptr(len(shaderCode) * 4),
        PCode:    &shaderCode[0],
    }
    result, err = VkCreateShaderModule(device, &shaderModuleCreateInfo, nil, &shaderModule)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }
    defer VkDestroyShaderModule(device, shaderModule, nil)

    // 7. VkCreateDescriptorSetLayout
    var descriptorSetLayout VulkanDescriptorSetLayout
    descriptorSetLayoutBindings := []VkDescriptorSetLayoutBinding{
        {
            Binding:            0,
            DescriptorType:     VK_DESCRIPTOR_TYPE_STORAGE_BUFFER,
            DescriptorCount:    1,
            StageFlags:         VK_SHADER_STAGE_COMPUTE_BIT,
        },
        {
            Binding:            1,
            DescriptorType:     VK_DESCRIPTOR_TYPE_STORAGE_BUFFER,
            DescriptorCount:    1,
            StageFlags:         VK_SHADER_STAGE_COMPUTE_BIT,
        },
    }
    descriptorSetLayoutCreateInfo := VkDescriptorSetLayoutCreateInfo{
        SType:        VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO,
        BindingCount: uint32(len(descriptorSetLayoutBindings)),
        PBindings:    &descriptorSetLayoutBindings[0],
    }
    result, err = VkCreateDescriptorSetLayout(device, &descriptorSetLayoutCreateInfo, nil, &descriptorSetLayout)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }
    defer VkDestroyDescriptorSetLayout(device, descriptorSetLayout, nil)

    // 8. VkCreatePipelineLayout
    var pipelineLayout VulkanPipelineLayout
    pipelineLayoutCreateInfo := VkPipelineLayoutCreateInfo{
        SType:          VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO,
        SetLayoutCount: 1,PSetLayouts:    &descriptorSetLayout,
    }
    result, err = VkCreatePipelineLayout(device, &pipelineLayoutCreateInfo, nil, &pipelineLayout)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }

    defer VkDestroyPipelineLayout(device, pipelineLayout, nil)

    // 9. VkCreateComputePipelines
    var pipeline VulkanPipeline
    pipelineShaderStageCreateInfo := VkPipelineShaderStageCreateInfo{
        SType:  VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO,
        Stage:  VK_SHADER_STAGE_COMPUTE_BIT,
        Module: shaderModule,
        PName:  (*byte)(unsafe.Pointer(&appName[0])),
    }
    computePipelineCreateInfo := VkComputePipelineCreateInfo{
        SType:  VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO,
        Stage:  pipelineShaderStageCreateInfo,
        Layout: pipelineLayout,
    }
    result, err = VkCreateComputePipelines(device, 0, 1, &computePipelineCreateInfo, nil, &pipeline)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }
    defer VkDestroyPipeline(device, pipeline, nil)

    // 10. VkCreateDescriptorPool
    var descriptorPool VulkanDescriptorPool
    descriptorPoolSizes := []VkDescriptorPoolSize{
        {
            Type:            VK_DESCRIPTOR_TYPE_STORAGE_BUFFER,
            DescriptorCount: 2,
        },
    }
    descriptorPoolCreateInfo := VkDescriptorPoolCreateInfo{
        SType:         VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO,
        MaxSets:       1,
        PoolSizeCount: uint32(len(descriptorPoolSizes)),
        PPoolSizes:    &descriptorPoolSizes[0],
    }
    result, err = VkCreateDescriptorPool(device, &descriptorPoolCreateInfo, nil, &descriptorPool)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }
    defer VkDestroyDescriptorPool(device, descriptorPool, nil)

    // 11. VkAllocateDescriptorSets
    var descriptorSet VulkanDescriptorSet
    descriptorSetAllocateInfo := VkDescriptorSetAllocateInfo{
        SType:              VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO,
        DescriptorPool:     descriptorPool,
        DescriptorSetCount: 1,
        PSetLayouts:        &descriptorSetLayout,
    }
    result, err = VkAllocateDescriptorSets(device, &descriptorSetAllocateInfo, &descriptorSet)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }

    // 12. VkUpdateDescriptorSets
    writeDescriptorSets := []VkWriteDescriptorSet{
        {
            SType:           VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET,
            DstSet:          descriptorSet,
            DstBinding:      0,
            DescriptorCount: 1,
            DescriptorType:  VK_DESCRIPTOR_TYPE_STORAGE_BUFFER,
            PBufferInfo: &VkDescriptorBufferInfo{
                Buffer: inputBuffer,
                Range:  uint64(unsafe.Sizeof(Vec3{}) * 2),
            },
        },
        {
            SType:           VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET,
            DstSet:          descriptorSet,
            DstBinding:      1,
            DescriptorCount: 1,
            DescriptorType:  VK_DESCRIPTOR_TYPE_STORAGE_BUFFER,
            PBufferInfo: &VkDescriptorBufferInfo{
                Buffer: outputBuffer,
                Range:  uint64(unsafe.Sizeof(Vec3{})),
            },
        },
    }
    VkUpdateDescriptorSets(device, uint32(len(writeDescriptorSets)), &writeDescriptorSets[0], 0, nil)

    // 13. VkCmdBindPipeline
    VkCmdBindPipeline(commandBuffer, VK_PIPELINE_BIND_POINT_COMPUTE, pipeline)

    // 14. VkCmdBindDescriptorSets
    VkCmdBindDescriptorSets(commandBuffer, VK_PIPELINE_BIND_POINT_COMPUTE, pipelineLayout, 0, 1, &descriptorSet, 0, nil)

    // 15. VkCmdDispatch
    VkCmdDispatch(commandBuffer, 1, 1, 1)

    // 16. VkEndCommandBuffer
    result, err = VkEndCommandBuffer(commandBuffer)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }

    // 17. VkQueueSubmit
    submitInfo := VkSubmitInfo{
        SType:                VK_STRUCTURE_TYPE_SUBMIT_INFO,
        CommandBufferCount: 1,
        PCommandBuffers:    &commandBuffer,
    }

    var queue VulkanQueue
    vkGetDeviceQueue(device, 0, 0, &queue)

    result, err = VkQueueSubmit(queue, 0, &submitInfo, 0)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }

    // 18. VkQueueWaitIdle
    result, err = VkQueueWaitIdle(queue)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }

    // 19. Čítanie dát z output bufferu
    var memoryRequirements VkMemoryRequirements
    vkGetBufferMemoryRequirements(device, outputBuffer, &memoryRequirements)
    type0, err := findMemoryType(physicalDevice, memoryRequirements.MemoryTypeBits, VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT|VK_MEMORY_PROPERTY_HOST_COHERENT_BIT)
    if err != nil {
        t.Fatalf("VkAllocateMemory failed: %v", err)
    }

    memoryAllocateInfo := VkMemoryAllocateInfo{
        SType:           VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO,
        AllocationSize:  memoryRequirements.Size,
        MemoryTypeIndex: type0,
    }

    var memory VulkanDeviceMemory
    result, err = VkAllocateMemory(device, &memoryAllocateInfo, nil, &memory)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }
    defer VkFreeMemory(device, memory, nil)

    result, err = VkBindBufferMemory(device, outputBuffer, memory, 0)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }

    var mappedMemory unsafe.Pointer
    result, err = VkMapMemory(device, memory, 0, uint64(unsafe.Sizeof(Vec3{})), 0, &mappedMemory)
    if err != nil || result != VK_SUCCESS {
        t.Fatalf("VkCreateInstance failed: %v, %s", err, HandleVkResult(result))
    }
    defer VkUnmapMemory(device, memory)

    resultVec := *(*Vec3)(mappedMemory)

    // 20. Kontrola výsledkov
    expectedVec := Vec3{5.0, 7.0, 9.0}
    if resultVec != expectedVec {
        t.Errorf("Vector addition failed, expected: %v, got: %v", expectedVec, resultVec)
    }
}