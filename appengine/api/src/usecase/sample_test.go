package usecase_test

import (
	"context"
	"testing"

	"github.com/aikizoku/rabbitgo/appengine/api/src/repository"
	"github.com/aikizoku/rabbitgo/appengine/api/src/usecase"
)

func Test_Sample_UnitTestMethod(t *testing.T) {
	// テスト実行時のinputの定義
	type args struct {
		hoge int
		fuga int
	}
	// 期待するoutputの定義
	type want struct {
		result int
		err    bool
	}
	// テストケースの定義（これはずっと固定）
	type testCase struct {
		name string
		args args
		want want
	}

	// テスト実行の準備
	ctx := context.Background()
	rSample := repository.NewSample(nil, nil)
	uSample := usecase.NewSample(rSample)

	// テストケースの定義
	tcs := []testCase{
		{
			name: "正常系",
			args: args{
				hoge: 1,
				fuga: 2,
			},
			want: want{
				result: 3,
				err:    false,
			},
		},
		{
			name: "異常系",
			args: args{
				hoge: -1,
				fuga: 2,
			},
			want: want{
				result: 0,
				err:    true,
			},
		},
	}

	// テストケース分のテスト実行
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// テスト対象メソッドの実行
			result, err := uSample.UnitTestMethod(ctx, tc.args.hoge, tc.args.fuga)

			// 結果を検証
			if result != tc.want.result {
				t.Errorf("result = %v, want %v", result, tc.want.result)
			}
			if (err != nil) != tc.want.err {
				t.Errorf("err = %v, want %v", err, tc.want.err)
			}
		})
	}
}
