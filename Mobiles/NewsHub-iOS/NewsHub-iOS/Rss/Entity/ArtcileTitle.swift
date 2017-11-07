import Foundation


struct Articles: Decodable {
    var Count: Int = 0
    var Articles = [ArticleTitle]()
}

struct ArticleTitle : Decodable {
    var Id = 0
    var IsRead = false
    var Link = ""
    var Title = ""
}
