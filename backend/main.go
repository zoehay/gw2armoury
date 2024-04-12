package main

type item struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Icon string `"json:icon"`
    Description string `"json:description"`
}

var items = []item{
    {ID: 28445, Name: "Strong Soft Wood Longbow of Fire", Icon: "https://render.guildwars2.com/file/C6110F52DF5AFE0F00A56F9E143E9732176DDDE9/65015.png", Description: ""},
    {ID:12452, Name: "Omnomberry Bar", Icon: "https://render.guildwars2.com/file/6BD5B65FBC6ED450219EC86DD570E59F4DA3791F/433643.png", Description: ""},
}

