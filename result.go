package pureVK

import "fmt"

const (
        VK_SUCCESS                                   VkResult = 0
        VK_NOT_READY                                 VkResult = 1
        VK_TIMEOUT                                   VkResult = 2
        VK_EVENT_SET                                 VkResult = 3
        VK_EVENT_RESET                               VkResult = 4
        VK_INCOMPLETE                                VkResult = 5
        VK_SUBOPTIMAL_KHR                            VkResult = 1000001003
        VK_ERROR_OUT_OF_HOST_MEMORY                   VkResult = -1
        VK_ERROR_OUT_OF_DEVICE_MEMORY                 VkResult = -2
        VK_ERROR_INITIALIZATION_FAILED                VkResult = -3
        VK_ERROR_DEVICE_LOST                         VkResult = -4
        VK_ERROR_MEMORY_MAP_FAILED                    VkResult = -5
        VK_ERROR_LAYER_NOT_PRESENT                    VkResult = -6
        VK_ERROR_EXTENSION_NOT_PRESENT                VkResult = -7
        VK_ERROR_FEATURE_NOT_PRESENT                  VkResult = -8
        VK_ERROR_INCOMPATIBLE_DRIVER                  VkResult = -9
        VK_ERROR_TOO_MANY_OBJECTS                     VkResult = -10
        VK_ERROR_FORMAT_NOT_SUPPORTED                 VkResult = -11
        VK_ERROR_FRAGMENTED_POOL                      VkResult = -12
        VK_ERROR_OUT_OF_DEVICE_MEMORY_SPECIFIC        VkResult = -13
        VK_ERROR_SURFACE_LOST_KHR                     VkResult = -1000000000
        VK_ERROR_NATIVE_WINDOW_IN_USE_KHR             VkResult = -1000000001
        VK_ERROR_OUT_OF_DATE_KHR                      VkResult = -1000001004
        VK_ERROR_INCOMPATIBLE_DISPLAY_KHR             VkResult = -1000003001
        VK_ERROR_INVALID_SHADER_NV                    VkResult = -1000012000
        VK_ERROR_OUT_OF_POOL_MEMORY                   VkResult = -1000069000
        VK_ERROR_INVALID_EXTERNAL_HANDLE              VkResult = -1000072003
        VK_ERROR_FRAGMENTATION                        VkResult = -1000161000
        VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS       VkResult = -1000257000
        VK_ERROR_PIPELINE_COMPILE_REQUIRED_EXT        VkResult = -1000297000
        VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT  VkResult = -1000255001
        VK_ERROR_VALIDATION_FAILED_EXT                VkResult = -1000011001
        VK_ERROR_INVALID_DRM_FORMAT_MODIFIER_PLANE_LAYOUT_EXT VkResult = -1000158007
        VK_ERROR_NOT_PERMITTED_EXT                    VkResult = -1000174001
        VK_ERROR_INVALID_DEVICE_ADDRESS_EXT           VkResult = -1000257001
        VK_ERROR_COMPRESSION_EXHAUSTED_EXT            VkResult = -1000338000
        VK_ERROR_UNKNOWN                              VkResult = -9
)

func HandleVkResult(result VkResult) string {
        switch result {
        case VK_SUCCESS:
                return "Operácia bola úspešná."
        case VK_NOT_READY:
                return "Operácia nie je dokončená."
        case VK_TIMEOUT:
                return "Operácia vypršala."
        case VK_EVENT_SET:
                return "Event bol nastavený."
        case VK_EVENT_RESET:
                return "Event bol resetovaný."
        case VK_INCOMPLETE:
                return "Operácia bola neúplná."
        case VK_SUBOPTIMAL_KHR:
                return "Swapchain je suboptimal."
        case VK_ERROR_OUT_OF_HOST_MEMORY:
                return "Chyba: Nedostatok pamäte hostiteľa (RAM)."
        case VK_ERROR_OUT_OF_DEVICE_MEMORY:
                return "Chyba: Nedostatok pamäte zariadenia (GPU)."
        case VK_ERROR_INITIALIZATION_FAILED:
                return "Chyba: Inicializácia zlyhala."
        case VK_ERROR_DEVICE_LOST:
                return "Chyba: Zariadenie (GPU) bolo stratené."
        case VK_ERROR_MEMORY_MAP_FAILED:
                return "Chyba: Mapovanie pamäte zlyhalo."
        case VK_ERROR_LAYER_NOT_PRESENT:
                return "Chyba: Vrstva nebola nájdená."
        case VK_ERROR_EXTENSION_NOT_PRESENT:
                return "Chyba: Rozšírenie nebolo nájdené."
        case VK_ERROR_FEATURE_NOT_PRESENT:
                return "Chyba: Funkcia nebola nájdená."
        case VK_ERROR_INCOMPATIBLE_DRIVER:
                return "Chyba: Nekompatibilný ovládač."
        case VK_ERROR_TOO_MANY_OBJECTS:
                return "Chyba: Príliš veľa objektov."
        case VK_ERROR_FORMAT_NOT_SUPPORTED:
                return "Chyba: Formát nie je podporovaný."
        case VK_ERROR_FRAGMENTED_POOL:
                return "Chyba: Fragmentovaný pool."
        case VK_ERROR_SURFACE_LOST_KHR:
                return "Chyba: Povrch (surface) bol stratený."
        case VK_ERROR_NATIVE_WINDOW_IN_USE_KHR:
                return "Chyba: Natívne okno sa používa."
        case VK_ERROR_OUT_OF_DATE_KHR:
                return "Chyba: Swapchain je zastaraný."
        case VK_ERROR_INCOMPATIBLE_DISPLAY_KHR:
                return "Chyba: Nekompatibilný displej."
        case VK_ERROR_INVALID_SHADER_NV:
                return "Chyba: Neplatný shader."
        case VK_ERROR_OUT_OF_POOL_MEMORY:
                return "Chyba: Nedostatok pamäte v pool-e."
        case VK_ERROR_INVALID_EXTERNAL_HANDLE:
                return "Chyba: Neplatný externý handle."
        case VK_ERROR_FRAGMENTATION:
                return "Chyba: Fragmentácia."
        case VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS:
                return "Chyba: Neplatná opaque capture adresa."
        case VK_ERROR_PIPELINE_COMPILE_REQUIRED_EXT:
                return "Chyba: Vyžaduje sa kompilácia pipeline."
        case VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT:
                return "Chyba: Exkluzívny režim celej obrazovky bol stratený."
        case VK_ERROR_VALIDATION_FAILED_EXT:
                return "Chyba: Validácia zlyhala."
        case VK_ERROR_INVALID_DRM_FORMAT_MODIFIER_PLANE_LAYOUT_EXT:
                return "Chyba: Neplatné rozloženie roviny modifikátora formátu DRM."
        case VK_ERROR_NOT_PERMITTED_EXT:
                return "Chyba: Operácia nie je povolená."
        case VK_ERROR_INVALID_DEVICE_ADDRESS_EXT:
                return "Chyba: Neplatná adresa zariadenia."
        case VK_ERROR_COMPRESSION_EXHAUSTED_EXT:
                return "Chyba: Kompresia vyčerpaná."
        case VK_ERROR_OUT_OF_DEVICE_MEMORY_SPECIFIC:
                return "Chyba: Nedostatok pamäte zariadenia (GPU) - špecifická chyba."
        default:
                return fmt.Sprint("Neznámy VkResult:", result)
        }
}