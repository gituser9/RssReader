//
//  VkVkConfigurator.swift
//  NewsHub-iOS
//
//  Created by user on 04/01/2018.
//  Copyright © 2018 user. All rights reserved.
//

import UIKit

class VkModuleConfigurator {

    func configureModuleForViewInput<UIViewController>(viewInput: UIViewController) {

        if let viewController = viewInput as? VkViewController {
            configure(viewController: viewController)
        }
    }

    private func configure(viewController: VkViewController) {

        let router = VkRouter()

        let presenter = VkPresenter()
        presenter.view = viewController
        presenter.router = router

        let interactor = VkInteractor()
        interactor.output = presenter

        presenter.interactor = interactor
        viewController.output = presenter
    }

}
