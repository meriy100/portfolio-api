package entities

import (
	"reflect"
	"testing"
)

func TestPost_ToHistories(t *testing.T) {
	tests := []struct {
		name   string
		post   Post
		got    []History
		gotErr bool
	}{
		{
			"success",
			Post{
				BodyMd: "# 株式会社 Photosynth.inc\r\n## スマートロック関連のプロダクト開発業務\r\n### startMonth\r\n2020/03\r\n### endMonth\r\n\r\n### description\r\n- サーバーサイド設計, 開発, レビュー\r\n- フロントエンド設計, 開発, レビュー\r\n\r\n###  technologyUsed \r\n- Ruby on Rails\r\n- Vue.js\r\n- HIDDEN\r\n\r\n# 合同会社フレイズ (業務委託)\r\n\r\n## スマートフォンブラウザでの映像撮影, 画像認識処理\r\n### startMonth\r\n 2019/11\r\n###  endMonth\r\n 2020/02\r\n###    description\r\n- 技術調査 (Web カメラ, 顔認識, 文字認識等)\r\n- フロントエンド設計, 開発\r\n###  technologyUsed \r\n- TypeScript\r\n- Google Cloud Platform\r\n- Vue.js\r\n- OpenCV\r\n- face api (tensorflow.js)\r\n\r\n## サンプルプロダクト\r\n### startMonth\r\n 2019/11\r\n###  endMonth\r\n 2020/02\r\n###    description\r\n- 技術調査 (Web カメラ, 顔認識, 文字認識等)\r\n- フロントエンド設計, 開発\r\n###  technologyUsed \r\n- TypeScript\r\n- Google Cloud Platform\r\n- Vue.js\r\n- OpenCV\r\n- face api (tensorflow.js)\r\n",
			},
			[]History{
				{
					"株式会社 Photosynth.inc",
					[]Product{
						{
							"スマートロック関連のプロダクト開発業務",
							Month{2020, 3},
							nil,
							[]string{
								"サーバーサイド設計, 開発, レビュー",
								"フロントエンド設計, 開発, レビュー",
							},
							[]string{
								"Ruby on Rails",
								"Vue.js",
								"HIDDEN",
							},
						},
					},
				},
				{
					"合同会社フレイズ (業務委託)",
					[]Product{
						{
							"スマートフォンブラウザでの映像撮影, 画像認識処理",
							Month{2019, 11},
							&Month{2020, 2},
							[]string{
								"技術調査 (Web カメラ, 顔認識, 文字認識等)",
								"フロントエンド設計, 開発",
							},
							[]string{
								"TypeScript",
								"Google Cloud Platform",
								"Vue.js",
								"OpenCV",
								"face api (tensorflow.js)",
							},
						},
						{
							"サンプルプロダクト",
							Month{2019, 11},
							&Month{2020, 2},
							[]string{
								"技術調査 (Web カメラ, 顔認識, 文字認識等)",
								"フロントエンド設計, 開発",
							},
							[]string{
								"TypeScript",
								"Google Cloud Platform",
								"Vue.js",
								"OpenCV",
								"face api (tensorflow.js)",
							},
						},
					},
				},
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			histories, err := tt.post.ToHistories()
			if (err != nil) != tt.gotErr {
				t.Errorf("ToHistory() : err = %v want gotErr %v\n", err, tt.gotErr)
			}
			if len(histories) != len(tt.got) {
				t.Errorf("length different : histories = %v want %v\n", histories, tt.got)
			}

			for i, got := range tt.got {
				h := histories[i]
				if histories[i].Organization != got.Organization {
					t.Errorf("histories.[%d].Organization = %v want %v\n", i, h.Organization, got.Organization)

				}

				if len(h.Products) != len(got.Products) {
					t.Errorf("histories.[%d].Products length different : Products = %v want %v\n", i, h.Products, got.Products)
				}

				for j, gotP := range got.Products {
					p := h.Products[j]
					if p.Title != gotP.Title {
						t.Errorf("histories.[%d].Products[%d].Title = %v want %v\n", i, j, p.Title, gotP.Title)
					}
					if p.StartMonth != gotP.StartMonth {
						t.Errorf("histories.[%d].Products[%d].StartMonth= %v want %v\n", i, j, p.StartMonth, gotP.StartMonth)
					}
					if (p.EndMonth == nil) != (gotP.EndMonth == nil) {
						t.Errorf("histories.[%d].Products[%d].EndMonth= %v want %v\n", i, j, p.EndMonth, gotP.EndMonth)
					} else if p.EndMonth != nil && gotP.EndMonth != nil {
						if *p.EndMonth != *gotP.EndMonth {
							t.Errorf("histories.[%d].Products[%d].EndMonth= %v want %v\n", i, j, p.EndMonth, gotP.EndMonth)
						}
					}

					if !reflect.DeepEqual(p.Description, gotP.Description) {
						t.Errorf("histories.[%d].Products[%d].Description= %s want %s\n", i, j, p.Description, gotP.Description)
					}

					if !reflect.DeepEqual(p.Technologies, gotP.Technologies) {
						t.Errorf("histories.[%d].Products[%d].Technologies= %v want %v\n", i, j, p.Technologies, gotP.Technologies)
					}
				}
			}
		})
	}
}
