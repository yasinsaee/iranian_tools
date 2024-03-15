package util

// دریافت اسم بانک از شماره کارت
// این متود فقط میتواند اسم بانک را از شماره کارت تشخیص دهد و کارایی دیگری ندارد
func GetBankNameFromCardNumber(cardNumber string) string {
	if len(cardNumber) < 6 {
		return "-"
	}
	cardNumber = cardNumber[:6]

	switch cardNumber {
	case "636214":
		return "بانک آینده"
	case "627381":
		return "بانک انصار"
	case "505785":
		return "بانک ایران زمین"
	case "622106", "627884", "639194":
		return "بانک پارسیان"
	case "502229", "639347":
		return "بانک پاسارگاد"
	case "627760":
		return "پست بانک ایران"
	case "585983", "627353":
		return "بانک تجارت"
	case "502908":
		return "بانک توسعه تعاون"
	case "504172":
		return "بانک رسالت"
	case "207177", "627648":
		return "بانک توسعه صادرات"
	case "636949":
		return "بانک حکمت ایرانیان"
	case "585947":
		return "بانک خاورمیانه"
	case "627412":
		return "بانک اقتصاد نوین"
	case "502938":
		return "بانک دی"
	case "589463":
		return "بانک رفاه کارگران"
	case "621986":
		return "بانک سامان"
	case "589210":
		return "بانک سپه"
	case "639607":
		return "بانک سرمایه"
	case "639346":
		return "بانک سینا"
	case "504706", "502806":
		return "بانک شهر"
	case "603769":
		return "بانک صادرات ایران"
	case "627961":
		return "بانک صنعت و معدن"
	case "639599":
		return "بانک قوامین"
	case "627488", "502910":
		return "بانک کارآفرین"
	case "639217", "603770":
		return "بانک کشاورزی"
	case "505416":
		return "بانک گردشگری"
	case "636795":
		return "بانک مرکزی"
	case "628023":
		return "بانک مسکن"
	case "991975", "610433":
		return "بانک ملت"
	case "603799":
		return "بانک ملی ایران"
	case "606373":
		return "بانک قرض الحسنه مهر ایران"
	case "505801":
		return "موسسه کوثر"
	case "606256":
		return "موسسه اعتباری ملل"
	case "628157":
		return "موسسه اعتباری توسعه"
	case "639370":
		return "بانک مهر اقتصاد"
	default:
		return "-"
	}
}
