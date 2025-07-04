package main

import (
	"math/rand"
	"time"
)

var (
	// Order fields
	TrackNumbers = []string{
		"TRK1234567890", "TRK0987654321", "TRK1122334455", "TRK5566778899", "TRK6677889900",
		"TRK1029384756", "TRK5647382910", "TRK0192837465", "TRK1928374650", "TRK9182736450",
	}
	Entries = []string{
		"WEB", "APP", "MOBILE", "CALLCENTER", "PARTNER",
		"API", "KIOSK", "STORE", "MARKETPLACE", "SOCIAL",
	}
	Locales = []string{
		"ru", "en", "fr", "de", "es",
		"it", "zh", "ja", "pt", "tr",
	}
	InternalSignatures = []string{
		"sig1", "sig2", "sig3", "sig4", "sig5",
		"sig6", "sig7", "sig8", "sig9", "sig10",
	}
	CustomerIds = []string{
		"cust001", "cust002", "cust003", "cust004", "cust005",
		"cust006", "cust007", "cust008", "cust009", "cust010",
	}
	DeliveryServices = []string{
		"DHL", "FedEx", "UPS", "Boxberry", "CDEK",
		"PonyExpress", "RussianPost", "DPD", "Hermes", "YandexDelivery",
	}
	ShardKeys = []string{
		"shardA", "shardB", "shardC", "shardD", "shardE",
		"shardF", "shardG", "shardH", "shardI", "shardJ",
	}
	SmIds = []int{
		101, 102, 103, 104, 105,
		106, 107, 108, 109, 110,
	}
	OofShards = []string{
		"oof1", "oof2", "oof3", "oof4", "oof5",
		"oof6", "oof7", "oof8", "oof9", "oof10",
	}
	DateCreateds = []time.Time{
		time.Date(2025, 6, 1, 12, 0, 0, 0, time.UTC),
		time.Date(2025, 6, 2, 13, 0, 0, 0, time.UTC),
		time.Date(2025, 6, 3, 14, 0, 0, 0, time.UTC),
		time.Date(2025, 6, 4, 15, 0, 0, 0, time.UTC),
		time.Date(2025, 6, 5, 16, 0, 0, 0, time.UTC),
		time.Date(2025, 6, 6, 17, 0, 0, 0, time.UTC),
		time.Date(2025, 6, 7, 18, 0, 0, 0, time.UTC),
		time.Date(2025, 6, 8, 19, 0, 0, 0, time.UTC),
		time.Date(2025, 6, 9, 20, 0, 0, 0, time.UTC),
		time.Date(2025, 6, 10, 21, 0, 0, 0, time.UTC),
	}

	// Delivery fields
	DeliveryNames = []string{
		"Иван Иванов", "Мария Петрова", "John Smith", "Anna Müller", "Li Wei",
		"Игорь Кузнецов", "Sara Rossi", "Tomáš Novák", "Emma Dubois", "Juan Pérez",
	}
	DeliveryPhones = []string{
		"+79001112233", "+79002223344", "+79003334455", "+79004445566", "+79005556677",
		"+79006667788", "+79007778899", "+79008889900", "+79009990011", "+79001001122",
	}
	DeliveryZips = []string{
		"101000", "102000", "103000", "104000", "105000",
		"106000", "107000", "108000", "109000", "110000",
	}
	DeliveryCities = []string{
		"Москва", "Санкт-Петербург", "Новосибирск", "Екатеринбург", "Казань",
		"Лондон", "Париж", "Берлин", "Рим", "Пекин",
	}
	DeliveryAddresses = []string{
		"ул. Ленина, 1", "пр. Мира, 10", "ул. Пушкина, 5", "ул. Садовая, 15", "ул. Советская, 20",
		"Baker St. 221B", "Champs-Élysées 10", "Alexanderplatz 3", "Via Roma 8", "Nanjing Rd. 100",
	}
	DeliveryRegions = []string{
		"Московская область", "Ленинградская область", "Новосибирская область", "Свердловская область", "Татарстан",
		"Greater London", "Île-de-France", "Berlin", "Lazio", "Beijing",
	}
	DeliveryEmails = []string{
		"ivanov@mail.ru", "petrova@yandex.ru", "smith@gmail.com", "mueller@web.de", "liwei@qq.com",
		"kuznetsov@mail.ru", "rossi@libero.it", "novak@seznam.cz", "dubois@orange.fr", "perez@gmail.com",
	}

	// Payment fields
	Transactions = []string{
		"txn001", "txn002", "txn003", "txn004", "txn005",
		"txn006", "txn007", "txn008", "txn009", "txn010",
	}
	RequestIds = []string{
		"req001", "req002", "req003", "req004", "req005",
		"req006", "req007", "req008", "req009", "req010",
	}
	Currencies = []string{
		"RUB", "USD", "EUR", "CNY", "GBP",
		"JPY", "KZT", "UAH", "BYN", "TRY",
	}
	Providers = []string{
		"Sberbank", "Tinkoff", "AlfaBank", "PayPal", "Stripe",
		"Qiwi", "YooMoney", "WebMoney", "Visa", "MasterCard",
	}
	Amounts = []int{
		1000, 2500, 3500, 4999, 1200,
		880, 15000, 6000, 750, 3200,
	}
	PaymentDts = []int64{
		1720000000, 1720003600, 1720007200, 1720010800, 1720014400,
		1720018000, 1720021600, 1720025200, 1720028800, 1720032400,
	}
	Banks = []string{
		"Сбербанк", "Тинькофф", "ВТБ", "Россельхозбанк", "Газпромбанк",
		"Raiffeisen", "UniCredit", "Открытие", "Росбанк", "Промсвязьбанк",
	}
	DeliveryCosts = []int{
		200, 300, 150, 400, 250,
		100, 350, 500, 180, 220,
	}
	GoodsTotals = []int{
		800, 2200, 3200, 4599, 950,
		780, 14000, 5700, 570, 2980,
	}
	CustomFees = []int{
		0, 50, 100, 75, 20,
		0, 30, 0, 10, 0,
	}

	// Item fields
	ChrtIds = []int64{
		123456, 234567, 345678, 456789, 567890,
		678901, 789012, 890123, 901234, 123457,
	}
	ItemTrackNumbers = []string{
		"ITM123456", "ITM234567", "ITM345678", "ITM456789", "ITM567890",
		"ITM678901", "ITM789012", "ITM890123", "ITM901234", "ITM012345",
	}
	Prices = []int{
		500, 1200, 800, 1500, 300,
		700, 2000, 1000, 450, 900,
	}
	Rids = []string{
		"RID001", "RID002", "RID003", "RID004", "RID005",
		"RID006", "RID007", "RID008", "RID009", "RID010",
	}
	ItemNames = []string{
		"Кроссовки", "Футболка", "Джинсы", "Пальто", "Сумка",
		"Часы", "Очки", "Платье", "Куртка", "Рюкзак",
	}
	Sales = []int{
		0, 10, 20, 15, 5,
		25, 30, 0, 12, 18,
	}
	Sizes = []string{
		"S", "M", "L", "XL", "XXL",
		"40", "42", "44", "46", "48",
	}

	NmIds = []int64{
		111111, 222222, 333333, 444444, 555555,
		666666, 777777, 888888, 999999, 101010,
	}
	Brands = []string{
		"Nike", "Adidas", "Puma", "Reebok", "Levi's",
		"Gucci", "Ray-Ban", "Zara", "Columbia", "Eastpak",
	}
	Statuses = []int{
		200, 201, 400, 401, 404,
	}
)

func GetRandomElement[T any](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))

	randomIndex := rand.Intn(len(s))
	return s[randomIndex]
}

func GenerateNewOrder() Order {
	var order Order
	var delivery Delivery
	var item1, item2, item3, item4, item5 Item
	var payment Payment

	delivery.Name = GetRandomElement(DeliveryNames)
	delivery.Phone = GetRandomElement(DeliveryPhones)
	delivery.Zip = GetRandomElement(DeliveryZips)
	delivery.City = GetRandomElement(DeliveryCities)
	delivery.Address = GetRandomElement(DeliveryAddresses)
	delivery.Region = GetRandomElement(DeliveryRegions)
	delivery.Email = GetRandomElement(DeliveryEmails)

	item1.Name = GetRandomElement(ItemNames)
	item1.Brand = GetRandomElement(Brands)
	item1.ChrtId = GetRandomElement(ChrtIds)
	item1.NmId = GetRandomElement(NmIds)
	item1.Price = GetRandomElement(Prices)
	item1.Rid = GetRandomElement(Rids)
	item1.Sale = GetRandomElement(Sales)
	item1.TotalPrice = item1.Price - (item1.Price * item1.Sale / 100)
	item1.Size = GetRandomElement(Sizes)
	item1.Status = GetRandomElement(Statuses)
	item1.TrackNumber = GetRandomElement(ItemTrackNumbers)

	item2.Name = GetRandomElement(ItemNames)
	item2.Brand = GetRandomElement(Brands)
	item2.ChrtId = GetRandomElement(ChrtIds)
	item2.NmId = GetRandomElement(NmIds)
	item2.Price = GetRandomElement(Prices)
	item2.Rid = GetRandomElement(Rids)
	item2.Sale = GetRandomElement(Sales)
	item2.TotalPrice = item2.Price - (item2.Price * item2.Sale / 100)
	item2.Size = GetRandomElement(Sizes)
	item2.Status = GetRandomElement(Statuses)
	item2.TrackNumber = GetRandomElement(ItemTrackNumbers)

	item3.Name = GetRandomElement(ItemNames)
	item3.Brand = GetRandomElement(Brands)
	item3.ChrtId = GetRandomElement(ChrtIds)
	item3.NmId = GetRandomElement(NmIds)
	item3.Price = GetRandomElement(Prices)
	item3.Rid = GetRandomElement(Rids)
	item3.Sale = GetRandomElement(Sales)
	item3.TotalPrice = item3.Price - (item3.Price * item3.Sale / 100)
	item3.Size = GetRandomElement(Sizes)
	item3.Status = GetRandomElement(Statuses)
	item3.TrackNumber = GetRandomElement(ItemTrackNumbers)

	item4.Name = GetRandomElement(ItemNames)
	item4.Brand = GetRandomElement(Brands)
	item4.ChrtId = GetRandomElement(ChrtIds)
	item4.NmId = GetRandomElement(NmIds)
	item4.Price = GetRandomElement(Prices)
	item4.Rid = GetRandomElement(Rids)
	item4.Sale = GetRandomElement(Sales)
	item4.TotalPrice = item4.Price - (item4.Price * item4.Sale / 100)
	item4.Size = GetRandomElement(Sizes)
	item4.Status = GetRandomElement(Statuses)
	item4.TrackNumber = GetRandomElement(ItemTrackNumbers)

	item5.Name = GetRandomElement(ItemNames)
	item5.Brand = GetRandomElement(Brands)
	item5.ChrtId = GetRandomElement(ChrtIds)
	item5.NmId = GetRandomElement(NmIds)
	item5.Price = GetRandomElement(Prices)
	item5.Rid = GetRandomElement(Rids)
	item5.Sale = GetRandomElement(Sales)
	item5.TotalPrice = item1.Price - (item1.Price * item1.Sale / 100)
	item5.Size = GetRandomElement(Sizes)
	item5.Status = GetRandomElement(Statuses)
	item5.TrackNumber = GetRandomElement(ItemTrackNumbers)

	payment.Amount = GetRandomElement(Amounts)
	payment.Bank = GetRandomElement(Banks)
	payment.Currency = GetRandomElement(Currencies)
	payment.CustomFee = GetRandomElement(CustomFees)
	payment.DeliveryCost = GetRandomElement(DeliveryCosts)
	payment.GoodsTotal = GetRandomElement(GoodsTotals)
	payment.PaymentDt = GetRandomElement(PaymentDts)
	payment.Provider = GetRandomElement(Providers)
	payment.Transaction = GetRandomElement(Transactions)
	payment.RequestId = GetRandomElement(RequestIds)

	order.CustomerId = GetRandomElement(CustomerIds)
	order.DateCreated = GetRandomElement(DateCreateds)
	order.Delivery = delivery
	order.DeliveryService = GetRandomElement(DeliveryServices)
	order.Entry = GetRandomElement(Entries)
	order.InternalSignature = GetRandomElement(InternalSignatures)
	order.Items = []Item{item1, item2, item3, item4, item5}
	order.Locale = GetRandomElement(Locales)
	order.OofShard = GetRandomElement(OofShards)
	order.ShardKey = GetRandomElement(ShardKeys)
	order.Payment = payment
	order.TrackNumber = GetRandomElement(TrackNumbers)

	return order
}
