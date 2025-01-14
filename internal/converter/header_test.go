package converter_test

import (
	"testing"
	"time"

	"github.com/daangn/clean-swift-scaffold/internal/converter"
	"github.com/daangn/clean-swift-scaffold/internal/model"
)

const dummySourceCode string = `//
//  __SCENE_NAME__Model.swift
//  __TARGET_PROJECT_NAME__
//
//  Created by __CREATOR__ on __DATE__.
//  Copyright © __YEAR__ __COPYRIGHT__. All rights reserved.
//

enum __SCENE_NAME__Model {

	// clean-swift-scaffold-generate-dto (do-not-remove-comments)
}`

const expectedSourceCode string = `//
//  ArticleDetailModel.swift
//  Miro
//
//  Created by Geektree0101 on 12/10/2020.
//  Copyright © 2020 miro. All rights reserved.
//

enum ArticleDetailModel {

	// clean-swift-scaffold-generate-dto (do-not-remove-comments)
}`

func TestHeader(t *testing.T) {

	t.Run("return expected header", func(t *testing.T) {
		// given
		config := model.Config{
			TargetProjectName: "Miro",
			Copyright:         "miro",
			TemplatePath:      "",
		}

		date := time.Date(2020, 10, 12, 0, 0, 0, 0, time.UTC)
		sut := converter.NewHeaderConverter(
			CreatorStub{
				GetSuccessStub: "Geektree0101",
				GetErrorStub:   nil,
			},
			&config,
			date,
		)

		usecaseName := "ArticleDetail"

		// when
		output := sut.Render(dummySourceCode, usecaseName)

		// then
		if output != expectedSourceCode {
			t.Errorf("Failed to render\n expected:\n%s\noutput:\n%s\n", expectedSourceCode, output)
		}
	})
}
