package pureVK

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
)

var (
	vulkanLib uintptr
)

// Vulkan funkcie
var (
    vkCreateInstance                  func(*VkInstanceCreateInfo, *VkAllocationCallbacks, *VulkanInstance) VkResult
    vkDestroyInstance                 func(VulkanInstance, *VkAllocationCallbacks)
    vkEnumeratePhysicalDevices        func(VulkanInstance, *uint32, *VulkanPhysicalDevice) VkResult

    vkCreateDevice                    func(VulkanPhysicalDevice, *VkDeviceCreateInfo, *VkAllocationCallbacks, *VulkanDevice) VkResult
    vkDestroyDevice                   func(VulkanDevice, *VkAllocationCallbacks)
    vkCreateBuffer                    func(VulkanDevice, *VkBufferCreateInfo, *VkAllocationCallbacks, *VulkanBuffer) VkResult
    vkDestroyBuffer                   func(VulkanDevice, VulkanBuffer, *VkAllocationCallbacks)
    vkCreateImage                     func(VulkanDevice, *VkImageCreateInfo, *VkAllocationCallbacks, *VulkanImage) VkResult
    vkDestroyImage                    func(VulkanDevice, VulkanImage, *VkAllocationCallbacks)
    vkCreateShaderModule              func(VulkanDevice, *VkShaderModuleCreateInfo, *VkAllocationCallbacks, *VulkanShaderModule) VkResult

    vkCreateCommandPool               func(VulkanDevice, *VkCommandPoolCreateInfo, *VkAllocationCallbacks, *VulkanCommandPool) VkResult
    vkAllocateCommandBuffers          func(VulkanDevice, *VkCommandBufferAllocateInfo, *VulkanCommandBuffer) VkResult
    vkBeginCommandBuffer              func(VulkanCommandBuffer, *VkCommandBufferBeginInfo) VkResult
    vkEndCommandBuffer                func(VulkanCommandBuffer) VkResult
    vkQueueSubmit                     func(VulkanQueue, uint32, *VkSubmitInfo, VulkanFence) VkResult
    vkQueueWaitIdle                   func(VulkanQueue) VkResult

    vkCreateDescriptorSetLayout       func(VulkanDevice, *VkDescriptorSetLayoutCreateInfo, *VkAllocationCallbacks, *VulkanDescriptorSetLayout) VkResult
    vkAllocateDescriptorSets          func(VulkanDevice, *VkDescriptorSetAllocateInfo, *VulkanDescriptorSet) VkResult
    vkUpdateDescriptorSets            func(VulkanDevice, uint32, *VkWriteDescriptorSet, uint32, *VkCopyDescriptorSet)

    vkCreatePipelineLayout            func(VulkanDevice, *VkPipelineLayoutCreateInfo, *VkAllocationCallbacks, *VulkanPipelineLayout) VkResult
    vkCreateComputePipelines          func(VulkanDevice, VulkanPipelineCache, uint32, *VkComputePipelineCreateInfo, *VkAllocationCallbacks, *VulkanPipeline) VkResult

    vkCreateFence                     func(VulkanDevice, *VkFenceCreateInfo, *VkAllocationCallbacks, *VulkanFence) VkResult
    vkWaitForFences                   func(VulkanDevice, uint32, *VulkanFence, uint32, uint64) VkResult

    vkGetPhysicalDeviceProperties     func(VulkanPhysicalDevice, *VkPhysicalDeviceProperties)
    vkCreateDescriptorPool            func(VulkanDevice, *VkDescriptorPoolCreateInfo, *VkAllocationCallbacks, *VulkanDescriptorPool) VkResult

    vkDestroyCommandPool              func(VulkanDevice, VulkanCommandPool, *VkAllocationCallbacks)
    vkDestroyShaderModule             func(VulkanDevice, VulkanShaderModule, *VkAllocationCallbacks)

    vkDestroyDescriptorSetLayout      func(VulkanDevice, VulkanDescriptorSetLayout, *VkAllocationCallbacks)
    vkDestroyPipelineLayout           func(VulkanDevice, VulkanPipelineLayout, *VkAllocationCallbacks)

    vkDestroyPipeline                 func(VulkanDevice, VulkanPipeline, *VkAllocationCallbacks)
    vkDestroyDescriptorPool           func(VulkanDevice, VulkanDescriptorPool, *VkAllocationCallbacks)

    vkCmdBindPipeline                 func(VulkanCommandBuffer, VkPipelineBindPoint, VulkanPipeline) VkResult // Upravené
    vkCmdBindDescriptorSets           func(VulkanCommandBuffer, VkPipelineBindPoint, VulkanPipelineLayout, uint32, uint32, *VulkanDescriptorSet, uint32, *uint32) VkResult // Upravené

    vkCmdDispatch                     func(VulkanCommandBuffer, uint32, uint32, uint32) VkResult // Upravené

    vkGetDeviceQueue                  func(VulkanDevice, uint32, uint32, *VulkanQueue) VkResult // Upravené
    vkMapMemory                       func(VulkanDevice, VulkanDeviceMemory, uint64, uint64, VkMemoryMapFlags, *unsafe.Pointer) VkResult
    vkUnmapMemory                     func(VulkanDevice, VulkanDeviceMemory) VkResult // Upravené

    vkGetBufferMemoryRequirements     func(VulkanDevice, VulkanBuffer, *VkMemoryRequirements) VkResult // Upravené
    vkAllocateMemory                  func(VulkanDevice, *VkMemoryAllocateInfo, *VkAllocationCallbacks, *VulkanDeviceMemory) VkResult
    vkFreeMemory                      func(VulkanDevice, VulkanDeviceMemory, *VkAllocationCallbacks)
    vkBindBufferMemory                func(VulkanDevice, VulkanBuffer, VulkanDeviceMemory, uint64) VkResult

    vkGetPhysicalDeviceMemoryProperties func(VulkanPhysicalDevice, *VkPhysicalDeviceMemoryProperties) VkResult // Upravené

	vkEnumerateDevices   func(VulkanInstance, *uint32, *VulkanPhysicalDevice) VkResult
)

// Registrácia Vulkan funkcií
func registerVulkanFuncs() error {
	err := registerFuncWithoutPanic(&vkCreateInstance, vulkanLib, "vkCreateInstance", nil)
	err = registerFuncWithoutPanic(&vkDestroyInstance, vulkanLib, "vkDestroyInstance", err)
	err = registerFuncWithoutPanic(&vkEnumerateDevices, vulkanLib, "vkEnumeratePhysicalDevices", err)
	err = registerFuncWithoutPanic(&vkCreateDevice, vulkanLib, "vkCreateDevice", err)
	err = registerFuncWithoutPanic(&vkDestroyDevice, vulkanLib, "vkDestroyDevice", err)
	err = registerFuncWithoutPanic(&vkCreateBuffer, vulkanLib, "vkCreateBuffer", err)
	err = registerFuncWithoutPanic(&vkDestroyBuffer, vulkanLib, "vkDestroyBuffer", err)
	err = registerFuncWithoutPanic(&vkCreateImage, vulkanLib, "vkCreateImage", err)
	err = registerFuncWithoutPanic(&vkDestroyImage, vulkanLib, "vkDestroyImage", err)
	err = registerFuncWithoutPanic(&vkCreateShaderModule, vulkanLib, "vkCreateShaderModule", err)

	err = registerFuncWithoutPanic(&vkCreateCommandPool, vulkanLib, "vkCreateCommandPool", err)
	err = registerFuncWithoutPanic(&vkAllocateCommandBuffers, vulkanLib, "vkAllocateCommandBuffers", err)
	err = registerFuncWithoutPanic(&vkBeginCommandBuffer, vulkanLib, "vkBeginCommandBuffer", err)
	err = registerFuncWithoutPanic(&vkEndCommandBuffer, vulkanLib, "vkEndCommandBuffer", err)
	err = registerFuncWithoutPanic(&vkQueueSubmit, vulkanLib, "vkQueueSubmit", err)
	err = registerFuncWithoutPanic(&vkQueueWaitIdle, vulkanLib, "vkQueueWaitIdle", err)

	err = registerFuncWithoutPanic(&vkCreateDescriptorSetLayout, vulkanLib, "vkCreateDescriptorSetLayout", err)
	err = registerFuncWithoutPanic(&vkAllocateDescriptorSets, vulkanLib, "vkAllocateDescriptorSets", err)
	err = registerFuncWithoutPanic(&vkUpdateDescriptorSets, vulkanLib, "vkUpdateDescriptorSets", err)
	err = registerFuncWithoutPanic(&vkCreatePipelineLayout, vulkanLib, "vkCreatePipelineLayout", err)
	err = registerFuncWithoutPanic(&vkCreateComputePipelines, vulkanLib, "vkCreateComputePipelines", err)

	err = registerFuncWithoutPanic(&vkCreateFence, vulkanLib, "vkCreateFence", err)
	err = registerFuncWithoutPanic(&vkWaitForFences, vulkanLib, "vkWaitForFences", err)

	err = registerFuncWithoutPanic(&vkDestroyCommandPool, vulkanLib, "vkDestroyCommandPool", err)
    err = registerFuncWithoutPanic(&vkDestroyShaderModule, vulkanLib, "vkDestroyShaderModule", err)
    
	err = registerFuncWithoutPanic(&vkDestroyDescriptorSetLayout, vulkanLib, "vkDestroyDescriptorSetLayout", err)
    err = registerFuncWithoutPanic(&vkDestroyPipelineLayout, vulkanLib, "vkDestroyPipelineLayout", err)
    
	err = registerFuncWithoutPanic(&vkDestroyPipeline, vulkanLib, "vkDestroyPipeline", err)
    err = registerFuncWithoutPanic(&vkDestroyDescriptorPool, vulkanLib, "vkDestroyDescriptorPool", err)
    
	err = registerFuncWithoutPanic(&vkCmdBindPipeline, vulkanLib, "vkCmdBindPipeline", err)
    err = registerFuncWithoutPanic(&vkCmdBindDescriptorSets, vulkanLib, "vkCmdBindDescriptorSets", err)
	
	err = registerFuncWithoutPanic(&vkCmdDispatch, vulkanLib, "vkCmdDispatch", err)
	
	err = registerFuncWithoutPanic(&vkGetDeviceQueue, vulkanLib, "vkGetDeviceQueue", err)
    err = registerFuncWithoutPanic(&vkMapMemory, vulkanLib, "vkMapMemory", err)
    err = registerFuncWithoutPanic(&vkUnmapMemory, vulkanLib, "vkUnmapMemory", err)
    
	err = registerFuncWithoutPanic(&vkGetBufferMemoryRequirements, vulkanLib, "vkGetBufferMemoryRequirements", err)
    err = registerFuncWithoutPanic(&vkAllocateMemory, vulkanLib, "vkAllocateMemory", err)
    err = registerFuncWithoutPanic(&vkFreeMemory, vulkanLib, "vkFreeMemory", err)
    err = registerFuncWithoutPanic(&vkBindBufferMemory, vulkanLib, "vkBindBufferMemory", err)
    
	err = registerFuncWithoutPanic(&vkEnumeratePhysicalDevices, vulkanLib, "vkEnumeratePhysicalDevices", err)

	err = registerFuncWithoutPanic(&vkGetPhysicalDeviceMemoryProperties, vulkanLib, "vkGetPhysicalDeviceMemoryProperties", err)


	err = registerFuncWithoutPanic(&vkCreateDescriptorPool, vulkanLib, "vkCreateDescriptorPool", err)
	return err
}

func registerFuncWithoutPanic(fptr interface{}, handle uintptr, name string, error0 error) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = ErrJoin(error0, errors.New(fmt.Sprint(r)))
		} else {
			e = error0
		}
	}()
	purego.RegisterLibFunc(fptr, handle, name)
	return
}

func myDlOpen() (uintptr, error) {
	// Zoznam možných ciest
	paths := []string{
		"libvulkan.so",
		"libvulkan.so.1",
		"/usr/lib/x86_64-linux-gnu/libvulkan.so",
		"/usr/local/lib/libvulkan.so",
		"/lib/x86_64-linux-gnu/libvulkan.so",
		"/opt/vulkan-sdk/lib/libvulkan.so",
	}

	var err error
	for _, path := range paths {
		lib, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			// fmt.Printf(" Vulkan knižnica načítaná: %s\n", path)
			return lib, nil
		}
	}

	return 0, fmt.Errorf("nepodarilo sa načítať Vulkan knižnicu: %v", err)
}

func VK_MAKE_VERSION(major, minor, patch uint32) uint32 {
	return (major << 22) | (minor << 12) | patch
}

// 📌 Upravená funkcia na načítanie Vulkan knižnice
func LoadVulkanLibrary() error {
	var err error
	vulkanLib, err = myDlOpen()
	if err != nil {
		return fmt.Errorf("failed to load Vulkan library: %v", err)
	}

	// Registrácia Vulkan funkcií
	if err := registerVulkanFuncs(); err != nil {
		return err
	}
	return nil
}

func UnloadVulkanLibrary() {
    if vulkanLib != 0 {
        purego.Dlclose(vulkanLib)
        vulkanLib = 0
    }
}

func GetBufferData[T any](data []T) *BufferData {
	size := unsafe.Sizeof(data[0])
	return &BufferData{
		TypeSize: size,
		DataSize: uintptr(len(data)) * size,
		Pointer:  unsafe.Pointer(&data[0]),
	}
}

func ErrJoin(e1, e2 error) error {
	if e1 != nil && e2 != nil {
		return errors.New(e1.Error() + ";\n" + e2.Error())
	}
	if e1 != nil {
		return e1
	}
	return e2
}

func VkCreateInstance(instanceInfo *VkInstanceCreateInfo, callback *VkAllocationCallbacks, instance *VulkanInstance) (VkResult, error){
	if vkCreateInstance == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkCreateInstance = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkCreateInstance(instanceInfo, callback, instance), nil
}

func VkDestroyInstance(instance VulkanInstance, callback *VkAllocationCallbacks) error {
	if vkDestroyInstance == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkDestroyInstance = nil
			return err
		}
	}
	vkDestroyInstance(instance, callback)
	return nil
}

func VkEnumerateDevices(instance VulkanInstance, number *uint32, device *VulkanPhysicalDevice) (VkResult, error){
	if vkEnumerateDevices == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkEnumerateDevices = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkEnumerateDevices(instance, number, device), nil
}

func VkCreateDevice(device0 VulkanPhysicalDevice, info *VkDeviceCreateInfo, callback *VkAllocationCallbacks, device *VulkanDevice) (VkResult, error){
	if vkCreateDevice == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkCreateDevice = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkCreateDevice(device0, info, callback, device), nil
}

func VkDestroyDevice(device VulkanDevice, callback *VkAllocationCallbacks) error {
	if vkDestroyDevice == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkDestroyDevice = nil
			return err
		}
	}
	vkDestroyDevice(device, callback)
	return nil
}

func VkCreateBuffer(device VulkanDevice, info *VkBufferCreateInfo, callback *VkAllocationCallbacks, buffer *VulkanBuffer)  (VkResult, error){
	if vkCreateBuffer == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkCreateBuffer = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkCreateBuffer(device, info, callback, buffer), nil
}

func VkDestroyBuffer(device VulkanDevice, buffer VulkanBuffer, callback *VkAllocationCallbacks) (error){
	if vkDestroyBuffer == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkDestroyBuffer = nil
			return err
		}
	}
	vkDestroyBuffer(device, buffer, callback)
	return nil
}


func VkCreateImage(device VulkanDevice, info *VkImageCreateInfo, callback *VkAllocationCallbacks, image *VulkanImage)  (VkResult, error){
	if vkCreateImage == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkCreateImage = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkCreateImage(device, info, callback, image), nil
}

func VkDestroyImage(device VulkanDevice, image VulkanImage, callback *VkAllocationCallbacks) error{
	if vkDestroyImage == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkDestroyImage = nil
			return err
		}
	}
	vkDestroyImage(device, image, callback)
	return nil
}

func VkCreateShaderModule(device VulkanDevice, info *VkShaderModuleCreateInfo, callback *VkAllocationCallbacks, module *VulkanShaderModule) (VkResult, error){
	if vkCreateShaderModule == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkCreateShaderModule = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkCreateShaderModule(device, info, callback, module), nil
}

func VkCreateCommandPool(device VulkanDevice, info *VkCommandPoolCreateInfo, callback *VkAllocationCallbacks, comandPool *VulkanCommandPool) (VkResult, error){
	if vkCreateCommandPool == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkCreateCommandPool = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkCreateCommandPool(device, info, callback, comandPool), nil
}

func VkAllocateCommandBuffers(device VulkanDevice, aInfo *VkCommandBufferAllocateInfo, cBuffer *VulkanCommandBuffer)  (VkResult, error){
	if vkAllocateCommandBuffers == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkAllocateCommandBuffers = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkAllocateCommandBuffers(device, aInfo, cBuffer), nil
}

func VkBeginCommandBuffer(cBuffer VulkanCommandBuffer, bInfo *VkCommandBufferBeginInfo)  (VkResult, error){
	if vkBeginCommandBuffer == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkBeginCommandBuffer = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkBeginCommandBuffer(cBuffer, bInfo), nil
}

func VkEndCommandBuffer(cBuffer VulkanCommandBuffer)  (VkResult, error){
	if vkEndCommandBuffer == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkEndCommandBuffer = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkEndCommandBuffer(cBuffer), nil
}

func VkQueueSubmit(queue VulkanQueue, f uint32, sInfo *VkSubmitInfo, fence VulkanFence)  (VkResult, error){
	if vkQueueSubmit == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkQueueSubmit = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkQueueSubmit(queue, f, sInfo, fence), nil
}

func VkQueueWaitIdle(queue VulkanQueue)  (VkResult, error){
	if vkQueueWaitIdle == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkQueueWaitIdle = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkQueueWaitIdle(queue), nil
}

func VkCreateDescriptorSetLayout(device VulkanDevice, dInfo *VkDescriptorSetLayoutCreateInfo, callback *VkAllocationCallbacks, dLayout *VulkanDescriptorSetLayout)  (VkResult, error){
	if vkCreateDescriptorSetLayout == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkCreateDescriptorSetLayout = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkCreateDescriptorSetLayout(device, dInfo, callback, dLayout), nil
}

func VkAllocateDescriptorSets(device VulkanDevice, aInfo *VkDescriptorSetAllocateInfo, dSet *VulkanDescriptorSet)  (VkResult, error){
	if vkAllocateDescriptorSets == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkAllocateDescriptorSets = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkAllocateDescriptorSets(device, aInfo, dSet), nil
}

func VkUpdateDescriptorSets(device VulkanDevice, f uint32, dWSet *VkWriteDescriptorSet, f2 uint32, dCSet *VkCopyDescriptorSet) error{
	if vkUpdateDescriptorSets == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkUpdateDescriptorSets = nil
			return err
		}
	}
	vkUpdateDescriptorSets(device, f, dWSet, f2, dCSet)
	return nil
}

func VkCreatePipelineLayout(device VulkanDevice, cInfo *VkPipelineLayoutCreateInfo, callback *VkAllocationCallbacks, pipelineLayout *VulkanPipelineLayout)  (VkResult, error){
	if vkCreatePipelineLayout == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkCreatePipelineLayout = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkCreatePipelineLayout(device, cInfo, callback, pipelineLayout), nil
}

func VkCreateComputePipelines(device VulkanDevice, cache VulkanPipelineCache, f uint32, pCInfo *VkComputePipelineCreateInfo, callback *VkAllocationCallbacks, pipeline *VulkanPipeline)  (VkResult, error){
	if vkCreateComputePipelines == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkCreateComputePipelines = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkCreateComputePipelines(device, cache, f, pCInfo, callback, pipeline), nil
}

func VkCreateFence(device VulkanDevice, fInfo *VkFenceCreateInfo, callback *VkAllocationCallbacks, fence *VulkanFence)  (VkResult, error){
	if vkCreateFence == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkCreateFence = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkCreateFence(device, fInfo, callback, fence), nil
}

func VkWaitForFences(device VulkanDevice, f uint32, fence *VulkanFence, f2 uint32, f3 uint64)  (VkResult, error){
	if vkWaitForFences == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkWaitForFences = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkWaitForFences(device, f, fence, f2, f3), nil
}

func VkGetPhysicalDeviceProperties(physicDevice VulkanPhysicalDevice, deviceProperties *VkPhysicalDeviceProperties) error{
	if vkGetPhysicalDeviceProperties == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkGetPhysicalDeviceProperties = nil
			return err
		}
	}
	vkGetPhysicalDeviceProperties(physicDevice, deviceProperties)
	return nil
}

func VkCreateDescriptorPool(device VulkanDevice, info *VkDescriptorPoolCreateInfo, callback *VkAllocationCallbacks, dPool *VulkanDescriptorPool) (VkResult, error){
	if vkCreateDescriptorPool == nil {
		err := LoadVulkanLibrary()
		if err != nil {
			vkCreateDescriptorPool = nil
			return VK_ERROR_UNKNOWN, err
		}
	}
	return vkCreateDescriptorPool(device, info, callback, dPool), nil
}
func VkDestroyCommandPool(device VulkanDevice, commandPool VulkanCommandPool, callback *VkAllocationCallbacks) error {
        if vkDestroyCommandPool == nil {
                err := LoadVulkanLibrary()
                if err != nil {
					vkDestroyCommandPool = nil
                    return err
                }
        }
        vkDestroyCommandPool(device, commandPool, callback)
        return nil
}

func VkDestroyShaderModule(device VulkanDevice, shaderModule VulkanShaderModule, callback *VkAllocationCallbacks) error {
        if vkDestroyShaderModule == nil {
                err := LoadVulkanLibrary()
                if err != nil {
					vkDestroyShaderModule = nil
                    return err
                }
        }
        vkDestroyShaderModule(device, shaderModule, callback)
        return nil
}

func VkDestroyDescriptorSetLayout(device VulkanDevice, descriptorSetLayout VulkanDescriptorSetLayout, callback *VkAllocationCallbacks) error {
    if vkDestroyDescriptorSetLayout == nil {
        err := LoadVulkanLibrary()
        if err != nil {
			vkDestroyShaderModule = nil
            return err
        }
    }
    vkDestroyDescriptorSetLayout(device, descriptorSetLayout, callback)
    return nil
}

func VkDestroyPipelineLayout(device VulkanDevice, pipelineLayout VulkanPipelineLayout, callback *VkAllocationCallbacks) error {
    if vkDestroyPipelineLayout == nil {
        err := LoadVulkanLibrary()
        if err != nil {
			vkDestroyPipelineLayout = nil
            return err
        }
    }
    vkDestroyPipelineLayout(device, pipelineLayout, callback)
    return nil
}

func VkDestroyPipeline(device VulkanDevice, pipeline VulkanPipeline, callback *VkAllocationCallbacks) error {
    if vkDestroyPipeline == nil {
        err := LoadVulkanLibrary()
        if err != nil {
            return err
        }
    }
    vkDestroyPipeline(device, pipeline, callback)
    return nil
}

func VkDestroyDescriptorPool(device VulkanDevice, descriptorPool VulkanDescriptorPool, callback *VkAllocationCallbacks) error {
    if vkDestroyDescriptorPool == nil{
        err := LoadVulkanLibrary()
        if err != nil {
			vkDestroyDescriptorPool = nil
            return err
        }
    }
    vkDestroyDescriptorPool(device, descriptorPool, callback)
    return nil
}

func VkMapMemory(device VulkanDevice, memory VulkanDeviceMemory, offset uint64, size uint64, flags VkMemoryMapFlags, ppData *unsafe.Pointer) (VkResult, error) {
    if vkMapMemory == nil {
        err := LoadVulkanLibrary()
        if err != nil {
			vkMapMemory = nil
            return VK_ERROR_UNKNOWN, err
        }
    }
    result := vkMapMemory(device, memory, offset, size, flags, ppData)
    return result, nil
}

func VkAllocateMemory(device VulkanDevice, allocateInfo *VkMemoryAllocateInfo, allocator *VkAllocationCallbacks, memory *VulkanDeviceMemory) (VkResult, error) {
    if vkAllocateMemory == nil {
        err := LoadVulkanLibrary()
        if err != nil {
			vkAllocateMemory = nil
            return VK_ERROR_UNKNOWN, err
        }
    }
    result := vkAllocateMemory(device, allocateInfo, allocator, memory)
    return result, nil
}

func VkFreeMemory(device VulkanDevice, memory VulkanDeviceMemory, allocator *VkAllocationCallbacks) (error){
    if vkFreeMemory == nil {
        err := LoadVulkanLibrary()
        if err != nil {
			vkFreeMemory = nil
            return err
        }
    }
    vkFreeMemory(device, memory, allocator)
	return nil
}

func VkBindBufferMemory(device VulkanDevice, buffer VulkanBuffer, memory VulkanDeviceMemory, memoryOffset uint64) (VkResult, error) {
    if vkBindBufferMemory == nil {
        err := LoadVulkanLibrary()
        if err != nil {
			vkBindBufferMemory = nil
            return VK_ERROR_UNKNOWN, err
        }
    }
    result := vkBindBufferMemory(device, buffer, memory, memoryOffset)
    return result, nil
}

func VkEnumeratePhysicalDevices(instance VulkanInstance, physicalDeviceCount *uint32, physicalDevices *VulkanPhysicalDevice) (VkResult, error) {
    if vkEnumeratePhysicalDevices == nil {
        err := LoadVulkanLibrary()
        if err != nil {
			vkEnumeratePhysicalDevices = nil
            return VK_ERROR_UNKNOWN, err
        }
    }
    result := vkEnumeratePhysicalDevices(instance, physicalDeviceCount, physicalDevices)
    return result, nil

}

func VkCmdBindPipeline(commandBuffer VulkanCommandBuffer, pipelineBindPoint VkPipelineBindPoint, pipeline VulkanPipeline) (VkResult, error) {
    if vkCmdBindPipeline == nil {
        err := LoadVulkanLibrary()
        if err != nil {
            vkCmdBindPipeline = nil
            return VK_ERROR_UNKNOWN, err
        }
    }
    result := vkCmdBindPipeline(commandBuffer, pipelineBindPoint, pipeline)
    return result, nil
}

func VkCmdBindDescriptorSets(commandBuffer VulkanCommandBuffer, pipelineBindPoint VkPipelineBindPoint, pipelineLayout VulkanPipelineLayout, firstSet uint32, descriptorSetCount uint32, descriptorSets *VulkanDescriptorSet, dynamicOffsetCount uint32, dynamicOffsets *uint32) (VkResult, error) {
    if vkCmdBindDescriptorSets == nil {
        err := LoadVulkanLibrary()
        if err != nil {
            vkCmdBindDescriptorSets = nil
            return VK_ERROR_UNKNOWN, err
        }
    }
    result := vkCmdBindDescriptorSets(commandBuffer, pipelineBindPoint, pipelineLayout, firstSet, descriptorSetCount, descriptorSets, dynamicOffsetCount, dynamicOffsets)
    return result, nil
}

func VkCmdDispatch(commandBuffer VulkanCommandBuffer, groupCountX uint32, groupCountY uint32, groupCountZ uint32) (VkResult, error) {
    if vkCmdDispatch == nil {
        err := LoadVulkanLibrary()
        if err != nil {
            vkCmdDispatch = nil
            return VK_ERROR_UNKNOWN, err
        }
    }
    result := vkCmdDispatch(commandBuffer, groupCountX, groupCountY, groupCountZ)
    return result, nil
}

func VkGetDeviceQueue(device VulkanDevice, queueFamilyIndex uint32, queueIndex uint32, queue *VulkanQueue) (VkResult, error) {
    if vkGetDeviceQueue == nil {
        err := LoadVulkanLibrary()
        if err != nil {
            vkGetDeviceQueue = nil
            return VK_ERROR_UNKNOWN, err
        }
    }
    result := vkGetDeviceQueue(device, queueFamilyIndex, queueIndex, queue)
    return result, nil
}

func VkUnmapMemory(device VulkanDevice, memory VulkanDeviceMemory) (VkResult, error) {
    if vkUnmapMemory == nil {
        err := LoadVulkanLibrary()
        if err != nil {
            vkUnmapMemory = nil
            return VK_ERROR_UNKNOWN, err
        }
    }
    result := vkUnmapMemory(device, memory)
    return result, nil
}

func VkGetBufferMemoryRequirements(device VulkanDevice, buffer VulkanBuffer, memoryRequirements *VkMemoryRequirements) (VkResult, error) {
    if vkGetBufferMemoryRequirements == nil {
        err := LoadVulkanLibrary()
        if err != nil {
            vkGetBufferMemoryRequirements = nil
            return VK_ERROR_UNKNOWN, err
        }
    }
    result := vkGetBufferMemoryRequirements(device, buffer, memoryRequirements)
    return result, nil
}

func VkGetPhysicalDeviceMemoryProperties(physicalDevice VulkanPhysicalDevice, memoryProperties *VkPhysicalDeviceMemoryProperties) (VkResult, error) {
    if vkGetPhysicalDeviceMemoryProperties == nil {
        err := LoadVulkanLibrary()
        if err != nil {
            vkGetPhysicalDeviceMemoryProperties = nil
            return VK_ERROR_UNKNOWN, err
        }
    }
    result := vkGetPhysicalDeviceMemoryProperties(physicalDevice, memoryProperties)
    return result, nil
}



func findMemoryType(physicalDevice VulkanPhysicalDevice, typeFilter uint32, properties VkMemoryPropertyFlags) (uint32, error) {
    var memoryProperties VkPhysicalDeviceMemoryProperties
	fmt.Printf("findMemoryType called with typeFilter: %d, properties: %d\n", typeFilter, properties) // Pridajte toto
    res, err := VkGetPhysicalDeviceMemoryProperties(physicalDevice, &memoryProperties)
	if err != nil || res < VK_SUCCESS{
		return 0, ErrJoin(err, errors.New(HandleVkResult(res)))
	}

    for i := uint32(0); i < memoryProperties.MemoryTypeCount; i++ {
        if (typeFilter & (1 << i)) != 0 && (memoryProperties.MemoryTypes[i].PropertyFlags&properties) == properties {
            return i, nil
        }
    }

    return 0, errors.New("not found")
}

