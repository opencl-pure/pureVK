package pureVK


// Structure Types
const (
    VK_STRUCTURE_TYPE_APPLICATION_INFO VkStructureType = iota
    VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO
    VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
    VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
    VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
    VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO
    VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO
    VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO
    VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
    VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
    VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO
    VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO
    VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO
    VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO
    VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO
    VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
    VK_STRUCTURE_TYPE_SUBMIT_INFO
    VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
    VK_STRUCTURE_TYPE_DESCRIPTOR_BUFFER_INFO // Pridané
)

// Shader Stage Flags
const (
    VK_SHADER_STAGE_VERTEX_BIT VkShaderStageFlagBits = 0x00000001 << iota
    VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT
    VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT
    VK_SHADER_STAGE_GEOMETRY_BIT
    VK_SHADER_STAGE_FRAGMENT_BIT
    VK_SHADER_STAGE_COMPUTE_BIT
)

// Physical Device Types
const (
        VK_PHYSICAL_DEVICE_TYPE_OTHER VkPhysicalDeviceType = iota
        VK_PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU
        VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU
        VK_PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU
        VK_PHYSICAL_DEVICE_TYPE_CPU
)

// Descriptor Pool Create Flags
const (
        VK_DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT VkDescriptorPoolCreateFlags = 0x00000001 << iota
        VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT
)

// Descriptor Types
const (
        VK_DESCRIPTOR_TYPE_SAMPLER VkDescriptorType = iota
        VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER
        VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE
        VK_DESCRIPTOR_TYPE_STORAGE_IMAGE
        VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER
        VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER
        VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER
        VK_DESCRIPTOR_TYPE_STORAGE_BUFFER
        VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC
        VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC
        VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT
)

// Boolean Values
const (
        VK_FALSE VkBool32 = 0
        VK_TRUE  VkBool32 = 1
)

// API Versions
const (
        VK_API_VERSION_1_0 = (1 << 22) | (0 << 12)
)

// Buffer Usage Flags
const (
        VK_BUFFER_USAGE_TRANSFER_SRC_BIT     uint32 = 0x00000001 << iota
        VK_BUFFER_USAGE_TRANSFER_DST_BIT
        VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT
        VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT
        VK_BUFFER_USAGE_UNIFORM_BUFFER_BIT
        VK_BUFFER_USAGE_VERTEX_BUFFER_BIT
        VK_BUFFER_USAGE_INDEX_BUFFER_BIT
        VK_BUFFER_USAGE_STORAGE_BUFFER_BIT
        VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT
        VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT = 0x00020000 // Toto nie je sekvenčné
)

const (
	VK_COMMAND_BUFFER_LEVEL_PRIMARY                VkCommandBufferLevel = 0
)

const (
    VK_PIPELINE_BIND_POINT_COMPUTE            VkPipelineBindPoint = 2
)

const (
    VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT   VkMemoryPropertyFlags = 1
    VK_MEMORY_PROPERTY_HOST_COHERENT_BIT  VkMemoryPropertyFlags = 2
)

const (
         VK_WHOLE_SIZE uint64 = ^uint64(0)
    VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT VkMemoryPropertyFlags = 4
)
const VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT VkCommandBufferUsageFlags = 1