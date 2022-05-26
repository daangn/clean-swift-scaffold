//
//  __SCENE_NAME__BuilderSpy.swift
//  __TARGET_PROJECT_NAME__Tests
//
//  Created by __CREATOR__ on __DATE__.
//  Copyright © __YEAR__ __COPYRIGHT__. All rights reserved.
//

import UIKit

import KarrotCore

final class __SCENE_NAME__BuilderSpy: __SCENE_NAME__Buildable {

  var buildCalled: Int = 0
  var buildParams: __SCENE_NAME__Payload?
  var buildStub: UIViewController = UIViewController()

  func build(payload: __SCENE_NAME__Payload) -> UIViewController {
    self.buildCalled += 1
    self.buildParams = payload
    return self.buildStub
  }
}