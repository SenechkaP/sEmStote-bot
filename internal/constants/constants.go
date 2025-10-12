package constants

const (
	LoadEnvErrorOutput      = "could not load env file %s: %v"
	EnvVariablesErrorOutput = "something wrong with .env file. It should include following variables:\n  TELEGRAM_BOT_TOKEN\n  COMMISSION_FOR_SNEAKERS\n  COMMISSION_FOR_SHIRTS\n  CNY_DEFAULT_RATE\n  EUR_DEFAULT_RATE"
	Greeting                = "Доброго времени суток! Чем я могу Вам помочь?"
	AdminAccount            = "https://t.me/s3mmm_7"
	StepsToMakeOrder        = "Для заказа отправьте @s3mmm_7 в лс:\n\n1️⃣Ссылку на товар\n2️⃣Укажите размер товара\n3️⃣При необходимости цвет товара"
	PoizonGuide             = "https://telegra.ph/Kak-poschitat-konechnuyu-cenu-za-tovar-i-zakazat-02-28"
	WayToLink               = "На странице выбранной Вами вещи нажимаем на кнопку поделиться (вместо нее иногда бывает зеленый значок) справа вверху, затем на кнопку скопировать ссылку во всплывающем окне"
	ChooseItemType          = "Выберите тип вещи"
	ChooseShoesType         = "Выберите тип обуви"
	ChooseClothesType       = "Выберите тип одежды"
	ChooseAccessoriesType   = "Выберите тип аксессуара"
	OtherItemTypeText       = "❗️В этой категории вы можете посчитать стоимость товара БЕЗ учета доставки до России"
	EnterPrice              = "Напишите стоимость товара в юанях\n\n❗️<i>Минимальная стоимость товара для заказа 20¥</i>"
	PriceOutput             = "Вы выбрали: <b>%s</b>\n\nСтоимость вашего товара: <b>%s ¥</b>\n\nИтоговая стоимость без учёта доставки до России: <b>%s ₽</b> ✅\n\n❗️<i>При заказе, к стоимости будет добавлена цена за доставку по России до вашего ПВЗ Boxberry</i>"
	RateOutput              = "Курс на сегодня: %.2f₽ за 1¥"
	RateErrorOutput         = "Не удалось получить информацию о текущем курсе юаня к рублю. Обратитесь к @s3mmm_7 для уточнения курса на сегодня."
)
