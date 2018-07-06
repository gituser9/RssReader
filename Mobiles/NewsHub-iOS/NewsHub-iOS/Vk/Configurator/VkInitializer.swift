//
//  VkVkInitializer.swift
//  NewsHub-iOS
//
//  Created by user on 04/01/2018.
//  Copyright © 2018 user. All rights reserved.
//

import UIKit

class VkModuleInitializer: NSObject {

    //Connect with object on storyboard
    @IBOutlet weak var vkViewController: VkViewController!

    override func awakeFromNib() {

        let configurator = VkModuleConfigurator()
        configurator.configureModuleForViewInput(viewInput: vkViewController)
    }

}
