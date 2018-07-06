//
//  VkVkConfiguratorTests.swift
//  NewsHub-iOS
//
//  Created by user on 04/01/2018.
//  Copyright © 2018 user. All rights reserved.
//

import XCTest

class VkModuleConfiguratorTests: XCTestCase {

    override func setUp() {
        super.setUp()
        // Put setup code here. This method is called before the invocation of each test method in the class.
    }

    override func tearDown() {
        // Put teardown code here. This method is called after the invocation of each test method in the class.
        super.tearDown()
    }

    func testConfigureModuleForViewController() {

        //given
        let viewController = VkViewControllerMock()
        let configurator = VkModuleConfigurator()

        //when
        configurator.configureModuleForViewInput(viewInput: viewController)

        //then
        XCTAssertNotNil(viewController.output, "VkViewController is nil after configuration")
        XCTAssertTrue(viewController.output is VkPresenter, "output is not VkPresenter")

        let presenter: VkPresenter = viewController.output as! VkPresenter
        XCTAssertNotNil(presenter.view, "view in VkPresenter is nil after configuration")
        XCTAssertNotNil(presenter.router, "router in VkPresenter is nil after configuration")
        XCTAssertTrue(presenter.router is VkRouter, "router is not VkRouter")

        let interactor: VkInteractor = presenter.interactor as! VkInteractor
        XCTAssertNotNil(interactor.output, "output in VkInteractor is nil after configuration")
    }

    class VkViewControllerMock: VkViewController {

        var setupInitialStateDidCall = false

        override func setupInitialState() {
            setupInitialStateDidCall = true
        }
    }
}
