//
//  ArticleDetailBuilder.swift
//  Miro
//
//  Created by Geektree0101 on 12/10/2020.
//  Copyright © 2020 miro. All rights reserved.
//

import UIKit

import KarrotCore

struct ArticleDetailDependency {

}

typealias ArticleDetailPayload = Void

protocol ArticleDetailBuildable {
  func build() -> UIViewController
}

final class ArticleDetailBuilder:
  KarrotBuilder<UIViewController, ArticleDetailDependency, ArticleDetailPayload>,
  ArticleDetailBuildable {

  override func build(payload: ArticleDetailPayload) -> UIViewController {
    let viewController = ArticleDetailViewController()
    let interactor = ArticleDetailInteractor()
    let presenter = ArticleDetailPresenter()
    let router = ArticleDetailRouter()

    interactor.presenter = presenter

    presenter.view = viewController

    router.viewController = viewController
    router.dataStore = interactor

    viewController.interactor = interactor
    viewController.router = router

    return viewController
  }
}