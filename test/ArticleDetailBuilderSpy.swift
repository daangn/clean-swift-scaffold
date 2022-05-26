//
//  ArticleDetailBuilderSpy.swift
//  MiroTests
//
//  Created by Geektree0101 on 12/10/2020.
//  Copyright © 2020 miro. All rights reserved.
//

import UIKit

import KarrotCore

final class ArticleDetailBuilderSpy: ArticleDetailBuildable {

  var buildCalled: Int = 0
  var buildPayload: ArticleDetailPayload?
  var buildStub: UIViewController = UIViewController()

  func build(payload: ArticleDetailPayload) -> UIViewController {
    self.buildCalled += 1
    self.buildPayload = payload
    return self.buildStub
  }
}