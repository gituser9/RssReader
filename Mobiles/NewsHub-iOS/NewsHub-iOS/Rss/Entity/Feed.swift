import Foundation


struct FeedModel: Decodable {
    var ArticlesCount = 0
    var ExistUnread = false
    var Feed: Feed? = nil
}

struct Feed: Decodable {
    var Id = 0
    var Name = ""
}
