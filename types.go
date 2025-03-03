package pureVK

import (
	"unsafe"
)

// Vulkan základné typy
type VkResult int32
type VulkanInstance uintptr
type VulkanDevice uintptr
type VulkanBuffer uintptr
type VulkanImage uintptr
type VulkanShaderModule uintptr
type VulkanPipeline uintptr
type VulkanDescriptorSetLayout uintptr
type VulkanDescriptorSet uintptr
type VulkanPipelineLayout uintptr
type VulkanPipelineCache uintptr
type VulkanDescriptorPool uintptr
type VulkanPhysicalDevice uintptr
type VulkanCommandPool uintptr
type VulkanCommandBuffer uintptr
type VulkanQueue uintptr
type VulkanFence uintptr
type VkCommandBufferLevel uint32
type VkStructureType uint32
type VkPipelineShaderStageCreateFlags uint32
type VkShaderStageFlagBits uint32
type VkPhysicalDeviceType uint32
type VkDescriptorPoolCreateFlags uint32
type VkDescriptorType uint32
type VkBool32 uint32
type VkPipelineBindPoint uint32
type VulkanRenderPass uintptr
type VulkanFramebuffer uintptr
type VkQueryControlFlags uint32
type VkQueryPipelineStatisticFlags uint32
type VkPipelineLayoutCreateFlags uint32
type VulkanSampler uintptr
type VkDescriptorSetLayoutCreateFlags uint32
type VkPipelineCreateFlags uint32
type VulkanBufferView uintptr
type VulkanImageView uintptr
type VkImageLayout uint32
type VulkanSemaphore uintptr
type VkPipelineStageFlags uint32
type VulkanDeviceMemory uintptr
type VkMemoryMapFlags uint32
type VkMemoryPropertyFlags uint32

type Vec3 struct {
    X float32
    Y float32
    Z float32
}

type VkMemoryRequirements struct {
    Size           uint64
    Alignment      uint64
    MemoryTypeBits uint32
}
type VkPhysicalDeviceMemoryProperties struct {
    MemoryTypeCount uint32
    MemoryTypes     [32]VkMemoryType
    MemoryHeapCount uint32
    MemoryHeaps     [16]VkMemoryHeap
}

type VkMemoryType struct {
    PropertyFlags VkMemoryPropertyFlags
    HeapIndex     uint32
}

type VkMemoryHeap struct {
    Size  uint64
    Flags VkMemoryHeapFlags
}

type VkMemoryHeapFlags uint32

type VkMemoryAllocateInfo struct {
    SType           VkStructureType
    PNext           unsafe.Pointer
    AllocationSize  uint64
    MemoryTypeIndex uint32
}

type VkDescriptorImageInfo struct {
    Sampler     VulkanSampler
    ImageView   VulkanImageView
    ImageLayout VkImageLayout
}

type VkSubmitInfo struct {
    SType                VkStructureType
    PNext                unsafe.Pointer
    WaitSemaphoreCount   uint32
    PWaitSemaphores      *VulkanSemaphore
    PWaitDstStageMask    *VkPipelineStageFlags
    CommandBufferCount   uint32
    PCommandBuffers      *VulkanCommandBuffer // Oprava typu
    SignalSemaphoreCount uint32
    PSignalSemaphores    *VulkanSemaphore
}
type VkWriteDescriptorSet struct {
    SType           VkStructureType
    PNext           unsafe.Pointer
    DstSet          VulkanDescriptorSet
    DstBinding      uint32
    DstArrayElement uint32
    DescriptorCount uint32
    DescriptorType  VkDescriptorType
    PImageInfo      *VkDescriptorImageInfo
    PBufferInfo     *VkDescriptorBufferInfo // Pridanie chýbajúceho poľa
    PTexelBufferView *VulkanBufferView
}
type VkComputePipelineCreateInfo struct {
    SType              VkStructureType
    PNext              unsafe.Pointer
    Flags              VkPipelineCreateFlags
    Stage              VkPipelineShaderStageCreateInfo
    Layout             VulkanPipelineLayout
    BasePipelineHandle VulkanPipeline
    BasePipelineIndex  int32
}

type VkCommandBufferBeginInfo struct {
    SType            VkStructureType
    PInheritanceInfo *VkCommandBufferInheritanceInfo
}
type VkDescriptorSetAllocateInfo struct {
    SType              VkStructureType
    PNext              unsafe.Pointer
    DescriptorPool     VulkanDescriptorPool
    DescriptorSetCount uint32
    PSetLayouts        *VulkanDescriptorSetLayout
}

type VkPipelineLayoutCreateInfo struct {
    SType                  VkStructureType
    PNext                  unsafe.Pointer
    Flags                  VkPipelineLayoutCreateFlags
    SetLayoutCount         uint32
    PSetLayouts            *VulkanDescriptorSetLayout
    PushConstantRangeCount uint32
    PPushConstantRanges    *VkPushConstantRange
}

type VkPushConstantRange struct {
    StageFlags VkShaderStageFlagBits
    Offset     uint32
    Size       uint32
}
type VkDescriptorSetLayoutBinding struct {
    Binding            uint32
    DescriptorType     VkDescriptorType
    DescriptorCount    uint32
    StageFlags         VkShaderStageFlagBits
    PImmutableSamplers *VulkanSampler
}

type VkDescriptorSetLayoutCreateInfo struct {
    SType        VkStructureType
    PNext        unsafe.Pointer
    Flags        VkDescriptorSetLayoutCreateFlags
    BindingCount uint32
    PBindings    *VkDescriptorSetLayoutBinding
}
type VkCommandBufferInheritanceInfo struct {
    SType              VkStructureType
    PNext              unsafe.Pointer
    RenderPass         VulkanRenderPass
    Subpass            uint32
    Framebuffer        VulkanFramebuffer
    OcclusionQueryEnable VkBool32
    QueryFlags         VkQueryControlFlags
    PipelineStatistics VkQueryPipelineStatisticFlags
}

type VkCommandBufferAllocateInfo struct {
    SType              VkStructureType
    PNext              unsafe.Pointer
    CommandPool        VulkanCommandPool
    Level              VkCommandBufferLevel
    CommandBufferCount uint32
}
type VkShaderModuleCreateInfo struct {
    SType    VkStructureType
    PNext    unsafe.Pointer
    Flags    uint32
    CodeSize uintptr
    PCode    *uint32
}
// Vulkan Application Info
type VkApplicationInfo struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	PApplicationName   *byte
	ApplicationVersion uint32
	PEngineName        *byte
	EngineVersion      uint32
	ApiVersion         uint32
}

// Vulkan Instance Create Info
type VkInstanceCreateInfo struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	Flags                 uint32
	PApplicationInfo      *VkApplicationInfo
	EnabledLayerCount     uint32
	PpEnabledLayerNames   **byte
	EnabledExtensionCount uint32
	PpEnabledExtensionNames **byte
}

// Vulkan Device Create Info
type VkDeviceCreateInfo struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	Flags                   uint32
	QueueCreateInfoCount    uint32
	PQueueCreateInfos       unsafe.Pointer
	EnabledLayerCount       uint32
	PpEnabledLayerNames     **byte
	EnabledExtensionCount   uint32
	PpEnabledExtensionNames **byte
	PEnabledFeatures        unsafe.Pointer
}

// Vulkan Buffer Create Info
type VkBufferCreateInfo struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	Flags                 uint32
	Size                  uint64
	Usage                 uint32
	SharingMode           uint32
	QueueFamilyIndexCount uint32
	PQueueFamilyIndices   *uint32
}

// Vulkan Image Create Info
type VkImageCreateInfo struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	Flags                 uint32
	ImageType             uint32
	Format                uint32
	Extent                [3]uint32
	MipLevels             uint32
	ArrayLayers           uint32
	Samples               uint32
	Tiling                uint32
	Usage                 uint32
	SharingMode           uint32
	QueueFamilyIndexCount uint32
	PQueueFamilyIndices   *uint32
	InitialLayout         uint32
}


// Vulkan Command Pool Create Info
type VkCommandPoolCreateInfo struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	Flags            uint32
	QueueFamilyIndex uint32
}

// Vulkan Fence Create Info
type VkFenceCreateInfo struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags uint32
}

type VkSpecializationMapEntry struct {
	ConstantID uint32
	Offset     uint32
	Size       uint
}

type VkDescriptorPoolSize struct {
	Type            VkDescriptorType
	DescriptorCount uint32
}

type VkPhysicalDeviceLimits struct {
	MaxImageDimension1D uint32
	MaxImageDimension2D uint32
	MaxImageDimension3D uint32
	MaxImageDimensionCube uint32
	MaxImageArrayLayers uint32
	MaxTexelBufferElements uint32
	MaxUniformBufferRange uint32
	MaxStorageBufferRange uint32
	MaxPushConstantsSize uint32
	// ... (existuje mnoho ďalších parametrov)
}

type VkPhysicalDeviceSparseProperties struct {
	ResidencyStandard2DBlockShape            VkBool32
	ResidencyStandard2DMultisampleBlockShape VkBool32
	ResidencyStandard3DBlockShape            VkBool32
	ResidencyAlignedMipSize                  VkBool32
	ResidencyNonResidentStrict               VkBool32
}

type VkSpecializationInfo struct {
	MapEntryCount uint32
	PMapEntries   *VkSpecializationMapEntry
	DataSize      uint
	PData         unsafe.Pointer
}

type VkDescriptorPoolCreateInfo struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	Flags         VkDescriptorPoolCreateFlags
	MaxSets       uint32
	PoolSizeCount uint32
	PPoolSizes    *VkDescriptorPoolSize
}

type VkPhysicalDeviceProperties struct {
	ApiVersion        uint32
	DriverVersion     uint32
	VendorID          uint32
	DeviceID          uint32
	DeviceType        VkPhysicalDeviceType
	DeviceName        [256]byte
	PipelineCacheUUID [16]byte
	Limits            VkPhysicalDeviceLimits
	SparseProperties  VkPhysicalDeviceSparseProperties
}

type VkPipelineShaderStageCreateInfo struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	Flags               VkPipelineShaderStageCreateFlags
	Stage               VkShaderStageFlagBits
	Module              VulkanShaderModule
	PName               *byte // C-string (napr. "main")
	PSpecializationInfo *VkSpecializationInfo
}

type BufferData struct {
	TypeSize uintptr
	DataSize uintptr
	Pointer  unsafe.Pointer
}

type ImageData struct {
	BufferData *BufferData
	Origin     [3]uint32
	Region     [3]uint32
	RowPitch   uint32
	SlicePitch uint32
}

// Vulkan alokačné callbacky (nepovinné, môže byť nil)
type VkAllocationCallbacks struct {
	PUserData       uintptr
	PfnAllocation   uintptr
	PfnReallocation uintptr
	PfnFree         uintptr
	PfnInternal     uintptr
}

type VkDescriptorBufferInfo struct {
    Buffer VulkanBuffer
    Offset uint64
    Range  uint64
}

type VkCopyDescriptorSet struct {
    SType           uint32
    PNext           uintptr
    SrcSet          VulkanDescriptorSet
    SrcBinding      uint32
    SrcArrayElement uint32
    DstSet          VulkanDescriptorSet
    DstBinding      uint32
    DstArrayElement uint32
    DescriptorCount uint32
}

type VkDeviceQueueCreateInfo struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	Flags            uint32
	QueueFamilyIndex uint32
	QueueCount       uint32
	PQueuePriorities *float32
}
