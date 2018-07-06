//
//  VkVkPresenterTests.swift
//  NewsHub-iOS
//
//  Created by user on 04/01/2018.
//  Copyright © 2018 user. All rights reserved.
//

import XCTest

class VkPresenterTest: XCTestCase {

    override func setUp() {
        super.setUp()
        // Put setup code here. This method is called before the invocation of each test method in the class.
    }

    override func tearDown() {
        // Put teardown code here. This method is called after the invocation of each test method in the class.
        super.tearDown()
    }

    class MockInteractor: VkInteractorInput {

    }

    class MockRouter: VkRouterInput {

    }

    class MockViewController: VkViewInput {

        func setupInitialState() {

        }
    }
}
