//
//  RssRssConfiguratorTests.swift
//  NewsHub-iOS
//
//  Created by user on 02/11/2017.
//  Copyright © 2017 user. All rights reserved.
//

import XCTest

class RssModuleConfiguratorTests: XCTestCase {

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
        let viewController = RssViewControllerMock()
        let configurator = RssModuleConfigurator()

        //when
        configurator.configureModuleForViewInput(viewInput: viewController)

        //then
        XCTAssertNotNil(viewController.output, "RssViewController is nil after configuration")
        XCTAssertTrue(viewController.output is RssPresenter, "output is not RssPresenter")

        let presenter: RssPresenter = viewController.output as! RssPresenter
        XCTAssertNotNil(presenter.view, "view in RssPresenter is nil after configuration")
        XCTAssertNotNil(presenter.router, "router in RssPresenter is nil after configuration")
        XCTAssertTrue(presenter.router is RssRouter, "router is not RssRouter")

        let interactor: RssInteractor = presenter.interactor as! RssInteractor
        XCTAssertNotNil(interactor.output, "output in RssInteractor is nil after configuration")
    }

    class RssViewControllerMock: RssViewController {

        var setupInitialStateDidCall = false

        override func setupInitialState() {
            setupInitialStateDidCall = true
        }
    }
}
