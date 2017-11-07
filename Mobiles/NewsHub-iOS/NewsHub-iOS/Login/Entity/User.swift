import Foundation


struct User : Decodable {
    var Id = 0
    var Name = ""
    var Settings: Settings? = nil
}
