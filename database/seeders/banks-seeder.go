package seeders

import (
	"github.com/akposieyefa/open-ai/core/models"
	"github.com/akposieyefa/open-ai/internals"
)

func BankSeeder() {
	banks := []models.Bank{
		{
			Name: "Access Bank",
			Code: "044",
			Ussd: "*901#",
			Logo: "https://nigerianbanks.xyz/logo/access-bank.png",
		},
		{
			Name: "Access Bank (Diamond)",
			Code: "063",
			Ussd: "*426#",
			Logo: "https://nigerianbanks.xyz/logo/access-bank-diamond.png",
		},
		{
			Name: "ALAT by WEMA",
			Code: "035A",
			Ussd: "*945*100#",
			Logo: "https://nigerianbanks.xyz/logo/alat-by-wema.png",
		},
		{
			Name: "ASO Savings and Loans",
			Code: "401",
			Ussd: "",
			Logo: "https://nigerianbanks.xyz/logo/asosavings.png",
		},
		{
			Name: "Bowen Microfinance Bank",
			Code: "50931",
			Ussd: "",
			Logo: "https://nigerianbanks.xyz/logo/default-image.png",
		},
		{
			Name: "CEMCS Microfinance Bank",
			Code: "50823",
			Ussd: "",
			Logo: "https://nigerianbanks.xyz/logo/cemcs-microfinance-bank.png",
		},
		{
			Name: "Citibank Nigeria",
			Code: "023",
			Ussd: "",
			Logo: "https://nigerianbanks.xyz/logo/citibank-nigeria.png",
		},
		{
			Name: "Ecobank Nigeria",
			Code: "050",
			Ussd: "*326#",
			Logo: "https://nigerianbanks.xyz/logo/ecobank-nigeria.png",
		},
		{
			Name: "Ekondo Microfinance Bank",
			Code: "562",
			Ussd: "*540*178#",
			Logo: "https://nigerianbanks.xyz/logo/ekondo-microfinance-bank.png",
		},
		{
			Name: "Fidelity Bank",
			Code: "070",
			Ussd: "*770#",
			Logo: "https://nigerianbanks.xyz/logo/fidelity-bank.png",
		},
		{
			Name: "First Bank of Nigeria",
			Code: "011",
			Ussd: "*894#",
			Logo: "https://nigerianbanks.xyz/logo/first-bank-of-nigeria.png",
		},
		{
			Name: "First City Monument Bank",
			Code: "214",
			Ussd: "*329#",
			Logo: "https://nigerianbanks.xyz/logo/first-city-monument-bank.png",
		},
		{
			Name: "Globus Bank",
			Code: "00103",
			Ussd: "*989#",
			Logo: "https://nigerianbanks.xyz/logo/globus-bank.png",
		},
		{
			Name: "Guaranty Trust Bank",
			Code: "058",
			Ussd: "*737#",
			Logo: "https://nigerianbanks.xyz/logo/guaranty-trust-bank.png",
		},
		{
			Name: "Hasal Microfinance Bank",
			Code: "50383",
			Ussd: "*322*127#",
			Logo: "https://nigerianbanks.xyz/logo/default-image.png",
		},
		{
			Name: "Heritage Bank",
			Code: "030",
			Ussd: "*322#",
			Logo: "https://nigerianbanks.xyz/logo/heritage-bank.png",
		},
		{
			Name: "Jaiz Bank",
			Code: "301",
			Ussd: "*389*301#",
			Logo: "https://nigerianbanks.xyz/logo/default-image.png",
		},
		{
			Name: "Keystone Bank",
			Code: "082",
			Ussd: "*7111#",
			Logo: "https://nigerianbanks.xyz/logo/keystone-bank.png",
		},
		{
			Name: "Kuda Bank",
			Code: "50211",
			Ussd: "",
			Logo: "https://nigerianbanks.xyz/logo/kuda-bank.png",
		},
		{
			Name: "One Finance",
			Code: "565",
			Ussd: "*1303#",
			Logo: "https://nigerianbanks.xyz/logo/default-image.png",
		},
		{
			Name: "Paga",
			Code: "327",
			Ussd: "",
			Logo: "https://nigerianbanks.xyz/logo/paga.png",
		},
		{
			Name: "Parallex Bank",
			Code: "526",
			Ussd: "*322*318*0#",
			Logo: "https://nigerianbanks.xyz/logo/default-image.png",
		},
		{
			Name: "PayCom",
			Code: "100004",
			Ussd: "*955#",
			Logo: "https://nigerianbanks.xyz/logo/default-image.png",
		},
		{
			Name: "Polaris Bank",
			Code: "076",
			Ussd: "*833#",
			Logo: "https://nigerianbanks.xyz/logo/polaris-bank.png",
		},
		{
			Name: "Providus Bank",
			Code: "101",
			Ussd: "",
			Logo: "https://nigerianbanks.xyz/logo/default-image.png",
		},
		{
			Name: "Rubies MFB",
			Code: "125",
			Ussd: "*7797#",
			Logo: "https://nigerianbanks.xyz/logo/default-image.png",
		},
		{
			Name: "Sparkle Microfinance Bank",
			Code: "51310",
			Ussd: "",
			Logo: "https://nigerianbanks.xyz/logo/sparkle-microfinance-bank.png",
		},
		{
			Name: "Stanbic IBTC Bank",
			Code: "221",
			Ussd: "*909#",
			Logo: "https://nigerianbanks.xyz/logo/stanbic-ibtc-bank.png",
		},
		{
			Name: "Standard Chartered Bank",
			Code: "068",
			Ussd: "",
			Logo: "https://nigerianbanks.xyz/logo/standard-chartered-bank.png",
		},
		{
			Name: "Sterling Bank",
			Code: "232",
			Ussd: "*822#",
			Logo: "https://nigerianbanks.xyz/logo/sterling-bank.png",
		},
		{
			Name: "Suntrust Bank",
			Code: "100",
			Ussd: "*5230#",
			Logo: "https://nigerianbanks.xyz/logo/default-image.png",
		},
		{
			Name: "TAJ Bank",
			Code: "302",
			Ussd: "*898#",
			Logo: "https://nigerianbanks.xyz/logo/taj-bank.png",
		},
		{
			Name: "TCF MFB",
			Code: "51211",
			Ussd: "*908#",
			Logo: "https://nigerianbanks.xyz/logo/default-image.png",
		},
		{
			Name: "Titan Trust Bank",
			Code: "102",
			Ussd: "*922#",
			Logo: "https://nigerianbanks.xyz/logo/default-image.png",
		},
		{
			Name: "Union Bank of Nigeria",
			Code: "032",
			Ussd: "*826#",
			Logo: "https://nigerianbanks.xyz/logo/union-bank-of-nigeria.png",
		},
		{
			Name: "United Bank For Africa",
			Code: "033",
			Ussd: "*919#",
			Logo: "https://nigerianbanks.xyz/logo/united-bank-for-africa.png",
		},
		{
			Name: "Unity Bank",
			Code: "215",
			Ussd: "*7799#",
			Logo: "https://nigerianbanks.xyz/logo/default-image.png",
		},
		{
			Name: "VFD",
			Code: "566",
			Ussd: "",
			Logo: "https://nigerianbanks.xyz/logo/default-image.png",
		},
		{
			Name: "Wema Bank",
			Code: "035",
			Ussd: "*945#",
			Logo: "https://nigerianbanks.xyz/logo/wema-bank.png",
		},
		{
			Name: "Zenith Bank",
			Code: "057",
			Ussd: "*966#",
			Logo: "https://nigerianbanks.xyz/logo/zenith-bank.png",
		},
	}
	internals.DB.Create(&banks)
}
