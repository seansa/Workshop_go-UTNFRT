package locales

var Ja = map[string]interface{}{
	"name": map[string]interface{}{
		"last_name": []string{
			"佐藤", "鈴木", "高橋", "田中", "渡辺", "伊藤", "山本", "中村", "小林", "加藤", "吉田", "山田", "佐々木", "山口", "斎藤", "松本", "井上", "木村", "林", "清水",
		},
		"first_name": []string{
			"大翔", "蓮", "颯太", "樹", "大和", "陽翔", "陸斗", "太一", "海翔", "蒼空", "翼", "陽菜", "結愛", "結衣", "杏", "莉子", "美羽", "結菜", "心愛", "愛菜", "美咲",
		},
		"name": []string{
			"#{name.last_name} #{name.first_name}",
		},
	},
	"address": map[string]interface{}{
		"city_suffix": []string{
			"市", "区", "町", "村",
		},
		"city": []string{
			"#{address.city_prefix}#{name.first_name}#{address.city_suffix}", "#{name.first_name}#{address.city_suffix}", "#{address.city_prefix}#{name.last_name}#{address.city_suffix}", "#{name.last_name}#{address.city_suffix}",
		},
		"street_name": []string{
			"#{name.first_name}#{address.street_suffix}", "#{name.last_name}#{address.street_suffix}",
		},
		"postcode": []string{
			"###-####",
		},
		"state": []string{
			"北海道", "青森県", "岩手県", "宮城県", "秋田県", "山形県", "福島県", "茨城県", "栃木県", "群馬県", "埼玉県", "千葉県", "東京都", "神奈川県", "新潟県", "富山県", "石川県", "福井県", "山梨県", "長野県", "岐阜県", "静岡県", "愛知県", "三重県", "滋賀県", "京都府", "大阪府", "兵庫県", "奈良県", "和歌山県", "鳥取県", "島根県", "岡山県", "広島県", "山口県", "徳島県", "香川県", "愛媛県", "高知県", "福岡県", "佐賀県", "長崎県", "熊本県", "大分県", "宮崎県", "鹿児島県", "沖縄県",
		},
		"state_abbr": []string{
			"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47",
		},
		"city_prefix": []string{
			"北", "東", "西", "南", "新", "湖", "港",
		},
	},
	"phone_number": map[string]interface{}{
		"formats": []string{
			"0####-#-####", "0###-##-####", "0##-###-####", "0#-####-####",
		},
	},
	"cell_phone": map[string]interface{}{
		"formats": []string{
			"090-####-####", "080-####-####", "070-####-####"}}}
