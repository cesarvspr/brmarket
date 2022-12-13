package brmarket

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	IBOV = "IBOV"
	IBXX = "IBXX"
	IBXL = "IBXL"
	IBRA = "IBRA"
	IGCX = "IGCX"
	ITAG = "ITAG"
	IGNM = "IGNM"
	IGCT = "IGCT"
	IDIV = "IDIV"
	MLCX = "MLCX"
	IVBX = "IVBX"
	ICO2 = "ICO2"
	ISEE = "ISEE"
	ICON = "ICON"
	IEEX = "IEEX"
	IFNC = "IFNC"
	IMOB = "IMOB"
	INDX = "INDX"
	IMAT = "IMAT"
	UTIL = "UTIL"
	IFIX = "IFIX"
	BDRX = "BDRX"
)

// Response type for B3 API
type B3IndexInfo struct {
	BizSts         bizSts           `json:"BizSts"`
	Index          index            `json:"Index"`
	UnderlyingList []underlyingList `json:"UnderlyingList"`
	Msg            msg              `json:"Msg"`
}
type bizSts struct {
	Cd string `json:"cd"`
}
type index struct {
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}
type underlyingList struct {
	IndexTheoreticalQty float64 `json:"indexTheoreticalQty"`
	IndxCmpnPctg        float64 `json:"indxCmpnPctg"`
	Symb                string  `json:"symb"`
	Desc                string  `json:"desc"`
}
type msg struct {
	DtTm string `json:"dtTm"`
}

// Return shares for reach index, only Brazilian indexes allowed.
// Example: GetSharesForIndex("IBOV") or GetSharesForIndex(brmarket.IBOV)
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
		return B3IndexInfo{}, fmt.Errorf("index %s not found", index)
	}

	URL := "https://cotacao.b3.com.br/mds/api/v1/IndexComposition/" + index
	resp, err := http.Get(URL)
	if err != nil {
		return B3IndexInfo{}, err
	}
	defer resp.Body.Close()

	//Read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return B3IndexInfo{}, err
	}
	var info B3IndexInfo
	if err := json.Unmarshal(data, &info); err != nil {
		return info, err
	}

	return info, nil
}
