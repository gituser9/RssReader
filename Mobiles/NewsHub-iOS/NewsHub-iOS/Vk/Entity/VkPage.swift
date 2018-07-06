import Foundation


struct VkPage : Decodable {
    var groups = [VkGroup]()
    var news = [VkNews]()
}

struct VkPageView {
    var groups = [Int:VkGroup]()
    var news = [VkNews]()
}
