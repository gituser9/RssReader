//
//  LoginLoginViewController.swift
//  NewsHub-iOS
//
//  Created by user on 04/11/2017.
//  Copyright © 2017 user. All rights reserved.
//

import UIKit

class LoginViewController: UIViewController, UITextFieldDelegate {

    @IBOutlet weak var loginTextField: UITextField!
    @IBOutlet weak var passwordTextField: UITextField!
    @IBOutlet weak var waitIndicator: UIActivityIndicatorView!

    var presenter: LoginPresenter?

    // MARK: Life cycle
    override func viewDidLoad() {
        super.viewDidLoad()

        presenter = LoginPresenter()
        presenter?.view = self
        
        waitIndicator.stopAnimating()
    }

    
    // MARK: UITextFieldDelegate
    func textFieldShouldReturn(_ textField: UITextField) -> Bool {
        if textField.isEqual(loginTextField) {
            passwordTextField.becomeFirstResponder()
        }
        if textField.isEqual(passwordTextField) {
            textField.resignFirstResponder()
            waitIndicator.startAnimating()
        }
        return true
    }
    
    
    // MARK: Actions
    @IBAction func loginAction(_ sender: UIButton) {
        waitIndicator.stopAnimating()
        presenter?.login(username: loginTextField.text, password: passwordTextField.text)
    }
    
}
