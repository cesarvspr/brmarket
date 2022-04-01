package brmarket

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type B3IndexInfo struct {
	BizSts         BizSts           `json:"BizSts"`
	Index          Index            `json:"Index"`
	UnderlyingList []UnderlyingList `json:"UnderlyingList"`
	Msg            Msg              `json:"Msg"`
}
type BizSts struct {
	Cd string `json:"cd"`
}
type Index struct {
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}
type UnderlyingList struct {
	IndexTheoreticalQty float64 `json:"indexTheoreticalQty"`
	IndxCmpnPctg        float64 `json:"indxCmpnPctg"`
	Symb                string  `json:"symb"`
	Desc                string  `json:"desc"`
}
type Msg struct {
	DtTm string `json:"dtTm"`
}

func GetSharesForIndex(index string) (B3IndexInfo, error) {
	mapIndexName := map[string]string{
		"IBOV": "IBOV",
		"IBXX": "IBXX",
		"IBXL": "IBXL",
		"IBRA": "IBRA",
		"IGCX": "IGCX",
		"ITAG": "ITAG",
		"IGNM": "IGNM",
		"IGCT": "IGCT",
		"IDIV": "IDIV",
		"MLCX": "MLCX",
		"IVBX": "IVBX",
		"ICO2": "ICO2",
		"ISEE": "ISEE",
		"ICON": "ICON",
		"IEEX": "IEEX",
		"IFNC": "IFNC",
		"IMOB": "IMOB",
		"INDX": "INDX",
		"IMAT": "IMAT",
		"UTIL": "UTIL",
		"IFIX": "IFIX",
		"BDRX": "BDRX",
	}

	if _, ok := mapIndexName[index]; !ok {
		return B3IndexInfo{}, fmt.Errorf("Index %s not found", index)
	}

	URL := "https://cotacao.b3.com.br/mds/api/v1/IndexComposition/" + index
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Errorf(fmt.Sprintf("Error when loading data from B3 %+v", err))
		return B3IndexInfo{}, err
	}

	//Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf(string(err.Error()))
		return B3IndexInfo{}, err
	}
	var info B3IndexInfo
	if err := json.Unmarshal(data, &info); err != nil {
		fmt.Errorf(string(err.Error()))
		return info, err
	}

	return info, nil
}
