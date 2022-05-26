//
//  __SCENE_NAME__Builder.swift
//  __TARGET_PROJECT_NAME__
//
//  Created by __CREATOR__ on __DATE__.
//  Copyright © __YEAR__ __COPYRIGHT__. All rights reserved.
//

import UIKit

import KarrotCore

struct __SCENE_NAME__Dependency {

}

typealias __SCENE_NAME__Payload = Void

protocol __SCENE_NAME__Buildable {
  func build() -> UIViewController
}

final class __SCENE_NAME__Builder:
  KarrotBuilder<UIViewController, __SCENE_NAME__Dependency, __SCENE_NAME__Payload>,
  __SCENE_NAME__Buildable {

  override func build(payload: __SCENE_NAME__Payload) -> UIViewController {
    let viewController = __SCENE_NAME__ViewController()
    let interactor = __SCENE_NAME__Interactor()
    let presenter = __SCENE_NAME__Presenter()
    let router = __SCENE_NAME__Router()

    interactor.presenter = presenter

    presenter.view = viewController

    router.viewController = viewController
    router.dataStore = interactor

    viewController.interactor = interactor
    viewController.router = router

    return viewController
  }
}