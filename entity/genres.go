package entity

// https://www.kinokuniya.co.jp/disp/CKnSfGenreSelect.jsp?dispNo=101001
var genres = []string{
	"文芸",
	"教養",
	"人文",
	"教育",
	"社会",
	"法律",
	"経済",
	"経営",
	"ビジネス",
	"就職・資格",
	"理学",
	"工学",
	"コンピュータ",
	"医学",
	"看護学",
	"薬学",
	"芸術",
	"語学",
	"辞典",
	"趣味・生活",
	"くらし・料理",
	"地図・ガイド",
	"文庫",
	"コミック",
	"エンターテイメント",
}

func GetGenres() []string {
	return genres
}
